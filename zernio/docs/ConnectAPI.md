# \ConnectAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteTelegramConnect**](ConnectAPI.md#CompleteTelegramConnect) | **Patch** /v1/connect/telegram | Check Telegram status
[**CompleteWhatsAppPhoneSelection**](ConnectAPI.md#CompleteWhatsAppPhoneSelection) | **Post** /v1/connect/whatsapp/select-phone-number | Complete WhatsApp phone number selection
[**ConfigureTikTokAdsBrandIdentity**](ConnectAPI.md#ConfigureTikTokAdsBrandIdentity) | **Patch** /v1/connect/tiktok-ads | Configure TikTok Ads Brand Identity
[**ConnectAds**](ConnectAPI.md#ConnectAds) | **Get** /v1/connect/{platform}/ads | Connect ads for a platform
[**ConnectBlueskyCredentials**](ConnectAPI.md#ConnectBlueskyCredentials) | **Post** /v1/connect/bluesky/credentials | Connect Bluesky account
[**ConnectWhatsAppCredentials**](ConnectAPI.md#ConnectWhatsAppCredentials) | **Post** /v1/connect/whatsapp/credentials | Connect WhatsApp via credentials
[**GetConnectUrl**](ConnectAPI.md#GetConnectUrl) | **Get** /v1/connect/{platform} | Get OAuth connect URL
[**GetFacebookPages**](ConnectAPI.md#GetFacebookPages) | **Get** /v1/accounts/{accountId}/facebook-page | List Facebook pages
[**GetGmbLocations**](ConnectAPI.md#GetGmbLocations) | **Get** /v1/accounts/{accountId}/gmb-locations | List GBP locations
[**GetLinkedInOrganizations**](ConnectAPI.md#GetLinkedInOrganizations) | **Get** /v1/accounts/{accountId}/linkedin-organizations | List LinkedIn orgs
[**GetPendingOAuthData**](ConnectAPI.md#GetPendingOAuthData) | **Get** /v1/connect/pending-data | Get pending OAuth data
[**GetPinterestBoards**](ConnectAPI.md#GetPinterestBoards) | **Get** /v1/accounts/{accountId}/pinterest-boards | List Pinterest boards
[**GetRedditFlairs**](ConnectAPI.md#GetRedditFlairs) | **Get** /v1/accounts/{accountId}/reddit-flairs | List subreddit flairs
[**GetRedditSubreddits**](ConnectAPI.md#GetRedditSubreddits) | **Get** /v1/accounts/{accountId}/reddit-subreddits | List Reddit subreddits
[**GetTelegramConnectStatus**](ConnectAPI.md#GetTelegramConnectStatus) | **Get** /v1/connect/telegram | Generate Telegram code
[**GetYoutubePlaylists**](ConnectAPI.md#GetYoutubePlaylists) | **Get** /v1/accounts/{accountId}/youtube-playlists | List YouTube playlists
[**HandleOAuthCallback**](ConnectAPI.md#HandleOAuthCallback) | **Post** /v1/connect/{platform} | Complete OAuth callback
[**InitiateTelegramConnect**](ConnectAPI.md#InitiateTelegramConnect) | **Post** /v1/connect/telegram | Connect Telegram directly
[**ListFacebookPages**](ConnectAPI.md#ListFacebookPages) | **Get** /v1/connect/facebook/select-page | List Facebook pages
[**ListGoogleBusinessLocations**](ConnectAPI.md#ListGoogleBusinessLocations) | **Get** /v1/connect/googlebusiness/locations | List GBP locations
[**ListLinkedInOrganizations**](ConnectAPI.md#ListLinkedInOrganizations) | **Get** /v1/connect/linkedin/organizations | List LinkedIn orgs
[**ListPinterestBoardsForSelection**](ConnectAPI.md#ListPinterestBoardsForSelection) | **Get** /v1/connect/pinterest/select-board | List Pinterest boards
[**ListSnapchatProfiles**](ConnectAPI.md#ListSnapchatProfiles) | **Get** /v1/connect/snapchat/select-profile | List Snapchat profiles
[**ListWhatsAppPhoneNumbers**](ConnectAPI.md#ListWhatsAppPhoneNumbers) | **Get** /v1/connect/whatsapp/select-phone-number | List WhatsApp phone numbers for selection
[**SelectFacebookPage**](ConnectAPI.md#SelectFacebookPage) | **Post** /v1/connect/facebook/select-page | Select Facebook page
[**SelectGoogleBusinessLocation**](ConnectAPI.md#SelectGoogleBusinessLocation) | **Post** /v1/connect/googlebusiness/select-location | Select GBP location
[**SelectLinkedInOrganization**](ConnectAPI.md#SelectLinkedInOrganization) | **Post** /v1/connect/linkedin/select-organization | Select LinkedIn org
[**SelectPinterestBoard**](ConnectAPI.md#SelectPinterestBoard) | **Post** /v1/connect/pinterest/select-board | Select Pinterest board
[**SelectSnapchatProfile**](ConnectAPI.md#SelectSnapchatProfile) | **Post** /v1/connect/snapchat/select-profile | Select Snapchat profile
[**UpdateFacebookPage**](ConnectAPI.md#UpdateFacebookPage) | **Put** /v1/accounts/{accountId}/facebook-page | Update Facebook page
[**UpdateGmbLocation**](ConnectAPI.md#UpdateGmbLocation) | **Put** /v1/accounts/{accountId}/gmb-locations | Update GBP location
[**UpdateLinkedInOrganization**](ConnectAPI.md#UpdateLinkedInOrganization) | **Put** /v1/accounts/{accountId}/linkedin-organization | Switch LinkedIn account type
[**UpdatePinterestBoards**](ConnectAPI.md#UpdatePinterestBoards) | **Put** /v1/accounts/{accountId}/pinterest-boards | Set default Pinterest board
[**UpdateRedditSubreddits**](ConnectAPI.md#UpdateRedditSubreddits) | **Put** /v1/accounts/{accountId}/reddit-subreddits | Set default subreddit
[**UpdateYoutubeDefaultPlaylist**](ConnectAPI.md#UpdateYoutubeDefaultPlaylist) | **Put** /v1/accounts/{accountId}/youtube-playlists | Set default YouTube playlist



## CompleteTelegramConnect

> CompleteTelegramConnect200Response CompleteTelegramConnect(ctx).Code(code).Execute()

Check Telegram status



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
	code := "ZRN-ABC123" // string | The access code to check status for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.CompleteTelegramConnect(context.Background()).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.CompleteTelegramConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteTelegramConnect`: CompleteTelegramConnect200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.CompleteTelegramConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteTelegramConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | The access code to check status for | 

### Return type

[**CompleteTelegramConnect200Response**](CompleteTelegramConnect200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteWhatsAppPhoneSelection

> CompleteWhatsAppPhoneSelection200Response CompleteWhatsAppPhoneSelection(ctx).CompleteWhatsAppPhoneSelectionRequest(completeWhatsAppPhoneSelectionRequest).XConnectToken(xConnectToken).Execute()

Complete WhatsApp phone number selection



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
	completeWhatsAppPhoneSelectionRequest := *openapiclient.NewCompleteWhatsAppPhoneSelectionRequest("ProfileId_example", "PhoneNumberId_example", "WabaId_example", "TempToken_example") // CompleteWhatsAppPhoneSelectionRequest | 
	xConnectToken := "xConnectToken_example" // string | Alternative auth for API users' end customers (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.CompleteWhatsAppPhoneSelection(context.Background()).CompleteWhatsAppPhoneSelectionRequest(completeWhatsAppPhoneSelectionRequest).XConnectToken(xConnectToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.CompleteWhatsAppPhoneSelection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteWhatsAppPhoneSelection`: CompleteWhatsAppPhoneSelection200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.CompleteWhatsAppPhoneSelection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteWhatsAppPhoneSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completeWhatsAppPhoneSelectionRequest** | [**CompleteWhatsAppPhoneSelectionRequest**](CompleteWhatsAppPhoneSelectionRequest.md) |  | 
 **xConnectToken** | **string** | Alternative auth for API users&#39; end customers | 

### Return type

[**CompleteWhatsAppPhoneSelection200Response**](CompleteWhatsAppPhoneSelection200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigureTikTokAdsBrandIdentity

> ConfigureTikTokAdsBrandIdentity200Response ConfigureTikTokAdsBrandIdentity(ctx).ConfigureTikTokAdsBrandIdentityRequest(configureTikTokAdsBrandIdentityRequest).Execute()

Configure TikTok Ads Brand Identity



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
	configureTikTokAdsBrandIdentityRequest := *openapiclient.NewConfigureTikTokAdsBrandIdentityRequest("AccountId_example", "DisplayName_example", "ImageUrl_example") // ConfigureTikTokAdsBrandIdentityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.ConfigureTikTokAdsBrandIdentity(context.Background()).ConfigureTikTokAdsBrandIdentityRequest(configureTikTokAdsBrandIdentityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.ConfigureTikTokAdsBrandIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigureTikTokAdsBrandIdentity`: ConfigureTikTokAdsBrandIdentity200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.ConfigureTikTokAdsBrandIdentity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigureTikTokAdsBrandIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configureTikTokAdsBrandIdentityRequest** | [**ConfigureTikTokAdsBrandIdentityRequest**](ConfigureTikTokAdsBrandIdentityRequest.md) |  | 

### Return type

[**ConfigureTikTokAdsBrandIdentity200Response**](ConfigureTikTokAdsBrandIdentity200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectAds

> ConnectAds200Response ConnectAds(ctx, platform).ProfileId(profileId).AccountId(accountId).RedirectUrl(redirectUrl).Headless(headless).AdAccountId(adAccountId).AdAccountIds(adAccountIds).Execute()

Connect ads for a platform



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
	platform := "platform_example" // string | Platform to connect ads for. Only platforms with ads support are accepted.
	profileId := "profileId_example" // string | Your Zernio profile ID
	accountId := "accountId_example" // string | Existing SocialAccount ID. Required for `twitter` (X Ads). Optional for `tiktok` — omit to enter ads-only mode (no TikTok posting account linked; ad creation uses a Brand Identity instead of a TT_USER). Ignored for same-token (`facebook`, `instagram`, `linkedin`, `pinterest`) and standalone (`googleads`) platforms.  (optional)
	redirectUrl := "redirectUrl_example" // string | Custom redirect URL after OAuth completes (same-token platforms only) (optional)
	headless := true // bool | Enable headless mode (same-token platforms only) (optional) (default to false)
	adAccountId := "act_1330190928038136" // string | Scope ad sync to a single platform ad account. Without this param, sync covers every ad account the connected token can see. Supported on `facebook`/`instagram` (Meta, `act_<digits>`), `linkedin` (bare numeric sponsored-account id), `googleads` (bare customer id digits) and `twitter` (X Ads, base36 account id). `tiktok` scopes advertisers at OAuth and `pinterest` has no ads discovery, so both ignore it. Meta ids are additionally validated against the connected token; unreachable IDs return 400. Setting a scope also removes already synced ads from de-scoped ad accounts. For multiple accounts use `adAccountIds` instead.  (optional)
	adAccountIds := []string{"Inner_example"} // []string | Scope ad sync to multiple platform ad accounts (same platform support and id shapes as `adAccountId`). Repeat the param (`?adAccountIds=act_1&adAccountIds=act_2`) or comma-separate (`?adAccountIds=act_1,act_2`). Persisted server-side; latest call wins, and de-scoped ad accounts have their synced ads removed. Omitting both `adAccountId` and `adAccountIds` keeps any previously persisted scope unchanged.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.ConnectAds(context.Background(), platform).ProfileId(profileId).AccountId(accountId).RedirectUrl(redirectUrl).Headless(headless).AdAccountId(adAccountId).AdAccountIds(adAccountIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.ConnectAds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectAds`: ConnectAds200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.ConnectAds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** | Platform to connect ads for. Only platforms with ads support are accepted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectAdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileId** | **string** | Your Zernio profile ID | 
 **accountId** | **string** | Existing SocialAccount ID. Required for &#x60;twitter&#x60; (X Ads). Optional for &#x60;tiktok&#x60; — omit to enter ads-only mode (no TikTok posting account linked; ad creation uses a Brand Identity instead of a TT_USER). Ignored for same-token (&#x60;facebook&#x60;, &#x60;instagram&#x60;, &#x60;linkedin&#x60;, &#x60;pinterest&#x60;) and standalone (&#x60;googleads&#x60;) platforms.  | 
 **redirectUrl** | **string** | Custom redirect URL after OAuth completes (same-token platforms only) | 
 **headless** | **bool** | Enable headless mode (same-token platforms only) | [default to false]
 **adAccountId** | **string** | Scope ad sync to a single platform ad account. Without this param, sync covers every ad account the connected token can see. Supported on &#x60;facebook&#x60;/&#x60;instagram&#x60; (Meta, &#x60;act_&lt;digits&gt;&#x60;), &#x60;linkedin&#x60; (bare numeric sponsored-account id), &#x60;googleads&#x60; (bare customer id digits) and &#x60;twitter&#x60; (X Ads, base36 account id). &#x60;tiktok&#x60; scopes advertisers at OAuth and &#x60;pinterest&#x60; has no ads discovery, so both ignore it. Meta ids are additionally validated against the connected token; unreachable IDs return 400. Setting a scope also removes already synced ads from de-scoped ad accounts. For multiple accounts use &#x60;adAccountIds&#x60; instead.  | 
 **adAccountIds** | **[]string** | Scope ad sync to multiple platform ad accounts (same platform support and id shapes as &#x60;adAccountId&#x60;). Repeat the param (&#x60;?adAccountIds&#x3D;act_1&amp;adAccountIds&#x3D;act_2&#x60;) or comma-separate (&#x60;?adAccountIds&#x3D;act_1,act_2&#x60;). Persisted server-side; latest call wins, and de-scoped ad accounts have their synced ads removed. Omitting both &#x60;adAccountId&#x60; and &#x60;adAccountIds&#x60; keeps any previously persisted scope unchanged.  | 

### Return type

[**ConnectAds200Response**](ConnectAds200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectBlueskyCredentials

> ConnectBlueskyCredentials200Response ConnectBlueskyCredentials(ctx).ConnectBlueskyCredentialsRequest(connectBlueskyCredentialsRequest).Execute()

Connect Bluesky account



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
	connectBlueskyCredentialsRequest := *openapiclient.NewConnectBlueskyCredentialsRequest("Identifier_example", "AppPassword_example", "6507a1b2c3d4e5f6a7b8c9d0-6507a1b2c3d4e5f6a7b8c9d1") // ConnectBlueskyCredentialsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.ConnectBlueskyCredentials(context.Background()).ConnectBlueskyCredentialsRequest(connectBlueskyCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.ConnectBlueskyCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectBlueskyCredentials`: ConnectBlueskyCredentials200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.ConnectBlueskyCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectBlueskyCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectBlueskyCredentialsRequest** | [**ConnectBlueskyCredentialsRequest**](ConnectBlueskyCredentialsRequest.md) |  | 

### Return type

[**ConnectBlueskyCredentials200Response**](ConnectBlueskyCredentials200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectWhatsAppCredentials

> ConnectWhatsAppCredentials200Response ConnectWhatsAppCredentials(ctx).ConnectWhatsAppCredentialsRequest(connectWhatsAppCredentialsRequest).Execute()

Connect WhatsApp via credentials



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
	connectWhatsAppCredentialsRequest := *openapiclient.NewConnectWhatsAppCredentialsRequest("ProfileId_example", "AccessToken_example", "WabaId_example", "PhoneNumberId_example") // ConnectWhatsAppCredentialsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.ConnectWhatsAppCredentials(context.Background()).ConnectWhatsAppCredentialsRequest(connectWhatsAppCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.ConnectWhatsAppCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectWhatsAppCredentials`: ConnectWhatsAppCredentials200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.ConnectWhatsAppCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectWhatsAppCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectWhatsAppCredentialsRequest** | [**ConnectWhatsAppCredentialsRequest**](ConnectWhatsAppCredentialsRequest.md) |  | 

### Return type

[**ConnectWhatsAppCredentials200Response**](ConnectWhatsAppCredentials200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectUrl

> GetConnectUrl200Response GetConnectUrl(ctx, platform).ProfileId(profileId).RedirectUrl(redirectUrl).Headless(headless).Execute()

Get OAuth connect URL



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
	platform := "platform_example" // string | Social media platform to connect
	profileId := "profileId_example" // string | Your Zernio profile ID (get from /v1/profiles)
	redirectUrl := "redirectUrl_example" // string | Your custom redirect URL after connection completes. Standard mode appends ?connected={platform}&profileId=X&accountId=Y&username=Z. Headless mode appends OAuth data params for platforms requiring selection (e.g. LinkedIn orgs, Facebook pages). If no selection is needed, the account is created directly and the redirect includes accountId. (optional)
	headless := true // bool | When true, the user is redirected to your redirect_url with raw OAuth data (code, state) instead of Zernio's default account selection UI. Use this to build a custom connect experience. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.GetConnectUrl(context.Background(), platform).ProfileId(profileId).RedirectUrl(redirectUrl).Headless(headless).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.GetConnectUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectUrl`: GetConnectUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.GetConnectUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** | Social media platform to connect | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileId** | **string** | Your Zernio profile ID (get from /v1/profiles) | 
 **redirectUrl** | **string** | Your custom redirect URL after connection completes. Standard mode appends ?connected&#x3D;{platform}&amp;profileId&#x3D;X&amp;accountId&#x3D;Y&amp;username&#x3D;Z. Headless mode appends OAuth data params for platforms requiring selection (e.g. LinkedIn orgs, Facebook pages). If no selection is needed, the account is created directly and the redirect includes accountId. | 
 **headless** | **bool** | When true, the user is redirected to your redirect_url with raw OAuth data (code, state) instead of Zernio&#39;s default account selection UI. Use this to build a custom connect experience. | [default to false]

### Return type

[**GetConnectUrl200Response**](GetConnectUrl200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFacebookPages

> GetFacebookPages200Response GetFacebookPages(ctx, accountId).Refresh(refresh).Execute()

List Facebook pages



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
	refresh := true // bool | When true, bypasses the page cache and fetches fresh pages from Meta. Rate-limited server-side to 1 refresh per 60s. Pages no longer accessible to the connected account will be removed from the list on refresh.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.GetFacebookPages(context.Background(), accountId).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.GetFacebookPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFacebookPages`: GetFacebookPages200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.GetFacebookPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFacebookPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refresh** | **bool** | When true, bypasses the page cache and fetches fresh pages from Meta. Rate-limited server-side to 1 refresh per 60s. Pages no longer accessible to the connected account will be removed from the list on refresh.  | 

### Return type

[**GetFacebookPages200Response**](GetFacebookPages200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGmbLocations

> GetGmbLocations200Response GetGmbLocations(ctx, accountId).Search(search).Filter(filter).Execute()

List GBP locations



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
	search := "search_example" // string | Free-text search on the business name, applied server-side by Google. Use for accounts with many locations. (optional)
	filter := "filter_example" // string | Raw Google Business Information API filter expression (advanced; takes precedence over search), e.g. storeCode=\"LH279411\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.GetGmbLocations(context.Background(), accountId).Search(search).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.GetGmbLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGmbLocations`: GetGmbLocations200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.GetGmbLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGmbLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | Free-text search on the business name, applied server-side by Google. Use for accounts with many locations. | 
 **filter** | **string** | Raw Google Business Information API filter expression (advanced; takes precedence over search), e.g. storeCode&#x3D;\&quot;LH279411\&quot;. | 

### Return type

[**GetGmbLocations200Response**](GetGmbLocations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedInOrganizations

> GetLinkedInOrganizations200Response GetLinkedInOrganizations(ctx, accountId).Execute()

List LinkedIn orgs



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
	resp, r, err := apiClient.ConnectAPI.GetLinkedInOrganizations(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.GetLinkedInOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedInOrganizations`: GetLinkedInOrganizations200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.GetLinkedInOrganizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedInOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLinkedInOrganizations200Response**](GetLinkedInOrganizations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPendingOAuthData

> GetPendingOAuthData200Response GetPendingOAuthData(ctx).Token(token).Execute()

Get pending OAuth data



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
	token := "token_example" // string | The pending data token from the OAuth redirect URL (pendingDataToken parameter)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.GetPendingOAuthData(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.GetPendingOAuthData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPendingOAuthData`: GetPendingOAuthData200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.GetPendingOAuthData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingOAuthDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The pending data token from the OAuth redirect URL (pendingDataToken parameter) | 

### Return type

[**GetPendingOAuthData200Response**](GetPendingOAuthData200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPinterestBoards

> ListPinterestBoardsForSelection200Response GetPinterestBoards(ctx, accountId).Execute()

List Pinterest boards



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
	resp, r, err := apiClient.ConnectAPI.GetPinterestBoards(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.GetPinterestBoards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPinterestBoards`: ListPinterestBoardsForSelection200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.GetPinterestBoards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPinterestBoardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPinterestBoardsForSelection200Response**](ListPinterestBoardsForSelection200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedditFlairs

> GetRedditFlairs200Response GetRedditFlairs(ctx, accountId).Subreddit(subreddit).Execute()

List subreddit flairs



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
	subreddit := "subreddit_example" // string | Subreddit name (without \"r/\" prefix) to fetch flairs for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.GetRedditFlairs(context.Background(), accountId).Subreddit(subreddit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.GetRedditFlairs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedditFlairs`: GetRedditFlairs200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.GetRedditFlairs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedditFlairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subreddit** | **string** | Subreddit name (without \&quot;r/\&quot; prefix) to fetch flairs for | 

### Return type

[**GetRedditFlairs200Response**](GetRedditFlairs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedditSubreddits

> GetRedditSubreddits200Response GetRedditSubreddits(ctx, accountId).Execute()

List Reddit subreddits



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
	resp, r, err := apiClient.ConnectAPI.GetRedditSubreddits(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.GetRedditSubreddits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedditSubreddits`: GetRedditSubreddits200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.GetRedditSubreddits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedditSubredditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRedditSubreddits200Response**](GetRedditSubreddits200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelegramConnectStatus

> GetTelegramConnectStatus200Response GetTelegramConnectStatus(ctx).ProfileId(profileId).Execute()

Generate Telegram code



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
	profileId := "profileId_example" // string | The profile ID to connect the Telegram account to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.GetTelegramConnectStatus(context.Background()).ProfileId(profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.GetTelegramConnectStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelegramConnectStatus`: GetTelegramConnectStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.GetTelegramConnectStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelegramConnectStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | The profile ID to connect the Telegram account to | 

### Return type

[**GetTelegramConnectStatus200Response**](GetTelegramConnectStatus200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetYoutubePlaylists

> GetYoutubePlaylists200Response GetYoutubePlaylists(ctx, accountId).Execute()

List YouTube playlists



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
	resp, r, err := apiClient.ConnectAPI.GetYoutubePlaylists(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.GetYoutubePlaylists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetYoutubePlaylists`: GetYoutubePlaylists200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.GetYoutubePlaylists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetYoutubePlaylistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetYoutubePlaylists200Response**](GetYoutubePlaylists200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleOAuthCallback

> HandleOAuthCallback(ctx, platform).HandleOAuthCallbackRequest(handleOAuthCallbackRequest).Execute()

Complete OAuth callback



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
	platform := "platform_example" // string | 
	handleOAuthCallbackRequest := *openapiclient.NewHandleOAuthCallbackRequest("Code_example", "State_example", "ProfileId_example") // HandleOAuthCallbackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectAPI.HandleOAuthCallback(context.Background(), platform).HandleOAuthCallbackRequest(handleOAuthCallbackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.HandleOAuthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleOAuthCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **handleOAuthCallbackRequest** | [**HandleOAuthCallbackRequest**](HandleOAuthCallbackRequest.md) |  | 

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


## InitiateTelegramConnect

> InitiateTelegramConnect200Response InitiateTelegramConnect(ctx).InitiateTelegramConnectRequest(initiateTelegramConnectRequest).Execute()

Connect Telegram directly



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
	initiateTelegramConnectRequest := *openapiclient.NewInitiateTelegramConnectRequest("ChatId_example", "ProfileId_example") // InitiateTelegramConnectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.InitiateTelegramConnect(context.Background()).InitiateTelegramConnectRequest(initiateTelegramConnectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.InitiateTelegramConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateTelegramConnect`: InitiateTelegramConnect200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.InitiateTelegramConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateTelegramConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **initiateTelegramConnectRequest** | [**InitiateTelegramConnectRequest**](InitiateTelegramConnectRequest.md) |  | 

### Return type

[**InitiateTelegramConnect200Response**](InitiateTelegramConnect200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFacebookPages

> ListFacebookPages200Response ListFacebookPages(ctx).ProfileId(profileId).TempToken(tempToken).Execute()

List Facebook pages



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
	profileId := "profileId_example" // string | Profile ID from your connection flow
	tempToken := "tempToken_example" // string | Temporary Facebook access token from the OAuth callback redirect

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.ListFacebookPages(context.Background()).ProfileId(profileId).TempToken(tempToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.ListFacebookPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFacebookPages`: ListFacebookPages200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.ListFacebookPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFacebookPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Profile ID from your connection flow | 
 **tempToken** | **string** | Temporary Facebook access token from the OAuth callback redirect | 

### Return type

[**ListFacebookPages200Response**](ListFacebookPages200Response.md)

### Authorization

[connectToken](../README.md#connectToken), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGoogleBusinessLocations

> ListGoogleBusinessLocations200Response ListGoogleBusinessLocations(ctx).ProfileId(profileId).PendingDataToken(pendingDataToken).TempToken(tempToken).Search(search).Filter(filter).Execute()

List GBP locations



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
	profileId := "profileId_example" // string | Profile ID from your connection flow. Required for auth validation when provided. (optional)
	pendingDataToken := "pendingDataToken_example" // string | Token from the OAuth callback redirect. Preferred over tempToken because it preserves server-side token storage. One of pendingDataToken or tempToken is required. (optional)
	tempToken := "tempToken_example" // string | Legacy. Direct Google access token. Use pendingDataToken instead when available. (optional)
	search := "search_example" // string | Free-text search on the business name, applied server-side by Google. Use this for accounts that own many locations (the response is bounded, see hasMore) so the user can find a specific location without loading the full list.  (optional)
	filter := "filter_example" // string | Raw Google Business Information API filter expression (advanced; takes precedence over search). Supports fields such as title, storeCode, storefront_address.postal_code, labels and categories, e.g. storeCode=\"LH279411\". See Google's \"Work with location data\" guide.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.ListGoogleBusinessLocations(context.Background()).ProfileId(profileId).PendingDataToken(pendingDataToken).TempToken(tempToken).Search(search).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.ListGoogleBusinessLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGoogleBusinessLocations`: ListGoogleBusinessLocations200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.ListGoogleBusinessLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGoogleBusinessLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Profile ID from your connection flow. Required for auth validation when provided. | 
 **pendingDataToken** | **string** | Token from the OAuth callback redirect. Preferred over tempToken because it preserves server-side token storage. One of pendingDataToken or tempToken is required. | 
 **tempToken** | **string** | Legacy. Direct Google access token. Use pendingDataToken instead when available. | 
 **search** | **string** | Free-text search on the business name, applied server-side by Google. Use this for accounts that own many locations (the response is bounded, see hasMore) so the user can find a specific location without loading the full list.  | 
 **filter** | **string** | Raw Google Business Information API filter expression (advanced; takes precedence over search). Supports fields such as title, storeCode, storefront_address.postal_code, labels and categories, e.g. storeCode&#x3D;\&quot;LH279411\&quot;. See Google&#39;s \&quot;Work with location data\&quot; guide.  | 

### Return type

[**ListGoogleBusinessLocations200Response**](ListGoogleBusinessLocations200Response.md)

### Authorization

[connectToken](../README.md#connectToken), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLinkedInOrganizations

> ListLinkedInOrganizations200Response ListLinkedInOrganizations(ctx).TempToken(tempToken).OrgIds(orgIds).Execute()

List LinkedIn orgs



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
	tempToken := "tempToken_example" // string | The temporary LinkedIn access token from the OAuth redirect
	orgIds := "12345678,87654321,11111111" // string | Comma-separated list of organization IDs to fetch details for (max 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.ListLinkedInOrganizations(context.Background()).TempToken(tempToken).OrgIds(orgIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.ListLinkedInOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLinkedInOrganizations`: ListLinkedInOrganizations200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.ListLinkedInOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLinkedInOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tempToken** | **string** | The temporary LinkedIn access token from the OAuth redirect | 
 **orgIds** | **string** | Comma-separated list of organization IDs to fetch details for (max 100) | 

### Return type

[**ListLinkedInOrganizations200Response**](ListLinkedInOrganizations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPinterestBoardsForSelection

> ListPinterestBoardsForSelection200Response ListPinterestBoardsForSelection(ctx).XConnectToken(xConnectToken).ProfileId(profileId).TempToken(tempToken).Execute()

List Pinterest boards



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
	xConnectToken := "xConnectToken_example" // string | Short-lived connect token from the OAuth redirect
	profileId := "profileId_example" // string | Your Zernio profile ID
	tempToken := "tempToken_example" // string | Temporary Pinterest access token from the OAuth callback redirect

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.ListPinterestBoardsForSelection(context.Background()).XConnectToken(xConnectToken).ProfileId(profileId).TempToken(tempToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.ListPinterestBoardsForSelection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPinterestBoardsForSelection`: ListPinterestBoardsForSelection200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.ListPinterestBoardsForSelection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPinterestBoardsForSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xConnectToken** | **string** | Short-lived connect token from the OAuth redirect | 
 **profileId** | **string** | Your Zernio profile ID | 
 **tempToken** | **string** | Temporary Pinterest access token from the OAuth callback redirect | 

### Return type

[**ListPinterestBoardsForSelection200Response**](ListPinterestBoardsForSelection200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSnapchatProfiles

> ListSnapchatProfiles200Response ListSnapchatProfiles(ctx).XConnectToken(xConnectToken).ProfileId(profileId).TempToken(tempToken).Execute()

List Snapchat profiles



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
	xConnectToken := "xConnectToken_example" // string | Short-lived connect token from the OAuth redirect
	profileId := "profileId_example" // string | Your Zernio profile ID
	tempToken := "tempToken_example" // string | Temporary Snapchat access token from the OAuth callback redirect

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.ListSnapchatProfiles(context.Background()).XConnectToken(xConnectToken).ProfileId(profileId).TempToken(tempToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.ListSnapchatProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSnapchatProfiles`: ListSnapchatProfiles200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.ListSnapchatProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSnapchatProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xConnectToken** | **string** | Short-lived connect token from the OAuth redirect | 
 **profileId** | **string** | Your Zernio profile ID | 
 **tempToken** | **string** | Temporary Snapchat access token from the OAuth callback redirect | 

### Return type

[**ListSnapchatProfiles200Response**](ListSnapchatProfiles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWhatsAppPhoneNumbers

> ListWhatsAppPhoneNumbers200Response ListWhatsAppPhoneNumbers(ctx).ProfileId(profileId).TempToken(tempToken).XConnectToken(xConnectToken).Execute()

List WhatsApp phone numbers for selection



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
	profileId := "profileId_example" // string | The Zernio profile ID from the headless redirect
	tempToken := "tempToken_example" // string | The temporary access token from the headless redirect
	xConnectToken := "xConnectToken_example" // string | Alternative auth for API users' end customers (used when the bearer token is scoped to a different user) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.ListWhatsAppPhoneNumbers(context.Background()).ProfileId(profileId).TempToken(tempToken).XConnectToken(xConnectToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.ListWhatsAppPhoneNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWhatsAppPhoneNumbers`: ListWhatsAppPhoneNumbers200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.ListWhatsAppPhoneNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWhatsAppPhoneNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | The Zernio profile ID from the headless redirect | 
 **tempToken** | **string** | The temporary access token from the headless redirect | 
 **xConnectToken** | **string** | Alternative auth for API users&#39; end customers (used when the bearer token is scoped to a different user) | 

### Return type

[**ListWhatsAppPhoneNumbers200Response**](ListWhatsAppPhoneNumbers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectFacebookPage

> SelectFacebookPage200Response SelectFacebookPage(ctx).SelectFacebookPageRequest(selectFacebookPageRequest).Execute()

Select Facebook page



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
	selectFacebookPageRequest := *openapiclient.NewSelectFacebookPageRequest("ProfileId_example", "PageId_example", "TempToken_example", *openapiclient.NewSelectFacebookPageRequestUserProfile()) // SelectFacebookPageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.SelectFacebookPage(context.Background()).SelectFacebookPageRequest(selectFacebookPageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.SelectFacebookPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectFacebookPage`: SelectFacebookPage200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.SelectFacebookPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelectFacebookPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selectFacebookPageRequest** | [**SelectFacebookPageRequest**](SelectFacebookPageRequest.md) |  | 

### Return type

[**SelectFacebookPage200Response**](SelectFacebookPage200Response.md)

### Authorization

[connectToken](../README.md#connectToken), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectGoogleBusinessLocation

> SelectGoogleBusinessLocation200Response SelectGoogleBusinessLocation(ctx).SelectGoogleBusinessLocationRequest(selectGoogleBusinessLocationRequest).Execute()

Select GBP location



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
	selectGoogleBusinessLocationRequest := *openapiclient.NewSelectGoogleBusinessLocationRequest("ProfileId_example", "LocationId_example", "PendingDataToken_example") // SelectGoogleBusinessLocationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.SelectGoogleBusinessLocation(context.Background()).SelectGoogleBusinessLocationRequest(selectGoogleBusinessLocationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.SelectGoogleBusinessLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectGoogleBusinessLocation`: SelectGoogleBusinessLocation200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.SelectGoogleBusinessLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelectGoogleBusinessLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selectGoogleBusinessLocationRequest** | [**SelectGoogleBusinessLocationRequest**](SelectGoogleBusinessLocationRequest.md) |  | 

### Return type

[**SelectGoogleBusinessLocation200Response**](SelectGoogleBusinessLocation200Response.md)

### Authorization

[connectToken](../README.md#connectToken), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectLinkedInOrganization

> SelectLinkedInOrganization200Response SelectLinkedInOrganization(ctx).SelectLinkedInOrganizationRequest(selectLinkedInOrganizationRequest).Execute()

Select LinkedIn org



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
	selectLinkedInOrganizationRequest := *openapiclient.NewSelectLinkedInOrganizationRequest("ProfileId_example", "TempToken_example", map[string]interface{}(123), "AccountType_example") // SelectLinkedInOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.SelectLinkedInOrganization(context.Background()).SelectLinkedInOrganizationRequest(selectLinkedInOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.SelectLinkedInOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectLinkedInOrganization`: SelectLinkedInOrganization200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.SelectLinkedInOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelectLinkedInOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selectLinkedInOrganizationRequest** | [**SelectLinkedInOrganizationRequest**](SelectLinkedInOrganizationRequest.md) |  | 

### Return type

[**SelectLinkedInOrganization200Response**](SelectLinkedInOrganization200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectPinterestBoard

> SelectPinterestBoard200Response SelectPinterestBoard(ctx).SelectPinterestBoardRequest(selectPinterestBoardRequest).Execute()

Select Pinterest board



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
	selectPinterestBoardRequest := *openapiclient.NewSelectPinterestBoardRequest("ProfileId_example", "BoardId_example", "TempToken_example") // SelectPinterestBoardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.SelectPinterestBoard(context.Background()).SelectPinterestBoardRequest(selectPinterestBoardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.SelectPinterestBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectPinterestBoard`: SelectPinterestBoard200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.SelectPinterestBoard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelectPinterestBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selectPinterestBoardRequest** | [**SelectPinterestBoardRequest**](SelectPinterestBoardRequest.md) |  | 

### Return type

[**SelectPinterestBoard200Response**](SelectPinterestBoard200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectSnapchatProfile

> SelectSnapchatProfile200Response SelectSnapchatProfile(ctx).SelectSnapchatProfileRequest(selectSnapchatProfileRequest).XConnectToken(xConnectToken).Execute()

Select Snapchat profile



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
	selectSnapchatProfileRequest := *openapiclient.NewSelectSnapchatProfileRequest("ProfileId_example", *openapiclient.NewSelectSnapchatProfileRequestSelectedPublicProfile("Id_example", "DisplayName_example"), "TempToken_example", map[string]interface{}(123)) // SelectSnapchatProfileRequest | 
	xConnectToken := "xConnectToken_example" // string | Short-lived connect token from the OAuth redirect (for API users) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.SelectSnapchatProfile(context.Background()).SelectSnapchatProfileRequest(selectSnapchatProfileRequest).XConnectToken(xConnectToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.SelectSnapchatProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectSnapchatProfile`: SelectSnapchatProfile200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.SelectSnapchatProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelectSnapchatProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selectSnapchatProfileRequest** | [**SelectSnapchatProfileRequest**](SelectSnapchatProfileRequest.md) |  | 
 **xConnectToken** | **string** | Short-lived connect token from the OAuth redirect (for API users) | 

### Return type

[**SelectSnapchatProfile200Response**](SelectSnapchatProfile200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFacebookPage

> UpdateFacebookPage200Response UpdateFacebookPage(ctx, accountId).UpdateFacebookPageRequest(updateFacebookPageRequest).Execute()

Update Facebook page



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
	updateFacebookPageRequest := *openapiclient.NewUpdateFacebookPageRequest("SelectedPageId_example") // UpdateFacebookPageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.UpdateFacebookPage(context.Background(), accountId).UpdateFacebookPageRequest(updateFacebookPageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.UpdateFacebookPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFacebookPage`: UpdateFacebookPage200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.UpdateFacebookPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFacebookPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFacebookPageRequest** | [**UpdateFacebookPageRequest**](UpdateFacebookPageRequest.md) |  | 

### Return type

[**UpdateFacebookPage200Response**](UpdateFacebookPage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGmbLocation

> UpdateGmbLocation200Response UpdateGmbLocation(ctx, accountId).UpdateGmbLocationRequest(updateGmbLocationRequest).Execute()

Update GBP location



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
	updateGmbLocationRequest := *openapiclient.NewUpdateGmbLocationRequest("SelectedLocationId_example") // UpdateGmbLocationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.UpdateGmbLocation(context.Background(), accountId).UpdateGmbLocationRequest(updateGmbLocationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.UpdateGmbLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGmbLocation`: UpdateGmbLocation200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.UpdateGmbLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGmbLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGmbLocationRequest** | [**UpdateGmbLocationRequest**](UpdateGmbLocationRequest.md) |  | 

### Return type

[**UpdateGmbLocation200Response**](UpdateGmbLocation200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLinkedInOrganization

> ConnectBlueskyCredentials200Response UpdateLinkedInOrganization(ctx, accountId).UpdateLinkedInOrganizationRequest(updateLinkedInOrganizationRequest).Execute()

Switch LinkedIn account type



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
	updateLinkedInOrganizationRequest := *openapiclient.NewUpdateLinkedInOrganizationRequest("AccountType_example") // UpdateLinkedInOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.UpdateLinkedInOrganization(context.Background(), accountId).UpdateLinkedInOrganizationRequest(updateLinkedInOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.UpdateLinkedInOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLinkedInOrganization`: ConnectBlueskyCredentials200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.UpdateLinkedInOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLinkedInOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLinkedInOrganizationRequest** | [**UpdateLinkedInOrganizationRequest**](UpdateLinkedInOrganizationRequest.md) |  | 

### Return type

[**ConnectBlueskyCredentials200Response**](ConnectBlueskyCredentials200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePinterestBoards

> ConnectBlueskyCredentials200Response UpdatePinterestBoards(ctx, accountId).UpdatePinterestBoardsRequest(updatePinterestBoardsRequest).Execute()

Set default Pinterest board



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
	updatePinterestBoardsRequest := *openapiclient.NewUpdatePinterestBoardsRequest("DefaultBoardId_example") // UpdatePinterestBoardsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.UpdatePinterestBoards(context.Background(), accountId).UpdatePinterestBoardsRequest(updatePinterestBoardsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.UpdatePinterestBoards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePinterestBoards`: ConnectBlueskyCredentials200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.UpdatePinterestBoards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePinterestBoardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePinterestBoardsRequest** | [**UpdatePinterestBoardsRequest**](UpdatePinterestBoardsRequest.md) |  | 

### Return type

[**ConnectBlueskyCredentials200Response**](ConnectBlueskyCredentials200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRedditSubreddits

> UpdateYoutubeDefaultPlaylist200Response UpdateRedditSubreddits(ctx, accountId).UpdateRedditSubredditsRequest(updateRedditSubredditsRequest).Execute()

Set default subreddit



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
	updateRedditSubredditsRequest := *openapiclient.NewUpdateRedditSubredditsRequest("DefaultSubreddit_example") // UpdateRedditSubredditsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.UpdateRedditSubreddits(context.Background(), accountId).UpdateRedditSubredditsRequest(updateRedditSubredditsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.UpdateRedditSubreddits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRedditSubreddits`: UpdateYoutubeDefaultPlaylist200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.UpdateRedditSubreddits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRedditSubredditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRedditSubredditsRequest** | [**UpdateRedditSubredditsRequest**](UpdateRedditSubredditsRequest.md) |  | 

### Return type

[**UpdateYoutubeDefaultPlaylist200Response**](UpdateYoutubeDefaultPlaylist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateYoutubeDefaultPlaylist

> UpdateYoutubeDefaultPlaylist200Response UpdateYoutubeDefaultPlaylist(ctx, accountId).UpdateYoutubeDefaultPlaylistRequest(updateYoutubeDefaultPlaylistRequest).Execute()

Set default YouTube playlist



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
	updateYoutubeDefaultPlaylistRequest := *openapiclient.NewUpdateYoutubeDefaultPlaylistRequest("DefaultPlaylistId_example") // UpdateYoutubeDefaultPlaylistRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.UpdateYoutubeDefaultPlaylist(context.Background(), accountId).UpdateYoutubeDefaultPlaylistRequest(updateYoutubeDefaultPlaylistRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.UpdateYoutubeDefaultPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateYoutubeDefaultPlaylist`: UpdateYoutubeDefaultPlaylist200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.UpdateYoutubeDefaultPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateYoutubeDefaultPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateYoutubeDefaultPlaylistRequest** | [**UpdateYoutubeDefaultPlaylistRequest**](UpdateYoutubeDefaultPlaylistRequest.md) |  | 

### Return type

[**UpdateYoutubeDefaultPlaylist200Response**](UpdateYoutubeDefaultPlaylist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

