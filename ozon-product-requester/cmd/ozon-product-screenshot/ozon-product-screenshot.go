package main

import (
	"github.com/alexflint/go-arg"
	"log"
	"os"
	"ozon-product-requester/internal/infrastructure/interactors"
)

var args struct {
	ID   string `arg:"positional,required"`
	Path string `arg:"positional,required"`
}

func main() {
	log.Println("Launching app")
	arg.MustParse(&args)
	makeScreenShotProductInteractor := interactors.NewMakeProductScreenshot()
	log.Println("Making product screenshot")
	imageBuffer, err := makeScreenShotProductInteractor.Call(args.ID)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(args.Path, *imageBuffer, 0600)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Screenshot saved to:", args.Path)
}
