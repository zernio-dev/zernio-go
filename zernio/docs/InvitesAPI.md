# \InvitesAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInviteToken**](InvitesAPI.md#CreateInviteToken) | **Post** /v1/invite/tokens | Create invite token



## CreateInviteToken

> CreateInviteToken201Response CreateInviteToken(ctx).CreateInviteTokenRequest(createInviteTokenRequest).Execute()

Create invite token



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
	createInviteTokenRequest := *openapiclient.NewCreateInviteTokenRequest("Scope_example") // CreateInviteTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitesAPI.CreateInviteToken(context.Background()).CreateInviteTokenRequest(createInviteTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.CreateInviteToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInviteToken`: CreateInviteToken201Response
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.CreateInviteToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInviteTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInviteTokenRequest** | [**CreateInviteTokenRequest**](CreateInviteTokenRequest.md) |  | 

### Return type

[**CreateInviteToken201Response**](CreateInviteToken201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

