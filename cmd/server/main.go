package main

import (
	"context"

	"github.com/nmramorov/gophkeeper/internal/api/server"
)

func main() {
	parent := context.Background()
	server.Run(parent)
}
