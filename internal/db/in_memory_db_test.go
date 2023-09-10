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
		Users:       sync.Map{},
		Credentials: sync.Map{},
	}
	suite.TestDB.Users.Store("test uuid", models.User{
		UUID:     "test uuid",
		Login:    "initial login",
		Password: "initial password",
	})
	suite.TestDB.Credentials.Store("initial UUID", models.CredentialsData{
		UUID:     "initial UUID",
		UserID:   "test uuid",
		Login:    "initial login",
		Password: "initial password",
		Meta:     "initial Meta",
	})
	suite.TestDB.Texts.Store("initial UUID", models.TextData{
		UUID: "initial UUID",
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
		UserID:   "test uuid",
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
	result, err := suite.TestDB.LoadUser(ctx, "test uuid")
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), models.User{
		UUID:     "test uuid",
		Login:    "initial login",
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

func (suite *InMemoryDBTestSuite) TestLoadUserNotFound() {
	ctx := context.Background()
	result, err := suite.TestDB.LoadUser(ctx, "some login")
	require.Error(suite.T(), ErrUserNotFound, err)
	require.Equal(suite.T(), result, models.User{})
}

func (suite *InMemoryDBTestSuite) TestSaveUser() {
	ctx := context.Background()
	newData := models.User{
		UUID:     "new uuid",
		Login:    "test",
		Password: "test",
		Token:    "test token",
	}
	err := suite.TestDB.SaveUser(ctx, newData)
	require.NoError(suite.T(), err)
	user, err := suite.TestDB.LoadUser(ctx, "new uuid")
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "test", user.Login)
	require.Equal(suite.T(), "new uuid", user.UUID)
	require.Equal(suite.T(), "test token", user.Token)
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
