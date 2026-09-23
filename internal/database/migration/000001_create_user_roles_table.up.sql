
CREATE TYPE role_type AS ENUM ('admin', 'user');
CREATE TABLE user_roles (
    role_id SERIAL PRIMARY KEY,
    role role_type NOT NULL UNIQUE
);
