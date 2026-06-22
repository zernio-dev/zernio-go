# \GMBMediaAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGoogleBusinessMedia**](GMBMediaAPI.md#CreateGoogleBusinessMedia) | **Post** /v1/accounts/{accountId}/gmb-media | Upload photo
[**DeleteGoogleBusinessMedia**](GMBMediaAPI.md#DeleteGoogleBusinessMedia) | **Delete** /v1/accounts/{accountId}/gmb-media | Delete photo
[**ListGoogleBusinessMedia**](GMBMediaAPI.md#ListGoogleBusinessMedia) | **Get** /v1/accounts/{accountId}/gmb-media | List media



## CreateGoogleBusinessMedia

> CreateGoogleBusinessMedia200Response CreateGoogleBusinessMedia(ctx, accountId).CreateGoogleBusinessMediaRequest(createGoogleBusinessMediaRequest).LocationId(locationId).Execute()

Upload photo



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
	createGoogleBusinessMediaRequest := *openapiclient.NewCreateGoogleBusinessMediaRequest("SourceUrl_example") // CreateGoogleBusinessMediaRequest | 
	locationId := "locationId_example" // string | Override which location to target. If omitted, uses the account's selected location. Use GET /gmb-locations to list valid IDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBMediaAPI.CreateGoogleBusinessMedia(context.Background(), accountId).CreateGoogleBusinessMediaRequest(createGoogleBusinessMediaRequest).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBMediaAPI.CreateGoogleBusinessMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGoogleBusinessMedia`: CreateGoogleBusinessMedia200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBMediaAPI.CreateGoogleBusinessMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGoogleBusinessMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGoogleBusinessMediaRequest** | [**CreateGoogleBusinessMediaRequest**](CreateGoogleBusinessMediaRequest.md) |  | 
 **locationId** | **string** | Override which location to target. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 

### Return type

[**CreateGoogleBusinessMedia200Response**](CreateGoogleBusinessMedia200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGoogleBusinessMedia

> DeleteGoogleBusinessMedia200Response DeleteGoogleBusinessMedia(ctx, accountId).MediaId(mediaId).LocationId(locationId).Execute()

Delete photo



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
	mediaId := "mediaId_example" // string | The media item ID to delete
	locationId := "locationId_example" // string | Override which location to target. If omitted, uses the account's selected location. Use GET /gmb-locations to list valid IDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBMediaAPI.DeleteGoogleBusinessMedia(context.Background(), accountId).MediaId(mediaId).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBMediaAPI.DeleteGoogleBusinessMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGoogleBusinessMedia`: DeleteGoogleBusinessMedia200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBMediaAPI.DeleteGoogleBusinessMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGoogleBusinessMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaId** | **string** | The media item ID to delete | 
 **locationId** | **string** | Override which location to target. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 

### Return type

[**DeleteGoogleBusinessMedia200Response**](DeleteGoogleBusinessMedia200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGoogleBusinessMedia

> ListGoogleBusinessMedia200Response ListGoogleBusinessMedia(ctx, accountId).LocationId(locationId).PageSize(pageSize).PageToken(pageToken).Execute()

List media



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
	locationId := "locationId_example" // string | Override which location to query. If omitted, uses the account's selected location. Use GET /gmb-locations to list valid IDs. (optional)
	pageSize := int32(56) // int32 | Number of items to return (max 100) (optional) (default to 100)
	pageToken := "pageToken_example" // string | Pagination token from previous response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBMediaAPI.ListGoogleBusinessMedia(context.Background(), accountId).LocationId(locationId).PageSize(pageSize).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBMediaAPI.ListGoogleBusinessMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGoogleBusinessMedia`: ListGoogleBusinessMedia200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBMediaAPI.ListGoogleBusinessMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGoogleBusinessMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationId** | **string** | Override which location to query. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 
 **pageSize** | **int32** | Number of items to return (max 100) | [default to 100]
 **pageToken** | **string** | Pagination token from previous response | 

### Return type

[**ListGoogleBusinessMedia200Response**](ListGoogleBusinessMedia200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

