-- Run once:  psql -U postgres -d employee_db -f schema.sql
-- (the program also creates this table itself if it's missing)
CREATE TABLE IF NOT EXISTS employees (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(100)  NOT NULL,
    email      VARCHAR(150)  NOT NULL UNIQUE,
    department VARCHAR(100)  NOT NULL,
    salary     NUMERIC(12,2) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_employees_department ON employees (department);
CREATE INDEX IF NOT EXISTS idx_employees_salary     ON employees (salary);
