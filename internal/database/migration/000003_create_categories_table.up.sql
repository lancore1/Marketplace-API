CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    title VARCHAR(30) NOT NULL UNIQUE
);
