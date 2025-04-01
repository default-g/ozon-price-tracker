package main

import (
	"log"
	"ozon-product-requester/internal/application/interactors"
	"ozon-product-requester/internal/infrastructure/http"
	"ozon-product-requester/internal/infrastructure/http/api/v1/controller"
	"ozon-product-requester/internal/infrastructure/ozon"
)

func main() {

	app := http.GetApp()

	ozonClient := ozon.NewOzonClientBuilder().Build()

	requestProductInteractor := interactors.NewRequestProductInteractor(ozonClient)
	controller.RegisterController(app, &requestProductInteractor)

	err := http.GetApp().Listen(":3000")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Finishing work")
}
