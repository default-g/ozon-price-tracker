package dto

type ProductDto struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Price     PriceDto `json:"price"`
	OzonPrice PriceDto `json:"ozon_price"`
}
