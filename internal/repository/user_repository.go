package repository

import (
	"context"
	"fmt"
	"module/internal/model"
	"module/internal/repository/queries"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Login     string
	Password  string
	FirstName string
	LastName  string
	Email     string
	Role      model.RoleType
}

type UserRepository interface {
	Create(ctx context.Context, input CreateUserInput) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *userRepository {
	return &userRepository{
		db: db,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (r *userRepository) Create(ctx context.Context, input CreateUserInput) (*model.User, error) {
	if !input.Role.Valid() {
		return nil, fmt.Errorf("invalid role: %s", input.Role)
	}

	hash, err := HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	uuid := uuid.New()

	var user model.User
	err = r.db.QueryRow(ctx, queries.CreateUser,
		uuid,
		input.Login,
		hash,
		input.FirstName,
		input.LastName,
		input.Email,
		input.Role,
	).Scan(
		&user.UserID,
		&user.Login,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow(ctx, queries.GetUserByEmail, email).Scan(
		&user.UserID,
		&user.Login,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow(ctx, queries.GetUserByID, id).Scan(
		&user.UserID,
		&user.Login,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
