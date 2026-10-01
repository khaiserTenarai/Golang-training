-- ============================================================
-- Schema for the Product Inventory system.
-- Run once against your database, e.g.:
--   psql -U postgres -d product_inventory_db -f sql/schema.sql
-- ============================================================

CREATE TABLE IF NOT EXISTS product (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(150)  NOT NULL,
    sku        VARCHAR(50)   NOT NULL UNIQUE,
    category   VARCHAR(100)  NOT NULL,
    price      NUMERIC(12,2) NOT NULL,
    quantity   INT           NOT NULL DEFAULT 0,
    created_at TIMESTAMP     NOT NULL DEFAULT NOW()
);

-- Helps the low-stock search and general lookups.
CREATE INDEX IF NOT EXISTS idx_product_quantity ON product (quantity);
CREATE INDEX IF NOT EXISTS idx_product_name     ON product (name);
