package repository

import (
	"context"
	"errors"
	"fmt"

	"order-management/services/product/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORMProductRepository implements ProductRepository using GORM.
type GORMProductRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a product repository.
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &GORMProductRepository{
		db: db,
	}
}

// Create inserts a new product into PostgreSQL.
func (r *GORMProductRepository) Create(
	ctx context.Context,
	product *model.Product,
) error {
	if err := r.db.WithContext(ctx).Create(product).Error; err != nil {
		return fmt.Errorf("create product: %w", err)
	}

	return nil
}

// GetByID retrieves a product by its ID.
func (r *GORMProductRepository) GetByID(
	ctx context.Context,
	id uint,
) (*model.Product, error) {
	var product model.Product

	err := r.db.WithContext(ctx).First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, fmt.Errorf("get product by ID: %w", err)
	}

	return &product, nil
}

// GetAll retrieves all products.
func (r *GORMProductRepository) GetAll(
	ctx context.Context,
) ([]model.Product, error) {
	var products []model.Product

	if err := r.db.WithContext(ctx).Order("id ASC").Find(&products).Error; err != nil {
		return nil, fmt.Errorf("get all products: %w", err)
	}

	return products, nil
}

// Update saves changes to an existing product.
func (r *GORMProductRepository) Update(
	ctx context.Context,
	product *model.Product,
) error {
	result := r.db.WithContext(ctx).Model(&model.Product{}).
		Where("id = ?", product.ID).
		Updates(map[string]interface{}{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"stock":       product.Stock,
		})

	if result.Error != nil {
		return fmt.Errorf("update product: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// Delete removes a product by its ID.
func (r *GORMProductRepository) Delete(
	ctx context.Context,
	id uint,
) error {
	result := r.db.WithContext(ctx).Delete(&model.Product{}, id)

	if result.Error != nil {
		return fmt.Errorf("delete product: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// ReserveStock atomically subtracts stock and returns the product snapshot used to price the order.
func (r *GORMProductRepository) ReserveStock(ctx context.Context, id uint, ref string, quantity int) (*model.Product, error) {
	if id == 0 || ref == "" || quantity <= 0 {
		return nil, fmt.Errorf("product ID, order reference and positive quantity are required")
	}
	var reservedProduct *model.Product
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var product model.Product
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, id).Error; err != nil {
			return err
		}
		var existing model.StockOperation
		err := tx.Where("order_reference = ? AND action = ?", ref, "RESERVE").First(&existing).Error
		if err == nil {
			if existing.ProductID == id && existing.Quantity == quantity {
				if existing.UnitPrice > 0 {
					product.Price = existing.UnitPrice
				}
				reservedProduct = &product
				return nil
			}
			return fmt.Errorf("order reference already used for a different reservation")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		var released model.StockOperation
		if err := tx.Where("order_reference = ? AND action = ?", ref, "RELEASE").First(&released).Error; err == nil {
			return fmt.Errorf("reservation was already released")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if product.Stock < quantity {
			return fmt.Errorf("insufficient stock: available %d, requested %d", product.Stock, quantity)
		}
		if err := tx.Model(&model.Product{}).Where("id = ? AND stock >= ?", id, quantity).Update("stock", gorm.Expr("stock - ?", quantity)).Error; err != nil {
			return err
		}
		if err := tx.Create(&model.StockOperation{OrderReference: ref, Action: "RESERVE", ProductID: id, Quantity: quantity, UnitPrice: product.Price}).Error; err != nil {
			return err
		}
		product.Stock -= quantity
		reservedProduct = &product
		return nil
	})
	if err != nil {
		return nil, err
	}
	return reservedProduct, nil
}

// ReleaseStock restores a reservation once. Repeated release requests are idempotent.
func (r *GORMProductRepository) ReleaseStock(ctx context.Context, id uint, ref string, quantity int) error {
	if id == 0 || ref == "" || quantity <= 0 {
		return fmt.Errorf("product ID, order reference and positive quantity are required")
	}
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var product model.Product
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, id).Error; err != nil {
			return err
		}
		var reserved model.StockOperation
		if err := tx.Where("order_reference = ? AND action = ?", ref, "RESERVE").First(&reserved).Error; err != nil {
			return fmt.Errorf("stock reservation not found: %w", err)
		}
		if reserved.ProductID != id || reserved.Quantity != quantity {
			return fmt.Errorf("release does not match original reservation")
		}
		var released model.StockOperation
		err := tx.Where("order_reference = ? AND action = ?", ref, "RELEASE").First(&released).Error
		if err == nil {
			return nil
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if err := tx.Model(&model.Product{}).Where("id = ?", id).Update("stock", gorm.Expr("stock + ?", quantity)).Error; err != nil {
			return err
		}
		return tx.Create(&model.StockOperation{OrderReference: ref, Action: "RELEASE", ProductID: id, Quantity: quantity}).Error
	})
}
