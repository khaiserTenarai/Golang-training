-- ============================================================
-- Schema for the Customer Management system.
-- Run once against your database, e.g.:
--   psql -U postgres -d customer_management_db -f sql/schema.sql
-- ============================================================

CREATE TABLE IF NOT EXISTS customer (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(150) NOT NULL,
    email      VARCHAR(150) NOT NULL UNIQUE,
    phone      VARCHAR(20)  NOT NULL,
    address    VARCHAR(255),
    created_at TIMESTAMP    NOT NULL DEFAULT NOW()
);

-- Helps the search feature (name/email/phone lookups).
CREATE INDEX IF NOT EXISTS idx_customer_name  ON customer (name);
CREATE INDEX IF NOT EXISTS idx_customer_email ON customer (email);
CREATE INDEX IF NOT EXISTS idx_customer_phone ON customer (phone);
