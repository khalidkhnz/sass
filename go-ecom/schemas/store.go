package schemas

import (
	"time"
)

type Store struct {
	ID         		  string    `json:"id"`
	StoreName		  string	`json:"store_name"`
	StoreEmail   	  string    `json:"store_email"`
	Password     	  string    `json:"password"`
	StorePhoneNumber  string    `json:"store_phone_number"`
	OwnerName		  string    `json:"owner_name"`
	OwnerPhoneNumber  string    `json:"owner_phone_number"`
	CreatedAt    	  time.Time `json:"created_at"`
	AddressIds		  []string  `json:"address_ids"` // Foreign Key
	UpdatedAt     	  time.Time `json:"updated_at"`
}


const CreateStoreTable = `
CREATE TABLE IF NOT EXISTS stores (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    store_name VARCHAR(255) NOT NULL,
    store_email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    store_phone_number VARCHAR(20),
    owner_name VARCHAR(100) NOT NULL,
    owner_phone_number VARCHAR(20),
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

CREATE TRIGGER update_store_updated_at
    BEFORE UPDATE ON stores
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
`
