# \RedditSearchAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRedditFeed**](RedditSearchAPI.md#GetRedditFeed) | **Get** /v1/reddit/feed | Get subreddit feed
[**SearchReddit**](RedditSearchAPI.md#SearchReddit) | **Get** /v1/reddit/search | Search posts



## GetRedditFeed

> SearchReddit200Response GetRedditFeed(ctx).AccountId(accountId).Subreddit(subreddit).Sort(sort).Limit(limit).After(after).T(t).Execute()

Get subreddit feed



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
	subreddit := "subreddit_example" // string |  (optional)
	sort := "sort_example" // string |  (optional) (default to "hot")
	limit := int32(56) // int32 |  (optional) (default to 25)
	after := "after_example" // string |  (optional)
	t := "t_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedditSearchAPI.GetRedditFeed(context.Background()).AccountId(accountId).Subreddit(subreddit).Sort(sort).Limit(limit).After(after).T(t).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedditSearchAPI.GetRedditFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedditFeed`: SearchReddit200Response
	fmt.Fprintf(os.Stdout, "Response from `RedditSearchAPI.GetRedditFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRedditFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **subreddit** | **string** |  | 
 **sort** | **string** |  | [default to &quot;hot&quot;]
 **limit** | **int32** |  | [default to 25]
 **after** | **string** |  | 
 **t** | **string** |  | 

### Return type

[**SearchReddit200Response**](SearchReddit200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchReddit

> SearchReddit200Response SearchReddit(ctx).AccountId(accountId).Q(q).Subreddit(subreddit).RestrictSr(restrictSr).Sort(sort).Limit(limit).After(after).Execute()

Search posts



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
	q := "q_example" // string | 
	subreddit := "subreddit_example" // string |  (optional)
	restrictSr := "restrictSr_example" // string |  (optional)
	sort := "sort_example" // string |  (optional) (default to "new")
	limit := int32(56) // int32 |  (optional) (default to 25)
	after := "after_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedditSearchAPI.SearchReddit(context.Background()).AccountId(accountId).Q(q).Subreddit(subreddit).RestrictSr(restrictSr).Sort(sort).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedditSearchAPI.SearchReddit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchReddit`: SearchReddit200Response
	fmt.Fprintf(os.Stdout, "Response from `RedditSearchAPI.SearchReddit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRedditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **q** | **string** |  | 
 **subreddit** | **string** |  | 
 **restrictSr** | **string** |  | 
 **sort** | **string** |  | [default to &quot;new&quot;]
 **limit** | **int32** |  | [default to 25]
 **after** | **string** |  | 

### Return type

[**SearchReddit200Response**](SearchReddit200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

