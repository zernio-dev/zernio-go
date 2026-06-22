# \GMBServicesAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGoogleBusinessServices**](GMBServicesAPI.md#GetGoogleBusinessServices) | **Get** /v1/accounts/{accountId}/gmb-services | Get services
[**UpdateGoogleBusinessServices**](GMBServicesAPI.md#UpdateGoogleBusinessServices) | **Put** /v1/accounts/{accountId}/gmb-services | Replace services



## GetGoogleBusinessServices

> GetGoogleBusinessServices200Response GetGoogleBusinessServices(ctx, accountId).LocationId(locationId).Execute()

Get services



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
	locationId := "locationId_example" // string | Override which location to query. If omitted, uses the account's selected location. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBServicesAPI.GetGoogleBusinessServices(context.Background(), accountId).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBServicesAPI.GetGoogleBusinessServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoogleBusinessServices`: GetGoogleBusinessServices200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBServicesAPI.GetGoogleBusinessServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoogleBusinessServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationId** | **string** | Override which location to query. If omitted, uses the account&#39;s selected location. | 

### Return type

[**GetGoogleBusinessServices200Response**](GetGoogleBusinessServices200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGoogleBusinessServices

> UpdateGoogleBusinessServices200Response UpdateGoogleBusinessServices(ctx, accountId).UpdateGoogleBusinessServicesRequest(updateGoogleBusinessServicesRequest).LocationId(locationId).Execute()

Replace services



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
	updateGoogleBusinessServicesRequest := *openapiclient.NewUpdateGoogleBusinessServicesRequest([]openapiclient.GetGoogleBusinessServices200ResponseServicesInner{*openapiclient.NewGetGoogleBusinessServices200ResponseServicesInner()}) // UpdateGoogleBusinessServicesRequest | 
	locationId := "locationId_example" // string | Override which location to target. If omitted, uses the account's selected location. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBServicesAPI.UpdateGoogleBusinessServices(context.Background(), accountId).UpdateGoogleBusinessServicesRequest(updateGoogleBusinessServicesRequest).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBServicesAPI.UpdateGoogleBusinessServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGoogleBusinessServices`: UpdateGoogleBusinessServices200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBServicesAPI.UpdateGoogleBusinessServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGoogleBusinessServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGoogleBusinessServicesRequest** | [**UpdateGoogleBusinessServicesRequest**](UpdateGoogleBusinessServicesRequest.md) |  | 
 **locationId** | **string** | Override which location to target. If omitted, uses the account&#39;s selected location. | 

### Return type

[**UpdateGoogleBusinessServices200Response**](UpdateGoogleBusinessServices200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

