# \WebhooksAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookSettings**](WebhooksAPI.md#CreateWebhookSettings) | **Post** /v1/webhooks/settings | Create webhook
[**DeleteWebhookSettings**](WebhooksAPI.md#DeleteWebhookSettings) | **Delete** /v1/webhooks/settings | Delete webhook
[**GetWebhookLogs**](WebhooksAPI.md#GetWebhookLogs) | **Get** /v1/webhooks/logs | List webhook delivery logs
[**GetWebhookSettings**](WebhooksAPI.md#GetWebhookSettings) | **Get** /v1/webhooks/settings | List webhooks
[**TestWebhook**](WebhooksAPI.md#TestWebhook) | **Post** /v1/webhooks/test | Send test webhook
[**UpdateWebhookSettings**](WebhooksAPI.md#UpdateWebhookSettings) | **Put** /v1/webhooks/settings | Update webhook



## CreateWebhookSettings

> UpdateWebhookSettings200Response CreateWebhookSettings(ctx).CreateWebhookSettingsRequest(createWebhookSettingsRequest).Execute()

Create webhook



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
	createWebhookSettingsRequest := *openapiclient.NewCreateWebhookSettingsRequest("Name_example", "Url_example", []string{"Events_example"}) // CreateWebhookSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.CreateWebhookSettings(context.Background()).CreateWebhookSettingsRequest(createWebhookSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CreateWebhookSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhookSettings`: UpdateWebhookSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CreateWebhookSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhookSettingsRequest** | [**CreateWebhookSettingsRequest**](CreateWebhookSettingsRequest.md) |  | 

### Return type

[**UpdateWebhookSettings200Response**](UpdateWebhookSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhookSettings

> UpdateYoutubeDefaultPlaylist200Response DeleteWebhookSettings(ctx).Id(id).Execute()

Delete webhook



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
	id := "id_example" // string | Webhook ID to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.DeleteWebhookSettings(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.DeleteWebhookSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhookSettings`: UpdateYoutubeDefaultPlaylist200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.DeleteWebhookSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Webhook ID to delete | 

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


## GetWebhookLogs

> GetWebhookLogs200Response GetWebhookLogs(ctx).Limit(limit).Skip(skip).Status(status).Event(event).WebhookId(webhookId).EventId(eventId).Execute()

List webhook delivery logs



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
	limit := int32(56) // int32 | Maximum number of logs to return (optional) (default to 50)
	skip := int32(56) // int32 | Number of logs to skip (offset-based pagination) (optional) (default to 0)
	status := "status_example" // string | Filter by delivery outcome (optional)
	event := "event_example" // string | Filter by event type (e.g. post.published) (optional)
	webhookId := "webhookId_example" // string | Filter by webhook configuration ID (optional)
	eventId := "eventId_example" // string | Filter by stable webhook event ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetWebhookLogs(context.Background()).Limit(limit).Skip(skip).Status(status).Event(event).WebhookId(webhookId).EventId(eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhookLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookLogs`: GetWebhookLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhookLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of logs to return | [default to 50]
 **skip** | **int32** | Number of logs to skip (offset-based pagination) | [default to 0]
 **status** | **string** | Filter by delivery outcome | 
 **event** | **string** | Filter by event type (e.g. post.published) | 
 **webhookId** | **string** | Filter by webhook configuration ID | 
 **eventId** | **string** | Filter by stable webhook event ID | 

### Return type

[**GetWebhookLogs200Response**](GetWebhookLogs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookSettings

> GetWebhookSettings200Response GetWebhookSettings(ctx).Execute()

List webhooks



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
	resp, r, err := apiClient.WebhooksAPI.GetWebhookSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhookSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookSettings`: GetWebhookSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhookSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookSettingsRequest struct via the builder pattern


### Return type

[**GetWebhookSettings200Response**](GetWebhookSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestWebhook

> UnpublishPost200Response TestWebhook(ctx).TestWebhookRequest(testWebhookRequest).Execute()

Send test webhook



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
	testWebhookRequest := *openapiclient.NewTestWebhookRequest("WebhookId_example") // TestWebhookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.TestWebhook(context.Background()).TestWebhookRequest(testWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.TestWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestWebhook`: UnpublishPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.TestWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testWebhookRequest** | [**TestWebhookRequest**](TestWebhookRequest.md) |  | 

### Return type

[**UnpublishPost200Response**](UnpublishPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookSettings

> UpdateWebhookSettings200Response UpdateWebhookSettings(ctx).UpdateWebhookSettingsRequest(updateWebhookSettingsRequest).Execute()

Update webhook



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
	updateWebhookSettingsRequest := *openapiclient.NewUpdateWebhookSettingsRequest("Id_example") // UpdateWebhookSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.UpdateWebhookSettings(context.Background()).UpdateWebhookSettingsRequest(updateWebhookSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.UpdateWebhookSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookSettings`: UpdateWebhookSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.UpdateWebhookSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateWebhookSettingsRequest** | [**UpdateWebhookSettingsRequest**](UpdateWebhookSettingsRequest.md) |  | 

### Return type

[**UpdateWebhookSettings200Response**](UpdateWebhookSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

