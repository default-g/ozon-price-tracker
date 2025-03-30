package main

import (
    "log"
    "ozon-product-requester/internal/infrastructure/http"
    "ozon-product-requester/internal/infrastructure/http/api/v1/controller"
    "ozon-product-requester/internal/infrastructure/ozon"
)

func main() {

    app := http.GetApp()

    ozonClient := ozon.NewOzonClientBuilder().Build()

    controller.RegisterController(app, ozonClient)

    err := http.GetApp().Listen(":3000")

    if err != nil {
        log.Fatal(err)
    }

    log.Println("Finishing work")
}
