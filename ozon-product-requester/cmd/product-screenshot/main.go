package main

import (
	"log"
	"ozon-product-requester/internal/application/interactors"
	"ozon-product-requester/internal/infrastructure/ozon"
)

func main() {
	ozonClientBuilder := ozon.NewOzonClientBuilder()
	interactorFactory := func() interactors.MakeProductScreenshot {
		return interactors.NewMakeProductScreenshot(ozonClientBuilder.Build())
	}
	args := GetProductScreenshotArgs{}

	command := NewGetProductScreenshotCommand(interactorFactory, &args)

	err := command.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
