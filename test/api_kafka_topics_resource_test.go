/*
kafka-self-service API

Testing KafkaTopicsResourceApiService

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

func Test_openapi_KafkaTopicsResourceApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KafkaTopicsResourceApiService ApisKafkaTopicV1NamespacesNamespaceTopicsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsGet(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KafkaTopicsResourceApiService ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KafkaTopicsResourceApiService ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KafkaTopicsResourceApiService ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KafkaTopicsResourceApiService ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KafkaTopicsResourceApiService ApisKafkaTopicV1NamespacesNamespaceTopicsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsPost(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KafkaTopicsResourceApiService ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KafkaTopicsResourceApiService ApisKafkaTopicV1TopicsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1TopicsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
