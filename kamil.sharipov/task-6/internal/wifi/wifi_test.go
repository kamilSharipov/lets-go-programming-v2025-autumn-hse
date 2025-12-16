package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	w "github.com/kamilSharipov/task-6/internal/wifi"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=WiFiHandle --dir=. --output=. --outpkg wifi_test --structname=MockWiFiHandle

var (
	errInterfacesError = errors.New("interfaces error")
	errNamesError      = errors.New("names error")
)

func TestGetAddressesSuccess(t *testing.T) {
	t.Parallel()

	mock := NewMockWiFiHandle(t)

	hwAddr, err := net.ParseMAC("38:d5:7a:eb:43:8f")
	require.NoError(t, err)

	interfaces := []*wifi.Interface{
		{HardwareAddr: hwAddr},
	}

	want := []net.HardwareAddr{hwAddr}

	mock.On("Interfaces").Return(interfaces, nil)

	service := w.New(mock)
	have, err := service.GetAddresses()

	require.NoError(t, err)
	assert.Equal(t, want, have)

	mock.AssertExpectations(t)
}

func TestGetAddressesError(t *testing.T) {
	t.Parallel()

	mock := NewMockWiFiHandle(t)

	mock.On("Interfaces").Return(nil, errInterfacesError)

	service := w.New(mock)
	have, err := service.GetAddresses()

	assert.Nil(t, have)
	require.ErrorContains(t, err, "getting interfaces: "+errInterfacesError.Error())

	mock.AssertExpectations(t)
}

func TestGetNamesSuccess(t *testing.T) {
	t.Parallel()

	mock := NewMockWiFiHandle(t)

	wifiInterfaceName := "wlp2s0"
	interfaces := []*wifi.Interface{
		{Name: wifiInterfaceName},
	}

	mock.On("Interfaces").Return(interfaces, nil)

	want := []string{wifiInterfaceName}
	service := w.New(mock)
	have, err := service.GetNames()

	require.NoError(t, err)
	assert.Equal(t, want, have)

	mock.AssertExpectations(t)
}

func TestGetNamesError(t *testing.T) {
	t.Parallel()

	mock := NewMockWiFiHandle(t)

	mock.On("Interfaces").Return(nil, errNamesError)

	service := w.New(mock)
	have, err := service.GetNames()

	assert.Nil(t, have)
	require.ErrorContains(t, err, "getting interfaces: "+errNamesError.Error())

	mock.AssertExpectations(t)
}
