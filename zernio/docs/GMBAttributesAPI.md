# \GMBAttributesAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGmbAttributeMetadata**](GMBAttributesAPI.md#GetGmbAttributeMetadata) | **Get** /v1/accounts/{accountId}/gmb-attribute-metadata | Get attribute metadata
[**GetGoogleBusinessAttributes**](GMBAttributesAPI.md#GetGoogleBusinessAttributes) | **Get** /v1/accounts/{accountId}/gmb-attributes | Get attributes
[**UpdateGoogleBusinessAttributes**](GMBAttributesAPI.md#UpdateGoogleBusinessAttributes) | **Put** /v1/accounts/{accountId}/gmb-attributes | Update attributes



## GetGmbAttributeMetadata

> GetGmbAttributeMetadata200Response GetGmbAttributeMetadata(ctx, accountId).LocationId(locationId).CategoryName(categoryName).RegionCode(regionCode).LanguageCode(languageCode).PageSize(pageSize).PageToken(pageToken).Execute()

Get attribute metadata



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
	locationId := "locationId_example" // string | GBP location ID (e.g. \"6257659026299438786\"). If omitted, uses the account's stored selectedLocationId. Mutually exclusive with categoryName.  (optional)
	categoryName := "categoryName_example" // string | Category resource name, must start with \"categories/\" (e.g. \"categories/gcid:plumber\"). Required together with regionCode. Mutually exclusive with locationId.  (optional)
	regionCode := "regionCode_example" // string | BCP-47 region code (e.g. \"US\", \"ES\"). Required when categoryName is provided.  (optional)
	languageCode := "languageCode_example" // string | BCP-47 language code for display names (e.g. \"en\", \"es\"). Optional when categoryName is provided. Omitted from the Google call when not supplied.  (optional)
	pageSize := int32(56) // int32 | Maximum number of attribute metadata items to return. Google defaults to 200. (optional)
	pageToken := "pageToken_example" // string | Pagination token from a previous response's nextPageToken field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBAttributesAPI.GetGmbAttributeMetadata(context.Background(), accountId).LocationId(locationId).CategoryName(categoryName).RegionCode(regionCode).LanguageCode(languageCode).PageSize(pageSize).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBAttributesAPI.GetGmbAttributeMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGmbAttributeMetadata`: GetGmbAttributeMetadata200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBAttributesAPI.GetGmbAttributeMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGmbAttributeMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationId** | **string** | GBP location ID (e.g. \&quot;6257659026299438786\&quot;). If omitted, uses the account&#39;s stored selectedLocationId. Mutually exclusive with categoryName.  | 
 **categoryName** | **string** | Category resource name, must start with \&quot;categories/\&quot; (e.g. \&quot;categories/gcid:plumber\&quot;). Required together with regionCode. Mutually exclusive with locationId.  | 
 **regionCode** | **string** | BCP-47 region code (e.g. \&quot;US\&quot;, \&quot;ES\&quot;). Required when categoryName is provided.  | 
 **languageCode** | **string** | BCP-47 language code for display names (e.g. \&quot;en\&quot;, \&quot;es\&quot;). Optional when categoryName is provided. Omitted from the Google call when not supplied.  | 
 **pageSize** | **int32** | Maximum number of attribute metadata items to return. Google defaults to 200. | 
 **pageToken** | **string** | Pagination token from a previous response&#39;s nextPageToken field. | 

### Return type

[**GetGmbAttributeMetadata200Response**](GetGmbAttributeMetadata200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGoogleBusinessAttributes

> GetGoogleBusinessAttributes200Response GetGoogleBusinessAttributes(ctx, accountId).LocationId(locationId).Execute()

Get attributes



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBAttributesAPI.GetGoogleBusinessAttributes(context.Background(), accountId).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBAttributesAPI.GetGoogleBusinessAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoogleBusinessAttributes`: GetGoogleBusinessAttributes200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBAttributesAPI.GetGoogleBusinessAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoogleBusinessAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationId** | **string** | Override which location to query. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 

### Return type

[**GetGoogleBusinessAttributes200Response**](GetGoogleBusinessAttributes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGoogleBusinessAttributes

> UpdateGoogleBusinessAttributes200Response UpdateGoogleBusinessAttributes(ctx, accountId).UpdateGoogleBusinessAttributesRequest(updateGoogleBusinessAttributesRequest).LocationId(locationId).Execute()

Update attributes



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
	updateGoogleBusinessAttributesRequest := *openapiclient.NewUpdateGoogleBusinessAttributesRequest([]openapiclient.UpdateGoogleBusinessAttributesRequestAttributesInner{*openapiclient.NewUpdateGoogleBusinessAttributesRequestAttributesInner()}, "AttributeMask_example") // UpdateGoogleBusinessAttributesRequest | 
	locationId := "locationId_example" // string | Override which location to target. If omitted, uses the account's selected location. Use GET /gmb-locations to list valid IDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBAttributesAPI.UpdateGoogleBusinessAttributes(context.Background(), accountId).UpdateGoogleBusinessAttributesRequest(updateGoogleBusinessAttributesRequest).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBAttributesAPI.UpdateGoogleBusinessAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGoogleBusinessAttributes`: UpdateGoogleBusinessAttributes200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBAttributesAPI.UpdateGoogleBusinessAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGoogleBusinessAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGoogleBusinessAttributesRequest** | [**UpdateGoogleBusinessAttributesRequest**](UpdateGoogleBusinessAttributesRequest.md) |  | 
 **locationId** | **string** | Override which location to target. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 

### Return type

[**UpdateGoogleBusinessAttributes200Response**](UpdateGoogleBusinessAttributes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

