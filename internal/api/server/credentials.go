package server

import (
	"context"
	"fmt"

	"github.com/nmramorov/gophkeeper/internal/api/models"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
)

func (s *StorageServer) SaveCredentials(ctx context.Context, in *pb.SaveCredentialsDataRequest) (*pb.SaveCredentialsDataResponse, error) {
	mctx, mcancel := mergeContext(ctx, s.gctx)

	defer mcancel()
	
	var response pb.SaveCredentialsDataResponse

	validationError := s.ValidateRequest(mctx, in.Token)
	response.Error = validationError

	newCredentials := models.CredentialsData{
		UUID:     in.Data.Uuid,
		Login:    in.Data.Login,
		Password: in.Data.Password,
	}
	err := s.Storage.SaveCredentials(mctx, newCredentials)
	if err != nil {
		response.Error = fmt.Sprintf("internal server error for data %s", in.Data.Uuid)
		return &response, nil
	}
	return &response, nil
}

func (s *StorageServer) LoadCredentials(ctx context.Context, in *pb.LoadCredentialsDataRequest) (*pb.LoadCredentialsDataResponse, error) {
	mctx, mcancel := mergeContext(ctx, s.gctx)

	defer mcancel()

	var response pb.LoadCredentialsDataResponse

	validationError := s.ValidateRequest(mctx, in.Token)
	response.Error = validationError

	credentials, err := s.Storage.LoadCredentials(mctx, in.Uuid)
	if err != nil {
		response.Error = fmt.Sprintf("internal server error for data %s", in.Uuid)
		return &response, nil
	}
	response.Data = &pb.CredentialsData{
		Uuid:     credentials.UUID,
		Login:    credentials.Login,
		Password: credentials.Password,
		Meta: &pb.Meta{
			Content: credentials.Meta,
		},
	}
	return &response, nil
}
