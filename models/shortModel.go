package models

import (
	"time"

	"github.com/speps/go-hashids"
)

type Url struct {
	ID   int    `json:"id"   db:"id"`
	Hash string `json:"hash" db:"hash"`
	Url  string `json:"url"  db:"url"`
}

func GetUrlRedirect(uri string) (string, error) {
	rows, err := DB.Query("SELECT id, hash, url FROM urls WHERE hash = $1", uri)

	if err != nil {
		return "", err
	}

	urlRedirect := ""

	for rows.Next() {
		data := Url{}
		rows.Scan(&data.ID, &data.Hash, &data.Url)

		urlRedirect = data.Url
	}

	return urlRedirect, nil
}

func VerifyUrl(url string) (string, error) {
	rows, err := DB.Query("SELECT url, hash FROM urls WHERE url = $1", url)

	if err != nil {
		return "", err
	}

	hash := ""

	for rows.Next() {
		data := Url{}
		rows.Scan(&data.Url, &data.Hash)

		hash = data.Hash
	}

	return hash, nil
}

func InsertUrl(Url string) (string, error) {
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	hash, _ := h.Encode([]int{int(now.Unix())})

	_, err := DB.Query("INSERT INTO urls (url, hash) VALUES ($1, $2)", Url, hash)

	if err != nil {
		return "", err
	}

	return hash, nil
}
