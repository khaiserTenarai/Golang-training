CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    balance NUMERIC(14, 2) NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    account_id INT REFERENCES accounts(id),
    type VARCHAR(20) NOT NULL,
    amount NUMERIC(14, 2) NOT NULL,
    related_account_id INT,
    created_at TIMESTAMP DEFAULT NOW()
);
