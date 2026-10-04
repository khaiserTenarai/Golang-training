package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	ID                int
	Name              string
	SKU               string
	Price             float64
	StockQuantity     int
	LowStockThreshold int
}

func main() {
	connStr := "host=localhost port=5432 user=postgres password=pgadmin dbname=gotraining sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	prodID := createProduct(db, "Wireless Keyboard", "WK-100", 49.99, 12, 15)
	fmt.Printf("Created Product ID: %d\n", prodID)

	p := getProduct(db, prodID)
	fmt.Printf("Read: %+v\n", p)

	updateProduct(db, prodID, "Wireless Keyboard V2", "WK-101", 54.99, 15)

	adjustStock(db, prodID, -10) 

	lowStockProds := getLowStockProducts(db)
	for _, prod := range lowStockProds {
		fmt.Printf("Low Stock Alert: %+v\n", prod)
	}

	deleteProduct(db, prodID)
}

func createProduct(db *sql.DB, name, sku string, price float64, stock, threshold int) int {
	var id int
	err := db.QueryRow(`
		INSERT INTO products (name, sku, price, stock_quantity, low_stock_threshold) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		name, sku, price, stock, threshold).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func getProduct(db *sql.DB, id int) Product {
	var p Product
	err := db.QueryRow(`
		SELECT id, name, sku, price, stock_quantity, low_stock_threshold 
		FROM products WHERE id = $1`, id).
		Scan(&p.ID, &p.Name, &p.SKU, &p.Price, &p.StockQuantity, &p.LowStockThreshold)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func updateProduct(db *sql.DB, id int, name, sku string, price float64, threshold int) {
	_, err := db.Exec(`
		UPDATE products SET name = $1, sku = $2, price = $3, low_stock_threshold = $4 
		WHERE id = $5`,
		name, sku, price, threshold, id)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteProduct(db *sql.DB, id int) {
	_, err := db.Exec(`DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
}

func adjustStock(db *sql.DB, id int, amount int) {
	_, err := db.Exec(`
		UPDATE products SET stock_quantity = stock_quantity + $1 
		WHERE id = $2`, amount, id)
	if err != nil {
		log.Fatal(err)
	}
}

func getLowStockProducts(db *sql.DB) []Product {
	rows, err := db.Query(`
		SELECT id, name, sku, price, stock_quantity, low_stock_threshold 
		FROM products WHERE stock_quantity <= low_stock_threshold`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.SKU, &p.Price, &p.StockQuantity, &p.LowStockThreshold); err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}
	
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	
	return products
}