# \CommentsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInboxComment**](CommentsAPI.md#DeleteInboxComment) | **Delete** /v1/inbox/comments/{postId} | Delete comment
[**GetInboxPostComments**](CommentsAPI.md#GetInboxPostComments) | **Get** /v1/inbox/comments/{postId} | Get post comments
[**HideInboxComment**](CommentsAPI.md#HideInboxComment) | **Post** /v1/inbox/comments/{postId}/{commentId}/hide | Hide comment
[**LikeInboxComment**](CommentsAPI.md#LikeInboxComment) | **Post** /v1/inbox/comments/{postId}/{commentId}/like | Like comment
[**ListInboxComments**](CommentsAPI.md#ListInboxComments) | **Get** /v1/inbox/comments | List commented posts
[**ReplyToInboxPost**](CommentsAPI.md#ReplyToInboxPost) | **Post** /v1/inbox/comments/{postId} | Reply to comment
[**SendPrivateReplyToComment**](CommentsAPI.md#SendPrivateReplyToComment) | **Post** /v1/inbox/comments/{postId}/{commentId}/private-reply | Send private reply
[**UnhideInboxComment**](CommentsAPI.md#UnhideInboxComment) | **Delete** /v1/inbox/comments/{postId}/{commentId}/hide | Unhide comment
[**UnlikeInboxComment**](CommentsAPI.md#UnlikeInboxComment) | **Delete** /v1/inbox/comments/{postId}/{commentId}/like | Unlike comment



## DeleteInboxComment

> DeleteInboxComment200Response DeleteInboxComment(ctx, postId).AccountId(accountId).CommentId(commentId).Execute()

Delete comment



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
	postId := "postId_example" // string | Zernio post ID or platform-specific post ID. LinkedIn third-party posts accept full activity URN or numeric ID.
	accountId := "accountId_example" // string | 
	commentId := "commentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.DeleteInboxComment(context.Background(), postId).AccountId(accountId).CommentId(commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.DeleteInboxComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInboxComment`: DeleteInboxComment200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.DeleteInboxComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | Zernio post ID or platform-specific post ID. LinkedIn third-party posts accept full activity URN or numeric ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInboxCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** |  | 
 **commentId** | **string** |  | 

### Return type

[**DeleteInboxComment200Response**](DeleteInboxComment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxPostComments

> GetInboxPostComments200Response GetInboxPostComments(ctx, postId).AccountId(accountId).Subreddit(subreddit).Limit(limit).Cursor(cursor).CommentId(commentId).Execute()

Get post comments



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
	postId := "postId_example" // string | Zernio post ID or platform-specific post ID. Zernio IDs are auto-resolved. LinkedIn third-party posts accept full activity URN or numeric ID.
	accountId := "accountId_example" // string | 
	subreddit := "subreddit_example" // string | (Reddit only) Subreddit name (optional)
	limit := int32(56) // int32 | Maximum number of comments to return (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor (optional)
	commentId := "commentId_example" // string | (Reddit only) Get replies to a specific comment (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.GetInboxPostComments(context.Background(), postId).AccountId(accountId).Subreddit(subreddit).Limit(limit).Cursor(cursor).CommentId(commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.GetInboxPostComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxPostComments`: GetInboxPostComments200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.GetInboxPostComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | Zernio post ID or platform-specific post ID. Zernio IDs are auto-resolved. LinkedIn third-party posts accept full activity URN or numeric ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxPostCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** |  | 
 **subreddit** | **string** | (Reddit only) Subreddit name | 
 **limit** | **int32** | Maximum number of comments to return | [default to 25]
 **cursor** | **string** | Pagination cursor | 
 **commentId** | **string** | (Reddit only) Get replies to a specific comment | 

### Return type

[**GetInboxPostComments200Response**](GetInboxPostComments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HideInboxComment

> HideInboxComment200Response HideInboxComment(ctx, postId, commentId).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()

Hide comment



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
	postId := "postId_example" // string | 
	commentId := "commentId_example" // string | 
	sendTypingIndicatorRequest := *openapiclient.NewSendTypingIndicatorRequest("AccountId_example") // SendTypingIndicatorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.HideInboxComment(context.Background(), postId, commentId).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.HideInboxComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HideInboxComment`: HideInboxComment200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.HideInboxComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHideInboxCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sendTypingIndicatorRequest** | [**SendTypingIndicatorRequest**](SendTypingIndicatorRequest.md) |  | 

### Return type

[**HideInboxComment200Response**](HideInboxComment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LikeInboxComment

> LikeInboxComment200Response LikeInboxComment(ctx, postId, commentId).LikeInboxCommentRequest(likeInboxCommentRequest).Execute()

Like comment



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
	postId := "postId_example" // string | 
	commentId := "commentId_example" // string | 
	likeInboxCommentRequest := *openapiclient.NewLikeInboxCommentRequest("AccountId_example") // LikeInboxCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.LikeInboxComment(context.Background(), postId, commentId).LikeInboxCommentRequest(likeInboxCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.LikeInboxComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LikeInboxComment`: LikeInboxComment200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.LikeInboxComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLikeInboxCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **likeInboxCommentRequest** | [**LikeInboxCommentRequest**](LikeInboxCommentRequest.md) |  | 

### Return type

[**LikeInboxComment200Response**](LikeInboxComment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInboxComments

> ListInboxComments200Response ListInboxComments(ctx).ProfileId(profileId).Platform(platform).MinComments(minComments).Since(since).SortBy(sortBy).SortOrder(sortOrder).Limit(limit).Cursor(cursor).AccountId(accountId).Execute()

List commented posts



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
	profileId := "profileId_example" // string | Filter by profile ID (optional)
	platform := "platform_example" // string | Filter by platform. `metaads` is a synthetic value meaning the user's ads (boosted/dark posts) only; `facebook`/`instagram` return organic posts only. (optional)
	minComments := int32(56) // int32 | Minimum comment count (optional)
	since := time.Now() // time.Time | Posts created after this date (optional)
	sortBy := "sortBy_example" // string | Sort field (optional) (default to "date")
	sortOrder := "sortOrder_example" // string | Sort order (optional) (default to "desc")
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)
	accountId := "accountId_example" // string | Filter by specific social account ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.ListInboxComments(context.Background()).ProfileId(profileId).Platform(platform).MinComments(minComments).Since(since).SortBy(sortBy).SortOrder(sortOrder).Limit(limit).Cursor(cursor).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.ListInboxComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInboxComments`: ListInboxComments200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.ListInboxComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInboxCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Filter by profile ID | 
 **platform** | **string** | Filter by platform. &#x60;metaads&#x60; is a synthetic value meaning the user&#39;s ads (boosted/dark posts) only; &#x60;facebook&#x60;/&#x60;instagram&#x60; return organic posts only. | 
 **minComments** | **int32** | Minimum comment count | 
 **since** | **time.Time** | Posts created after this date | 
 **sortBy** | **string** | Sort field | [default to &quot;date&quot;]
 **sortOrder** | **string** | Sort order | [default to &quot;desc&quot;]
 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 
 **accountId** | **string** | Filter by specific social account ID | 

### Return type

[**ListInboxComments200Response**](ListInboxComments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplyToInboxPost

> ReplyToInboxPost200Response ReplyToInboxPost(ctx, postId).ReplyToInboxPostRequest(replyToInboxPostRequest).Execute()

Reply to comment



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
	postId := "postId_example" // string | Zernio post ID or platform-specific post ID. LinkedIn third-party posts accept full activity URN or numeric ID.
	replyToInboxPostRequest := *openapiclient.NewReplyToInboxPostRequest("AccountId_example", "Message_example") // ReplyToInboxPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.ReplyToInboxPost(context.Background(), postId).ReplyToInboxPostRequest(replyToInboxPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.ReplyToInboxPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplyToInboxPost`: ReplyToInboxPost200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.ReplyToInboxPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | Zernio post ID or platform-specific post ID. LinkedIn third-party posts accept full activity URN or numeric ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplyToInboxPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replyToInboxPostRequest** | [**ReplyToInboxPostRequest**](ReplyToInboxPostRequest.md) |  | 

### Return type

[**ReplyToInboxPost200Response**](ReplyToInboxPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendPrivateReplyToComment

> SendPrivateReplyToComment200Response SendPrivateReplyToComment(ctx, postId, commentId).SendPrivateReplyToCommentRequest(sendPrivateReplyToCommentRequest).Execute()

Send private reply



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
	postId := "postId_example" // string | The media/post ID (Instagram media ID or Facebook post ID)
	commentId := "commentId_example" // string | The comment ID to send a private reply to
	sendPrivateReplyToCommentRequest := *openapiclient.NewSendPrivateReplyToCommentRequest("AccountId_example", "Message_example") // SendPrivateReplyToCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.SendPrivateReplyToComment(context.Background(), postId, commentId).SendPrivateReplyToCommentRequest(sendPrivateReplyToCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.SendPrivateReplyToComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendPrivateReplyToComment`: SendPrivateReplyToComment200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.SendPrivateReplyToComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | The media/post ID (Instagram media ID or Facebook post ID) | 
**commentId** | **string** | The comment ID to send a private reply to | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendPrivateReplyToCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sendPrivateReplyToCommentRequest** | [**SendPrivateReplyToCommentRequest**](SendPrivateReplyToCommentRequest.md) |  | 

### Return type

[**SendPrivateReplyToComment200Response**](SendPrivateReplyToComment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnhideInboxComment

> HideInboxComment200Response UnhideInboxComment(ctx, postId, commentId).AccountId(accountId).Execute()

Unhide comment



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
	postId := "postId_example" // string | 
	commentId := "commentId_example" // string | 
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.UnhideInboxComment(context.Background(), postId, commentId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.UnhideInboxComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnhideInboxComment`: HideInboxComment200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.UnhideInboxComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnhideInboxCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **string** |  | 

### Return type

[**HideInboxComment200Response**](HideInboxComment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlikeInboxComment

> UnlikeInboxComment200Response UnlikeInboxComment(ctx, postId, commentId).AccountId(accountId).LikeUri(likeUri).Execute()

Unlike comment



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
	postId := "postId_example" // string | 
	commentId := "commentId_example" // string | 
	accountId := "accountId_example" // string | 
	likeUri := "likeUri_example" // string | (Bluesky only) The like URI returned when liking (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.UnlikeInboxComment(context.Background(), postId, commentId).AccountId(accountId).LikeUri(likeUri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.UnlikeInboxComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnlikeInboxComment`: UnlikeInboxComment200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.UnlikeInboxComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlikeInboxCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **string** |  | 
 **likeUri** | **string** | (Bluesky only) The like URI returned when liking | 

### Return type

[**UnlikeInboxComment200Response**](UnlikeInboxComment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

