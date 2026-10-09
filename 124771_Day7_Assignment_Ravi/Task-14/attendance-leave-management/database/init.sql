CREATE DATABASE attendance_db;


CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS attendance (
    id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL,
    attendance_date DATE NOT NULL,
    check_in TIMESTAMP,
    check_out TIMESTAMP,

    CONSTRAINT fk_attendance_employee
        FOREIGN KEY (employee_id)
        REFERENCES employees(id)
        ON DELETE CASCADE,

    CONSTRAINT unique_employee_date
        UNIQUE (employee_id, attendance_date)
);

CREATE TABLE IF NOT EXISTS leaves (
    id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL,
    leave_date DATE NOT NULL,
    reason VARCHAR(255) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING',

    CONSTRAINT fk_leave_employee
        FOREIGN KEY (employee_id)
        REFERENCES employees(id)
        ON DELETE CASCADE,

    CONSTRAINT valid_leave_status
        CHECK (status IN ('PENDING', 'APPROVED', 'REJECTED'))
);

CREATE INDEX IF NOT EXISTS idx_attendance_employee
ON attendance(employee_id);

CREATE INDEX IF NOT EXISTS idx_attendance_date
ON attendance(attendance_date);

CREATE INDEX IF NOT EXISTS idx_leaves_employee
ON leaves(employee_id);

INSERT INTO employees (name, email)
VALUES
    ('Ravi Ranjan', 'ravi@example.com'),
    ('Amit Kumar', 'amit@example.com'),
    ('Priya Sharma', 'priya@example.com')
ON CONFLICT (email) DO NOTHING;
