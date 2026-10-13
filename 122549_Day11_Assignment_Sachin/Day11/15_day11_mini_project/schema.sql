CREATE TABLE IF NOT EXISTS employees (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(100)  NOT NULL,
    email      VARCHAR(150)  NOT NULL UNIQUE,
    department VARCHAR(100)  NOT NULL,
    salary     NUMERIC(12,2) NOT NULL
);
