/*
Infoblox THREATINSIGHT API

Testing ThreatinsightCloudclientAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package threatinsight

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/infoblox-nios-go-client/threatinsight"
)

func TestThreatinsightCloudclientAPIService(t *testing.T) {

	apiClient := threatinsight.NewAPIClient()

	t.Run("Test ThreatinsightCloudclientAPIService List", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ThreatinsightCloudclientAPI.List(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThreatinsightCloudclientAPIService Read", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		resp, httpRes, err := apiClient.ThreatinsightCloudclientAPI.Read(context.Background(), reference).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThreatinsightCloudclientAPIService Update", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reference string

		resp, httpRes, err := apiClient.ThreatinsightCloudclientAPI.Update(context.Background(), reference).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
