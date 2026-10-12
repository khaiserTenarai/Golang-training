package main

import "fmt"

func producer(jobs chan<- Employee) {
	list := []Employee{
		{1, "Asha", "Sales", 30000},
		{2, "Ravi", "Tech", 45000},
		{3, "Meena", "HR", 38000},
		{4, "Karthik", "Support", 28000},
		{5, "Divya", "Tech", 52000},
		{6, "Arun", "Sales", 33000},
		{7, "Sneha", "HR", 36000},
		{8, "Vikram", "Tech", 48000},
	}

	for _, e := range list {
		fmt.Println("producer: queued", e.Name)
		jobs <- e
	}
	close(jobs)
}
