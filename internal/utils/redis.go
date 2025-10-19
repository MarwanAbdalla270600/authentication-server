package utils

import (
	"authentication-server/internal/entity"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func CreateRefreshTokenRedis(ctx context.Context, rdb *redis.Client, user *entity.UserDAO, duration time.Duration) (string, error) {
	token := uuid.NewString()
	key := fmt.Sprintf("refresh:token:%s", token)
	value := fmt.Sprintf(`{"userId":"%s","email":"%s","role":"%s"}`, user.Id, user.Email, user.Role)
	err := rdb.Set(ctx, key, value, duration).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}
