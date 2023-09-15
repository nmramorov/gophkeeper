package client

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	pb "github.com/nmramorov/gophkeeper/internal/proto"
)

type ClientApp struct {
	Client pb.StorageClient
}

func NewClientApp() ClientApp {
	return ClientApp{}
}

func (c *ClientApp) Run(ctx context.Context) {
	saveCommand := flag.NewFlagSet("save", flag.ExitOnError)
	registerCommand := flag.NewFlagSet("register", flag.ExitOnError)
	loadCommand := flag.NewFlagSet("load", flag.ExitOnError)
	loginCommand := flag.NewFlagSet("login", flag.ExitOnError)

	credentialsFlags := flag.NewFlagSet("credentials", flag.ExitOnError)
	textFlags := flag.NewFlagSet("text", flag.ExitOnError)
	binaryFlags := flag.NewFlagSet("binary", flag.ExitOnError)
	bankCardFlags := flag.NewFlagSet("card", flag.ExitOnError)

	newUserLogin := registerCommand.String("login", "", "new user login")
	newUserPassword := registerCommand.String("password", "", "new user password")

	userLogin := loginCommand.String("login", "", "user login")
	userPassword := loginCommand.String("password", "", "user password")

	credLogin := credentialsFlags.String("login", "", "credential login")
	credPassword := credentialsFlags.String("password", "", "credential password")
	credAlias := credentialsFlags.String("alias", "", "credentials alias")
	credMeta := credentialsFlags.String("meta", "", "credential meta")
	credToken := credentialsFlags.String("token", "", "token")

	textData := textFlags.String("data", "", "text to save")
	textAlias := textFlags.String("alias", "", "texts alias")
	textMeta := textFlags.String("meta", "", "text meta")
	textToken := textFlags.String("token", "", "token")

	binData := binaryFlags.String("data", "", "bin data in encoded form to save")
	binAlias := binaryFlags.String("alias", "", "binary data alias")
	binMeta := binaryFlags.String("meta", "", "bin meta")
	binToken := binaryFlags.String("token", "", "token")

	cardNumber := bankCardFlags.String("number", "", "bank card number")
	cardOwner := bankCardFlags.String("owner", "", "bank card owner")
	cardExpires := bankCardFlags.String("expires at", "", "bank card expiration date")
	cardSecret := bankCardFlags.String("secret key", "", "bank card s ecret key")
	cardPIN := bankCardFlags.String("pin", "", "bank card PIN code")
	cardAlias := bankCardFlags.String("alias", "", "cards alias")
	cardMeta := bankCardFlags.String("meta", "", "bank card meta")
	cardToken := bankCardFlags.String("token", "", "token")

	if len(os.Args) < 2 {
		fmt.Println("save/register/load/login subcommand required")
		os.Exit(1)
	}
	if len(os.Args) < 3 {
		fmt.Println("credentials/text/binary/ subcommand required")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "save":
		saveCommand.Parse(os.Args[2:])
	case "load":
		loadCommand.Parse(os.Args[2:])
	case "login":
		loginCommand.Parse(os.Args[2:])
	case "register":
		registerCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	switch os.Args[2] {
	case "credentials":
		credentialsFlags.Parse(os.Args[3:])
	case "text":
		textFlags.Parse(os.Args[3:])
	case "binary":
		binaryFlags.Parse(os.Args[3:])
	case "card":
		bankCardFlags.Parse(os.Args[3:])
	}

	if saveCommand.Parsed() {
		if credentialsFlags.Parsed() {
			c.SaveCredentials(ctx, *credLogin, *credPassword, *credMeta, *credAlias, *credToken)
		} else if textFlags.Parsed() {
			c.SaveText(ctx, *textData, *textMeta, *textAlias, *textToken)
		} else if binaryFlags.Parsed() {
			c.SaveBinary(ctx, *binData, *binMeta, *binAlias, *binToken)
		} else {
			c.SaveCard(ctx, *cardNumber, *cardOwner, *cardExpires, *cardSecret, *cardPIN, *cardMeta, *cardAlias, *cardToken)
		}
	} else if loadCommand.Parsed() {
		if credentialsFlags.Parsed() {
			c.LoadCredentials(ctx, *credAlias, *credToken)
		} else if textFlags.Parsed() {
			c.LoadText(ctx, *textAlias, *textToken)
		} else if binaryFlags.Parsed() {
			c.LoadBinary(ctx, *binAlias, *binToken)
		} else {
			c.LoadCard(ctx, *cardAlias, *cardToken)
		}
	} else if loginCommand.Parsed() {
		c.Login(ctx, *userLogin, *userPassword)
	} else {
		c.Register(ctx, *newUserLogin, *newUserPassword)
	}
}

func (c *ClientApp) SaveText(parent context.Context, text, meta, alias, token string) error {
	if token == "" {
		fmt.Println("register user or authorize before saving data")
		return nil
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	_, err := c.Client.SaveText(ctx, &pb.SaveTextDataRequest{
		Token: token,
		Data: &pb.TextData{
			Uuid: alias,
			Data: text,
			Meta: &pb.Meta{
				Content: meta,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientApp) SaveBinary(parent context.Context, binData, meta, alias, token string) error {
	if token == "" {
		fmt.Println("register user or authorize before saving data")
		return nil
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	_, err := c.Client.SaveBinary(ctx, &pb.SaveBinaryDataRequest{
		Token: token,
		Data: &pb.BinaryData{
			Uuid: alias,
			Data: []byte(binData),
			Meta: &pb.Meta{
				Content: meta,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientApp) SaveCard(parent context.Context, cardNumber, cardOwner, cardExpires, cardSecret, cardPIN, meta, alias, token string) error {
	if token == "" {
		fmt.Println("register user or authorize before saving data")
		return nil
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	_, err := c.Client.SaveBankCard(ctx, &pb.SaveBankCardDataRequest{
		Token: token,
		Data: &pb.BankCardData{
			Uuid:       alias,
			Number:     cardNumber,
			Owner:      cardOwner,
			ExpiresAt:  cardExpires,
			SecretCode: cardSecret,
			PinCode:    cardPIN,
			Meta: &pb.Meta{
				Content: meta,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientApp) SaveCredentials(parent context.Context, login, password, meta, alias, token string) error {
	if token == "" {
		fmt.Println("register user or authorize before saving data")
		return nil
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	_, err := c.Client.SaveCredentials(ctx, &pb.SaveCredentialsDataRequest{
		Token: token,
		Data: &pb.CredentialsData{
			Uuid:     alias,
			Login:    login,
			Password: password,
			Meta: &pb.Meta{
				Content: meta,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientApp) LoadCredentials(parent context.Context, alias, token string) string {
	if token == "" {
		return "register user or authorize before saving data"
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	req, err := c.Client.LoadCredentials(ctx, &pb.LoadCredentialsDataRequest{
		Token: token,
		Uuid:  alias,
	})
	if err != nil {
		return err.Error()
	}
	if req.Error != "" {
		return req.Error
	}
	return fmt.Sprintln(req.Data)
}

func (c *ClientApp) LoadText(parent context.Context, alias, token string) string {
	if token == "" {
		return "register user or authorize before saving data"
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	req, err := c.Client.LoadText(ctx, &pb.LoadTextDataRequest{
		Token: token,
		Uuid:  alias,
	})
	if err != nil {
		return err.Error()
	}
	if req.Error != "" {
		return req.Error
	}
	return fmt.Sprintln(req.Data)
}

func (c *ClientApp) LoadBinary(parent context.Context, alias, token string) string {
	if token == "" {
		return "register user or authorize before saving data"
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	req, err := c.Client.LoadBinary(ctx, &pb.LoadBinaryDataRequest{
		Token: token,
		Uuid:  alias,
	})
	if err != nil {
		return err.Error()
	}
	if req.Error != "" {
		return req.Error
	}
	return fmt.Sprintln(req.Data)
}

func (c *ClientApp) LoadCard(parent context.Context, alias, token string) string {
	if token == "" {
		return "register user or authorize before saving data"
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	req, err := c.Client.LoadBankCard(ctx, &pb.LoadBankCardDataRequest{
		Token: token,
		Uuid:  alias,
	})
	if err != nil {
		return err.Error()
	}
	if req.Error != "" {
		return req.Error
	}
	return fmt.Sprintln(req.Data)
}

func (c *ClientApp) Login(parent context.Context, login, password string) string {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	req, err := c.Client.LoginUser(ctx, &pb.LoginUserRequest{
		User: &pb.User{
			Login:    login,
			Password: password,
		},
	})
	if err != nil {
		return err.Error()
	}
	if req.Error != "" {
		return req.Error
	}
	return fmt.Sprintf("user authorized. token: %s", req.Token)
}

func (c *ClientApp) Register(parent context.Context, login, password string) string {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	req, err := c.Client.RegisterUser(ctx, &pb.RegisterUserRequest{
		User: &pb.User{
			Login:    login,
			Password: password,
		},
	})
	if err != nil {
		return err.Error()
	}
	if req.Error != "" {
		return req.Error
	}
	return "user registered, please login to work further"
}
