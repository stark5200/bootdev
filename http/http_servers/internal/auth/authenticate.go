package auth

import (
	//"context"
	//"database/sql"
	//"errors"
	"golang.org/x/crypto/bcrypt"
)

/*
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Authenticator struct {
	db *sql.DB
}

func NewAuthenticator(db *sql.DB) *Authenticator {
	return &Authenticator{db: db}
}

func (a *Authenticator) Authenticate(ctx context.Context, email, password string) error {
	var hashedPassword string
	err := a.db.QueryRowContext(ctx, "SELECT password_hash FROM users WHERE email = $1", email).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrInvalidCredentials
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return ErrInvalidCredentials
	}

	return nil
}
*/

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}