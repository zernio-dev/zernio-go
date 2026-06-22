# \InboxAnalyticsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInboxConversationAnalytics**](InboxAnalyticsAPI.md#GetInboxConversationAnalytics) | **Get** /v1/analytics/inbox/conversations/{conversationId} | Get analytics for a single conversation
[**GetInboxHeatmap**](InboxAnalyticsAPI.md#GetInboxHeatmap) | **Get** /v1/analytics/inbox/heatmap | Get inbox day-of-week × hour-of-day heatmap
[**GetInboxResponseTime**](InboxAnalyticsAPI.md#GetInboxResponseTime) | **Get** /v1/analytics/inbox/response-time | Get inbox response-time stats
[**GetInboxSourceBreakdown**](InboxAnalyticsAPI.md#GetInboxSourceBreakdown) | **Get** /v1/analytics/inbox/source-breakdown | Get inbox source breakdown
[**GetInboxTopAccounts**](InboxAnalyticsAPI.md#GetInboxTopAccounts) | **Get** /v1/analytics/inbox/top-accounts | Get top accounts by inbox volume
[**GetInboxVolume**](InboxAnalyticsAPI.md#GetInboxVolume) | **Get** /v1/analytics/inbox/volume | Get inbox messaging volume
[**ListInboxConversationAnalytics**](InboxAnalyticsAPI.md#ListInboxConversationAnalytics) | **Get** /v1/analytics/inbox/conversations | List conversations with inbox analytics



## GetInboxConversationAnalytics

> GetInboxConversationAnalytics200Response GetInboxConversationAnalytics(ctx, conversationId).FromDate(fromDate).ToDate(toDate).Execute()

Get analytics for a single conversation



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
	conversationId := "conversationId_example" // string | Mongo _id or platformConversationId.
	fromDate := time.Now() // string | 
	toDate := time.Now() // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAnalyticsAPI.GetInboxConversationAnalytics(context.Background(), conversationId).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAnalyticsAPI.GetInboxConversationAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxConversationAnalytics`: GetInboxConversationAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `InboxAnalyticsAPI.GetInboxConversationAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | Mongo _id or platformConversationId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxConversationAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 

### Return type

[**GetInboxConversationAnalytics200Response**](GetInboxConversationAnalytics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxHeatmap

> GetInboxHeatmap200Response GetInboxHeatmap(ctx).FromDate(fromDate).ToDate(toDate).ProfileId(profileId).Platform(platform).AccountId(accountId).Source(source).Action(action).Execute()

Get inbox day-of-week × hour-of-day heatmap



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
	fromDate := time.Now() // string | 
	toDate := time.Now() // string |  (optional)
	profileId := "profileId_example" // string |  (optional)
	platform := "platform_example" // string |  (optional)
	accountId := "accountId_example" // string |  (optional)
	source := "source_example" // string |  (optional)
	action := "action_example" // string | Narrow to a single event type. \"all\" or omitted means no filter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAnalyticsAPI.GetInboxHeatmap(context.Background()).FromDate(fromDate).ToDate(toDate).ProfileId(profileId).Platform(platform).AccountId(accountId).Source(source).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAnalyticsAPI.GetInboxHeatmap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxHeatmap`: GetInboxHeatmap200Response
	fmt.Fprintf(os.Stdout, "Response from `InboxAnalyticsAPI.GetInboxHeatmap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxHeatmapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 
 **profileId** | **string** |  | 
 **platform** | **string** |  | 
 **accountId** | **string** |  | 
 **source** | **string** |  | 
 **action** | **string** | Narrow to a single event type. \&quot;all\&quot; or omitted means no filter. | 

### Return type

[**GetInboxHeatmap200Response**](GetInboxHeatmap200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxResponseTime

> GetInboxResponseTime200Response GetInboxResponseTime(ctx).FromDate(fromDate).ToDate(toDate).ProfileId(profileId).Platform(platform).AccountId(accountId).Execute()

Get inbox response-time stats



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
	fromDate := time.Now() // string | 
	toDate := time.Now() // string |  (optional)
	profileId := "profileId_example" // string |  (optional)
	platform := "platform_example" // string |  (optional)
	accountId := "accountId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAnalyticsAPI.GetInboxResponseTime(context.Background()).FromDate(fromDate).ToDate(toDate).ProfileId(profileId).Platform(platform).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAnalyticsAPI.GetInboxResponseTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxResponseTime`: GetInboxResponseTime200Response
	fmt.Fprintf(os.Stdout, "Response from `InboxAnalyticsAPI.GetInboxResponseTime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxResponseTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 
 **profileId** | **string** |  | 
 **platform** | **string** |  | 
 **accountId** | **string** |  | 

### Return type

[**GetInboxResponseTime200Response**](GetInboxResponseTime200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxSourceBreakdown

> GetInboxSourceBreakdown200Response GetInboxSourceBreakdown(ctx).FromDate(fromDate).ToDate(toDate).ProfileId(profileId).Platform(platform).AccountId(accountId).Execute()

Get inbox source breakdown



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
	fromDate := time.Now() // string | 
	toDate := time.Now() // string |  (optional)
	profileId := "profileId_example" // string |  (optional)
	platform := "platform_example" // string |  (optional)
	accountId := "accountId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAnalyticsAPI.GetInboxSourceBreakdown(context.Background()).FromDate(fromDate).ToDate(toDate).ProfileId(profileId).Platform(platform).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAnalyticsAPI.GetInboxSourceBreakdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxSourceBreakdown`: GetInboxSourceBreakdown200Response
	fmt.Fprintf(os.Stdout, "Response from `InboxAnalyticsAPI.GetInboxSourceBreakdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxSourceBreakdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 
 **profileId** | **string** |  | 
 **platform** | **string** |  | 
 **accountId** | **string** |  | 

### Return type

[**GetInboxSourceBreakdown200Response**](GetInboxSourceBreakdown200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxTopAccounts

> GetInboxTopAccounts200Response GetInboxTopAccounts(ctx).FromDate(fromDate).ToDate(toDate).ProfileId(profileId).Platform(platform).Source(source).Limit(limit).Execute()

Get top accounts by inbox volume



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
	fromDate := time.Now() // string | 
	toDate := time.Now() // string |  (optional)
	profileId := "profileId_example" // string |  (optional)
	platform := "platform_example" // string |  (optional)
	source := "source_example" // string |  (optional)
	limit := int32(56) // int32 | Cap on returned rows. Lower than the posting listing's 100 because each row triggers a SocialAccount Mongo lookup. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAnalyticsAPI.GetInboxTopAccounts(context.Background()).FromDate(fromDate).ToDate(toDate).ProfileId(profileId).Platform(platform).Source(source).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAnalyticsAPI.GetInboxTopAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxTopAccounts`: GetInboxTopAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `InboxAnalyticsAPI.GetInboxTopAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxTopAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 
 **profileId** | **string** |  | 
 **platform** | **string** |  | 
 **source** | **string** |  | 
 **limit** | **int32** | Cap on returned rows. Lower than the posting listing&#39;s 100 because each row triggers a SocialAccount Mongo lookup. | [default to 10]

### Return type

[**GetInboxTopAccounts200Response**](GetInboxTopAccounts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxVolume

> GetInboxVolume200Response GetInboxVolume(ctx).FromDate(fromDate).ToDate(toDate).ProfileId(profileId).Platform(platform).AccountId(accountId).Source(source).Execute()

Get inbox messaging volume



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
	fromDate := time.Now() // string | Inclusive lower bound (YYYY-MM-DD). Required.
	toDate := time.Now() // string | Inclusive upper bound (YYYY-MM-DD). Defaults to today. (optional)
	profileId := "profileId_example" // string |  (optional)
	platform := "platform_example" // string | Filter by single platform (facebook, instagram, twitter, etc.). (optional)
	accountId := "accountId_example" // string |  (optional)
	source := "source_example" // string | Filter by metadata.source lineage (human, workflow, sequence, broadcast, comment_automation, api, contact, platform). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAnalyticsAPI.GetInboxVolume(context.Background()).FromDate(fromDate).ToDate(toDate).ProfileId(profileId).Platform(platform).AccountId(accountId).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAnalyticsAPI.GetInboxVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxVolume`: GetInboxVolume200Response
	fmt.Fprintf(os.Stdout, "Response from `InboxAnalyticsAPI.GetInboxVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **string** | Inclusive lower bound (YYYY-MM-DD). Required. | 
 **toDate** | **string** | Inclusive upper bound (YYYY-MM-DD). Defaults to today. | 
 **profileId** | **string** |  | 
 **platform** | **string** | Filter by single platform (facebook, instagram, twitter, etc.). | 
 **accountId** | **string** |  | 
 **source** | **string** | Filter by metadata.source lineage (human, workflow, sequence, broadcast, comment_automation, api, contact, platform). | 

### Return type

[**GetInboxVolume200Response**](GetInboxVolume200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInboxConversationAnalytics

> ListInboxConversationAnalytics200Response ListInboxConversationAnalytics(ctx).FromDate(fromDate).ToDate(toDate).ProfileId(profileId).Platform(platform).AccountId(accountId).Source(source).Limit(limit).Page(page).SortBy(sortBy).Order(order).Execute()

List conversations with inbox analytics



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
	fromDate := time.Now() // string | 
	toDate := time.Now() // string |  (optional)
	profileId := "profileId_example" // string |  (optional)
	platform := "platform_example" // string |  (optional)
	accountId := "accountId_example" // string |  (optional)
	source := "source_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	page := int32(56) // int32 |  (optional) (default to 1)
	sortBy := "sortBy_example" // string |  (optional) (default to "lastMessageAt")
	order := "order_example" // string |  (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAnalyticsAPI.ListInboxConversationAnalytics(context.Background()).FromDate(fromDate).ToDate(toDate).ProfileId(profileId).Platform(platform).AccountId(accountId).Source(source).Limit(limit).Page(page).SortBy(sortBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAnalyticsAPI.ListInboxConversationAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInboxConversationAnalytics`: ListInboxConversationAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `InboxAnalyticsAPI.ListInboxConversationAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInboxConversationAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 
 **profileId** | **string** |  | 
 **platform** | **string** |  | 
 **accountId** | **string** |  | 
 **source** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **page** | **int32** |  | [default to 1]
 **sortBy** | **string** |  | [default to &quot;lastMessageAt&quot;]
 **order** | **string** |  | [default to &quot;desc&quot;]

### Return type

[**ListInboxConversationAnalytics200Response**](ListInboxConversationAnalytics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

