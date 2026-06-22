# \ProfilesAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfile**](ProfilesAPI.md#CreateProfile) | **Post** /v1/profiles | Create profile
[**DeleteProfile**](ProfilesAPI.md#DeleteProfile) | **Delete** /v1/profiles/{profileId} | Delete profile
[**GetProfile**](ProfilesAPI.md#GetProfile) | **Get** /v1/profiles/{profileId} | Get profile
[**ListProfiles**](ProfilesAPI.md#ListProfiles) | **Get** /v1/profiles | List profiles
[**UpdateProfile**](ProfilesAPI.md#UpdateProfile) | **Put** /v1/profiles/{profileId} | Update profile



## CreateProfile

> ProfileCreateResponse CreateProfile(ctx).CreateProfileRequest(createProfileRequest).Execute()

Create profile



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
	createProfileRequest := *openapiclient.NewCreateProfileRequest("Name_example") // CreateProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateProfile(context.Background()).CreateProfileRequest(createProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProfile`: ProfileCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProfileRequest** | [**CreateProfileRequest**](CreateProfileRequest.md) |  | 

### Return type

[**ProfileCreateResponse**](ProfileCreateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProfile

> DeleteAccountGroup200Response DeleteProfile(ctx, profileId).Execute()

Delete profile



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
	profileId := "profileId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeleteProfile(context.Background(), profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProfile`: DeleteAccountGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfileRequest struct via the builder pattern


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


## GetProfile

> GetProfile200Response GetProfile(ctx, profileId).Execute()

Get profile



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
	profileId := "profileId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetProfile(context.Background(), profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfile`: GetProfile200Response
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetProfile200Response**](GetProfile200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfiles

> ProfilesListResponse ListProfiles(ctx).IncludeOverLimit(includeOverLimit).Execute()

List profiles



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
	includeOverLimit := true // bool | When true, includes over-limit profiles (marked with isOverLimit: true). (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ListProfiles(context.Background()).IncludeOverLimit(includeOverLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ListProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProfiles`: ProfilesListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ListProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeOverLimit** | **bool** | When true, includes over-limit profiles (marked with isOverLimit: true). | [default to false]

### Return type

[**ProfilesListResponse**](ProfilesListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> UpdateProfile200Response UpdateProfile(ctx, profileId).UpdateProfileRequest(updateProfileRequest).Execute()

Update profile



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
	profileId := "profileId_example" // string | 
	updateProfileRequest := *openapiclient.NewUpdateProfileRequest() // UpdateProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.UpdateProfile(context.Background(), profileId).UpdateProfileRequest(updateProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.UpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProfile`: UpdateProfile200Response
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.UpdateProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProfileRequest** | [**UpdateProfileRequest**](UpdateProfileRequest.md) |  | 

### Return type

[**UpdateProfile200Response**](UpdateProfile200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

