CREATE TABLE IF NOT EXISTS employees (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    department VARCHAR(100) NOT NULL,
    salary NUMERIC(12, 2) NOT NULL CHECK (salary >= 0)
);

CREATE TABLE IF NOT EXISTS salary_history (
    id BIGSERIAL PRIMARY KEY,

    employee_id BIGINT NOT NULL,

    old_salary NUMERIC(12, 2) NOT NULL,

    new_salary NUMERIC(12, 2) NOT NULL,

    changed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_salary_history_employee
        FOREIGN KEY (employee_id)
        REFERENCES employees(id)
        ON DELETE CASCADE,

    CONSTRAINT salary_change_check
        CHECK (old_salary <> new_salary)
);

INSERT INTO employees (
    name,
    department,
    salary
)
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

CREATE TABLE IF NOT EXISTS salary_history (
    id BIGSERIAL PRIMARY KEY,

    employee_id BIGINT NOT NULL,

    old_salary NUMERIC(12, 2) NOT NULL,

    new_salary NUMERIC(12, 2) NOT NULL,

    changed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_salary_history_employee
        FOREIGN KEY (employee_id)
        REFERENCES employees(id)
        ON DELETE CASCADE,

    CONSTRAINT salary_change_check
        CHECK (old_salary <> new_salary)
);
