package main

import (
	"log"
	"os"
	"ozon-product-requester/internal/infrastructure/interactors"
	ozon_client "ozon-product-requester/internal/ozon_client"
	"github.com/alexflint/go-arg"
)

var args struct {
	ID   string `arg:"positional,required"`
	Path string `arg:"positional,required"`
}

func main() {
	log.Println("Launching app")
	arg.MustParse(&args)
	ozonClient := ozon_client.NewOzonClient()
	makeScreenShotProductInteractor := interactors.NewMakeProductScreenshot(ozonClient)
	log.Println("Making product screenshot")
	imageBuffer, err := makeScreenShotProductInteractor.Call(args.ID)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(args.Path, imageBuffer, 0600)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Screenshot saved to:", args.Path)
}
