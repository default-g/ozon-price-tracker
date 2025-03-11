package main

import (
	"log"
	"ozon-product-requester/internal/infrastructure/interactors"
	ozon_client "ozon-product-requester/internal/infrastructure/ozon_client"

	"github.com/alexflint/go-arg"
)

var args struct {
	ID   string `arg:"positional,required"`
}

func main() {
	arg.MustParse(&args)
	ozonClient := ozon_client.NewOzonClient()
	interactor := interactors.NewRequestProductInteractor(ozonClient)
	product, err := interactor.Call(args.ID)
	if err != nil {
		log.Fatal(err)
	}

	
	log.Println(product)
}

