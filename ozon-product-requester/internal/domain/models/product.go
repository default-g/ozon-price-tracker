package models

import (
	"slices"
)

type ProductError struct {
	message string
}

func (e ProductError) Error() string {
	return e.message
}

var (
	ErrInvalidOzonCardPrice = ProductError{"invalid ozon card price"}
	ErrInvalidPrice         = ProductError{"invalid price"}
	ErrInvalidName          = ProductError{"name cannot be empty"}
)

type ProductID string
type Currency string
type Name string
type Availability bool

type Price struct {
	Value    float64
	Currency Currency
}

type Product struct {
	ID            ProductID
	OzonCardPrice Price
	Price         Price
	Name          Name
	Available     Availability
}

const (
	RUB Currency = "RUB"
)

const (
	Available    Availability = true
	NotAvailable Availability = false
)

func GetCurrencies() []Currency {
	return []Currency{RUB}
}

func (currency Currency) IsValid() bool {
	return slices.Contains(GetCurrencies(), currency)
}

func (price Price) IsValid() bool {
	return price.Value > 0 && price.Currency.IsValid()
}

func NewProduct(ID ProductID, ozonCardPrice Price, price Price, name Name, available Availability) (*Product, error) {

	if !ozonCardPrice.IsValid() {
		return nil, ErrInvalidOzonCardPrice
	}

	if !price.IsValid() {
		return nil, ErrInvalidPrice
	}

	return &Product{
		ID:            ID,
		OzonCardPrice: ozonCardPrice,
		Price:         price,
		Name:          name,
		Available:     available,
	}, nil
}
