package ozon

import product "ozon-product-requester/internal/domain/models"

type Client interface {
	ProductRequester
	ProductScreenshotMaker
}

type ProductRequester interface {
	GetProduct(id string) (*product.Product, error)
}

type ProductScreenshotMaker interface {
	MakeScreenshot(id string) ([]byte, error)
}
