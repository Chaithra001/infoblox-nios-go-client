/*
Infoblox MISC API

Testing ScheduledtaskAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package misc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"
)

func TestScheduledtaskAPIService(t *testing.T) {

	apiClient := misc.NewAPIClient()

	t.Run("Test ScheduledtaskAPIService Delete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		httpRes, err := apiClient.ScheduledtaskAPI.Delete(context.Background(), reference).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScheduledtaskAPIService List", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ScheduledtaskAPI.List(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScheduledtaskAPIService Read", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		resp, httpRes, err := apiClient.ScheduledtaskAPI.Read(context.Background(), reference).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScheduledtaskAPIService Update", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		resp, httpRes, err := apiClient.ScheduledtaskAPI.Update(context.Background(), reference).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
