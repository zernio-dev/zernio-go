# \BroadcastsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBroadcastRecipients**](BroadcastsAPI.md#AddBroadcastRecipients) | **Post** /v1/broadcasts/{broadcastId}/recipients | Add recipients to a broadcast
[**CancelBroadcast**](BroadcastsAPI.md#CancelBroadcast) | **Post** /v1/broadcasts/{broadcastId}/cancel | Cancel broadcast
[**CreateBroadcast**](BroadcastsAPI.md#CreateBroadcast) | **Post** /v1/broadcasts | Create broadcast draft
[**DeleteBroadcast**](BroadcastsAPI.md#DeleteBroadcast) | **Delete** /v1/broadcasts/{broadcastId} | Delete broadcast
[**GetBroadcast**](BroadcastsAPI.md#GetBroadcast) | **Get** /v1/broadcasts/{broadcastId} | Get broadcast details
[**ListBroadcastRecipients**](BroadcastsAPI.md#ListBroadcastRecipients) | **Get** /v1/broadcasts/{broadcastId}/recipients | List broadcast recipients
[**ListBroadcasts**](BroadcastsAPI.md#ListBroadcasts) | **Get** /v1/broadcasts | List broadcasts
[**ScheduleBroadcast**](BroadcastsAPI.md#ScheduleBroadcast) | **Post** /v1/broadcasts/{broadcastId}/schedule | Schedule broadcast for later
[**SendBroadcast**](BroadcastsAPI.md#SendBroadcast) | **Post** /v1/broadcasts/{broadcastId}/send | Send broadcast now
[**UpdateBroadcast**](BroadcastsAPI.md#UpdateBroadcast) | **Patch** /v1/broadcasts/{broadcastId} | Update broadcast



## AddBroadcastRecipients

> AddBroadcastRecipients200Response AddBroadcastRecipients(ctx, broadcastId).AddBroadcastRecipientsRequest(addBroadcastRecipientsRequest).Execute()

Add recipients to a broadcast



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
	broadcastId := "broadcastId_example" // string | 
	addBroadcastRecipientsRequest := *openapiclient.NewAddBroadcastRecipientsRequest() // AddBroadcastRecipientsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.AddBroadcastRecipients(context.Background(), broadcastId).AddBroadcastRecipientsRequest(addBroadcastRecipientsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.AddBroadcastRecipients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddBroadcastRecipients`: AddBroadcastRecipients200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.AddBroadcastRecipients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddBroadcastRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addBroadcastRecipientsRequest** | [**AddBroadcastRecipientsRequest**](AddBroadcastRecipientsRequest.md) |  | 

### Return type

[**AddBroadcastRecipients200Response**](AddBroadcastRecipients200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelBroadcast

> CancelBroadcast200Response CancelBroadcast(ctx, broadcastId).Execute()

Cancel broadcast



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
	broadcastId := "broadcastId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.CancelBroadcast(context.Background(), broadcastId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.CancelBroadcast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelBroadcast`: CancelBroadcast200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.CancelBroadcast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelBroadcastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CancelBroadcast200Response**](CancelBroadcast200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBroadcast

> CreateBroadcast200Response CreateBroadcast(ctx).CreateBroadcastRequest(createBroadcastRequest).Execute()

Create broadcast draft



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
	createBroadcastRequest := *openapiclient.NewCreateBroadcastRequest("ProfileId_example", "AccountId_example", "Platform_example", "Name_example") // CreateBroadcastRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.CreateBroadcast(context.Background()).CreateBroadcastRequest(createBroadcastRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.CreateBroadcast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBroadcast`: CreateBroadcast200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.CreateBroadcast`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBroadcastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBroadcastRequest** | [**CreateBroadcastRequest**](CreateBroadcastRequest.md) |  | 

### Return type

[**CreateBroadcast200Response**](CreateBroadcast200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBroadcast

> DeleteBroadcast(ctx, broadcastId).Execute()

Delete broadcast



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
	broadcastId := "broadcastId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BroadcastsAPI.DeleteBroadcast(context.Background(), broadcastId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.DeleteBroadcast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBroadcastRequest struct via the builder pattern


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


## GetBroadcast

> GetBroadcast200Response GetBroadcast(ctx, broadcastId).Execute()

Get broadcast details



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
	broadcastId := "broadcastId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.GetBroadcast(context.Background(), broadcastId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.GetBroadcast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBroadcast`: GetBroadcast200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.GetBroadcast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBroadcastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBroadcast200Response**](GetBroadcast200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBroadcastRecipients

> ListBroadcastRecipients200Response ListBroadcastRecipients(ctx, broadcastId).Status(status).Limit(limit).Skip(skip).Execute()

List broadcast recipients



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
	broadcastId := "broadcastId_example" // string | 
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	skip := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.ListBroadcastRecipients(context.Background(), broadcastId).Status(status).Limit(limit).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.ListBroadcastRecipients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBroadcastRecipients`: ListBroadcastRecipients200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.ListBroadcastRecipients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBroadcastRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **skip** | **int32** |  | [default to 0]

### Return type

[**ListBroadcastRecipients200Response**](ListBroadcastRecipients200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBroadcasts

> ListBroadcasts200Response ListBroadcasts(ctx).ProfileId(profileId).Status(status).Platform(platform).Limit(limit).Skip(skip).Execute()

List broadcasts



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
	status := "status_example" // string |  (optional)
	platform := "platform_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	skip := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.ListBroadcasts(context.Background()).ProfileId(profileId).Status(status).Platform(platform).Limit(limit).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.ListBroadcasts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBroadcasts`: ListBroadcasts200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.ListBroadcasts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBroadcastsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Filter by profile. Omit to list across all profiles | 
 **status** | **string** |  | 
 **platform** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **skip** | **int32** |  | [default to 0]

### Return type

[**ListBroadcasts200Response**](ListBroadcasts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleBroadcast

> ScheduleBroadcast200Response ScheduleBroadcast(ctx, broadcastId).ScheduleBroadcastRequest(scheduleBroadcastRequest).Execute()

Schedule broadcast for later



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/zernio-dev/zernio-go"
)

func main() {
	broadcastId := "broadcastId_example" // string | 
	scheduleBroadcastRequest := *openapiclient.NewScheduleBroadcastRequest(time.Now()) // ScheduleBroadcastRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.ScheduleBroadcast(context.Background(), broadcastId).ScheduleBroadcastRequest(scheduleBroadcastRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.ScheduleBroadcast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleBroadcast`: ScheduleBroadcast200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.ScheduleBroadcast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleBroadcastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduleBroadcastRequest** | [**ScheduleBroadcastRequest**](ScheduleBroadcastRequest.md) |  | 

### Return type

[**ScheduleBroadcast200Response**](ScheduleBroadcast200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendBroadcast

> SendBroadcast200Response SendBroadcast(ctx, broadcastId).Execute()

Send broadcast now



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
	broadcastId := "broadcastId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.SendBroadcast(context.Background(), broadcastId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.SendBroadcast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendBroadcast`: SendBroadcast200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.SendBroadcast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendBroadcastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SendBroadcast200Response**](SendBroadcast200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBroadcast

> UpdateBroadcast200Response UpdateBroadcast(ctx, broadcastId).UpdateBroadcastRequest(updateBroadcastRequest).Execute()

Update broadcast



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
	broadcastId := "broadcastId_example" // string | 
	updateBroadcastRequest := *openapiclient.NewUpdateBroadcastRequest() // UpdateBroadcastRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.UpdateBroadcast(context.Background(), broadcastId).UpdateBroadcastRequest(updateBroadcastRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.UpdateBroadcast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBroadcast`: UpdateBroadcast200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.UpdateBroadcast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBroadcastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBroadcastRequest** | [**UpdateBroadcastRequest**](UpdateBroadcastRequest.md) |  | 

### Return type

[**UpdateBroadcast200Response**](UpdateBroadcast200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

