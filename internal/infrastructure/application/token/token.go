package token

import (
	"github.com/SaiHLu/rest-template/common"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateAccessToken(userId uuid.UUID, ttl int64) (string, error) {
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": ttl,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(common.GetEnv().ACCESS_TOKEN_SECRET))

	return t, err
}

func GenerateRefreshToken(userId uuid.UUID, ttl int64) (string, error) {
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": ttl,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(common.GetEnv().REFRESH_TOKEN_SECRET))

	return t, err
}
