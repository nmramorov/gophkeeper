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

	// now try login with user which not exists
	resp, err := client.LoginUser(ctx, &pb.LoginUserRequest{
		User: &pb.User{
			Login:    "stray",
			Password: "gachipower",
		},
	})
	require.Error(t, ErrUserNotFound, err)

	// then we login with proper user
	resp, err = client.LoginUser(ctx, &pb.LoginUserRequest{
		User: &pb.User{
			Login:    "stray228",
			Password: "gachipower",
		},
	})
	require.NoError(t, err)
	require.Equal(t, "", resp.Error)
	require.Equal(t, "stray228/gachipower/salt", resp.Token)
	token := resp.Token

	// let's save some data
	r, errResp := client.SaveCredentials(ctx, &pb.SaveCredentialsDataRequest{
		Token: token,
		Data: &pb.CredentialsData{
			Uuid:     "new uuid",
			Login:    "twitch login",
			Password: "twitch password",
			Meta: &pb.Meta{
				Content: "twitch metadata",
			},
		},
	})
	require.NoError(t, errResp)
	require.Equal(t, "", r.Error)

	// then let's load this data
	loadResp, err := client.LoadCredentials(ctx, &pb.LoadCredentialsDataRequest{
		Token: token,
		Uuid:  "new uuid",
	})
	require.NoError(t, err)
	require.EqualValues(t, pb.CredentialsData{
		Uuid:     "new uuid",
		Login:    "twitch login",
		Password: "twitch password",
		Meta: &pb.Meta{
			Content: "twitch metadata",
		},
	}, *loadResp.Data)
}
