CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    category VARCHAR(100) NOT NULL,
    price NUMERIC(12,2) NOT NULL CHECK (price >= 0),
    stock_quantity INT NOT NULL CHECK (stock_quantity >= 0),
    reorder_level INT NOT NULL DEFAULT 10 CHECK (reorder_level >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO products (name, category, price, stock_quantity, reorder_level) VALUES
('Wireless Mouse', 'Electronics', 29.99, 50, 15),
('Mechanical Keyboard', 'Electronics', 89.99, 8, 10),
('USB-C Cable', 'Electronics', 12.50, 5, 20)
ON CONFLICT DO NOTHING;