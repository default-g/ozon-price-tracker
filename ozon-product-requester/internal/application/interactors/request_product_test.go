package interactors_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ozon-product-requester/internal/application/interactors"
	"ozon-product-requester/internal/domain/models"
)

type FakeOzonClientRequester struct {
	getProductFunc func(id string) (*models.Product, error)
}

func (f FakeOzonClientRequester) GetProduct(id string) (*models.Product, error) {
	return f.getProductFunc(id)
}

func TestRequestProductInteractor_Call(t *testing.T) {
	tests := []struct {
		name            string
		inputID         string
		mockProduct     *models.Product
		mockErr         error
		expectedErr     error
		expectedProduct *models.Product
	}{
		{
			name:            "successful product retrieval",
			inputID:         "123",
			mockProduct:     &models.Product{ID: "123"},
			mockErr:         nil,
			expectedProduct: &models.Product{ID: "123"},
			expectedErr:     nil,
		},
		{
			name:            "client returns error",
			inputID:         "456",
			mockProduct:     nil,
			mockErr:         errors.New("client error"),
			expectedProduct: nil,
			expectedErr:     errors.New("client error"),
		},
		{
			name:            "nil product without error",
			inputID:         "789",
			mockProduct:     nil,
			mockErr:         nil,
			expectedProduct: nil,
			expectedErr:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var capturedID string
			mockClient := &FakeOzonClientRequester{
				getProductFunc: func(id string) (*models.Product, error) {
					capturedID = id
					return tt.mockProduct, tt.mockErr
				},
			}

			interactor := interactors.NewRequestProductInteractor(mockClient)
			product, err := interactor.Call(tt.inputID)

			assert.Equal(t, tt.inputID, capturedID, "Expected client.GetProduct to receive ID %s, got %s", tt.inputID, capturedID)

			assert.Equal(t, tt.expectedProduct, product, "Product mismatch")
			assert.Equal(t, tt.expectedErr, err, "Error mismatch")
		})
	}
}
