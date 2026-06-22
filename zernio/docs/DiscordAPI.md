# \DiscordAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDiscordMemberRole**](DiscordAPI.md#AddDiscordMemberRole) | **Put** /v1/discord/guilds/{guildId}/members/{userId}/roles/{roleId} | Assign a role to a guild member
[**CreateDiscordScheduledEvent**](DiscordAPI.md#CreateDiscordScheduledEvent) | **Post** /v1/discord/guilds/{guildId}/events | Create a Discord scheduled event
[**DeleteDiscordScheduledEvent**](DiscordAPI.md#DeleteDiscordScheduledEvent) | **Delete** /v1/discord/guilds/{guildId}/events/{eventId} | Delete a Discord scheduled event
[**GetDiscordChannels**](DiscordAPI.md#GetDiscordChannels) | **Get** /v1/accounts/{accountId}/discord-channels | List Discord guild channels
[**GetDiscordScheduledEvent**](DiscordAPI.md#GetDiscordScheduledEvent) | **Get** /v1/discord/guilds/{guildId}/events/{eventId} | Get a Discord scheduled event
[**GetDiscordSettings**](DiscordAPI.md#GetDiscordSettings) | **Get** /v1/accounts/{accountId}/discord-settings | Get Discord account settings
[**ListDiscordGuildMembers**](DiscordAPI.md#ListDiscordGuildMembers) | **Get** /v1/discord/guilds/{guildId}/members | List Discord guild members
[**ListDiscordGuildRoles**](DiscordAPI.md#ListDiscordGuildRoles) | **Get** /v1/discord/guilds/{guildId}/roles | List Discord guild roles
[**ListDiscordPinnedMessages**](DiscordAPI.md#ListDiscordPinnedMessages) | **Get** /v1/discord/channels/{channelId}/pins | List pinned messages in a Discord channel
[**ListDiscordScheduledEvents**](DiscordAPI.md#ListDiscordScheduledEvents) | **Get** /v1/discord/guilds/{guildId}/events | List Discord scheduled events
[**PinDiscordMessage**](DiscordAPI.md#PinDiscordMessage) | **Put** /v1/discord/channels/{channelId}/pins/{messageId} | Pin a Discord message
[**RemoveDiscordMemberRole**](DiscordAPI.md#RemoveDiscordMemberRole) | **Delete** /v1/discord/guilds/{guildId}/members/{userId}/roles/{roleId} | Remove a role from a guild member
[**SendDiscordDirectMessage**](DiscordAPI.md#SendDiscordDirectMessage) | **Post** /v1/discord/dms | Send a Discord Direct Message
[**UnpinDiscordMessage**](DiscordAPI.md#UnpinDiscordMessage) | **Delete** /v1/discord/channels/{channelId}/pins/{messageId} | Unpin a Discord message
[**UpdateDiscordScheduledEvent**](DiscordAPI.md#UpdateDiscordScheduledEvent) | **Patch** /v1/discord/guilds/{guildId}/events/{eventId} | Update a Discord scheduled event
[**UpdateDiscordSettings**](DiscordAPI.md#UpdateDiscordSettings) | **Patch** /v1/accounts/{accountId}/discord-settings | Update Discord settings



## AddDiscordMemberRole

> AddDiscordMemberRole200Response AddDiscordMemberRole(ctx, guildId, userId, roleId).AccountId(accountId).Execute()

Assign a role to a guild member



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
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string | Discord user snowflake to assign the role to.
	roleId := "roleId_example" // string | 
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.AddDiscordMemberRole(context.Background(), guildId, userId, roleId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.AddDiscordMemberRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDiscordMemberRole`: AddDiscordMemberRole200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.AddDiscordMemberRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**userId** | **string** | Discord user snowflake to assign the role to. | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDiscordMemberRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accountId** | **string** |  | 

### Return type

[**AddDiscordMemberRole200Response**](AddDiscordMemberRole200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDiscordScheduledEvent

> CreateDiscordScheduledEvent200Response CreateDiscordScheduledEvent(ctx, guildId).CreateDiscordScheduledEventRequest(createDiscordScheduledEventRequest).Execute()

Create a Discord scheduled event



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
	guildId := "guildId_example" // string | 
	createDiscordScheduledEventRequest := *openapiclient.NewCreateDiscordScheduledEventRequest("AccountId_example", "Name_example", time.Now(), openapiclient.createDiscordScheduledEvent_request_entity{CreateDiscordScheduledEventRequestEntityOneOf: openapiclient.NewCreateDiscordScheduledEventRequestEntityOneOf("Type_example", "Location_example", time.Now())}) // CreateDiscordScheduledEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.CreateDiscordScheduledEvent(context.Background(), guildId).CreateDiscordScheduledEventRequest(createDiscordScheduledEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.CreateDiscordScheduledEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDiscordScheduledEvent`: CreateDiscordScheduledEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.CreateDiscordScheduledEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiscordScheduledEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDiscordScheduledEventRequest** | [**CreateDiscordScheduledEventRequest**](CreateDiscordScheduledEventRequest.md) |  | 

### Return type

[**CreateDiscordScheduledEvent200Response**](CreateDiscordScheduledEvent200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDiscordScheduledEvent

> DeleteDiscordScheduledEvent200Response DeleteDiscordScheduledEvent(ctx, guildId, eventId).AccountId(accountId).Execute()

Delete a Discord scheduled event



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
	guildId := "guildId_example" // string | 
	eventId := "eventId_example" // string | 
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.DeleteDiscordScheduledEvent(context.Background(), guildId, eventId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.DeleteDiscordScheduledEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDiscordScheduledEvent`: DeleteDiscordScheduledEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.DeleteDiscordScheduledEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDiscordScheduledEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **string** |  | 

### Return type

[**DeleteDiscordScheduledEvent200Response**](DeleteDiscordScheduledEvent200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscordChannels

> GetDiscordChannels200Response GetDiscordChannels(ctx, accountId).Execute()

List Discord guild channels



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
	resp, r, err := apiClient.DiscordAPI.GetDiscordChannels(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.GetDiscordChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiscordChannels`: GetDiscordChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.GetDiscordChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscordChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDiscordChannels200Response**](GetDiscordChannels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscordScheduledEvent

> CreateDiscordScheduledEvent200Response GetDiscordScheduledEvent(ctx, guildId, eventId).AccountId(accountId).Execute()

Get a Discord scheduled event

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
	guildId := "guildId_example" // string | 
	eventId := "eventId_example" // string | 
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.GetDiscordScheduledEvent(context.Background(), guildId, eventId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.GetDiscordScheduledEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiscordScheduledEvent`: CreateDiscordScheduledEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.GetDiscordScheduledEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscordScheduledEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **string** |  | 

### Return type

[**CreateDiscordScheduledEvent200Response**](CreateDiscordScheduledEvent200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscordSettings

> GetDiscordSettings200Response GetDiscordSettings(ctx, accountId).Execute()

Get Discord account settings



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
	resp, r, err := apiClient.DiscordAPI.GetDiscordSettings(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.GetDiscordSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiscordSettings`: GetDiscordSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.GetDiscordSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscordSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDiscordSettings200Response**](GetDiscordSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiscordGuildMembers

> ListDiscordGuildMembers200Response ListDiscordGuildMembers(ctx, guildId).AccountId(accountId).Limit(limit).After(after).Execute()

List Discord guild members



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
	guildId := "guildId_example" // string | 
	accountId := "accountId_example" // string | 
	limit := int32(56) // int32 | Page size (1-1000). (optional) (default to 100)
	after := "after_example" // string | Snowflake of the last member from the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.ListDiscordGuildMembers(context.Background(), guildId).AccountId(accountId).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.ListDiscordGuildMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDiscordGuildMembers`: ListDiscordGuildMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.ListDiscordGuildMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDiscordGuildMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** |  | 
 **limit** | **int32** | Page size (1-1000). | [default to 100]
 **after** | **string** | Snowflake of the last member from the previous page. | 

### Return type

[**ListDiscordGuildMembers200Response**](ListDiscordGuildMembers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiscordGuildRoles

> ListDiscordGuildRoles200Response ListDiscordGuildRoles(ctx, guildId).AccountId(accountId).Execute()

List Discord guild roles



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
	guildId := "guildId_example" // string | Discord guild snowflake ID
	accountId := "accountId_example" // string | SocialAccount _id of the Discord account bound to this guild

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.ListDiscordGuildRoles(context.Background(), guildId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.ListDiscordGuildRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDiscordGuildRoles`: ListDiscordGuildRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.ListDiscordGuildRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** | Discord guild snowflake ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDiscordGuildRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | SocialAccount _id of the Discord account bound to this guild | 

### Return type

[**ListDiscordGuildRoles200Response**](ListDiscordGuildRoles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiscordPinnedMessages

> ListDiscordPinnedMessages200Response ListDiscordPinnedMessages(ctx, channelId).AccountId(accountId).Execute()

List pinned messages in a Discord channel



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
	channelId := "channelId_example" // string | Discord channel snowflake.
	accountId := "accountId_example" // string | SocialAccount _id of any Discord account in the same guild.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.ListDiscordPinnedMessages(context.Background(), channelId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.ListDiscordPinnedMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDiscordPinnedMessages`: ListDiscordPinnedMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.ListDiscordPinnedMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | Discord channel snowflake. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDiscordPinnedMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | SocialAccount _id of any Discord account in the same guild. | 

### Return type

[**ListDiscordPinnedMessages200Response**](ListDiscordPinnedMessages200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiscordScheduledEvents

> ListDiscordScheduledEvents200Response ListDiscordScheduledEvents(ctx, guildId).AccountId(accountId).WithUserCount(withUserCount).Execute()

List Discord scheduled events



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
	guildId := "guildId_example" // string | 
	accountId := "accountId_example" // string | 
	withUserCount := true // bool | Include user_count on each event. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.ListDiscordScheduledEvents(context.Background(), guildId).AccountId(accountId).WithUserCount(withUserCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.ListDiscordScheduledEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDiscordScheduledEvents`: ListDiscordScheduledEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.ListDiscordScheduledEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDiscordScheduledEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** |  | 
 **withUserCount** | **bool** | Include user_count on each event. | 

### Return type

[**ListDiscordScheduledEvents200Response**](ListDiscordScheduledEvents200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PinDiscordMessage

> PinDiscordMessage200Response PinDiscordMessage(ctx, channelId, messageId).AccountId(accountId).Execute()

Pin a Discord message



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
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.PinDiscordMessage(context.Background(), channelId, messageId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.PinDiscordMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PinDiscordMessage`: PinDiscordMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.PinDiscordMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPinDiscordMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **string** |  | 

### Return type

[**PinDiscordMessage200Response**](PinDiscordMessage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDiscordMemberRole

> RemoveDiscordMemberRole200Response RemoveDiscordMemberRole(ctx, guildId, userId, roleId).AccountId(accountId).Execute()

Remove a role from a guild member



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
	guildId := "guildId_example" // string | 
	userId := "userId_example" // string | 
	roleId := "roleId_example" // string | 
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.RemoveDiscordMemberRole(context.Background(), guildId, userId, roleId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.RemoveDiscordMemberRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDiscordMemberRole`: RemoveDiscordMemberRole200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.RemoveDiscordMemberRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**userId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDiscordMemberRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accountId** | **string** |  | 

### Return type

[**RemoveDiscordMemberRole200Response**](RemoveDiscordMemberRole200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendDiscordDirectMessage

> SendDiscordDirectMessage200Response SendDiscordDirectMessage(ctx).SendDiscordDirectMessageRequest(sendDiscordDirectMessageRequest).Execute()

Send a Discord Direct Message



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
	sendDiscordDirectMessageRequest := *openapiclient.NewSendDiscordDirectMessageRequest("65a1b2c3d4e5f60718293a4b", "1234567890123456789") // SendDiscordDirectMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.SendDiscordDirectMessage(context.Background()).SendDiscordDirectMessageRequest(sendDiscordDirectMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.SendDiscordDirectMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendDiscordDirectMessage`: SendDiscordDirectMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.SendDiscordDirectMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendDiscordDirectMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendDiscordDirectMessageRequest** | [**SendDiscordDirectMessageRequest**](SendDiscordDirectMessageRequest.md) |  | 

### Return type

[**SendDiscordDirectMessage200Response**](SendDiscordDirectMessage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpinDiscordMessage

> UnpinDiscordMessage200Response UnpinDiscordMessage(ctx, channelId, messageId).AccountId(accountId).Execute()

Unpin a Discord message



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
	channelId := "channelId_example" // string | 
	messageId := "messageId_example" // string | 
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.UnpinDiscordMessage(context.Background(), channelId, messageId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.UnpinDiscordMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnpinDiscordMessage`: UnpinDiscordMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.UnpinDiscordMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpinDiscordMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **string** |  | 

### Return type

[**UnpinDiscordMessage200Response**](UnpinDiscordMessage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDiscordScheduledEvent

> CreateDiscordScheduledEvent200Response UpdateDiscordScheduledEvent(ctx, guildId, eventId).UpdateDiscordScheduledEventRequest(updateDiscordScheduledEventRequest).Execute()

Update a Discord scheduled event



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
	guildId := "guildId_example" // string | 
	eventId := "eventId_example" // string | 
	updateDiscordScheduledEventRequest := *openapiclient.NewUpdateDiscordScheduledEventRequest("AccountId_example") // UpdateDiscordScheduledEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.UpdateDiscordScheduledEvent(context.Background(), guildId, eventId).UpdateDiscordScheduledEventRequest(updateDiscordScheduledEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.UpdateDiscordScheduledEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDiscordScheduledEvent`: CreateDiscordScheduledEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.UpdateDiscordScheduledEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guildId** | **string** |  | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDiscordScheduledEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDiscordScheduledEventRequest** | [**UpdateDiscordScheduledEventRequest**](UpdateDiscordScheduledEventRequest.md) |  | 

### Return type

[**CreateDiscordScheduledEvent200Response**](CreateDiscordScheduledEvent200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDiscordSettings

> UpdateDiscordSettings200Response UpdateDiscordSettings(ctx, accountId).UpdateDiscordSettingsRequest(updateDiscordSettingsRequest).Execute()

Update Discord settings



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
	updateDiscordSettingsRequest := *openapiclient.NewUpdateDiscordSettingsRequest() // UpdateDiscordSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscordAPI.UpdateDiscordSettings(context.Background(), accountId).UpdateDiscordSettingsRequest(updateDiscordSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscordAPI.UpdateDiscordSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDiscordSettings`: UpdateDiscordSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `DiscordAPI.UpdateDiscordSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDiscordSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDiscordSettingsRequest** | [**UpdateDiscordSettingsRequest**](UpdateDiscordSettingsRequest.md) |  | 

### Return type

[**UpdateDiscordSettings200Response**](UpdateDiscordSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

