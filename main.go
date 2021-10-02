package main

import (
	"labs/shortener/controllers"
	"labs/shortener/models"
)

func main() {
	db := models.ConnectDataBase()
	defer db.Close()

	controllers.Routes()
}
