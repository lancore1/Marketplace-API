CREATE TABLE cart_items (
    cart_id BIGINT REFERENCES carts(cart_id) ON DELETE CASCADE,
    product_id BIGINT REFERENCES products(product_id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),
    PRIMARY KEY (cart_id, product_id)
);
