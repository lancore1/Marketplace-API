CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    login VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100) NOT NULL UNIQUE,
    role_id INT NOT NULL REFERENCES user_roles(role_id) ON DELETE RESTRICT
);
