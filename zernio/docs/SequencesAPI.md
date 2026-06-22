# \SequencesAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSequence**](SequencesAPI.md#ActivateSequence) | **Post** /v1/sequences/{sequenceId}/activate | Activate sequence
[**CreateSequence**](SequencesAPI.md#CreateSequence) | **Post** /v1/sequences | Create sequence
[**DeleteSequence**](SequencesAPI.md#DeleteSequence) | **Delete** /v1/sequences/{sequenceId} | Delete sequence
[**EnrollContacts**](SequencesAPI.md#EnrollContacts) | **Post** /v1/sequences/{sequenceId}/enroll | Enroll contacts in a sequence
[**GetSequence**](SequencesAPI.md#GetSequence) | **Get** /v1/sequences/{sequenceId} | Get sequence with steps
[**ListSequenceEnrollments**](SequencesAPI.md#ListSequenceEnrollments) | **Get** /v1/sequences/{sequenceId}/enrollments | List enrollments for a sequence
[**ListSequences**](SequencesAPI.md#ListSequences) | **Get** /v1/sequences | List sequences
[**PauseSequence**](SequencesAPI.md#PauseSequence) | **Post** /v1/sequences/{sequenceId}/pause | Pause sequence
[**UnenrollContact**](SequencesAPI.md#UnenrollContact) | **Delete** /v1/sequences/{sequenceId}/enroll/{contactId} | Unenroll contact
[**UpdateSequence**](SequencesAPI.md#UpdateSequence) | **Patch** /v1/sequences/{sequenceId} | Update sequence



## ActivateSequence

> ActivateSequence200Response ActivateSequence(ctx, sequenceId).Execute()

Activate sequence



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
	sequenceId := "sequenceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SequencesAPI.ActivateSequence(context.Background(), sequenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SequencesAPI.ActivateSequence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateSequence`: ActivateSequence200Response
	fmt.Fprintf(os.Stdout, "Response from `SequencesAPI.ActivateSequence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateSequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivateSequence200Response**](ActivateSequence200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSequence

> CreateSequence200Response CreateSequence(ctx).CreateSequenceRequest(createSequenceRequest).Execute()

Create sequence



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
	createSequenceRequest := *openapiclient.NewCreateSequenceRequest("ProfileId_example", "AccountId_example", "Platform_example", "Name_example") // CreateSequenceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SequencesAPI.CreateSequence(context.Background()).CreateSequenceRequest(createSequenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SequencesAPI.CreateSequence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSequence`: CreateSequence200Response
	fmt.Fprintf(os.Stdout, "Response from `SequencesAPI.CreateSequence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSequenceRequest** | [**CreateSequenceRequest**](CreateSequenceRequest.md) |  | 

### Return type

[**CreateSequence200Response**](CreateSequence200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSequence

> DeleteSequence(ctx, sequenceId).Execute()

Delete sequence



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
	sequenceId := "sequenceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SequencesAPI.DeleteSequence(context.Background(), sequenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SequencesAPI.DeleteSequence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSequenceRequest struct via the builder pattern


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


## EnrollContacts

> EnrollContacts200Response EnrollContacts(ctx, sequenceId).EnrollContactsRequest(enrollContactsRequest).Execute()

Enroll contacts in a sequence



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
	sequenceId := "sequenceId_example" // string | 
	enrollContactsRequest := *openapiclient.NewEnrollContactsRequest([]string{"ContactIds_example"}) // EnrollContactsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SequencesAPI.EnrollContacts(context.Background(), sequenceId).EnrollContactsRequest(enrollContactsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SequencesAPI.EnrollContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnrollContacts`: EnrollContacts200Response
	fmt.Fprintf(os.Stdout, "Response from `SequencesAPI.EnrollContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnrollContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollContactsRequest** | [**EnrollContactsRequest**](EnrollContactsRequest.md) |  | 

### Return type

[**EnrollContacts200Response**](EnrollContacts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSequence

> GetSequence200Response GetSequence(ctx, sequenceId).Execute()

Get sequence with steps



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
	sequenceId := "sequenceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SequencesAPI.GetSequence(context.Background(), sequenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SequencesAPI.GetSequence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSequence`: GetSequence200Response
	fmt.Fprintf(os.Stdout, "Response from `SequencesAPI.GetSequence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSequence200Response**](GetSequence200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSequenceEnrollments

> ListSequenceEnrollments200Response ListSequenceEnrollments(ctx, sequenceId).Status(status).Limit(limit).Skip(skip).Execute()

List enrollments for a sequence



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
	sequenceId := "sequenceId_example" // string | 
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	skip := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SequencesAPI.ListSequenceEnrollments(context.Background(), sequenceId).Status(status).Limit(limit).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SequencesAPI.ListSequenceEnrollments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSequenceEnrollments`: ListSequenceEnrollments200Response
	fmt.Fprintf(os.Stdout, "Response from `SequencesAPI.ListSequenceEnrollments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSequenceEnrollmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **skip** | **int32** |  | [default to 0]

### Return type

[**ListSequenceEnrollments200Response**](ListSequenceEnrollments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSequences

> ListSequences200Response ListSequences(ctx).ProfileId(profileId).Status(status).Limit(limit).Skip(skip).Execute()

List sequences



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
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	skip := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SequencesAPI.ListSequences(context.Background()).ProfileId(profileId).Status(status).Limit(limit).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SequencesAPI.ListSequences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSequences`: ListSequences200Response
	fmt.Fprintf(os.Stdout, "Response from `SequencesAPI.ListSequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Filter by profile. Omit to list across all profiles | 
 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **skip** | **int32** |  | [default to 0]

### Return type

[**ListSequences200Response**](ListSequences200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseSequence

> ActivateSequence200Response PauseSequence(ctx, sequenceId).Execute()

Pause sequence



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
	sequenceId := "sequenceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SequencesAPI.PauseSequence(context.Background(), sequenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SequencesAPI.PauseSequence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseSequence`: ActivateSequence200Response
	fmt.Fprintf(os.Stdout, "Response from `SequencesAPI.PauseSequence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseSequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivateSequence200Response**](ActivateSequence200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnenrollContact

> UnenrollContact(ctx, sequenceId, contactId).Execute()

Unenroll contact



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
	sequenceId := "sequenceId_example" // string | 
	contactId := "contactId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SequencesAPI.UnenrollContact(context.Background(), sequenceId, contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SequencesAPI.UnenrollContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sequenceId** | **string** |  | 
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnenrollContactRequest struct via the builder pattern


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


## UpdateSequence

> UpdateSequence200Response UpdateSequence(ctx, sequenceId).UpdateSequenceRequest(updateSequenceRequest).Execute()

Update sequence



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
	sequenceId := "sequenceId_example" // string | 
	updateSequenceRequest := *openapiclient.NewUpdateSequenceRequest() // UpdateSequenceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SequencesAPI.UpdateSequence(context.Background(), sequenceId).UpdateSequenceRequest(updateSequenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SequencesAPI.UpdateSequence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSequence`: UpdateSequence200Response
	fmt.Fprintf(os.Stdout, "Response from `SequencesAPI.UpdateSequence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSequenceRequest** | [**UpdateSequenceRequest**](UpdateSequenceRequest.md) |  | 

### Return type

[**UpdateSequence200Response**](UpdateSequence200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

