

CREATE TABLE IF NOT EXISTS department (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS new_employee (
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(100)  NOT NULL,
    email         VARCHAR(150)  NOT NULL UNIQUE,
    age           INT           NOT NULL,
    salary        NUMERIC(12,2) NOT NULL,
    department_id INT REFERENCES department(id) ON DELETE RESTRICT,
    created_at    TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS salary_history (
    id          SERIAL PRIMARY KEY,
    employee_id INT NOT NULL REFERENCES new_employee(id) ON DELETE CASCADE,
    old_salary  NUMERIC(12,2) NOT NULL,
    new_salary  NUMERIC(12,2) NOT NULL,
    changed_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_new_employee_name       ON new_employee (name);
CREATE INDEX IF NOT EXISTS idx_new_employee_salary     ON new_employee (salary);
CREATE INDEX IF NOT EXISTS idx_new_employee_department ON new_employee (department_id);
