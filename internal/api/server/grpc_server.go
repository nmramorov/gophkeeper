package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nmramorov/gophkeeper/internal/db"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
	"google.golang.org/grpc"
)

type StorageServer struct {
	gctx context.Context
	pb.UnimplementedStorageServer
	Storage db.DBAPI
}

func (s StorageServer) Run(parent context.Context) {
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}
	globalCtx, globalCancel := context.WithCancel(parent)
	defer globalCancel()
	server := grpc.NewServer()

	pb.RegisterStorageServer(server, &StorageServer{
		gctx: globalCtx,
	})

	fmt.Println("gRPC server starts working")
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

func mergeContext(a, b context.Context) (context.Context, context.CancelFunc) {
	// merge client and server contexts into one `mctx`
	// (client context will cancel if client disconnects)
	// (server context will cancel if service Ctrl-C'ed)

	mctx, mcancel := context.WithCancel(a) // will cancel if `a` cancels

	go func() {
		select {
		case <-mctx.Done(): // don't leak go-routine on clean gRPC run
		case <-b.Done():
			mcancel() // b canceled, so cancel mctx
		}
	}()

	return mctx, mcancel
}
