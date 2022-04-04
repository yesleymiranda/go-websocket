package jwtoken

import (
	"errors"

	"github.com/yesleymiranda/go-websocket/server/account"

	"github.com/golang-jwt/jwt"
)

type service struct {
	secret []byte
}

type Service interface {
	Create(acc *account.Account) (string, error)
	ReadAndValidate(token string) (*account.Account, error)
}

func NewService() Service {
	secret := []byte("yeye::super-secret-=D")
	return &service{
		secret,
	}
}

func (s service) Create(acc *account.Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Data{
		Username: acc.Username,
	})

	return token.SignedString(s.secret)
}

func (s service) ReadAndValidate(tokenString string) (*account.Account, error) {
	if tokenString == "" {
		return nil, errors.New("token is required")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Data{}, func(token *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Data)
	if !token.Valid || !ok {
		return nil, errors.New("token is invalid")
	}

	acc := &account.Account{
		Username: claims.Username,
		Token:    tokenString,
	}
	return acc, nil
}
