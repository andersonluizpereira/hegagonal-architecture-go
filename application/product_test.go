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

func TestProduct_GetID(t *testing.T) {
	 ID := uuid.NewV4().String()
	product := application.Product{
		ID:ID,
		Name: "Hello",
		Status: application.ENABLED,
		Price: 10,
	}

	err := product.Enable()
	require.Nil(t, err)
	newId := product.GetID()
	require.Equal(t, ID, newId)
}

func TestProduct_GetName(t *testing.T) {
	Name := "Hello"
	product := application.Product{
		ID: uuid.NewV4().String(),
		Name: Name,
		Status: application.ENABLED,
		Price: 10,
	}

	err := product.Enable()
	require.Nil(t, err)
	getName := product.GetName()
	require.Equal(t, Name, getName)
}

func TestProduct_GetStatusEnable(t *testing.T) {
	Status := application.ENABLED
	product := application.Product{
		ID: uuid.NewV4().String(),
		Name: "Hello",
		Status: Status,
		Price: 10,
	}

	err := product.Enable()
	require.Nil(t, err)
	getStatusEnable := product.GetStatus()
	hasValid, err := product.IsValid()
	require.Nil(t, err)
	require.Equal(t, Status, getStatusEnable)
	require.Equal(t, true, hasValid)
}

func TestProduct_GetStatusDisabled(t *testing.T) {
	Status := application.ENABLED
	product := application.Product{
		ID: uuid.NewV4().String(),
		Name: "Hello",
		Status: Status,
		Price: 0,
	}
	product.Disabled()
	getStatusEnable := product.GetStatus()
	require.Equal(t, application.DISABLED, getStatusEnable)

}

func TestProduct_GetPrice(t *testing.T) {
	Price := float64(100)
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "Hello",
		Status: application.ENABLED,
		Price: Price,
	}

	err := product.Enable()
	require.Nil(t, err)
	newPrice := product.GetPrice()
	require.Equal(t, Price, newPrice)
}

func TestProduct_InvalidStruct(t *testing.T) {
	product := application.Product{
		ID: "123",
		Name: "Hello",
		Status: "",
		Price: 10,
	}
	hasValidStruct, err := product.IsValid()
	require.Equal(t, "ID: 123 does not validate as uuidv4",err.Error())
	require.Equal(t, false, hasValidStruct)
}