/*
kafka-self-service API

Testing ProjectsResourceApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/shnplr/kssapi"
)

func Test_openapi_ProjectsResourceApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectsResourceApiService ApisProjectV1ProjectsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsResourceApi.ApisProjectV1ProjectsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsResourceApiService ApisProjectV1ProjectsNameDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.ProjectsResourceApi.ApisProjectV1ProjectsNameDelete(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsResourceApiService ApisProjectV1ProjectsNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.ProjectsResourceApi.ApisProjectV1ProjectsNameGet(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsResourceApiService ApisProjectV1ProjectsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsResourceApi.ApisProjectV1ProjectsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
