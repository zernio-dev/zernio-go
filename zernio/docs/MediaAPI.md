# \MediaAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMediaPresignedUrl**](MediaAPI.md#GetMediaPresignedUrl) | **Post** /v1/media/presign | Get upload URL



## GetMediaPresignedUrl

> GetMediaPresignedUrl200Response GetMediaPresignedUrl(ctx).GetMediaPresignedUrlRequest(getMediaPresignedUrlRequest).Execute()

Get upload URL



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
	getMediaPresignedUrlRequest := *openapiclient.NewGetMediaPresignedUrlRequest("my-video.mp4", "video/mp4") // GetMediaPresignedUrlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.GetMediaPresignedUrl(context.Background()).GetMediaPresignedUrlRequest(getMediaPresignedUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.GetMediaPresignedUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaPresignedUrl`: GetMediaPresignedUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.GetMediaPresignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaPresignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getMediaPresignedUrlRequest** | [**GetMediaPresignedUrlRequest**](GetMediaPresignedUrlRequest.md) |  | 

### Return type

[**GetMediaPresignedUrl200Response**](GetMediaPresignedUrl200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

