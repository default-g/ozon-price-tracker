package interactors

import ozon_client "ozon-product-requester/internal/ozon-client"

type MakeProductScreenshot struct{}

func NewMakeProductScreenshot() *MakeProductScreenshot {
	return &MakeProductScreenshot{}
}

func (p *MakeProductScreenshot) Call(id string) (*[]byte, error) {
	return ozon_client.MakeScreeshot(id)
}
