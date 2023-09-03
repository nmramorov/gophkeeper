package db

import "errors"

var ErrInMemoryDB = errors.New("InMemoryDB internal error")
var ErrContextTimeout = errors.New("context timeout called")
var ErrDatabaseUnreachable = errors.New("database unreachable")
