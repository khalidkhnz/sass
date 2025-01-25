package services

import (
	"context"

	"github.com/khalidkhnz/sass/go-ecom/schemas"
)

func InsertStoreToDB(s *schemas.Store) (*schemas.Store, error) {
	query := `
		INSERT INTO stores (store_name, store_email, password, store_phone_number, owner_name, owner_phone_number)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, store_name, store_email, password, store_phone_number, owner_name, owner_phone_number, created_at, updated_at`

	err := DATABASE.QueryRow(
		context.TODO(),
		query,
		s.StoreName,
		s.StoreEmail,
		s.Password,
		s.StorePhoneNumber,
		s.OwnerName,
		s.OwnerPhoneNumber,
	).Scan(
		&s.ID,
		&s.StoreName,
		&s.StoreEmail,
		&s.Password,
		&s.StorePhoneNumber,
		&s.OwnerName,
		&s.OwnerPhoneNumber,
		&s.CreatedAt,
		&s.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return s, nil
}