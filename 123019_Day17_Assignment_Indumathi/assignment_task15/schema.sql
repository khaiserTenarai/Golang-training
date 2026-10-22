CREATE TABLE IF NOT EXISTS employees (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    department VARCHAR(100) NOT NULL,
    salary NUMERIC(12,2) NOT NULL
);

INSERT INTO employees (name, email, department, salary) VALUES
('Rajesh Kumar', 'rajesh@example.com', 'IT', 75000),
('Priya Sharma', 'priya@example.com', 'HR', 65000),
('Amit Verma', 'amit@example.com', 'Finance', 70000)
ON CONFLICT (email) DO NOTHING;