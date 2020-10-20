package authentication

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ServiceI ...
type ServiceI interface {
	CreateAuthToken(int) (string, error)
	TokenValid(string) error
	ExtractTokenMetadata(string) (*AccessDetails, error)
}

// Service ...
type Service struct{}

// NewService ...
func NewService() *Service {
	return &Service{}
}

// CreateAuthToken ...
func (a *Service) CreateAuthToken(userID int) (string, error) {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

// TokenValid ...
func (a *Service) TokenValid(tokenString string) error {
	token, err := verifyToken(tokenString)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// AccessDetails ...
type AccessDetails struct {
	UserID int
}

// ExtractTokenMetadata ...
func (a *Service) ExtractTokenMetadata(tokenString string) (*AccessDetails, error) {
	token, err := verifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			UserID: int(userID),
		}, nil
	}
	return nil, err
}

// VerifyToken ...
func verifyToken(tokenString string) (*jwt.Token, error) {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
