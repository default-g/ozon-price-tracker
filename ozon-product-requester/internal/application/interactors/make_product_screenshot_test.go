package interactors_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type FakeOzonProductScreenshotMaker struct {
	makeProductScreenshot func(id string) ([]byte, error)
}

func (f FakeOzonProductScreenshotMaker) MakeScreenshot(id string) ([]byte, error) {
	return f.makeProductScreenshot(id)
}

func TestMakeProductScreenshot_Call(t *testing.T) {
	tests := []struct {
		name      string
		productId string
		mockBytes []byte
		mockError error
	}{
		{
			name:      "Make product screenshot success",
			productId: "productId",
			mockBytes: []byte(`123`),
			mockError: nil,
		},
		{
			name:      "Make product screenshot error",
			productId: "productId",
			mockBytes: nil,
			mockError: errors.New("error making request"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeOzonClient := FakeOzonProductScreenshotMaker{
				makeProductScreenshot: func(id string) ([]byte, error) {
					return tt.mockBytes, tt.mockError
				},
			}

			productImage, err := fakeOzonClient.MakeScreenshot(tt.productId)

			assert.Equal(t, tt.mockBytes, productImage)
			assert.Equal(t, tt.mockError, err)
		})
	}
}
