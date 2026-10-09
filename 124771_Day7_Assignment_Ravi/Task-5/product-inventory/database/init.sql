CREATE TABLE IF NOT EXISTS public.products (
    id BIGSERIAL PRIMARY KEY,

    name VARCHAR(150) NOT NULL,

    description VARCHAR(500),

    price NUMERIC(12, 2) NOT NULL
        CHECK (price >= 0),

    stock_quantity INTEGER NOT NULL
        CHECK (stock_quantity >= 0),

    low_stock_threshold INTEGER NOT NULL DEFAULT 10
        CHECK (low_stock_threshold >= 0),

    created_at TIMESTAMP NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    updated_at TIMESTAMP NOT NULL
        DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO public.products (
    name,
    description,
    price,
    stock_quantity,
    low_stock_threshold
)
VALUES
(
    'Laptop',
    'Business laptop',
    75000,
    15,
    5
),
(
    'Mouse',
    'Wireless mouse',
    1200,
    8,
    10
),
(
    'Keyboard',
    'Mechanical keyboard',
    3500,
    20,
    5
),
(
    'Monitor',
    '24 inch monitor',
    15000,
    4,
    5
),
(
    'Headphones',
    'Wireless headphones',
    5000,
    25,
    10
);
