# \AccountsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccount**](AccountsAPI.md#DeleteAccount) | **Delete** /v1/accounts/{accountId} | Disconnect account
[**GetAccountHealth**](AccountsAPI.md#GetAccountHealth) | **Get** /v1/accounts/{accountId}/health | Check account health
[**GetAllAccountsHealth**](AccountsAPI.md#GetAllAccountsHealth) | **Get** /v1/accounts/health | Check accounts health
[**GetFollowerStats**](AccountsAPI.md#GetFollowerStats) | **Get** /v1/accounts/follower-stats | Get follower stats
[**GetTikTokCreatorInfo**](AccountsAPI.md#GetTikTokCreatorInfo) | **Get** /v1/accounts/{accountId}/tiktok/creator-info | Get TikTok creator info
[**ListAccounts**](AccountsAPI.md#ListAccounts) | **Get** /v1/accounts | List accounts
[**MoveAccountToProfile**](AccountsAPI.md#MoveAccountToProfile) | **Patch** /v1/accounts/{accountId} | Move account to a different profile
[**UpdateAccount**](AccountsAPI.md#UpdateAccount) | **Put** /v1/accounts/{accountId} | Update account



## DeleteAccount

> DeleteAccountGroup200Response DeleteAccount(ctx, accountId).Execute()

Disconnect account



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.DeleteAccount(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccount`: DeleteAccountGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


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


## GetAccountHealth

> GetAccountHealth200Response GetAccountHealth(ctx, accountId).Execute()

Check account health



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
	accountId := "accountId_example" // string | The account ID to check

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountHealth(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountHealth`: GetAccountHealth200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The account ID to check | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAccountHealth200Response**](GetAccountHealth200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAccountsHealth

> GetAllAccountsHealth200Response GetAllAccountsHealth(ctx).ProfileId(profileId).Platform(platform).Status(status).Execute()

Check accounts health



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
	profileId := "profileId_example" // string | Filter by profile ID (optional)
	platform := "platform_example" // string | Filter by platform (optional)
	status := "status_example" // string | Filter by health status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAllAccountsHealth(context.Background()).ProfileId(profileId).Platform(platform).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAllAccountsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAccountsHealth`: GetAllAccountsHealth200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAllAccountsHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAccountsHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Filter by profile ID | 
 **platform** | **string** | Filter by platform | 
 **status** | **string** | Filter by health status | 

### Return type

[**GetAllAccountsHealth200Response**](GetAllAccountsHealth200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFollowerStats

> GetFollowerStats200Response GetFollowerStats(ctx).AccountIds(accountIds).ProfileId(profileId).FromDate(fromDate).ToDate(toDate).Granularity(granularity).Execute()

Get follower stats



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
	accountIds := "accountIds_example" // string | Comma-separated list of account IDs (optional, defaults to all user's accounts) (optional)
	profileId := "profileId_example" // string | Filter by profile ID (optional)
	fromDate := time.Now() // string | Start date in YYYY-MM-DD format (defaults to 30 days ago) (optional)
	toDate := time.Now() // string | End date in YYYY-MM-DD format (defaults to today) (optional)
	granularity := "granularity_example" // string | Data aggregation level (optional) (default to "daily")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetFollowerStats(context.Background()).AccountIds(accountIds).ProfileId(profileId).FromDate(fromDate).ToDate(toDate).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetFollowerStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFollowerStats`: GetFollowerStats200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetFollowerStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowerStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIds** | **string** | Comma-separated list of account IDs (optional, defaults to all user&#39;s accounts) | 
 **profileId** | **string** | Filter by profile ID | 
 **fromDate** | **string** | Start date in YYYY-MM-DD format (defaults to 30 days ago) | 
 **toDate** | **string** | End date in YYYY-MM-DD format (defaults to today) | 
 **granularity** | **string** | Data aggregation level | [default to &quot;daily&quot;]

### Return type

[**GetFollowerStats200Response**](GetFollowerStats200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTikTokCreatorInfo

> GetTikTokCreatorInfo200Response GetTikTokCreatorInfo(ctx, accountId).MediaType(mediaType).Execute()

Get TikTok creator info



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
	accountId := "accountId_example" // string | The TikTok account ID
	mediaType := "mediaType_example" // string | The media type to get creator info for (affects available interaction settings) (optional) (default to "video")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetTikTokCreatorInfo(context.Background(), accountId).MediaType(mediaType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetTikTokCreatorInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTikTokCreatorInfo`: GetTikTokCreatorInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetTikTokCreatorInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The TikTok account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTikTokCreatorInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaType** | **string** | The media type to get creator info for (affects available interaction settings) | [default to &quot;video&quot;]

### Return type

[**GetTikTokCreatorInfo200Response**](GetTikTokCreatorInfo200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> ListAccounts200Response ListAccounts(ctx).ProfileId(profileId).Platform(platform).Status(status).IncludeOverLimit(includeOverLimit).Page(page).Limit(limit).Execute()

List accounts



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
	profileId := "profileId_example" // string | Filter accounts by profile ID (optional)
	platform := "platform_example" // string | Filter accounts by platform (e.g. \"instagram\", \"twitter\"). (optional)
	status := "status_example" // string | Filter accounts by connection status. `connected` returns healthy accounts; `disconnected` returns accounts that need reconnection (per the same reconnection check surfaced in the dashboard). Omit to return accounts in any status. When combined with page/limit, pagination totals reflect the filtered result set.  (optional)
	includeOverLimit := true // bool | When true, includes accounts from over-limit profiles. (optional) (default to false)
	page := int32(56) // int32 | Page number (1-based). When provided with limit, enables server-side pagination. Omit for all accounts. (optional)
	limit := int32(56) // int32 | Page size. Required alongside page for pagination. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ListAccounts(context.Background()).ProfileId(profileId).Platform(platform).Status(status).IncludeOverLimit(includeOverLimit).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: ListAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Filter accounts by profile ID | 
 **platform** | **string** | Filter accounts by platform (e.g. \&quot;instagram\&quot;, \&quot;twitter\&quot;). | 
 **status** | **string** | Filter accounts by connection status. &#x60;connected&#x60; returns healthy accounts; &#x60;disconnected&#x60; returns accounts that need reconnection (per the same reconnection check surfaced in the dashboard). Omit to return accounts in any status. When combined with page/limit, pagination totals reflect the filtered result set.  | 
 **includeOverLimit** | **bool** | When true, includes accounts from over-limit profiles. | [default to false]
 **page** | **int32** | Page number (1-based). When provided with limit, enables server-side pagination. Omit for all accounts. | 
 **limit** | **int32** | Page size. Required alongside page for pagination. | 

### Return type

[**ListAccounts200Response**](ListAccounts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveAccountToProfile

> MoveAccountToProfile200Response MoveAccountToProfile(ctx, accountId).MoveAccountToProfileRequest(moveAccountToProfileRequest).Execute()

Move account to a different profile



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
	moveAccountToProfileRequest := *openapiclient.NewMoveAccountToProfileRequest("ProfileId_example") // MoveAccountToProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.MoveAccountToProfile(context.Background(), accountId).MoveAccountToProfileRequest(moveAccountToProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.MoveAccountToProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveAccountToProfile`: MoveAccountToProfile200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.MoveAccountToProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveAccountToProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveAccountToProfileRequest** | [**MoveAccountToProfileRequest**](MoveAccountToProfileRequest.md) |  | 

### Return type

[**MoveAccountToProfile200Response**](MoveAccountToProfile200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> UpdateAccount200Response UpdateAccount(ctx, accountId).UpdateAccountRequest(updateAccountRequest).Execute()

Update account



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
	updateAccountRequest := *openapiclient.NewUpdateAccountRequest() // UpdateAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.UpdateAccount(context.Background(), accountId).UpdateAccountRequest(updateAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.UpdateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccount`: UpdateAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAccountRequest** | [**UpdateAccountRequest**](UpdateAccountRequest.md) |  | 

### Return type

[**UpdateAccount200Response**](UpdateAccount200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

