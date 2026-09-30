CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    quantity INT NOT NULL
);
-- add sample products

INSERT INTO products
(name, price, quantity)
VALUES
('Laptop', 55000, 10),
('Mouse', 500, 25),
('Keyboard', 1200, 5),
('Monitor', 15000, 3),
('Headphone', 2500, 15);
