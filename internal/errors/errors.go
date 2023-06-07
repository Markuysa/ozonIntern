package errors

import "github.com/pkg/errors"

var (
	ErrAlreadyExists = errors.New("the url already exists")
	ErrUrlNotFound   = errors.New("the url not found")
)
