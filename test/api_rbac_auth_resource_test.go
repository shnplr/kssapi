/*
kafka-self-service API

Testing RbacAuthResourceApiService

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

func Test_openapi_RbacAuthResourceApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test RbacAuthResourceApiService ApisRbacAuthorizationV1ClusterrolebindingsDelete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsDelete(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RbacAuthResourceApiService ApisRbacAuthorizationV1ClusterrolebindingsGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RbacAuthResourceApiService ApisRbacAuthorizationV1ClusterrolebindingsPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RbacAuthResourceApiService ApisRbacAuthorizationV1ClusterrolesGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolesGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RbacAuthResourceApiService ApisRbacAuthorizationV1ClusterrolesNameGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var name string

        resp, httpRes, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolesNameGet(context.Background(), name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RbacAuthResourceApiService ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsDelete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var namespace string

        resp, httpRes, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsDelete(context.Background(), namespace).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RbacAuthResourceApiService ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var namespace string

        resp, httpRes, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet(context.Background(), namespace).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RbacAuthResourceApiService ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var namespace string

        resp, httpRes, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost(context.Background(), namespace).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
