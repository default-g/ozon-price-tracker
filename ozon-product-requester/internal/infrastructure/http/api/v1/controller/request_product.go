package controller

import (
    "encoding/json"
    "github.com/gofiber/fiber/v3"
    "ozon-product-requester/internal/infrastructure/http/api/v1/dto"
    "ozon-product-requester/internal/infrastructure/ozon"
    "strings"
)

type ProductController struct {
    ozonClient ozon.Client
}

func RegisterController(app fiber.Router, ozonClient ozon.Client) *ProductController {
    controller := &ProductController{ozonClient}

    app.Get("/api/v1", controller.getProduct)

    return controller
}

func (pc *ProductController) getProduct(ctx fiber.Ctx) error {

    id := ctx.Query("id", "")

    if strings.TrimSpace(id) == "" {
        return fiber.ErrBadRequest
    }

    product, err := pc.ozonClient.GetProduct(id)

    if err != nil {
        return err
    }

    productDto := dto.ProductDto{
        ID:   string(product.ID),
        Name: string(product.Name),
        OzonPrice: dto.PriceDto{
            Currency: string(product.OzonCardPrice.Currency),
            Value:    float32(product.OzonCardPrice.Value),
        },
        Price: dto.PriceDto{
            Currency: string(product.Price.Currency),
            Value:    float32(product.Price.Value),
        },
    }

    marshal, err := json.Marshal(productDto)
    if err != nil {
        return err
    }

    return ctx.Send(marshal)
}
