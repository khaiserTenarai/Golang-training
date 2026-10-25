/*
*TEST 1: Get all employees with nested department

	query {
	  employees {
	    id
	    name
	    email
	    department {
	      name
	    }
	    salary
	  }
	}

TEST 2: Get one employee

	query {
	  employee(id: "4") {
	    id
	    name
	    email
	    department {
	      name
	    }
	    salary
	  }
	}

TEST 3: Create employee

	mutation {
	  createEmployee(input: {
	    name: "Task4 Test User"
	    email: "task4test@example.com"
	    department: "IT"
	    salary: 50000
	  }) {
	    id
	    name
	    email
	    department {
	      name
	    }
	    salary
	  }
	}

TEST 4: Update employee

	mutation {
	  updateEmployee(
	    id: "4"
	    input: {
	      name: "Task4 Updated User"
	      email: "task4test@example.com"
	      department: "Engineering"
	      salary: 60000
	    }
	  ) {
	    id
	    name
	    email
	    department {
	      name
	    }
	    salary
	  }
	}

TEST 5: Delete employee
WARNING: This deletes employee ID 4.

	mutation {
	  deleteEmployee(id: "4")
	}
*/
package graphql
