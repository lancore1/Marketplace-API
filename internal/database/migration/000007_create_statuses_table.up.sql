CREATE TYPE status_type AS ENUM ('pending', 'paid', 'shipped', 'delivered', 'cancelled');

CREATE TABLE statuses (
    status_id SERIAL PRIMARY KEY,
    status status_type NOT NULL
);
