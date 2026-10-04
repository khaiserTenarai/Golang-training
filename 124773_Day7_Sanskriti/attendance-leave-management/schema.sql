CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    department VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS attendance (
    id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL REFERENCES employees(id) ON DELETE CASCADE,
    attendance_date DATE NOT NULL DEFAULT CURRENT_DATE,
    check_in TIMESTAMP,
    check_out TIMESTAMP,
    UNIQUE(employee_id, attendance_date)
);

CREATE TABLE IF NOT EXISTS leaves (
    id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL REFERENCES employees(id) ON DELETE CASCADE,
    leave_type VARCHAR(50) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    reason TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'Pending',
    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CHECK (end_date >= start_date),
    CHECK (status IN ('Pending', 'Approved', 'Rejected'))
);

INSERT INTO employees (name, email, department) VALUES
('Sanskriti Sahu', 'sanskriti@example.com', 'IT'),
('Rahul Sharma', 'rahul@example.com', 'HR'),
('Priya Singh', 'priya@example.com', 'Finance')
ON CONFLICT (email) DO NOTHING;
