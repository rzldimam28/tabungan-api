package domain

import "errors"

var (
	ErrAccountNotFound = errors.New("error account not found")
	ErrMutationsNotFound = errors.New("error mutations not found")
	ErrNotEnoughBalance = errors.New("error not enough balance")
)