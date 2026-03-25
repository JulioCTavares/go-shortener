package shortener

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/julioctavares/go-shortener/pkg/config"
)

func CreateShortUrl(originalUrl string) (string, error) {

	var cachedUrl, err = cachedUrl(originalUrl)

	if err == nil {
		return cachedUrl, nil
	}

	var db = config.DB

	var existingCode, someErr = existingUrl(originalUrl, db)

	if someErr != nil {
		return "", someErr
	} else if existingCode != "" {
		return existingCode, nil
	}

	var generatedCode string
	for {
		generatedCode, err = GenerateShortURL()
		if err != nil {
			return "", err
		}

		var exists bool

		err = db.QueryRow(context.Background(),
			"SELECT EXISTS(SELECT 1 FROM urls WHERE code = $1)", generatedCode,
		).Scan(&exists)
		if err != nil {
			return "", err
		}
		if !exists {
			break
		}
	}

	expiresAfter3Days := expireAfter(3)

	row := db.QueryRow(context.Background(),
		"INSERT INTO urls (code, original_url, expires_at) VALUES ($1, $2, $3) RETURNING id",
		generatedCode, originalUrl, expiresAfter3Days)

	if row == nil {
		return "", errors.New("Failed to create short URL")
	}

	saveUrlInCache(originalUrl, generatedCode)

	return generatedCode, nil
}

func expireAfter(days int) time.Time {
	return time.Now().AddDate(0, 0, days)
}

func cachedUrl(originalUrl string) (string, error) {
	var redisClient = config.RedisClient

	cachedUrl := redisClient.Get("short_url:" + originalUrl)

	if cachedUrl != nil {
		return cachedUrl.Result()
	}

	return "", nil
}

func existingUrl(originalUrl string, conn *pgx.Conn) (string, error) {

	var existingCode string

	err := conn.QueryRow(context.Background(),
		"SELECT code FROM urls WHERE original_url = $1", originalUrl).Scan(&existingCode)

	if err == nil {
		return existingCode, nil
	} else if errors.Is(err, pgx.ErrNoRows) {
		return "", nil
	}

	return "", nil
}

func saveUrlInCache(originalUrl string, shortUrl string) error {
	var redisClient = config.ConnectRedis()

	err := redisClient.Set("short_url:"+originalUrl, shortUrl, 3*24*time.Hour).Err()

	if err != nil {
		return err
	}

	return nil
}
