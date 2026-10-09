package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"product-inventory/repository"
	"product-inventory/utility"
	"product-inventory/view"
)

type ProductServiceImpl struct {
	db         *pgxpool.Pool
	repository repository.ProductRepository
}

func NewProductService(
	db *pgxpool.Pool,
	repository repository.ProductRepository,
) *ProductServiceImpl {

	return &ProductServiceImpl{
		db:         db,
		repository: repository,
	}
}

// ----------------------------------------------------
// Create Product
// ----------------------------------------------------

func (s *ProductServiceImpl) CreateProduct(
	ctx context.Context,
	req view.CreateProductRequest,
) (*view.ProductResponse, error) {

	if err := utility.ValidateProductName(
		req.Name,
	); err != nil {
		return nil, err
	}

	if err := utility.ValidatePrice(
		req.Price,
	); err != nil {
		return nil, err
	}

	if err := utility.ValidateStock(
		req.StockQuantity,
	); err != nil {
		return nil, err
	}

	if err := utility.ValidateThreshold(
		req.LowStockThreshold,
	); err != nil {
		return nil, err
	}

	product := &viewProductToModel{
		Name:              req.Name,
		Description:       req.Description,
		Price:             req.Price,
		StockQuantity:     req.StockQuantity,
		LowStockThreshold: req.LowStockThreshold,
	}

	created, err := s.repository.Create(
		ctx,
		product.ToModel(),
	)

	if err != nil {
		return nil, err
	}

	return &view.ProductResponse{
		Product: *created,
	}, nil
}

// ----------------------------------------------------
// Get Product
// ----------------------------------------------------

func (s *ProductServiceImpl) GetProduct(
	ctx context.Context,
	id int64,
) (*view.ProductResponse, error) {

	if err := utility.ValidateProductID(id); err != nil {
		return nil, err
	}

	product, err :=
		s.repository.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return &view.ProductResponse{
		Product: *product,
	}, nil
}

// ----------------------------------------------------
// Get All Products
// ----------------------------------------------------

func (s *ProductServiceImpl) GetAllProducts(
	ctx context.Context,
) (*view.ProductListResponse, error) {

	products, err :=
		s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return &view.ProductListResponse{
		Products: products,
	}, nil
}

// ----------------------------------------------------
// Update Product
// ----------------------------------------------------

func (s *ProductServiceImpl) UpdateProduct(
	ctx context.Context,
	req view.UpdateProductRequest,
) (*view.ProductResponse, error) {

	if err := utility.ValidateProductID(
		req.ID,
	); err != nil {
		return nil, err
	}

	if err := utility.ValidateProductName(
		req.Name,
	); err != nil {
		return nil, err
	}

	if err := utility.ValidatePrice(
		req.Price,
	); err != nil {
		return nil, err
	}

	if err := utility.ValidateThreshold(
		req.LowStockThreshold,
	); err != nil {
		return nil, err
	}

	product := &viewProductToModel{
		ID:                req.ID,
		Name:              req.Name,
		Description:       req.Description,
		Price:             req.Price,
		LowStockThreshold: req.LowStockThreshold,
	}

	updated, err :=
		s.repository.Update(
			ctx,
			product.ToModel(),
		)

	if err != nil {
		return nil, err
	}

	return &view.ProductResponse{
		Product: *updated,
	}, nil
}

// ----------------------------------------------------
// Delete Product
// ----------------------------------------------------

func (s *ProductServiceImpl) DeleteProduct(
	ctx context.Context,
	id int64,
) error {

	if err := utility.ValidateProductID(id); err != nil {
		return err
	}

	return s.repository.Delete(
		ctx,
		id,
	)
}

// ----------------------------------------------------
// Increase Stock
// ----------------------------------------------------

func (s *ProductServiceImpl) IncreaseStock(
	ctx context.Context,
	req view.StockRequest,
) (*view.ProductResponse, error) {

	return s.changeStock(
		ctx,
		req,
		true,
	)
}

// ----------------------------------------------------
// Decrease Stock
// ----------------------------------------------------

func (s *ProductServiceImpl) DecreaseStock(
	ctx context.Context,
	req view.StockRequest,
) (*view.ProductResponse, error) {

	return s.changeStock(
		ctx,
		req,
		false,
	)
}

// ----------------------------------------------------
// Change Stock
// ----------------------------------------------------

func (s *ProductServiceImpl) changeStock(
	ctx context.Context,
	req view.StockRequest,
	increase bool,
) (*view.ProductResponse, error) {

	if err := utility.ValidateProductID(
		req.ProductID,
	); err != nil {
		return nil, err
	}

	if err := utility.ValidateStockChange(
		req.Quantity,
	); err != nil {
		return nil, err
	}

	tx, err := s.db.BeginTx(
		ctx,
		pgx.TxOptions{},
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to begin transaction: %w",
			err,
		)
	}

	committed := false

	defer func() {

		if !committed {
			_ = tx.Rollback(ctx)
		}

	}()

	product, err :=
		s.repository.GetByIDForUpdate(
			ctx,
			tx,
			req.ProductID,
		)

	if err != nil {
		return nil, err
	}

	newStock := product.StockQuantity

	if increase {

		newStock += req.Quantity

	} else {

		if req.Quantity > product.StockQuantity {
			return nil, fmt.Errorf(
				"insufficient stock: available=%d requested=%d",
				product.StockQuantity,
				req.Quantity,
			)
		}

		newStock -= req.Quantity
	}

	err = s.repository.UpdateStock(
		ctx,
		tx,
		product.ID,
		newStock,
	)

	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf(
			"failed to commit stock transaction: %w",
			err,
		)
	}

	committed = true

	product.StockQuantity = newStock

	return &view.ProductResponse{
		Product: *product,
	}, nil
}

// ----------------------------------------------------
// Low Stock
// ----------------------------------------------------

func (s *ProductServiceImpl) GetLowStockProducts(
	ctx context.Context,
) (*view.ProductListResponse, error) {

	products, err :=
		s.repository.GetLowStockProducts(ctx)

	if err != nil {
		return nil, err
	}

	return &view.ProductListResponse{
		Products: products,
	}, nil
}
