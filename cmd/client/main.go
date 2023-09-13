package main

import (
	"context"
	"fmt"

	c "github.com/nmramorov/gophkeeper/internal/api/client"
)

var (
	Version   string
	BuildTime string
)

func main() {
	fmt.Printf("version=%s, time=%s\n", Version, BuildTime)
	parent := context.Background()
	client := c.NewClientApp()
	for {
		client.Run(parent)
	}
}
