package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
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

	// Check if time is 0
	if claims.ExpiresAt != nil && claims.ExpiresAt.IsZero() {
		return claims, nil
	}

	// Check if expired
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now().UTC()) {
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

func GenerateSignedToken(userUID, videoID, filePath any, expiry time.Time) (*string, error) {
	payload := map[string]interface{}{
		"user_id":   userUID,
		"video_id":  videoID,
		"file_path": filePath,
		"expiry":    expiry.UnixMilli(),
	}
	byts, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	h := hmac.New(sha256.New, []byte(config.SHARE_SECRET))
	h.Write(byts)
	digest := h.Sum(nil)

	sig := base64.StdEncoding.EncodeToString(digest)

	payload["signature"] = sig
	encodedPayload, _ := json.Marshal(payload)
	str := base64.StdEncoding.EncodeToString(encodedPayload)

	return &str, nil
}

func VerifySignedToken(token string) (map[string]interface{}, error) {
	// Decode the base64-encoded token
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, errors.New("invalid token encoding")
	}

	// Unmarshal the JSON payload
	var payload map[string]interface{}
	err = json.Unmarshal(decodedToken, &payload)
	if err != nil {
		return nil, errors.New("invalid token payload")
	}

	// Extract the signature from the payload
	sig, ok := payload["signature"].(string)
	if !ok {
		return nil, errors.New("signature not found in token")
	}
	delete(payload, "signature")

	// Recompute the signature
	byts, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	h := hmac.New(sha256.New, []byte(config.SHARE_SECRET))
	h.Write(byts)
	expectedDigest := h.Sum(nil)
	expectedSig := base64.StdEncoding.EncodeToString(expectedDigest)

	// Compare  signatures
	if sig != expectedSig {
		return nil, errors.New("invalid token signature")
	}

	expiry, ok := payload["expiry"].(float64)
	if !ok {
		return nil, errors.New("expiry not found in token")
	}
	if time.Now().UnixMilli() > int64(expiry) {
		return nil, errors.New("token expired")
	}

	return payload, nil
}
