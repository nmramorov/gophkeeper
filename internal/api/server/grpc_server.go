package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/nmramorov/gophkeeper/internal/db"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
	"google.golang.org/grpc"
)

type StorageServer struct {
	pb.UnimplementedStorageServer
	Storage db.DBAPI
}

func Run(parent context.Context) {
	idleConnsClosed := make(chan struct{})
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	go func() {
		<-sigint
		fmt.Println("received shutdown signal")
		fmt.Println("running gracefull shutdown")
		server.GracefulStop()
		fmt.Println("server shutted down")
		close(idleConnsClosed)
	}()

	pb.RegisterStorageServer(server, &StorageServer{
		Storage: &db.InMemoryDB{},
	})

	fmt.Println("gRPC server starts working")
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
	<-idleConnsClosed
}
