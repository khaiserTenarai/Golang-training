package main

import (
	"context"

	"fmt"

	"log"

	"time"

	"github.com/jackc/pgx/v5"
)

func main() {

	connString :=
		"postgres://postgres:Info%40131@localhost:5432/go_training"

	conn, err := pgx.Connect(

		context.Background(),

		connString,
	)

	if err != nil {

		log.Fatal(err)

	}

	defer conn.Close(context.Background())

	ctx := context.Background()

	fmt.Println("PostgreSQL connected successfully")

	employeeID := 1

	// =====================================
	// CHECK IN
	// =====================================

	fmt.Println("\n----- CHECK IN -----")

	var attendanceID int

	err = conn.QueryRow(

		ctx,

		`
		INSERT INTO attendance
		(employee_id,check_in)

		VALUES($1,$2)

		RETURNING id
		`,

		employeeID,

		time.Now(),
	).Scan(&attendanceID)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println(
			"Check-in successful ID:",
			attendanceID,
		)

	}

	// =====================================
	// CHECK OUT
	// =====================================

	fmt.Println("\n----- CHECK OUT -----")

	_, err = conn.Exec(

		ctx,

		`
		UPDATE attendance

		SET check_out=$1

		WHERE id=$2

		`,

		time.Now(),

		attendanceID,
	)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Check-out successful")

	}

	// =====================================
	// APPLY LEAVE
	// =====================================

	fmt.Println("\n----- APPLY LEAVE -----")

	var leaveID int

	err = conn.QueryRow(

		ctx,

		`
		INSERT INTO leave

		(employee_id,reason,status)

		VALUES($1,$2,$3)

		RETURNING id
		`,

		employeeID,

		"Medical Leave",

		"PENDING",
	).Scan(&leaveID)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println(
			"Leave applied ID:",
			leaveID,
		)

	}

	// =====================================
	// APPROVE LEAVE
	// =====================================

	fmt.Println("\n----- APPROVE LEAVE -----")

	_, err = conn.Exec(

		ctx,

		`
		UPDATE leave

		SET status=$1

		WHERE id=$2
		`,

		"APPROVED",

		leaveID,
	)

	if err != nil {

		log.Println(err)

	} else {

		fmt.Println("Leave approved")

	}

	// =====================================
	// ATTENDANCE REPORT
	// =====================================

	fmt.Println("\n----- ATTENDANCE REPORT -----")

	rows, err := conn.Query(

		ctx,

		`
		SELECT

		e.name,

		a.check_in,

		a.check_out


		FROM employees e


		JOIN attendance a

		ON e.id=a.employee_id

		`,
	)

	if err != nil {

		log.Println(err)

	} else {

		defer rows.Close()

		for rows.Next() {

			var name string

			var checkIn time.Time

			var checkOut time.Time

			rows.Scan(

				&name,

				&checkIn,

				&checkOut,
			)

			fmt.Println("----------------")

			fmt.Println("Name:", name)

			fmt.Println(
				"Check In:",
				checkIn,
			)

			fmt.Println(
				"Check Out:",
				checkOut,
			)

		}

	}

}
