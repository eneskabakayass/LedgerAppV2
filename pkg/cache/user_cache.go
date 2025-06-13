package cache

import (
	"LedgerV2/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const userCacheTTL = time.Minute * 5

func GetUserFromCache(userID string) (*models.User, error) {
	key := fmt.Sprintf("user:%s", userID)

	data, err := RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func SetUserToCache(user *models.User) error {
	key := fmt.Sprintf("user:%s", user.ID)

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return RedisClient.Set(context.Background(), key, data, userCacheTTL).Err()
}
