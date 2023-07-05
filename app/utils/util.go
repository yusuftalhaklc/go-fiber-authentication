package utils

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/config"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/models"
)

// HashPassword generates a SHA256 hash of the provided password.
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	hashedPassword := fmt.Sprintf("%x", hash)
	return hashedPassword
}

// VerifyPassword compares the provided password with the hashed password and returns true if they match.
func VerifyPassword(password, hashedPassword string) bool {
	hashedInput := HashPassword(password)
	if hashedInput == hashedPassword {
		return true
	}
	return false
}

var secretKey = []byte(config.Config("SECRET_KEY"))

// CreateToken generates a JWT token for the provided user.
func CreateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Failed to create token:", err)
		return "", err
	}
	return tokenString, nil
}

// VerifyToken verifies the provided JWT token and returns the claims if the token is valid.
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, errors.New("Invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid token claims")
	}
	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	currentTime := time.Now()
	fmt.Println("exp : ", expirationTime, "\n", "current : ", currentTime)
	if currentTime.After(expirationTime) {
		return nil, errors.New("Invalid token")
	}
	if !token.Valid {
		return nil, errors.New("Invalid token")
	}
	return claims, nil
}

// InvalidateToken invalidates the provided JWT token by setting its expiration time to the current time.
func InvalidateToken(tokenString string) (string, error) {
	// Verify the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	// Invalidate the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("Invalid token claims")
	}
	claims["exp"] = time.Now().Unix() // Set the expiration time to the current time

	// Return the invalidated token
	newToken := jwt.NewWithClaims(token.Method, claims)
	invalidatedTokenString, err := newToken.SigningString()
	if err != nil {
		return "", err
	}

	return invalidatedTokenString, nil
}
