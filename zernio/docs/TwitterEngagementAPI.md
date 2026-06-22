# \TwitterEngagementAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookmarkPost**](TwitterEngagementAPI.md#BookmarkPost) | **Post** /v1/twitter/bookmark | Bookmark a tweet
[**FollowUser**](TwitterEngagementAPI.md#FollowUser) | **Post** /v1/twitter/follow | Follow a user
[**RemoveBookmark**](TwitterEngagementAPI.md#RemoveBookmark) | **Delete** /v1/twitter/bookmark | Remove bookmark
[**RetweetPost**](TwitterEngagementAPI.md#RetweetPost) | **Post** /v1/twitter/retweet | Retweet a post
[**UndoRetweet**](TwitterEngagementAPI.md#UndoRetweet) | **Delete** /v1/twitter/retweet | Undo retweet
[**UnfollowUser**](TwitterEngagementAPI.md#UnfollowUser) | **Delete** /v1/twitter/follow | Unfollow a user



## BookmarkPost

> BookmarkPost200Response BookmarkPost(ctx).RetweetPostRequest(retweetPostRequest).Execute()

Bookmark a tweet



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
	retweetPostRequest := *openapiclient.NewRetweetPostRequest("AccountId_example", "TweetId_example") // RetweetPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TwitterEngagementAPI.BookmarkPost(context.Background()).RetweetPostRequest(retweetPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwitterEngagementAPI.BookmarkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookmarkPost`: BookmarkPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TwitterEngagementAPI.BookmarkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookmarkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **retweetPostRequest** | [**RetweetPostRequest**](RetweetPostRequest.md) |  | 

### Return type

[**BookmarkPost200Response**](BookmarkPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowUser

> FollowUser200Response FollowUser(ctx).FollowUserRequest(followUserRequest).Execute()

Follow a user



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
	followUserRequest := *openapiclient.NewFollowUserRequest("AccountId_example", "TargetUserId_example") // FollowUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TwitterEngagementAPI.FollowUser(context.Background()).FollowUserRequest(followUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwitterEngagementAPI.FollowUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FollowUser`: FollowUser200Response
	fmt.Fprintf(os.Stdout, "Response from `TwitterEngagementAPI.FollowUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFollowUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **followUserRequest** | [**FollowUserRequest**](FollowUserRequest.md) |  | 

### Return type

[**FollowUser200Response**](FollowUser200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveBookmark

> BookmarkPost200Response RemoveBookmark(ctx).AccountId(accountId).TweetId(tweetId).Execute()

Remove bookmark



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
	tweetId := "tweetId_example" // string | The ID of the tweet to unbookmark

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TwitterEngagementAPI.RemoveBookmark(context.Background()).AccountId(accountId).TweetId(tweetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwitterEngagementAPI.RemoveBookmark``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveBookmark`: BookmarkPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TwitterEngagementAPI.RemoveBookmark`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveBookmarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **tweetId** | **string** | The ID of the tweet to unbookmark | 

### Return type

[**BookmarkPost200Response**](BookmarkPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetweetPost

> RetweetPost200Response RetweetPost(ctx).RetweetPostRequest(retweetPostRequest).Execute()

Retweet a post



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
	retweetPostRequest := *openapiclient.NewRetweetPostRequest("AccountId_example", "TweetId_example") // RetweetPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TwitterEngagementAPI.RetweetPost(context.Background()).RetweetPostRequest(retweetPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwitterEngagementAPI.RetweetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetweetPost`: RetweetPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TwitterEngagementAPI.RetweetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetweetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **retweetPostRequest** | [**RetweetPostRequest**](RetweetPostRequest.md) |  | 

### Return type

[**RetweetPost200Response**](RetweetPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UndoRetweet

> RetweetPost200Response UndoRetweet(ctx).AccountId(accountId).TweetId(tweetId).Execute()

Undo retweet



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
	tweetId := "tweetId_example" // string | The ID of the original tweet to un-retweet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TwitterEngagementAPI.UndoRetweet(context.Background()).AccountId(accountId).TweetId(tweetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwitterEngagementAPI.UndoRetweet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UndoRetweet`: RetweetPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TwitterEngagementAPI.UndoRetweet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUndoRetweetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **tweetId** | **string** | The ID of the original tweet to un-retweet | 

### Return type

[**RetweetPost200Response**](RetweetPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnfollowUser

> UnfollowUser200Response UnfollowUser(ctx).AccountId(accountId).TargetUserId(targetUserId).Execute()

Unfollow a user



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
	targetUserId := "targetUserId_example" // string | The Twitter ID of the user to unfollow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TwitterEngagementAPI.UnfollowUser(context.Background()).AccountId(accountId).TargetUserId(targetUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwitterEngagementAPI.UnfollowUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnfollowUser`: UnfollowUser200Response
	fmt.Fprintf(os.Stdout, "Response from `TwitterEngagementAPI.UnfollowUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnfollowUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **targetUserId** | **string** | The Twitter ID of the user to unfollow | 

### Return type

[**UnfollowUser200Response**](UnfollowUser200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

