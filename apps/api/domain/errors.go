package domain

import "errors"

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrCustomerExists   = errors.New("customer already exists")
	ErrIncompleteInfo   = errors.New("incomplete info")
	ErrHistoryNotFound  = errors.New("history not found")
)
