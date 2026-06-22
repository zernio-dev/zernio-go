# \ReviewsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInboxReviewReply**](ReviewsAPI.md#DeleteInboxReviewReply) | **Delete** /v1/inbox/reviews/{reviewId}/reply | Delete review reply
[**ListInboxReviews**](ReviewsAPI.md#ListInboxReviews) | **Get** /v1/inbox/reviews | List reviews
[**ReplyToInboxReview**](ReviewsAPI.md#ReplyToInboxReview) | **Post** /v1/inbox/reviews/{reviewId}/reply | Reply to review



## DeleteInboxReviewReply

> DeleteInboxReviewReply200Response DeleteInboxReviewReply(ctx, reviewId).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()

Delete review reply



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
	reviewId := "reviewId_example" // string | 
	sendTypingIndicatorRequest := *openapiclient.NewSendTypingIndicatorRequest("AccountId_example") // SendTypingIndicatorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReviewsAPI.DeleteInboxReviewReply(context.Background(), reviewId).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.DeleteInboxReviewReply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInboxReviewReply`: DeleteInboxReviewReply200Response
	fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.DeleteInboxReviewReply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reviewId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInboxReviewReplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendTypingIndicatorRequest** | [**SendTypingIndicatorRequest**](SendTypingIndicatorRequest.md) |  | 

### Return type

[**DeleteInboxReviewReply200Response**](DeleteInboxReviewReply200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInboxReviews

> ListInboxReviews200Response ListInboxReviews(ctx).ProfileId(profileId).Platform(platform).MinRating(minRating).MaxRating(maxRating).Replied(replied).SortBy(sortBy).SortOrder(sortOrder).Limit(limit).Cursor(cursor).AccountId(accountId).Execute()

List reviews



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
	profileId := "profileId_example" // string |  (optional)
	platform := "platform_example" // string |  (optional)
	minRating := int32(56) // int32 |  (optional)
	maxRating := int32(56) // int32 |  (optional)
	replied := true // bool | Filter by reply status (optional)
	sortBy := "sortBy_example" // string |  (optional) (default to "date")
	sortOrder := "sortOrder_example" // string |  (optional) (default to "desc")
	limit := int32(56) // int32 |  (optional) (default to 25)
	cursor := "cursor_example" // string |  (optional)
	accountId := "accountId_example" // string | Filter by specific social account ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReviewsAPI.ListInboxReviews(context.Background()).ProfileId(profileId).Platform(platform).MinRating(minRating).MaxRating(maxRating).Replied(replied).SortBy(sortBy).SortOrder(sortOrder).Limit(limit).Cursor(cursor).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.ListInboxReviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInboxReviews`: ListInboxReviews200Response
	fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.ListInboxReviews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInboxReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** |  | 
 **platform** | **string** |  | 
 **minRating** | **int32** |  | 
 **maxRating** | **int32** |  | 
 **replied** | **bool** | Filter by reply status | 
 **sortBy** | **string** |  | [default to &quot;date&quot;]
 **sortOrder** | **string** |  | [default to &quot;desc&quot;]
 **limit** | **int32** |  | [default to 25]
 **cursor** | **string** |  | 
 **accountId** | **string** | Filter by specific social account ID | 

### Return type

[**ListInboxReviews200Response**](ListInboxReviews200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplyToInboxReview

> ReplyToInboxReview200Response ReplyToInboxReview(ctx, reviewId).ReplyToInboxReviewRequest(replyToInboxReviewRequest).Execute()

Reply to review



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
	reviewId := "reviewId_example" // string | Review ID (URL-encoded for Google Business)
	replyToInboxReviewRequest := *openapiclient.NewReplyToInboxReviewRequest("AccountId_example", "Message_example") // ReplyToInboxReviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReviewsAPI.ReplyToInboxReview(context.Background(), reviewId).ReplyToInboxReviewRequest(replyToInboxReviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.ReplyToInboxReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplyToInboxReview`: ReplyToInboxReview200Response
	fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.ReplyToInboxReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reviewId** | **string** | Review ID (URL-encoded for Google Business) | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplyToInboxReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replyToInboxReviewRequest** | [**ReplyToInboxReviewRequest**](ReplyToInboxReviewRequest.md) |  | 

### Return type

[**ReplyToInboxReview200Response**](ReplyToInboxReview200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

