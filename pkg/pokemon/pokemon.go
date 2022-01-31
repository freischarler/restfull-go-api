package pokemon

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Pokemon of the system.
type Pokemon struct {
	ID           int       `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Element      string    `json:"element,omitempty"`
	Health       int       `json:"health,omitempty"`
	Picture      string    `json:"picture,omitempty"`
	Password     string    `json:"password,omitempty"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

func (u *Pokemon) HashPassword() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(passwordHash)

	return nil
}

func (u Pokemon) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	return err == nil
}
