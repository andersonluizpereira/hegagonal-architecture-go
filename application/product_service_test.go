package application_test

import (
	"errors"
	"github.com/acpereira/go-hexagonal/application"
	mock_application "github.com/acpereira/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}
	result, err := service.Get("123")
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_Get_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, errors.New("ID not found")).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}
	result, err := service.Get("123")
	require.Equal(t, "ID not found",err.Error())
	require.Equal(t, nil, result)

}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}
	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, errors.New("Server Error")).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}
	_, err := service.Create("Product 1", 10)
	require.Equal(t, "Server Error",err.Error())

}

func TestProductService_Create_Product_Invalid_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, errors.New("the price must be greater or equal zero\"")).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}
	_, err := service.Create("Product 1", -1)
	require.Equal(t, "the price must be greater or equal zero",err.Error())

}