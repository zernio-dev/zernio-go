# \AccountGroupsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountGroup**](AccountGroupsAPI.md#CreateAccountGroup) | **Post** /v1/account-groups | Create group
[**DeleteAccountGroup**](AccountGroupsAPI.md#DeleteAccountGroup) | **Delete** /v1/account-groups/{groupId} | Delete group
[**ListAccountGroups**](AccountGroupsAPI.md#ListAccountGroups) | **Get** /v1/account-groups | List groups
[**UpdateAccountGroup**](AccountGroupsAPI.md#UpdateAccountGroup) | **Put** /v1/account-groups/{groupId} | Update group



## CreateAccountGroup

> CreateAccountGroup201Response CreateAccountGroup(ctx).CreateAccountGroupRequest(createAccountGroupRequest).Execute()

Create group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zernio-dev/zernio-go"
)

func main() {
	createAccountGroupRequest := *openapiclient.NewCreateAccountGroupRequest("Name_example", []string{"AccountIds_example"}) // CreateAccountGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountGroupsAPI.CreateAccountGroup(context.Background()).CreateAccountGroupRequest(createAccountGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.CreateAccountGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountGroup`: CreateAccountGroup201Response
	fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.CreateAccountGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountGroupRequest** | [**CreateAccountGroupRequest**](CreateAccountGroupRequest.md) |  | 

### Return type

[**CreateAccountGroup201Response**](CreateAccountGroup201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountGroup

> DeleteAccountGroup200Response DeleteAccountGroup(ctx, groupId).Execute()

Delete group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zernio-dev/zernio-go"
)

func main() {
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountGroupsAPI.DeleteAccountGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.DeleteAccountGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccountGroup`: DeleteAccountGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.DeleteAccountGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteAccountGroup200Response**](DeleteAccountGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountGroups

> ListAccountGroups200Response ListAccountGroups(ctx).Execute()

List groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zernio-dev/zernio-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountGroupsAPI.ListAccountGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.ListAccountGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccountGroups`: ListAccountGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.ListAccountGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountGroupsRequest struct via the builder pattern


### Return type

[**ListAccountGroups200Response**](ListAccountGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountGroup

> UpdateAccountGroup200Response UpdateAccountGroup(ctx, groupId).UpdateAccountGroupRequest(updateAccountGroupRequest).Execute()

Update group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zernio-dev/zernio-go"
)

func main() {
	groupId := "groupId_example" // string | 
	updateAccountGroupRequest := *openapiclient.NewUpdateAccountGroupRequest() // UpdateAccountGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountGroupsAPI.UpdateAccountGroup(context.Background(), groupId).UpdateAccountGroupRequest(updateAccountGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.UpdateAccountGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountGroup`: UpdateAccountGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.UpdateAccountGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAccountGroupRequest** | [**UpdateAccountGroupRequest**](UpdateAccountGroupRequest.md) |  | 

### Return type

[**UpdateAccountGroup200Response**](UpdateAccountGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

