package server

import (
	"context"
	"sync"
	"testing"

	"github.com/nmramorov/gophkeeper/internal/api/models"
	"github.com/nmramorov/gophkeeper/internal/db"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CardTestSuite struct {
	suite.Suite

	Server StorageServer
	ctx    context.Context
}

func (suite *CardTestSuite) SetupTest() {
	suite.Server = StorageServer{
		gctx: context.Background(),
		Storage: &db.InMemoryDB{
			Users:    sync.Map{},
			Binaries: sync.Map{},
		},
	}
	suite.ctx = context.Background()
	sessions := []models.Session{
		{
			UUID:  "test session 1",
			Token: "initial login/initial password/salt",
		},
		{
			UUID:  "test session 2",
			Token: "initial login/initial password/test salt 2",
		},
	}
	err := suite.Server.Storage.SaveUser(suite.ctx, models.User{
		Login:    "initial login",
		Password: "initial password",
		Sessions: sessions,
	})
	if err != nil {
		suite.T().Errorf("Error setup - saving user: %e", err)
	}
	err = suite.Server.Storage.SaveCard(suite.ctx, models.BankCardData{
		UUID:       "initial UUID",
		Number:     "test number",
		Owner:      "test owner",
		ExpiresAt:  "never",
		SecretCode: "secret code",
		PinCode:    "test pin",
		Meta:       "initial Meta",
	})
	if err != nil {
		suite.T().Errorf("Error setup - saving text: %e", err)
	}
}

func (suite *CardTestSuite) TestSaveCardInvalidToken() {
	req := &pb.SaveBankCardDataRequest{
		Token: "initial login/initial password/invalid salt",
		Data: &pb.BankCardData{
			Uuid: "test uuid",
			Meta: &pb.Meta{},
		},
	}
	resp, err := suite.Server.SaveCard(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "authorization error: invalid token", resp.Error)
}

func (suite *CardTestSuite) TestSaveCardSuccess() {
	req := &pb.SaveBankCardDataRequest{
		Token: "initial login/initial password/salt",
		Data: &pb.BankCardData{
			Uuid:       "test UUID",
			Number:     "test number 1",
			Owner:      "test owner 1",
			ExpiresAt:  "never ever",
			SecretCode: "secret code 1",
			PinCode:    "test pin 1",
			Meta: &pb.Meta{
				Content: "test content",
			},
		},
	}
	resp, err := suite.Server.SaveCard(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
}

func (suite *CardTestSuite) TestLoadCardSuccess() {
	req := &pb.LoadBankCardDataRequest{
		Token: "initial login/initial password/salt",
		Uuid:  "initial UUID",
	}
	resp, err := suite.Server.LoadCard(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
	require.Equal(suite.T(), "initial UUID", resp.Data.Uuid)
	require.Equal(suite.T(), "test number", resp.Data.Number)
	require.Equal(suite.T(), "test owner", resp.Data.Owner)
	require.Equal(suite.T(), "never", resp.Data.ExpiresAt)
	require.Equal(suite.T(), "secret code", resp.Data.SecretCode)
	require.Equal(suite.T(), "test pin", resp.Data.PinCode)
	require.Equal(suite.T(), "initial Meta", resp.Data.Meta.Content)
}

func (suite *CardTestSuite) TestLoadCardAuthError() {
	req := &pb.LoadBankCardDataRequest{
		Token: "initial login/initial password/wrong salt",
		Uuid:  "initial UUID",
	}
	resp, err := suite.Server.LoadCard(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "authorization error: invalid token", resp.Error)
}

func (suite *CardTestSuite) TestLoadCardNoSuchID() {
	req := &pb.LoadBankCardDataRequest{
		Token: "initial login/initial password/salt",
		Uuid:  "wrong UUID",
	}
	resp, err := suite.Server.LoadCard(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "internal server error for data wrong UUID", resp.Error)
}

func TestCardTestSuite(t *testing.T) {
	suite.Run(t, new(CardTestSuite))
}
