# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0-SNAPSHOT
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/shnplr/kssapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:9080*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthorizationResourceApi* | [**ApisAuthorizationV1NamespacesNamespaceLocalresourceaccessreviewsPost**](docs/AuthorizationResourceApi.md#apisauthorizationv1namespacesnamespacelocalresourceaccessreviewspost) | **Post** /apis/authorization/v1/namespaces/{namespace}/localresourceaccessreviews | 
*AuthorizationResourceApi* | [**ApisAuthorizationV1ResourceaccessreviewsPost**](docs/AuthorizationResourceApi.md#apisauthorizationv1resourceaccessreviewspost) | **Post** /apis/authorization/v1/resourceaccessreviews | 
*AuthorizationResourceApi* | [**ApisAuthorizationV1SubjectaccessreviewsPost**](docs/AuthorizationResourceApi.md#apisauthorizationv1subjectaccessreviewspost) | **Post** /apis/authorization/v1/subjectaccessreviews | 
*KafkaRbacResourceApi* | [**ApisKafkaRbacV1NamespacesNamespaceBindingsPost**](docs/KafkaRbacResourceApi.md#apiskafkarbacv1namespacesnamespacebindingspost) | **Post** /apis/kafka.rbac/v1/namespaces/{namespace}/bindings | 
*ProjectsResourceApi* | [**ApisProjectV1ProjectsGet**](docs/ProjectsResourceApi.md#apisprojectv1projectsget) | **Get** /apis/project/v1/projects | 
*ProjectsResourceApi* | [**ApisProjectV1ProjectsNameGet**](docs/ProjectsResourceApi.md#apisprojectv1projectsnameget) | **Get** /apis/project/v1/projects/{name} | 
*ProjectsResourceApi* | [**ApisProjectV1ProjectsPost**](docs/ProjectsResourceApi.md#apisprojectv1projectspost) | **Post** /apis/project/v1/projects | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1ClusterrolebindingsDelete**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1clusterrolebindingsdelete) | **Delete** /apis/rbac.authorization/v1/clusterrolebindings | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1ClusterrolebindingsGet**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1clusterrolebindingsget) | **Get** /apis/rbac.authorization/v1/clusterrolebindings | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1ClusterrolebindingsPost**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1clusterrolebindingspost) | **Post** /apis/rbac.authorization/v1/clusterrolebindings | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1ClusterrolesGet**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1clusterrolesget) | **Get** /apis/rbac.authorization/v1/clusterroles | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1ClusterrolesNameGet**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1clusterrolesnameget) | **Get** /apis/rbac.authorization/v1/clusterroles/{name} | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsDelete**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1namespacesnamespacerolebindingsdelete) | **Delete** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1namespacesnamespacerolebindingsget) | **Get** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1namespacesnamespacerolebindingspost) | **Post** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings | 
*TopicsResourceApi* | [**ApisTopicV1NamespacesNamespaceTopicsGet**](docs/TopicsResourceApi.md#apistopicv1namespacesnamespacetopicsget) | **Get** /apis/topic/v1/namespaces/{namespace}/topics | 
*TopicsResourceApi* | [**ApisTopicV1NamespacesNamespaceTopicsPost**](docs/TopicsResourceApi.md#apistopicv1namespacesnamespacetopicspost) | **Post** /apis/topic/v1/namespaces/{namespace}/topics | 
*UserGroupsResourceApi* | [**ApisUserV1GroupsGet**](docs/UserGroupsResourceApi.md#apisuserv1groupsget) | **Get** /apis/user/v1/groups | 
*UserGroupsResourceApi* | [**ApisUserV1GroupsNameGet**](docs/UserGroupsResourceApi.md#apisuserv1groupsnameget) | **Get** /apis/user/v1/groups/{name} | 
*UserGroupsResourceApi* | [**ApisUserV1GroupsNamePut**](docs/UserGroupsResourceApi.md#apisuserv1groupsnameput) | **Put** /apis/user/v1/groups/{name} | 
*UserGroupsResourceApi* | [**ApisUserV1GroupsPost**](docs/UserGroupsResourceApi.md#apisuserv1groupspost) | **Post** /apis/user/v1/groups | 
*UserGroupsResourceApi* | [**ApisUserV1UsersNameGet**](docs/UserGroupsResourceApi.md#apisuserv1usersnameget) | **Get** /apis/user/v1/users/{name} | 


## Documentation For Models

 - [ApiStatus](docs/ApiStatus.md)
 - [CreateTopicRequestData](docs/CreateTopicRequestData.md)
 - [CreateTopicRequestDataConfigsInner](docs/CreateTopicRequestDataConfigsInner.md)
 - [CreateTopicRequestDataReplicasAssignmentsInner](docs/CreateTopicRequestDataReplicasAssignmentsInner.md)
 - [GenericListGroup](docs/GenericListGroup.md)
 - [GenericListProject](docs/GenericListProject.md)
 - [GenericListRole](docs/GenericListRole.md)
 - [GenericListRoleBinding](docs/GenericListRoleBinding.md)
 - [Group](docs/Group.md)
 - [KafkaTopic](docs/KafkaTopic.md)
 - [Project](docs/Project.md)
 - [RbacRoleBindingRequest](docs/RbacRoleBindingRequest.md)
 - [RbacRoleBindingResponse](docs/RbacRoleBindingResponse.md)
 - [Relationship](docs/Relationship.md)
 - [ResourceAccessReview](docs/ResourceAccessReview.md)
 - [ResourceAccessReviewResponse](docs/ResourceAccessReviewResponse.md)
 - [ResourceMetadata](docs/ResourceMetadata.md)
 - [Role](docs/Role.md)
 - [RoleBinding](docs/RoleBinding.md)
 - [Subject](docs/Subject.md)
 - [SubjectAccessReview](docs/SubjectAccessReview.md)
 - [SubjectAccessReviewResponse](docs/SubjectAccessReviewResponse.md)
 - [TopicData](docs/TopicData.md)
 - [User](docs/User.md)


## Documentation For Authorization



### SecurityScheme


- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: https://login.microsoftonline.com/c6198395-22e6-4778-9809-ac8dcfd76901/v2.0/protocol/openid-connect/auth
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



