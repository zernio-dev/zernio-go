# \CustomFieldsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearContactFieldValue**](CustomFieldsAPI.md#ClearContactFieldValue) | **Delete** /v1/contacts/{contactId}/fields/{slug} | Clear custom field value
[**CreateCustomField**](CustomFieldsAPI.md#CreateCustomField) | **Post** /v1/custom-fields | Create custom field
[**DeleteCustomField**](CustomFieldsAPI.md#DeleteCustomField) | **Delete** /v1/custom-fields/{fieldId} | Delete custom field
[**ListCustomFields**](CustomFieldsAPI.md#ListCustomFields) | **Get** /v1/custom-fields | List custom field definitions
[**SetContactFieldValue**](CustomFieldsAPI.md#SetContactFieldValue) | **Put** /v1/contacts/{contactId}/fields/{slug} | Set custom field value
[**UpdateCustomField**](CustomFieldsAPI.md#UpdateCustomField) | **Patch** /v1/custom-fields/{fieldId} | Update custom field



## ClearContactFieldValue

> ClearContactFieldValue(ctx, contactId, slug).Execute()

Clear custom field value



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
	contactId := "contactId_example" // string | 
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomFieldsAPI.ClearContactFieldValue(context.Background(), contactId, slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.ClearContactFieldValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearContactFieldValueRequest struct via the builder pattern


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


## CreateCustomField

> CreateCustomField200Response CreateCustomField(ctx).CreateCustomFieldRequest(createCustomFieldRequest).Execute()

Create custom field



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
	createCustomFieldRequest := *openapiclient.NewCreateCustomFieldRequest("ProfileId_example", "Name_example", "Type_example") // CreateCustomFieldRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.CreateCustomField(context.Background()).CreateCustomFieldRequest(createCustomFieldRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.CreateCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomField`: CreateCustomField200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.CreateCustomField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomFieldRequest** | [**CreateCustomFieldRequest**](CreateCustomFieldRequest.md) |  | 

### Return type

[**CreateCustomField200Response**](CreateCustomField200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomField

> DeleteCustomField(ctx, fieldId).Execute()

Delete custom field



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
	fieldId := "fieldId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomFieldsAPI.DeleteCustomField(context.Background(), fieldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.DeleteCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomFieldRequest struct via the builder pattern


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


## ListCustomFields

> ListCustomFields200Response ListCustomFields(ctx).ProfileId(profileId).Execute()

List custom field definitions



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
	profileId := "profileId_example" // string | Filter by profile. Omit to list across all profiles (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.ListCustomFields(context.Background()).ProfileId(profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.ListCustomFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomFields`: ListCustomFields200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.ListCustomFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Filter by profile. Omit to list across all profiles | 

### Return type

[**ListCustomFields200Response**](ListCustomFields200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetContactFieldValue

> SetContactFieldValue(ctx, contactId, slug).SetContactFieldValueRequest(setContactFieldValueRequest).Execute()

Set custom field value



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
	contactId := "contactId_example" // string | 
	slug := "slug_example" // string | 
	setContactFieldValueRequest := *openapiclient.NewSetContactFieldValueRequest(interface{}(123)) // SetContactFieldValueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomFieldsAPI.SetContactFieldValue(context.Background(), contactId, slug).SetContactFieldValueRequest(setContactFieldValueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.SetContactFieldValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetContactFieldValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setContactFieldValueRequest** | [**SetContactFieldValueRequest**](SetContactFieldValueRequest.md) |  | 

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


## UpdateCustomField

> UpdateCustomField200Response UpdateCustomField(ctx, fieldId).UpdateCustomFieldRequest(updateCustomFieldRequest).Execute()

Update custom field



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
	fieldId := "fieldId_example" // string | 
	updateCustomFieldRequest := *openapiclient.NewUpdateCustomFieldRequest() // UpdateCustomFieldRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.UpdateCustomField(context.Background(), fieldId).UpdateCustomFieldRequest(updateCustomFieldRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.UpdateCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomField`: UpdateCustomField200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.UpdateCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomFieldRequest** | [**UpdateCustomFieldRequest**](UpdateCustomFieldRequest.md) |  | 

### Return type

[**UpdateCustomField200Response**](UpdateCustomField200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

