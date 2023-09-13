package server

import (
	"context"
	"fmt"

	"github.com/nmramorov/gophkeeper/internal/api/models"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
)

func (s *StorageServer) SaveBinary(ctx context.Context, in *pb.SaveBinaryDataRequest) (*pb.SaveBinaryDataResponse, error) {
	var response pb.SaveBinaryDataResponse

	validationError := s.ValidateRequest(ctx, in.Token)
	response.Error = validationError

	newBytes := models.BinaryData{
		UUID: in.Data.Uuid,
		Data: in.Data.Data,
		Meta: in.Data.Meta.Content,
	}
	err := s.Storage.SaveBinary(ctx, newBytes)
	if err != nil {
		response.Error = fmt.Sprintf("internal server error for data %s", in.Data.Uuid)
		return &response, nil
	}
	return &response, nil
}

func (s *StorageServer) LoadBinary(ctx context.Context, in *pb.LoadBinaryDataRequest) (*pb.LoadBinaryDataResponse, error) {
	var response pb.LoadBinaryDataResponse

	validationError := s.ValidateRequest(ctx, in.Token)
	response.Error = validationError

	bin, err := s.Storage.LoadBinary(ctx, in.Uuid)
	if err != nil {
		response.Error = fmt.Sprintf("internal server error for data %s", in.Uuid)
		return &response, nil
	}
	response.Data = &pb.BinaryData{
		Uuid: bin.UUID,
		Data: bin.Data,
		Meta: &pb.Meta{
			Content: bin.Meta,
		},
	}
	return &response, nil
}
