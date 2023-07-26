package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	JwtParseError    = 500
	JwtClaimsInvalid = 400
	JwtTokenInvalid  = 403
)

type NewJwtClaim struct {
	UserID string
	jwt.RegisteredClaims
}

type JWT struct {
	SigningKey interface{}
	ExpireHour int
	CheckUser  string
	JwtString  string
}

func (j *JWT) Creating() (string, error) {
	claims := NewJwtClaim{
		j.CheckUser,
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

func (j *JWT) Parsing() (int, error) {
	token, err := jwt.ParseWithClaims(
		j.JwtString,
		&NewJwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return j.SigningKey, nil
		},
	)

	if token.Valid {
		if claims, ok := token.Claims.(*NewJwtClaim); ok && token.Valid {
			if j.CheckUser == claims.UserID {
				return 200, err
			} else {
				return JwtTokenInvalid, errors.New("无效的token")
			}
		}
	}
	if errors.Is(err, jwt.ErrTokenMalformed) {
		return JwtParseError, errors.New("token解析失败")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return JwtTokenInvalid, errors.New("无效的token")
	} else {
		return JwtClaimsInvalid, err
	}
}
