package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {

	connString := "postgres://postgres:admin@localhost:5432/gotraining"

	conn, err := pgx.Connect(
		context.Background(),
		connString,
	)

	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	defer conn.Close(context.Background())

	fmt.Println("PostgreSQL connected successfully")

	// =====================================
	// CREATE PRODUCT
	// =====================================

	fmt.Println("\n----- CREATE PRODUCT -----")

	var productID int

	err = conn.QueryRow(
		context.Background(),
		`
		INSERT INTO products_q5
		(name,price,stock)

		VALUES($1,$2,$3)

		RETURNING id
		`,
		"Mouse",
		800,
		20,
	).Scan(&productID)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Product created ID:", productID)

	}

	// =====================================
	// READ PRODUCTS
	// =====================================

	fmt.Println("\n----- ALL PRODUCTS -----")

	rows, err := conn.Query(
		context.Background(),
		`
		SELECT id,name,price,stock
		FROM products_q5
		ORDER BY id
		`,
	)

	if err != nil {

		log.Println(err)

	} else {

		for rows.Next() {

			var id int
			var name string
			var price float64
			var stock int

			rows.Scan(
				&id,
				&name,
				&price,
				&stock,
			)

			fmt.Println(
				id,
				name,
				price,
				stock,
			)

		}

		rows.Close()

	}

	// =====================================
	// UPDATE PRODUCT PRICE
	// =====================================

	fmt.Println("\n----- UPDATE PRODUCT -----")

	result, err := conn.Exec(
		context.Background(),
		`
		UPDATE products_q5

		SET price=$1

		WHERE id=$2
		`,
		900,
		productID,
	)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Product updated")
		fmt.Println("Rows:", result.RowsAffected())

	}

	// =====================================
	// STOCK INCREASE
	// =====================================

	fmt.Println("\n----- STOCK INCREASE -----")

	result, err = conn.Exec(
		context.Background(),
		`
		UPDATE products_q5

		SET stock = stock + $1

		WHERE id=$2
		`,
		10,
		productID,
	)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Stock increased")
		fmt.Println("Rows:", result.RowsAffected())

	}

	// =====================================
	// STOCK DECREASE
	// =====================================

	fmt.Println("\n----- STOCK DECREASE -----")

	result, err = conn.Exec(
		context.Background(),
		`
		UPDATE products_q5

		SET stock = stock - $1

		WHERE id=$2
		AND stock >= $1
		`,
		5,
		productID,
	)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Stock decreased")
		fmt.Println("Rows:", result.RowsAffected())

	}

	// =====================================
	// LOW STOCK SEARCH
	// =====================================

	fmt.Println("\n----- LOW STOCK PRODUCTS -----")

	rows, err = conn.Query(
		context.Background(),
		`
		SELECT id,name,stock

		FROM products_q5

		WHERE stock < $1
		`,
		5,
	)

	if err != nil {

		log.Println(err)

	} else {

		for rows.Next() {

			var id int
			var name string
			var stock int

			rows.Scan(
				&id,
				&name,
				&stock,
			)

			fmt.Println(
				id,
				name,
				stock,
			)

		}

		rows.Close()

	}

	// =====================================
	// DELETE PRODUCT
	// =====================================

	fmt.Println("\n----- DELETE PRODUCT -----")

	result, err = conn.Exec(
		context.Background(),
		`
		DELETE FROM products_q5
		WHERE id=$1
		`,
		productID,
	)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Product deleted")
		fmt.Println("Rows:", result.RowsAffected())

	}

}
