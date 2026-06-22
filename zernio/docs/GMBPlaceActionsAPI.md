# \GMBPlaceActionsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGoogleBusinessPlaceAction**](GMBPlaceActionsAPI.md#CreateGoogleBusinessPlaceAction) | **Post** /v1/accounts/{accountId}/gmb-place-actions | Create action link
[**DeleteGoogleBusinessPlaceAction**](GMBPlaceActionsAPI.md#DeleteGoogleBusinessPlaceAction) | **Delete** /v1/accounts/{accountId}/gmb-place-actions | Delete action link
[**ListGoogleBusinessPlaceActions**](GMBPlaceActionsAPI.md#ListGoogleBusinessPlaceActions) | **Get** /v1/accounts/{accountId}/gmb-place-actions | List action links
[**UpdateGoogleBusinessPlaceAction**](GMBPlaceActionsAPI.md#UpdateGoogleBusinessPlaceAction) | **Patch** /v1/accounts/{accountId}/gmb-place-actions | Update action link



## CreateGoogleBusinessPlaceAction

> CreateGoogleBusinessPlaceAction200Response CreateGoogleBusinessPlaceAction(ctx, accountId).CreateGoogleBusinessPlaceActionRequest(createGoogleBusinessPlaceActionRequest).LocationId(locationId).Execute()

Create action link



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
	createGoogleBusinessPlaceActionRequest := *openapiclient.NewCreateGoogleBusinessPlaceActionRequest("Uri_example", "PlaceActionType_example") // CreateGoogleBusinessPlaceActionRequest | 
	locationId := "locationId_example" // string | Override which location to target. If omitted, uses the account's selected location. Use GET /gmb-locations to list valid IDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBPlaceActionsAPI.CreateGoogleBusinessPlaceAction(context.Background(), accountId).CreateGoogleBusinessPlaceActionRequest(createGoogleBusinessPlaceActionRequest).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBPlaceActionsAPI.CreateGoogleBusinessPlaceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGoogleBusinessPlaceAction`: CreateGoogleBusinessPlaceAction200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBPlaceActionsAPI.CreateGoogleBusinessPlaceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGoogleBusinessPlaceActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGoogleBusinessPlaceActionRequest** | [**CreateGoogleBusinessPlaceActionRequest**](CreateGoogleBusinessPlaceActionRequest.md) |  | 
 **locationId** | **string** | Override which location to target. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 

### Return type

[**CreateGoogleBusinessPlaceAction200Response**](CreateGoogleBusinessPlaceAction200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGoogleBusinessPlaceAction

> DeleteGoogleBusinessPlaceAction200Response DeleteGoogleBusinessPlaceAction(ctx, accountId).Name(name).LocationId(locationId).Execute()

Delete action link



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
	name := "name_example" // string | The resource name of the place action link (e.g. locations/123/placeActionLinks/456)
	locationId := "locationId_example" // string | Override which location to target. If omitted, uses the account's selected location. Use GET /gmb-locations to list valid IDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBPlaceActionsAPI.DeleteGoogleBusinessPlaceAction(context.Background(), accountId).Name(name).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBPlaceActionsAPI.DeleteGoogleBusinessPlaceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGoogleBusinessPlaceAction`: DeleteGoogleBusinessPlaceAction200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBPlaceActionsAPI.DeleteGoogleBusinessPlaceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGoogleBusinessPlaceActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The resource name of the place action link (e.g. locations/123/placeActionLinks/456) | 
 **locationId** | **string** | Override which location to target. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 

### Return type

[**DeleteGoogleBusinessPlaceAction200Response**](DeleteGoogleBusinessPlaceAction200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGoogleBusinessPlaceActions

> ListGoogleBusinessPlaceActions200Response ListGoogleBusinessPlaceActions(ctx, accountId).LocationId(locationId).PageSize(pageSize).PageToken(pageToken).Execute()

List action links



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
	pageSize := int32(56) // int32 |  (optional) (default to 100)
	pageToken := "pageToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBPlaceActionsAPI.ListGoogleBusinessPlaceActions(context.Background(), accountId).LocationId(locationId).PageSize(pageSize).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBPlaceActionsAPI.ListGoogleBusinessPlaceActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGoogleBusinessPlaceActions`: ListGoogleBusinessPlaceActions200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBPlaceActionsAPI.ListGoogleBusinessPlaceActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGoogleBusinessPlaceActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationId** | **string** | Override which location to query. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 
 **pageSize** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**ListGoogleBusinessPlaceActions200Response**](ListGoogleBusinessPlaceActions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGoogleBusinessPlaceAction

> CreateGoogleBusinessPlaceAction200Response UpdateGoogleBusinessPlaceAction(ctx, accountId).UpdateGoogleBusinessPlaceActionRequest(updateGoogleBusinessPlaceActionRequest).LocationId(locationId).Execute()

Update action link



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
	updateGoogleBusinessPlaceActionRequest := *openapiclient.NewUpdateGoogleBusinessPlaceActionRequest("Name_example") // UpdateGoogleBusinessPlaceActionRequest | 
	locationId := "locationId_example" // string | Override which location to target. If omitted, uses the account's selected location. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBPlaceActionsAPI.UpdateGoogleBusinessPlaceAction(context.Background(), accountId).UpdateGoogleBusinessPlaceActionRequest(updateGoogleBusinessPlaceActionRequest).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBPlaceActionsAPI.UpdateGoogleBusinessPlaceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGoogleBusinessPlaceAction`: CreateGoogleBusinessPlaceAction200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBPlaceActionsAPI.UpdateGoogleBusinessPlaceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGoogleBusinessPlaceActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGoogleBusinessPlaceActionRequest** | [**UpdateGoogleBusinessPlaceActionRequest**](UpdateGoogleBusinessPlaceActionRequest.md) |  | 
 **locationId** | **string** | Override which location to target. If omitted, uses the account&#39;s selected location. | 

### Return type

[**CreateGoogleBusinessPlaceAction200Response**](CreateGoogleBusinessPlaceAction200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

