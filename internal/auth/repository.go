package auth

import (
	"errors"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/umdalecs/todo-list-api/internal/db"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

var ErrDuplicatedEmail = errors.New("email is already taken")

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateUser(userDto *RegisterDto, user *User) error {
	hashedPassword, err := HashPassword(userDto.Password)
	if err != nil {
		return err
	}

	err = pgxscan.Get(db.Ctx, r.db, user, `
		INSERT INTO users(name, email, password) 
		VALUES ($1,$2,$3) 
		RETURNING id, name, email, password`,
		userDto.Name, userDto.Email, hashedPassword)

	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			return ErrDuplicatedEmail
		}
		return err
	}

	return nil
}

func (r *AuthRepository) GetUserByEmail(email string, user *User) error {
	err := pgxscan.Get(db.Ctx, r.db, user, `
		SELECT id, password FROM users 
		WHERE email = $1`,
		email)

	if err != nil {
		return err
	}

	return nil
}
