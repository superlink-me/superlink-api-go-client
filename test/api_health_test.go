/*
Superlink API

Testing HealthAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ResolutionAPI

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/superlink-me/resolution-api-go-client"
)

func Test_ResolutionAPI_HealthAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HealthAPIService HealthCheck", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.HealthAPI.HealthCheck(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
