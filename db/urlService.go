package db

import (
	"fmt"
	"log"
	"medievalgoose/url-shortener/util"
)

type CustomURL struct {
	Id       int
	PlainUrl string `json:"plain_url"`
	ShortUrl string `json:"short_url"`
}

func AddNewUrl(newUrl CustomURL) (string, error) {
	db := OpenConnection()
	defer db.Close()

	createQuery := "INSERT INTO urls (plain_url) VALUES ($1) RETURNING id;"
	err := db.QueryRow(createQuery, newUrl.PlainUrl).Scan(&newUrl.Id)
	if err != nil {
		return "", err
	}

	newUrl.ShortUrl = util.EncodeToBase62(newUrl.Id)

	updateUrlQuery := "UPDATE urls SET short_url = $1 WHERE id = $2;"
	_, err = db.Exec(updateUrlQuery, newUrl.ShortUrl, newUrl.Id)
	if err != nil {
		return "", err
	}

	return newUrl.ShortUrl, nil
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
