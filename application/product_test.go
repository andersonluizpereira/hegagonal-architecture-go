package application_test

import (
	"github.com/acpereira/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)


func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		Name: "Hello",
		Status: application.ENABLED,
		Price: 10,
	}

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must ben greater than zero to enable the product", err.Error())
}

func TestProduct_Disabled(t *testing.T) {
	product := application.Product{
		Name: "Hello",
		Status: application.ENABLED,
		Price: 0,
	}

	err := product.Disabled()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disabled()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID: uuid.NewV4().String(),
		Name: "Hello",
		Status: application.DISABLED,
		Price: 10,
	}
	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "The status must enable or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "The price must be greater or equal zero", err.Error())

}