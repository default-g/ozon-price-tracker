package main

import (
	"fmt"
	"ozon-product-requester/internal/infrastructure/interactors"
	ozon_client "ozon-product-requester/internal/ozon_client"
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
		fmt.Println(err)
	}

	
	fmt.Println(product)
}

