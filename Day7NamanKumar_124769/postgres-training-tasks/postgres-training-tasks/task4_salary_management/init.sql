CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    salary NUMERIC(12, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS salary_history (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id),
    old_salary NUMERIC(12, 2) NOT NULL,
    new_salary NUMERIC(12, 2) NOT NULL,
    changed_at TIMESTAMP DEFAULT NOW()
);
