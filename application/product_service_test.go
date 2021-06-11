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