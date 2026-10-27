package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	consulapi "github.com/hashicorp/consul/api"
	"order-management/config"
	"order-management/shared/discovery"
	"order-management/shared/middleware"
)

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

type ProductClient struct {
	httpClient    *http.Client
	balancer      *discovery.RoundRobin
	internalToken string
	timeout       time.Duration
	mu            sync.Mutex
	failures      int
	openUntil     time.Time
}

func NewProductClient() *ProductClient {
	fallback := config.Get("PRODUCT_SERVICE_URL", "http://localhost:8081")
	var consulClient *consulapi.Client
	if cc, err := discovery.NewConsulClient(); err == nil {
		consulClient = cc
	}
	timeout := config.GetDuration("PRODUCT_REQUEST_TIMEOUT", 2*time.Second)
	return &ProductClient{
		httpClient:    &http.Client{Timeout: timeout},
		balancer:      discovery.NewRoundRobin(consulClient, config.Get("PRODUCT_SERVICE_NAME", "product-service"), fallback),
		internalToken: os.Getenv("INTERNAL_SERVICE_TOKEN"),
		timeout:       timeout,
	}
}

func (c *ProductClient) GetProduct(ctx context.Context, id uint) (*Product, error) {
	var product Product
	if err := c.callJSON(ctx, http.MethodGet, fmt.Sprintf("/products/%d", id), nil, &product, false); err != nil {
		return nil, err
	}
	if product.ID == 0 || product.Price <= 0 {
		return nil, errors.New("product service returned invalid product data")
	}
	return &product, nil
}

func (c *ProductClient) ReserveStock(ctx context.Context, id uint, reference string, quantity int) (*Product, error) {
	body := map[string]interface{}{"orderReference": reference, "quantity": quantity}
	var product Product
	if err := c.callJSON(ctx, http.MethodPost, fmt.Sprintf("/internal/products/%d/reserve", id), body, &product, true); err != nil {
		return nil, err
	}
	if product.ID == 0 || product.Price <= 0 {
		return nil, errors.New("product service returned invalid product data after stock reservation")
	}
	return &product, nil
}

func (c *ProductClient) ReleaseStock(ctx context.Context, id uint, reference string, quantity int) error {
	body := map[string]interface{}{"orderReference": reference, "quantity": quantity}
	return c.callJSON(ctx, http.MethodPost, fmt.Sprintf("/internal/products/%d/release", id), body, nil, true)
}

func (c *ProductClient) callJSON(ctx context.Context, method, path string, body interface{}, output interface{}, internal bool) error {
	c.mu.Lock()
	if time.Now().Before(c.openUntil) {
		until := c.openUntil
		c.mu.Unlock()
		return fmt.Errorf("product service circuit breaker is open until %s", until.Format(time.RFC3339))
	}
	c.mu.Unlock()

	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		lastErr = nil
		if err := ctx.Err(); err != nil {
			return err
		}
		base, err := c.balancer.Next()
		if err != nil {
			lastErr = err
		} else {
			target, parseErr := url.Parse(base.String() + path)
			if parseErr != nil {
				lastErr = parseErr
			} else {
				requestCtx, cancel := context.WithTimeout(ctx, c.timeout)
				var data []byte
				if body != nil {
					data, err = json.Marshal(body)
					if err != nil {
						cancel()
						return err
					}
				}
				req, reqErr := http.NewRequestWithContext(requestCtx, method, target.String(), bytes.NewReader(data))
				if reqErr != nil {
					cancel()
					return reqErr
				}
				req.Header.Set("Accept", "application/json")
				if requestID := middleware.RequestIDFromContext(ctx); requestID != "" {
					req.Header.Set("X-Request-ID", requestID)
				}
				if body != nil {
					req.Header.Set("Content-Type", "application/json")
				}
				if internal {
					req.Header.Set("X-Internal-Token", c.internalToken)
				}
				resp, callErr := c.httpClient.Do(req)
				if callErr != nil {
					lastErr = callErr
				} else {
					responseBody, readErr := io.ReadAll(io.LimitReader(resp.Body, 2<<20))
					_ = resp.Body.Close()
					if readErr != nil {
						lastErr = readErr
					} else if resp.StatusCode >= 500 || resp.StatusCode == http.StatusTooManyRequests {
						lastErr = fmt.Errorf("product service returned HTTP %d: %s", resp.StatusCode, string(responseBody))
					} else if resp.StatusCode < 200 || resp.StatusCode >= 300 {
						cancel()
						c.recordResult(false, false)
						return fmt.Errorf("product service returned HTTP %d: %s", resp.StatusCode, string(responseBody))
					} else if output != nil {
						if err := json.Unmarshal(responseBody, output); err != nil {
							cancel()
							c.recordResult(false, true)
							return fmt.Errorf("decode product service response: %w", err)
						}
					} else if len(responseBody) > 0 {
						var ignored interface{}
						_ = json.Unmarshal(responseBody, &ignored)
					}
					if lastErr == nil {
						cancel()
						c.recordResult(true, false)
						return nil
					}
				}
				cancel()
			}
		}
		if attempt < 2 {
			delay := time.Duration(100*(1<<attempt)) * time.Millisecond
			timer := time.NewTimer(delay)
			select {
			case <-ctx.Done():
				timer.Stop()
				return ctx.Err()
			case <-timer.C:
			}
		}
	}
	c.recordResult(false, true)
	return fmt.Errorf("product service request failed after retries: %w", lastErr)
}

func (c *ProductClient) recordResult(success, transient bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if success {
		c.failures = 0
		c.openUntil = time.Time{}
		return
	}
	if transient {
		c.failures++
		if c.failures >= 5 {
			c.openUntil = time.Now().Add(10 * time.Second)
			c.failures = 0
		}
	}
}
