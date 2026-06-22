# \WhatsAppTemplatesAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWhatsAppLibraryTemplate**](WhatsAppTemplatesAPI.md#GetWhatsAppLibraryTemplate) | **Get** /v1/whatsapp/template-library | Look up a library template



## GetWhatsAppLibraryTemplate

> GetWhatsAppLibraryTemplate200Response GetWhatsAppLibraryTemplate(ctx).AccountId(accountId).Name(name).Execute()

Look up a library template



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
	name := "name_example" // string | Exact library template name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppTemplatesAPI.GetWhatsAppLibraryTemplate(context.Background()).AccountId(accountId).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppTemplatesAPI.GetWhatsAppLibraryTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppLibraryTemplate`: GetWhatsAppLibraryTemplate200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppTemplatesAPI.GetWhatsAppLibraryTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppLibraryTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | WhatsApp social account ID | 
 **name** | **string** | Exact library template name | 

### Return type

[**GetWhatsAppLibraryTemplate200Response**](GetWhatsAppLibraryTemplate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

