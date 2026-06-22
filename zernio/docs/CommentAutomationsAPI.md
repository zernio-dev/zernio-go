# \CommentAutomationsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommentAutomation**](CommentAutomationsAPI.md#CreateCommentAutomation) | **Post** /v1/comment-automations | Create comment-to-DM automation
[**DeleteCommentAutomation**](CommentAutomationsAPI.md#DeleteCommentAutomation) | **Delete** /v1/comment-automations/{automationId} | Delete automation
[**GetCommentAutomation**](CommentAutomationsAPI.md#GetCommentAutomation) | **Get** /v1/comment-automations/{automationId} | Get automation details
[**ListCommentAutomationLogs**](CommentAutomationsAPI.md#ListCommentAutomationLogs) | **Get** /v1/comment-automations/{automationId}/logs | List automation logs
[**ListCommentAutomations**](CommentAutomationsAPI.md#ListCommentAutomations) | **Get** /v1/comment-automations | List comment-to-DM automations
[**UpdateCommentAutomation**](CommentAutomationsAPI.md#UpdateCommentAutomation) | **Patch** /v1/comment-automations/{automationId} | Update automation settings



## CreateCommentAutomation

> CreateCommentAutomation200Response CreateCommentAutomation(ctx).CreateCommentAutomationRequest(createCommentAutomationRequest).Execute()

Create comment-to-DM automation



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
	createCommentAutomationRequest := *openapiclient.NewCreateCommentAutomationRequest("ProfileId_example", "AccountId_example", "Name_example", "DmMessage_example") // CreateCommentAutomationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAutomationsAPI.CreateCommentAutomation(context.Background()).CreateCommentAutomationRequest(createCommentAutomationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAutomationsAPI.CreateCommentAutomation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCommentAutomation`: CreateCommentAutomation200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentAutomationsAPI.CreateCommentAutomation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentAutomationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCommentAutomationRequest** | [**CreateCommentAutomationRequest**](CreateCommentAutomationRequest.md) |  | 

### Return type

[**CreateCommentAutomation200Response**](CreateCommentAutomation200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommentAutomation

> DeleteCommentAutomation(ctx, automationId).Execute()

Delete automation



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
	automationId := "automationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommentAutomationsAPI.DeleteCommentAutomation(context.Background(), automationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAutomationsAPI.DeleteCommentAutomation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**automationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentAutomationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentAutomation

> GetCommentAutomation200Response GetCommentAutomation(ctx, automationId).Execute()

Get automation details



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
	automationId := "automationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAutomationsAPI.GetCommentAutomation(context.Background(), automationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAutomationsAPI.GetCommentAutomation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentAutomation`: GetCommentAutomation200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentAutomationsAPI.GetCommentAutomation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**automationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentAutomationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCommentAutomation200Response**](GetCommentAutomation200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommentAutomationLogs

> ListCommentAutomationLogs200Response ListCommentAutomationLogs(ctx, automationId).Status(status).Limit(limit).Skip(skip).Execute()

List automation logs



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
	automationId := "automationId_example" // string | 
	status := "status_example" // string | Filter by result status (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	skip := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAutomationsAPI.ListCommentAutomationLogs(context.Background(), automationId).Status(status).Limit(limit).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAutomationsAPI.ListCommentAutomationLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCommentAutomationLogs`: ListCommentAutomationLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentAutomationsAPI.ListCommentAutomationLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**automationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCommentAutomationLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter by result status | 
 **limit** | **int32** |  | [default to 50]
 **skip** | **int32** |  | [default to 0]

### Return type

[**ListCommentAutomationLogs200Response**](ListCommentAutomationLogs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommentAutomations

> ListCommentAutomations200Response ListCommentAutomations(ctx).ProfileId(profileId).Execute()

List comment-to-DM automations



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
	profileId := "profileId_example" // string | Filter by profile. Omit to list across all profiles (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAutomationsAPI.ListCommentAutomations(context.Background()).ProfileId(profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAutomationsAPI.ListCommentAutomations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCommentAutomations`: ListCommentAutomations200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentAutomationsAPI.ListCommentAutomations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCommentAutomationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Filter by profile. Omit to list across all profiles | 

### Return type

[**ListCommentAutomations200Response**](ListCommentAutomations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCommentAutomation

> UpdateCommentAutomation200Response UpdateCommentAutomation(ctx, automationId).UpdateCommentAutomationRequest(updateCommentAutomationRequest).Execute()

Update automation settings



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
	automationId := "automationId_example" // string | 
	updateCommentAutomationRequest := *openapiclient.NewUpdateCommentAutomationRequest() // UpdateCommentAutomationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentAutomationsAPI.UpdateCommentAutomation(context.Background(), automationId).UpdateCommentAutomationRequest(updateCommentAutomationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentAutomationsAPI.UpdateCommentAutomation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCommentAutomation`: UpdateCommentAutomation200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentAutomationsAPI.UpdateCommentAutomation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**automationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommentAutomationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCommentAutomationRequest** | [**UpdateCommentAutomationRequest**](UpdateCommentAutomationRequest.md) |  | 

### Return type

[**UpdateCommentAutomation200Response**](UpdateCommentAutomation200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

