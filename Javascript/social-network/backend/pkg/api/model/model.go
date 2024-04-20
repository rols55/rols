package model

import "errors"

var (
	// No records found
	ErrNotFound = errors.New("not found in database")
)
