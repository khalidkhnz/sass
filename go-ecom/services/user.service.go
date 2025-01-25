package services

import (
	"context"

	"github.com/khalidkhnz/sass/go-ecom/schemas"

	"github.com/khalidkhnz/sass/go-ecom/types"
)



func InsertUserToDB(u *schemas.User) (*schemas.User, error) {
	query := `
		INSERT INTO users (email, password, name, phone_number)
		VALUES ($1, $2, $3, $4)
		RETURNING id, email, password, name, phone_number, created_at, updated_at`

	err := DATABASE.QueryRow(
		context.TODO(),
		query,
		u.Email,
		u.Password,
		u.Name,
		u.PhoneNumber,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.Name,
		&u.PhoneNumber,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}



func GetByEmail(UserEmail string) (*schemas.User, error) {
	query := `
		SELECT id, email, password, name, store_id, phone_number, address_id, created_at, updated_at 
		FROM users 
		WHERE email = $1`

	var user schemas.User
	err := DATABASE.QueryRow(
		context.TODO(),
		query,
		UserEmail,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.StoreId,
		&user.PhoneNumber,
		&user.AddressIds,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserById(UserId string) (*schemas.User, error) {
	query := `
		SELECT id, email, password, name, store_id, phone_number, address_id, created_at, updated_at 
		FROM users 
		WHERE id = $1`

	var user schemas.User
	err := DATABASE.QueryRow(
		context.TODO(),
		query,
		UserId,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.StoreId,
		&user.PhoneNumber,
		&user.AddressIds,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsersOfStoreByStoreId(StoreId string, filter types.CommonFilters) ([]*schemas.User, error) {
	// Set default values if filters are missing
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	offset := (filter.Page - 1) * filter.Limit

	// Build base query
	query := `
		SELECT id, email, password, name, store_id, phone_number, address_id, created_at, updated_at 
		FROM users 
		WHERE store_id = $1`

	// Only add search condition if search term provided
	if filter.Search != "" {
		query += `
		AND (
			LOWER(name) LIKE LOWER($4) OR
			LOWER(email) LIKE LOWER($4) OR 
			LOWER(phone_number) LIKE LOWER($4)
		)`
	}

	query += `
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3`

	// Build args slice based on whether search is included
	var args []interface{}
	args = append(args, StoreId, filter.Limit, offset)
	if filter.Search != "" {
		args = append(args, "%"+filter.Search+"%")
	}

	rows, err := DATABASE.Query(
		context.TODO(),
		query,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*schemas.User
	for rows.Next() {
		var user schemas.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.Name,
			&user.StoreId,
			&user.PhoneNumber,
			&user.AddressIds,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if len(users) == 0 {
		return []*schemas.User{}, nil
	}

	return users, nil
}

func GetAllUsers(filter types.CommonFilters) ([]*schemas.User, error) {
	offset := (filter.Page - 1) * filter.Limit
	query := `
		SELECT id, email, password, name, store_id, phone_number, address_id, created_at, updated_at 
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := DATABASE.Query(
		context.TODO(),
		query,
		filter.Limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*schemas.User
	for rows.Next() {
		var user schemas.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.Name,
			&user.StoreId,
			&user.PhoneNumber,
			&user.AddressIds,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if len(users) == 0 {
		return []*schemas.User{}, nil
	}

	return users, nil
}