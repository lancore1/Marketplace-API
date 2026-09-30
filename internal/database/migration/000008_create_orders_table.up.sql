CREATE TABLE orders (
    order_id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(user_id) ON DELETE SET NULL,
    status_id INT NOT NULL REFERENCES statuses(status_id) ON DELETE RESTRICT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);
