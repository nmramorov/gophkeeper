package server

import "errors"

var ErrDatabaseError = errors.New("database internal error")
var ErrInvalidToken = errors.New("invalid token")
var ErrUserAlreadyExists = errors.New("user already exists")
var ErrTimeout = errors.New("timeout error")
var ErrUserNotFound = errors.New("user not found")
