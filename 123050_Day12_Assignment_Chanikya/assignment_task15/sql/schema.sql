-- Table structural schema documentation reference blueprint
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    department VARCHAR(100) NOT NULL,
    salary NUMERIC(15, 2) NOT NULL
);
