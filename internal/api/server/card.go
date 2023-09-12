package server

import (
	"context"
	"fmt"

	"github.com/nmramorov/gophkeeper/internal/api/models"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
)

func (s *StorageServer) SaveCard(ctx context.Context, in *pb.SaveBankCardDataRequest) (*pb.SaveBankCardDataResponse, error) {
	mctx, mcancel := mergeContext(ctx, s.gctx)

	defer mcancel()

	var response pb.SaveBankCardDataResponse

	validationError := s.ValidateRequest(mctx, in.Token)
	response.Error = validationError

	newCard := models.BankCardData{
		UUID:       in.Data.Uuid,
		Number:     in.Data.Number,
		Owner:      in.Data.Owner,
		ExpiresAt:  in.Data.ExpiresAt,
		SecretCode: in.Data.SecretCode,
		PinCode:    in.Data.PinCode,
		Meta:       in.Data.Meta.Content,
	}
	err := s.Storage.SaveCard(mctx, newCard)
	if err != nil {
		response.Error = fmt.Sprintf("internal server error for data %s", in.Data.Uuid)
		return &response, nil
	}
	return &response, nil
}

func (s *StorageServer) LoadCard(ctx context.Context, in *pb.LoadBankCardDataRequest) (*pb.LoadBankCardDataResponse, error) {
	mctx, mcancel := mergeContext(ctx, s.gctx)

	defer mcancel()

	var response pb.LoadBankCardDataResponse

	validationError := s.ValidateRequest(mctx, in.Token)
	response.Error = validationError

	card, err := s.Storage.LoadCard(mctx, in.Uuid)
	if err != nil {
		response.Error = fmt.Sprintf("internal server error for data %s", in.Uuid)
		return &response, nil
	}
	response.Data = &pb.BankCardData{
		Uuid:       card.UUID,
		Number:     card.Number,
		Owner:      card.Owner,
		ExpiresAt:  card.ExpiresAt,
		SecretCode: card.SecretCode,
		PinCode:    card.PinCode,
		Meta: &pb.Meta{
			Content: card.Meta,
		},
	}
	return &response, nil
}
