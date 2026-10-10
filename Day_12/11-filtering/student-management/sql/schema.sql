CREATE TABLE IF NOT EXISTS students (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT NOT NULL,
    grade VARCHAR(20) NOT NULL
);

INSERT INTO students (name, age, grade)
VALUES
    ('Ravi', 20, 'A'),
    ('Priya', 21, 'B'),
    ('Arun', 22, 'A');
