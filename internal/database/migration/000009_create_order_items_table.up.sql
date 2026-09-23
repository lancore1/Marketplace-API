CREATE TABLE order_items (
    order_id BIGINT REFERENCES orders(order_id) ON DELETE CASCADE,
    product_id BIGINT REFERENCES products(product_id) ON DELETE RESTRICT,
    quantity INT NOT NULL CHECK (quantity > 0),
    price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    PRIMARY KEY (order_id, product_id)
);
