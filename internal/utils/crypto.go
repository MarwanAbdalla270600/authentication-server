package utils

import (
	"authentication-server/internal/entity"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)



func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func loadPrivateKey(path string) (*ecdsa.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("invalid PEM block")
	}
	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}


func CreateAccessToken(user *entity.UserDAO) (string, error) {
	privKey, err := loadPrivateKey("keys/private.pem")
	if err != nil {
		return "", err
	}

	// Define claims
	claims := jwt.MapClaims{
		"iss":   "auth-server",
		"sub":   user.Id,
		"email": user.Email,
		"role":  user.Role,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour).Unix(),
	}

	// Create ES256 token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// Sign with the private key
	return token.SignedString(privKey)
}