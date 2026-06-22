# \AnalyticsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnalytics**](AnalyticsAPI.md#GetAnalytics) | **Get** /v1/analytics | Get post analytics
[**GetBestTimeToPost**](AnalyticsAPI.md#GetBestTimeToPost) | **Get** /v1/analytics/best-time | Get best times to post
[**GetContentDecay**](AnalyticsAPI.md#GetContentDecay) | **Get** /v1/analytics/content-decay | Get content performance decay
[**GetDailyMetrics**](AnalyticsAPI.md#GetDailyMetrics) | **Get** /v1/analytics/daily-metrics | Get daily aggregated metrics
[**GetFacebookPageInsights**](AnalyticsAPI.md#GetFacebookPageInsights) | **Get** /v1/analytics/facebook/page-insights | Get Facebook Page insights
[**GetFollowerStats**](AnalyticsAPI.md#GetFollowerStats) | **Get** /v1/accounts/follower-stats | Get follower stats
[**GetGoogleBusinessPerformance**](AnalyticsAPI.md#GetGoogleBusinessPerformance) | **Get** /v1/analytics/googlebusiness/performance | Get GBP performance metrics
[**GetGoogleBusinessSearchKeywords**](AnalyticsAPI.md#GetGoogleBusinessSearchKeywords) | **Get** /v1/analytics/googlebusiness/search-keywords | Get GBP search keywords
[**GetInstagramAccountInsights**](AnalyticsAPI.md#GetInstagramAccountInsights) | **Get** /v1/analytics/instagram/account-insights | Get Instagram insights
[**GetInstagramDemographics**](AnalyticsAPI.md#GetInstagramDemographics) | **Get** /v1/analytics/instagram/demographics | Get Instagram demographics
[**GetInstagramFollowerHistory**](AnalyticsAPI.md#GetInstagramFollowerHistory) | **Get** /v1/analytics/instagram/follower-history | Get Instagram follower history
[**GetLinkedInAggregateAnalytics**](AnalyticsAPI.md#GetLinkedInAggregateAnalytics) | **Get** /v1/accounts/{accountId}/linkedin-aggregate-analytics | Get LinkedIn aggregate stats
[**GetLinkedInOrgAggregateAnalytics**](AnalyticsAPI.md#GetLinkedInOrgAggregateAnalytics) | **Get** /v1/analytics/linkedin/org-aggregate-analytics | Get LinkedIn organization page aggregate analytics
[**GetLinkedInPostAnalytics**](AnalyticsAPI.md#GetLinkedInPostAnalytics) | **Get** /v1/accounts/{accountId}/linkedin-post-analytics | Get LinkedIn post stats
[**GetLinkedInPostReactions**](AnalyticsAPI.md#GetLinkedInPostReactions) | **Get** /v1/accounts/{accountId}/linkedin-post-reactions | Get LinkedIn post reactions
[**GetPostTimeline**](AnalyticsAPI.md#GetPostTimeline) | **Get** /v1/analytics/post-timeline | Get post analytics timeline
[**GetPostingFrequency**](AnalyticsAPI.md#GetPostingFrequency) | **Get** /v1/analytics/posting-frequency | Get frequency vs engagement
[**GetTikTokAccountInsights**](AnalyticsAPI.md#GetTikTokAccountInsights) | **Get** /v1/analytics/tiktok/account-insights | Get TikTok account-level insights
[**GetYouTubeChannelInsights**](AnalyticsAPI.md#GetYouTubeChannelInsights) | **Get** /v1/analytics/youtube/channel-insights | Get YouTube channel-level insights
[**GetYouTubeDailyViews**](AnalyticsAPI.md#GetYouTubeDailyViews) | **Get** /v1/analytics/youtube/daily-views | Get YouTube daily views
[**GetYouTubeDemographics**](AnalyticsAPI.md#GetYouTubeDemographics) | **Get** /v1/analytics/youtube/demographics | Get YouTube demographics
[**GetYouTubeVideoRetention**](AnalyticsAPI.md#GetYouTubeVideoRetention) | **Get** /v1/analytics/youtube/video-retention | Get YouTube video retention curve



## GetAnalytics

> GetAnalytics200Response GetAnalytics(ctx).PostId(postId).Platform(platform).ProfileId(profileId).AccountId(accountId).Source(source).FromDate(fromDate).ToDate(toDate).Limit(limit).Page(page).SortBy(sortBy).Order(order).Execute()

Get post analytics



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
	postId := "postId_example" // string | Returns analytics for a single post. Accepts both Zernio Post IDs and External Post IDs. Zernio IDs are auto-resolved to External Post analytics. (optional)
	platform := "platform_example" // string | Filter by platform (default \"all\") (optional)
	profileId := "profileId_example" // string | Filter by profile ID (default \"all\") (optional)
	accountId := "accountId_example" // string | Filter by social account ID (optional)
	source := "source_example" // string | Filter by post source: late (posted via Zernio API), external (synced from platform), all (default) (optional) (default to "all")
	fromDate := time.Now() // string | Inclusive lower bound (YYYY-MM-DD). Defaults to 90 days ago if omitted. Max range is 366 days. (optional)
	toDate := time.Now() // string | Inclusive upper bound (YYYY-MM-DD). Defaults to today if omitted. (optional)
	limit := int32(56) // int32 | Page size (default 50) (optional) (default to 50)
	page := int32(56) // int32 | Page number (default 1) (optional) (default to 1)
	sortBy := "sortBy_example" // string | Sort by date, engagement, or a specific metric (optional) (default to "date")
	order := "order_example" // string | Sort order (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetAnalytics(context.Background()).PostId(postId).Platform(platform).ProfileId(profileId).AccountId(accountId).Source(source).FromDate(fromDate).ToDate(toDate).Limit(limit).Page(page).SortBy(sortBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalytics`: GetAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postId** | **string** | Returns analytics for a single post. Accepts both Zernio Post IDs and External Post IDs. Zernio IDs are auto-resolved to External Post analytics. | 
 **platform** | **string** | Filter by platform (default \&quot;all\&quot;) | 
 **profileId** | **string** | Filter by profile ID (default \&quot;all\&quot;) | 
 **accountId** | **string** | Filter by social account ID | 
 **source** | **string** | Filter by post source: late (posted via Zernio API), external (synced from platform), all (default) | [default to &quot;all&quot;]
 **fromDate** | **string** | Inclusive lower bound (YYYY-MM-DD). Defaults to 90 days ago if omitted. Max range is 366 days. | 
 **toDate** | **string** | Inclusive upper bound (YYYY-MM-DD). Defaults to today if omitted. | 
 **limit** | **int32** | Page size (default 50) | [default to 50]
 **page** | **int32** | Page number (default 1) | [default to 1]
 **sortBy** | **string** | Sort by date, engagement, or a specific metric | [default to &quot;date&quot;]
 **order** | **string** | Sort order | [default to &quot;desc&quot;]

### Return type

[**GetAnalytics200Response**](GetAnalytics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBestTimeToPost

> GetBestTimeToPost200Response GetBestTimeToPost(ctx).Platform(platform).ProfileId(profileId).AccountId(accountId).Source(source).Execute()

Get best times to post



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
	platform := "platform_example" // string | Filter by platform (e.g. \"instagram\", \"tiktok\"). Omit for all platforms. (optional)
	profileId := "profileId_example" // string | Filter by profile ID. Omit for all profiles. (optional)
	accountId := "accountId_example" // string | Filter by social account ID. Omit for all accounts. (optional)
	source := "source_example" // string | Filter by post origin. \"late\" for posts published via Zernio, \"external\" for posts imported from platforms. (optional) (default to "all")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetBestTimeToPost(context.Background()).Platform(platform).ProfileId(profileId).AccountId(accountId).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetBestTimeToPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBestTimeToPost`: GetBestTimeToPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetBestTimeToPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBestTimeToPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **string** | Filter by platform (e.g. \&quot;instagram\&quot;, \&quot;tiktok\&quot;). Omit for all platforms. | 
 **profileId** | **string** | Filter by profile ID. Omit for all profiles. | 
 **accountId** | **string** | Filter by social account ID. Omit for all accounts. | 
 **source** | **string** | Filter by post origin. \&quot;late\&quot; for posts published via Zernio, \&quot;external\&quot; for posts imported from platforms. | [default to &quot;all&quot;]

### Return type

[**GetBestTimeToPost200Response**](GetBestTimeToPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentDecay

> GetContentDecay200Response GetContentDecay(ctx).Platform(platform).ProfileId(profileId).AccountId(accountId).Source(source).Execute()

Get content performance decay



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
	platform := "platform_example" // string | Filter by platform (e.g. \"instagram\", \"tiktok\"). Omit for all platforms. (optional)
	profileId := "profileId_example" // string | Filter by profile ID. Omit for all profiles. (optional)
	accountId := "accountId_example" // string | Filter by social account ID. Omit for all accounts. (optional)
	source := "source_example" // string | Filter by post origin. \"late\" for posts published via Zernio, \"external\" for posts imported from platforms. (optional) (default to "all")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetContentDecay(context.Background()).Platform(platform).ProfileId(profileId).AccountId(accountId).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetContentDecay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentDecay`: GetContentDecay200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetContentDecay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContentDecayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **string** | Filter by platform (e.g. \&quot;instagram\&quot;, \&quot;tiktok\&quot;). Omit for all platforms. | 
 **profileId** | **string** | Filter by profile ID. Omit for all profiles. | 
 **accountId** | **string** | Filter by social account ID. Omit for all accounts. | 
 **source** | **string** | Filter by post origin. \&quot;late\&quot; for posts published via Zernio, \&quot;external\&quot; for posts imported from platforms. | [default to &quot;all&quot;]

### Return type

[**GetContentDecay200Response**](GetContentDecay200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDailyMetrics

> GetDailyMetrics200Response GetDailyMetrics(ctx).Platform(platform).ProfileId(profileId).AccountId(accountId).FromDate(fromDate).ToDate(toDate).Source(source).Attribution(attribution).Execute()

Get daily aggregated metrics



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
	platform := "platform_example" // string | Filter by platform (e.g. \"instagram\", \"tiktok\"). Omit for all platforms. (optional)
	profileId := "profileId_example" // string | Filter by profile ID. Omit for all profiles. (optional)
	accountId := "accountId_example" // string | Filter by social account ID (optional)
	fromDate := time.Now() // time.Time | Inclusive start date (ISO 8601). Defaults to 180 days ago. (optional)
	toDate := time.Now() // time.Time | Inclusive end date (ISO 8601). Defaults to now. (optional)
	source := "source_example" // string | Filter by post origin. \"late\" for posts published via Zernio, \"external\" for posts imported from platforms. (optional) (default to "all")
	attribution := "attribution_example" // string | How each post's engagement is attributed to a day. \"publish\" (default) sums each post's lifetime total on its publish date. \"received\" buckets the per-day increase in engagement by the day it actually arrived (engagement-over-time), so engagement on older posts appears on the day it was gained rather than the post's publish date.  (optional) (default to "publish")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetDailyMetrics(context.Background()).Platform(platform).ProfileId(profileId).AccountId(accountId).FromDate(fromDate).ToDate(toDate).Source(source).Attribution(attribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetDailyMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDailyMetrics`: GetDailyMetrics200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetDailyMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **string** | Filter by platform (e.g. \&quot;instagram\&quot;, \&quot;tiktok\&quot;). Omit for all platforms. | 
 **profileId** | **string** | Filter by profile ID. Omit for all profiles. | 
 **accountId** | **string** | Filter by social account ID | 
 **fromDate** | **time.Time** | Inclusive start date (ISO 8601). Defaults to 180 days ago. | 
 **toDate** | **time.Time** | Inclusive end date (ISO 8601). Defaults to now. | 
 **source** | **string** | Filter by post origin. \&quot;late\&quot; for posts published via Zernio, \&quot;external\&quot; for posts imported from platforms. | [default to &quot;all&quot;]
 **attribution** | **string** | How each post&#39;s engagement is attributed to a day. \&quot;publish\&quot; (default) sums each post&#39;s lifetime total on its publish date. \&quot;received\&quot; buckets the per-day increase in engagement by the day it actually arrived (engagement-over-time), so engagement on older posts appears on the day it was gained rather than the post&#39;s publish date.  | [default to &quot;publish&quot;]

### Return type

[**GetDailyMetrics200Response**](GetDailyMetrics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFacebookPageInsights

> InstagramAccountInsightsResponse GetFacebookPageInsights(ctx).AccountId(accountId).Metrics(metrics).Since(since).Until(until).MetricType(metricType).Execute()

Get Facebook Page insights



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
	accountId := "accountId_example" // string | The Zernio SocialAccount ID for the connected Facebook Page.
	metrics := "metrics_example" // string | Comma-separated list of metrics. Defaults to \"page_media_view,page_post_engagements,page_follows,followers_gained,followers_lost\".  Live Meta metrics (current names, post-Nov-2025):   - page_media_view       (replaces deprecated page_impressions)   - page_views_total   - page_post_engagements   - page_video_views   - page_video_view_time   - page_follows          (replaces deprecated page_fans)  Zernio-synthesized from daily follower snapshots (filling the Nov-2025 gap left by the page_fan_adds / page_fan_removes deprecation):   - followers_gained   - followers_lost  (optional)
	since := time.Now() // string | Start date (YYYY-MM-DD). Defaults to 30 days ago. (optional)
	until := time.Now() // string | End date (YYYY-MM-DD). Defaults to today. (optional)
	metricType := "metricType_example" // string | \"total_value\" (default) returns aggregated totals only. \"time_series\" returns daily values in the \"values\" array.  (optional) (default to "total_value")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetFacebookPageInsights(context.Background()).AccountId(accountId).Metrics(metrics).Since(since).Until(until).MetricType(metricType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetFacebookPageInsights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFacebookPageInsights`: InstagramAccountInsightsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetFacebookPageInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFacebookPageInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The Zernio SocialAccount ID for the connected Facebook Page. | 
 **metrics** | **string** | Comma-separated list of metrics. Defaults to \&quot;page_media_view,page_post_engagements,page_follows,followers_gained,followers_lost\&quot;.  Live Meta metrics (current names, post-Nov-2025):   - page_media_view       (replaces deprecated page_impressions)   - page_views_total   - page_post_engagements   - page_video_views   - page_video_view_time   - page_follows          (replaces deprecated page_fans)  Zernio-synthesized from daily follower snapshots (filling the Nov-2025 gap left by the page_fan_adds / page_fan_removes deprecation):   - followers_gained   - followers_lost  | 
 **since** | **string** | Start date (YYYY-MM-DD). Defaults to 30 days ago. | 
 **until** | **string** | End date (YYYY-MM-DD). Defaults to today. | 
 **metricType** | **string** | \&quot;total_value\&quot; (default) returns aggregated totals only. \&quot;time_series\&quot; returns daily values in the \&quot;values\&quot; array.  | [default to &quot;total_value&quot;]

### Return type

[**InstagramAccountInsightsResponse**](InstagramAccountInsightsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFollowerStats

> GetFollowerStats200Response GetFollowerStats(ctx).AccountIds(accountIds).ProfileId(profileId).FromDate(fromDate).ToDate(toDate).Granularity(granularity).Execute()

Get follower stats



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
	accountIds := "accountIds_example" // string | Comma-separated list of account IDs (optional, defaults to all user's accounts) (optional)
	profileId := "profileId_example" // string | Filter by profile ID (optional)
	fromDate := time.Now() // string | Start date in YYYY-MM-DD format (defaults to 30 days ago) (optional)
	toDate := time.Now() // string | End date in YYYY-MM-DD format (defaults to today) (optional)
	granularity := "granularity_example" // string | Data aggregation level (optional) (default to "daily")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetFollowerStats(context.Background()).AccountIds(accountIds).ProfileId(profileId).FromDate(fromDate).ToDate(toDate).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetFollowerStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFollowerStats`: GetFollowerStats200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetFollowerStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowerStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIds** | **string** | Comma-separated list of account IDs (optional, defaults to all user&#39;s accounts) | 
 **profileId** | **string** | Filter by profile ID | 
 **fromDate** | **string** | Start date in YYYY-MM-DD format (defaults to 30 days ago) | 
 **toDate** | **string** | End date in YYYY-MM-DD format (defaults to today) | 
 **granularity** | **string** | Data aggregation level | [default to &quot;daily&quot;]

### Return type

[**GetFollowerStats200Response**](GetFollowerStats200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGoogleBusinessPerformance

> GetGoogleBusinessPerformance200Response GetGoogleBusinessPerformance(ctx).AccountId(accountId).Metrics(metrics).StartDate(startDate).EndDate(endDate).Execute()

Get GBP performance metrics



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
	accountId := "accountId_example" // string | The Zernio SocialAccount ID for the Google Business Profile account.
	metrics := "metrics_example" // string | Comma-separated metric names. Defaults to all available metrics. Valid values: BUSINESS_IMPRESSIONS_DESKTOP_MAPS, BUSINESS_IMPRESSIONS_DESKTOP_SEARCH, BUSINESS_IMPRESSIONS_MOBILE_MAPS, BUSINESS_IMPRESSIONS_MOBILE_SEARCH, BUSINESS_CONVERSATIONS, BUSINESS_DIRECTION_REQUESTS, CALL_CLICKS, WEBSITE_CLICKS, BUSINESS_BOOKINGS, BUSINESS_FOOD_ORDERS, BUSINESS_FOOD_MENU_CLICKS  (optional)
	startDate := time.Now() // string | Start date (YYYY-MM-DD). Defaults to 30 days ago. Max 18 months back. (optional)
	endDate := time.Now() // string | End date (YYYY-MM-DD). Defaults to today. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetGoogleBusinessPerformance(context.Background()).AccountId(accountId).Metrics(metrics).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetGoogleBusinessPerformance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoogleBusinessPerformance`: GetGoogleBusinessPerformance200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetGoogleBusinessPerformance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGoogleBusinessPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The Zernio SocialAccount ID for the Google Business Profile account. | 
 **metrics** | **string** | Comma-separated metric names. Defaults to all available metrics. Valid values: BUSINESS_IMPRESSIONS_DESKTOP_MAPS, BUSINESS_IMPRESSIONS_DESKTOP_SEARCH, BUSINESS_IMPRESSIONS_MOBILE_MAPS, BUSINESS_IMPRESSIONS_MOBILE_SEARCH, BUSINESS_CONVERSATIONS, BUSINESS_DIRECTION_REQUESTS, CALL_CLICKS, WEBSITE_CLICKS, BUSINESS_BOOKINGS, BUSINESS_FOOD_ORDERS, BUSINESS_FOOD_MENU_CLICKS  | 
 **startDate** | **string** | Start date (YYYY-MM-DD). Defaults to 30 days ago. Max 18 months back. | 
 **endDate** | **string** | End date (YYYY-MM-DD). Defaults to today. | 

### Return type

[**GetGoogleBusinessPerformance200Response**](GetGoogleBusinessPerformance200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGoogleBusinessSearchKeywords

> GetGoogleBusinessSearchKeywords200Response GetGoogleBusinessSearchKeywords(ctx).AccountId(accountId).StartMonth(startMonth).EndMonth(endMonth).Execute()

Get GBP search keywords



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
	accountId := "accountId_example" // string | The Zernio SocialAccount ID for the Google Business Profile account.
	startMonth := "startMonth_example" // string | Start month (YYYY-MM). Defaults to 3 months ago. (optional)
	endMonth := "endMonth_example" // string | End month (YYYY-MM). Defaults to current month. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetGoogleBusinessSearchKeywords(context.Background()).AccountId(accountId).StartMonth(startMonth).EndMonth(endMonth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetGoogleBusinessSearchKeywords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoogleBusinessSearchKeywords`: GetGoogleBusinessSearchKeywords200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetGoogleBusinessSearchKeywords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGoogleBusinessSearchKeywordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The Zernio SocialAccount ID for the Google Business Profile account. | 
 **startMonth** | **string** | Start month (YYYY-MM). Defaults to 3 months ago. | 
 **endMonth** | **string** | End month (YYYY-MM). Defaults to current month. | 

### Return type

[**GetGoogleBusinessSearchKeywords200Response**](GetGoogleBusinessSearchKeywords200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstagramAccountInsights

> InstagramAccountInsightsResponse GetInstagramAccountInsights(ctx).AccountId(accountId).Metrics(metrics).Since(since).Until(until).MetricType(metricType).Breakdown(breakdown).Execute()

Get Instagram insights



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
	accountId := "accountId_example" // string | The Zernio SocialAccount ID for the Instagram account
	metrics := "metrics_example" // string | Comma-separated list of metrics. Defaults to \"reach,views,accounts_engaged,total_interactions\". Valid metrics: reach, views, accounts_engaged, total_interactions, comments, likes, saves, shares, replies, reposts, follows_and_unfollows, profile_links_taps. Note: only \"reach\" supports metricType=time_series. All other metrics (including follows_and_unfollows) are total_value only. This is an Instagram Graph API limitation, not a Zernio limitation - the IG API does not return time-series data for these metrics. For a daily running follower count, use /v1/analytics/instagram/follower-history instead.  (optional)
	since := time.Now() // string | Start date (YYYY-MM-DD). Defaults to 30 days ago. (optional)
	until := time.Now() // string | End date (YYYY-MM-DD). Defaults to today. (optional)
	metricType := "metricType_example" // string | \"total_value\" (default) returns aggregated totals and supports breakdowns. \"time_series\" returns daily values but only works with the \"reach\" metric.  (optional) (default to "total_value")
	breakdown := "breakdown_example" // string | Breakdown dimension (only valid with metricType=total_value). Valid values depend on the metric: media_product_type, follow_type, follower_type, contact_button_type.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetInstagramAccountInsights(context.Background()).AccountId(accountId).Metrics(metrics).Since(since).Until(until).MetricType(metricType).Breakdown(breakdown).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetInstagramAccountInsights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstagramAccountInsights`: InstagramAccountInsightsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetInstagramAccountInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstagramAccountInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The Zernio SocialAccount ID for the Instagram account | 
 **metrics** | **string** | Comma-separated list of metrics. Defaults to \&quot;reach,views,accounts_engaged,total_interactions\&quot;. Valid metrics: reach, views, accounts_engaged, total_interactions, comments, likes, saves, shares, replies, reposts, follows_and_unfollows, profile_links_taps. Note: only \&quot;reach\&quot; supports metricType&#x3D;time_series. All other metrics (including follows_and_unfollows) are total_value only. This is an Instagram Graph API limitation, not a Zernio limitation - the IG API does not return time-series data for these metrics. For a daily running follower count, use /v1/analytics/instagram/follower-history instead.  | 
 **since** | **string** | Start date (YYYY-MM-DD). Defaults to 30 days ago. | 
 **until** | **string** | End date (YYYY-MM-DD). Defaults to today. | 
 **metricType** | **string** | \&quot;total_value\&quot; (default) returns aggregated totals and supports breakdowns. \&quot;time_series\&quot; returns daily values but only works with the \&quot;reach\&quot; metric.  | [default to &quot;total_value&quot;]
 **breakdown** | **string** | Breakdown dimension (only valid with metricType&#x3D;total_value). Valid values depend on the metric: media_product_type, follow_type, follower_type, contact_button_type.  | 

### Return type

[**InstagramAccountInsightsResponse**](InstagramAccountInsightsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstagramDemographics

> InstagramDemographicsResponse GetInstagramDemographics(ctx).AccountId(accountId).Metric(metric).Breakdown(breakdown).Timeframe(timeframe).Execute()

Get Instagram demographics



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
	accountId := "accountId_example" // string | The Zernio SocialAccount ID for the Instagram account
	metric := "metric_example" // string | \"follower_demographics\" for follower audience data, or \"engaged_audience_demographics\" for engaged viewers.  (optional) (default to "follower_demographics")
	breakdown := "breakdown_example" // string | Comma-separated list of demographic dimensions: age, city, country, gender. Defaults to all four if omitted.  (optional)
	timeframe := "timeframe_example" // string | Time period for demographic data. Defaults to \"this_month\".  (optional) (default to "this_month")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetInstagramDemographics(context.Background()).AccountId(accountId).Metric(metric).Breakdown(breakdown).Timeframe(timeframe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetInstagramDemographics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstagramDemographics`: InstagramDemographicsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetInstagramDemographics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstagramDemographicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The Zernio SocialAccount ID for the Instagram account | 
 **metric** | **string** | \&quot;follower_demographics\&quot; for follower audience data, or \&quot;engaged_audience_demographics\&quot; for engaged viewers.  | [default to &quot;follower_demographics&quot;]
 **breakdown** | **string** | Comma-separated list of demographic dimensions: age, city, country, gender. Defaults to all four if omitted.  | 
 **timeframe** | **string** | Time period for demographic data. Defaults to \&quot;this_month\&quot;.  | [default to &quot;this_month&quot;]

### Return type

[**InstagramDemographicsResponse**](InstagramDemographicsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstagramFollowerHistory

> InstagramAccountInsightsResponse GetInstagramFollowerHistory(ctx).AccountId(accountId).Metrics(metrics).Since(since).Until(until).MetricType(metricType).Execute()

Get Instagram follower history



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
	accountId := "accountId_example" // string | The Zernio SocialAccount ID for the Instagram account.
	metrics := "metrics_example" // string | Comma-separated list. Defaults to \"follower_count,followers_gained,followers_lost\".   - follower_count   : per-day raw follower count   - followers_gained : sum of positive daily deltas   - followers_lost   : sum of absolute negative daily deltas  (optional)
	since := time.Now() // string | Start date (YYYY-MM-DD). Defaults to 30 days ago. (optional)
	until := time.Now() // string | End date (YYYY-MM-DD). Defaults to today. (optional)
	metricType := "metricType_example" // string | \"total_value\" returns aggregated totals (latest for follower_count, sum for gained/lost). \"time_series\" returns per-day values in the \"values\" array.  (optional) (default to "total_value")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetInstagramFollowerHistory(context.Background()).AccountId(accountId).Metrics(metrics).Since(since).Until(until).MetricType(metricType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetInstagramFollowerHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstagramFollowerHistory`: InstagramAccountInsightsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetInstagramFollowerHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstagramFollowerHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The Zernio SocialAccount ID for the Instagram account. | 
 **metrics** | **string** | Comma-separated list. Defaults to \&quot;follower_count,followers_gained,followers_lost\&quot;.   - follower_count   : per-day raw follower count   - followers_gained : sum of positive daily deltas   - followers_lost   : sum of absolute negative daily deltas  | 
 **since** | **string** | Start date (YYYY-MM-DD). Defaults to 30 days ago. | 
 **until** | **string** | End date (YYYY-MM-DD). Defaults to today. | 
 **metricType** | **string** | \&quot;total_value\&quot; returns aggregated totals (latest for follower_count, sum for gained/lost). \&quot;time_series\&quot; returns per-day values in the \&quot;values\&quot; array.  | [default to &quot;total_value&quot;]

### Return type

[**InstagramAccountInsightsResponse**](InstagramAccountInsightsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedInAggregateAnalytics

> GetLinkedInAggregateAnalytics200Response GetLinkedInAggregateAnalytics(ctx, accountId).Aggregation(aggregation).StartDate(startDate).EndDate(endDate).Metrics(metrics).Execute()

Get LinkedIn aggregate stats



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
	accountId := "accountId_example" // string | The ID of the LinkedIn personal account
	aggregation := "aggregation_example" // string | TOTAL (default, lifetime totals) or DAILY (time series). MEMBERS_REACHED not available with DAILY. (optional) (default to "TOTAL")
	startDate := time.Now() // string | Start date (YYYY-MM-DD). If omitted, returns lifetime analytics. (optional)
	endDate := time.Now() // string | End date (YYYY-MM-DD, exclusive). Defaults to today if omitted. (optional)
	metrics := "IMPRESSION,REACTION,COMMENT,POST_SAVE,POST_SEND" // string | Comma-separated metrics: IMPRESSION, MEMBERS_REACHED, REACTION, COMMENT, RESHARE, POST_SAVE, POST_SEND. Omit for all. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetLinkedInAggregateAnalytics(context.Background(), accountId).Aggregation(aggregation).StartDate(startDate).EndDate(endDate).Metrics(metrics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetLinkedInAggregateAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedInAggregateAnalytics`: GetLinkedInAggregateAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetLinkedInAggregateAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the LinkedIn personal account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedInAggregateAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aggregation** | **string** | TOTAL (default, lifetime totals) or DAILY (time series). MEMBERS_REACHED not available with DAILY. | [default to &quot;TOTAL&quot;]
 **startDate** | **string** | Start date (YYYY-MM-DD). If omitted, returns lifetime analytics. | 
 **endDate** | **string** | End date (YYYY-MM-DD, exclusive). Defaults to today if omitted. | 
 **metrics** | **string** | Comma-separated metrics: IMPRESSION, MEMBERS_REACHED, REACTION, COMMENT, RESHARE, POST_SAVE, POST_SEND. Omit for all. | 

### Return type

[**GetLinkedInAggregateAnalytics200Response**](GetLinkedInAggregateAnalytics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedInOrgAggregateAnalytics

> InstagramAccountInsightsResponse GetLinkedInOrgAggregateAnalytics(ctx).AccountId(accountId).Metrics(metrics).Since(since).Until(until).MetricType(metricType).Execute()

Get LinkedIn organization page aggregate analytics



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
	accountId := "accountId_example" // string | The Zernio SocialAccount ID for the LinkedIn organization account.
	metrics := "metrics_example" // string | Comma-separated list. Defaults to \"impressions,clicks,engagement_rate,organic_followers_gained,followers_gained,followers_lost\".  Share statistics (support both total_value and time_series):   - impressions   - unique_impressions   - clicks   - likes   - comments   - shares   - engagement_rate       (0..1, LinkedIn-computed)  Follower-gain statistics (support total_value and time_series):   - organic_followers_gained   (per-day organic gains for time_series; sum of organic gains over the range for total_value)   - paid_followers_gained      (per-day paid gains for time_series; sum of paid gains over the range for total_value)  Page-view statistics (total_value ONLY - LinkedIn platform limit):   - page_views_total   - page_views_overview   - page_views_careers   - page_views_jobs   - page_views_life  Zernio-synthesized from daily follower snapshots:   - followers_gained   - followers_lost  (optional)
	since := time.Now() // string | Start date (YYYY-MM-DD). Defaults to 30 days ago. (optional)
	until := time.Now() // string | End date (YYYY-MM-DD). Defaults to today. (optional)
	metricType := "metricType_example" // string |  (optional) (default to "total_value")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetLinkedInOrgAggregateAnalytics(context.Background()).AccountId(accountId).Metrics(metrics).Since(since).Until(until).MetricType(metricType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetLinkedInOrgAggregateAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedInOrgAggregateAnalytics`: InstagramAccountInsightsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetLinkedInOrgAggregateAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedInOrgAggregateAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The Zernio SocialAccount ID for the LinkedIn organization account. | 
 **metrics** | **string** | Comma-separated list. Defaults to \&quot;impressions,clicks,engagement_rate,organic_followers_gained,followers_gained,followers_lost\&quot;.  Share statistics (support both total_value and time_series):   - impressions   - unique_impressions   - clicks   - likes   - comments   - shares   - engagement_rate       (0..1, LinkedIn-computed)  Follower-gain statistics (support total_value and time_series):   - organic_followers_gained   (per-day organic gains for time_series; sum of organic gains over the range for total_value)   - paid_followers_gained      (per-day paid gains for time_series; sum of paid gains over the range for total_value)  Page-view statistics (total_value ONLY - LinkedIn platform limit):   - page_views_total   - page_views_overview   - page_views_careers   - page_views_jobs   - page_views_life  Zernio-synthesized from daily follower snapshots:   - followers_gained   - followers_lost  | 
 **since** | **string** | Start date (YYYY-MM-DD). Defaults to 30 days ago. | 
 **until** | **string** | End date (YYYY-MM-DD). Defaults to today. | 
 **metricType** | **string** |  | [default to &quot;total_value&quot;]

### Return type

[**InstagramAccountInsightsResponse**](InstagramAccountInsightsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedInPostAnalytics

> GetLinkedInPostAnalytics200Response GetLinkedInPostAnalytics(ctx, accountId).Urn(urn).Execute()

Get LinkedIn post stats



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
	accountId := "accountId_example" // string | The ID of the LinkedIn account
	urn := "urn:li:share:7123456789012345678" // string | The LinkedIn post URN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetLinkedInPostAnalytics(context.Background(), accountId).Urn(urn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetLinkedInPostAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedInPostAnalytics`: GetLinkedInPostAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetLinkedInPostAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the LinkedIn account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedInPostAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urn** | **string** | The LinkedIn post URN | 

### Return type

[**GetLinkedInPostAnalytics200Response**](GetLinkedInPostAnalytics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedInPostReactions

> GetLinkedInPostReactions200Response GetLinkedInPostReactions(ctx, accountId).Urn(urn).Limit(limit).Cursor(cursor).Execute()

Get LinkedIn post reactions



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
	accountId := "accountId_example" // string | The ID of the LinkedIn organization account
	urn := "urn:li:share:7123456789012345678" // string | The LinkedIn post URN
	limit := int32(56) // int32 | Maximum number of reactions to return per page (optional) (default to 25)
	cursor := "cursor_example" // string | Offset-based pagination start index (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetLinkedInPostReactions(context.Background(), accountId).Urn(urn).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetLinkedInPostReactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedInPostReactions`: GetLinkedInPostReactions200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetLinkedInPostReactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the LinkedIn organization account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedInPostReactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **urn** | **string** | The LinkedIn post URN | 
 **limit** | **int32** | Maximum number of reactions to return per page | [default to 25]
 **cursor** | **string** | Offset-based pagination start index | 

### Return type

[**GetLinkedInPostReactions200Response**](GetLinkedInPostReactions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostTimeline

> GetPostTimeline200Response GetPostTimeline(ctx).PostId(postId).FromDate(fromDate).ToDate(toDate).Execute()

Get post analytics timeline



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
	postId := "postId_example" // string | The post to fetch timeline for. Accepts an ExternalPost ID, a platformPostId, or a Zernio Post ID. 
	fromDate := time.Now() // time.Time | Start of date range (ISO 8601). Defaults to 90 days ago. (optional)
	toDate := time.Now() // time.Time | End of date range (ISO 8601). Defaults to now. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetPostTimeline(context.Background()).PostId(postId).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetPostTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostTimeline`: GetPostTimeline200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetPostTimeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postId** | **string** | The post to fetch timeline for. Accepts an ExternalPost ID, a platformPostId, or a Zernio Post ID.  | 
 **fromDate** | **time.Time** | Start of date range (ISO 8601). Defaults to 90 days ago. | 
 **toDate** | **time.Time** | End of date range (ISO 8601). Defaults to now. | 

### Return type

[**GetPostTimeline200Response**](GetPostTimeline200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostingFrequency

> GetPostingFrequency200Response GetPostingFrequency(ctx).Platform(platform).ProfileId(profileId).AccountId(accountId).Source(source).Execute()

Get frequency vs engagement



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
	platform := "platform_example" // string | Filter by platform (e.g. \"instagram\", \"tiktok\"). Omit for all platforms. (optional)
	profileId := "profileId_example" // string | Filter by profile ID. Omit for all profiles. (optional)
	accountId := "accountId_example" // string | Filter by social account ID. Omit for all accounts. (optional)
	source := "source_example" // string | Filter by post origin. \"late\" for posts published via Zernio, \"external\" for posts imported from platforms. (optional) (default to "all")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetPostingFrequency(context.Background()).Platform(platform).ProfileId(profileId).AccountId(accountId).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetPostingFrequency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostingFrequency`: GetPostingFrequency200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetPostingFrequency`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostingFrequencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **string** | Filter by platform (e.g. \&quot;instagram\&quot;, \&quot;tiktok\&quot;). Omit for all platforms. | 
 **profileId** | **string** | Filter by profile ID. Omit for all profiles. | 
 **accountId** | **string** | Filter by social account ID. Omit for all accounts. | 
 **source** | **string** | Filter by post origin. \&quot;late\&quot; for posts published via Zernio, \&quot;external\&quot; for posts imported from platforms. | [default to &quot;all&quot;]

### Return type

[**GetPostingFrequency200Response**](GetPostingFrequency200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTikTokAccountInsights

> InstagramAccountInsightsResponse GetTikTokAccountInsights(ctx).AccountId(accountId).Metrics(metrics).Since(since).Until(until).MetricType(metricType).Execute()

Get TikTok account-level insights



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
	accountId := "accountId_example" // string | The Zernio SocialAccount ID for the TikTok account.
	metrics := "metrics_example" // string | Comma-separated list. Defaults to \"follower_count,likes_count,video_count,followers_gained,followers_lost\".  Live from /v2/user/info/ (requires user.info.stats scope):   - follower_count  (cumulative; time series joined from AccountStats)   - following_count (cumulative; time series joined from AccountStats.metadata)   - likes_count     (cumulative; time series joined from AccountStats.metadata)   - video_count     (cumulative; time series joined from AccountStats.metadata)  Zernio-synthesized:   - followers_gained  (sum of positive daily follower deltas)   - followers_lost    (sum of absolute negative daily deltas)  (optional)
	since := time.Now() // string | Start date (YYYY-MM-DD). Defaults to 30 days ago. (optional)
	until := time.Now() // string | End date (YYYY-MM-DD). Defaults to today. (optional)
	metricType := "metricType_example" // string | \"total_value\" returns the latest cumulative counter value. \"time_series\" returns daily values joined from AccountStats snapshots.  (optional) (default to "total_value")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetTikTokAccountInsights(context.Background()).AccountId(accountId).Metrics(metrics).Since(since).Until(until).MetricType(metricType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetTikTokAccountInsights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTikTokAccountInsights`: InstagramAccountInsightsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetTikTokAccountInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTikTokAccountInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The Zernio SocialAccount ID for the TikTok account. | 
 **metrics** | **string** | Comma-separated list. Defaults to \&quot;follower_count,likes_count,video_count,followers_gained,followers_lost\&quot;.  Live from /v2/user/info/ (requires user.info.stats scope):   - follower_count  (cumulative; time series joined from AccountStats)   - following_count (cumulative; time series joined from AccountStats.metadata)   - likes_count     (cumulative; time series joined from AccountStats.metadata)   - video_count     (cumulative; time series joined from AccountStats.metadata)  Zernio-synthesized:   - followers_gained  (sum of positive daily follower deltas)   - followers_lost    (sum of absolute negative daily deltas)  | 
 **since** | **string** | Start date (YYYY-MM-DD). Defaults to 30 days ago. | 
 **until** | **string** | End date (YYYY-MM-DD). Defaults to today. | 
 **metricType** | **string** | \&quot;total_value\&quot; returns the latest cumulative counter value. \&quot;time_series\&quot; returns daily values joined from AccountStats snapshots.  | [default to &quot;total_value&quot;]

### Return type

[**InstagramAccountInsightsResponse**](InstagramAccountInsightsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetYouTubeChannelInsights

> InstagramAccountInsightsResponse GetYouTubeChannelInsights(ctx).AccountId(accountId).Metrics(metrics).Since(since).Until(until).MetricType(metricType).Execute()

Get YouTube channel-level insights



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
	accountId := "accountId_example" // string | The Zernio SocialAccount ID for the YouTube account.
	metrics := "metrics_example" // string | Comma-separated list. Defaults to \"views,estimatedMinutesWatched,subscribersGained,subscribersLost\".  Live YouTube Analytics v2 metrics:   - views   - estimatedMinutesWatched   - averageViewDuration          (ratio - weighted mean computed across days)   - subscribersGained   - subscribersLost  Zernio-synthesized from daily follower snapshots (cross-platform parity):   - followers_gained   - followers_lost  (optional)
	since := time.Now() // string | Start date (YYYY-MM-DD). Defaults to 30 days ago. (optional)
	until := time.Now() // string | End date (YYYY-MM-DD). Defaults to today. YouTube Analytics has a 2-3 day delay, so the fetch is internally clamped to 3 days ago; any requested range extending beyond that returns zero values for the tail days. The response's dateRange.until field reflects your requested value.  (optional)
	metricType := "metricType_example" // string | \"total_value\" (default) returns aggregated totals. \"time_series\" returns per-day values in the \"values\" array.  (optional) (default to "total_value")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetYouTubeChannelInsights(context.Background()).AccountId(accountId).Metrics(metrics).Since(since).Until(until).MetricType(metricType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetYouTubeChannelInsights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetYouTubeChannelInsights`: InstagramAccountInsightsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetYouTubeChannelInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetYouTubeChannelInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The Zernio SocialAccount ID for the YouTube account. | 
 **metrics** | **string** | Comma-separated list. Defaults to \&quot;views,estimatedMinutesWatched,subscribersGained,subscribersLost\&quot;.  Live YouTube Analytics v2 metrics:   - views   - estimatedMinutesWatched   - averageViewDuration          (ratio - weighted mean computed across days)   - subscribersGained   - subscribersLost  Zernio-synthesized from daily follower snapshots (cross-platform parity):   - followers_gained   - followers_lost  | 
 **since** | **string** | Start date (YYYY-MM-DD). Defaults to 30 days ago. | 
 **until** | **string** | End date (YYYY-MM-DD). Defaults to today. YouTube Analytics has a 2-3 day delay, so the fetch is internally clamped to 3 days ago; any requested range extending beyond that returns zero values for the tail days. The response&#39;s dateRange.until field reflects your requested value.  | 
 **metricType** | **string** | \&quot;total_value\&quot; (default) returns aggregated totals. \&quot;time_series\&quot; returns per-day values in the \&quot;values\&quot; array.  | [default to &quot;total_value&quot;]

### Return type

[**InstagramAccountInsightsResponse**](InstagramAccountInsightsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetYouTubeDailyViews

> YouTubeDailyViewsResponse GetYouTubeDailyViews(ctx).VideoId(videoId).AccountId(accountId).StartDate(startDate).EndDate(endDate).Execute()

Get YouTube daily views



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
	videoId := "videoId_example" // string | The YouTube video ID (e.g., \"dQw4w9WgXcQ\")
	accountId := "accountId_example" // string | The Zernio account ID for the YouTube account
	startDate := time.Now() // string | Start date (YYYY-MM-DD). Defaults to 30 days ago. (optional)
	endDate := time.Now() // string | End date (YYYY-MM-DD). Defaults to 3 days ago (YouTube data latency). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetYouTubeDailyViews(context.Background()).VideoId(videoId).AccountId(accountId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetYouTubeDailyViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetYouTubeDailyViews`: YouTubeDailyViewsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetYouTubeDailyViews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetYouTubeDailyViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **string** | The YouTube video ID (e.g., \&quot;dQw4w9WgXcQ\&quot;) | 
 **accountId** | **string** | The Zernio account ID for the YouTube account | 
 **startDate** | **string** | Start date (YYYY-MM-DD). Defaults to 30 days ago. | 
 **endDate** | **string** | End date (YYYY-MM-DD). Defaults to 3 days ago (YouTube data latency). | 

### Return type

[**YouTubeDailyViewsResponse**](YouTubeDailyViewsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetYouTubeDemographics

> YouTubeDemographicsResponse GetYouTubeDemographics(ctx).AccountId(accountId).Breakdown(breakdown).StartDate(startDate).EndDate(endDate).Execute()

Get YouTube demographics



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
	accountId := "accountId_example" // string | The Zernio SocialAccount ID for the YouTube account
	breakdown := "breakdown_example" // string | Comma-separated list of demographic dimensions: age, gender, country. Defaults to all three if omitted.  (optional)
	startDate := time.Now() // string | Start date in YYYY-MM-DD format. Defaults to 90 days ago.  (optional)
	endDate := time.Now() // string | End date in YYYY-MM-DD format. Defaults to 3 days ago (YouTube data latency).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetYouTubeDemographics(context.Background()).AccountId(accountId).Breakdown(breakdown).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetYouTubeDemographics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetYouTubeDemographics`: YouTubeDemographicsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetYouTubeDemographics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetYouTubeDemographicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The Zernio SocialAccount ID for the YouTube account | 
 **breakdown** | **string** | Comma-separated list of demographic dimensions: age, gender, country. Defaults to all three if omitted.  | 
 **startDate** | **string** | Start date in YYYY-MM-DD format. Defaults to 90 days ago.  | 
 **endDate** | **string** | End date in YYYY-MM-DD format. Defaults to 3 days ago (YouTube data latency).  | 

### Return type

[**YouTubeDemographicsResponse**](YouTubeDemographicsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetYouTubeVideoRetention

> YouTubeVideoRetentionResponse GetYouTubeVideoRetention(ctx).VideoId(videoId).AccountId(accountId).StartDate(startDate).EndDate(endDate).Execute()

Get YouTube video retention curve



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
	videoId := "videoId_example" // string | The YouTube video ID (e.g., \"dQw4w9WgXcQ\")
	accountId := "accountId_example" // string | The Zernio account ID for the YouTube account
	startDate := time.Now() // string | Start date (YYYY-MM-DD). Defaults to the video's publish date (lifetime curve). (optional)
	endDate := time.Now() // string | End date (YYYY-MM-DD). Defaults to 3 days ago (YouTube data latency). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetYouTubeVideoRetention(context.Background()).VideoId(videoId).AccountId(accountId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetYouTubeVideoRetention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetYouTubeVideoRetention`: YouTubeVideoRetentionResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetYouTubeVideoRetention`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetYouTubeVideoRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **string** | The YouTube video ID (e.g., \&quot;dQw4w9WgXcQ\&quot;) | 
 **accountId** | **string** | The Zernio account ID for the YouTube account | 
 **startDate** | **string** | Start date (YYYY-MM-DD). Defaults to the video&#39;s publish date (lifetime curve). | 
 **endDate** | **string** | End date (YYYY-MM-DD). Defaults to 3 days ago (YouTube data latency). | 

### Return type

[**YouTubeVideoRetentionResponse**](YouTubeVideoRetentionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

