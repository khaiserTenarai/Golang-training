-- ============================================================
-- Schema for the Employee / Department service.
-- Run this once against your database, e.g.:
--   psql -U postgres -d employee_department_db -f sql/schema.sql
-- ============================================================

-- One row per department (Engineering, HR, Sales, etc).
CREATE TABLE IF NOT EXISTS department (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Employee table (named "new_employee" as requested).
-- department_id is a foreign key back to department.id - this is
-- the "employee <-> department mapping".
--
-- ON DELETE RESTRICT means Postgres will refuse to delete a
-- department while any employee still points to it. That's what
-- protects the data - the Go code doesn't have to check this itself,
-- the database enforces it for us.
CREATE TABLE IF NOT EXISTS new_employee (
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(100)  NOT NULL,
    email         VARCHAR(150)  NOT NULL UNIQUE,
    age           INT           NOT NULL,
    salary        NUMERIC(12,2) NOT NULL,
    department_id INT REFERENCES department(id) ON DELETE RESTRICT,
    created_at    TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Every time an employee's salary changes, a new row is added here
-- instead of throwing the old value away - this is the "salary
-- history" required for Salary Management.
--
-- ON DELETE CASCADE here (unlike department above) because a salary
-- history row is meaningless once the employee it belongs to is gone.
CREATE TABLE IF NOT EXISTS salary_history (
    id          SERIAL PRIMARY KEY,
    employee_id INT NOT NULL REFERENCES new_employee(id) ON DELETE CASCADE,
    old_salary  NUMERIC(12,2) NOT NULL,
    new_salary  NUMERIC(12,2) NOT NULL,
    changed_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Indexes to help the search feature (name / department / salary
-- filtering and sorting all use these columns).
CREATE INDEX IF NOT EXISTS idx_new_employee_name       ON new_employee (name);
CREATE INDEX IF NOT EXISTS idx_new_employee_salary     ON new_employee (salary);
CREATE INDEX IF NOT EXISTS idx_new_employee_department ON new_employee (department_id);
