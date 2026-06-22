# \LogsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLogs**](LogsAPI.md#ListLogs) | **Get** /v1/logs | List activity logs



## ListLogs

> ListLogs200Response ListLogs(ctx).Type_(type_).Status(status).Platform(platform).Action(action).Search(search).Days(days).Limit(limit).Skip(skip).Execute()

List activity logs



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
	type_ := "type__example" // string | Log category to query (optional) (default to "publishing")
	status := "status_example" // string | Filter by status (optional)
	platform := "platform_example" // string | Filter by platform (optional)
	action := "action_example" // string | Filter by action (e.g., post.published, message.sent, account.connected, webhook.delivered) (optional)
	search := "search_example" // string | Free-text search across log fields (optional)
	days := int32(56) // int32 | Number of days to look back (max 90) (optional) (default to 90)
	limit := int32(56) // int32 | Maximum number of logs to return (max 100) (optional) (default to 50)
	skip := int32(56) // int32 | Number of logs to skip (for pagination) (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.ListLogs(context.Background()).Type_(type_).Status(status).Platform(platform).Action(action).Search(search).Days(days).Limit(limit).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.ListLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLogs`: ListLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.ListLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Log category to query | [default to &quot;publishing&quot;]
 **status** | **string** | Filter by status | 
 **platform** | **string** | Filter by platform | 
 **action** | **string** | Filter by action (e.g., post.published, message.sent, account.connected, webhook.delivered) | 
 **search** | **string** | Free-text search across log fields | 
 **days** | **int32** | Number of days to look back (max 90) | [default to 90]
 **limit** | **int32** | Maximum number of logs to return (max 100) | [default to 50]
 **skip** | **int32** | Number of logs to skip (for pagination) | [default to 0]

### Return type

[**ListLogs200Response**](ListLogs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

