package main

import (
	"log"
	"ozon-product-requester/internal/infrastructure/http"
	"ozon-product-requester/internal/infrastructure/http/api/v1/controller"
)

func main() {

	app := http.GetApp()

	controller.ProductController(app)

	err := http.GetApp().Listen(":3000")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Finishing work")
}
