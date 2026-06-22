# \GMBLocationDetailsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGoogleBusinessLocationDetails**](GMBLocationDetailsAPI.md#GetGoogleBusinessLocationDetails) | **Get** /v1/accounts/{accountId}/gmb-location-details | Get location details
[**UpdateGoogleBusinessLocationDetails**](GMBLocationDetailsAPI.md#UpdateGoogleBusinessLocationDetails) | **Put** /v1/accounts/{accountId}/gmb-location-details | Update location details



## GetGoogleBusinessLocationDetails

> GetGoogleBusinessLocationDetails200Response GetGoogleBusinessLocationDetails(ctx, accountId).LocationId(locationId).ReadMask(readMask).Execute()

Get location details



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
	readMask := "readMask_example" // string | Comma-separated fields to return. Available: name, title, phoneNumbers, categories, storefrontAddress, websiteUri, regularHours, specialHours, serviceArea, serviceItems, profile, openInfo, metadata, moreHours. `title` and `metadata` are always included in the response so the `location` summary block can be populated, even if you omit them here. Note: `location` is a derived response field, not a Google readMask value, passing it returns 400.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBLocationDetailsAPI.GetGoogleBusinessLocationDetails(context.Background(), accountId).LocationId(locationId).ReadMask(readMask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBLocationDetailsAPI.GetGoogleBusinessLocationDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoogleBusinessLocationDetails`: GetGoogleBusinessLocationDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBLocationDetailsAPI.GetGoogleBusinessLocationDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The Zernio account ID (from /v1/accounts) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoogleBusinessLocationDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationId** | **string** | Override which location to query. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 
 **readMask** | **string** | Comma-separated fields to return. Available: name, title, phoneNumbers, categories, storefrontAddress, websiteUri, regularHours, specialHours, serviceArea, serviceItems, profile, openInfo, metadata, moreHours. &#x60;title&#x60; and &#x60;metadata&#x60; are always included in the response so the &#x60;location&#x60; summary block can be populated, even if you omit them here. Note: &#x60;location&#x60; is a derived response field, not a Google readMask value, passing it returns 400.  | 

### Return type

[**GetGoogleBusinessLocationDetails200Response**](GetGoogleBusinessLocationDetails200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGoogleBusinessLocationDetails

> UpdateGoogleBusinessLocationDetails200Response UpdateGoogleBusinessLocationDetails(ctx, accountId).UpdateGoogleBusinessLocationDetailsRequest(updateGoogleBusinessLocationDetailsRequest).LocationId(locationId).Execute()

Update location details



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
	updateGoogleBusinessLocationDetailsRequest := *openapiclient.NewUpdateGoogleBusinessLocationDetailsRequest("UpdateMask_example") // UpdateGoogleBusinessLocationDetailsRequest | 
	locationId := "locationId_example" // string | Override which location to target. If omitted, uses the account's selected location. Use GET /gmb-locations to list valid IDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBLocationDetailsAPI.UpdateGoogleBusinessLocationDetails(context.Background(), accountId).UpdateGoogleBusinessLocationDetailsRequest(updateGoogleBusinessLocationDetailsRequest).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBLocationDetailsAPI.UpdateGoogleBusinessLocationDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGoogleBusinessLocationDetails`: UpdateGoogleBusinessLocationDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBLocationDetailsAPI.UpdateGoogleBusinessLocationDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The Zernio account ID (from /v1/accounts) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGoogleBusinessLocationDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGoogleBusinessLocationDetailsRequest** | [**UpdateGoogleBusinessLocationDetailsRequest**](UpdateGoogleBusinessLocationDetailsRequest.md) |  | 
 **locationId** | **string** | Override which location to target. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 

### Return type

[**UpdateGoogleBusinessLocationDetails200Response**](UpdateGoogleBusinessLocationDetails200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

