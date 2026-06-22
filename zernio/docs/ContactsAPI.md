# \ContactsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateContacts**](ContactsAPI.md#BulkCreateContacts) | **Post** /v1/contacts/bulk | Bulk create contacts
[**CreateContact**](ContactsAPI.md#CreateContact) | **Post** /v1/contacts | Create contact
[**DeleteContact**](ContactsAPI.md#DeleteContact) | **Delete** /v1/contacts/{contactId} | Delete contact
[**GetContact**](ContactsAPI.md#GetContact) | **Get** /v1/contacts/{contactId} | Get contact
[**GetContactChannels**](ContactsAPI.md#GetContactChannels) | **Get** /v1/contacts/{contactId}/channels | List channels for a contact
[**ListContacts**](ContactsAPI.md#ListContacts) | **Get** /v1/contacts | List contacts
[**UpdateContact**](ContactsAPI.md#UpdateContact) | **Patch** /v1/contacts/{contactId} | Update contact



## BulkCreateContacts

> BulkCreateContacts200Response BulkCreateContacts(ctx).BulkCreateContactsRequest(bulkCreateContactsRequest).Execute()

Bulk create contacts



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
	bulkCreateContactsRequest := *openapiclient.NewBulkCreateContactsRequest("ProfileId_example", "AccountId_example", "Platform_example", []openapiclient.BulkCreateContactsRequestContactsInner{*openapiclient.NewBulkCreateContactsRequestContactsInner("Name_example", "PlatformIdentifier_example")}) // BulkCreateContactsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.BulkCreateContacts(context.Background()).BulkCreateContactsRequest(bulkCreateContactsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.BulkCreateContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkCreateContacts`: BulkCreateContacts200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.BulkCreateContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkCreateContactsRequest** | [**BulkCreateContactsRequest**](BulkCreateContactsRequest.md) |  | 

### Return type

[**BulkCreateContacts200Response**](BulkCreateContacts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContact

> CreateContact200Response CreateContact(ctx).CreateContactRequest(createContactRequest).Execute()

Create contact



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
	createContactRequest := *openapiclient.NewCreateContactRequest("ProfileId_example", "Name_example") // CreateContactRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.CreateContact(context.Background()).CreateContactRequest(createContactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.CreateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContact`: CreateContact200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.CreateContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContactRequest** | [**CreateContactRequest**](CreateContactRequest.md) |  | 

### Return type

[**CreateContact200Response**](CreateContact200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContact

> DeleteContact(ctx, contactId).Execute()

Delete contact



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContactsAPI.DeleteContact(context.Background(), contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.DeleteContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactRequest struct via the builder pattern


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


## GetContact

> GetContact200Response GetContact(ctx, contactId).Execute()

Get contact



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.GetContact(context.Background(), contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.GetContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContact`: GetContact200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.GetContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContact200Response**](GetContact200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactChannels

> GetContactChannels200Response GetContactChannels(ctx, contactId).Execute()

List channels for a contact



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.GetContactChannels(context.Background(), contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.GetContactChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactChannels`: GetContactChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.GetContactChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContactChannels200Response**](GetContactChannels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContacts

> ListContacts200Response ListContacts(ctx).ProfileId(profileId).Search(search).Tag(tag).Platform(platform).IsSubscribed(isSubscribed).Limit(limit).Skip(skip).Execute()

List contacts



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
	search := "search_example" // string |  (optional)
	tag := "tag_example" // string |  (optional)
	platform := "platform_example" // string |  (optional)
	isSubscribed := "isSubscribed_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	skip := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.ListContacts(context.Background()).ProfileId(profileId).Search(search).Tag(tag).Platform(platform).IsSubscribed(isSubscribed).Limit(limit).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.ListContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContacts`: ListContacts200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.ListContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Filter by profile. Omit to list across all profiles | 
 **search** | **string** |  | 
 **tag** | **string** |  | 
 **platform** | **string** |  | 
 **isSubscribed** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **skip** | **int32** |  | [default to 0]

### Return type

[**ListContacts200Response**](ListContacts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContact

> UpdateContact200Response UpdateContact(ctx, contactId).UpdateContactRequest(updateContactRequest).Execute()

Update contact



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
	updateContactRequest := *openapiclient.NewUpdateContactRequest() // UpdateContactRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.UpdateContact(context.Background(), contactId).UpdateContactRequest(updateContactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.UpdateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContact`: UpdateContact200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.UpdateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateContactRequest** | [**UpdateContactRequest**](UpdateContactRequest.md) |  | 

### Return type

[**UpdateContact200Response**](UpdateContact200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

