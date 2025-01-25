package services

import (
	"context"

	"github.com/khalidkhnz/sass/go-ecom/schemas"
)


func GetAllAddressesOfUser(UserId string) ([]*schemas.Address, error) {
	// First get the user to get their address IDs
	user, err := GetUserById(UserId)
	if err != nil {
		return nil, err
	}

	// Then get all addresses by those IDs
	return GetAddressesByIds(user.AddressIds)
}


func GetAddressesByIds(AddressIds []string) ([]*schemas.Address, error) {
	if len(AddressIds) == 0 {
		return []*schemas.Address{}, nil
	}

	query := `
		SELECT id, street, city, state, country, postal_code, created_at, updated_at 
		FROM addresses 
		WHERE id = ANY($1)`

	rows, err := DATABASE.Query(
		context.TODO(),
		query,
		AddressIds,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []*schemas.Address
	for rows.Next() {
		var address schemas.Address
		err := rows.Scan(
			&address.ID,
			&address.Street,
			&address.City,
			&address.State,
			&address.Country,
			&address.PostalCode,
			&address.CreatedAt,
			&address.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, &address)
	}

	if len(addresses) == 0 {
		return []*schemas.Address{}, nil
	}

	return addresses, nil
}


func GetAddressById(AddressId string) (*schemas.Address, error) {
	query := `
		SELECT id, street, city, state, country, postal_code, created_at, updated_at 
		FROM addresses 
		WHERE id = $1`

	var address schemas.Address
	err := DATABASE.QueryRow(
		context.TODO(),
		query,
		AddressId,
	).Scan(
		&address.ID,
		&address.Street,
		&address.City,
		&address.State,
		&address.Country,
		&address.PostalCode,
		&address.CreatedAt,
		&address.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &address, nil
}

func InsertAddressToDB(a *schemas.Address) (*schemas.Address, error) {
	query := `
		INSERT INTO addresses (street, city, state, country, postal_code)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, street, city, state, country, postal_code, created_at, updated_at`

	err := DATABASE.QueryRow(
		context.TODO(),
		query,
		a.Street,
		a.City,
		a.State,
		a.Country,
		a.PostalCode,
	).Scan(
		&a.ID,
		&a.Street,
		&a.City,
		&a.State,
		&a.Country,
		&a.PostalCode,
		&a.CreatedAt,
		&a.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return a, nil
}