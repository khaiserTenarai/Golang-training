-- Run against department_db
CREATE TABLE IF NOT EXISTS departments (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
INSERT INTO departments (name) VALUES ('IT'), ('HR'), ('Finance')
ON CONFLICT (name) DO NOTHING;
