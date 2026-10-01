
CREATE TABLE IF NOT EXISTS account (
    id                  SERIAL PRIMARY KEY,
    account_holder_name VARCHAR(150)  NOT NULL,
    balance             NUMERIC(12,2) NOT NULL DEFAULT 0,
    created_at          TIMESTAMP     NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transaction (
    id            SERIAL PRIMARY KEY,
    account_id    INT NOT NULL REFERENCES account(id) ON DELETE CASCADE,
    type          VARCHAR(10)   NOT NULL CHECK (type IN ('DEPOSIT', 'WITHDRAWAL')),
    amount        NUMERIC(12,2) NOT NULL,
    balance_after NUMERIC(12,2) NOT NULL,
    created_at    TIMESTAMP     NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_transaction_account ON transaction (account_id);
