package service
import (
	"fmt"
	"sync"
	"time"

	"ems/model"
	"ems/repository"
	"ems/utility"
)

/*
	EmployeeServiceImpl implements
	EmployeeService.
*/
type EmployeeServiceImpl struct {

	// Existing Day 7 repository.
	repository repository.EmployeeRepository

	/*
		Day 9:

		Number of workers.
	*/
	workers int

	/*
		Day 9:

		Channel buffer size.
	*/
	buffer int
}

/*
	Constructor.
*/
func NewEmployeeService(
	repository repository.EmployeeRepository,
	workers int,
	buffer int,
) EmployeeService {

	return &EmployeeServiceImpl{
		repository: repository,
		workers:    workers,
		buffer:     buffer,
	}
}

/*
	==================================================
	EXISTING DAY 7 METHODS
	==================================================
*/

/*
	Save employee.
*/
func (s EmployeeServiceImpl) Save(
	employee model.Employee,
) error {

	err := utility.ValidateEmployee(employee)

	if err != nil {
		return err
	}

	err = utility.ValidateEmail(employee.Email)

	if err != nil {
		return err
	}

	return s.repository.Save(employee)
}

/*
	Find employee by ID.
*/
func (s EmployeeServiceImpl) FindByID(
	id int,
) (model.Employee, error) {

	err := utility.ValidateEmployeeID(id)

	if err != nil {
		return model.Employee{}, err
	}

	return s.repository.FindByID(id)
}

/*
	Find all employees.
*/
func (s EmployeeServiceImpl) FindAll() (
	[]model.Employee,
	error,
) {

	return s.repository.FindAll()
}

/*
	Update employee.
*/
func (s EmployeeServiceImpl) Update(
	employee model.Employee,
) error {

	err := utility.ValidateEmployee(employee)

	if err != nil {
		return err
	}

	err = utility.ValidateEmail(employee.Email)

	if err != nil {
		return err
	}

	return s.repository.Update(employee)
}

/*
	Delete employee.
*/
func (s EmployeeServiceImpl) Delete(
	id int,
) error {

	err := utility.ValidateEmployeeID(id)

	if err != nil {
		return err
	}

	return s.repository.Delete(id)
}

/*
	==================================================
	DAY 9 - CONCURRENT PROCESSING
	==================================================
*/

/*
	ProcessEmployees performs concurrent
	employee processing.

	Flow:

	Database
	    ↓
	Repository
	    ↓
	Service
	    ↓
	Producer
	    ↓
	Buffered Channel
	    ↓
	Workers
	    ↓
	Processing
*/
func (s EmployeeServiceImpl) ProcessEmployees() {

	/*
		Get all employees from PostgreSQL.
	*/
	employees, err := s.repository.FindAll()

	if err != nil {
		fmt.Println("Error while fetching employees:", err)
		return
	}

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	fmt.Println()
	fmt.Println("========== Concurrent Processing ==========")
	fmt.Println("Total Employees:", len(employees))
	fmt.Println("Workers:", s.workers)
	fmt.Println("Channel Buffer:", s.buffer)

	/*
		==================================================
		CHANNEL
		==================================================

		This is a buffered channel.

		If buffer = 2:

		    Employee 1
		    Employee 2
		    ----------------
		    Channel capacity = 2

		The producer can place two employees
		without waiting.

		If the channel becomes full,
		the producer has to wait.

		This waiting is called BACKPRESSURE.
	*/
	employeeChannel := make(
		chan model.Employee,
		s.buffer,
	)

	/*
		WaitGroup is used to wait for
		all workers to finish.
	*/
	var wg sync.WaitGroup

	/*
		==================================================
		WORKERS
		==================================================

		Create multiple goroutines.

		Each goroutine is a worker.

		For example:

		Worker 1
		Worker 2
		Worker 3
	*/
	for i := 1; i <= s.workers; i++ {

		wg.Add(1)

		go s.worker(
			i,
			employeeChannel,
			&wg,
		)
	}

	/*
		==================================================
		PRODUCER
		==================================================

		The producer sends employees
		into the channel.

		It runs as a goroutine.
	*/
	go func() {

		/*
			Close the channel after all
			employees have been sent.

			Workers use this to know
			that there are no more employees.
		*/
		defer close(employeeChannel)

		for _, employee := range employees {

			/*
				SELECT

				select waits for a channel
				operation to become ready.

				Here we select between:

				1. Sending employee to channel
				2. Waiting for a small timeout

				The timeout demonstrates
				that select can handle multiple
				possible events.
			*/
			select {

			case employeeChannel <- employee:

				fmt.Println(
					"Producer sent Employee:",
					employee.ID,
				)

			case <-time.After(2 * time.Second):

				fmt.Println(
					"Producer timeout for Employee:",
					employee.ID,
				)

				return
			}

			/*
				Small delay makes the producer
				and worker activity easier to observe.

				It also helps us understand
				backpressure.
			*/
			time.Sleep(100 * time.Millisecond)
		}

		fmt.Println("Producer completed.")

	}()

	/*
		Wait until all workers finish.
	*/
	wg.Wait()

	fmt.Println()
	fmt.Println("All employees processed.")
}

/*
	==================================================
	WORKER
	==================================================

	A worker is a consumer.

	Workers receive employees from
	the employee channel.
*/
func (s EmployeeServiceImpl) worker(
	workerID int,
	employeeChannel <-chan model.Employee,
	wg *sync.WaitGroup,
) {

	/*
		Tell WaitGroup that this worker
		has completed.
	*/
	defer wg.Done()

	for {

		/*
			SELECT

			The worker waits for an employee
			from the channel.

			When the channel is closed,
			ok becomes false.
		*/
		select {

		case employee, ok := <-employeeChannel:

			/*
				Channel is closed and there
				are no more employees.
			*/
			if !ok {

				fmt.Println(
					"Worker",
					workerID,
					"stopped.",
				)

				return
			}

			/*
				Process this employee.
			*/
			s.processEmployee(
				workerID,
				employee,
			)
		}
	}
}

/*
	processEmployee contains the actual
	employee processing logic.
*/
func (s EmployeeServiceImpl) processEmployee(
	workerID int,
	employee model.Employee,
) {

	fmt.Printf(
		"Worker %d processing Employee %d - %s\n",
		workerID,
		employee.ID,
		employee.Name,
	)

	/*
		Simulate some processing time.

		Because different workers run
		at the same time, employees can
		be processed concurrently.
	*/
	time.Sleep(500 * time.Millisecond)

	/*
		Validate employee before processing.
	*/
	err := utility.ValidateEmployee(employee)

	if err != nil {

		fmt.Printf(
			"Worker %d -> Employee %d failed: %v\n",
			workerID,
			employee.ID,
			err,
		)

		return
	}

	/*
		Example Day 9 processing:

		Calculate a 10% salary increment.

		The database is not changed.
		This is only concurrent processing.
	*/
	newSalary := employee.Salary +
		(employee.Salary * 10 / 100)

	fmt.Printf(
		"Worker %d completed Employee %d | Old Salary: %.2f | New Salary: %.2f\n",
		workerID,
		employee.ID,
		employee.Salary,
		newSalary,
	)
}

/*
	This helper is optional if you want to convert
	worker/buffer values inside the service.

	It is kept simple.

	Currently main.go performs the conversion.
*/
