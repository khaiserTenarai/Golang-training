package repository

import (
    "context"
    "errors"
    "fmt"

    "example.com/employee-management/model"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
)

type PostgresProductRepository struct {
    db *pgxpool.Pool
}

func NewPostgresProductRepository(db *pgxpool.Pool) *PostgresProductRepository {
    return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) Save(p model.Product) error {
    query := `
        INSERT INTO products (name, category, price, stock_quantity, reorder_level)
        VALUES ($1, $2, $3, $4, $5) RETURNING id
    `
    return r.db.QueryRow(context.Background(), query, p.Name, p.Category, p.Price, p.StockQuantity, p.ReorderLevel).Scan(&p.ID)
}

func (r *PostgresProductRepository) FindByID(id int64) (model.Product, error) {
    var p model.Product
    query := `SELECT id, name, category, price, stock_quantity, reorder_level FROM products WHERE id = $1`
    err := r.db.QueryRow(context.Background(), query, id).Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.StockQuantity, &p.ReorderLevel)
    if errors.Is(err, pgx.ErrNoRows) {
        return p, ErrProductNotFound
    }
    return p, err
}

func (r *PostgresProductRepository) FindAll() ([]model.Product, error) {
    query := `SELECT id, name, category, price, stock_quantity, reorder_level FROM products ORDER BY id`
    rows, err := r.db.Query(context.Background(), query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []model.Product
    for rows.Next() {
        var p model.Product
        if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.StockQuantity, &p.ReorderLevel); err != nil {
            return nil, err
        }
        products = append(products, p)
    }
    return products, nil
}

func (r *PostgresProductRepository) Update(p model.Product) error {
    query := `
        UPDATE products
        SET name = $1, category = $2, price = $3, stock_quantity = $4, reorder_level = $5, updated_at = CURRENT_TIMESTAMP
        WHERE id = $6
    `
    result, err := r.db.Exec(context.Background(), query, p.Name, p.Category, p.Price, p.StockQuantity, p.ReorderLevel, p.ID)
    if err != nil {
        return err
    }
    if result.RowsAffected() == 0 {
        return ErrProductNotFound
    }
    return nil
}

func (r *PostgresProductRepository) Delete(id int64) error {
    result, err := r.db.Exec(context.Background(), `DELETE FROM products WHERE id = $1`, id)
    if err != nil {
        return err
    }
    if result.RowsAffected() == 0 {
        return ErrProductNotFound
    }
    return nil
}

func (r *PostgresProductRepository) AdjustStock(id int64, amount int) error {
    ctx := context.Background()
    tx, err := r.db.Begin(ctx)
    if err != nil {
        return fmt.Errorf("failed to start transaction: %w", err)
    }
    defer tx.Rollback(ctx)

    var currentStock int
    err = tx.QueryRow(ctx, `SELECT stock_quantity FROM products WHERE id = $1 FOR UPDATE`, id).Scan(&currentStock)
    if errors.Is(err, pgx.ErrNoRows) {
        return ErrProductNotFound
    }
    if err != nil {
        return err
    }

    newStock := currentStock + amount
    if newStock < 0 {
        return ErrInsufficientStock
    }

    _, err = tx.Exec(ctx, `UPDATE products SET stock_quantity = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`, newStock, id)
    if err != nil {
        return err
    }

    return tx.Commit(ctx)
}

func (r *PostgresProductRepository) FindLowStockProducts() ([]model.Product, error) {
    query := `
        SELECT id, name, category, price, stock_quantity, reorder_level
        FROM products
        WHERE stock_quantity <= reorder_level
        ORDER BY stock_quantity ASC
    `
    rows, err := r.db.Query(context.Background(), query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []model.Product
    for rows.Next() {
        var p model.Product
        if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.StockQuantity, &p.ReorderLevel); err != nil {
            return nil, err
        }
        products = append(products, p)
    }
    return products, nil
}