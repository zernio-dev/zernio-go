# \WhatsAppSandboxAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWhatsAppSandboxSession**](WhatsAppSandboxAPI.md#CreateWhatsAppSandboxSession) | **Post** /v1/whatsapp/sandbox/sessions | Start a sandbox activation for a phone
[**DeleteWhatsAppSandboxSession**](WhatsAppSandboxAPI.md#DeleteWhatsAppSandboxSession) | **Delete** /v1/whatsapp/sandbox/sessions/{sessionId} | Revoke a sandbox session
[**ListWhatsAppSandboxSessions**](WhatsAppSandboxAPI.md#ListWhatsAppSandboxSessions) | **Get** /v1/whatsapp/sandbox/sessions | List your sandbox sessions



## CreateWhatsAppSandboxSession

> CreateWhatsAppSandboxSession200Response CreateWhatsAppSandboxSession(ctx).CreateWhatsAppSandboxSessionRequest(createWhatsAppSandboxSessionRequest).Execute()

Start a sandbox activation for a phone



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
	createWhatsAppSandboxSessionRequest := *openapiclient.NewCreateWhatsAppSandboxSessionRequest("+34688246216") // CreateWhatsAppSandboxSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppSandboxAPI.CreateWhatsAppSandboxSession(context.Background()).CreateWhatsAppSandboxSessionRequest(createWhatsAppSandboxSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppSandboxAPI.CreateWhatsAppSandboxSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWhatsAppSandboxSession`: CreateWhatsAppSandboxSession200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppSandboxAPI.CreateWhatsAppSandboxSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWhatsAppSandboxSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWhatsAppSandboxSessionRequest** | [**CreateWhatsAppSandboxSessionRequest**](CreateWhatsAppSandboxSessionRequest.md) |  | 

### Return type

[**CreateWhatsAppSandboxSession200Response**](CreateWhatsAppSandboxSession200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWhatsAppSandboxSession

> UpdateYoutubeDefaultPlaylist200Response DeleteWhatsAppSandboxSession(ctx, sessionId).Execute()

Revoke a sandbox session



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
	sessionId := "sessionId_example" // string | The session id returned by POST /v1/whatsapp/sandbox/sessions.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppSandboxAPI.DeleteWhatsAppSandboxSession(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppSandboxAPI.DeleteWhatsAppSandboxSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWhatsAppSandboxSession`: UpdateYoutubeDefaultPlaylist200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppSandboxAPI.DeleteWhatsAppSandboxSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The session id returned by POST /v1/whatsapp/sandbox/sessions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWhatsAppSandboxSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UpdateYoutubeDefaultPlaylist200Response**](UpdateYoutubeDefaultPlaylist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWhatsAppSandboxSessions

> ListWhatsAppSandboxSessions200Response ListWhatsAppSandboxSessions(ctx).Execute()

List your sandbox sessions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppSandboxAPI.ListWhatsAppSandboxSessions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppSandboxAPI.ListWhatsAppSandboxSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWhatsAppSandboxSessions`: ListWhatsAppSandboxSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppSandboxAPI.ListWhatsAppSandboxSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWhatsAppSandboxSessionsRequest struct via the builder pattern


### Return type

[**ListWhatsAppSandboxSessions200Response**](ListWhatsAppSandboxSessions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

