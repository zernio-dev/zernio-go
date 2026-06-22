# \LinkedInMentionsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLinkedInMentions**](LinkedInMentionsAPI.md#GetLinkedInMentions) | **Get** /v1/accounts/{accountId}/linkedin-mentions | Resolve LinkedIn mention



## GetLinkedInMentions

> GetLinkedInMentions200Response GetLinkedInMentions(ctx, accountId).Url(url).DisplayName(displayName).Execute()

Resolve LinkedIn mention



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
	accountId := "accountId_example" // string | The LinkedIn account ID
	url := "miquelpalet" // string | LinkedIn profile URL, company URL, or vanity name.
	displayName := "Miquel Palet" // string | Exact display name as shown on LinkedIn. Required for person mentions to be clickable. Optional for org mentions. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinkedInMentionsAPI.GetLinkedInMentions(context.Background(), accountId).Url(url).DisplayName(displayName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkedInMentionsAPI.GetLinkedInMentions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedInMentions`: GetLinkedInMentions200Response
	fmt.Fprintf(os.Stdout, "Response from `LinkedInMentionsAPI.GetLinkedInMentions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The LinkedIn account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedInMentionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **url** | **string** | LinkedIn profile URL, company URL, or vanity name. | 
 **displayName** | **string** | Exact display name as shown on LinkedIn. Required for person mentions to be clickable. Optional for org mentions. | 

### Return type

[**GetLinkedInMentions200Response**](GetLinkedInMentions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

