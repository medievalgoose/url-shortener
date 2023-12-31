package db

import (
	"fmt"
	"log"
)

type CustomURL struct {
	PlainUrl string
	ShortUrl string
}

func AddNewUrl(newUrl CustomURL) error {
	db := OpenConnection()
	defer db.Close()

	createQuery := "INSERT INTO urls (plain_url, short_url) VALUES ($1, $2);"
	_, err := db.Exec(createQuery, newUrl.PlainUrl, newUrl.ShortUrl)
	if err != nil {
		return err
	}

	return nil
}

func IsValidUrl(url string) bool {
	db := OpenConnection()
	defer db.Close()

	validUrlId := 0

	selectQuery := "SELECT id FROM urls WHERE short_url = $1;"
	err := db.QueryRow(selectQuery, url).Scan(&validUrlId)

	return err == nil
}

func GetPlainUrl(shortUrl string) (string, error) {
	db := OpenConnection()
	defer db.Close()

	urlValid := IsValidUrl(shortUrl)

	if urlValid {
		plainUrl := ""

		plainUrlQuery := "SELECT plain_url FROM urls WHERE short_url = $1;"
		err := db.QueryRow(plainUrlQuery, shortUrl).Scan(&plainUrl)
		if err != nil {
			log.Fatal(err)
		}

		return plainUrl, nil
	} else {
		return "", fmt.Errorf("short URL doesn't exist")
	}
}
