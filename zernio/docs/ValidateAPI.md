# \ValidateAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateMedia**](ValidateAPI.md#ValidateMedia) | **Post** /v1/tools/validate/media | Validate media URL
[**ValidatePost**](ValidateAPI.md#ValidatePost) | **Post** /v1/tools/validate/post | Validate post content
[**ValidatePostLength**](ValidateAPI.md#ValidatePostLength) | **Post** /v1/tools/validate/post-length | Validate character count
[**ValidateSubreddit**](ValidateAPI.md#ValidateSubreddit) | **Get** /v1/tools/validate/subreddit | Check subreddit existence



## ValidateMedia

> ValidateMedia200Response ValidateMedia(ctx).ValidateMediaRequest(validateMediaRequest).Execute()

Validate media URL



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
	validateMediaRequest := *openapiclient.NewValidateMediaRequest("https://example.com/image.jpg") // ValidateMediaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidateAPI.ValidateMedia(context.Background()).ValidateMediaRequest(validateMediaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidateAPI.ValidateMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateMedia`: ValidateMedia200Response
	fmt.Fprintf(os.Stdout, "Response from `ValidateAPI.ValidateMedia`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateMediaRequest** | [**ValidateMediaRequest**](ValidateMediaRequest.md) |  | 

### Return type

[**ValidateMedia200Response**](ValidateMedia200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePost

> ValidatePost200Response ValidatePost(ctx).ValidatePostRequest(validatePostRequest).Execute()

Validate post content



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
	validatePostRequest := *openapiclient.NewValidatePostRequest([]openapiclient.ValidatePostRequestPlatformsInner{*openapiclient.NewValidatePostRequestPlatformsInner("Platform_example")}) // ValidatePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidateAPI.ValidatePost(context.Background()).ValidatePostRequest(validatePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidateAPI.ValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidatePost`: ValidatePost200Response
	fmt.Fprintf(os.Stdout, "Response from `ValidateAPI.ValidatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validatePostRequest** | [**ValidatePostRequest**](ValidatePostRequest.md) |  | 

### Return type

[**ValidatePost200Response**](ValidatePost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePostLength

> ValidatePostLength200Response ValidatePostLength(ctx).ValidatePostLengthRequest(validatePostLengthRequest).Execute()

Validate character count



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
	validatePostLengthRequest := *openapiclient.NewValidatePostLengthRequest("Check out https://zernio.com for scheduling posts!") // ValidatePostLengthRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidateAPI.ValidatePostLength(context.Background()).ValidatePostLengthRequest(validatePostLengthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidateAPI.ValidatePostLength``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidatePostLength`: ValidatePostLength200Response
	fmt.Fprintf(os.Stdout, "Response from `ValidateAPI.ValidatePostLength`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidatePostLengthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validatePostLengthRequest** | [**ValidatePostLengthRequest**](ValidatePostLengthRequest.md) |  | 

### Return type

[**ValidatePostLength200Response**](ValidatePostLength200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSubreddit

> ValidateSubreddit200Response ValidateSubreddit(ctx).Name(name).AccountId(accountId).Execute()

Check subreddit existence



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
	name := "programming" // string | Subreddit name (with or without \"r/\" prefix)
	accountId := "accountId_example" // string | Reddit social account ID for authenticated lookup (recommended for reliable results) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidateAPI.ValidateSubreddit(context.Background()).Name(name).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidateAPI.ValidateSubreddit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateSubreddit`: ValidateSubreddit200Response
	fmt.Fprintf(os.Stdout, "Response from `ValidateAPI.ValidateSubreddit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateSubredditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Subreddit name (with or without \&quot;r/\&quot; prefix) | 
 **accountId** | **string** | Reddit social account ID for authenticated lookup (recommended for reliable results) | 

### Return type

[**ValidateSubreddit200Response**](ValidateSubreddit200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

