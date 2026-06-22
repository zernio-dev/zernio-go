# \GMBReviewsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchGetGoogleBusinessReviews**](GMBReviewsAPI.md#BatchGetGoogleBusinessReviews) | **Post** /v1/accounts/{accountId}/gmb-reviews/batch | Batch get reviews
[**DeleteGoogleBusinessReviewReply**](GMBReviewsAPI.md#DeleteGoogleBusinessReviewReply) | **Delete** /v1/accounts/{accountId}/gmb-reviews/{reviewId}/reply | Delete a review reply
[**GetGoogleBusinessReviews**](GMBReviewsAPI.md#GetGoogleBusinessReviews) | **Get** /v1/accounts/{accountId}/gmb-reviews | Get reviews
[**ReplyToGoogleBusinessReview**](GMBReviewsAPI.md#ReplyToGoogleBusinessReview) | **Post** /v1/accounts/{accountId}/gmb-reviews/{reviewId}/reply | Reply to a review



## BatchGetGoogleBusinessReviews

> BatchGetGoogleBusinessReviews200Response BatchGetGoogleBusinessReviews(ctx, accountId).BatchGetGoogleBusinessReviewsRequest(batchGetGoogleBusinessReviewsRequest).Execute()

Batch get reviews



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
	batchGetGoogleBusinessReviewsRequest := *openapiclient.NewBatchGetGoogleBusinessReviewsRequest([]string{"LocationNames_example"}) // BatchGetGoogleBusinessReviewsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBReviewsAPI.BatchGetGoogleBusinessReviews(context.Background(), accountId).BatchGetGoogleBusinessReviewsRequest(batchGetGoogleBusinessReviewsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBReviewsAPI.BatchGetGoogleBusinessReviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchGetGoogleBusinessReviews`: BatchGetGoogleBusinessReviews200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBReviewsAPI.BatchGetGoogleBusinessReviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchGetGoogleBusinessReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchGetGoogleBusinessReviewsRequest** | [**BatchGetGoogleBusinessReviewsRequest**](BatchGetGoogleBusinessReviewsRequest.md) |  | 

### Return type

[**BatchGetGoogleBusinessReviews200Response**](BatchGetGoogleBusinessReviews200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGoogleBusinessReviewReply

> DeleteGoogleBusinessReviewReply200Response DeleteGoogleBusinessReviewReply(ctx, accountId, reviewId).Execute()

Delete a review reply



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
	accountId := "accountId_example" // string | The Zernio account ID (from /v1/accounts)
	reviewId := "reviewId_example" // string | The review ID portion (e.g. \"AIe9_BGx1234567890\"), not the full resource name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBReviewsAPI.DeleteGoogleBusinessReviewReply(context.Background(), accountId, reviewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBReviewsAPI.DeleteGoogleBusinessReviewReply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGoogleBusinessReviewReply`: DeleteGoogleBusinessReviewReply200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBReviewsAPI.DeleteGoogleBusinessReviewReply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The Zernio account ID (from /v1/accounts) | 
**reviewId** | **string** | The review ID portion (e.g. \&quot;AIe9_BGx1234567890\&quot;), not the full resource name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGoogleBusinessReviewReplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteGoogleBusinessReviewReply200Response**](DeleteGoogleBusinessReviewReply200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGoogleBusinessReviews

> GetGoogleBusinessReviews200Response GetGoogleBusinessReviews(ctx, accountId).LocationId(locationId).PageSize(pageSize).PageToken(pageToken).Execute()

Get reviews



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
	accountId := "accountId_example" // string | The Zernio account ID (from /v1/accounts)
	locationId := "locationId_example" // string | Override which location to query. If omitted, uses the account's selected location. Use GET /gmb-locations to list valid IDs. (optional)
	pageSize := int32(56) // int32 | Number of reviews to fetch per page (max 50) (optional) (default to 50)
	pageToken := "pageToken_example" // string | Pagination token from previous response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBReviewsAPI.GetGoogleBusinessReviews(context.Background(), accountId).LocationId(locationId).PageSize(pageSize).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBReviewsAPI.GetGoogleBusinessReviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoogleBusinessReviews`: GetGoogleBusinessReviews200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBReviewsAPI.GetGoogleBusinessReviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The Zernio account ID (from /v1/accounts) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoogleBusinessReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationId** | **string** | Override which location to query. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 
 **pageSize** | **int32** | Number of reviews to fetch per page (max 50) | [default to 50]
 **pageToken** | **string** | Pagination token from previous response | 

### Return type

[**GetGoogleBusinessReviews200Response**](GetGoogleBusinessReviews200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplyToGoogleBusinessReview

> ReplyToGoogleBusinessReview200Response ReplyToGoogleBusinessReview(ctx, accountId, reviewId).ReplyToGoogleBusinessReviewRequest(replyToGoogleBusinessReviewRequest).Execute()

Reply to a review



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
	accountId := "accountId_example" // string | The Zernio account ID (from /v1/accounts)
	reviewId := "reviewId_example" // string | The review ID portion (e.g. \"AIe9_BGx1234567890\"), not the full resource name
	replyToGoogleBusinessReviewRequest := *openapiclient.NewReplyToGoogleBusinessReviewRequest("Comment_example") // ReplyToGoogleBusinessReviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBReviewsAPI.ReplyToGoogleBusinessReview(context.Background(), accountId, reviewId).ReplyToGoogleBusinessReviewRequest(replyToGoogleBusinessReviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBReviewsAPI.ReplyToGoogleBusinessReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplyToGoogleBusinessReview`: ReplyToGoogleBusinessReview200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBReviewsAPI.ReplyToGoogleBusinessReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The Zernio account ID (from /v1/accounts) | 
**reviewId** | **string** | The review ID portion (e.g. \&quot;AIe9_BGx1234567890\&quot;), not the full resource name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplyToGoogleBusinessReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replyToGoogleBusinessReviewRequest** | [**ReplyToGoogleBusinessReviewRequest**](ReplyToGoogleBusinessReviewRequest.md) |  | 

### Return type

[**ReplyToGoogleBusinessReview200Response**](ReplyToGoogleBusinessReview200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

