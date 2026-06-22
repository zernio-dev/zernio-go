# \MessagesAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMessageReaction**](MessagesAPI.md#AddMessageReaction) | **Post** /v1/inbox/conversations/{conversationId}/messages/{messageId}/reactions | Add reaction
[**CreateInboxConversation**](MessagesAPI.md#CreateInboxConversation) | **Post** /v1/inbox/conversations | Create conversation (send a WhatsApp template)
[**DeleteInboxMessage**](MessagesAPI.md#DeleteInboxMessage) | **Delete** /v1/inbox/conversations/{conversationId}/messages/{messageId} | Delete message
[**EditInboxMessage**](MessagesAPI.md#EditInboxMessage) | **Patch** /v1/inbox/conversations/{conversationId}/messages/{messageId} | Edit message
[**GetInboxConversation**](MessagesAPI.md#GetInboxConversation) | **Get** /v1/inbox/conversations/{conversationId} | Get conversation
[**GetInboxConversationMessages**](MessagesAPI.md#GetInboxConversationMessages) | **Get** /v1/inbox/conversations/{conversationId}/messages | List messages
[**ListInboxConversations**](MessagesAPI.md#ListInboxConversations) | **Get** /v1/inbox/conversations | List conversations
[**MarkConversationRead**](MessagesAPI.md#MarkConversationRead) | **Post** /v1/inbox/conversations/{conversationId}/read | Mark a conversation as read
[**RemoveMessageReaction**](MessagesAPI.md#RemoveMessageReaction) | **Delete** /v1/inbox/conversations/{conversationId}/messages/{messageId}/reactions | Remove reaction
[**SendInboxMessage**](MessagesAPI.md#SendInboxMessage) | **Post** /v1/inbox/conversations/{conversationId}/messages | Send message
[**SendTypingIndicator**](MessagesAPI.md#SendTypingIndicator) | **Post** /v1/inbox/conversations/{conversationId}/typing | Send typing indicator
[**UpdateInboxConversation**](MessagesAPI.md#UpdateInboxConversation) | **Put** /v1/inbox/conversations/{conversationId} | Update conversation status
[**UploadMediaDirect**](MessagesAPI.md#UploadMediaDirect) | **Post** /v1/media/upload-direct | Upload media file



## AddMessageReaction

> UpdateYoutubeDefaultPlaylist200Response AddMessageReaction(ctx, conversationId, messageId).AddMessageReactionRequest(addMessageReactionRequest).Execute()

Add reaction



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
	conversationId := "conversationId_example" // string | The conversation ID
	messageId := "messageId_example" // string | The platform message ID to react to
	addMessageReactionRequest := *openapiclient.NewAddMessageReactionRequest("AccountId_example", "👍") // AddMessageReactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.AddMessageReaction(context.Background(), conversationId, messageId).AddMessageReactionRequest(addMessageReactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.AddMessageReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMessageReaction`: UpdateYoutubeDefaultPlaylist200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.AddMessageReaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The conversation ID | 
**messageId** | **string** | The platform message ID to react to | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMessageReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addMessageReactionRequest** | [**AddMessageReactionRequest**](AddMessageReactionRequest.md) |  | 

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


## CreateInboxConversation

> CreateInboxConversation201Response CreateInboxConversation(ctx).CreateInboxConversationRequest(createInboxConversationRequest).Execute()

Create conversation (send a WhatsApp template)



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
	createInboxConversationRequest := *openapiclient.NewCreateInboxConversationRequest("AccountId_example") // CreateInboxConversationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.CreateInboxConversation(context.Background()).CreateInboxConversationRequest(createInboxConversationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.CreateInboxConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInboxConversation`: CreateInboxConversation201Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.CreateInboxConversation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInboxConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInboxConversationRequest** | [**CreateInboxConversationRequest**](CreateInboxConversationRequest.md) |  | 

### Return type

[**CreateInboxConversation201Response**](CreateInboxConversation201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInboxMessage

> UpdateYoutubeDefaultPlaylist200Response DeleteInboxMessage(ctx, conversationId, messageId).AccountId(accountId).Execute()

Delete message



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
	conversationId := "conversationId_example" // string | The conversation ID
	messageId := "messageId_example" // string | The platform message ID to delete
	accountId := "accountId_example" // string | Social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.DeleteInboxMessage(context.Background(), conversationId, messageId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.DeleteInboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInboxMessage`: UpdateYoutubeDefaultPlaylist200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.DeleteInboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The conversation ID | 
**messageId** | **string** | The platform message ID to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **string** | Social account ID | 

### Return type

[**UpdateYoutubeDefaultPlaylist200Response**](UpdateYoutubeDefaultPlaylist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditInboxMessage

> EditInboxMessage200Response EditInboxMessage(ctx, conversationId, messageId).EditInboxMessageRequest(editInboxMessageRequest).Execute()

Edit message



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
	conversationId := "conversationId_example" // string | The conversation ID
	messageId := "messageId_example" // string | The Telegram message ID to edit
	editInboxMessageRequest := *openapiclient.NewEditInboxMessageRequest("AccountId_example") // EditInboxMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.EditInboxMessage(context.Background(), conversationId, messageId).EditInboxMessageRequest(editInboxMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.EditInboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditInboxMessage`: EditInboxMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.EditInboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The conversation ID | 
**messageId** | **string** | The Telegram message ID to edit | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditInboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **editInboxMessageRequest** | [**EditInboxMessageRequest**](EditInboxMessageRequest.md) |  | 

### Return type

[**EditInboxMessage200Response**](EditInboxMessage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxConversation

> GetInboxConversation200Response GetInboxConversation(ctx, conversationId).AccountId(accountId).Execute()

Get conversation



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
	conversationId := "conversationId_example" // string | The conversation ID (id field from list conversations endpoint). This is the platform-specific conversation identifier, not an internal database ID.
	accountId := "accountId_example" // string | The social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.GetInboxConversation(context.Background(), conversationId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetInboxConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxConversation`: GetInboxConversation200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetInboxConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The conversation ID (id field from list conversations endpoint). This is the platform-specific conversation identifier, not an internal database ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | The social account ID | 

### Return type

[**GetInboxConversation200Response**](GetInboxConversation200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxConversationMessages

> GetInboxConversationMessages200Response GetInboxConversationMessages(ctx, conversationId).AccountId(accountId).Limit(limit).Cursor(cursor).SortOrder(sortOrder).Execute()

List messages



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
	conversationId := "conversationId_example" // string | The conversation ID (id field from list conversations endpoint). This is the platform-specific conversation identifier, not an internal database ID.
	accountId := "accountId_example" // string | Social account ID
	limit := int32(56) // int32 | Number of messages to return per page. Default 100, max 100. (optional) (default to 100)
	cursor := "cursor_example" // string | Opaque pagination cursor. Pass `pagination.nextCursor` from a prior response. (optional)
	sortOrder := "sortOrder_example" // string | Order of returned messages. Default `asc` (oldest first, chat style). Twitter, Instagram, Telegram, WhatsApp and Reddit honor this order across cursor pages. For Facebook and Bluesky, only intra-page ordering is affected — pages always walk newest→oldest. See `sortOrderApplied` in the response.  (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.GetInboxConversationMessages(context.Background(), conversationId).AccountId(accountId).Limit(limit).Cursor(cursor).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetInboxConversationMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxConversationMessages`: GetInboxConversationMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetInboxConversationMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The conversation ID (id field from list conversations endpoint). This is the platform-specific conversation identifier, not an internal database ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxConversationMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | Social account ID | 
 **limit** | **int32** | Number of messages to return per page. Default 100, max 100. | [default to 100]
 **cursor** | **string** | Opaque pagination cursor. Pass &#x60;pagination.nextCursor&#x60; from a prior response. | 
 **sortOrder** | **string** | Order of returned messages. Default &#x60;asc&#x60; (oldest first, chat style). Twitter, Instagram, Telegram, WhatsApp and Reddit honor this order across cursor pages. For Facebook and Bluesky, only intra-page ordering is affected — pages always walk newest→oldest. See &#x60;sortOrderApplied&#x60; in the response.  | [default to &quot;asc&quot;]

### Return type

[**GetInboxConversationMessages200Response**](GetInboxConversationMessages200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInboxConversations

> ListInboxConversations200Response ListInboxConversations(ctx).ProfileId(profileId).Platform(platform).Status(status).SortOrder(sortOrder).Limit(limit).Cursor(cursor).AccountId(accountId).Execute()

List conversations



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
	status := "status_example" // string | Filter by conversation status (optional)
	sortOrder := "sortOrder_example" // string | Sort order by updated time (optional) (default to "desc")
	limit := int32(56) // int32 | Maximum number of conversations to return (optional) (default to 50)
	cursor := "cursor_example" // string | Pagination cursor for next page (optional)
	accountId := "accountId_example" // string | Filter by specific social account ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.ListInboxConversations(context.Background()).ProfileId(profileId).Platform(platform).Status(status).SortOrder(sortOrder).Limit(limit).Cursor(cursor).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.ListInboxConversations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInboxConversations`: ListInboxConversations200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.ListInboxConversations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInboxConversationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Filter by profile ID | 
 **platform** | **string** | Filter by platform | 
 **status** | **string** | Filter by conversation status | 
 **sortOrder** | **string** | Sort order by updated time | [default to &quot;desc&quot;]
 **limit** | **int32** | Maximum number of conversations to return | [default to 50]
 **cursor** | **string** | Pagination cursor for next page | 
 **accountId** | **string** | Filter by specific social account ID | 

### Return type

[**ListInboxConversations200Response**](ListInboxConversations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkConversationRead

> MarkConversationRead200Response MarkConversationRead(ctx, conversationId).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()

Mark a conversation as read



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
	conversationId := "conversationId_example" // string | The conversation ID
	sendTypingIndicatorRequest := *openapiclient.NewSendTypingIndicatorRequest("AccountId_example") // SendTypingIndicatorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.MarkConversationRead(context.Background(), conversationId).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MarkConversationRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkConversationRead`: MarkConversationRead200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MarkConversationRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The conversation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkConversationReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendTypingIndicatorRequest** | [**SendTypingIndicatorRequest**](SendTypingIndicatorRequest.md) |  | 

### Return type

[**MarkConversationRead200Response**](MarkConversationRead200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMessageReaction

> UpdateYoutubeDefaultPlaylist200Response RemoveMessageReaction(ctx, conversationId, messageId).AccountId(accountId).Execute()

Remove reaction



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
	conversationId := "conversationId_example" // string | The conversation ID
	messageId := "messageId_example" // string | The platform message ID
	accountId := "accountId_example" // string | Social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.RemoveMessageReaction(context.Background(), conversationId, messageId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.RemoveMessageReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveMessageReaction`: UpdateYoutubeDefaultPlaylist200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.RemoveMessageReaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The conversation ID | 
**messageId** | **string** | The platform message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMessageReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **string** | Social account ID | 

### Return type

[**UpdateYoutubeDefaultPlaylist200Response**](UpdateYoutubeDefaultPlaylist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendInboxMessage

> SendInboxMessage200Response SendInboxMessage(ctx, conversationId).SendInboxMessageRequest(sendInboxMessageRequest).Execute()

Send message



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
	conversationId := "conversationId_example" // string | The conversation ID (id field from list conversations endpoint). This is the platform-specific conversation identifier, not an internal database ID.
	sendInboxMessageRequest := *openapiclient.NewSendInboxMessageRequest("AccountId_example") // SendInboxMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.SendInboxMessage(context.Background(), conversationId).SendInboxMessageRequest(sendInboxMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.SendInboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendInboxMessage`: SendInboxMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.SendInboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The conversation ID (id field from list conversations endpoint). This is the platform-specific conversation identifier, not an internal database ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendInboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendInboxMessageRequest** | [**SendInboxMessageRequest**](SendInboxMessageRequest.md) |  | 

### Return type

[**SendInboxMessage200Response**](SendInboxMessage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTypingIndicator

> UpdateYoutubeDefaultPlaylist200Response SendTypingIndicator(ctx, conversationId).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()

Send typing indicator



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
	conversationId := "conversationId_example" // string | The conversation ID
	sendTypingIndicatorRequest := *openapiclient.NewSendTypingIndicatorRequest("AccountId_example") // SendTypingIndicatorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.SendTypingIndicator(context.Background(), conversationId).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.SendTypingIndicator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTypingIndicator`: UpdateYoutubeDefaultPlaylist200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.SendTypingIndicator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The conversation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTypingIndicatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendTypingIndicatorRequest** | [**SendTypingIndicatorRequest**](SendTypingIndicatorRequest.md) |  | 

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


## UpdateInboxConversation

> UpdateInboxConversation200Response UpdateInboxConversation(ctx, conversationId).UpdateInboxConversationRequest(updateInboxConversationRequest).Execute()

Update conversation status



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
	conversationId := "conversationId_example" // string | The conversation ID (id field from list conversations endpoint). This is the platform-specific conversation identifier, not an internal database ID.
	updateInboxConversationRequest := *openapiclient.NewUpdateInboxConversationRequest("AccountId_example", "Status_example") // UpdateInboxConversationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.UpdateInboxConversation(context.Background(), conversationId).UpdateInboxConversationRequest(updateInboxConversationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.UpdateInboxConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInboxConversation`: UpdateInboxConversation200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.UpdateInboxConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The conversation ID (id field from list conversations endpoint). This is the platform-specific conversation identifier, not an internal database ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInboxConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInboxConversationRequest** | [**UpdateInboxConversationRequest**](UpdateInboxConversationRequest.md) |  | 

### Return type

[**UpdateInboxConversation200Response**](UpdateInboxConversation200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadMediaDirect

> UploadMediaDirect200Response UploadMediaDirect(ctx).File(file).ContentType(contentType).Execute()

Upload media file



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
	file := os.NewFile(1234, "some_file") // *os.File | The file to upload (max 25MB)
	contentType := "contentType_example" // string | Override MIME type (e.g. \\\"image/jpeg\\\"). Auto-detected from file if not provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.UploadMediaDirect(context.Background()).File(file).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.UploadMediaDirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadMediaDirect`: UploadMediaDirect200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.UploadMediaDirect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadMediaDirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The file to upload (max 25MB) | 
 **contentType** | **string** | Override MIME type (e.g. \\\&quot;image/jpeg\\\&quot;). Auto-detected from file if not provided. | 

### Return type

[**UploadMediaDirect200Response**](UploadMediaDirect200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

