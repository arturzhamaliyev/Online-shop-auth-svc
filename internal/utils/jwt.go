package utils

import (
	"errors"
	"time"

	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/models"
	"github.com/golang-jwt/jwt"
)

type JwtWrapper struct {
	SecretKey      string
	Issuer         string
	ExpirationTime int64
}

type jwtClaims struct {
	jwt.StandardClaims
	Id    int64
	Email string
}

func (w *JwtWrapper) GenerateToken(user models.User) (string, error) {
	claims := &jwtClaims{
		Id:    user.Id,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationTime)).Unix(),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (*jwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)
	if err != nil {
		return &jwtClaims{}, err
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil
}
