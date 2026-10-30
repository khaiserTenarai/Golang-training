-- Schema setup for Microservices Assignment

CREATE TABLE IF NOT EXISTS departments (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL,
    headcount_limit INT NOT NULL DEFAULT 50,
    current_headcount INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS employees (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    department_id BIGINT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Seed Data
INSERT INTO departments (name, code, headcount_limit, current_headcount) VALUES
('Engineering', 'ENG-01', 50, 2),
('Human Resources', 'HR-01', 10, 1)
ON CONFLICT DO NOTHING;

INSERT INTO employees (first_name, last_name, email, department_id, status) VALUES
('John', 'Doe', 'john.doe@company.com', 1, 'ACTIVE'),
('Jane', 'Smith', 'jane.smith@company.com', 2, 'PENDING')
ON CONFLICT DO NOTHING;
