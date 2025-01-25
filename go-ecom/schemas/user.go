package schemas

import (
	"time"
)

type User struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Name    	string    `json:"name"`
	StoreId		string 	  `json:"store_id"` // Foreign Key
	PhoneNumber string    `json:"phone_number"`
	AddressIds	[]string  `json:"address_ids"` // Foreign Key
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}




const CreateUserTable = `
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100),
    store_id UUID REFERENCES stores(id),
    phone_number VARCHAR(20),
    address_ids UUID[] DEFAULT '{}',
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
