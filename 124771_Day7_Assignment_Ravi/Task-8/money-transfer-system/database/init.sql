CREATE DATABASE money_transfer_db;


CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    account_number VARCHAR(20) UNIQUE NOT NULL,
    customer_name VARCHAR(100) NOT NULL,
    balance NUMERIC(12,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    from_account_id INT,
    to_account_id INT,
    transaction_type VARCHAR(30) NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_from_account
        FOREIGN KEY (from_account_id)
        REFERENCES accounts(id),

    CONSTRAINT fk_to_account
        FOREIGN KEY (to_account_id)
        REFERENCES accounts(id)
);


INSERT INTO accounts (account_number, customer_name, balance)
VALUES
    ('10001', 'Ravi Ranjan', 10000.00),
    ('10002', 'Amit Kumar', 5000.00)
ON CONFLICT (account_number) DO NOTHING;
