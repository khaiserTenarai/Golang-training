-- Employee Service Table
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

-- Transactional Outbox Pattern Table for Saga / Asynchronous Messaging
CREATE TABLE IF NOT EXISTS outbox_events (
    id BIGSERIAL PRIMARY KEY,
    aggregate_type VARCHAR(100) NOT NULL,
    aggregate_id BIGINT NOT NULL,
    event_type VARCHAR(100) NOT NULL,
    payload JSONB NOT NULL,
    status VARCHAR(50) DEFAULT 'PENDING',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Department Service Table
CREATE TABLE IF NOT EXISTS departments (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL,
    headcount_limit INT NOT NULL DEFAULT 50,
    current_headcount INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 1. Seed Department Data
INSERT INTO departments (name, code, headcount_limit, current_headcount) 
VALUES 
    ('Engineering', 'ENG-01', 50, 2),
    ('Human Resources', 'HR-01', 10, 1),
    ('Product Management', 'PROD-01', 15, 0);

-- 2. Seed Employee Data (department_id references department primary keys 1 and 2)
INSERT INTO employees (first_name, last_name, email, department_id, status) 
VALUES 
    ('Alice', 'Smith', 'alice.smith@company.com', 1, 'ACTIVE'),
    ('Bob', 'Johnson', 'bob.johnson@company.com', 1, 'PENDING'),
    ('Carol', 'Williams', 'carol.williams@company.com', 2, 'ACTIVE');

-- 3. Seed Outbox Events Data (simulating initial onboarding events for Saga/Messaging)
INSERT INTO outbox_events (aggregate_type, aggregate_id, event_type, payload, status) 
VALUES 
    ('EMPLOYEE', 1, 'EmployeeOnboardingInitiated', '{"employee_id": 1, "department_id": 1, "email": "alice.smith@company.com"}', 'COMPLETED'),
    ('EMPLOYEE', 2, 'EmployeeOnboardingInitiated', '{"employee_id": 2, "department_id": 1, "email": "bob.johnson@company.com"}', 'PENDING');