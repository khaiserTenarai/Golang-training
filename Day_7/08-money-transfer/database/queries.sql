CREATE DATABASE money_transfer_db;


CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    balance NUMERIC(12,2) NOT NULL DEFAULT 0
);

INSERT INTO accounts
(name, balance)
VALUES
('Rahul', 10000),
('Amit', 5000),
('Priya', 8000);



SELECT * FROM accounts;
