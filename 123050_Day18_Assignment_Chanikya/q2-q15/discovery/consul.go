package discovery

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const defaultConsulAddress = "http://127.0.0.1:8500"

var consulClient = &http.Client{
	Timeout: 3 * time.Second,
}

type serviceInstance struct {
	Service struct {
		Address string `json:"Address"`
		Port    int    `json:"Port"`
	} `json:"Service"`

	Node struct {
		Address string `json:"Address"`
	} `json:"Node"`
}

// Lookup discovers a healthy service instance from Consul.
func Lookup(serviceName string) (*url.URL, error) {
	consulAddress := strings.TrimRight(
		os.Getenv("CONSUL_HTTP_ADDR"),
		"/",
	)

	if consulAddress == "" {
		consulAddress = defaultConsulAddress
	}

	endpoint := fmt.Sprintf(
		"%s/v1/health/service/%s?passing=true",
		consulAddress,
		url.PathEscape(serviceName),
	)

	resp, err := consulClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to contact Consul: %w",
			err,
		)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"Consul returned HTTP status %d",
			resp.StatusCode,
		)
	}

	var instances []serviceInstance

	if err := json.NewDecoder(resp.Body).Decode(&instances); err != nil {
		return nil, fmt.Errorf(
			"failed to decode Consul response: %w",
			err,
		)
	}

	if len(instances) == 0 {
		return nil, fmt.Errorf(
			"no healthy instances found for service %q",
			serviceName,
		)
	}

	// Use the first healthy instance returned by Consul.
	instance := instances[0]

	address := strings.TrimSpace(instance.Service.Address)

	if address == "" {
		address = strings.TrimSpace(instance.Node.Address)
	}

	if address == "" || instance.Service.Port < 1 {
		return nil, fmt.Errorf(
			"invalid address or port for service %q",
			serviceName,
		)
	}

	target := &url.URL{
		Scheme: "http",
		Host: net.JoinHostPort(
			address,
			fmt.Sprintf("%d", instance.Service.Port),
		),
	}

	return target, nil
}

// Proxy discovers a healthy service instance and forwards the request.
// stripPrefix is removed from the incoming path before forwarding.
func Proxy(serviceName, stripPrefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		target, err := Lookup(serviceName)

		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error":   "service unavailable",
				"service": serviceName,
			})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(target)

		originalDirector := proxy.Director

		proxy.Director = func(req *http.Request) {
			originalDirector(req)

			path := req.URL.Path

			if stripPrefix != "" &&
				strings.HasPrefix(path, stripPrefix) {
				path = strings.TrimPrefix(path, stripPrefix)
			}

			if path == "" {
				path = "/"
			}

			if !strings.HasPrefix(path, "/") {
				path = "/" + path
			}

			req.URL.Path = path
			req.Host = target.Host
		}

		proxy.ErrorHandler = func(
			w http.ResponseWriter,
			req *http.Request,
			proxyErr error,
		) {
			w.Header().Set(
				"Content-Type",
				"application/json",
			)
			w.WriteHeader(http.StatusBadGateway)

			_, _ = w.Write([]byte(
				`{"error":"upstream service request failed"}`,
			))
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
