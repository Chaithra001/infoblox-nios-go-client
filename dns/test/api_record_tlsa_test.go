/*
Infoblox DNS API

Testing RecordTlsaAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dns

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func TestRecordTlsaAPIService(t *testing.T) {

	apiClient := dns.NewAPIClient()

	t.Run("Test RecordTlsaAPIService Create", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RecordTlsaAPI.Create(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RecordTlsaAPIService Delete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		httpRes, err := apiClient.RecordTlsaAPI.Delete(context.Background(), reference).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RecordTlsaAPIService List", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RecordTlsaAPI.List(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RecordTlsaAPIService Read", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		resp, httpRes, err := apiClient.RecordTlsaAPI.Read(context.Background(), reference).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RecordTlsaAPIService Update", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		resp, httpRes, err := apiClient.RecordTlsaAPI.Update(context.Background(), reference).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
