# \WebhookEventsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OnAccountAdsInitialSyncCompleted**](WebhookEventsAPI.md#OnAccountAdsInitialSyncCompleted) | **Post** /account.ads.initial_sync_completed | Ads initial sync completed event
[**OnAccountConnected**](WebhookEventsAPI.md#OnAccountConnected) | **Post** /account.connected | Account connected event
[**OnAccountDisconnected**](WebhookEventsAPI.md#OnAccountDisconnected) | **Post** /account.disconnected | Account disconnected event
[**OnAdStatusChanged**](WebhookEventsAPI.md#OnAdStatusChanged) | **Post** /ad.status_changed | Ad status changed event
[**OnCallEnded**](WebhookEventsAPI.md#OnCallEnded) | **Post** /call.ended | Call ended event
[**OnCallFailed**](WebhookEventsAPI.md#OnCallFailed) | **Post** /call.failed | Call failed event
[**OnCallPermissionRequest**](WebhookEventsAPI.md#OnCallPermissionRequest) | **Post** /call.permission_request | Call permission request reply event
[**OnCallReceived**](WebhookEventsAPI.md#OnCallReceived) | **Post** /call.received | Call received event
[**OnCommentReceived**](WebhookEventsAPI.md#OnCommentReceived) | **Post** /comment.received | Comment received event
[**OnConversationStarted**](WebhookEventsAPI.md#OnConversationStarted) | **Post** /conversation.started | Conversation started event
[**OnLeadReceived**](WebhookEventsAPI.md#OnLeadReceived) | **Post** /lead.received | Lead received event
[**OnMessageDeleted**](WebhookEventsAPI.md#OnMessageDeleted) | **Post** /message.deleted | Message deleted event
[**OnMessageDelivered**](WebhookEventsAPI.md#OnMessageDelivered) | **Post** /message.delivered | Message delivered event
[**OnMessageEdited**](WebhookEventsAPI.md#OnMessageEdited) | **Post** /message.edited | Message edited event
[**OnMessageFailed**](WebhookEventsAPI.md#OnMessageFailed) | **Post** /message.failed | Message delivery failed event
[**OnMessageRead**](WebhookEventsAPI.md#OnMessageRead) | **Post** /message.read | Message read event
[**OnMessageReceived**](WebhookEventsAPI.md#OnMessageReceived) | **Post** /message.received | Message received event
[**OnMessageSent**](WebhookEventsAPI.md#OnMessageSent) | **Post** /message.sent | Message sent event
[**OnPostCancelled**](WebhookEventsAPI.md#OnPostCancelled) | **Post** /post.cancelled | Post cancelled event
[**OnPostExternalCreated**](WebhookEventsAPI.md#OnPostExternalCreated) | **Post** /post.external.created | External post created event
[**OnPostExternalDeleted**](WebhookEventsAPI.md#OnPostExternalDeleted) | **Post** /post.external.deleted | External post deleted event
[**OnPostExternalUpdated**](WebhookEventsAPI.md#OnPostExternalUpdated) | **Post** /post.external.updated | External post updated event
[**OnPostFailed**](WebhookEventsAPI.md#OnPostFailed) | **Post** /post.failed | Post failed event
[**OnPostPartial**](WebhookEventsAPI.md#OnPostPartial) | **Post** /post.partial | Post partial event
[**OnPostPlatformFailed**](WebhookEventsAPI.md#OnPostPlatformFailed) | **Post** /post.platform.failed | Post platform failed event
[**OnPostPlatformPublished**](WebhookEventsAPI.md#OnPostPlatformPublished) | **Post** /post.platform.published | Post platform published event
[**OnPostPublished**](WebhookEventsAPI.md#OnPostPublished) | **Post** /post.published | Post published event
[**OnPostRecycled**](WebhookEventsAPI.md#OnPostRecycled) | **Post** /post.recycled | Post recycled event
[**OnPostScheduled**](WebhookEventsAPI.md#OnPostScheduled) | **Post** /post.scheduled | Post scheduled event
[**OnReactionReceived**](WebhookEventsAPI.md#OnReactionReceived) | **Post** /reaction.received | Reaction received event
[**OnReviewNew**](WebhookEventsAPI.md#OnReviewNew) | **Post** /review.new | Review new event
[**OnReviewUpdated**](WebhookEventsAPI.md#OnReviewUpdated) | **Post** /review.updated | Review updated event
[**OnWebhookTest**](WebhookEventsAPI.md#OnWebhookTest) | **Post** /webhook.test | Webhook test event
[**OnWhatsAppNumberActionRequired**](WebhookEventsAPI.md#OnWhatsAppNumberActionRequired) | **Post** /whatsapp.number.action_required | WhatsApp number action required event
[**OnWhatsAppNumberActivated**](WebhookEventsAPI.md#OnWhatsAppNumberActivated) | **Post** /whatsapp.number.activated | WhatsApp number activated event
[**OnWhatsAppNumberDeclined**](WebhookEventsAPI.md#OnWhatsAppNumberDeclined) | **Post** /whatsapp.number.declined | WhatsApp number declined event
[**OnWhatsAppNumberReactivated**](WebhookEventsAPI.md#OnWhatsAppNumberReactivated) | **Post** /whatsapp.number.reactivated | WhatsApp number reactivated event
[**OnWhatsAppNumberReleased**](WebhookEventsAPI.md#OnWhatsAppNumberReleased) | **Post** /whatsapp.number.released | WhatsApp number released event
[**OnWhatsAppNumberSuspended**](WebhookEventsAPI.md#OnWhatsAppNumberSuspended) | **Post** /whatsapp.number.suspended | WhatsApp number suspended event
[**OnWhatsAppNumberVerificationRequired**](WebhookEventsAPI.md#OnWhatsAppNumberVerificationRequired) | **Post** /whatsapp.number.verification_required | WhatsApp number verification-required event
[**OnWhatsAppTemplateStatusUpdated**](WebhookEventsAPI.md#OnWhatsAppTemplateStatusUpdated) | **Post** /whatsapp.template.status_updated | WhatsApp template status updated event



## OnAccountAdsInitialSyncCompleted

> OnAccountAdsInitialSyncCompleted(ctx).WebhookPayloadAccountAdsInitialSyncCompleted(webhookPayloadAccountAdsInitialSyncCompleted).Execute()

Ads initial sync completed event



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
	webhookPayloadAccountAdsInitialSyncCompleted :=  // WebhookPayloadAccountAdsInitialSyncCompleted | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnAccountAdsInitialSyncCompleted(context.Background()).WebhookPayloadAccountAdsInitialSyncCompleted(webhookPayloadAccountAdsInitialSyncCompleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnAccountAdsInitialSyncCompleted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnAccountAdsInitialSyncCompletedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadAccountAdsInitialSyncCompleted** | [**WebhookPayloadAccountAdsInitialSyncCompleted**](WebhookPayloadAccountAdsInitialSyncCompleted.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnAccountConnected

> OnAccountConnected(ctx).WebhookPayloadAccountConnected(webhookPayloadAccountConnected).Execute()

Account connected event



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
	webhookPayloadAccountConnected :=  // WebhookPayloadAccountConnected | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnAccountConnected(context.Background()).WebhookPayloadAccountConnected(webhookPayloadAccountConnected).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnAccountConnected``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnAccountConnectedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadAccountConnected** | [**WebhookPayloadAccountConnected**](WebhookPayloadAccountConnected.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnAccountDisconnected

> OnAccountDisconnected(ctx).WebhookPayloadAccountDisconnected(webhookPayloadAccountDisconnected).Execute()

Account disconnected event



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
	webhookPayloadAccountDisconnected :=  // WebhookPayloadAccountDisconnected | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnAccountDisconnected(context.Background()).WebhookPayloadAccountDisconnected(webhookPayloadAccountDisconnected).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnAccountDisconnected``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnAccountDisconnectedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadAccountDisconnected** | [**WebhookPayloadAccountDisconnected**](WebhookPayloadAccountDisconnected.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnAdStatusChanged

> OnAdStatusChanged(ctx).WebhookPayloadAdStatusChanged(webhookPayloadAdStatusChanged).Execute()

Ad status changed event



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
	webhookPayloadAdStatusChanged :=  // WebhookPayloadAdStatusChanged | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnAdStatusChanged(context.Background()).WebhookPayloadAdStatusChanged(webhookPayloadAdStatusChanged).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnAdStatusChanged``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnAdStatusChangedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadAdStatusChanged** | [**WebhookPayloadAdStatusChanged**](WebhookPayloadAdStatusChanged.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnCallEnded

> OnCallEnded(ctx).WebhookPayloadCallEnded(webhookPayloadCallEnded).Execute()

Call ended event



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
	webhookPayloadCallEnded :=  // WebhookPayloadCallEnded | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnCallEnded(context.Background()).WebhookPayloadCallEnded(webhookPayloadCallEnded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnCallEnded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnCallEndedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadCallEnded** | [**WebhookPayloadCallEnded**](WebhookPayloadCallEnded.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnCallFailed

> OnCallFailed(ctx).WebhookPayloadCallFailed(webhookPayloadCallFailed).Execute()

Call failed event



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
	webhookPayloadCallFailed :=  // WebhookPayloadCallFailed | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnCallFailed(context.Background()).WebhookPayloadCallFailed(webhookPayloadCallFailed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnCallFailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnCallFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadCallFailed** | [**WebhookPayloadCallFailed**](WebhookPayloadCallFailed.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnCallPermissionRequest

> OnCallPermissionRequest(ctx).WebhookPayloadCallPermissionRequest(webhookPayloadCallPermissionRequest).Execute()

Call permission request reply event



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
	webhookPayloadCallPermissionRequest :=  // WebhookPayloadCallPermissionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnCallPermissionRequest(context.Background()).WebhookPayloadCallPermissionRequest(webhookPayloadCallPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnCallPermissionRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnCallPermissionRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadCallPermissionRequest** | [**WebhookPayloadCallPermissionRequest**](WebhookPayloadCallPermissionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnCallReceived

> OnCallReceived(ctx).WebhookPayloadCallReceived(webhookPayloadCallReceived).Execute()

Call received event



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
	webhookPayloadCallReceived :=  // WebhookPayloadCallReceived | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnCallReceived(context.Background()).WebhookPayloadCallReceived(webhookPayloadCallReceived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnCallReceived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnCallReceivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadCallReceived** | [**WebhookPayloadCallReceived**](WebhookPayloadCallReceived.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnCommentReceived

> OnCommentReceived(ctx).WebhookPayloadComment(webhookPayloadComment).Execute()

Comment received event



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
	webhookPayloadComment :=  // WebhookPayloadComment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnCommentReceived(context.Background()).WebhookPayloadComment(webhookPayloadComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnCommentReceived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnCommentReceivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadComment** | [**WebhookPayloadComment**](WebhookPayloadComment.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnConversationStarted

> OnConversationStarted(ctx).WebhookPayloadConversationStarted(webhookPayloadConversationStarted).Execute()

Conversation started event



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
	webhookPayloadConversationStarted :=  // WebhookPayloadConversationStarted | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnConversationStarted(context.Background()).WebhookPayloadConversationStarted(webhookPayloadConversationStarted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnConversationStarted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnConversationStartedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadConversationStarted** | [**WebhookPayloadConversationStarted**](WebhookPayloadConversationStarted.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnLeadReceived

> OnLeadReceived(ctx).WebhookPayloadLead(webhookPayloadLead).Execute()

Lead received event



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
	webhookPayloadLead :=  // WebhookPayloadLead | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnLeadReceived(context.Background()).WebhookPayloadLead(webhookPayloadLead).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnLeadReceived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnLeadReceivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadLead** | [**WebhookPayloadLead**](WebhookPayloadLead.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnMessageDeleted

> OnMessageDeleted(ctx).WebhookPayloadMessageDeleted(webhookPayloadMessageDeleted).Execute()

Message deleted event



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
	webhookPayloadMessageDeleted :=  // WebhookPayloadMessageDeleted | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnMessageDeleted(context.Background()).WebhookPayloadMessageDeleted(webhookPayloadMessageDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnMessageDeleted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnMessageDeletedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadMessageDeleted** | [**WebhookPayloadMessageDeleted**](WebhookPayloadMessageDeleted.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnMessageDelivered

> OnMessageDelivered(ctx).WebhookPayloadMessageDeliveryStatus(webhookPayloadMessageDeliveryStatus).Execute()

Message delivered event



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
	webhookPayloadMessageDeliveryStatus :=  // WebhookPayloadMessageDeliveryStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnMessageDelivered(context.Background()).WebhookPayloadMessageDeliveryStatus(webhookPayloadMessageDeliveryStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnMessageDelivered``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnMessageDeliveredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadMessageDeliveryStatus** | [**WebhookPayloadMessageDeliveryStatus**](WebhookPayloadMessageDeliveryStatus.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnMessageEdited

> OnMessageEdited(ctx).WebhookPayloadMessageEdited(webhookPayloadMessageEdited).Execute()

Message edited event



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
	webhookPayloadMessageEdited :=  // WebhookPayloadMessageEdited | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnMessageEdited(context.Background()).WebhookPayloadMessageEdited(webhookPayloadMessageEdited).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnMessageEdited``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnMessageEditedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadMessageEdited** | [**WebhookPayloadMessageEdited**](WebhookPayloadMessageEdited.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnMessageFailed

> OnMessageFailed(ctx).WebhookPayloadMessageDeliveryStatus(webhookPayloadMessageDeliveryStatus).Execute()

Message delivery failed event



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
	webhookPayloadMessageDeliveryStatus :=  // WebhookPayloadMessageDeliveryStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnMessageFailed(context.Background()).WebhookPayloadMessageDeliveryStatus(webhookPayloadMessageDeliveryStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnMessageFailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnMessageFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadMessageDeliveryStatus** | [**WebhookPayloadMessageDeliveryStatus**](WebhookPayloadMessageDeliveryStatus.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnMessageRead

> OnMessageRead(ctx).WebhookPayloadMessageDeliveryStatus(webhookPayloadMessageDeliveryStatus).Execute()

Message read event



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
	webhookPayloadMessageDeliveryStatus :=  // WebhookPayloadMessageDeliveryStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnMessageRead(context.Background()).WebhookPayloadMessageDeliveryStatus(webhookPayloadMessageDeliveryStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnMessageRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnMessageReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadMessageDeliveryStatus** | [**WebhookPayloadMessageDeliveryStatus**](WebhookPayloadMessageDeliveryStatus.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnMessageReceived

> OnMessageReceived(ctx).WebhookPayloadMessage(webhookPayloadMessage).Execute()

Message received event



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
	webhookPayloadMessage :=  // WebhookPayloadMessage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnMessageReceived(context.Background()).WebhookPayloadMessage(webhookPayloadMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnMessageReceived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnMessageReceivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadMessage** | [**WebhookPayloadMessage**](WebhookPayloadMessage.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnMessageSent

> OnMessageSent(ctx).WebhookPayloadMessageSent(webhookPayloadMessageSent).Execute()

Message sent event



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
	webhookPayloadMessageSent :=  // WebhookPayloadMessageSent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnMessageSent(context.Background()).WebhookPayloadMessageSent(webhookPayloadMessageSent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnMessageSent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnMessageSentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadMessageSent** | [**WebhookPayloadMessageSent**](WebhookPayloadMessageSent.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPostCancelled

> OnPostCancelled(ctx).WebhookPayloadPost(webhookPayloadPost).Execute()

Post cancelled event



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
	webhookPayloadPost :=  // WebhookPayloadPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnPostCancelled(context.Background()).WebhookPayloadPost(webhookPayloadPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnPostCancelled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPostCancelledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadPost** | [**WebhookPayloadPost**](WebhookPayloadPost.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPostExternalCreated

> OnPostExternalCreated(ctx).WebhookPayloadExternalPost(webhookPayloadExternalPost).Execute()

External post created event



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
	webhookPayloadExternalPost :=  // WebhookPayloadExternalPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnPostExternalCreated(context.Background()).WebhookPayloadExternalPost(webhookPayloadExternalPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnPostExternalCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPostExternalCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadExternalPost** | [**WebhookPayloadExternalPost**](WebhookPayloadExternalPost.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPostExternalDeleted

> OnPostExternalDeleted(ctx).WebhookPayloadExternalPost(webhookPayloadExternalPost).Execute()

External post deleted event



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
	webhookPayloadExternalPost :=  // WebhookPayloadExternalPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnPostExternalDeleted(context.Background()).WebhookPayloadExternalPost(webhookPayloadExternalPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnPostExternalDeleted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPostExternalDeletedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadExternalPost** | [**WebhookPayloadExternalPost**](WebhookPayloadExternalPost.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPostExternalUpdated

> OnPostExternalUpdated(ctx).WebhookPayloadExternalPost(webhookPayloadExternalPost).Execute()

External post updated event



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
	webhookPayloadExternalPost :=  // WebhookPayloadExternalPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnPostExternalUpdated(context.Background()).WebhookPayloadExternalPost(webhookPayloadExternalPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnPostExternalUpdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPostExternalUpdatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadExternalPost** | [**WebhookPayloadExternalPost**](WebhookPayloadExternalPost.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPostFailed

> OnPostFailed(ctx).WebhookPayloadPost(webhookPayloadPost).Execute()

Post failed event



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
	webhookPayloadPost :=  // WebhookPayloadPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnPostFailed(context.Background()).WebhookPayloadPost(webhookPayloadPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnPostFailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPostFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadPost** | [**WebhookPayloadPost**](WebhookPayloadPost.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPostPartial

> OnPostPartial(ctx).WebhookPayloadPost(webhookPayloadPost).Execute()

Post partial event



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
	webhookPayloadPost :=  // WebhookPayloadPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnPostPartial(context.Background()).WebhookPayloadPost(webhookPayloadPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnPostPartial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPostPartialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadPost** | [**WebhookPayloadPost**](WebhookPayloadPost.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPostPlatformFailed

> OnPostPlatformFailed(ctx).WebhookPayloadPostPlatform(webhookPayloadPostPlatform).Execute()

Post platform failed event



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
	webhookPayloadPostPlatform :=  // WebhookPayloadPostPlatform | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnPostPlatformFailed(context.Background()).WebhookPayloadPostPlatform(webhookPayloadPostPlatform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnPostPlatformFailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPostPlatformFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadPostPlatform** | [**WebhookPayloadPostPlatform**](WebhookPayloadPostPlatform.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPostPlatformPublished

> OnPostPlatformPublished(ctx).WebhookPayloadPostPlatform(webhookPayloadPostPlatform).Execute()

Post platform published event



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
	webhookPayloadPostPlatform :=  // WebhookPayloadPostPlatform | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnPostPlatformPublished(context.Background()).WebhookPayloadPostPlatform(webhookPayloadPostPlatform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnPostPlatformPublished``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPostPlatformPublishedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadPostPlatform** | [**WebhookPayloadPostPlatform**](WebhookPayloadPostPlatform.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPostPublished

> OnPostPublished(ctx).WebhookPayloadPost(webhookPayloadPost).Execute()

Post published event



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
	webhookPayloadPost :=  // WebhookPayloadPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnPostPublished(context.Background()).WebhookPayloadPost(webhookPayloadPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnPostPublished``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPostPublishedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadPost** | [**WebhookPayloadPost**](WebhookPayloadPost.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPostRecycled

> OnPostRecycled(ctx).WebhookPayloadPost(webhookPayloadPost).Execute()

Post recycled event



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
	webhookPayloadPost :=  // WebhookPayloadPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnPostRecycled(context.Background()).WebhookPayloadPost(webhookPayloadPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnPostRecycled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPostRecycledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadPost** | [**WebhookPayloadPost**](WebhookPayloadPost.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPostScheduled

> OnPostScheduled(ctx).WebhookPayloadPost(webhookPayloadPost).Execute()

Post scheduled event



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
	webhookPayloadPost :=  // WebhookPayloadPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnPostScheduled(context.Background()).WebhookPayloadPost(webhookPayloadPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnPostScheduled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnPostScheduledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadPost** | [**WebhookPayloadPost**](WebhookPayloadPost.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnReactionReceived

> OnReactionReceived(ctx).WebhookPayloadReaction(webhookPayloadReaction).Execute()

Reaction received event



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
	webhookPayloadReaction :=  // WebhookPayloadReaction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnReactionReceived(context.Background()).WebhookPayloadReaction(webhookPayloadReaction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnReactionReceived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnReactionReceivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadReaction** | [**WebhookPayloadReaction**](WebhookPayloadReaction.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnReviewNew

> OnReviewNew(ctx).WebhookPayloadReviewNew(webhookPayloadReviewNew).Execute()

Review new event



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
	webhookPayloadReviewNew :=  // WebhookPayloadReviewNew | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnReviewNew(context.Background()).WebhookPayloadReviewNew(webhookPayloadReviewNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnReviewNew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnReviewNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadReviewNew** | [**WebhookPayloadReviewNew**](WebhookPayloadReviewNew.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnReviewUpdated

> OnReviewUpdated(ctx).WebhookPayloadReviewUpdated(webhookPayloadReviewUpdated).Execute()

Review updated event



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
	webhookPayloadReviewUpdated :=  // WebhookPayloadReviewUpdated | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnReviewUpdated(context.Background()).WebhookPayloadReviewUpdated(webhookPayloadReviewUpdated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnReviewUpdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnReviewUpdatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadReviewUpdated** | [**WebhookPayloadReviewUpdated**](WebhookPayloadReviewUpdated.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnWebhookTest

> OnWebhookTest(ctx).WebhookPayloadTest(webhookPayloadTest).Execute()

Webhook test event



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
	webhookPayloadTest :=  // WebhookPayloadTest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnWebhookTest(context.Background()).WebhookPayloadTest(webhookPayloadTest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnWebhookTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnWebhookTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadTest** | [**WebhookPayloadTest**](WebhookPayloadTest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnWhatsAppNumberActionRequired

> OnWhatsAppNumberActionRequired(ctx).OnWhatsAppNumberActionRequiredRequest(onWhatsAppNumberActionRequiredRequest).Execute()

WhatsApp number action required event



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
	onWhatsAppNumberActionRequiredRequest :=  // OnWhatsAppNumberActionRequiredRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnWhatsAppNumberActionRequired(context.Background()).OnWhatsAppNumberActionRequiredRequest(onWhatsAppNumberActionRequiredRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnWhatsAppNumberActionRequired``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnWhatsAppNumberActionRequiredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onWhatsAppNumberActionRequiredRequest** | [**OnWhatsAppNumberActionRequiredRequest**](OnWhatsAppNumberActionRequiredRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnWhatsAppNumberActivated

> OnWhatsAppNumberActivated(ctx).OnWhatsAppNumberActivatedRequest(onWhatsAppNumberActivatedRequest).Execute()

WhatsApp number activated event



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
	onWhatsAppNumberActivatedRequest :=  // OnWhatsAppNumberActivatedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnWhatsAppNumberActivated(context.Background()).OnWhatsAppNumberActivatedRequest(onWhatsAppNumberActivatedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnWhatsAppNumberActivated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnWhatsAppNumberActivatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onWhatsAppNumberActivatedRequest** | [**OnWhatsAppNumberActivatedRequest**](OnWhatsAppNumberActivatedRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnWhatsAppNumberDeclined

> OnWhatsAppNumberDeclined(ctx).OnWhatsAppNumberDeclinedRequest(onWhatsAppNumberDeclinedRequest).Execute()

WhatsApp number declined event



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
	onWhatsAppNumberDeclinedRequest :=  // OnWhatsAppNumberDeclinedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnWhatsAppNumberDeclined(context.Background()).OnWhatsAppNumberDeclinedRequest(onWhatsAppNumberDeclinedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnWhatsAppNumberDeclined``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnWhatsAppNumberDeclinedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onWhatsAppNumberDeclinedRequest** | [**OnWhatsAppNumberDeclinedRequest**](OnWhatsAppNumberDeclinedRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnWhatsAppNumberReactivated

> OnWhatsAppNumberReactivated(ctx).OnWhatsAppNumberReactivatedRequest(onWhatsAppNumberReactivatedRequest).Execute()

WhatsApp number reactivated event



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
	onWhatsAppNumberReactivatedRequest :=  // OnWhatsAppNumberReactivatedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnWhatsAppNumberReactivated(context.Background()).OnWhatsAppNumberReactivatedRequest(onWhatsAppNumberReactivatedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnWhatsAppNumberReactivated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnWhatsAppNumberReactivatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onWhatsAppNumberReactivatedRequest** | [**OnWhatsAppNumberReactivatedRequest**](OnWhatsAppNumberReactivatedRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnWhatsAppNumberReleased

> OnWhatsAppNumberReleased(ctx).OnWhatsAppNumberReleasedRequest(onWhatsAppNumberReleasedRequest).Execute()

WhatsApp number released event



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
	onWhatsAppNumberReleasedRequest :=  // OnWhatsAppNumberReleasedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnWhatsAppNumberReleased(context.Background()).OnWhatsAppNumberReleasedRequest(onWhatsAppNumberReleasedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnWhatsAppNumberReleased``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnWhatsAppNumberReleasedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onWhatsAppNumberReleasedRequest** | [**OnWhatsAppNumberReleasedRequest**](OnWhatsAppNumberReleasedRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnWhatsAppNumberSuspended

> OnWhatsAppNumberSuspended(ctx).OnWhatsAppNumberSuspendedRequest(onWhatsAppNumberSuspendedRequest).Execute()

WhatsApp number suspended event



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
	onWhatsAppNumberSuspendedRequest :=  // OnWhatsAppNumberSuspendedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnWhatsAppNumberSuspended(context.Background()).OnWhatsAppNumberSuspendedRequest(onWhatsAppNumberSuspendedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnWhatsAppNumberSuspended``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnWhatsAppNumberSuspendedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onWhatsAppNumberSuspendedRequest** | [**OnWhatsAppNumberSuspendedRequest**](OnWhatsAppNumberSuspendedRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnWhatsAppNumberVerificationRequired

> OnWhatsAppNumberVerificationRequired(ctx).OnWhatsAppNumberVerificationRequiredRequest(onWhatsAppNumberVerificationRequiredRequest).Execute()

WhatsApp number verification-required event



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
	onWhatsAppNumberVerificationRequiredRequest :=  // OnWhatsAppNumberVerificationRequiredRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnWhatsAppNumberVerificationRequired(context.Background()).OnWhatsAppNumberVerificationRequiredRequest(onWhatsAppNumberVerificationRequiredRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnWhatsAppNumberVerificationRequired``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnWhatsAppNumberVerificationRequiredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onWhatsAppNumberVerificationRequiredRequest** | [**OnWhatsAppNumberVerificationRequiredRequest**](OnWhatsAppNumberVerificationRequiredRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnWhatsAppTemplateStatusUpdated

> OnWhatsAppTemplateStatusUpdated(ctx).WebhookPayloadWhatsAppTemplateStatusUpdated(webhookPayloadWhatsAppTemplateStatusUpdated).Execute()

WhatsApp template status updated event



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
	webhookPayloadWhatsAppTemplateStatusUpdated :=  // WebhookPayloadWhatsAppTemplateStatusUpdated | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookEventsAPI.OnWhatsAppTemplateStatusUpdated(context.Background()).WebhookPayloadWhatsAppTemplateStatusUpdated(webhookPayloadWhatsAppTemplateStatusUpdated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEventsAPI.OnWhatsAppTemplateStatusUpdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnWhatsAppTemplateStatusUpdatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPayloadWhatsAppTemplateStatusUpdated** | [**WebhookPayloadWhatsAppTemplateStatusUpdated**](WebhookPayloadWhatsAppTemplateStatusUpdated.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

