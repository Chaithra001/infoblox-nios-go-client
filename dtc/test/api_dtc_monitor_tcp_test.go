/*
Infoblox DTC API

Testing DtcMonitorTcpAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dtc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
)

func TestDtcMonitorTcpAPIService(t *testing.T) {

	apiClient := dtc.NewAPIClient()

	t.Run("Test DtcMonitorTcpAPIService Create", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DtcMonitorTcpAPI.Create(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DtcMonitorTcpAPIService Delete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		httpRes, err := apiClient.DtcMonitorTcpAPI.Delete(context.Background(), reference).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DtcMonitorTcpAPIService List", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DtcMonitorTcpAPI.List(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DtcMonitorTcpAPIService Read", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		resp, httpRes, err := apiClient.DtcMonitorTcpAPI.Read(context.Background(), reference).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DtcMonitorTcpAPIService Update", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		resp, httpRes, err := apiClient.DtcMonitorTcpAPI.Update(context.Background(), reference).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
