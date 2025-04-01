package controller

import (
    "github.com/gofiber/fiber/v3"
    "ozon-product-requester/internal/application/interactors"
    "ozon-product-requester/internal/infrastructure/http/api/v1/dto"
    "strings"
)

type ProductController struct {
    requestProductInteractor *interactors.RequestProductInteractor
}

func RegisterController(app fiber.Router, getProductInteractor *interactors.RequestProductInteractor) *ProductController {
    controller := &ProductController{getProductInteractor}

    app.Get("/api/v1", controller.getProduct)

    return controller
}

func (pc *ProductController) getProduct(ctx fiber.Ctx) error {

    id := ctx.Query("id", "")

    if strings.TrimSpace(id) == "" {
        return fiber.ErrBadRequest
    }

    product, err := pc.requestProductInteractor.Call(id)

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

    return ctx.JSON(productDto)
}
