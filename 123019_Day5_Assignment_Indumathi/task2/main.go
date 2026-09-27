package main

type Address struct{
	City string
	State string
	Pincode int
}

type Department struct{
	ID int
	Name string
	ManagerID int
}

type Employee struct{
	ID int
	Name string
	Age int
	Salary float64
	Address Address
	Department Department
	Email string
	IsActive bool
}