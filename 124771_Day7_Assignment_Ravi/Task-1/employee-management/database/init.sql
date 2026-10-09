CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    age INT NOT NULL CHECK (age >= 18),
    salary NUMERIC(12,2) NOT NULL CHECK (salary >= 0),
    city VARCHAR(100) NOT NULL,
    state VARCHAR(100) NOT NULL,
    pincode VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO employees
(name, email, age, salary, city, state, pincode)
VALUES
('Rajesh', 'rajesh@gmail.com', 35, 50000, 'Bengaluru', 'Karnataka', '560001'),
('Amit', 'amit@gmail.com', 30, 60000, 'Mumbai', 'Maharashtra', '400001')
ON CONFLICT (email) DO NOTHING;
