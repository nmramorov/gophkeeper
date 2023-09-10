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
		UUID:   "initial UUID",
		UserID: "test uuid",
		Data:   "test text data",
		Meta:   "test text meta",
	})
	suite.TestDB.Binaries.Store("initial UUID", models.BinaryData{
		UUID:   "initial UUID",
		UserID: "test uuid",
		Data:   []byte{255, 255, 255},
		Meta:   "test binary meta",
	})
	suite.TestDB.Cards.Store("initial UUID", models.BankCardData{
		UUID:       "initial UUID",
		UserID:     "test uuid",
		Number:     "1234 5678 9101 2134",
		Owner:      "Jeff Jeff",
		ExpiresAt:  "Never",
		SecretCode: "228",
		PinCode:    "0451",
		Meta:       "test card meta",
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

func (suite *InMemoryDBTestSuite) TestSaveTextSuccess() {
	ctx := context.Background()
	newData := models.TextData{
		UUID:   "test text uuid",
		Data:   "test",
		UserID: "test uuid",
		Meta:   "test",
	}
	err := suite.TestDB.SaveText(ctx, newData)
	require.NoError(suite.T(), err)
}

func (suite *InMemoryDBTestSuite) TestSaveTextContextTimeout() {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Millisecond)
	defer cancel()
	time.Sleep(time.Duration(20) * time.Millisecond)
	newData := models.TextData{
		UUID:   "test text uuid",
		Data:   "test",
		UserID: "test uuid",
		Meta:   "test",
	}
	err := suite.TestDB.SaveText(newCtx, newData)
	require.ErrorIs(suite.T(), err, ErrContextTimeout)
}

func (suite *InMemoryDBTestSuite) TestLoadText() {
	ctx := context.Background()
	result, err := suite.TestDB.LoadText(ctx, "initial UUID")
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), models.TextData{
		UUID:   "initial UUID",
		UserID: "test uuid",
		Data:   "test text data",
		Meta:   "test text meta",
	}, result)
}

func (suite *InMemoryDBTestSuite) TestLoadTextContextTimeout() {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Millisecond)
	defer cancel()
	time.Sleep(time.Duration(30) * time.Millisecond)
	result, err := suite.TestDB.LoadText(newCtx, "initial UUID")
	require.ErrorIs(suite.T(), err, ErrContextTimeout)
	require.Equal(suite.T(), models.TextData{}, result)
}

func (suite *InMemoryDBTestSuite) TestSaveBinarySuccess() {
	ctx := context.Background()
	newData := models.BinaryData{
		UUID:   "test binary uuid",
		Data:   []byte{0, 0, 0},
		UserID: "test uuid",
		Meta:   "test",
	}
	err := suite.TestDB.SaveBinary(ctx, newData)
	require.NoError(suite.T(), err)
}

func (suite *InMemoryDBTestSuite) TestSaveBinaryContextTimeout() {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Millisecond)
	defer cancel()
	time.Sleep(time.Duration(20) * time.Millisecond)
	newData := models.BinaryData{
		UUID:   "test text uuid",
		Data:   []byte{},
		UserID: "test uuid",
		Meta:   "test",
	}
	err := suite.TestDB.SaveBinary(newCtx, newData)
	require.ErrorIs(suite.T(), err, ErrContextTimeout)
}

func (suite *InMemoryDBTestSuite) TestLoadBinary() {
	ctx := context.Background()
	result, err := suite.TestDB.LoadBinary(ctx, "initial UUID")
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), models.BinaryData{
		UUID:   "initial UUID",
		UserID: "test uuid",
		Data:   []byte{255, 255, 255},
		Meta:   "test binary meta",
	}, result)
}

func (suite *InMemoryDBTestSuite) TestLoadBinaryContextTimeout() {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Millisecond)
	defer cancel()
	time.Sleep(time.Duration(30) * time.Millisecond)
	result, err := suite.TestDB.LoadBinary(newCtx, "initial UUID")
	require.ErrorIs(suite.T(), err, ErrContextTimeout)
	require.Equal(suite.T(), models.BinaryData{}, result)
}

func (suite *InMemoryDBTestSuite) TestSaveCardSuccess() {
	ctx := context.Background()
	newData := models.BankCardData{
		UUID:       "test text uuid",
		UserID:     "test uuid",
		Number:     "test number",
		Owner:      "test owner",
		ExpiresAt:  "test expiration",
		SecretCode: "test secret",
		PinCode:    "test pin",
		Meta:       "test",
	}
	err := suite.TestDB.SaveCard(ctx, newData)
	require.NoError(suite.T(), err)
}

func (suite *InMemoryDBTestSuite) TestSaveCardContextTimeout() {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Millisecond)
	defer cancel()
	time.Sleep(time.Duration(20) * time.Millisecond)
	newData := models.BankCardData{}
	err := suite.TestDB.SaveCard(newCtx, newData)
	require.ErrorIs(suite.T(), err, ErrContextTimeout)
}

func (suite *InMemoryDBTestSuite) TestLoadCard() {
	ctx := context.Background()
	result, err := suite.TestDB.LoadCard(ctx, "initial UUID")
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), models.BankCardData{
		UUID:       "initial UUID",
		UserID:     "test uuid",
		Number:     "1234 5678 9101 2134",
		Owner:      "Jeff Jeff",
		ExpiresAt:  "Never",
		SecretCode: "228",
		PinCode:    "0451",
		Meta:       "test card meta",
	}, result)
}

func (suite *InMemoryDBTestSuite) TestLoadCardContextTimeout() {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Millisecond)
	defer cancel()
	time.Sleep(time.Duration(30) * time.Millisecond)
	result, err := suite.TestDB.LoadCard(newCtx, "initial UUID")
	require.ErrorIs(suite.T(), err, ErrContextTimeout)
	require.Equal(suite.T(), models.BankCardData{}, result)
}

func (suite *InMemoryDBTestSuite) TestFindUser() {
	ctx := context.Background()
	result, err := suite.TestDB.FindUser(ctx, "initial login", "initial password")
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), models.User{
		UUID:     "test uuid",
		Login:    "initial login",
		Password: "initial password",
	}, result)
}

func (suite *InMemoryDBTestSuite) TestFindUserContextTimeout() {
	ctx := context.Background()
	newCtx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Millisecond)
	defer cancel()
	time.Sleep(time.Duration(50) * time.Millisecond)
	result, err := suite.TestDB.FindUser(newCtx, "initial login", "initial password")
	require.ErrorIs(suite.T(), err, ErrContextTimeout)
	require.Equal(suite.T(), models.User{}, result)
}

func (suite *InMemoryDBTestSuite) TestFindUserNotFound() {
	ctx := context.Background()
	result, err := suite.TestDB.FindUser(ctx, "fake login", "fake password")
	require.Error(suite.T(), ErrUserNotFound, err)
	require.Equal(suite.T(), result, models.User{})
}

func TestInMemoryDBTestSuite(t *testing.T) {
	suite.Run(t, new(InMemoryDBTestSuite))
}
