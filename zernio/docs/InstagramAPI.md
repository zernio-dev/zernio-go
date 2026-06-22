# \InstagramAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstagramStoryInsights**](InstagramAPI.md#GetInstagramStoryInsights) | **Get** /v1/accounts/{accountId}/instagram/stories/{storyId}/insights | Get Instagram story insights
[**ListInstagramStories**](InstagramAPI.md#ListInstagramStories) | **Get** /v1/accounts/{accountId}/instagram/stories | List active Instagram stories



## GetInstagramStoryInsights

> GetInstagramStoryInsights200Response GetInstagramStoryInsights(ctx, accountId, storyId).Execute()

Get Instagram story insights



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
	accountId := "accountId_example" // string | The Instagram account ID
	storyId := "storyId_example" // string | The Instagram media ID of the story.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstagramAPI.GetInstagramStoryInsights(context.Background(), accountId, storyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstagramAPI.GetInstagramStoryInsights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstagramStoryInsights`: GetInstagramStoryInsights200Response
	fmt.Fprintf(os.Stdout, "Response from `InstagramAPI.GetInstagramStoryInsights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The Instagram account ID | 
**storyId** | **string** | The Instagram media ID of the story. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstagramStoryInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetInstagramStoryInsights200Response**](GetInstagramStoryInsights200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstagramStories

> ListInstagramStories200Response ListInstagramStories(ctx, accountId).Execute()

List active Instagram stories



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
	accountId := "accountId_example" // string | The Instagram account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstagramAPI.ListInstagramStories(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstagramAPI.ListInstagramStories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstagramStories`: ListInstagramStories200Response
	fmt.Fprintf(os.Stdout, "Response from `InstagramAPI.ListInstagramStories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The Instagram account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstagramStoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListInstagramStories200Response**](ListInstagramStories200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

