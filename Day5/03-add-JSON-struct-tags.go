package main

type Employee struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
    Salary float64 `json:"salary"`
    Email string `json:"email"`
    Phone string `json:"phone"`
    Address Address `json:"address"`
    Department Department `json:"department"`
}

type Address struct {
    City string `json:"city"`
    State string `json:"state"`
    Pincode int `json:"pincode"`
}

type Department struct {
    Name string `json:"name"`
    Location string `json:"location"`
}