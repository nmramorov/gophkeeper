package server

// import (
// 	"context"
// 	"fmt"

// 	"github.com/nmramorov/gophkeeper/internal/api/models"
// 	pb "github.com/nmramorov/gophkeeper/internal/proto"
// )

// func (s *StorageServer) SaveText(ctx context.Context, in *pb.SaveTextDataRequest) (*pb.SaveTextDataResponse, error) {
// 	var response pb.SaveTextDataResponse

// 	validationError := s.ValidateRequest(ctx, in.Token)
// 	response.Error = validationError

// 	newText := models.TextData{
// 		UUID:     in.Data.Uuid,
// 		Data: in.Data.Data,
// 		Meta: in.Data.Meta.Content,
// 	}
// 	err := s.Storage.SaveText(ctx, newText)
// 	if err != nil {
// 		response.Error = fmt.Sprintf("internal server error for data %s", in.Data.Uuid)
// 		return &response, nil
// 	}
// 	return &response, nil
// }

// func (s *StorageServer) LoadText(ctx context.Context, in *pb.LoadTextDataRequest) (*pb.LoadTextDataResponse, error) {
// 	var response pb.LoadTextDataResponse

// 	validationError := s.ValidateRequest(ctx, in.Token)
// 	response.Error = validationError

// 	text, err := s.Storage.LoadText(ctx, in.Uuid)
// 	if err != nil {
// 		response.Error = fmt.Sprintf("internal server error for data %s", in.Uuid)
// 		return &response, nil
// 	}
// 	response.Data = &pb.TextData{
// 		Uuid:     text.UUID,
// 		Data: text.Data,
// 		Meta: &pb.Meta{
// 			Content: text.Meta,
// 		},
// 	}
// 	return &response, nil
// }
