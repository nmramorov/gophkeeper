package server

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"

	"github.com/nmramorov/gophkeeper/internal/db"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterStorageServer(s, &StorageServer{
		Storage: &db.InMemoryDB{},
	})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestStorageServer(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewStorageClient(conn)

	response, errResp := client.RegisterUser(ctx, &pb.RegisterUserRequest{
		User: &pb.User{
			Login:    "stray228",
			Password: "gachipower",
		},
	})
	require.NoError(t, errResp)
	require.Equal(t, "", response.Error)
}
