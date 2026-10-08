-- Create the employee table if it does not already exist.
CREATE TABLE IF NOT EXISTS employees (
    -- Generate a unique numeric employee ID.
    id BIGSERIAL PRIMARY KEY,
    -- Store the employee name.
    name VARCHAR(100) NOT NULL,
    -- Store the employee email.
    email VARCHAR(150) NOT NULL,
    -- Store the employee department.
    department VARCHAR(100) NOT NULL,
    -- Store the employee salary.
    salary NUMERIC(12,2) NOT NULL
);

-- Insert sample employee data for testing.
INSERT INTO employees (name, email, department, salary) VALUES
('ganesh', 'ganesh@example.com', 'IT', 75000),
('Priya Sharma', 'priya@example.com', 'HR', 65000),
('Amit Verma', 'amit@example.com', 'Finance', 70000);
