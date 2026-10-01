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

	ctx := context.Background()

	fmt.Println("PostgreSQL connected successfully")

	// =====================================
	// CREATE ORDER
	// =====================================

	fmt.Println("\n----- CREATE ORDER -----")

	var orderID int

	err = conn.QueryRow(
		ctx,
		`
		INSERT INTO orders_q9
		(customer_id)

		VALUES($1)

		RETURNING id
		`,
		1,
	).Scan(&orderID)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Order created ID:", orderID)

	}

	// =====================================
	// ADD ORDER ITEMS
	// =====================================

	fmt.Println("\n----- ADD ORDER ITEMS -----")

	_, err = conn.Exec(
		ctx,
		`
		INSERT INTO order_items_q9
		(order_id,product_id,quantity)

		VALUES($1,$2,$3)
		`,
		orderID,
		1,
		1,
	)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Laptop added")

	}

	_, err = conn.Exec(
		ctx,
		`
		INSERT INTO order_items_q9
		(order_id,product_id,quantity)

		VALUES($1,$2,$3)
		`,
		orderID,
		2,
		2,
	)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Mouse added")

	}

	// =====================================
	// COMPLETE ORDER DETAILS USING JOIN
	// =====================================

	fmt.Println("\n----- ORDER DETAILS -----")

	rows, err := conn.Query(
		ctx,
		`
		SELECT

		c.name,

		o.id,

		p.name,

		p.price,

		oi.quantity,

		(p.price * oi.quantity) AS total


		FROM customers_q9 c


		JOIN orders_q9 o

		ON c.id=o.customer_id


		JOIN order_items_q9 oi

		ON o.id=oi.order_id


		JOIN products_q9 p

		ON p.id=oi.product_id


		WHERE o.id=$1
		`,
		orderID,
	)

	if err != nil {

		log.Println(err)

	} else {

		defer rows.Close()

		for rows.Next() {

			var customer string
			var order int
			var product string
			var price float64
			var quantity int
			var total float64

			err := rows.Scan(
				&customer,
				&order,
				&product,
				&price,
				&quantity,
				&total,
			)

			if err != nil {

				log.Println(err)
				continue

			}

			fmt.Println("----------------")
			fmt.Println("Customer:", customer)
			fmt.Println("Order ID:", order)
			fmt.Println("Product:", product)
			fmt.Println("Price:", price)
			fmt.Println("Quantity:", quantity)
			fmt.Println("Total:", total)

		}

	}

}
