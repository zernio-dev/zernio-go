# \UsageAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsageStats**](UsageAPI.md#GetUsageStats) | **Get** /v1/usage-stats | Get plan and usage stats
[**GetXApiPricing**](UsageAPI.md#GetXApiPricing) | **Get** /v1/billing/x-pricing | Get X/Twitter API pricing table



## GetUsageStats

> UsageStats GetUsageStats(ctx).Reconcile(reconcile).Execute()

Get plan and usage stats



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
	reconcile := true // bool | For Stripe subscription users, `true` forces a subscription reconciliation pass even when cached plan data looks complete. Omit the parameter, or pass `false`, to use the default first-time-only reconciliation behavior. Invalid boolean values are rejected.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.GetUsageStats(context.Background()).Reconcile(reconcile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetUsageStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageStats`: UsageStats
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.GetUsageStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reconcile** | **bool** | For Stripe subscription users, &#x60;true&#x60; forces a subscription reconciliation pass even when cached plan data looks complete. Omit the parameter, or pass &#x60;false&#x60;, to use the default first-time-only reconciliation behavior. Invalid boolean values are rejected.  | 

### Return type

[**UsageStats**](UsageStats.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetXApiPricing

> XApiPricing GetXApiPricing(ctx).Execute()

Get X/Twitter API pricing table



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
	resp, r, err := apiClient.UsageAPI.GetXApiPricing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetXApiPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetXApiPricing`: XApiPricing
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.GetXApiPricing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetXApiPricingRequest struct via the builder pattern


### Return type

[**XApiPricing**](XApiPricing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

