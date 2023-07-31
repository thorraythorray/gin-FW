package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type NewJwtClaim struct {
	User string
	jwt.RegisteredClaims
}

type JWT struct {
	SigningKey interface{}
	ExpireHour int
}

func (j *JWT) Obtaining(u string) (string, error) {
	claims := NewJwtClaim{
		u,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.ExpireHour) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(j.SigningKey)
	return ss, err
}

func (j *JWT) Authenticating(s string) (int, error) {
	token, err := jwt.ParseWithClaims(
		s,
		&NewJwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return j.SigningKey, nil
		},
	)

	if token.Valid {
		if claims, ok := token.Claims.(*NewJwtClaim); ok && token.Valid {
			// todo: query in db to validate user
			fmt.Printf("jwt check user:%s\n", claims.User)
		}
	}
	if errors.Is(err, jwt.ErrTokenMalformed) {
		return http.StatusBadRequest, errors.New("token解析失败")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return http.StatusForbidden, errors.New("无效的token")
	} else {
		return http.StatusBadRequest, err
	}
}
