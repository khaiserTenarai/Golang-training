package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"order-management/services/order/client"
	"order-management/services/order/model"
	"order-management/services/order/repository"

	"github.com/google/uuid"
)

type orderService struct {
	repo          repository.OrderRepository
	productClient *client.ProductClient
}

func NewOrderService(repo repository.OrderRepository, productClient *client.ProductClient) OrderService {
	return &orderService{repo: repo, productClient: productClient}
}

func (s *orderService) Create(ctx context.Context, customerID, productID uint, quantity int) (*model.Order, error) {
	if customerID == 0 || productID == 0 || quantity <= 0 {
		return nil, errors.New("customerId, productId and positive quantity are required")
	}
	// Product Service checks and reserves stock atomically and returns the price snapshot.
	// This avoids a separate GET request and prevents a stale stock check from being trusted.
	stockReference := uuid.New().String()
	product, err := s.productClient.ReserveStock(ctx, productID, stockReference, quantity)
	if err != nil {
		return nil, fmt.Errorf("reserve product stock: %w", err)
	}

	order := &model.Order{
		CustomerID: customerID, ProductID: productID, Quantity: quantity,
		TotalAmount: product.Price * float64(quantity), Status: "CREATED",
		StockReference: stockReference, StockReleased: false,
	}
	if err := s.repo.Create(ctx, order); err != nil {
		// Compensation uses a separate short context so a canceled HTTP request
		// does not prevent stock restoration after a database failure.
		compensationCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if releaseErr := s.productClient.ReleaseStock(compensationCtx, productID, stockReference, quantity); releaseErr != nil {
			return nil, fmt.Errorf("create order failed (%v) and stock compensation failed (%w); reservation reference=%s", err, releaseErr, stockReference)
		}
		return nil, fmt.Errorf("create order: %w", err)
	}
	return order, nil
}

func (s *orderService) GetByID(ctx context.Context, id uint) (*model.Order, error) {
	if id == 0 {
		return nil, errors.New("invalid order ID")
	}
	return s.repo.GetByID(ctx, id)
}
func (s *orderService) GetAll(ctx context.Context) ([]model.Order, error) { return s.repo.GetAll(ctx) }

func (s *orderService) UpdateStatus(ctx context.Context, id uint, status string) (*model.Order, error) {
	if id == 0 {
		return nil, errors.New("invalid order ID")
	}
	status = strings.ToUpper(strings.TrimSpace(status))
	if status != "CONFIRMED" && status != "CANCELLED" {
		return nil, errors.New("status must be CONFIRMED or CANCELLED")
	}

	if status == "CONFIRMED" {
		_, err := s.repo.SetStatus(ctx, id, "CONFIRMED", []string{"CREATED"})
		if err != nil {
			return nil, err
		}
		return s.repo.GetByID(ctx, id)
	}

	// Persist the cancellation first. If stock release fails, StockReleased remains
	// false and repeating this same cancellation request safely retries the release.
	if _, err := s.repo.SetStatus(ctx, id, "CANCELLED", []string{"CREATED", "CONFIRMED"}); err != nil {
		return nil, err
	}
	order, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if !order.StockReleased {
		// Orders created before inventory tracking have no reservation to restore.
		if order.StockReference == "" {
			if err := s.repo.MarkStockReleased(ctx, id); err != nil {
				return nil, err
			}
			order.StockReleased = true
		} else if err := s.productClient.ReleaseStock(ctx, order.ProductID, order.StockReference, order.Quantity); err != nil {
			return nil, fmt.Errorf("order is cancelled but stock release is pending; retry cancellation: %w", err)
		}
		if order.StockReference != "" {
			if err := s.repo.MarkStockReleased(ctx, id); err != nil {
				return nil, fmt.Errorf("stock released but failed to record release status; retry cancellation safely: %w", err)
			}
			order.StockReleased = true
		}
	}
	order.Status = "CANCELLED"
	return order, nil
}
