# \AdsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConversionAssociations**](AdsAPI.md#AddConversionAssociations) | **Post** /v1/accounts/{accountId}/conversion-destinations/{destinationId}/associations | Associate campaigns with a conversion destination
[**AdjustConversions**](AdsAPI.md#AdjustConversions) | **Post** /v1/ads/conversions/adjustments | Adjust already-uploaded conversions (Google only)
[**ArchiveLeadForm**](AdsAPI.md#ArchiveLeadForm) | **Delete** /v1/ads/lead-forms/{formId} | Archive a Lead Gen form
[**BoostPost**](AdsAPI.md#BoostPost) | **Post** /v1/ads/boost | Boost post as ad
[**CreateConversionDestination**](AdsAPI.md#CreateConversionDestination) | **Post** /v1/accounts/{accountId}/conversion-destinations | Create a conversion destination (LinkedIn, Google Ads)
[**CreateCtwaAd**](AdsAPI.md#CreateCtwaAd) | **Post** /v1/ads/ctwa | Create Click-to-WhatsApp ad(s)
[**CreateLeadForm**](AdsAPI.md#CreateLeadForm) | **Post** /v1/ads/lead-forms | Create a Lead Gen (Instant) form
[**CreateStandaloneAd**](AdsAPI.md#CreateStandaloneAd) | **Post** /v1/ads/create | Create standalone ad
[**CreateTestLead**](AdsAPI.md#CreateTestLead) | **Post** /v1/ads/lead-forms/{formId}/test-leads | Create a synthetic test lead
[**DeleteAd**](AdsAPI.md#DeleteAd) | **Delete** /v1/ads/{adId} | Cancel an ad
[**DeleteConversionDestination**](AdsAPI.md#DeleteConversionDestination) | **Delete** /v1/accounts/{accountId}/conversion-destinations/{destinationId} | Soft-delete a conversion destination
[**EstimateAdReach**](AdsAPI.md#EstimateAdReach) | **Post** /v1/ads/targeting/reach-estimate | Estimate audience reach
[**GetAd**](AdsAPI.md#GetAd) | **Get** /v1/ads/{adId} | Get ad details
[**GetAdAnalytics**](AdsAPI.md#GetAdAnalytics) | **Get** /v1/ads/{adId}/analytics | Get ad analytics
[**GetAdComments**](AdsAPI.md#GetAdComments) | **Get** /v1/ads/{adId}/comments | List comments on an ad
[**GetAdTrackingTags**](AdsAPI.md#GetAdTrackingTags) | **Get** /v1/ads/{adId}/tracking-tags | Read an ad&#39;s click-URL tracking tags
[**GetConversionDestination**](AdsAPI.md#GetConversionDestination) | **Get** /v1/accounts/{accountId}/conversion-destinations/{destinationId} | Fetch a single conversion destination
[**GetConversionMetrics**](AdsAPI.md#GetConversionMetrics) | **Get** /v1/accounts/{accountId}/conversion-destinations/{destinationId}/metrics | Fetch attribution metrics for a conversion destination
[**GetConversionsQuality**](AdsAPI.md#GetConversionsQuality) | **Get** /v1/ads/conversions/quality | Read Event Match Quality + coverage for a Meta pixel
[**GetLeadForm**](AdsAPI.md#GetLeadForm) | **Get** /v1/ads/lead-forms/{formId} | Get a single Lead Gen form
[**ListAdAccounts**](AdsAPI.md#ListAdAccounts) | **Get** /v1/ads/accounts | List ad accounts
[**ListAdCatalogProductSets**](AdsAPI.md#ListAdCatalogProductSets) | **Get** /v1/ads/catalogs/{catalogId}/product-sets | List a catalog&#39;s product sets
[**ListAdCatalogs**](AdsAPI.md#ListAdCatalogs) | **Get** /v1/ads/catalogs | List Meta product catalogs
[**ListAds**](AdsAPI.md#ListAds) | **Get** /v1/ads | List ads
[**ListAdsBusinessCenters**](AdsAPI.md#ListAdsBusinessCenters) | **Get** /v1/ads/business-centers | List TikTok Business Centers
[**ListConversionAssociations**](AdsAPI.md#ListConversionAssociations) | **Get** /v1/accounts/{accountId}/conversion-destinations/{destinationId}/associations | List campaigns associated with a conversion destination
[**ListConversionDestinations**](AdsAPI.md#ListConversionDestinations) | **Get** /v1/accounts/{accountId}/conversion-destinations | List destinations for the Conversions API
[**ListFormLeads**](AdsAPI.md#ListFormLeads) | **Get** /v1/ads/lead-forms/{formId}/leads | List leads for a single form
[**ListLeadForms**](AdsAPI.md#ListLeadForms) | **Get** /v1/ads/lead-forms | List Lead Gen (Instant) forms
[**ListLeads**](AdsAPI.md#ListLeads) | **Get** /v1/ads/leads | List submitted leads (cross-form CRM view)
[**ListWhatsAppConversions**](AdsAPI.md#ListWhatsAppConversions) | **Get** /v1/whatsapp/conversions | List recent WhatsApp conversion events
[**RemoveConversionAssociations**](AdsAPI.md#RemoveConversionAssociations) | **Delete** /v1/accounts/{accountId}/conversion-destinations/{destinationId}/associations | Remove campaign↔conversion associations
[**SearchAdInterests**](AdsAPI.md#SearchAdInterests) | **Get** /v1/ads/interests | Search targeting interests (deprecated)
[**SearchAdTargeting**](AdsAPI.md#SearchAdTargeting) | **Get** /v1/ads/targeting/search | Search targeting options
[**SendConversions**](AdsAPI.md#SendConversions) | **Post** /v1/ads/conversions | Send conversion events to an ad platform
[**SendWhatsAppConversion**](AdsAPI.md#SendWhatsAppConversion) | **Post** /v1/whatsapp/conversions | Send WhatsApp conversion event
[**UpdateAd**](AdsAPI.md#UpdateAd) | **Put** /v1/ads/{adId} | Update ad
[**UpdateAdTrackingTags**](AdsAPI.md#UpdateAdTrackingTags) | **Patch** /v1/ads/{adId}/tracking-tags | Set/update an ad&#39;s click-URL tracking tags
[**UpdateConversionDestination**](AdsAPI.md#UpdateConversionDestination) | **Patch** /v1/accounts/{accountId}/conversion-destinations/{destinationId} | Update a conversion destination



## AddConversionAssociations

> AddConversionAssociations200Response AddConversionAssociations(ctx, accountId, destinationId).AddConversionAssociationsRequest(addConversionAssociationsRequest).Execute()

Associate campaigns with a conversion destination



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
	accountId := "accountId_example" // string | 
	destinationId := "destinationId_example" // string | 
	addConversionAssociationsRequest := *openapiclient.NewAddConversionAssociationsRequest("AdAccountId_example", []string{"CampaignIds_example"}) // AddConversionAssociationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.AddConversionAssociations(context.Background(), accountId, destinationId).AddConversionAssociationsRequest(addConversionAssociationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.AddConversionAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddConversionAssociations`: AddConversionAssociations200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.AddConversionAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddConversionAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addConversionAssociationsRequest** | [**AddConversionAssociationsRequest**](AddConversionAssociationsRequest.md) |  | 

### Return type

[**AddConversionAssociations200Response**](AddConversionAssociations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdjustConversions

> AdjustConversions200Response AdjustConversions(ctx).AdjustConversionsRequest(adjustConversionsRequest).Execute()

Adjust already-uploaded conversions (Google only)



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
	adjustConversionsRequest := *openapiclient.NewAdjustConversionsRequest("AccountId_example", "DestinationId_example", []openapiclient.AdjustConversionsRequestAdjustmentsInner{*openapiclient.NewAdjustConversionsRequestAdjustmentsInner("AdjustmentType_example", float32(123))}) // AdjustConversionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.AdjustConversions(context.Background()).AdjustConversionsRequest(adjustConversionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.AdjustConversions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdjustConversions`: AdjustConversions200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.AdjustConversions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdjustConversionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adjustConversionsRequest** | [**AdjustConversionsRequest**](AdjustConversionsRequest.md) |  | 

### Return type

[**AdjustConversions200Response**](AdjustConversions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveLeadForm

> ArchiveLeadForm200Response ArchiveLeadForm(ctx, formId).AccountId(accountId).Execute()

Archive a Lead Gen form



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
	formId := "formId_example" // string | 
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.ArchiveLeadForm(context.Background(), formId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.ArchiveLeadForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveLeadForm`: ArchiveLeadForm200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.ArchiveLeadForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveLeadFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** |  | 

### Return type

[**ArchiveLeadForm200Response**](ArchiveLeadForm200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoostPost

> UpdateAd200Response BoostPost(ctx).BoostPostRequest(boostPostRequest).Execute()

Boost post as ad



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
	boostPostRequest := *openapiclient.NewBoostPostRequest("AccountId_example", "AdAccountId_example", "Name_example", "Goal_example", *openapiclient.NewUpdateAdCampaignRequestBudget(float32(123), "Type_example")) // BoostPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.BoostPost(context.Background()).BoostPostRequest(boostPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.BoostPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoostPost`: UpdateAd200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.BoostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBoostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **boostPostRequest** | [**BoostPostRequest**](BoostPostRequest.md) |  | 

### Return type

[**UpdateAd200Response**](UpdateAd200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConversionDestination

> CreateConversionDestination201Response CreateConversionDestination(ctx, accountId).CreateConversionDestinationRequest(createConversionDestinationRequest).Execute()

Create a conversion destination (LinkedIn, Google Ads)



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
	accountId := "accountId_example" // string | SocialAccount ID (linkedinads or googleads).
	createConversionDestinationRequest := *openapiclient.NewCreateConversionDestinationRequest("AdAccountId_example", "Name_example", "Type_example") // CreateConversionDestinationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.CreateConversionDestination(context.Background(), accountId).CreateConversionDestinationRequest(createConversionDestinationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.CreateConversionDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConversionDestination`: CreateConversionDestination201Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.CreateConversionDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | SocialAccount ID (linkedinads or googleads). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConversionDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createConversionDestinationRequest** | [**CreateConversionDestinationRequest**](CreateConversionDestinationRequest.md) |  | 

### Return type

[**CreateConversionDestination201Response**](CreateConversionDestination201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCtwaAd

> CreateCtwaAd201Response CreateCtwaAd(ctx).CreateCtwaAdRequest(createCtwaAdRequest).Execute()

Create Click-to-WhatsApp ad(s)



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
	createCtwaAdRequest := *openapiclient.NewCreateCtwaAdRequest("AccountId_example", "AdAccountId_example", "Name_example", float32(123), "BudgetType_example") // CreateCtwaAdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.CreateCtwaAd(context.Background()).CreateCtwaAdRequest(createCtwaAdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.CreateCtwaAd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCtwaAd`: CreateCtwaAd201Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.CreateCtwaAd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCtwaAdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCtwaAdRequest** | [**CreateCtwaAdRequest**](CreateCtwaAdRequest.md) |  | 

### Return type

[**CreateCtwaAd201Response**](CreateCtwaAd201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLeadForm

> CreateLeadForm200Response CreateLeadForm(ctx).CreateLeadFormRequest(createLeadFormRequest).Execute()

Create a Lead Gen (Instant) form



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
	createLeadFormRequest := *openapiclient.NewCreateLeadFormRequest("AccountId_example", "Name_example", []openapiclient.CreateLeadFormRequestQuestionsInner{*openapiclient.NewCreateLeadFormRequestQuestionsInner("Type_example")}, "PrivacyPolicyUrl_example") // CreateLeadFormRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.CreateLeadForm(context.Background()).CreateLeadFormRequest(createLeadFormRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.CreateLeadForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLeadForm`: CreateLeadForm200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.CreateLeadForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLeadFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLeadFormRequest** | [**CreateLeadFormRequest**](CreateLeadFormRequest.md) |  | 

### Return type

[**CreateLeadForm200Response**](CreateLeadForm200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStandaloneAd

> CreateStandaloneAd201Response CreateStandaloneAd(ctx).CreateStandaloneAdRequest(createStandaloneAdRequest).IdempotencyKey(idempotencyKey).Execute()

Create standalone ad



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
	createStandaloneAdRequest := *openapiclient.NewCreateStandaloneAdRequest("AccountId_example", "AdAccountId_example", "Name_example") // CreateStandaloneAdRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Optional client-generated unique key (e.g. a UUID) that makes create retries safe. Same key + same body replays the original response; same key + different body → 422; key still processing → 409. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.CreateStandaloneAd(context.Background()).CreateStandaloneAdRequest(createStandaloneAdRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.CreateStandaloneAd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStandaloneAd`: CreateStandaloneAd201Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.CreateStandaloneAd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStandaloneAdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStandaloneAdRequest** | [**CreateStandaloneAdRequest**](CreateStandaloneAdRequest.md) |  | 
 **idempotencyKey** | **string** | Optional client-generated unique key (e.g. a UUID) that makes create retries safe. Same key + same body replays the original response; same key + different body → 422; key still processing → 409. | 

### Return type

[**CreateStandaloneAd201Response**](CreateStandaloneAd201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTestLead

> CreateTestLead200Response CreateTestLead(ctx, formId).CreateTestLeadRequest(createTestLeadRequest).Execute()

Create a synthetic test lead



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
	formId := "formId_example" // string | 
	createTestLeadRequest := *openapiclient.NewCreateTestLeadRequest("AccountId_example", []openapiclient.CreateTestLeadRequestFieldDataInner{*openapiclient.NewCreateTestLeadRequestFieldDataInner("Name_example", []string{"Values_example"})}) // CreateTestLeadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.CreateTestLead(context.Background(), formId).CreateTestLeadRequest(createTestLeadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.CreateTestLead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTestLead`: CreateTestLead200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.CreateTestLead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTestLeadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTestLeadRequest** | [**CreateTestLeadRequest**](CreateTestLeadRequest.md) |  | 

### Return type

[**CreateTestLead200Response**](CreateTestLead200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAd

> DeleteAccountGroup200Response DeleteAd(ctx, adId).Execute()

Cancel an ad



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
	adId := "adId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.DeleteAd(context.Background(), adId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.DeleteAd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAd`: DeleteAccountGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.DeleteAd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdRequest struct via the builder pattern


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


## DeleteConversionDestination

> DeleteConversionDestination(ctx, accountId, destinationId).AdAccountId(adAccountId).Execute()

Soft-delete a conversion destination



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
	accountId := "accountId_example" // string | 
	destinationId := "destinationId_example" // string | 
	adAccountId := "adAccountId_example" // string | Required as query OR in JSON body. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdsAPI.DeleteConversionDestination(context.Background(), accountId, destinationId).AdAccountId(adAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.DeleteConversionDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConversionDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **adAccountId** | **string** | Required as query OR in JSON body. | 

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


## EstimateAdReach

> EstimateAdReach200Response EstimateAdReach(ctx).EstimateAdReachRequest(estimateAdReachRequest).Execute()

Estimate audience reach



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
	estimateAdReachRequest := *openapiclient.NewEstimateAdReachRequest("AccountId_example", "AdAccountId_example", *openapiclient.NewTargetingSpec()) // EstimateAdReachRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.EstimateAdReach(context.Background()).EstimateAdReachRequest(estimateAdReachRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.EstimateAdReach``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EstimateAdReach`: EstimateAdReach200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.EstimateAdReach`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEstimateAdReachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **estimateAdReachRequest** | [**EstimateAdReachRequest**](EstimateAdReachRequest.md) |  | 

### Return type

[**EstimateAdReach200Response**](EstimateAdReach200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAd

> GetAd200Response GetAd(ctx, adId).Execute()

Get ad details



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
	adId := "adId_example" // string | Zernio `_id` (hex), Meta `platformAdId` (numeric), or one of the creative's effective story/media IDs. See description for details. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.GetAd(context.Background(), adId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.GetAd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAd`: GetAd200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.GetAd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adId** | **string** | Zernio &#x60;_id&#x60; (hex), Meta &#x60;platformAdId&#x60; (numeric), or one of the creative&#39;s effective story/media IDs. See description for details.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAd200Response**](GetAd200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdAnalytics

> GetAdAnalytics200Response GetAdAnalytics(ctx, adId).FromDate(fromDate).ToDate(toDate).Breakdowns(breakdowns).Execute()

Get ad analytics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/zernio-dev/zernio-go"
)

func main() {
	adId := "adId_example" // string | 
	fromDate := time.Now() // string | Start of date range (YYYY-MM-DD). Defaults to 90 days ago. (optional)
	toDate := time.Now() // string | End of date range (YYYY-MM-DD). Defaults to today. Max 730-day range. (optional)
	breakdowns := "breakdowns_example" // string | Comma-separated breakdown dimensions. Meta: age, gender, country, publisher_platform, device_platform, region. TikTok: gender, age, country_code, platform, ac, language. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.GetAdAnalytics(context.Background(), adId).FromDate(fromDate).ToDate(toDate).Breakdowns(breakdowns).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.GetAdAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdAnalytics`: GetAdAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.GetAdAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | Start of date range (YYYY-MM-DD). Defaults to 90 days ago. | 
 **toDate** | **string** | End of date range (YYYY-MM-DD). Defaults to today. Max 730-day range. | 
 **breakdowns** | **string** | Comma-separated breakdown dimensions. Meta: age, gender, country, publisher_platform, device_platform, region. TikTok: gender, age, country_code, platform, ac, language. | 

### Return type

[**GetAdAnalytics200Response**](GetAdAnalytics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdComments

> GetAdComments200Response GetAdComments(ctx, adId).Placement(placement).Limit(limit).Cursor(cursor).Execute()

List comments on an ad



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
	adId := "adId_example" // string | Internal Zernio ad ID (ObjectId).
	placement := "placement_example" // string | Which side of the ad to return comments for. Omit to default to the Instagram side when present, else Facebook. Returns ad_not_commentable if the ad has no such placement. (optional)
	limit := int32(56) // int32 |  (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor from a previous response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.GetAdComments(context.Background(), adId).Placement(placement).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.GetAdComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdComments`: GetAdComments200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.GetAdComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adId** | **string** | Internal Zernio ad ID (ObjectId). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **placement** | **string** | Which side of the ad to return comments for. Omit to default to the Instagram side when present, else Facebook. Returns ad_not_commentable if the ad has no such placement. | 
 **limit** | **int32** |  | [default to 25]
 **cursor** | **string** | Pagination cursor from a previous response. | 

### Return type

[**GetAdComments200Response**](GetAdComments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdTrackingTags

> GetAdTrackingTags200Response GetAdTrackingTags(ctx, adId).Execute()

Read an ad's click-URL tracking tags



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
	adId := "adId_example" // string | Ad id (hex _id, platformAdId, or effective story/media id).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.GetAdTrackingTags(context.Background(), adId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.GetAdTrackingTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdTrackingTags`: GetAdTrackingTags200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.GetAdTrackingTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adId** | **string** | Ad id (hex _id, platformAdId, or effective story/media id). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdTrackingTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAdTrackingTags200Response**](GetAdTrackingTags200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversionDestination

> GetConversionDestination200Response GetConversionDestination(ctx, accountId, destinationId).AdAccountId(adAccountId).Execute()

Fetch a single conversion destination



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
	accountId := "accountId_example" // string | 
	destinationId := "destinationId_example" // string | 
	adAccountId := "adAccountId_example" // string | Numeric ID or full `urn:li:sponsoredAccount:{id}` URN.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.GetConversionDestination(context.Background(), accountId, destinationId).AdAccountId(adAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.GetConversionDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversionDestination`: GetConversionDestination200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.GetConversionDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversionDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **adAccountId** | **string** | Numeric ID or full &#x60;urn:li:sponsoredAccount:{id}&#x60; URN. | 

### Return type

[**GetConversionDestination200Response**](GetConversionDestination200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversionMetrics

> GetConversionMetrics200Response GetConversionMetrics(ctx, accountId, destinationId).AdAccountId(adAccountId).StartDate(startDate).EndDate(endDate).Granularity(granularity).Execute()

Fetch attribution metrics for a conversion destination



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
	accountId := "accountId_example" // string | 
	destinationId := "destinationId_example" // string | 
	adAccountId := "adAccountId_example" // string | 
	startDate := "startDate_example" // string | 
	endDate := "endDate_example" // string |  (optional)
	granularity := "granularity_example" // string |  (optional) (default to "DAILY")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.GetConversionMetrics(context.Background(), accountId, destinationId).AdAccountId(adAccountId).StartDate(startDate).EndDate(endDate).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.GetConversionMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversionMetrics`: GetConversionMetrics200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.GetConversionMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversionMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **adAccountId** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **granularity** | **string** |  | [default to &quot;DAILY&quot;]

### Return type

[**GetConversionMetrics200Response**](GetConversionMetrics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversionsQuality

> GetConversionsQuality200Response GetConversionsQuality(ctx).AccountId(accountId).DestinationId(destinationId).Execute()

Read Event Match Quality + coverage for a Meta pixel



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
	accountId := "accountId_example" // string | SocialAccount _id (must be a metaads account).
	destinationId := "destinationId_example" // string | Meta pixel/dataset ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.GetConversionsQuality(context.Background()).AccountId(accountId).DestinationId(destinationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.GetConversionsQuality``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversionsQuality`: GetConversionsQuality200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.GetConversionsQuality`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConversionsQualityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | SocialAccount _id (must be a metaads account). | 
 **destinationId** | **string** | Meta pixel/dataset ID. | 

### Return type

[**GetConversionsQuality200Response**](GetConversionsQuality200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeadForm

> GetLeadForm200Response GetLeadForm(ctx, formId).AccountId(accountId).Execute()

Get a single Lead Gen form

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
	formId := "formId_example" // string | 
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.GetLeadForm(context.Background(), formId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.GetLeadForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeadForm`: GetLeadForm200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.GetLeadForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeadFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** |  | 

### Return type

[**GetLeadForm200Response**](GetLeadForm200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAdAccounts

> ListAdAccounts200Response ListAdAccounts(ctx).AccountId(accountId).AdAccountId(adAccountId).Limit(limit).Execute()

List ad accounts



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
	adAccountId := "adAccountId_example" // string | Filter response to a single platform ad account ID (e.g. `act_123` for Meta, advertiser_id for TikTok). Returns at most one item. (optional)
	limit := int32(56) // int32 | Clamp the returned `accounts[]` length. Useful for typeahead pickers on agency tokens with hundreds of advertisers. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.ListAdAccounts(context.Background()).AccountId(accountId).AdAccountId(adAccountId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.ListAdAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAdAccounts`: ListAdAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.ListAdAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAdAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Social account ID | 
 **adAccountId** | **string** | Filter response to a single platform ad account ID (e.g. &#x60;act_123&#x60; for Meta, advertiser_id for TikTok). Returns at most one item. | 
 **limit** | **int32** | Clamp the returned &#x60;accounts[]&#x60; length. Useful for typeahead pickers on agency tokens with hundreds of advertisers. | 

### Return type

[**ListAdAccounts200Response**](ListAdAccounts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAdCatalogProductSets

> ListAdCatalogProductSets200Response ListAdCatalogProductSets(ctx, catalogId).AccountId(accountId).Execute()

List a catalog's product sets



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
	catalogId := "catalogId_example" // string | Meta product catalog ID (from GET /v1/ads/catalogs)
	accountId := "accountId_example" // string | A facebook, instagram, or metaads social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.ListAdCatalogProductSets(context.Background(), catalogId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.ListAdCatalogProductSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAdCatalogProductSets`: ListAdCatalogProductSets200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.ListAdCatalogProductSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogId** | **string** | Meta product catalog ID (from GET /v1/ads/catalogs) | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAdCatalogProductSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | A facebook, instagram, or metaads social account ID | 

### Return type

[**ListAdCatalogProductSets200Response**](ListAdCatalogProductSets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAdCatalogs

> ListAdCatalogs200Response ListAdCatalogs(ctx).AccountId(accountId).AdAccountId(adAccountId).Execute()

List Meta product catalogs



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
	accountId := "accountId_example" // string | A facebook, instagram, or metaads social account ID
	adAccountId := "adAccountId_example" // string | Meta ad account ID (act_...)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.ListAdCatalogs(context.Background()).AccountId(accountId).AdAccountId(adAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.ListAdCatalogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAdCatalogs`: ListAdCatalogs200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.ListAdCatalogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAdCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | A facebook, instagram, or metaads social account ID | 
 **adAccountId** | **string** | Meta ad account ID (act_...) | 

### Return type

[**ListAdCatalogs200Response**](ListAdCatalogs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAds

> ListAds200Response ListAds(ctx).Page(page).Limit(limit).Source(source).Status(status).Platform(platform).AccountId(accountId).AdAccountId(adAccountId).ProfileId(profileId).CampaignId(campaignId).PlatformAdId(platformAdId).EffectiveObjectStoryId(effectiveObjectStoryId).EffectiveInstagramMediaId(effectiveInstagramMediaId).FromDate(fromDate).ToDate(toDate).Execute()

List ads



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/zernio-dev/zernio-go"
)

func main() {
	page := int32(56) // int32 | Page number (1-based) (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 50)
	source := "source_example" // string | all (default) = Zernio-created + platform-discovered ads. zernio = restrict to Zernio-created only. (optional) (default to "all")
	status := openapiclient.AdStatus("active") // AdStatus |  (optional)
	platform := "platform_example" // string |  (optional)
	accountId := "accountId_example" // string | Social account ID (optional)
	adAccountId := "adAccountId_example" // string | Platform ad account ID (e.g. act_123 for Meta). Mirrors the same filter on /v1/ads/campaigns and /v1/ads/tree. (optional)
	profileId := "profileId_example" // string | Profile ID (optional)
	campaignId := "campaignId_example" // string | Platform campaign ID (filter ads within a campaign) (optional)
	platformAdId := "platformAdId_example" // string | Meta ad ID. Returns the ad with this platform-side ad ID. (optional)
	effectiveObjectStoryId := "effectiveObjectStoryId_example" // string | Facebook `{pageId}_{postId}` of the post the ad's engagement lives on (Meta `effective_object_story_id`). Use to map a Business-Manager-visible post back to the Zernio ad. (optional)
	effectiveInstagramMediaId := "effectiveInstagramMediaId_example" // string | Instagram media ID of the boosted post (Meta `effective_instagram_media_id`). Use to map a Business-Manager-visible IG post back to the Zernio ad. (optional)
	fromDate := time.Now() // string | Start of metrics date range (YYYY-MM-DD). Defaults to 90 days ago. (optional)
	toDate := time.Now() // string | End of metrics date range (YYYY-MM-DD). Defaults to today. Max 730-day range. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.ListAds(context.Background()).Page(page).Limit(limit).Source(source).Status(status).Platform(platform).AccountId(accountId).AdAccountId(adAccountId).ProfileId(profileId).CampaignId(campaignId).PlatformAdId(platformAdId).EffectiveObjectStoryId(effectiveObjectStoryId).EffectiveInstagramMediaId(effectiveInstagramMediaId).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.ListAds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAds`: ListAds200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.ListAds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-based) | [default to 1]
 **limit** | **int32** |  | [default to 50]
 **source** | **string** | all (default) &#x3D; Zernio-created + platform-discovered ads. zernio &#x3D; restrict to Zernio-created only. | [default to &quot;all&quot;]
 **status** | [**AdStatus**](AdStatus.md) |  | 
 **platform** | **string** |  | 
 **accountId** | **string** | Social account ID | 
 **adAccountId** | **string** | Platform ad account ID (e.g. act_123 for Meta). Mirrors the same filter on /v1/ads/campaigns and /v1/ads/tree. | 
 **profileId** | **string** | Profile ID | 
 **campaignId** | **string** | Platform campaign ID (filter ads within a campaign) | 
 **platformAdId** | **string** | Meta ad ID. Returns the ad with this platform-side ad ID. | 
 **effectiveObjectStoryId** | **string** | Facebook &#x60;{pageId}_{postId}&#x60; of the post the ad&#39;s engagement lives on (Meta &#x60;effective_object_story_id&#x60;). Use to map a Business-Manager-visible post back to the Zernio ad. | 
 **effectiveInstagramMediaId** | **string** | Instagram media ID of the boosted post (Meta &#x60;effective_instagram_media_id&#x60;). Use to map a Business-Manager-visible IG post back to the Zernio ad. | 
 **fromDate** | **string** | Start of metrics date range (YYYY-MM-DD). Defaults to 90 days ago. | 
 **toDate** | **string** | End of metrics date range (YYYY-MM-DD). Defaults to today. Max 730-day range. | 

### Return type

[**ListAds200Response**](ListAds200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAdsBusinessCenters

> ListAdsBusinessCenters200Response ListAdsBusinessCenters(ctx).AccountId(accountId).Execute()

List TikTok Business Centers



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
	accountId := "accountId_example" // string | ID of the `tiktokads` (or parent `tiktok` posting) SocialAccount

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.ListAdsBusinessCenters(context.Background()).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.ListAdsBusinessCenters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAdsBusinessCenters`: ListAdsBusinessCenters200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.ListAdsBusinessCenters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAdsBusinessCentersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | ID of the &#x60;tiktokads&#x60; (or parent &#x60;tiktok&#x60; posting) SocialAccount | 

### Return type

[**ListAdsBusinessCenters200Response**](ListAdsBusinessCenters200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversionAssociations

> ListConversionAssociations200Response ListConversionAssociations(ctx, accountId, destinationId).AdAccountId(adAccountId).Execute()

List campaigns associated with a conversion destination



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
	accountId := "accountId_example" // string | 
	destinationId := "destinationId_example" // string | 
	adAccountId := "adAccountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.ListConversionAssociations(context.Background(), accountId, destinationId).AdAccountId(adAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.ListConversionAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConversionAssociations`: ListConversionAssociations200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.ListConversionAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConversionAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **adAccountId** | **string** |  | 

### Return type

[**ListConversionAssociations200Response**](ListConversionAssociations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversionDestinations

> ListConversionDestinations200Response ListConversionDestinations(ctx, accountId).Execute()

List destinations for the Conversions API



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
	accountId := "accountId_example" // string | SocialAccount ID (metaads, googleads, linkedinads, or tiktokads).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.ListConversionDestinations(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.ListConversionDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConversionDestinations`: ListConversionDestinations200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.ListConversionDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | SocialAccount ID (metaads, googleads, linkedinads, or tiktokads). | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConversionDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListConversionDestinations200Response**](ListConversionDestinations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFormLeads

> ListFormLeads200Response ListFormLeads(ctx, formId).AccountId(accountId).Limit(limit).Cursor(cursor).Since(since).Execute()

List leads for a single form



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
	formId := "formId_example" // string | 
	accountId := "accountId_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 25)
	cursor := "cursor_example" // string |  (optional)
	since := int32(56) // int32 | Unix seconds. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.ListFormLeads(context.Background(), formId).AccountId(accountId).Limit(limit).Cursor(cursor).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.ListFormLeads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFormLeads`: ListFormLeads200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.ListFormLeads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFormLeadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** |  | 
 **limit** | **int32** |  | [default to 25]
 **cursor** | **string** |  | 
 **since** | **int32** | Unix seconds. | 

### Return type

[**ListFormLeads200Response**](ListFormLeads200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLeadForms

> ListLeadForms200Response ListLeadForms(ctx).AccountId(accountId).Limit(limit).Cursor(cursor).Execute()

List Lead Gen (Instant) forms



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
	accountId := "accountId_example" // string | Connected facebook account id.
	limit := int32(56) // int32 |  (optional) (default to 25)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.ListLeadForms(context.Background()).AccountId(accountId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.ListLeadForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLeadForms`: ListLeadForms200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.ListLeadForms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLeadFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Connected facebook account id. | 
 **limit** | **int32** |  | [default to 25]
 **cursor** | **string** |  | 

### Return type

[**ListLeadForms200Response**](ListLeadForms200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLeads

> ListLeads200Response ListLeads(ctx).FormId(formId).AccountId(accountId).Limit(limit).Since(since).Cursor(cursor).Execute()

List submitted leads (cross-form CRM view)



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
	formId := "formId_example" // string | Filter to a single lead form. (optional)
	accountId := "accountId_example" // string | Filter to a single connected account. (optional)
	limit := int32(56) // int32 |  (optional) (default to 25)
	since := int32(56) // int32 | Unix seconds; only leads created at/after this Meta timestamp. (optional)
	cursor := "cursor_example" // string | Keyset cursor from a previous response's pagination.cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.ListLeads(context.Background()).FormId(formId).AccountId(accountId).Limit(limit).Since(since).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.ListLeads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLeads`: ListLeads200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.ListLeads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLeadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **formId** | **string** | Filter to a single lead form. | 
 **accountId** | **string** | Filter to a single connected account. | 
 **limit** | **int32** |  | [default to 25]
 **since** | **int32** | Unix seconds; only leads created at/after this Meta timestamp. | 
 **cursor** | **string** | Keyset cursor from a previous response&#39;s pagination.cursor. | 

### Return type

[**ListLeads200Response**](ListLeads200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWhatsAppConversions

> ListWhatsAppConversions200Response ListWhatsAppConversions(ctx).AccountId(accountId).Limit(limit).Execute()

List recent WhatsApp conversion events



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
	limit := int32(56) // int32 | Max events to return (1-200, default 50). (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.ListWhatsAppConversions(context.Background()).AccountId(accountId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.ListWhatsAppConversions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWhatsAppConversions`: ListWhatsAppConversions200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.ListWhatsAppConversions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWhatsAppConversionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | WhatsApp social account ID | 
 **limit** | **int32** | Max events to return (1-200, default 50). | [default to 50]

### Return type

[**ListWhatsAppConversions200Response**](ListWhatsAppConversions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveConversionAssociations

> AddConversionAssociations200Response RemoveConversionAssociations(ctx, accountId, destinationId).AdAccountId(adAccountId).CampaignIds(campaignIds).Execute()

Remove campaign↔conversion associations



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
	accountId := "accountId_example" // string | 
	destinationId := "destinationId_example" // string | 
	adAccountId := "adAccountId_example" // string | 
	campaignIds := "campaignIds_example" // string | Comma-separated list of campaign IDs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.RemoveConversionAssociations(context.Background(), accountId, destinationId).AdAccountId(adAccountId).CampaignIds(campaignIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.RemoveConversionAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveConversionAssociations`: AddConversionAssociations200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.RemoveConversionAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveConversionAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **adAccountId** | **string** |  | 
 **campaignIds** | **string** | Comma-separated list of campaign IDs. | 

### Return type

[**AddConversionAssociations200Response**](AddConversionAssociations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAdInterests

> SearchAdInterests200Response SearchAdInterests(ctx).Q(q).AccountId(accountId).Execute()

Search targeting interests (deprecated)



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
	q := "q_example" // string | Search query
	accountId := "accountId_example" // string | Social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.SearchAdInterests(context.Background()).Q(q).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.SearchAdInterests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAdInterests`: SearchAdInterests200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.SearchAdInterests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAdInterestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Search query | 
 **accountId** | **string** | Social account ID | 

### Return type

[**SearchAdInterests200Response**](SearchAdInterests200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAdTargeting

> SearchAdTargeting200Response SearchAdTargeting(ctx).AccountId(accountId).Q(q).Dimension(dimension).GeoType(geoType).CountryCode(countryCode).Limit(limit).Execute()

Search targeting options



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
	accountId := "accountId_example" // string | Social account ID (a connected account on the target ad platform).
	q := "q_example" // string | Search query. For geo, the locality name only (no region/country suffix).
	dimension := "dimension_example" // string | What to search. `geo` resolves locations (scope further with `geoType`), `interest`/`behavior` resolve audience entities, `income` resolves income-tier options. Defaults to `interest` for backward compatibility with the deprecated /v1/ads/interests alias. (optional) (default to "interest")
	geoType := "geoType_example" // string | Only used when `dimension=geo`. The kind of location to resolve. Defaults to `city`. (optional) (default to "city")
	countryCode := "countryCode_example" // string | ISO 3166-1 alpha-2 country code (e.g. NL) to scope a geo search. (optional)
	limit := int32(56) // int32 | Maximum results to return. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.SearchAdTargeting(context.Background()).AccountId(accountId).Q(q).Dimension(dimension).GeoType(geoType).CountryCode(countryCode).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.SearchAdTargeting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAdTargeting`: SearchAdTargeting200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.SearchAdTargeting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAdTargetingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Social account ID (a connected account on the target ad platform). | 
 **q** | **string** | Search query. For geo, the locality name only (no region/country suffix). | 
 **dimension** | **string** | What to search. &#x60;geo&#x60; resolves locations (scope further with &#x60;geoType&#x60;), &#x60;interest&#x60;/&#x60;behavior&#x60; resolve audience entities, &#x60;income&#x60; resolves income-tier options. Defaults to &#x60;interest&#x60; for backward compatibility with the deprecated /v1/ads/interests alias. | [default to &quot;interest&quot;]
 **geoType** | **string** | Only used when &#x60;dimension&#x3D;geo&#x60;. The kind of location to resolve. Defaults to &#x60;city&#x60;. | [default to &quot;city&quot;]
 **countryCode** | **string** | ISO 3166-1 alpha-2 country code (e.g. NL) to scope a geo search. | 
 **limit** | **int32** | Maximum results to return. | [default to 25]

### Return type

[**SearchAdTargeting200Response**](SearchAdTargeting200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendConversions

> SendConversions200Response SendConversions(ctx).SendConversionsRequest(sendConversionsRequest).Execute()

Send conversion events to an ad platform



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
	sendConversionsRequest := *openapiclient.NewSendConversionsRequest("AccountId_example", "DestinationId_example", []openapiclient.ConversionEvent{*openapiclient.NewConversionEvent("Purchase", int32(1744732800), "order_abc_123", *openapiclient.NewConversionEventUser())}) // SendConversionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.SendConversions(context.Background()).SendConversionsRequest(sendConversionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.SendConversions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendConversions`: SendConversions200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.SendConversions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendConversionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendConversionsRequest** | [**SendConversionsRequest**](SendConversionsRequest.md) |  | 

### Return type

[**SendConversions200Response**](SendConversions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendWhatsAppConversion

> SendWhatsAppConversion200Response SendWhatsAppConversion(ctx).SendWhatsAppConversionRequest(sendWhatsAppConversionRequest).Execute()

Send WhatsApp conversion event



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
	sendWhatsAppConversionRequest := *openapiclient.NewSendWhatsAppConversionRequest("AccountId_example", "EventName_example", "EventId_example") // SendWhatsAppConversionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.SendWhatsAppConversion(context.Background()).SendWhatsAppConversionRequest(sendWhatsAppConversionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.SendWhatsAppConversion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendWhatsAppConversion`: SendWhatsAppConversion200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.SendWhatsAppConversion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendWhatsAppConversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendWhatsAppConversionRequest** | [**SendWhatsAppConversionRequest**](SendWhatsAppConversionRequest.md) |  | 

### Return type

[**SendWhatsAppConversion200Response**](SendWhatsAppConversion200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAd

> UpdateAd200Response UpdateAd(ctx, adId).UpdateAdRequest(updateAdRequest).Execute()

Update ad



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
	adId := "adId_example" // string | 
	updateAdRequest := *openapiclient.NewUpdateAdRequest() // UpdateAdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.UpdateAd(context.Background(), adId).UpdateAdRequest(updateAdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.UpdateAd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAd`: UpdateAd200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.UpdateAd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAdRequest** | [**UpdateAdRequest**](UpdateAdRequest.md) |  | 

### Return type

[**UpdateAd200Response**](UpdateAd200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdTrackingTags

> UpdateAdTrackingTags(ctx, adId).UpdateAdTrackingTagsRequest(updateAdTrackingTagsRequest).Execute()

Set/update an ad's click-URL tracking tags



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
	adId := "adId_example" // string | 
	updateAdTrackingTagsRequest := *openapiclient.NewUpdateAdTrackingTagsRequest() // UpdateAdTrackingTagsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdsAPI.UpdateAdTrackingTags(context.Background(), adId).UpdateAdTrackingTagsRequest(updateAdTrackingTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.UpdateAdTrackingTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdTrackingTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAdTrackingTagsRequest** | [**UpdateAdTrackingTagsRequest**](UpdateAdTrackingTagsRequest.md) |  | 

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


## UpdateConversionDestination

> GetConversionDestination200Response UpdateConversionDestination(ctx, accountId, destinationId).UpdateConversionDestinationRequest(updateConversionDestinationRequest).Execute()

Update a conversion destination



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
	accountId := "accountId_example" // string | 
	destinationId := "destinationId_example" // string | 
	updateConversionDestinationRequest := *openapiclient.NewUpdateConversionDestinationRequest("AdAccountId_example") // UpdateConversionDestinationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdsAPI.UpdateConversionDestination(context.Background(), accountId, destinationId).UpdateConversionDestinationRequest(updateConversionDestinationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.UpdateConversionDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConversionDestination`: GetConversionDestination200Response
	fmt.Fprintf(os.Stdout, "Response from `AdsAPI.UpdateConversionDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConversionDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateConversionDestinationRequest** | [**UpdateConversionDestinationRequest**](UpdateConversionDestinationRequest.md) |  | 

### Return type

[**GetConversionDestination200Response**](GetConversionDestination200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

