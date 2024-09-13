package service

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/oscargh945/go-Chat/domain/entities"
	"os"
	"time"
)

func GenerateTokens(user entities.User) (entities.LoginResponse, error) {
	accessTokenID := uuid.New()
	refreshTokenID := uuid.New()

	accessDuration := time.Minute * 24
	refreshDuration := time.Minute * 25

	accessTokenClaims := jwt.MapClaims{
		"sub":       user.ID,
		"user_name": user.UserName,
		"email":     user.Email,
		"exp":       time.Now().Add(accessDuration).Unix(),
		"jti":       accessTokenID.String(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	refreshTokenClaims := jwt.MapClaims{
		"sub":       user.ID,
		"user_name": user.UserName,
		"email":     user.Email,
		"exp":       time.Now().Add(refreshDuration).Unix(),
		"jti":       refreshTokenID.String(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	tokenAccessString, err := accessToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return entities.LoginResponse{}, err
	}

	tokenRefreshString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return entities.LoginResponse{}, err
	}

	return entities.LoginResponse{
		AccessToken:  tokenAccessString,
		RefreshToken: tokenRefreshString,
	}, nil
}

func RefreshToken(refreshToken string) (*entities.LoginResponse, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, err
	}

	var user entities.User
	tokens, _ := GenerateTokens(user)
	response := &entities.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
	return response, nil
}

func ValidateToken(tokenString string) (string, error) {
	secretKey := []byte(os.Getenv("SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return "", nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimsJSON, err := json.Marshal(map[string]interface{}{
			"sub":       claims["sub"].(string),
			"user_name": claims["user_name"].(string),
			"email":     claims["email"].(string),
			"exp":       claims["exp"].(float64),
		})
		if err != nil {
			return "", nil
		}
		return string(claimsJSON), nil
	}

	return "", nil
}
