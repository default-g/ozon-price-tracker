package main

import (
	"log"
	"ozon-product-requester/internal/application/interactors"
	"ozon-product-requester/internal/infrastructure/ozon"
)

func main() {
	ozonClient := ozon.NewOzonClientBuilder().Build()
	interactor := interactors.NewRequestProductInteractor(ozonClient)

	args := GetProductInfoArgs{}
	command := NewGetProductInfoCommand(interactor, &args)
	err := command.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
