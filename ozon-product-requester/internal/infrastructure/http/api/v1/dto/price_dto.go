package dto

type PriceDto struct {
	Currency string  `json:"currency"`
	Value    float32 `json:"value"`
}
