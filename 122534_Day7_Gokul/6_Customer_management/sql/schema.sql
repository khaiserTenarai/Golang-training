CREATE TABLE IF NOT EXISTS customer (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(150) NOT NULL,
    email      VARCHAR(150) NOT NULL UNIQUE,
    phone      VARCHAR(20)  NOT NULL,
    address    VARCHAR(255),
    created_at TIMESTAMP    NOT NULL DEFAULT NOW()
);



