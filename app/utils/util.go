package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/models"
)

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))

	hashedPassword := fmt.Sprintf("%x", hash)
	return hashedPassword
}
func VerifyPassword(password, hashedPassword string) bool {
	hashedInput := HashPassword(password)

	if hashedInput == hashedPassword {
		return true
	}
	return false
}

var privateKey *ecdsa.PrivateKey

func init() {
	var err error
	privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
}

func CreateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id":   user.ID,
		"mail": user.Email,
		"exp":  time.Now().Add(time.Minute * 30).Unix(),
	})

	// ECDSA ile imzala
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Token doğrula
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return &privateKey.PublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	// Token geçerlilik süresini kontrol et
	if !token.Valid {
		return nil, errors.New("Token is expired")
	}
	return token, nil
}

func InvalidateToken(tokenString string) (string, error) {
	// Token doğrula
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return privateKey.PublicKey, nil
	})
	if err != nil {
		return "", err
	}

	// Token geçersizleştir
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("Invalid token claims")
	}

	claims["exp"] = time.Now().Unix() // Geçerlilik süresini şu anın tarihine ayarla

	// Geçersizleştirilmiş tokenı döndür
	newToken := jwt.NewWithClaims(token.Method, claims)
	invalidatedTokenString, err := newToken.SigningString()
	if err != nil {
		return "", err
	}

	return invalidatedTokenString, nil
}
