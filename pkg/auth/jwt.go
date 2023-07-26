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

type JwtVerify struct {
	Status int
	Msg    string
}

type NewJwtClaim struct {
	UserID string
	jwt.RegisteredClaims
}

type JWT struct {
	SigningKey interface{}
}

func (j *JWT) CreateJwtToken(userid string, expireAt int) (string, error) {
	claims := NewJwtClaim{
		userid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireAt) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(j.SigningKey)
	return ss, err
}

func (j *JWT) ParseJwtToken(tokenString string) (*NewJwtClaim, int, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&NewJwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return j.SigningKey, nil
		},
	)

	if token.Valid {
		if claims, ok := token.Claims.(*NewJwtClaim); ok && token.Valid {
			return claims, JwtTokenInvalid, nil
		}
	}
	if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, JwtParseError, errors.New("token解析失败")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return nil, JwtTokenInvalid, errors.New("无效的token")
	} else {
		return nil, JwtClaimsInvalid, err
	}
}
