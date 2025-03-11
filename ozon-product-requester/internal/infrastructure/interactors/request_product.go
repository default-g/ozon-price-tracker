package interactors

import (
	product "ozon-product-requester/internal/domain/product/entity"
	ozon_client "ozon-product-requester/internal/infrastructure/ozon_client"
)

type RequestProductInteractor struct {
	ozonClient *ozon_client.Client
}

func NewRequestProductInteractor(ozonClient *ozon_client.Client) RequestProductInteractor {
	return RequestProductInteractor{
		ozonClient: ozonClient,
	}
}

func (i RequestProductInteractor) Call(id string) (*product.Product, error) {
	return i.ozonClient.GetProduct(id)
}
