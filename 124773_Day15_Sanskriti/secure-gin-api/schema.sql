CREATE TABLE IF NOT EXISTS users (
 id BIGSERIAL PRIMARY KEY,
 username VARCHAR(50) UNIQUE NOT NULL,
 password_hash TEXT NOT NULL,
 role VARCHAR(10) NOT NULL CHECK (role IN ('user','admin'))
);
CREATE TABLE IF NOT EXISTS employees (
 id BIGSERIAL PRIMARY KEY,
 name VARCHAR(100) NOT NULL,
 email VARCHAR(150) UNIQUE NOT NULL,
 age INT NOT NULL CHECK (age >= 18),
 salary NUMERIC(12,2) NOT NULL CHECK (salary >= 0),
 created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Create password hashes with bcrypt; never store plaintext passwords.
-- Example one-time Go helper: bcrypt.GenerateFromPassword([]byte("change-this"), bcrypt.DefaultCost)
-- Insert users with the resulting hash and role admin/user.
