package interactors

import (
	product "ozon-product-requester/internal/domain/models"
	"ozon-product-requester/internal/infrastructure/ozon"
)

type RequestProductInteractor struct {
	ozonProductRequester ozon.ProductRequester
}

func NewRequestProductInteractor(ozonProductRequester ozon.ProductRequester) RequestProductInteractor {
	return RequestProductInteractor{
		ozonProductRequester: ozonProductRequester,
	}
}

func (i RequestProductInteractor) Call(id string) (*product.Product, error) {
	return i.ozonProductRequester.GetProduct(id)
}
