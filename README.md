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
*ApisResourceApi* | [**ApisResourcesGet**](docs/ApisResourceApi.md#apisresourcesget) | **Get** /apis/resources | 
*AuthorizationResourceApi* | [**ApisAuthorizationV1NamespacesNamespaceLocalresourceaccessreviewsPost**](docs/AuthorizationResourceApi.md#apisauthorizationv1namespacesnamespacelocalresourceaccessreviewspost) | **Post** /apis/authorization/v1/namespaces/{namespace}/localresourceaccessreviews | 
*AuthorizationResourceApi* | [**ApisAuthorizationV1ResourceaccessreviewsPost**](docs/AuthorizationResourceApi.md#apisauthorizationv1resourceaccessreviewspost) | **Post** /apis/authorization/v1/resourceaccessreviews | 
*AuthorizationResourceApi* | [**ApisAuthorizationV1SelfsubjectaccessreviewsPost**](docs/AuthorizationResourceApi.md#apisauthorizationv1selfsubjectaccessreviewspost) | **Post** /apis/authorization/v1/selfsubjectaccessreviews | 
*AuthorizationResourceApi* | [**ApisAuthorizationV1SubjectaccessreviewsPost**](docs/AuthorizationResourceApi.md#apisauthorizationv1subjectaccessreviewspost) | **Post** /apis/authorization/v1/subjectaccessreviews | 
*KafkaAppsResourceApi* | [**ApisKafkaAppsV1AppsGet**](docs/KafkaAppsResourceApi.md#apiskafkaappsv1appsget) | **Get** /apis/kafka.apps/v1/apps | 
*KafkaAppsResourceApi* | [**ApisKafkaAppsV1NamespacesNamespaceAppsGet**](docs/KafkaAppsResourceApi.md#apiskafkaappsv1namespacesnamespaceappsget) | **Get** /apis/kafka.apps/v1/namespaces/{namespace}/apps | 
*KafkaAppsResourceApi* | [**ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete**](docs/KafkaAppsResourceApi.md#apiskafkaappsv1namespacesnamespaceappsnamedelete) | **Delete** /apis/kafka.apps/v1/namespaces/{namespace}/apps/{name} | 
*KafkaAppsResourceApi* | [**ApisKafkaAppsV1NamespacesNamespaceAppsNameGet**](docs/KafkaAppsResourceApi.md#apiskafkaappsv1namespacesnamespaceappsnameget) | **Get** /apis/kafka.apps/v1/namespaces/{namespace}/apps/{name} | 
*KafkaAppsResourceApi* | [**ApisKafkaAppsV1NamespacesNamespaceAppsNamePut**](docs/KafkaAppsResourceApi.md#apiskafkaappsv1namespacesnamespaceappsnameput) | **Put** /apis/kafka.apps/v1/namespaces/{namespace}/apps/{name} | 
*KafkaAppsResourceApi* | [**ApisKafkaAppsV1NamespacesNamespaceAppsPost**](docs/KafkaAppsResourceApi.md#apiskafkaappsv1namespacesnamespaceappspost) | **Post** /apis/kafka.apps/v1/namespaces/{namespace}/apps | 
*KafkaRbacResourceApi* | [**ApisKafkaRbacV1NamespacesNamespaceBindingsPost**](docs/KafkaRbacResourceApi.md#apiskafkarbacv1namespacesnamespacebindingspost) | **Post** /apis/kafka.rbac/v1/namespaces/{namespace}/bindings | 
*KafkaRbacResourceApi* | [**ApisKafkaRbacV1NamespacesNamespaceResourcesGet**](docs/KafkaRbacResourceApi.md#apiskafkarbacv1namespacesnamespaceresourcesget) | **Get** /apis/kafka.rbac/v1/namespaces/{namespace}/resources | 
*KafkaRbacResourceApi* | [**ApisKafkaRbacV1ResourcesGet**](docs/KafkaRbacResourceApi.md#apiskafkarbacv1resourcesget) | **Get** /apis/kafka.rbac/v1/resources | 
*KafkaTopicsResourceApi* | [**ApisKafkaTopicV1NamespacesNamespaceTopicsGet**](docs/KafkaTopicsResourceApi.md#apiskafkatopicv1namespacesnamespacetopicsget) | **Get** /apis/kafka.topic/v1/namespaces/{namespace}/topics | 
*KafkaTopicsResourceApi* | [**ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost**](docs/KafkaTopicsResourceApi.md#apiskafkatopicv1namespacesnamespacetopicsnameconfigsalterpost) | **Post** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name}/configs:alter | 
*KafkaTopicsResourceApi* | [**ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete**](docs/KafkaTopicsResourceApi.md#apiskafkatopicv1namespacesnamespacetopicsnamedelete) | **Delete** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name} | 
*KafkaTopicsResourceApi* | [**ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet**](docs/KafkaTopicsResourceApi.md#apiskafkatopicv1namespacesnamespacetopicsnamedescribeget) | **Get** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name}/describe | 
*KafkaTopicsResourceApi* | [**ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet**](docs/KafkaTopicsResourceApi.md#apiskafkatopicv1namespacesnamespacetopicsnameget) | **Get** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name} | 
*KafkaTopicsResourceApi* | [**ApisKafkaTopicV1NamespacesNamespaceTopicsPost**](docs/KafkaTopicsResourceApi.md#apiskafkatopicv1namespacesnamespacetopicspost) | **Post** /apis/kafka.topic/v1/namespaces/{namespace}/topics | 
*KafkaTopicsResourceApi* | [**ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost**](docs/KafkaTopicsResourceApi.md#apiskafkatopicv1namespacesnamespacetopicssyncpost) | **Post** /apis/kafka.topic/v1/namespaces/{namespace}/topics/sync | 
*KafkaTopicsResourceApi* | [**ApisKafkaTopicV1TopicsGet**](docs/KafkaTopicsResourceApi.md#apiskafkatopicv1topicsget) | **Get** /apis/kafka.topic/v1/topics | 
*ProjectsResourceApi* | [**ApisProjectV1ProjectsGet**](docs/ProjectsResourceApi.md#apisprojectv1projectsget) | **Get** /apis/project/v1/projects | 
*ProjectsResourceApi* | [**ApisProjectV1ProjectsNameDelete**](docs/ProjectsResourceApi.md#apisprojectv1projectsnamedelete) | **Delete** /apis/project/v1/projects/{name} | 
*ProjectsResourceApi* | [**ApisProjectV1ProjectsNameGet**](docs/ProjectsResourceApi.md#apisprojectv1projectsnameget) | **Get** /apis/project/v1/projects/{name} | 
*ProjectsResourceApi* | [**ApisProjectV1ProjectsPost**](docs/ProjectsResourceApi.md#apisprojectv1projectspost) | **Post** /apis/project/v1/projects | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1ClusterrolebindingsGet**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1clusterrolebindingsget) | **Get** /apis/rbac.authorization/v1/clusterrolebindings | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1ClusterrolebindingsNameDelete**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1clusterrolebindingsnamedelete) | **Delete** /apis/rbac.authorization/v1/clusterrolebindings/{name} | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1ClusterrolebindingsNameGet**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1clusterrolebindingsnameget) | **Get** /apis/rbac.authorization/v1/clusterrolebindings/{name} | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1ClusterrolebindingsNamePut**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1clusterrolebindingsnameput) | **Put** /apis/rbac.authorization/v1/clusterrolebindings/{name} | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1ClusterrolesGet**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1clusterrolesget) | **Get** /apis/rbac.authorization/v1/clusterroles | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1ClusterrolesNameGet**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1clusterrolesnameget) | **Get** /apis/rbac.authorization/v1/clusterroles/{name} | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1namespacesnamespacerolebindingsget) | **Get** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameDelete**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1namespacesnamespacerolebindingsnamedelete) | **Delete** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings/{name} | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameGet**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1namespacesnamespacerolebindingsnameget) | **Get** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings/{name} | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNamePut**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1namespacesnamespacerolebindingsnameput) | **Put** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings/{name} | 
*RbacAuthResourceApi* | [**ApisRbacAuthorizationV1RolebindingsGet**](docs/RbacAuthResourceApi.md#apisrbacauthorizationv1rolebindingsget) | **Get** /apis/rbac.authorization/v1/rolebindings | 
*UserGroupsResourceApi* | [**ApisUserV1GroupsGet**](docs/UserGroupsResourceApi.md#apisuserv1groupsget) | **Get** /apis/user/v1/groups | 
*UserGroupsResourceApi* | [**ApisUserV1GroupsNameGet**](docs/UserGroupsResourceApi.md#apisuserv1groupsnameget) | **Get** /apis/user/v1/groups/{name} | 
*UserGroupsResourceApi* | [**ApisUserV1GroupsNamePut**](docs/UserGroupsResourceApi.md#apisuserv1groupsnameput) | **Put** /apis/user/v1/groups/{name} | 
*UserGroupsResourceApi* | [**ApisUserV1GroupsPost**](docs/UserGroupsResourceApi.md#apisuserv1groupspost) | **Post** /apis/user/v1/groups | 
*UserGroupsResourceApi* | [**ApisUserV1UsersGet**](docs/UserGroupsResourceApi.md#apisuserv1usersget) | **Get** /apis/user/v1/users | 
*UserGroupsResourceApi* | [**ApisUserV1UsersNameGet**](docs/UserGroupsResourceApi.md#apisuserv1usersnameget) | **Get** /apis/user/v1/users/{name} | 


## Documentation For Models

 - [ApiResource](docs/ApiResource.md)
 - [ApiResourceList](docs/ApiResourceList.md)
 - [ClusterRole](docs/ClusterRole.md)
 - [ClusterRoleBinding](docs/ClusterRoleBinding.md)
 - [ClusterRoleBindingList](docs/ClusterRoleBindingList.md)
 - [ClusterRoleList](docs/ClusterRoleList.md)
 - [ConfigItem](docs/ConfigItem.md)
 - [Group](docs/Group.md)
 - [GroupList](docs/GroupList.md)
 - [KafkaApplication](docs/KafkaApplication.md)
 - [KafkaApplicationList](docs/KafkaApplicationList.md)
 - [KafkaRbacSummary](docs/KafkaRbacSummary.md)
 - [KafkaRbacSummaryList](docs/KafkaRbacSummaryList.md)
 - [KafkaRoleBindingRequest](docs/KafkaRoleBindingRequest.md)
 - [KafkaTopic](docs/KafkaTopic.md)
 - [KafkaTopicList](docs/KafkaTopicList.md)
 - [KafkaTopicRequest](docs/KafkaTopicRequest.md)
 - [LocalResourceAccessReview](docs/LocalResourceAccessReview.md)
 - [NonResourceAttributes](docs/NonResourceAttributes.md)
 - [PartitionInfo](docs/PartitionInfo.md)
 - [PolicyRule](docs/PolicyRule.md)
 - [Project](docs/Project.md)
 - [ProjectList](docs/ProjectList.md)
 - [ResourceAccessReview](docs/ResourceAccessReview.md)
 - [ResourceAccessReviewResponse](docs/ResourceAccessReviewResponse.md)
 - [ResourceAttributes](docs/ResourceAttributes.md)
 - [ResourcePattern](docs/ResourcePattern.md)
 - [RoleBinding](docs/RoleBinding.md)
 - [RoleBindingList](docs/RoleBindingList.md)
 - [RoleRef](docs/RoleRef.md)
 - [SelfSubjectAccessReview](docs/SelfSubjectAccessReview.md)
 - [SelfSubjectAccessReviewSpec](docs/SelfSubjectAccessReviewSpec.md)
 - [Status](docs/Status.md)
 - [StatusCause](docs/StatusCause.md)
 - [StatusDetails](docs/StatusDetails.md)
 - [Subject](docs/Subject.md)
 - [SubjectAccessReview](docs/SubjectAccessReview.md)
 - [SubjectAccessReviewResponse](docs/SubjectAccessReviewResponse.md)
 - [SubjectAccessReviewStatus](docs/SubjectAccessReviewStatus.md)
 - [User](docs/User.md)
 - [UserList](docs/UserList.md)


## Documentation For Authorization


Authentication schemes defined for the API:
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



