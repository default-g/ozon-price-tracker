package dto

type ProductDto struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	OzonPrice float32 `json:"ozon_price"`
}
