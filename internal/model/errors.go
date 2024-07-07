package model

import "errors"

var (
	ErrAlreadyExists = errors.New("user with this passport already exists")
	ErrNotFound      = errors.New("user not found")
)
