package server

import (
	"github.com/nmramorov/gophkeeper/internal/db"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
)


type StorageServer struct {
	pb.UnimplementedStorageServer
	Storage db.DBAPI
}
