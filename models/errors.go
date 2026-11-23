package models

import "errors"

var (
	ErrNotFound    = errors.New("not found")
	ErrTimeout     = errors.New("timeout")
	ErrUnavailable = errors.New("service unavailable")
	ErrBadRequest  = errors.New("bad request")
)
