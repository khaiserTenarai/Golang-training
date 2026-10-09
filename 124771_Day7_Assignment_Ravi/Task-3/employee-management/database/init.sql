CREATE TABLE IF NOT EXISTS employees (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    department VARCHAR(100) NOT NULL,
    salary NUMERIC(12, 2) NOT NULL CHECK (salary >= 0)
);

INSERT INTO employees (name, department, salary)
VALUES
    ('Ranjan Kumar', 'IT', 75000),
    ('Rahul Sharma', 'HR', 55000),
    ('Priya Singh', 'IT', 85000),
    ('Amit Verma', 'Finance', 65000),
    ('Neha Gupta', 'IT', 92000),
    ('Ankit Patel', 'Finance', 72000),
    ('Sneha Joshi', 'HR', 58000),
    ('Vikas Mehta', 'IT', 68000),
    ('Pooja Shah', 'Marketing', 62000),
    ('Arjun Rao', 'IT', 88000);
