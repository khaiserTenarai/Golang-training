CREATE DATABASE concurrency_emp_db;


CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT NOT NULL,
    email VARCHAR(100) NOT NULL,
    salary NUMERIC(10,2) NOT NULL
);

INSERT INTO employees (name, age, email, salary) VALUES
('Ganesh', 23, 'ganesh@gmail.com', 45000),
('Ravi', 25, 'ravi@gmail.com', 50000),
('Suresh', 28, 'suresh@gmail.com', 55000),
('Kiran', 24, 'kiran@gmail.com', 48000),
('Arun', 30, 'arun@gmail.com', 60000),
('Vijay', 27, 'vijay@gmail.com', 52000),
('Raj', 29, 'raj@gmail.com', 58000),
('Kumar', 26, 'kumar@gmail.com', 47000);