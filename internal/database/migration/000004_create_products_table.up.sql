CREATE TABLE products (
    product_id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category_id INT REFERENCES categories(category_id) ON DELETE SET NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL CHECK (price > 0),
    quantity INT NOT NULL DEFAULT 0 CHECK (quantity >= 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);
