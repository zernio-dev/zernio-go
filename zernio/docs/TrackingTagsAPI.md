# \TrackingTagsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTrackingTagSharedAccount**](TrackingTagsAPI.md#AddTrackingTagSharedAccount) | **Post** /v1/accounts/{accountId}/tracking-tags/{tagId}/shared-accounts | Share a tracking tag with an ad account
[**CreateTrackingTag**](TrackingTagsAPI.md#CreateTrackingTag) | **Post** /v1/accounts/{accountId}/tracking-tags | Create a tracking tag (Meta Pixel)
[**GetTrackingTag**](TrackingTagsAPI.md#GetTrackingTag) | **Get** /v1/accounts/{accountId}/tracking-tags/{tagId} | Fetch a single tracking tag (Meta Pixel)
[**GetTrackingTagStats**](TrackingTagsAPI.md#GetTrackingTagStats) | **Get** /v1/accounts/{accountId}/tracking-tags/{tagId}/stats | Aggregated event stats for a tracking tag (Meta Pixel)
[**ListTrackingTagSharedAccounts**](TrackingTagsAPI.md#ListTrackingTagSharedAccounts) | **Get** /v1/accounts/{accountId}/tracking-tags/{tagId}/shared-accounts | List ad accounts a tracking tag is shared with
[**ListTrackingTags**](TrackingTagsAPI.md#ListTrackingTags) | **Get** /v1/accounts/{accountId}/tracking-tags | List tracking tags (Meta Pixels)
[**RemoveTrackingTagSharedAccount**](TrackingTagsAPI.md#RemoveTrackingTagSharedAccount) | **Delete** /v1/accounts/{accountId}/tracking-tags/{tagId}/shared-accounts | Stop sharing a tracking tag with an ad account
[**UpdateTrackingTag**](TrackingTagsAPI.md#UpdateTrackingTag) | **Patch** /v1/accounts/{accountId}/tracking-tags/{tagId} | Update a tracking tag (Meta Pixel)



## AddTrackingTagSharedAccount

> AddTrackingTagSharedAccount201Response AddTrackingTagSharedAccount(ctx, accountId, tagId).AddTrackingTagSharedAccountRequest(addTrackingTagSharedAccountRequest).Execute()

Share a tracking tag with an ad account



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
	accountId := "accountId_example" // string | 
	tagId := "tagId_example" // string | Pixel id.
	addTrackingTagSharedAccountRequest := *openapiclient.NewAddTrackingTagSharedAccountRequest("AdAccountId_example") // AddTrackingTagSharedAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackingTagsAPI.AddTrackingTagSharedAccount(context.Background(), accountId, tagId).AddTrackingTagSharedAccountRequest(addTrackingTagSharedAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackingTagsAPI.AddTrackingTagSharedAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTrackingTagSharedAccount`: AddTrackingTagSharedAccount201Response
	fmt.Fprintf(os.Stdout, "Response from `TrackingTagsAPI.AddTrackingTagSharedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**tagId** | **string** | Pixel id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTrackingTagSharedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addTrackingTagSharedAccountRequest** | [**AddTrackingTagSharedAccountRequest**](AddTrackingTagSharedAccountRequest.md) |  | 

### Return type

[**AddTrackingTagSharedAccount201Response**](AddTrackingTagSharedAccount201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrackingTag

> CreateTrackingTag201Response CreateTrackingTag(ctx, accountId).CreateTrackingTagRequest(createTrackingTagRequest).Execute()

Create a tracking tag (Meta Pixel)



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
	accountId := "accountId_example" // string | Meta ads SocialAccount id (platform `metaads`).
	createTrackingTagRequest := *openapiclient.NewCreateTrackingTagRequest("AdAccountId_example", "Name_example") // CreateTrackingTagRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackingTagsAPI.CreateTrackingTag(context.Background(), accountId).CreateTrackingTagRequest(createTrackingTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackingTagsAPI.CreateTrackingTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTrackingTag`: CreateTrackingTag201Response
	fmt.Fprintf(os.Stdout, "Response from `TrackingTagsAPI.CreateTrackingTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Meta ads SocialAccount id (platform &#x60;metaads&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrackingTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTrackingTagRequest** | [**CreateTrackingTagRequest**](CreateTrackingTagRequest.md) |  | 

### Return type

[**CreateTrackingTag201Response**](CreateTrackingTag201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackingTag

> CreateTrackingTag201Response GetTrackingTag(ctx, accountId, tagId).Execute()

Fetch a single tracking tag (Meta Pixel)



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
	accountId := "accountId_example" // string | 
	tagId := "tagId_example" // string | Pixel id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackingTagsAPI.GetTrackingTag(context.Background(), accountId, tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackingTagsAPI.GetTrackingTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrackingTag`: CreateTrackingTag201Response
	fmt.Fprintf(os.Stdout, "Response from `TrackingTagsAPI.GetTrackingTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**tagId** | **string** | Pixel id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackingTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateTrackingTag201Response**](CreateTrackingTag201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackingTagStats

> GetTrackingTagStats200Response GetTrackingTagStats(ctx, accountId, tagId).Aggregation(aggregation).StartTime(startTime).EndTime(endTime).Execute()

Aggregated event stats for a tracking tag (Meta Pixel)



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
	accountId := "accountId_example" // string | 
	tagId := "tagId_example" // string | Pixel id.
	aggregation := "aggregation_example" // string | Aggregation dimension. Defaults to `event`. (optional) (default to "event")
	startTime := int32(56) // int32 | Unix seconds lower bound. (optional)
	endTime := int32(56) // int32 | Unix seconds upper bound. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackingTagsAPI.GetTrackingTagStats(context.Background(), accountId, tagId).Aggregation(aggregation).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackingTagsAPI.GetTrackingTagStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrackingTagStats`: GetTrackingTagStats200Response
	fmt.Fprintf(os.Stdout, "Response from `TrackingTagsAPI.GetTrackingTagStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**tagId** | **string** | Pixel id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackingTagStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aggregation** | **string** | Aggregation dimension. Defaults to &#x60;event&#x60;. | [default to &quot;event&quot;]
 **startTime** | **int32** | Unix seconds lower bound. | 
 **endTime** | **int32** | Unix seconds upper bound. | 

### Return type

[**GetTrackingTagStats200Response**](GetTrackingTagStats200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrackingTagSharedAccounts

> ListTrackingTagSharedAccounts200Response ListTrackingTagSharedAccounts(ctx, accountId, tagId).Execute()

List ad accounts a tracking tag is shared with



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
	accountId := "accountId_example" // string | 
	tagId := "tagId_example" // string | Pixel id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackingTagsAPI.ListTrackingTagSharedAccounts(context.Background(), accountId, tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackingTagsAPI.ListTrackingTagSharedAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrackingTagSharedAccounts`: ListTrackingTagSharedAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `TrackingTagsAPI.ListTrackingTagSharedAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**tagId** | **string** | Pixel id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTrackingTagSharedAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListTrackingTagSharedAccounts200Response**](ListTrackingTagSharedAccounts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrackingTags

> ListTrackingTags200Response ListTrackingTags(ctx, accountId).AdAccountId(adAccountId).Execute()

List tracking tags (Meta Pixels)



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
	accountId := "accountId_example" // string | Meta ads SocialAccount id (platform `metaads`).
	adAccountId := "adAccountId_example" // string | Optional. Scope to one ad account, e.g. `act_123456789`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackingTagsAPI.ListTrackingTags(context.Background(), accountId).AdAccountId(adAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackingTagsAPI.ListTrackingTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrackingTags`: ListTrackingTags200Response
	fmt.Fprintf(os.Stdout, "Response from `TrackingTagsAPI.ListTrackingTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Meta ads SocialAccount id (platform &#x60;metaads&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTrackingTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adAccountId** | **string** | Optional. Scope to one ad account, e.g. &#x60;act_123456789&#x60;. | 

### Return type

[**ListTrackingTags200Response**](ListTrackingTags200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTrackingTagSharedAccount

> RemoveTrackingTagSharedAccount(ctx, accountId, tagId).AdAccountId(adAccountId).Execute()

Stop sharing a tracking tag with an ad account



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
	accountId := "accountId_example" // string | 
	tagId := "tagId_example" // string | Pixel id.
	adAccountId := "adAccountId_example" // string | Ad account to unshare, e.g. `act_123456789`. May also be sent in the JSON body. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TrackingTagsAPI.RemoveTrackingTagSharedAccount(context.Background(), accountId, tagId).AdAccountId(adAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackingTagsAPI.RemoveTrackingTagSharedAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**tagId** | **string** | Pixel id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTrackingTagSharedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **adAccountId** | **string** | Ad account to unshare, e.g. &#x60;act_123456789&#x60;. May also be sent in the JSON body. | 

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


## UpdateTrackingTag

> CreateTrackingTag201Response UpdateTrackingTag(ctx, accountId, tagId).UpdateTrackingTagRequest(updateTrackingTagRequest).Execute()

Update a tracking tag (Meta Pixel)



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
	accountId := "accountId_example" // string | 
	tagId := "tagId_example" // string | Pixel id.
	updateTrackingTagRequest := *openapiclient.NewUpdateTrackingTagRequest() // UpdateTrackingTagRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackingTagsAPI.UpdateTrackingTag(context.Background(), accountId, tagId).UpdateTrackingTagRequest(updateTrackingTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackingTagsAPI.UpdateTrackingTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTrackingTag`: CreateTrackingTag201Response
	fmt.Fprintf(os.Stdout, "Response from `TrackingTagsAPI.UpdateTrackingTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**tagId** | **string** | Pixel id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrackingTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateTrackingTagRequest** | [**UpdateTrackingTagRequest**](UpdateTrackingTagRequest.md) |  | 

### Return type

[**CreateTrackingTag201Response**](CreateTrackingTag201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

