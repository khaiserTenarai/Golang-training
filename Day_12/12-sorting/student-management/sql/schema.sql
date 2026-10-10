CREATE TABLE IF NOT EXISTS students (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    grade VARCHAR(20) NOT NULL
);

INSERT INTO students (name, grade) VALUES
('Ravi', 'A'), ('Priya', 'B'), ('Ganesh', 'A')
ON CONFLICT DO NOTHING;
