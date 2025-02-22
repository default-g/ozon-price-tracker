package interactors

import (
	entity "ozon-product-requester/internal/domain/product/entity"
)

type RequestProduct interface {
	Call(id string) (*entity.Product, error)
}
