# \GMBVerificationsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteGoogleBusinessVerification**](GMBVerificationsAPI.md#CompleteGoogleBusinessVerification) | **Post** /v1/accounts/{accountId}/gmb-verifications/{verificationId}/complete | Complete a verification
[**FetchGoogleBusinessVerificationOptions**](GMBVerificationsAPI.md#FetchGoogleBusinessVerificationOptions) | **Post** /v1/accounts/{accountId}/gmb-verifications/options | Fetch verification options
[**GetGoogleBusinessVerifications**](GMBVerificationsAPI.md#GetGoogleBusinessVerifications) | **Get** /v1/accounts/{accountId}/gmb-verifications | Get verification state
[**StartGoogleBusinessVerification**](GMBVerificationsAPI.md#StartGoogleBusinessVerification) | **Post** /v1/accounts/{accountId}/gmb-verifications | Start a verification



## CompleteGoogleBusinessVerification

> StartGoogleBusinessVerification200Response CompleteGoogleBusinessVerification(ctx, accountId, verificationId).CompleteGoogleBusinessVerificationRequest(completeGoogleBusinessVerificationRequest).LocationId(locationId).Execute()

Complete a verification



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
	verificationId := "verificationId_example" // string | The last segment of a verification `name` from GET /gmb-verifications.
	completeGoogleBusinessVerificationRequest := *openapiclient.NewCompleteGoogleBusinessVerificationRequest("Pin_example") // CompleteGoogleBusinessVerificationRequest | 
	locationId := "locationId_example" // string | Override which location to target. If omitted, uses the account's selected location. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBVerificationsAPI.CompleteGoogleBusinessVerification(context.Background(), accountId, verificationId).CompleteGoogleBusinessVerificationRequest(completeGoogleBusinessVerificationRequest).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBVerificationsAPI.CompleteGoogleBusinessVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteGoogleBusinessVerification`: StartGoogleBusinessVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBVerificationsAPI.CompleteGoogleBusinessVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The Zernio account ID (from /v1/accounts) | 
**verificationId** | **string** | The last segment of a verification &#x60;name&#x60; from GET /gmb-verifications. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteGoogleBusinessVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **completeGoogleBusinessVerificationRequest** | [**CompleteGoogleBusinessVerificationRequest**](CompleteGoogleBusinessVerificationRequest.md) |  | 
 **locationId** | **string** | Override which location to target. If omitted, uses the account&#39;s selected location. | 

### Return type

[**StartGoogleBusinessVerification200Response**](StartGoogleBusinessVerification200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchGoogleBusinessVerificationOptions

> FetchGoogleBusinessVerificationOptions200Response FetchGoogleBusinessVerificationOptions(ctx, accountId).FetchGoogleBusinessVerificationOptionsRequest(fetchGoogleBusinessVerificationOptionsRequest).LocationId(locationId).Execute()

Fetch verification options



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
	fetchGoogleBusinessVerificationOptionsRequest := *openapiclient.NewFetchGoogleBusinessVerificationOptionsRequest("en-US") // FetchGoogleBusinessVerificationOptionsRequest | 
	locationId := "locationId_example" // string | Override which location to query. If omitted, uses the account's selected location. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBVerificationsAPI.FetchGoogleBusinessVerificationOptions(context.Background(), accountId).FetchGoogleBusinessVerificationOptionsRequest(fetchGoogleBusinessVerificationOptionsRequest).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBVerificationsAPI.FetchGoogleBusinessVerificationOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchGoogleBusinessVerificationOptions`: FetchGoogleBusinessVerificationOptions200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBVerificationsAPI.FetchGoogleBusinessVerificationOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The Zernio account ID (from /v1/accounts) | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchGoogleBusinessVerificationOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchGoogleBusinessVerificationOptionsRequest** | [**FetchGoogleBusinessVerificationOptionsRequest**](FetchGoogleBusinessVerificationOptionsRequest.md) |  | 
 **locationId** | **string** | Override which location to query. If omitted, uses the account&#39;s selected location. | 

### Return type

[**FetchGoogleBusinessVerificationOptions200Response**](FetchGoogleBusinessVerificationOptions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGoogleBusinessVerifications

> GetGoogleBusinessVerifications200Response GetGoogleBusinessVerifications(ctx, accountId).LocationId(locationId).Execute()

Get verification state



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
	resp, r, err := apiClient.GMBVerificationsAPI.GetGoogleBusinessVerifications(context.Background(), accountId).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBVerificationsAPI.GetGoogleBusinessVerifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoogleBusinessVerifications`: GetGoogleBusinessVerifications200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBVerificationsAPI.GetGoogleBusinessVerifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The Zernio account ID (from /v1/accounts) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoogleBusinessVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationId** | **string** | Override which location to query. If omitted, uses the account&#39;s selected location. Use GET /gmb-locations to list valid IDs. | 

### Return type

[**GetGoogleBusinessVerifications200Response**](GetGoogleBusinessVerifications200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartGoogleBusinessVerification

> StartGoogleBusinessVerification200Response StartGoogleBusinessVerification(ctx, accountId).StartGoogleBusinessVerificationRequest(startGoogleBusinessVerificationRequest).LocationId(locationId).Execute()

Start a verification



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
	startGoogleBusinessVerificationRequest := *openapiclient.NewStartGoogleBusinessVerificationRequest("Method_example") // StartGoogleBusinessVerificationRequest | 
	locationId := "locationId_example" // string | Override which location to target. If omitted, uses the account's selected location. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GMBVerificationsAPI.StartGoogleBusinessVerification(context.Background(), accountId).StartGoogleBusinessVerificationRequest(startGoogleBusinessVerificationRequest).LocationId(locationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GMBVerificationsAPI.StartGoogleBusinessVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartGoogleBusinessVerification`: StartGoogleBusinessVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `GMBVerificationsAPI.StartGoogleBusinessVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The Zernio account ID (from /v1/accounts) | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartGoogleBusinessVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startGoogleBusinessVerificationRequest** | [**StartGoogleBusinessVerificationRequest**](StartGoogleBusinessVerificationRequest.md) |  | 
 **locationId** | **string** | Override which location to target. If omitted, uses the account&#39;s selected location. | 

### Return type

[**StartGoogleBusinessVerification200Response**](StartGoogleBusinessVerification200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

