package repository

import "cloud.google.com/go/firestore"

type AddressRepository interface {
	Get(id string) (Address, error)
	Create(address Address) (Address, error)
	Update(address Address) (Address, error)
}

type addressRepository struct {
	db *firestore.Client
}

func NewAddressRepository(db *firestore.Client) AddressRepository {
	return &addressRepository{db: db}
}

func (ref *addressRepository)Get(id string) (Address, error) {
	return Address{}, nil
}

func (ref *addressRepository)Create(address Address) (Address, error) {
	return Address{}, nil
}

func (ref *addressRepository)Update(address Address) (Address, error) {
	return Address{}, nil
}
