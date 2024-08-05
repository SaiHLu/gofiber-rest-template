package constant

import (
	"fmt"

	"github.com/google/uuid"
)

func GetAccessToken(userId uuid.UUID) string {
	return fmt.Sprintf("accessToken%s", userId.String())
}

func GetRefreshToken(userId uuid.UUID) string {
	return fmt.Sprintf("refreshToken%s", userId.String())
}
