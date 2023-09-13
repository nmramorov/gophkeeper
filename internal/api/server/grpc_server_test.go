package server

import (
	"context"
	"log"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/nmramorov/gophkeeper/internal/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/credentials/insecure"
)

func TestStorageServer(t *testing.T) {
	ctx := context.Background()
	go Run(ctx)
	conn, err := grpc.DialContext(ctx, ":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewStorageClient(conn)

	// register new user
	response, errResp := client.RegisterUser(ctx, &pb.RegisterUserRequest{
		User: &pb.User{
			Login:    "stray228",
			Password: "gachipower",
		},
	})
	require.NoError(t, errResp)
	require.Equal(t, "", response.Error)

}
