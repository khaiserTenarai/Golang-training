CREATE DATABASE ecommerce_db;

CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    price NUMERIC(12,2) NOT NULL CHECK (price >= 0),
    stock INT NOT NULL DEFAULT 0 CHECK (stock >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(30) NOT NULL DEFAULT 'PLACED',

    CONSTRAINT fk_order_customer
        FOREIGN KEY (customer_id)
        REFERENCES customers(id)
);

CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    price NUMERIC(12,2) NOT NULL CHECK (price >= 0),

    CONSTRAINT fk_item_order
        FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_item_product
        FOREIGN KEY (product_id)
        REFERENCES products(id)
);

CREATE INDEX IF NOT EXISTS idx_orders_customer
ON orders(customer_id);

CREATE INDEX IF NOT EXISTS idx_order_items_order
ON order_items(order_id);

CREATE INDEX IF NOT EXISTS idx_order_items_product
ON order_items(product_id);

-- Sample data
INSERT INTO customers (name, email)
VALUES
    ('Ravi Ranjan', 'ravi@example.com'),
    ('Amit Kumar', 'amit@example.com')
ON CONFLICT (email) DO NOTHING;

INSERT INTO products (name, price, stock)
VALUES
    ('Laptop', 55000.00, 10),
    ('Mouse', 800.00, 50),
    ('Keyboard', 1500.00, 30),
    ('Monitor', 12000.00, 15)
ON CONFLICT DO NOTHING;
