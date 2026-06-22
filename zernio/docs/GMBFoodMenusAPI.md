# \GMBFoodMenusAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGoogleBusinessFoodMenus**](GMBFoodMenusAPI.md#GetGoogleBusinessFoodMenus) | **Get** /v1/accounts/{accountId}/gmb-food-menus | Get food menus
[**UpdateGoogleBusinessFoodMenus**](GMBFoodMenusAPI.md#UpdateGoogleBusinessFoodMenus) | **Put** /v1/accounts/{accountId}/gmb-food-menus | Update food menus



## GetGoogleBusinessFoodMenus

> GetGoogleBusinessFoodMenus200Response GetGoogleBusinessFoodMenus(ctx, accountId).LocationId(locationId).Execute()

Get food menus



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBFoodMenusAPI.GetGoogleBusinessFoodMenus(context.Background(), accountId).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBFoodMenusAPI.GetGoogleBusinessFoodMenus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoogleBusinessFoodMenus`: GetGoogleBusinessFoodMenus200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBFoodMenusAPI.GetGoogleBusinessFoodMenus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The Zernio account ID (from /v1/accounts) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoogleBusinessFoodMenusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationId** | **string** | Override which location to query. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 

### Return type

[**GetGoogleBusinessFoodMenus200Response**](GetGoogleBusinessFoodMenus200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGoogleBusinessFoodMenus

> GetGoogleBusinessFoodMenus200Response UpdateGoogleBusinessFoodMenus(ctx, accountId).UpdateGoogleBusinessFoodMenusRequest(updateGoogleBusinessFoodMenusRequest).LocationId(locationId).Execute()

Update food menus



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
	updateGoogleBusinessFoodMenusRequest := *openapiclient.NewUpdateGoogleBusinessFoodMenusRequest([]openapiclient.FoodMenu{*openapiclient.NewFoodMenu([]openapiclient.FoodMenuLabel{*openapiclient.NewFoodMenuLabel("DisplayName_example")})}) // UpdateGoogleBusinessFoodMenusRequest | 
	locationId := "locationId_example" // string | Override which location to target. If omitted, uses the account's selected location. Use GET /gmb-locations to list valid IDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBFoodMenusAPI.UpdateGoogleBusinessFoodMenus(context.Background(), accountId).UpdateGoogleBusinessFoodMenusRequest(updateGoogleBusinessFoodMenusRequest).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBFoodMenusAPI.UpdateGoogleBusinessFoodMenus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGoogleBusinessFoodMenus`: GetGoogleBusinessFoodMenus200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBFoodMenusAPI.UpdateGoogleBusinessFoodMenus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The Zernio account ID (from /v1/accounts) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGoogleBusinessFoodMenusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGoogleBusinessFoodMenusRequest** | [**UpdateGoogleBusinessFoodMenusRequest**](UpdateGoogleBusinessFoodMenusRequest.md) |  | 
 **locationId** | **string** | Override which location to target. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 

### Return type

[**GetGoogleBusinessFoodMenus200Response**](GetGoogleBusinessFoodMenus200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

