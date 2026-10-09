CREATE DATABASE customer_db;



CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    phone VARCHAR(15) NOT NULL,
    city VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_customers_name
ON customers(name);

CREATE INDEX IF NOT EXISTS idx_customers_email
ON customers(email);

CREATE INDEX IF NOT EXISTS idx_customers_phone
ON customers(phone);
