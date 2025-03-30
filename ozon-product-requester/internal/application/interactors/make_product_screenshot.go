package interactors

import (
	"ozon-product-requester/internal/infrastructure/ozon"
)

type MakeProductScreenshot struct {
	ozonClient ozon.Client
}

func NewMakeProductScreenshot(ozonClient ozon.Client) MakeProductScreenshot {
	return MakeProductScreenshot{
		ozonClient: ozonClient,
	}
}

func (p MakeProductScreenshot) Call(id string) ([]byte, error) {
	return p.ozonClient.MakeScreenshot(id)
}
