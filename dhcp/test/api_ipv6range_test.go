/*
Infoblox DHCP API

Testing Ipv6rangeAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dhcp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

func TestIpv6rangeAPIService(t *testing.T) {

	apiClient := dhcp.NewAPIClient()

	t.Run("Test Ipv6rangeAPIService Create", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.Ipv6rangeAPI.Create(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test Ipv6rangeAPIService Delete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		httpRes, err := apiClient.Ipv6rangeAPI.Delete(context.Background(), reference).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test Ipv6rangeAPIService List", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.Ipv6rangeAPI.List(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test Ipv6rangeAPIService Read", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		resp, httpRes, err := apiClient.Ipv6rangeAPI.Read(context.Background(), reference).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test Ipv6rangeAPIService Update", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		resp, httpRes, err := apiClient.Ipv6rangeAPI.Update(context.Background(), reference).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
