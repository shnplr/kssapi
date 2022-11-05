/*
kafka-self-service API

Testing KafkaRbacResourceApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_openapi_KafkaRbacResourceApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test KafkaRbacResourceApiService ApisKafkaRbacV1NamespacesNameBindingsPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var name string

        resp, httpRes, err := apiClient.KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNameBindingsPost(context.Background(), name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
