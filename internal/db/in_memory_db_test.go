package db

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/nmramorov/gophkeeper/internal/api/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type InMemoryDBTestSuite struct {
	suite.Suite

	TestDB InMemoryDB
}

func (suite *InMemoryDBTestSuite) SetupTest() {
	suite.TestDB = InMemoryDB{
		users:       sync.Map{},
		credentials: sync.Map{},
	}
	suite.TestDB.users.Store("initial login", models.User{
		Login:    "initial login",
		Password: "initial password",
	})
	suite.TestDB.credentials.Store("initial UUID", models.CredentialsData{
		UUID:     "initial UUID",
		Login:    "initial login",
		Password: "initial password",
		Meta:     "initial Meta",
	})
}

func (suite *InMemoryDBTestSuite) TestSaveCredentialsSuccess() {
	ctx := context.Background()
	newData := models.CredentialsData{
		UUID:     "test uuid",
		Login:    "test",
		Password: "test",
		Meta:     "test",
	}
	err := suite.TestDB.SaveCredentials(ctx, newData)
	require.NoError(suite.T(), err)
}

func (suite *InMemoryDBTestSuite) TestSaveCredentialsContextTimeout() {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Millisecond)
	defer cancel()
	time.Sleep(time.Duration(20) * time.Millisecond)
	newData := models.CredentialsData{
		UUID:     "test uuid",
		Login:    "test",
		Password: "test",
		Meta:     "test",
	}
	err := suite.TestDB.SaveCredentials(newCtx, newData)
	require.ErrorIs(suite.T(), err, ErrContextTimeout)
}

func (suite *InMemoryDBTestSuite) TestLoadCredentials() {
	ctx := context.Background()
	result, err := suite.TestDB.LoadCredentials(ctx, "initial UUID")
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), models.CredentialsData{
		UUID:     "initial UUID",
		Login:    "initial login",
		Password: "initial password",
		Meta:     "initial Meta",
	}, result)
}

func (suite *InMemoryDBTestSuite) TestLoadCredentialsContextTimeout() {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Millisecond)
	defer cancel()
	time.Sleep(time.Duration(30) * time.Millisecond)
	result, err := suite.TestDB.LoadCredentials(newCtx, "initial UUID")
	require.ErrorIs(suite.T(), err, ErrContextTimeout)
	require.Equal(suite.T(), models.CredentialsData{
		UUID:     "",
		Login:    "",
		Password: "",
		Meta:     "",
	}, result)
}

func (suite *InMemoryDBTestSuite) TestPing() {
	ctx := context.Background()
	err := suite.TestDB.Ping(ctx)
	require.NoError(suite.T(), err)
}

func (suite *InMemoryDBTestSuite) TestLPingContextTimeout() {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Millisecond)
	defer cancel()
	time.Sleep(time.Duration(50) * time.Millisecond)
	err := suite.TestDB.Ping(newCtx)
	require.ErrorIs(suite.T(), err, ErrDatabaseUnreachable)
}

func (suite *InMemoryDBTestSuite) TestLoadUser() {
	ctx := context.Background()
	result, err := suite.TestDB.LoadUser(ctx, "initial login")
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), models.User{
		Login: "initial login",
		Password: "initial password",
	}, result)
}

func (suite *InMemoryDBTestSuite) TestLoadUserContextTimeout() {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Millisecond)
	defer cancel()
	time.Sleep(time.Duration(50) * time.Millisecond)
	result, err := suite.TestDB.LoadUser(newCtx, "initial login")
	require.ErrorIs(suite.T(), err, ErrContextTimeout)
	require.Equal(suite.T(), models.User{
		Login:    "",
		Password: "",
	}, result)
}

func (suite *InMemoryDBTestSuite) TestSaveUser() {
	ctx := context.Background()
	newData := models.User{
		Login:    "test",
		Password: "test",
	}
	err := suite.TestDB.SaveUser(ctx, newData)
	require.NoError(suite.T(), err)
}

func (suite *InMemoryDBTestSuite) TestSaveUserContextTimeout() {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Millisecond)
	defer cancel()
	time.Sleep(time.Duration(50) * time.Millisecond)
	newData := models.User{
		Login:    "test",
		Password: "test",
	}
	err := suite.TestDB.SaveUser(newCtx, newData)
	require.ErrorIs(suite.T(), err, ErrContextTimeout)
}

func TestInMemoryDBTestSuite(t *testing.T) {
	suite.Run(t, new(InMemoryDBTestSuite))
}
