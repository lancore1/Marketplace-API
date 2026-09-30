package queries

const CreateUser = `
    INSERT INTO users (user_id, login, password_hash, first_name, last_name, email, role_id)
    SELECT $1, $2, $3, $4, $5, ur.role_id
    FROM user_roles ur
    WHERE ur.role = $6
    RETURNING user_id, login, password_hash, first_name, last_name, email, role_id, created_at, updated_at;
	`

const GetUserByEmail = `
    SELECT user_id, login, password_hash, first_name, last_name, email, role_id, created_at, updated_at
    FROM users
    WHERE email = $1;
	`
const GetUserByID = `
    SELECT user_id, login, password_hash, first_name, last_name, email, role_id, created_at, updated_at
    FROM users
    WHERE user_id = $1;
	`
