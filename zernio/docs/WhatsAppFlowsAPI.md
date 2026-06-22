# \WhatsAppFlowsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWhatsAppFlow**](WhatsAppFlowsAPI.md#CreateWhatsAppFlow) | **Post** /v1/whatsapp/flows | Create flow
[**DeleteWhatsAppFlow**](WhatsAppFlowsAPI.md#DeleteWhatsAppFlow) | **Delete** /v1/whatsapp/flows/{flowId} | Delete flow
[**DeprecateWhatsAppFlow**](WhatsAppFlowsAPI.md#DeprecateWhatsAppFlow) | **Post** /v1/whatsapp/flows/{flowId}/deprecate | Deprecate flow
[**GetWhatsAppFlow**](WhatsAppFlowsAPI.md#GetWhatsAppFlow) | **Get** /v1/whatsapp/flows/{flowId} | Get flow
[**GetWhatsAppFlowJson**](WhatsAppFlowsAPI.md#GetWhatsAppFlowJson) | **Get** /v1/whatsapp/flows/{flowId}/json | Get flow JSON asset
[**GetWhatsAppFlowPreview**](WhatsAppFlowsAPI.md#GetWhatsAppFlowPreview) | **Get** /v1/whatsapp/flows/{flowId}/preview | Get flow preview URL
[**ListWhatsAppFlowResponses**](WhatsAppFlowsAPI.md#ListWhatsAppFlowResponses) | **Get** /v1/whatsapp/flow-responses | List flow responses
[**ListWhatsAppFlowVersions**](WhatsAppFlowsAPI.md#ListWhatsAppFlowVersions) | **Get** /v1/whatsapp/flows/{flowId}/versions | List flow versions
[**ListWhatsAppFlows**](WhatsAppFlowsAPI.md#ListWhatsAppFlows) | **Get** /v1/whatsapp/flows | List flows
[**PublishWhatsAppFlow**](WhatsAppFlowsAPI.md#PublishWhatsAppFlow) | **Post** /v1/whatsapp/flows/{flowId}/publish | Publish flow
[**SendWhatsAppFlowMessage**](WhatsAppFlowsAPI.md#SendWhatsAppFlowMessage) | **Post** /v1/whatsapp/flows/send | Send flow message
[**UpdateWhatsAppFlow**](WhatsAppFlowsAPI.md#UpdateWhatsAppFlow) | **Patch** /v1/whatsapp/flows/{flowId} | Update flow
[**UploadWhatsAppFlowJson**](WhatsAppFlowsAPI.md#UploadWhatsAppFlowJson) | **Put** /v1/whatsapp/flows/{flowId}/json | Upload flow JSON



## CreateWhatsAppFlow

> CreateWhatsAppFlow200Response CreateWhatsAppFlow(ctx).CreateWhatsAppFlowRequest(createWhatsAppFlowRequest).Execute()

Create flow



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
	createWhatsAppFlowRequest := *openapiclient.NewCreateWhatsAppFlowRequest("AccountId_example", "Name_example", []string{"Categories_example"}) // CreateWhatsAppFlowRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppFlowsAPI.CreateWhatsAppFlow(context.Background()).CreateWhatsAppFlowRequest(createWhatsAppFlowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppFlowsAPI.CreateWhatsAppFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWhatsAppFlow`: CreateWhatsAppFlow200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppFlowsAPI.CreateWhatsAppFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWhatsAppFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWhatsAppFlowRequest** | [**CreateWhatsAppFlowRequest**](CreateWhatsAppFlowRequest.md) |  | 

### Return type

[**CreateWhatsAppFlow200Response**](CreateWhatsAppFlow200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWhatsAppFlow

> UpdateYoutubeDefaultPlaylist200Response DeleteWhatsAppFlow(ctx, flowId).AccountId(accountId).Execute()

Delete flow



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
	flowId := "flowId_example" // string | Flow ID
	accountId := "accountId_example" // string | WhatsApp social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppFlowsAPI.DeleteWhatsAppFlow(context.Background(), flowId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppFlowsAPI.DeleteWhatsAppFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWhatsAppFlow`: UpdateYoutubeDefaultPlaylist200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppFlowsAPI.DeleteWhatsAppFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWhatsAppFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**UpdateYoutubeDefaultPlaylist200Response**](UpdateYoutubeDefaultPlaylist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeprecateWhatsAppFlow

> UpdateYoutubeDefaultPlaylist200Response DeprecateWhatsAppFlow(ctx, flowId).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()

Deprecate flow



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
	flowId := "flowId_example" // string | Flow ID
	sendTypingIndicatorRequest := *openapiclient.NewSendTypingIndicatorRequest("AccountId_example") // SendTypingIndicatorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppFlowsAPI.DeprecateWhatsAppFlow(context.Background(), flowId).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppFlowsAPI.DeprecateWhatsAppFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeprecateWhatsAppFlow`: UpdateYoutubeDefaultPlaylist200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppFlowsAPI.DeprecateWhatsAppFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecateWhatsAppFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendTypingIndicatorRequest** | [**SendTypingIndicatorRequest**](SendTypingIndicatorRequest.md) |  | 

### Return type

[**UpdateYoutubeDefaultPlaylist200Response**](UpdateYoutubeDefaultPlaylist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppFlow

> GetWhatsAppFlow200Response GetWhatsAppFlow(ctx, flowId).AccountId(accountId).Fields(fields).Execute()

Get flow



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
	flowId := "flowId_example" // string | Flow ID
	accountId := "accountId_example" // string | WhatsApp social account ID
	fields := "fields_example" // string | Comma-separated fields to return (default: id,name,status,categories,validation_errors,json_version,preview,data_api_version,endpoint_uri) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppFlowsAPI.GetWhatsAppFlow(context.Background(), flowId).AccountId(accountId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppFlowsAPI.GetWhatsAppFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppFlow`: GetWhatsAppFlow200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppFlowsAPI.GetWhatsAppFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 
 **fields** | **string** | Comma-separated fields to return (default: id,name,status,categories,validation_errors,json_version,preview,data_api_version,endpoint_uri) | 

### Return type

[**GetWhatsAppFlow200Response**](GetWhatsAppFlow200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppFlowJson

> GetWhatsAppFlowJson200Response GetWhatsAppFlowJson(ctx, flowId).AccountId(accountId).Execute()

Get flow JSON asset



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
	flowId := "flowId_example" // string | Flow ID
	accountId := "accountId_example" // string | WhatsApp social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppFlowsAPI.GetWhatsAppFlowJson(context.Background(), flowId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppFlowsAPI.GetWhatsAppFlowJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppFlowJson`: GetWhatsAppFlowJson200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppFlowsAPI.GetWhatsAppFlowJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppFlowJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**GetWhatsAppFlowJson200Response**](GetWhatsAppFlowJson200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppFlowPreview

> GetWhatsAppFlowPreview200Response GetWhatsAppFlowPreview(ctx, flowId).AccountId(accountId).Invalidate(invalidate).Execute()

Get flow preview URL



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
	flowId := "flowId_example" // string | Flow ID
	accountId := "accountId_example" // string | WhatsApp social account ID
	invalidate := true // bool | Mint a fresh preview link (default false) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppFlowsAPI.GetWhatsAppFlowPreview(context.Background(), flowId).AccountId(accountId).Invalidate(invalidate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppFlowsAPI.GetWhatsAppFlowPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppFlowPreview`: GetWhatsAppFlowPreview200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppFlowsAPI.GetWhatsAppFlowPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppFlowPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 
 **invalidate** | **bool** | Mint a fresh preview link (default false) | 

### Return type

[**GetWhatsAppFlowPreview200Response**](GetWhatsAppFlowPreview200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWhatsAppFlowResponses

> ListWhatsAppFlowResponses200Response ListWhatsAppFlowResponses(ctx).AccountId(accountId).FlowId(flowId).Limit(limit).Execute()

List flow responses



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
	accountId := "accountId_example" // string | WhatsApp social account ID
	flowId := "flowId_example" // string | Scope to responses for this flow (optional)
	limit := int32(56) // int32 | Max responses to return (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppFlowsAPI.ListWhatsAppFlowResponses(context.Background()).AccountId(accountId).FlowId(flowId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppFlowsAPI.ListWhatsAppFlowResponses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWhatsAppFlowResponses`: ListWhatsAppFlowResponses200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppFlowsAPI.ListWhatsAppFlowResponses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWhatsAppFlowResponsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | WhatsApp social account ID | 
 **flowId** | **string** | Scope to responses for this flow | 
 **limit** | **int32** | Max responses to return | [default to 50]

### Return type

[**ListWhatsAppFlowResponses200Response**](ListWhatsAppFlowResponses200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWhatsAppFlowVersions

> ListWhatsAppFlowVersions200Response ListWhatsAppFlowVersions(ctx, flowId).AccountId(accountId).Execute()

List flow versions



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
	flowId := "flowId_example" // string | Flow ID
	accountId := "accountId_example" // string | WhatsApp social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppFlowsAPI.ListWhatsAppFlowVersions(context.Background(), flowId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppFlowsAPI.ListWhatsAppFlowVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWhatsAppFlowVersions`: ListWhatsAppFlowVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppFlowsAPI.ListWhatsAppFlowVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWhatsAppFlowVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**ListWhatsAppFlowVersions200Response**](ListWhatsAppFlowVersions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWhatsAppFlows

> ListWhatsAppFlows200Response ListWhatsAppFlows(ctx).AccountId(accountId).Execute()

List flows



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
	accountId := "accountId_example" // string | WhatsApp social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppFlowsAPI.ListWhatsAppFlows(context.Background()).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppFlowsAPI.ListWhatsAppFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWhatsAppFlows`: ListWhatsAppFlows200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppFlowsAPI.ListWhatsAppFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWhatsAppFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**ListWhatsAppFlows200Response**](ListWhatsAppFlows200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishWhatsAppFlow

> UpdateYoutubeDefaultPlaylist200Response PublishWhatsAppFlow(ctx, flowId).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()

Publish flow



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
	flowId := "flowId_example" // string | Flow ID
	sendTypingIndicatorRequest := *openapiclient.NewSendTypingIndicatorRequest("AccountId_example") // SendTypingIndicatorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppFlowsAPI.PublishWhatsAppFlow(context.Background(), flowId).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppFlowsAPI.PublishWhatsAppFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishWhatsAppFlow`: UpdateYoutubeDefaultPlaylist200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppFlowsAPI.PublishWhatsAppFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishWhatsAppFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendTypingIndicatorRequest** | [**SendTypingIndicatorRequest**](SendTypingIndicatorRequest.md) |  | 

### Return type

[**UpdateYoutubeDefaultPlaylist200Response**](UpdateYoutubeDefaultPlaylist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendWhatsAppFlowMessage

> SendWhatsAppFlowMessage200Response SendWhatsAppFlowMessage(ctx).SendWhatsAppFlowMessageRequest(sendWhatsAppFlowMessageRequest).Execute()

Send flow message



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
	sendWhatsAppFlowMessageRequest := *openapiclient.NewSendWhatsAppFlowMessageRequest("AccountId_example", "To_example", "FlowId_example", "FlowCta_example", "Body_example") // SendWhatsAppFlowMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppFlowsAPI.SendWhatsAppFlowMessage(context.Background()).SendWhatsAppFlowMessageRequest(sendWhatsAppFlowMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppFlowsAPI.SendWhatsAppFlowMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendWhatsAppFlowMessage`: SendWhatsAppFlowMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppFlowsAPI.SendWhatsAppFlowMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendWhatsAppFlowMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendWhatsAppFlowMessageRequest** | [**SendWhatsAppFlowMessageRequest**](SendWhatsAppFlowMessageRequest.md) |  | 

### Return type

[**SendWhatsAppFlowMessage200Response**](SendWhatsAppFlowMessage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWhatsAppFlow

> UpdateYoutubeDefaultPlaylist200Response UpdateWhatsAppFlow(ctx, flowId).UpdateWhatsAppFlowRequest(updateWhatsAppFlowRequest).Execute()

Update flow



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
	flowId := "flowId_example" // string | Flow ID
	updateWhatsAppFlowRequest := *openapiclient.NewUpdateWhatsAppFlowRequest("AccountId_example") // UpdateWhatsAppFlowRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppFlowsAPI.UpdateWhatsAppFlow(context.Background(), flowId).UpdateWhatsAppFlowRequest(updateWhatsAppFlowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppFlowsAPI.UpdateWhatsAppFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWhatsAppFlow`: UpdateYoutubeDefaultPlaylist200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppFlowsAPI.UpdateWhatsAppFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWhatsAppFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWhatsAppFlowRequest** | [**UpdateWhatsAppFlowRequest**](UpdateWhatsAppFlowRequest.md) |  | 

### Return type

[**UpdateYoutubeDefaultPlaylist200Response**](UpdateYoutubeDefaultPlaylist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadWhatsAppFlowJson

> UploadWhatsAppFlowJson200Response UploadWhatsAppFlowJson(ctx, flowId).UploadWhatsAppFlowJsonRequest(uploadWhatsAppFlowJsonRequest).Execute()

Upload flow JSON



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
	flowId := "flowId_example" // string | Flow ID
	uploadWhatsAppFlowJsonRequest := *openapiclient.NewUploadWhatsAppFlowJsonRequest("AccountId_example", openapiclient.uploadWhatsAppFlowJson_request_flow_json{MapmapOfStringAny: new(map[string]interface{})}) // UploadWhatsAppFlowJsonRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppFlowsAPI.UploadWhatsAppFlowJson(context.Background(), flowId).UploadWhatsAppFlowJsonRequest(uploadWhatsAppFlowJsonRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppFlowsAPI.UploadWhatsAppFlowJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadWhatsAppFlowJson`: UploadWhatsAppFlowJson200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppFlowsAPI.UploadWhatsAppFlowJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadWhatsAppFlowJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uploadWhatsAppFlowJsonRequest** | [**UploadWhatsAppFlowJsonRequest**](UploadWhatsAppFlowJsonRequest.md) |  | 

### Return type

[**UploadWhatsAppFlowJson200Response**](UploadWhatsAppFlowJson200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

