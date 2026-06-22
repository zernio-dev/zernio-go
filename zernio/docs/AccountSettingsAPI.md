# \AccountSettingsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInstagramIceBreakers**](AccountSettingsAPI.md#DeleteInstagramIceBreakers) | **Delete** /v1/accounts/{accountId}/instagram-ice-breakers | Delete IG ice breakers
[**DeleteMessengerMenu**](AccountSettingsAPI.md#DeleteMessengerMenu) | **Delete** /v1/accounts/{accountId}/messenger-menu | Delete FB persistent menu
[**DeleteTelegramCommands**](AccountSettingsAPI.md#DeleteTelegramCommands) | **Delete** /v1/accounts/{accountId}/telegram-commands | Delete TG bot commands
[**GetInstagramIceBreakers**](AccountSettingsAPI.md#GetInstagramIceBreakers) | **Get** /v1/accounts/{accountId}/instagram-ice-breakers | Get IG ice breakers
[**GetMessengerMenu**](AccountSettingsAPI.md#GetMessengerMenu) | **Get** /v1/accounts/{accountId}/messenger-menu | Get FB persistent menu
[**GetTelegramCommands**](AccountSettingsAPI.md#GetTelegramCommands) | **Get** /v1/accounts/{accountId}/telegram-commands | Get TG bot commands
[**SetInstagramIceBreakers**](AccountSettingsAPI.md#SetInstagramIceBreakers) | **Put** /v1/accounts/{accountId}/instagram-ice-breakers | Set IG ice breakers
[**SetMessengerMenu**](AccountSettingsAPI.md#SetMessengerMenu) | **Put** /v1/accounts/{accountId}/messenger-menu | Set FB persistent menu
[**SetTelegramCommands**](AccountSettingsAPI.md#SetTelegramCommands) | **Put** /v1/accounts/{accountId}/telegram-commands | Set TG bot commands



## DeleteInstagramIceBreakers

> DeleteInstagramIceBreakers(ctx, accountId).Execute()

Delete IG ice breakers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountSettingsAPI.DeleteInstagramIceBreakers(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountSettingsAPI.DeleteInstagramIceBreakers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstagramIceBreakersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteMessengerMenu

> DeleteMessengerMenu(ctx, accountId).Execute()

Delete FB persistent menu



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountSettingsAPI.DeleteMessengerMenu(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountSettingsAPI.DeleteMessengerMenu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessengerMenuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteTelegramCommands

> DeleteTelegramCommands(ctx, accountId).Execute()

Delete TG bot commands



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountSettingsAPI.DeleteTelegramCommands(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountSettingsAPI.DeleteTelegramCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTelegramCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetInstagramIceBreakers

> GetMessengerMenu200Response GetInstagramIceBreakers(ctx, accountId).Execute()

Get IG ice breakers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountSettingsAPI.GetInstagramIceBreakers(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountSettingsAPI.GetInstagramIceBreakers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstagramIceBreakers`: GetMessengerMenu200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountSettingsAPI.GetInstagramIceBreakers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstagramIceBreakersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMessengerMenu200Response**](GetMessengerMenu200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessengerMenu

> GetMessengerMenu200Response GetMessengerMenu(ctx, accountId).Execute()

Get FB persistent menu



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountSettingsAPI.GetMessengerMenu(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountSettingsAPI.GetMessengerMenu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessengerMenu`: GetMessengerMenu200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountSettingsAPI.GetMessengerMenu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessengerMenuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMessengerMenu200Response**](GetMessengerMenu200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelegramCommands

> GetTelegramCommands200Response GetTelegramCommands(ctx, accountId).Execute()

Get TG bot commands



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountSettingsAPI.GetTelegramCommands(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountSettingsAPI.GetTelegramCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelegramCommands`: GetTelegramCommands200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountSettingsAPI.GetTelegramCommands`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelegramCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTelegramCommands200Response**](GetTelegramCommands200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInstagramIceBreakers

> SetInstagramIceBreakers(ctx, accountId).SetInstagramIceBreakersRequest(setInstagramIceBreakersRequest).Execute()

Set IG ice breakers



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
	setInstagramIceBreakersRequest := *openapiclient.NewSetInstagramIceBreakersRequest([]openapiclient.SetInstagramIceBreakersRequestIceBreakersInner{*openapiclient.NewSetInstagramIceBreakersRequestIceBreakersInner("Question_example", "Payload_example")}) // SetInstagramIceBreakersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountSettingsAPI.SetInstagramIceBreakers(context.Background(), accountId).SetInstagramIceBreakersRequest(setInstagramIceBreakersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountSettingsAPI.SetInstagramIceBreakers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetInstagramIceBreakersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setInstagramIceBreakersRequest** | [**SetInstagramIceBreakersRequest**](SetInstagramIceBreakersRequest.md) |  | 

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


## SetMessengerMenu

> SetMessengerMenu(ctx, accountId).SetMessengerMenuRequest(setMessengerMenuRequest).Execute()

Set FB persistent menu



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
	setMessengerMenuRequest := *openapiclient.NewSetMessengerMenuRequest([]map[string]interface{}{map[string]interface{}(123)}) // SetMessengerMenuRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountSettingsAPI.SetMessengerMenu(context.Background(), accountId).SetMessengerMenuRequest(setMessengerMenuRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountSettingsAPI.SetMessengerMenu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMessengerMenuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setMessengerMenuRequest** | [**SetMessengerMenuRequest**](SetMessengerMenuRequest.md) |  | 

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


## SetTelegramCommands

> SetTelegramCommands(ctx, accountId).SetTelegramCommandsRequest(setTelegramCommandsRequest).Execute()

Set TG bot commands



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
	setTelegramCommandsRequest := *openapiclient.NewSetTelegramCommandsRequest([]openapiclient.SetTelegramCommandsRequestCommandsInner{*openapiclient.NewSetTelegramCommandsRequestCommandsInner("Command_example", "Description_example")}) // SetTelegramCommandsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountSettingsAPI.SetTelegramCommands(context.Background(), accountId).SetTelegramCommandsRequest(setTelegramCommandsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountSettingsAPI.SetTelegramCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetTelegramCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setTelegramCommandsRequest** | [**SetTelegramCommandsRequest**](SetTelegramCommandsRequest.md) |  | 

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

