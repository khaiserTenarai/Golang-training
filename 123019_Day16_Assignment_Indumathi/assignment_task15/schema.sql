CREATE TABLE IF NOT EXISTS employees (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    department_id INT NOT NULL,
    department_name VARCHAR(100) NOT NULL,
    salary NUMERIC(12,2) NOT NULL
);

INSERT INTO employees (name, email, department_id, department_name, salary) VALUES
('Vimal', 'vimal@example.com', 101, 'IT', 75000),
('Aditya', 'aditya@example.com', 102, 'HR', 65000),
('Adi', 'adi@example.com', 103, 'Finance', 70000)
ON CONFLICT (email) DO NOTHING;