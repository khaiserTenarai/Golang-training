-- ============================================================
-- Schema for the Bank Account System.
-- Run once against your database, e.g.:
--   psql -U postgres -d bank_account_db -f sql/schema.sql
-- ============================================================

-- The account's own ID (the "id" column) doubles as its account
-- number in this simplified system - no separate account number
-- generator is needed.
CREATE TABLE IF NOT EXISTS account (
    id                  SERIAL PRIMARY KEY,
    account_holder_name VARCHAR(150)  NOT NULL,
    balance             NUMERIC(12,2) NOT NULL DEFAULT 0,
    created_at          TIMESTAMP     NOT NULL DEFAULT NOW()
);

-- Every deposit and withdrawal gets its own row here - this is the
-- "transaction history". balance_after records what the account's
-- balance became right after this transaction, which makes the
-- history self-explanatory without needing to replay every row.
CREATE TABLE IF NOT EXISTS transaction (
    id            SERIAL PRIMARY KEY,
    account_id    INT NOT NULL REFERENCES account(id) ON DELETE CASCADE,
    type          VARCHAR(10)   NOT NULL CHECK (type IN ('DEPOSIT', 'WITHDRAWAL')),
    amount        NUMERIC(12,2) NOT NULL,
    balance_after NUMERIC(12,2) NOT NULL,
    created_at    TIMESTAMP     NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_transaction_account ON transaction (account_id);
