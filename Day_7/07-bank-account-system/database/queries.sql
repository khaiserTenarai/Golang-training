CREATE DATABASE bank_account_db;


CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    balance NUMERIC(12,2) NOT NULL DEFAULT 0
);



CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    type VARCHAR(20) NOT NULL,
    amount NUMERIC(12,2) NOT NULL,

    FOREIGN KEY (account_id)
        REFERENCES accounts(id)
);
