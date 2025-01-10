package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"videoverse/pkg/config"
	"videoverse/pkg/models"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPayload struct {
	models.JwtContextData
}

type tokenManager struct {
	signingKey string
}

type tokenDetails struct {
	TokenPayload
	jwt.RegisteredClaims
}

type ITokenManager interface {
	parse(accessToken string, signingKey string) (*tokenDetails, error)
	NewJWT(payload TokenPayload) (*string, error)
	NewRefreshToken() (string, error)
	VerifyToken(token *string) (*tokenDetails, error)
}

func NewTokenManager() ITokenManager {
	return &tokenManager{config.JWT_SECRET}
}

func (tkn *tokenManager) parse(accessToken string, signingKey string) (*tokenDetails, error) {
	var err error

	token, err := jwt.ParseWithClaims(accessToken, &tokenDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*tokenDetails)
	if !ok {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt.IsZero() {
		return claims, nil
	}

	// Check if expired
	if claims.ExpiresAt.Before(time.Now().UTC()) {
		return nil, errors.New("token expired")

	}
	return claims, nil
}

func (tkn *tokenManager) NewJWT(payload TokenPayload) (*string, error) {
	var key = []byte(tkn.signingKey)

	claims := tokenDetails{}
	claims.UserID = payload.UserID
	claims.Issuer = config.SERVICE_NAME
	claims.IssuedAt = jwt.NewNumericDate(time.Now().UTC())
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24).UTC())

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (tkn *tokenManager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().UTC().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}

func (tkn *tokenManager) VerifyToken(token *string) (*tokenDetails, error) {
	return tkn.parse(*token, tkn.signingKey)
}
