package interactors

import ozon_client "ozon-product-requester/internal/infrastructure/ozon_client"

type MakeProductScreenshot struct {
	ozonClient *ozon_client.Client
}

func NewMakeProductScreenshot(ozonClient *ozon_client.Client) MakeProductScreenshot {
	return MakeProductScreenshot{
		ozonClient: ozonClient,
	}
}

func (p MakeProductScreenshot) Call(id string) ([]byte, error) {
	return p.ozonClient.MakeScreeshot(id)
}
