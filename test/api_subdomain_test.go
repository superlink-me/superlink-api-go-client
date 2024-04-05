/*
Superlink

Testing SubdomainAPIService

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

func Test_superlink_SubdomainAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SubdomainAPIService ParentdomainPurchase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var parentDomain string

		httpRes, err := apiClient.SubdomainAPI.ParentdomainPurchase(context.Background(), parentDomain).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubdomainAPIService ParentdomainSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var parentDomain string

		resp, httpRes, err := apiClient.SubdomainAPI.ParentdomainSearch(context.Background(), parentDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubdomainAPIService ParentdomainValidation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var parentDomain string

		resp, httpRes, err := apiClient.SubdomainAPI.ParentdomainValidation(context.Background(), parentDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubdomainAPIService SubdomainAvailable", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var parentDomain string
		var subDomainName string

		resp, httpRes, err := apiClient.SubdomainAPI.SubdomainAvailable(context.Background(), parentDomain, subDomainName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubdomainAPIService SubdomainClaimed", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var parentDomain string
		var ethAddress string

		resp, httpRes, err := apiClient.SubdomainAPI.SubdomainClaimed(context.Background(), parentDomain, ethAddress).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubdomainAPIService SubdomainList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var parentDomain string

		resp, httpRes, err := apiClient.SubdomainAPI.SubdomainList(context.Background(), parentDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubdomainAPIService SubdomainMint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var parentDomain string
		var subDomainName string

		httpRes, err := apiClient.SubdomainAPI.SubdomainMint(context.Background(), parentDomain, subDomainName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubdomainAPIService SubdomainMintSig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var parentDomain string
		var subDomainName string

		httpRes, err := apiClient.SubdomainAPI.SubdomainMintSig(context.Background(), parentDomain, subDomainName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubdomainAPIService SubdomainPerDay", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var parentDomain string

		resp, httpRes, err := apiClient.SubdomainAPI.SubdomainPerDay(context.Background(), parentDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubdomainAPIService SubdomainResolve", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var parentDomain string
		var subDomainName string

		resp, httpRes, err := apiClient.SubdomainAPI.SubdomainResolve(context.Background(), parentDomain, subDomainName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubdomainAPIService SubdomainReverseResolve", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var parentDomain string
		var ethAddress string

		resp, httpRes, err := apiClient.SubdomainAPI.SubdomainReverseResolve(context.Background(), parentDomain, ethAddress).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubdomainAPIService SubdomainRevoke", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var parentDomain string
		var subDomainName string

		httpRes, err := apiClient.SubdomainAPI.SubdomainRevoke(context.Background(), parentDomain, subDomainName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubdomainAPIService SubdomainTotal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var parentDomain string

		resp, httpRes, err := apiClient.SubdomainAPI.SubdomainTotal(context.Background(), parentDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
