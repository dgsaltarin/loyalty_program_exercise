package entity

import (
	"github.com/google/uuid"
)

type Commerce struct {
	ID       string
	Name     string
	Address  string
	Branches []*Branch
}

type Branch struct {
	ID         string
	CommerceID string
	Name       string
	Address    string
}

func NewCommerce(name, address string) *Commerce {
	return &Commerce{
		ID:      uuid.New().String(),
		Name:    name,
		Address: address,
	}
}

func NewBranch(commerceID, name, address string) *Branch {
	return &Branch{
		ID:         uuid.New().String(),
		CommerceID: commerceID,
		Name:       name,
		Address:    address,
	}
}
