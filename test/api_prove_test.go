/*
Superlink

Testing ProveAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package superlink

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/superlink-me/superlink-api-go-client"
)

func Test_superlink_ProveAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProveAPIService GetEthWalletMessage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerAddress string

		resp, httpRes, err := apiClient.ProveAPI.GetEthWalletMessage(context.Background(), ownerAddress).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
