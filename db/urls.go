package db

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/3nt3/urlshortener/structs"
)

func CreateURL(shortURL structs.ShortURL) (string, error) {
	id, err := generateID()
	if err != nil {
		log.Panicf("[ - ] error generating id: %s\n", err.Error())
		return "", err
	}

	statement := "insert into urls (id, original_url, created_at) VALUES ($1, $2, $3);"
	_, err = database.Exec(statement, id, shortURL.OriginalURL, time.Now())

	return id, err
}

func generateID() (string, error) {
	rand.Seed(time.Now().Unix())

	const chars string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJLKMNOPQRSTUVWXYZ"
	unique := false

	var output strings.Builder
	const n = 4

	for !unique {
		output.Reset()
		for i := 0; i < n; i++ {
			random := rand.Intn(len(chars))
			randomChar := chars[random]
			output.WriteString(string(randomChar))
		}

		var count int
		err := database.QueryRow("SELECT COUNT(*) FROM urls WHERE id = $1", output.String()).Scan(&count)
		if err != nil {
			return "", err
		}

		if count == 0 {
			unique = true
		}
	}
	return output.String(), nil
}

func GetURLByID(id string) (structs.ShortURL, error) {
	query := "SELECT * FROM urls WHERE id = $1;"

	var url structs.ShortURL
	err := database.QueryRow(query, id).Scan(&url.ID, &url.OriginalURL, &url.CreatedAt)
	return url, err
}
