CREATE TABLE IF NOT EXISTS students (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT NOT NULL,
    grade VARCHAR(20) NOT NULL
);

INSERT INTO students (name, age, grade)
SELECT 'Ravi', 20, 'A'
WHERE NOT EXISTS (
    SELECT 1 FROM students
    WHERE name = 'Ravi' AND grade = 'A'
);

INSERT INTO students (name, age, grade)
SELECT 'Priya', 21, 'B'
WHERE NOT EXISTS (
    SELECT 1 FROM students
    WHERE name = 'Priya' AND grade = 'B'
);

INSERT INTO students (name, age, grade)
SELECT 'Arun', 22, 'A'
WHERE NOT EXISTS (
    SELECT 1 FROM students
    WHERE name = 'Arun' AND grade = 'A'
);