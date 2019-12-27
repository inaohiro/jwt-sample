package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// todo: store environment variables
const PRIVKEY = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEICZ656/YaTOEVhSjWEblNEaflw5WDg5O9ucKrygwVv1IoAoGCCqGSM49
AwEHoUQDQgAEgWz7eN6MXctMJ8NLpHPwFki+kJ2WpUvqcr35UTfveJnbpb4neK7j
0l/BO/VOjfjRdgQsvgo96F9rbgOGrmWgZw==
-----END EC PRIVATE KEY-----`

const PUBKEY = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEgWz7eN6MXctMJ8NLpHPwFki+kJ2W
pUvqcr35UTfveJnbpb4neK7j0l/BO/VOjfjRdgQsvgo96F9rbgOGrmWgZw==
-----END PUBLIC KEY-----`

type Token struct {
	token string
}

func NewToken() (*Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	// todo: get private key from env
	ecdsaKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(PRIVKEY))
	if err != nil {
		return nil, err
	}
	tokenString, err := token.SignedString(ecdsaKey)
	if err != nil {
		return nil, err
	}

	return &Token{token: tokenString}, nil
}

func (token *Token) Token() string {
	if token != nil {
		return token.token
	}
	return ""
}

func (token *Token) parse() (*jwt.Token, error) {
	if token == nil {
		return nil, errors.New("nil token is not valid")
	}
	t, err := jwt.Parse(token.Token(), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		// todo: get public key from env
		ecdsaKey, err := jwt.ParseECPublicKeyFromPEM([]byte(PUBKEY))
		if err != nil {
			return nil, fmt.Errorf("parse ec pubkey: %v", err)
		}
		return ecdsaKey, nil
	})
	if err != nil {
		// fmt.Errorf
		return nil, err
	}

	return t, nil
}

func Validate(token *Token) error {
	t, err := token.parse()
	if err != nil {
		return err
	}

	if !t.Valid {
		return errors.New("token is not valid")
	}
	return nil
}
