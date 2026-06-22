# \AdCampaignsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateAdCampaignStatus**](AdCampaignsAPI.md#BulkUpdateAdCampaignStatus) | **Post** /v1/ads/campaigns/bulk-status | Pause or resume many campaigns
[**DeleteAdCampaign**](AdCampaignsAPI.md#DeleteAdCampaign) | **Delete** /v1/ads/campaigns/{campaignId} | Delete a campaign
[**DuplicateAdCampaign**](AdCampaignsAPI.md#DuplicateAdCampaign) | **Post** /v1/ads/campaigns/{campaignId}/duplicate | Duplicate a campaign
[**GetAdTree**](AdCampaignsAPI.md#GetAdTree) | **Get** /v1/ads/tree | Get campaign tree
[**GetAdsTimeline**](AdCampaignsAPI.md#GetAdsTimeline) | **Get** /v1/ads/timeline | Get daily aggregate ad metrics for an account
[**ListAdCampaigns**](AdCampaignsAPI.md#ListAdCampaigns) | **Get** /v1/ads/campaigns | List campaigns
[**UpdateAdCampaign**](AdCampaignsAPI.md#UpdateAdCampaign) | **Put** /v1/ads/campaigns/{campaignId} | Update a campaign (budget and/or bid strategy)
[**UpdateAdCampaignStatus**](AdCampaignsAPI.md#UpdateAdCampaignStatus) | **Put** /v1/ads/campaigns/{campaignId}/status | Pause or resume a campaign
[**UpdateAdSet**](AdCampaignsAPI.md#UpdateAdSet) | **Put** /v1/ads/ad-sets/{adSetId} | Update an ad set (budget, status, and/or bid strategy)
[**UpdateAdSetStatus**](AdCampaignsAPI.md#UpdateAdSetStatus) | **Put** /v1/ads/ad-sets/{adSetId}/status | Pause or resume a single ad set



## BulkUpdateAdCampaignStatus

> BulkUpdateAdCampaignStatus200Response BulkUpdateAdCampaignStatus(ctx).BulkUpdateAdCampaignStatusRequest(bulkUpdateAdCampaignStatusRequest).Execute()

Pause or resume many campaigns



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
	bulkUpdateAdCampaignStatusRequest := *openapiclient.NewBulkUpdateAdCampaignStatusRequest("Status_example", []openapiclient.BulkUpdateAdCampaignStatusRequestCampaignsInner{*openapiclient.NewBulkUpdateAdCampaignStatusRequestCampaignsInner("PlatformCampaignId_example", "Platform_example")}) // BulkUpdateAdCampaignStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdCampaignsAPI.BulkUpdateAdCampaignStatus(context.Background()).BulkUpdateAdCampaignStatusRequest(bulkUpdateAdCampaignStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdCampaignsAPI.BulkUpdateAdCampaignStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUpdateAdCampaignStatus`: BulkUpdateAdCampaignStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `AdCampaignsAPI.BulkUpdateAdCampaignStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateAdCampaignStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkUpdateAdCampaignStatusRequest** | [**BulkUpdateAdCampaignStatusRequest**](BulkUpdateAdCampaignStatusRequest.md) |  | 

### Return type

[**BulkUpdateAdCampaignStatus200Response**](BulkUpdateAdCampaignStatus200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAdCampaign

> DeleteAdCampaign200Response DeleteAdCampaign(ctx, campaignId).DeleteAdCampaignRequest(deleteAdCampaignRequest).Execute()

Delete a campaign



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
	campaignId := "campaignId_example" // string | Platform campaign ID
	deleteAdCampaignRequest := *openapiclient.NewDeleteAdCampaignRequest("Platform_example") // DeleteAdCampaignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdCampaignsAPI.DeleteAdCampaign(context.Background(), campaignId).DeleteAdCampaignRequest(deleteAdCampaignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdCampaignsAPI.DeleteAdCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAdCampaign`: DeleteAdCampaign200Response
	fmt.Fprintf(os.Stdout, "Response from `AdCampaignsAPI.DeleteAdCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Platform campaign ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteAdCampaignRequest** | [**DeleteAdCampaignRequest**](DeleteAdCampaignRequest.md) |  | 

### Return type

[**DeleteAdCampaign200Response**](DeleteAdCampaign200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuplicateAdCampaign

> DuplicateAdCampaign200Response DuplicateAdCampaign(ctx, campaignId).DuplicateAdCampaignRequest(duplicateAdCampaignRequest).Execute()

Duplicate a campaign



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
	campaignId := "campaignId_example" // string | Source platform campaign ID
	duplicateAdCampaignRequest := *openapiclient.NewDuplicateAdCampaignRequest("Platform_example") // DuplicateAdCampaignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdCampaignsAPI.DuplicateAdCampaign(context.Background(), campaignId).DuplicateAdCampaignRequest(duplicateAdCampaignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdCampaignsAPI.DuplicateAdCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DuplicateAdCampaign`: DuplicateAdCampaign200Response
	fmt.Fprintf(os.Stdout, "Response from `AdCampaignsAPI.DuplicateAdCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Source platform campaign ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuplicateAdCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duplicateAdCampaignRequest** | [**DuplicateAdCampaignRequest**](DuplicateAdCampaignRequest.md) |  | 

### Return type

[**DuplicateAdCampaign200Response**](DuplicateAdCampaign200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdTree

> GetAdTree200Response GetAdTree(ctx).Page(page).Limit(limit).Source(source).Platform(platform).Status(status).AdAccountId(adAccountId).AccountId(accountId).ProfileId(profileId).CampaignId(campaignId).FromDate(fromDate).ToDate(toDate).Sort(sort).Execute()

Get campaign tree



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
	page := int32(56) // int32 | Page number (1-based) (optional) (default to 1)
	limit := int32(56) // int32 | Campaigns per page (optional) (default to 20)
	source := "source_example" // string | `all` (default) returns both Zernio-created ads and those discovered from the platform's ad manager — matches the web UI's default view. Pass `zernio` to restrict to isExternal=false only. Status is NOT filtered by default — use the `status` param for that. (optional) (default to "all")
	platform := "platform_example" // string |  (optional)
	status := openapiclient.AdStatus("active") // AdStatus | Filter by derived campaign status (post-aggregation) (optional)
	adAccountId := "adAccountId_example" // string | Platform ad account ID (optional)
	accountId := "accountId_example" // string | Social account ID (optional)
	profileId := "profileId_example" // string | Profile ID (optional)
	campaignId := "campaignId_example" // string | Restrict the tree to a single campaign by its platform campaign id (the id the platform assigns, e.g. Meta's numeric campaign id). Filters the campaign set itself, so it works regardless of account size and pagination — pass this when you already hold a campaign id instead of paging the tree to find it. Mirrors the `campaignId` filter on GET /v1/ads. (optional)
	fromDate := time.Now() // string | Start of the METRICS date range (YYYY-MM-DD). Affects only the spend/impression numbers overlaid on each node, NOT which campaigns are returned. Defaults to 90 days ago. (optional)
	toDate := time.Now() // string | End of metrics date range (YYYY-MM-DD). Defaults to today. Max 730-day range. (optional)
	sort := "sort_example" // string | Campaign-level sort order. `newest` (default) / `oldest` order by the campaign's newest-ad createdAt. `spend_desc` / `spend_asc` order by aggregated spend in the requested date range; campaigns with no spend land at the end. (optional) (default to "newest")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdCampaignsAPI.GetAdTree(context.Background()).Page(page).Limit(limit).Source(source).Platform(platform).Status(status).AdAccountId(adAccountId).AccountId(accountId).ProfileId(profileId).CampaignId(campaignId).FromDate(fromDate).ToDate(toDate).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdCampaignsAPI.GetAdTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdTree`: GetAdTree200Response
	fmt.Fprintf(os.Stdout, "Response from `AdCampaignsAPI.GetAdTree`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-based) | [default to 1]
 **limit** | **int32** | Campaigns per page | [default to 20]
 **source** | **string** | &#x60;all&#x60; (default) returns both Zernio-created ads and those discovered from the platform&#39;s ad manager — matches the web UI&#39;s default view. Pass &#x60;zernio&#x60; to restrict to isExternal&#x3D;false only. Status is NOT filtered by default — use the &#x60;status&#x60; param for that. | [default to &quot;all&quot;]
 **platform** | **string** |  | 
 **status** | [**AdStatus**](AdStatus.md) | Filter by derived campaign status (post-aggregation) | 
 **adAccountId** | **string** | Platform ad account ID | 
 **accountId** | **string** | Social account ID | 
 **profileId** | **string** | Profile ID | 
 **campaignId** | **string** | Restrict the tree to a single campaign by its platform campaign id (the id the platform assigns, e.g. Meta&#39;s numeric campaign id). Filters the campaign set itself, so it works regardless of account size and pagination — pass this when you already hold a campaign id instead of paging the tree to find it. Mirrors the &#x60;campaignId&#x60; filter on GET /v1/ads. | 
 **fromDate** | **string** | Start of the METRICS date range (YYYY-MM-DD). Affects only the spend/impression numbers overlaid on each node, NOT which campaigns are returned. Defaults to 90 days ago. | 
 **toDate** | **string** | End of metrics date range (YYYY-MM-DD). Defaults to today. Max 730-day range. | 
 **sort** | **string** | Campaign-level sort order. &#x60;newest&#x60; (default) / &#x60;oldest&#x60; order by the campaign&#39;s newest-ad createdAt. &#x60;spend_desc&#x60; / &#x60;spend_asc&#x60; order by aggregated spend in the requested date range; campaigns with no spend land at the end. | [default to &quot;newest&quot;]

### Return type

[**GetAdTree200Response**](GetAdTree200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdsTimeline

> GetAdsTimeline200Response GetAdsTimeline(ctx).AccountId(accountId).AdAccountId(adAccountId).FromDate(fromDate).ToDate(toDate).Platform(platform).Execute()

Get daily aggregate ad metrics for an account



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
	accountId := "accountId_example" // string | Social account ID. Sibling-expanded to its linked posting↔ads pair.
	adAccountId := "adAccountId_example" // string | Optional platform-native ad account ID (e.g. Meta `act_…`, TikTok advertiser ID). Use when the connection wraps multiple platform ad accounts and the chart should show one only. Note: rows ingested before 2026-05-13 don't carry this column; the recurring 7-day re-sync repopulates them naturally. (optional)
	fromDate := time.Now() // string | Inclusive start of metrics range (YYYY-MM-DD). Defaults to 90 days ago. (optional)
	toDate := time.Now() // string | Inclusive end of metrics range (YYYY-MM-DD). Defaults to today. Max 730-day range. (optional)
	platform := "platform_example" // string | Restrict to one platform. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdCampaignsAPI.GetAdsTimeline(context.Background()).AccountId(accountId).AdAccountId(adAccountId).FromDate(fromDate).ToDate(toDate).Platform(platform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdCampaignsAPI.GetAdsTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdsTimeline`: GetAdsTimeline200Response
	fmt.Fprintf(os.Stdout, "Response from `AdCampaignsAPI.GetAdsTimeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdsTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Social account ID. Sibling-expanded to its linked posting↔ads pair. | 
 **adAccountId** | **string** | Optional platform-native ad account ID (e.g. Meta &#x60;act_…&#x60;, TikTok advertiser ID). Use when the connection wraps multiple platform ad accounts and the chart should show one only. Note: rows ingested before 2026-05-13 don&#39;t carry this column; the recurring 7-day re-sync repopulates them naturally. | 
 **fromDate** | **string** | Inclusive start of metrics range (YYYY-MM-DD). Defaults to 90 days ago. | 
 **toDate** | **string** | Inclusive end of metrics range (YYYY-MM-DD). Defaults to today. Max 730-day range. | 
 **platform** | **string** | Restrict to one platform. | 

### Return type

[**GetAdsTimeline200Response**](GetAdsTimeline200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAdCampaigns

> ListAdCampaigns200Response ListAdCampaigns(ctx).Page(page).Limit(limit).Source(source).Platform(platform).Status(status).AdAccountId(adAccountId).AccountId(accountId).ProfileId(profileId).FromDate(fromDate).ToDate(toDate).Execute()

List campaigns



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
	page := int32(56) // int32 | Page number (1-based) (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 20)
	source := "source_example" // string | `all` (default) returns both Zernio-created ads and those discovered from the platform's ad manager — matches the web UI's default view. Pass `zernio` to restrict to isExternal=false only. Status is NOT filtered by default — use the `status` param for that. (optional) (default to "all")
	platform := "platform_example" // string |  (optional)
	status := openapiclient.AdStatus("active") // AdStatus | Filter by derived campaign status (post-aggregation) (optional)
	adAccountId := "adAccountId_example" // string | Platform ad account ID (e.g. act_123 for Meta) (optional)
	accountId := "accountId_example" // string | Social account ID (optional)
	profileId := "profileId_example" // string | Profile ID (optional)
	fromDate := time.Now() // string | Start of metrics date range (YYYY-MM-DD, inclusive). Defaults to 90 days ago when both date params are omitted. (optional)
	toDate := time.Now() // string | End of metrics date range (YYYY-MM-DD, inclusive). Defaults to today. Max 730-day range. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdCampaignsAPI.ListAdCampaigns(context.Background()).Page(page).Limit(limit).Source(source).Platform(platform).Status(status).AdAccountId(adAccountId).AccountId(accountId).ProfileId(profileId).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdCampaignsAPI.ListAdCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAdCampaigns`: ListAdCampaigns200Response
	fmt.Fprintf(os.Stdout, "Response from `AdCampaignsAPI.ListAdCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAdCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-based) | [default to 1]
 **limit** | **int32** |  | [default to 20]
 **source** | **string** | &#x60;all&#x60; (default) returns both Zernio-created ads and those discovered from the platform&#39;s ad manager — matches the web UI&#39;s default view. Pass &#x60;zernio&#x60; to restrict to isExternal&#x3D;false only. Status is NOT filtered by default — use the &#x60;status&#x60; param for that. | [default to &quot;all&quot;]
 **platform** | **string** |  | 
 **status** | [**AdStatus**](AdStatus.md) | Filter by derived campaign status (post-aggregation) | 
 **adAccountId** | **string** | Platform ad account ID (e.g. act_123 for Meta) | 
 **accountId** | **string** | Social account ID | 
 **profileId** | **string** | Profile ID | 
 **fromDate** | **string** | Start of metrics date range (YYYY-MM-DD, inclusive). Defaults to 90 days ago when both date params are omitted. | 
 **toDate** | **string** | End of metrics date range (YYYY-MM-DD, inclusive). Defaults to today. Max 730-day range. | 

### Return type

[**ListAdCampaigns200Response**](ListAdCampaigns200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdCampaign

> UpdateAdCampaign200Response UpdateAdCampaign(ctx, campaignId).UpdateAdCampaignRequest(updateAdCampaignRequest).Execute()

Update a campaign (budget and/or bid strategy)



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
	campaignId := "campaignId_example" // string | Platform campaign ID
	updateAdCampaignRequest := *openapiclient.NewUpdateAdCampaignRequest("Platform_example") // UpdateAdCampaignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdCampaignsAPI.UpdateAdCampaign(context.Background(), campaignId).UpdateAdCampaignRequest(updateAdCampaignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdCampaignsAPI.UpdateAdCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdCampaign`: UpdateAdCampaign200Response
	fmt.Fprintf(os.Stdout, "Response from `AdCampaignsAPI.UpdateAdCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Platform campaign ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAdCampaignRequest** | [**UpdateAdCampaignRequest**](UpdateAdCampaignRequest.md) |  | 

### Return type

[**UpdateAdCampaign200Response**](UpdateAdCampaign200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdCampaignStatus

> UpdateAdCampaignStatus200Response UpdateAdCampaignStatus(ctx, campaignId).UpdateAdCampaignStatusRequest(updateAdCampaignStatusRequest).Execute()

Pause or resume a campaign



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
	campaignId := "campaignId_example" // string | Platform campaign ID
	updateAdCampaignStatusRequest := *openapiclient.NewUpdateAdCampaignStatusRequest("Status_example", "Platform_example") // UpdateAdCampaignStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdCampaignsAPI.UpdateAdCampaignStatus(context.Background(), campaignId).UpdateAdCampaignStatusRequest(updateAdCampaignStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdCampaignsAPI.UpdateAdCampaignStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdCampaignStatus`: UpdateAdCampaignStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `AdCampaignsAPI.UpdateAdCampaignStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Platform campaign ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdCampaignStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAdCampaignStatusRequest** | [**UpdateAdCampaignStatusRequest**](UpdateAdCampaignStatusRequest.md) |  | 

### Return type

[**UpdateAdCampaignStatus200Response**](UpdateAdCampaignStatus200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdSet

> UpdateAdSet200Response UpdateAdSet(ctx, adSetId).UpdateAdSetRequest(updateAdSetRequest).Execute()

Update an ad set (budget, status, and/or bid strategy)



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
	adSetId := "adSetId_example" // string | Platform ad set ID
	updateAdSetRequest := *openapiclient.NewUpdateAdSetRequest("Platform_example") // UpdateAdSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdCampaignsAPI.UpdateAdSet(context.Background(), adSetId).UpdateAdSetRequest(updateAdSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdCampaignsAPI.UpdateAdSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdSet`: UpdateAdSet200Response
	fmt.Fprintf(os.Stdout, "Response from `AdCampaignsAPI.UpdateAdSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adSetId** | **string** | Platform ad set ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAdSetRequest** | [**UpdateAdSetRequest**](UpdateAdSetRequest.md) |  | 

### Return type

[**UpdateAdSet200Response**](UpdateAdSet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdSetStatus

> UpdateAdSetStatus200Response UpdateAdSetStatus(ctx, adSetId).UpdateAdCampaignStatusRequest(updateAdCampaignStatusRequest).Execute()

Pause or resume a single ad set



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
	adSetId := "adSetId_example" // string | Platform ad set ID
	updateAdCampaignStatusRequest := *openapiclient.NewUpdateAdCampaignStatusRequest("Status_example", "Platform_example") // UpdateAdCampaignStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdCampaignsAPI.UpdateAdSetStatus(context.Background(), adSetId).UpdateAdCampaignStatusRequest(updateAdCampaignStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdCampaignsAPI.UpdateAdSetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdSetStatus`: UpdateAdSetStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `AdCampaignsAPI.UpdateAdSetStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adSetId** | **string** | Platform ad set ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdSetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAdCampaignStatusRequest** | [**UpdateAdCampaignStatusRequest**](UpdateAdCampaignStatusRequest.md) |  | 

### Return type

[**UpdateAdSetStatus200Response**](UpdateAdSetStatus200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

