# \WhatsAppPhoneNumbersAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckWhatsAppNumberAvailability**](WhatsAppPhoneNumbersAPI.md#CheckWhatsAppNumberAvailability) | **Get** /v1/whatsapp/phone-numbers/availability | Check a country&#39;s availability + address constraint
[**GetWhatsAppNumberInfo**](WhatsAppPhoneNumbersAPI.md#GetWhatsAppNumberInfo) | **Get** /v1/whatsapp/number-info | Get number status
[**GetWhatsAppNumberKycForm**](WhatsAppPhoneNumbersAPI.md#GetWhatsAppNumberKycForm) | **Get** /v1/whatsapp/phone-numbers/kyc | Get regulated-number KYC form spec
[**GetWhatsAppNumberRemediation**](WhatsAppPhoneNumbersAPI.md#GetWhatsAppNumberRemediation) | **Get** /v1/whatsapp/phone-numbers/{id}/remediate | Get the declined requirements to fix
[**GetWhatsAppPhoneNumber**](WhatsAppPhoneNumbersAPI.md#GetWhatsAppPhoneNumber) | **Get** /v1/whatsapp/phone-numbers/{phoneNumberId} | Get phone number
[**GetWhatsAppPhoneNumbers**](WhatsAppPhoneNumbersAPI.md#GetWhatsAppPhoneNumbers) | **Get** /v1/whatsapp/phone-numbers | List phone numbers
[**ListWhatsAppNumberCountries**](WhatsAppPhoneNumbersAPI.md#ListWhatsAppNumberCountries) | **Get** /v1/whatsapp/phone-numbers/countries | List offerable number countries
[**PurchaseWhatsAppPhoneNumber**](WhatsAppPhoneNumbersAPI.md#PurchaseWhatsAppPhoneNumber) | **Post** /v1/whatsapp/phone-numbers/purchase | Purchase phone number
[**ReleaseWhatsAppPhoneNumber**](WhatsAppPhoneNumbersAPI.md#ReleaseWhatsAppPhoneNumber) | **Delete** /v1/whatsapp/phone-numbers/{phoneNumberId} | Release phone number
[**RemediateWhatsAppNumber**](WhatsAppPhoneNumbersAPI.md#RemediateWhatsAppNumber) | **Post** /v1/whatsapp/phone-numbers/{id}/remediate | Fix a declined number and re-submit
[**SearchAvailableWhatsAppNumbers**](WhatsAppPhoneNumbersAPI.md#SearchAvailableWhatsAppNumbers) | **Get** /v1/whatsapp/phone-numbers/available | Search available numbers to purchase
[**SubmitWhatsAppNumberKyc**](WhatsAppPhoneNumbersAPI.md#SubmitWhatsAppNumberKyc) | **Post** /v1/whatsapp/phone-numbers/kyc | Submit regulated-number KYC
[**UploadWhatsAppNumberKycDocument**](WhatsAppPhoneNumbersAPI.md#UploadWhatsAppNumberKycDocument) | **Post** /v1/whatsapp/phone-numbers/kyc/upload-document | Upload a single regulated-number KYC document
[**ValidateWhatsAppNumberKycAddress**](WhatsAppPhoneNumbersAPI.md#ValidateWhatsAppNumberKycAddress) | **Post** /v1/whatsapp/phone-numbers/kyc/validate-address | Pre-validate a regulated-number KYC address (Tier 4)



## CheckWhatsAppNumberAvailability

> CheckWhatsAppNumberAvailability200Response CheckWhatsAppNumberAvailability(ctx).Country(country).Execute()

Check a country's availability + address constraint



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
	country := "country_example" // string | ISO-2 country code.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.CheckWhatsAppNumberAvailability(context.Background()).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.CheckWhatsAppNumberAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckWhatsAppNumberAvailability`: CheckWhatsAppNumberAvailability200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.CheckWhatsAppNumberAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckWhatsAppNumberAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | ISO-2 country code. | 

### Return type

[**CheckWhatsAppNumberAvailability200Response**](CheckWhatsAppNumberAvailability200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppNumberInfo

> GetWhatsAppNumberInfo200Response GetWhatsAppNumberInfo(ctx).AccountId(accountId).Execute()

Get number status



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.GetWhatsAppNumberInfo(context.Background()).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.GetWhatsAppNumberInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppNumberInfo`: GetWhatsAppNumberInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.GetWhatsAppNumberInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppNumberInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**GetWhatsAppNumberInfo200Response**](GetWhatsAppNumberInfo200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppNumberKycForm

> GetWhatsAppNumberKycForm200Response GetWhatsAppNumberKycForm(ctx).Country(country).ProfileId(profileId).Execute()

Get regulated-number KYC form spec



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
	country := "country_example" // string | 
	profileId := "profileId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.GetWhatsAppNumberKycForm(context.Background()).Country(country).ProfileId(profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.GetWhatsAppNumberKycForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppNumberKycForm`: GetWhatsAppNumberKycForm200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.GetWhatsAppNumberKycForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppNumberKycFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** |  | 
 **profileId** | **string** |  | 

### Return type

[**GetWhatsAppNumberKycForm200Response**](GetWhatsAppNumberKycForm200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppNumberRemediation

> GetWhatsAppNumberRemediation200Response GetWhatsAppNumberRemediation(ctx, id).Execute()

Get the declined requirements to fix



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
	id := "id_example" // string | WhatsAppPhoneNumber id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.GetWhatsAppNumberRemediation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.GetWhatsAppNumberRemediation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppNumberRemediation`: GetWhatsAppNumberRemediation200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.GetWhatsAppNumberRemediation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WhatsAppPhoneNumber id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppNumberRemediationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWhatsAppNumberRemediation200Response**](GetWhatsAppNumberRemediation200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppPhoneNumber

> GetWhatsAppPhoneNumber200Response GetWhatsAppPhoneNumber(ctx, phoneNumberId).Execute()

Get phone number



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
	phoneNumberId := "phoneNumberId_example" // string | Phone number record ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.GetWhatsAppPhoneNumber(context.Background(), phoneNumberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.GetWhatsAppPhoneNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppPhoneNumber`: GetWhatsAppPhoneNumber200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.GetWhatsAppPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | **string** | Phone number record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWhatsAppPhoneNumber200Response**](GetWhatsAppPhoneNumber200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppPhoneNumbers

> GetWhatsAppPhoneNumbers200Response GetWhatsAppPhoneNumbers(ctx).Status(status).ProfileId(profileId).Execute()

List phone numbers



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
	status := "status_example" // string | Filter by status (by default excludes released numbers). NOTE: `status=pending_regulatory` returns the \"provisioning\" view — numbers still in review PLUS recently-declined (last 30 days) ones, so a failed registration surfaces (with `regulatoryDeclineReason`) instead of silently disappearing. Declined numbers can be re-submitted via POST /v1/whatsapp/phone-numbers/{id}/remediate. `verifying` is the short-lived state after the number is provisioned on our side while WhatsApp confirms the activation code; the number is not billed until it reaches `active`.  (optional)
	profileId := "profileId_example" // string | Filter by profile (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.GetWhatsAppPhoneNumbers(context.Background()).Status(status).ProfileId(profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.GetWhatsAppPhoneNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppPhoneNumbers`: GetWhatsAppPhoneNumbers200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.GetWhatsAppPhoneNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppPhoneNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter by status (by default excludes released numbers). NOTE: &#x60;status&#x3D;pending_regulatory&#x60; returns the \&quot;provisioning\&quot; view — numbers still in review PLUS recently-declined (last 30 days) ones, so a failed registration surfaces (with &#x60;regulatoryDeclineReason&#x60;) instead of silently disappearing. Declined numbers can be re-submitted via POST /v1/whatsapp/phone-numbers/{id}/remediate. &#x60;verifying&#x60; is the short-lived state after the number is provisioned on our side while WhatsApp confirms the activation code; the number is not billed until it reaches &#x60;active&#x60;.  | 
 **profileId** | **string** | Filter by profile | 

### Return type

[**GetWhatsAppPhoneNumbers200Response**](GetWhatsAppPhoneNumbers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWhatsAppNumberCountries

> ListWhatsAppNumberCountries200Response ListWhatsAppNumberCountries(ctx).Execute()

List offerable number countries



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
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.ListWhatsAppNumberCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.ListWhatsAppNumberCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWhatsAppNumberCountries`: ListWhatsAppNumberCountries200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.ListWhatsAppNumberCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWhatsAppNumberCountriesRequest struct via the builder pattern


### Return type

[**ListWhatsAppNumberCountries200Response**](ListWhatsAppNumberCountries200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseWhatsAppPhoneNumber

> PurchaseWhatsAppPhoneNumber200Response PurchaseWhatsAppPhoneNumber(ctx).PurchaseWhatsAppPhoneNumberRequest(purchaseWhatsAppPhoneNumberRequest).Execute()

Purchase phone number



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
	purchaseWhatsAppPhoneNumberRequest := *openapiclient.NewPurchaseWhatsAppPhoneNumberRequest("ProfileId_example") // PurchaseWhatsAppPhoneNumberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.PurchaseWhatsAppPhoneNumber(context.Background()).PurchaseWhatsAppPhoneNumberRequest(purchaseWhatsAppPhoneNumberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.PurchaseWhatsAppPhoneNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseWhatsAppPhoneNumber`: PurchaseWhatsAppPhoneNumber200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.PurchaseWhatsAppPhoneNumber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseWhatsAppPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purchaseWhatsAppPhoneNumberRequest** | [**PurchaseWhatsAppPhoneNumberRequest**](PurchaseWhatsAppPhoneNumberRequest.md) |  | 

### Return type

[**PurchaseWhatsAppPhoneNumber200Response**](PurchaseWhatsAppPhoneNumber200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseWhatsAppPhoneNumber

> ReleaseWhatsAppPhoneNumber200Response ReleaseWhatsAppPhoneNumber(ctx, phoneNumberId).Execute()

Release phone number



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
	phoneNumberId := "phoneNumberId_example" // string | Phone number record ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.ReleaseWhatsAppPhoneNumber(context.Background(), phoneNumberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.ReleaseWhatsAppPhoneNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReleaseWhatsAppPhoneNumber`: ReleaseWhatsAppPhoneNumber200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.ReleaseWhatsAppPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | **string** | Phone number record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseWhatsAppPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReleaseWhatsAppPhoneNumber200Response**](ReleaseWhatsAppPhoneNumber200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemediateWhatsAppNumber

> RemediateWhatsAppNumber200Response RemediateWhatsAppNumber(ctx, id).RemediateWhatsAppNumberRequest(remediateWhatsAppNumberRequest).Execute()

Fix a declined number and re-submit



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
	id := "id_example" // string | 
	remediateWhatsAppNumberRequest := *openapiclient.NewRemediateWhatsAppNumberRequest() // RemediateWhatsAppNumberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.RemediateWhatsAppNumber(context.Background(), id).RemediateWhatsAppNumberRequest(remediateWhatsAppNumberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.RemediateWhatsAppNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemediateWhatsAppNumber`: RemediateWhatsAppNumber200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.RemediateWhatsAppNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemediateWhatsAppNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remediateWhatsAppNumberRequest** | [**RemediateWhatsAppNumberRequest**](RemediateWhatsAppNumberRequest.md) |  | 

### Return type

[**RemediateWhatsAppNumber200Response**](RemediateWhatsAppNumber200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAvailableWhatsAppNumbers

> SearchAvailableWhatsAppNumbers200Response SearchAvailableWhatsAppNumbers(ctx).Country(country).Type_(type_).Prefix(prefix).Locality(locality).Contains(contains).Limit(limit).Execute()

Search available numbers to purchase



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
	country := "country_example" // string |  (optional) (default to "US")
	type_ := "type__example" // string | Number type; defaults to the country's WhatsApp-safe type (optional)
	prefix := "prefix_example" // string | Area code (optional)
	locality := "locality_example" // string | City (optional)
	contains := "contains_example" // string | Pattern to match within the number (optional)
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.SearchAvailableWhatsAppNumbers(context.Background()).Country(country).Type_(type_).Prefix(prefix).Locality(locality).Contains(contains).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.SearchAvailableWhatsAppNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAvailableWhatsAppNumbers`: SearchAvailableWhatsAppNumbers200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.SearchAvailableWhatsAppNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAvailableWhatsAppNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** |  | [default to &quot;US&quot;]
 **type_** | **string** | Number type; defaults to the country&#39;s WhatsApp-safe type | 
 **prefix** | **string** | Area code | 
 **locality** | **string** | City | 
 **contains** | **string** | Pattern to match within the number | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**SearchAvailableWhatsAppNumbers200Response**](SearchAvailableWhatsAppNumbers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitWhatsAppNumberKyc

> SubmitWhatsAppNumberKyc200Response SubmitWhatsAppNumberKyc(ctx).SubmitWhatsAppNumberKycRequest(submitWhatsAppNumberKycRequest).Execute()

Submit regulated-number KYC



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
	submitWhatsAppNumberKycRequest := *openapiclient.NewSubmitWhatsAppNumberKycRequest("ProfileId_example", "Country_example") // SubmitWhatsAppNumberKycRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.SubmitWhatsAppNumberKyc(context.Background()).SubmitWhatsAppNumberKycRequest(submitWhatsAppNumberKycRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.SubmitWhatsAppNumberKyc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitWhatsAppNumberKyc`: SubmitWhatsAppNumberKyc200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.SubmitWhatsAppNumberKyc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitWhatsAppNumberKycRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitWhatsAppNumberKycRequest** | [**SubmitWhatsAppNumberKycRequest**](SubmitWhatsAppNumberKycRequest.md) |  | 

### Return type

[**SubmitWhatsAppNumberKyc200Response**](SubmitWhatsAppNumberKyc200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadWhatsAppNumberKycDocument

> UploadWhatsAppNumberKycDocument200Response UploadWhatsAppNumberKycDocument(ctx).XFilename(xFilename).Body(body).Execute()

Upload a single regulated-number KYC document



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
	xFilename := "xFilename_example" // string | URL-encoded original filename.
	body := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.UploadWhatsAppNumberKycDocument(context.Background()).XFilename(xFilename).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.UploadWhatsAppNumberKycDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadWhatsAppNumberKycDocument`: UploadWhatsAppNumberKycDocument200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.UploadWhatsAppNumberKycDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadWhatsAppNumberKycDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFilename** | **string** | URL-encoded original filename. | 
 **body** | ***os.File** |  | 

### Return type

[**UploadWhatsAppNumberKycDocument200Response**](UploadWhatsAppNumberKycDocument200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateWhatsAppNumberKycAddress

> ValidateWhatsAppNumberKycAddress200Response ValidateWhatsAppNumberKycAddress(ctx).ValidateWhatsAppNumberKycAddressRequest(validateWhatsAppNumberKycAddressRequest).Execute()

Pre-validate a regulated-number KYC address (Tier 4)



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
	validateWhatsAppNumberKycAddressRequest := *openapiclient.NewValidateWhatsAppNumberKycAddressRequest("Country_example", "StreetAddress_example", "Locality_example", "PostalCode_example") // ValidateWhatsAppNumberKycAddressRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppPhoneNumbersAPI.ValidateWhatsAppNumberKycAddress(context.Background()).ValidateWhatsAppNumberKycAddressRequest(validateWhatsAppNumberKycAddressRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppPhoneNumbersAPI.ValidateWhatsAppNumberKycAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateWhatsAppNumberKycAddress`: ValidateWhatsAppNumberKycAddress200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppPhoneNumbersAPI.ValidateWhatsAppNumberKycAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateWhatsAppNumberKycAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateWhatsAppNumberKycAddressRequest** | [**ValidateWhatsAppNumberKycAddressRequest**](ValidateWhatsAppNumberKycAddressRequest.md) |  | 

### Return type

[**ValidateWhatsAppNumberKycAddress200Response**](ValidateWhatsAppNumberKycAddress200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

