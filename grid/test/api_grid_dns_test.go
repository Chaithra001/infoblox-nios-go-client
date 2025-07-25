/*
Infoblox GRID API

Testing GridDnsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package grid

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
)

func TestGridDnsAPIService(t *testing.T) {

	apiClient := grid.NewAPIClient()

	t.Run("Test GridDnsAPIService List", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GridDnsAPI.List(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridDnsAPIService Read", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		resp, httpRes, err := apiClient.GridDnsAPI.Read(context.Background(), reference).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridDnsAPIService Update", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		resp, httpRes, err := apiClient.GridDnsAPI.Update(context.Background(), reference).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
