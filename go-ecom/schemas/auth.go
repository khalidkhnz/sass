package schemas

import (
	"context"
	"time"

	"github.com/khalidkhnz/sass/go-ecom/services"
)

type User struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}


func (u *User) InsertToDB() *User {
	query := `
		INSERT INTO users (email, password, first_name, last_name, phone_number)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, email, password, first_name, last_name, phone_number, created_at, updated_at`

	err := services.DATABASE.QueryRow(
		context.TODO(),
		query,
		u.Email,
		u.Password,
		u.FirstName,
		u.LastName,
		u.PhoneNumber,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.FirstName,
		&u.LastName,
		&u.PhoneNumber,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return nil
	}

	return u
}


const CreateUserTable = `
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    phone_number VARCHAR(20),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
`
