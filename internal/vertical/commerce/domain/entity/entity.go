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

func (c *Commerce) NewCommerce(name, address string, branches []*Branch) *Commerce {
	return &Commerce{
		ID:       uuid.New().String(),
		Name:     name,
		Address:  address,
		Branches: branches,
	}
}

func (b *Branch) NewBranch(commerceID, name, address string) *Branch {
	return &Branch{
		ID:         uuid.New().String(),
		CommerceID: commerceID,
		Name:       name,
		Address:    address,
	}
}
