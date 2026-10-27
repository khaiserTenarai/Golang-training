package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"order-management/services/order/model"
)

type GORMOrderRepository struct{ db *gorm.DB }

func NewOrderRepository(db *gorm.DB) OrderRepository { return &GORMOrderRepository{db: db} }
func (r *GORMOrderRepository) Create(ctx context.Context, order *model.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}
func (r *GORMOrderRepository) GetByID(ctx context.Context, id uint) (*model.Order, error) {
	var order model.Order
	if err := r.db.WithContext(ctx).First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
func (r *GORMOrderRepository) GetAll(ctx context.Context) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.WithContext(ctx).Order("id DESC").Limit(500).Find(&orders).Error
	return orders, err
}
func (r *GORMOrderRepository) Update(ctx context.Context, order *model.Order) error {
	return r.db.WithContext(ctx).Save(order).Error
}

// SetStatus uses a conditional update to avoid concurrent status overwrites.
func (r *GORMOrderRepository) SetStatus(ctx context.Context, id uint, newStatus string, allowedFrom []string) (bool, error) {
	result := r.db.WithContext(ctx).Model(&model.Order{}).Where("id = ? AND status IN ?", id, allowedFrom).Updates(map[string]interface{}{"status": newStatus})
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected > 0 {
		return true, nil
	}
	var existing model.Order
	if err := r.db.WithContext(ctx).First(&existing, id).Error; err != nil {
		return false, err
	}
	if existing.Status == newStatus {
		return false, nil
	}
	return false, fmt.Errorf("cannot change order status from %s to %s", existing.Status, newStatus)
}
func (r *GORMOrderRepository) MarkStockReleased(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&model.Order{}).Where("id = ?", id).Update("stock_released", true).Error
}
