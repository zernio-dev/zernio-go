# \AdAudiencesAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUsersToAdAudience**](AdAudiencesAPI.md#AddUsersToAdAudience) | **Post** /v1/ads/audiences/{audienceId}/users | Add users to audience
[**CreateAdAudience**](AdAudiencesAPI.md#CreateAdAudience) | **Post** /v1/ads/audiences | Create custom audience
[**DeleteAdAudience**](AdAudiencesAPI.md#DeleteAdAudience) | **Delete** /v1/ads/audiences/{audienceId} | Delete custom audience
[**GetAdAudience**](AdAudiencesAPI.md#GetAdAudience) | **Get** /v1/ads/audiences/{audienceId} | Get audience details
[**ListAdAudiences**](AdAudiencesAPI.md#ListAdAudiences) | **Get** /v1/ads/audiences | List custom audiences



## AddUsersToAdAudience

> AddUsersToAdAudience200Response AddUsersToAdAudience(ctx, audienceId).AddUsersToAdAudienceRequest(addUsersToAdAudienceRequest).Execute()

Add users to audience



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
	audienceId := "audienceId_example" // string | 
	addUsersToAdAudienceRequest := *openapiclient.NewAddUsersToAdAudienceRequest([]openapiclient.AddUsersToAdAudienceRequestUsersInner{*openapiclient.NewAddUsersToAdAudienceRequestUsersInner()}) // AddUsersToAdAudienceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdAudiencesAPI.AddUsersToAdAudience(context.Background(), audienceId).AddUsersToAdAudienceRequest(addUsersToAdAudienceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAudiencesAPI.AddUsersToAdAudience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUsersToAdAudience`: AddUsersToAdAudience200Response
	fmt.Fprintf(os.Stdout, "Response from `AdAudiencesAPI.AddUsersToAdAudience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUsersToAdAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addUsersToAdAudienceRequest** | [**AddUsersToAdAudienceRequest**](AddUsersToAdAudienceRequest.md) |  | 

### Return type

[**AddUsersToAdAudience200Response**](AddUsersToAdAudience200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdAudience

> CreateAdAudience201Response CreateAdAudience(ctx).CreateAdAudienceRequest(createAdAudienceRequest).Execute()

Create custom audience



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
	createAdAudienceRequest := openapiclient.createAdAudience_request{SavedTargetingAudience: openapiclient.NewSavedTargetingAudience("Type_example", "AccountId_example", "Name_example", "TODO")} // CreateAdAudienceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdAudiencesAPI.CreateAdAudience(context.Background()).CreateAdAudienceRequest(createAdAudienceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAudiencesAPI.CreateAdAudience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAdAudience`: CreateAdAudience201Response
	fmt.Fprintf(os.Stdout, "Response from `AdAudiencesAPI.CreateAdAudience`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAdAudienceRequest** | [**CreateAdAudienceRequest**](CreateAdAudienceRequest.md) |  | 

### Return type

[**CreateAdAudience201Response**](CreateAdAudience201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAdAudience

> DeleteAccountGroup200Response DeleteAdAudience(ctx, audienceId).Execute()

Delete custom audience



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
	audienceId := "audienceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdAudiencesAPI.DeleteAdAudience(context.Background(), audienceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAudiencesAPI.DeleteAdAudience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAdAudience`: DeleteAccountGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `AdAudiencesAPI.DeleteAdAudience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteAccountGroup200Response**](DeleteAccountGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdAudience

> GetAdAudience200Response GetAdAudience(ctx, audienceId).Execute()

Get audience details



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
	audienceId := "audienceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdAudiencesAPI.GetAdAudience(context.Background(), audienceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAudiencesAPI.GetAdAudience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdAudience`: GetAdAudience200Response
	fmt.Fprintf(os.Stdout, "Response from `AdAudiencesAPI.GetAdAudience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAdAudience200Response**](GetAdAudience200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAdAudiences

> ListAdAudiences200Response ListAdAudiences(ctx).AccountId(accountId).AdAccountId(adAccountId).Platform(platform).Type_(type_).Execute()

List custom audiences



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
	accountId := "accountId_example" // string | Social account ID
	adAccountId := "adAccountId_example" // string | Platform ad account ID
	platform := "platform_example" // string |  (optional)
	type_ := "type__example" // string | Filter to one audience type. `saved_targeting` returns stored TargetingSpec audiences (each item carries a `spec`); the other types return uploaded/derived audiences. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdAudiencesAPI.ListAdAudiences(context.Background()).AccountId(accountId).AdAccountId(adAccountId).Platform(platform).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAudiencesAPI.ListAdAudiences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAdAudiences`: ListAdAudiences200Response
	fmt.Fprintf(os.Stdout, "Response from `AdAudiencesAPI.ListAdAudiences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAdAudiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Social account ID | 
 **adAccountId** | **string** | Platform ad account ID | 
 **platform** | **string** |  | 
 **type_** | **string** | Filter to one audience type. &#x60;saved_targeting&#x60; returns stored TargetingSpec audiences (each item carries a &#x60;spec&#x60;); the other types return uploaded/derived audiences. | 

### Return type

[**ListAdAudiences200Response**](ListAdAudiences200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

