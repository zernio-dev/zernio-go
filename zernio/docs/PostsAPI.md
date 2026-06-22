# \PostsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUploadPosts**](PostsAPI.md#BulkUploadPosts) | **Post** /v1/posts/bulk-upload | Bulk upload from CSV
[**CreatePost**](PostsAPI.md#CreatePost) | **Post** /v1/posts | Create post
[**DeletePost**](PostsAPI.md#DeletePost) | **Delete** /v1/posts/{postId} | Delete post
[**EditPost**](PostsAPI.md#EditPost) | **Post** /v1/posts/{postId}/edit | Edit published post
[**GetPost**](PostsAPI.md#GetPost) | **Get** /v1/posts/{postId} | Get post
[**ListPosts**](PostsAPI.md#ListPosts) | **Get** /v1/posts | List posts
[**RetryPost**](PostsAPI.md#RetryPost) | **Post** /v1/posts/{postId}/retry | Retry failed post
[**UnpublishPost**](PostsAPI.md#UnpublishPost) | **Post** /v1/posts/{postId}/unpublish | Unpublish post
[**UpdatePost**](PostsAPI.md#UpdatePost) | **Put** /v1/posts/{postId} | Update post
[**UpdatePostMetadata**](PostsAPI.md#UpdatePostMetadata) | **Post** /v1/posts/{postId}/update-metadata | Update post metadata



## BulkUploadPosts

> BulkUploadResult BulkUploadPosts(ctx).DryRun(dryRun).File(file).Execute()

Bulk upload from CSV



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
	dryRun := true // bool |  (optional) (default to false)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.BulkUploadPosts(context.Background()).DryRun(dryRun).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.BulkUploadPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUploadPosts`: BulkUploadResult
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.BulkUploadPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUploadPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dryRun** | **bool** |  | [default to false]
 **file** | ***os.File** |  | 

### Return type

[**BulkUploadResult**](BulkUploadResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePost

> PostCreateResponse CreatePost(ctx).CreatePostRequest(createPostRequest).XRequestId(xRequestId).Execute()

Create post



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
	createPostRequest := *openapiclient.NewCreatePostRequest() // CreatePostRequest | 
	xRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional client-generated request identifier for safe retry (idempotency). When two requests carry the same value, the second is treated as a retry of the first and returns the original post (HTTP 200) instead of creating a duplicate. Window is ~5 minutes from the first request. Generate a UUID per logical call. SDKs do this automatically; HTTP clients should set it themselves or omit it. See the operation description for the full idempotency contract.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.CreatePost(context.Background()).CreatePostRequest(createPostRequest).XRequestId(xRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.CreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePost`: PostCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.CreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPostRequest** | [**CreatePostRequest**](CreatePostRequest.md) |  | 
 **xRequestId** | **string** | Optional client-generated request identifier for safe retry (idempotency). When two requests carry the same value, the second is treated as a retry of the first and returns the original post (HTTP 200) instead of creating a duplicate. Window is ~5 minutes from the first request. Generate a UUID per logical call. SDKs do this automatically; HTTP clients should set it themselves or omit it. See the operation description for the full idempotency contract.  | 

### Return type

[**PostCreateResponse**](PostCreateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePost

> PostDeleteResponse DeletePost(ctx, postId).Execute()

Delete post



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.DeletePost(context.Background(), postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.DeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePost`: PostDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.DeletePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostDeleteResponse**](PostDeleteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditPost

> EditPost200Response EditPost(ctx, postId).EditPostRequest(editPostRequest).Execute()

Edit published post



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
	editPostRequest := *openapiclient.NewEditPostRequest("Platform_example", "Content_example") // EditPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.EditPost(context.Background(), postId).EditPostRequest(editPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.EditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditPost`: EditPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.EditPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editPostRequest** | [**EditPostRequest**](EditPostRequest.md) |  | 

### Return type

[**EditPost200Response**](EditPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPost

> PostGetResponse GetPost(ctx, postId).Execute()

Get post



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.GetPost(context.Background(), postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.GetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPost`: PostGetResponse
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.GetPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostGetResponse**](PostGetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPosts

> PostsListResponse ListPosts(ctx).Page(page).Limit(limit).Status(status).Platform(platform).ProfileId(profileId).CreatedBy(createdBy).DateFrom(dateFrom).DateTo(dateTo).IncludeHidden(includeHidden).Search(search).SortBy(sortBy).AccountId(accountId).Execute()

List posts



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
	limit := int32(56) // int32 | Page size (optional) (default to 10)
	status := "status_example" // string |  (optional)
	platform := "twitter" // string |  (optional)
	profileId := "profileId_example" // string |  (optional)
	createdBy := "createdBy_example" // string |  (optional)
	dateFrom := time.Now() // string |  (optional)
	dateTo := time.Now() // string |  (optional)
	includeHidden := true // bool |  (optional) (default to false)
	search := "search_example" // string | Search posts by text content. (optional)
	sortBy := "sortBy_example" // string | Sort order for results. (optional) (default to "scheduled-desc")
	accountId := "accountId_example" // string | Filter posts to those published via a specific social account (24-char hex ObjectId). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.ListPosts(context.Background()).Page(page).Limit(limit).Status(status).Platform(platform).ProfileId(profileId).CreatedBy(createdBy).DateFrom(dateFrom).DateTo(dateTo).IncludeHidden(includeHidden).Search(search).SortBy(sortBy).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.ListPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPosts`: PostsListResponse
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.ListPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-based) | [default to 1]
 **limit** | **int32** | Page size | [default to 10]
 **status** | **string** |  | 
 **platform** | **string** |  | 
 **profileId** | **string** |  | 
 **createdBy** | **string** |  | 
 **dateFrom** | **string** |  | 
 **dateTo** | **string** |  | 
 **includeHidden** | **bool** |  | [default to false]
 **search** | **string** | Search posts by text content. | 
 **sortBy** | **string** | Sort order for results. | [default to &quot;scheduled-desc&quot;]
 **accountId** | **string** | Filter posts to those published via a specific social account (24-char hex ObjectId). | 

### Return type

[**PostsListResponse**](PostsListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryPost

> PostRetryResponse RetryPost(ctx, postId).Execute()

Retry failed post



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.RetryPost(context.Background(), postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.RetryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryPost`: PostRetryResponse
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.RetryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostRetryResponse**](PostRetryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpublishPost

> UnpublishPost200Response UnpublishPost(ctx, postId).UnpublishPostRequest(unpublishPostRequest).Execute()

Unpublish post



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
	unpublishPostRequest := *openapiclient.NewUnpublishPostRequest("Platform_example") // UnpublishPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.UnpublishPost(context.Background(), postId).UnpublishPostRequest(unpublishPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.UnpublishPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnpublishPost`: UnpublishPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.UnpublishPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unpublishPostRequest** | [**UnpublishPostRequest**](UnpublishPostRequest.md) |  | 

### Return type

[**UnpublishPost200Response**](UnpublishPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePost

> PostUpdateResponse UpdatePost(ctx, postId).UpdatePostRequest(updatePostRequest).Execute()

Update post



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
	updatePostRequest := *openapiclient.NewUpdatePostRequest() // UpdatePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.UpdatePost(context.Background(), postId).UpdatePostRequest(updatePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.UpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePost`: PostUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.UpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePostRequest** | [**UpdatePostRequest**](UpdatePostRequest.md) |  | 

### Return type

[**PostUpdateResponse**](PostUpdateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePostMetadata

> UpdatePostMetadata200Response UpdatePostMetadata(ctx, postId).UpdatePostMetadataRequest(updatePostMetadataRequest).Execute()

Update post metadata



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
	postId := "postId_example" // string | Zernio post ID, or \"_\" when using direct video ID mode
	updatePostMetadataRequest := *openapiclient.NewUpdatePostMetadataRequest("Platform_example") // UpdatePostMetadataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.UpdatePostMetadata(context.Background(), postId).UpdatePostMetadataRequest(updatePostMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.UpdatePostMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePostMetadata`: UpdatePostMetadata200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.UpdatePostMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postId** | **string** | Zernio post ID, or \&quot;_\&quot; when using direct video ID mode | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePostMetadataRequest** | [**UpdatePostMetadataRequest**](UpdatePostMetadataRequest.md) |  | 

### Return type

[**UpdatePostMetadata200Response**](UpdatePostMetadata200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

