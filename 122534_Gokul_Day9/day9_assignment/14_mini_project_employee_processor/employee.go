package main

type Employee struct {
	ID         int
	Name       string
	Department string
	Basic      float64
}

type Payslip struct {
	Employee Employee
	Gross    float64
	Tax      float64
	Net      float64
	WorkerID int
}

type Summary struct {
	Count    int
	TotalNet float64
}

func calculatePayslip(e Employee, workerID int) Payslip {
	gross := e.Basic * 1.20
	tax := gross * 0.10

	return Payslip{
		Employee: e,
		Gross:    gross,
		Tax:      tax,
		Net:      gross - tax,
		WorkerID: workerID,
	}
}
