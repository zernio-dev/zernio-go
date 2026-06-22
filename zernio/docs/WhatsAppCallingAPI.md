# \WhatsAppCallingAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableWhatsAppCalling**](WhatsAppCallingAPI.md#DisableWhatsAppCalling) | **Delete** /v1/whatsapp/phone-numbers/{id}/calling | Disable calling on a number
[**EnableWhatsAppCalling**](WhatsAppCallingAPI.md#EnableWhatsAppCalling) | **Post** /v1/whatsapp/phone-numbers/{id}/calling | Enable calling on a number
[**GetWhatsAppCall**](WhatsAppCallingAPI.md#GetWhatsAppCall) | **Get** /v1/whatsapp/calls/{callId} | Get a single call
[**GetWhatsAppCallEstimate**](WhatsAppCallingAPI.md#GetWhatsAppCallEstimate) | **Get** /v1/whatsapp/calls/estimate | Estimate per-minute cost for a destination
[**GetWhatsAppCallPermissions**](WhatsAppCallingAPI.md#GetWhatsAppCallPermissions) | **Get** /v1/whatsapp/call-permissions | Check call permission for a consumer
[**GetWhatsAppCallingConfig**](WhatsAppCallingAPI.md#GetWhatsAppCallingConfig) | **Get** /v1/whatsapp/calling | Get calling config for an account
[**InitiateWhatsAppCall**](WhatsAppCallingAPI.md#InitiateWhatsAppCall) | **Post** /v1/whatsapp/calls | Initiate outbound call
[**ListWhatsAppCalls**](WhatsAppCallingAPI.md#ListWhatsAppCalls) | **Get** /v1/whatsapp/calls | List call history for an account
[**UpdateWhatsAppCalling**](WhatsAppCallingAPI.md#UpdateWhatsAppCalling) | **Patch** /v1/whatsapp/phone-numbers/{id}/calling | Update calling config



## DisableWhatsAppCalling

> DisableWhatsAppCalling(ctx, id).AccountId(accountId).Execute()

Disable calling on a number



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
	id := "id_example" // string | 
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WhatsAppCallingAPI.DisableWhatsAppCalling(context.Background(), id).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppCallingAPI.DisableWhatsAppCalling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableWhatsAppCallingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableWhatsAppCalling

> EnableWhatsAppCalling200Response EnableWhatsAppCalling(ctx, id).EnableWhatsAppCallingRequest(enableWhatsAppCallingRequest).Execute()

Enable calling on a number



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
	id := "id_example" // string | WhatsAppPhoneNumber Mongo ID
	enableWhatsAppCallingRequest := *openapiclient.NewEnableWhatsAppCallingRequest("AccountId_example", "ForwardTo_example") // EnableWhatsAppCallingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppCallingAPI.EnableWhatsAppCalling(context.Background(), id).EnableWhatsAppCallingRequest(enableWhatsAppCallingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppCallingAPI.EnableWhatsAppCalling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableWhatsAppCalling`: EnableWhatsAppCalling200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppCallingAPI.EnableWhatsAppCalling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WhatsAppPhoneNumber Mongo ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableWhatsAppCallingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableWhatsAppCallingRequest** | [**EnableWhatsAppCallingRequest**](EnableWhatsAppCallingRequest.md) |  | 

### Return type

[**EnableWhatsAppCalling200Response**](EnableWhatsAppCalling200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppCall

> GetWhatsAppCall200Response GetWhatsAppCall(ctx, callId).AccountId(accountId).Execute()

Get a single call

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
	callId := "callId_example" // string | 
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppCallingAPI.GetWhatsAppCall(context.Background(), callId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppCallingAPI.GetWhatsAppCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppCall`: GetWhatsAppCall200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppCallingAPI.GetWhatsAppCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** |  | 

### Return type

[**GetWhatsAppCall200Response**](GetWhatsAppCall200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppCallEstimate

> GetWhatsAppCallEstimate200Response GetWhatsAppCallEstimate(ctx).AccountId(accountId).To(to).Minutes(minutes).Recording(recording).Execute()

Estimate per-minute cost for a destination



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
	to := "to_example" // string | 
	minutes := int32(56) // int32 |  (optional)
	recording := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppCallingAPI.GetWhatsAppCallEstimate(context.Background()).AccountId(accountId).To(to).Minutes(minutes).Recording(recording).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppCallingAPI.GetWhatsAppCallEstimate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppCallEstimate`: GetWhatsAppCallEstimate200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppCallingAPI.GetWhatsAppCallEstimate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppCallEstimateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **to** | **string** |  | 
 **minutes** | **int32** |  | 
 **recording** | **bool** |  | 

### Return type

[**GetWhatsAppCallEstimate200Response**](GetWhatsAppCallEstimate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppCallPermissions

> GetWhatsAppCallPermissions200Response GetWhatsAppCallPermissions(ctx).AccountId(accountId).To(to).Execute()

Check call permission for a consumer



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
	to := "to_example" // string | Consumer wa_id (E.164, leading + optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppCallingAPI.GetWhatsAppCallPermissions(context.Background()).AccountId(accountId).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppCallingAPI.GetWhatsAppCallPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppCallPermissions`: GetWhatsAppCallPermissions200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppCallingAPI.GetWhatsAppCallPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppCallPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **to** | **string** | Consumer wa_id (E.164, leading + optional) | 

### Return type

[**GetWhatsAppCallPermissions200Response**](GetWhatsAppCallPermissions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppCallingConfig

> GetWhatsAppCallingConfig200Response GetWhatsAppCallingConfig(ctx).AccountId(accountId).Execute()

Get calling config for an account



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
	accountId := "accountId_example" // string | WhatsApp social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppCallingAPI.GetWhatsAppCallingConfig(context.Background()).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppCallingAPI.GetWhatsAppCallingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppCallingConfig`: GetWhatsAppCallingConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppCallingAPI.GetWhatsAppCallingConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppCallingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**GetWhatsAppCallingConfig200Response**](GetWhatsAppCallingConfig200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateWhatsAppCall

> InitiateWhatsAppCall200Response InitiateWhatsAppCall(ctx).InitiateWhatsAppCallRequest(initiateWhatsAppCallRequest).Execute()

Initiate outbound call



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
	initiateWhatsAppCallRequest := *openapiclient.NewInitiateWhatsAppCallRequest("AccountId_example", "To_example") // InitiateWhatsAppCallRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppCallingAPI.InitiateWhatsAppCall(context.Background()).InitiateWhatsAppCallRequest(initiateWhatsAppCallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppCallingAPI.InitiateWhatsAppCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateWhatsAppCall`: InitiateWhatsAppCall200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppCallingAPI.InitiateWhatsAppCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateWhatsAppCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **initiateWhatsAppCallRequest** | [**InitiateWhatsAppCallRequest**](InitiateWhatsAppCallRequest.md) |  | 

### Return type

[**InitiateWhatsAppCall200Response**](InitiateWhatsAppCall200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWhatsAppCalls

> ListWhatsAppCalls200Response ListWhatsAppCalls(ctx).AccountId(accountId).Status(status).Direction(direction).Since(since).Until(until).Limit(limit).Execute()

List call history for an account



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
	accountId := "accountId_example" // string | 
	status := "status_example" // string |  (optional)
	direction := "direction_example" // string |  (optional)
	since := time.Now() // time.Time |  (optional)
	until := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppCallingAPI.ListWhatsAppCalls(context.Background()).AccountId(accountId).Status(status).Direction(direction).Since(since).Until(until).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppCallingAPI.ListWhatsAppCalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWhatsAppCalls`: ListWhatsAppCalls200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppCallingAPI.ListWhatsAppCalls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWhatsAppCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **status** | **string** |  | 
 **direction** | **string** |  | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **limit** | **int32** |  | 

### Return type

[**ListWhatsAppCalls200Response**](ListWhatsAppCalls200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWhatsAppCalling

> UpdateWhatsAppCalling(ctx, id).UpdateWhatsAppCallingRequest(updateWhatsAppCallingRequest).Execute()

Update calling config



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
	id := "id_example" // string | 
	updateWhatsAppCallingRequest := *openapiclient.NewUpdateWhatsAppCallingRequest("AccountId_example") // UpdateWhatsAppCallingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WhatsAppCallingAPI.UpdateWhatsAppCalling(context.Background(), id).UpdateWhatsAppCallingRequest(updateWhatsAppCallingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppCallingAPI.UpdateWhatsAppCalling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWhatsAppCallingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWhatsAppCallingRequest** | [**UpdateWhatsAppCallingRequest**](UpdateWhatsAppCallingRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

