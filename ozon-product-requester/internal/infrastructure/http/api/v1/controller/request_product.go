package controller

import (
    "fmt"
    "github.com/gofiber/fiber/v3"
)

func requestProduct(ctx fiber.Ctx) error {
    msg := fmt.Sprintf("Hello, %s!", ctx.Path())
    return ctx.Send([]byte(msg))
}

func ProductController(app fiber.Router) {
    const ROUTE = "/api/v1/"

    app.Get(ROUTE, requestProduct)
}
