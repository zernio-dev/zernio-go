# \WhatsAppAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWhatsAppGroupParticipants**](WhatsAppAPI.md#AddWhatsAppGroupParticipants) | **Post** /v1/whatsapp/wa-groups/{groupId}/participants | Add participants
[**ApproveWhatsAppGroupJoinRequests**](WhatsAppAPI.md#ApproveWhatsAppGroupJoinRequests) | **Post** /v1/whatsapp/wa-groups/{groupId}/join-requests | Approve join requests
[**BlockWhatsAppUsers**](WhatsAppAPI.md#BlockWhatsAppUsers) | **Post** /v1/whatsapp/block-users | Block users
[**CreateWhatsAppDataset**](WhatsAppAPI.md#CreateWhatsAppDataset) | **Post** /v1/whatsapp/dataset | Provision CTWA conversions dataset
[**CreateWhatsAppGroupChat**](WhatsAppAPI.md#CreateWhatsAppGroupChat) | **Post** /v1/whatsapp/wa-groups | Create group
[**CreateWhatsAppGroupInviteLink**](WhatsAppAPI.md#CreateWhatsAppGroupInviteLink) | **Post** /v1/whatsapp/wa-groups/{groupId}/invite-link | Create invite link
[**CreateWhatsAppTemplate**](WhatsAppAPI.md#CreateWhatsAppTemplate) | **Post** /v1/whatsapp/templates | Create template
[**DeleteWhatsAppGroupChat**](WhatsAppAPI.md#DeleteWhatsAppGroupChat) | **Delete** /v1/whatsapp/wa-groups/{groupId} | Delete group
[**DeleteWhatsAppTemplate**](WhatsAppAPI.md#DeleteWhatsAppTemplate) | **Delete** /v1/whatsapp/templates/{templateName} | Delete template
[**GetWhatsAppBlockStatus**](WhatsAppAPI.md#GetWhatsAppBlockStatus) | **Get** /v1/whatsapp/block-users/status | Check if a user is blocked
[**GetWhatsAppBlockedUsers**](WhatsAppAPI.md#GetWhatsAppBlockedUsers) | **Get** /v1/whatsapp/block-users | List blocked users
[**GetWhatsAppBusinessProfile**](WhatsAppAPI.md#GetWhatsAppBusinessProfile) | **Get** /v1/whatsapp/business-profile | Get business profile
[**GetWhatsAppDataset**](WhatsAppAPI.md#GetWhatsAppDataset) | **Get** /v1/whatsapp/dataset | Get CTWA conversions dataset
[**GetWhatsAppDisplayName**](WhatsAppAPI.md#GetWhatsAppDisplayName) | **Get** /v1/whatsapp/business-profile/display-name | Get display name status
[**GetWhatsAppGroupChat**](WhatsAppAPI.md#GetWhatsAppGroupChat) | **Get** /v1/whatsapp/wa-groups/{groupId} | Get group info
[**GetWhatsAppTemplate**](WhatsAppAPI.md#GetWhatsAppTemplate) | **Get** /v1/whatsapp/templates/{templateName} | Get template
[**GetWhatsAppTemplates**](WhatsAppAPI.md#GetWhatsAppTemplates) | **Get** /v1/whatsapp/templates | List templates
[**ListWhatsAppConversions**](WhatsAppAPI.md#ListWhatsAppConversions) | **Get** /v1/whatsapp/conversions | List recent WhatsApp conversion events
[**ListWhatsAppGroupChats**](WhatsAppAPI.md#ListWhatsAppGroupChats) | **Get** /v1/whatsapp/wa-groups | List active groups
[**ListWhatsAppGroupJoinRequests**](WhatsAppAPI.md#ListWhatsAppGroupJoinRequests) | **Get** /v1/whatsapp/wa-groups/{groupId}/join-requests | List join requests
[**RejectWhatsAppGroupJoinRequests**](WhatsAppAPI.md#RejectWhatsAppGroupJoinRequests) | **Delete** /v1/whatsapp/wa-groups/{groupId}/join-requests | Reject join requests
[**RemoveWhatsAppGroupParticipants**](WhatsAppAPI.md#RemoveWhatsAppGroupParticipants) | **Delete** /v1/whatsapp/wa-groups/{groupId}/participants | Remove participants
[**SendWhatsAppConversion**](WhatsAppAPI.md#SendWhatsAppConversion) | **Post** /v1/whatsapp/conversions | Send WhatsApp conversion event
[**UnblockWhatsAppUsers**](WhatsAppAPI.md#UnblockWhatsAppUsers) | **Delete** /v1/whatsapp/block-users | Unblock users
[**UpdateWhatsAppBusinessProfile**](WhatsAppAPI.md#UpdateWhatsAppBusinessProfile) | **Post** /v1/whatsapp/business-profile | Update business profile
[**UpdateWhatsAppDisplayName**](WhatsAppAPI.md#UpdateWhatsAppDisplayName) | **Post** /v1/whatsapp/business-profile/display-name | Request display name change
[**UpdateWhatsAppGroupChat**](WhatsAppAPI.md#UpdateWhatsAppGroupChat) | **Post** /v1/whatsapp/wa-groups/{groupId} | Update group settings
[**UpdateWhatsAppTemplate**](WhatsAppAPI.md#UpdateWhatsAppTemplate) | **Patch** /v1/whatsapp/templates/{templateName} | Update template
[**UploadWhatsAppProfilePhoto**](WhatsAppAPI.md#UploadWhatsAppProfilePhoto) | **Post** /v1/whatsapp/business-profile/photo | Upload profile picture



## AddWhatsAppGroupParticipants

> UnpublishPost200Response AddWhatsAppGroupParticipants(ctx, groupId).AccountId(accountId).AddWhatsAppGroupParticipantsRequest(addWhatsAppGroupParticipantsRequest).Execute()

Add participants



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
	groupId := "groupId_example" // string | Group ID
	accountId := "accountId_example" // string | WhatsApp social account ID
	addWhatsAppGroupParticipantsRequest := *openapiclient.NewAddWhatsAppGroupParticipantsRequest([]string{"PhoneNumbers_example"}) // AddWhatsAppGroupParticipantsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.AddWhatsAppGroupParticipants(context.Background(), groupId).AccountId(accountId).AddWhatsAppGroupParticipantsRequest(addWhatsAppGroupParticipantsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.AddWhatsAppGroupParticipants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddWhatsAppGroupParticipants`: UnpublishPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.AddWhatsAppGroupParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWhatsAppGroupParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 
 **addWhatsAppGroupParticipantsRequest** | [**AddWhatsAppGroupParticipantsRequest**](AddWhatsAppGroupParticipantsRequest.md) |  | 

### Return type

[**UnpublishPost200Response**](UnpublishPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApproveWhatsAppGroupJoinRequests

> UnpublishPost200Response ApproveWhatsAppGroupJoinRequests(ctx, groupId).AccountId(accountId).RemoveWhatsAppGroupParticipantsRequest(removeWhatsAppGroupParticipantsRequest).Execute()

Approve join requests



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
	groupId := "groupId_example" // string | Group ID
	accountId := "accountId_example" // string | WhatsApp social account ID
	removeWhatsAppGroupParticipantsRequest := *openapiclient.NewRemoveWhatsAppGroupParticipantsRequest([]string{"PhoneNumbers_example"}) // RemoveWhatsAppGroupParticipantsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.ApproveWhatsAppGroupJoinRequests(context.Background(), groupId).AccountId(accountId).RemoveWhatsAppGroupParticipantsRequest(removeWhatsAppGroupParticipantsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.ApproveWhatsAppGroupJoinRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveWhatsAppGroupJoinRequests`: UnpublishPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.ApproveWhatsAppGroupJoinRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveWhatsAppGroupJoinRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 
 **removeWhatsAppGroupParticipantsRequest** | [**RemoveWhatsAppGroupParticipantsRequest**](RemoveWhatsAppGroupParticipantsRequest.md) |  | 

### Return type

[**UnpublishPost200Response**](UnpublishPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockWhatsAppUsers

> BlockWhatsAppUsers200Response BlockWhatsAppUsers(ctx).BlockWhatsAppUsersRequest(blockWhatsAppUsersRequest).Execute()

Block users



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
	blockWhatsAppUsersRequest := *openapiclient.NewBlockWhatsAppUsersRequest("AccountId_example", []string{"Users_example"}) // BlockWhatsAppUsersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.BlockWhatsAppUsers(context.Background()).BlockWhatsAppUsersRequest(blockWhatsAppUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.BlockWhatsAppUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockWhatsAppUsers`: BlockWhatsAppUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.BlockWhatsAppUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlockWhatsAppUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockWhatsAppUsersRequest** | [**BlockWhatsAppUsersRequest**](BlockWhatsAppUsersRequest.md) |  | 

### Return type

[**BlockWhatsAppUsers200Response**](BlockWhatsAppUsers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWhatsAppDataset

> CreateWhatsAppDataset200Response CreateWhatsAppDataset(ctx).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()

Provision CTWA conversions dataset



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
	sendTypingIndicatorRequest := *openapiclient.NewSendTypingIndicatorRequest("AccountId_example") // SendTypingIndicatorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.CreateWhatsAppDataset(context.Background()).SendTypingIndicatorRequest(sendTypingIndicatorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.CreateWhatsAppDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWhatsAppDataset`: CreateWhatsAppDataset200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.CreateWhatsAppDataset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWhatsAppDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendTypingIndicatorRequest** | [**SendTypingIndicatorRequest**](SendTypingIndicatorRequest.md) |  | 

### Return type

[**CreateWhatsAppDataset200Response**](CreateWhatsAppDataset200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWhatsAppGroupChat

> CreateWhatsAppGroupChat201Response CreateWhatsAppGroupChat(ctx).CreateWhatsAppGroupChatRequest(createWhatsAppGroupChatRequest).Execute()

Create group



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
	createWhatsAppGroupChatRequest := *openapiclient.NewCreateWhatsAppGroupChatRequest("AccountId_example", "Subject_example") // CreateWhatsAppGroupChatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.CreateWhatsAppGroupChat(context.Background()).CreateWhatsAppGroupChatRequest(createWhatsAppGroupChatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.CreateWhatsAppGroupChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWhatsAppGroupChat`: CreateWhatsAppGroupChat201Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.CreateWhatsAppGroupChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWhatsAppGroupChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWhatsAppGroupChatRequest** | [**CreateWhatsAppGroupChatRequest**](CreateWhatsAppGroupChatRequest.md) |  | 

### Return type

[**CreateWhatsAppGroupChat201Response**](CreateWhatsAppGroupChat201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWhatsAppGroupInviteLink

> CreateWhatsAppGroupInviteLink200Response CreateWhatsAppGroupInviteLink(ctx, groupId).AccountId(accountId).Execute()

Create invite link



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
	groupId := "groupId_example" // string | Group ID
	accountId := "accountId_example" // string | WhatsApp social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.CreateWhatsAppGroupInviteLink(context.Background(), groupId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.CreateWhatsAppGroupInviteLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWhatsAppGroupInviteLink`: CreateWhatsAppGroupInviteLink200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.CreateWhatsAppGroupInviteLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWhatsAppGroupInviteLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**CreateWhatsAppGroupInviteLink200Response**](CreateWhatsAppGroupInviteLink200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWhatsAppTemplate

> CreateWhatsAppTemplate200Response CreateWhatsAppTemplate(ctx).CreateWhatsAppTemplateRequest(createWhatsAppTemplateRequest).Execute()

Create template



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
	createWhatsAppTemplateRequest := *openapiclient.NewCreateWhatsAppTemplateRequest("AccountId_example", "Name_example", "Category_example", "Language_example") // CreateWhatsAppTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.CreateWhatsAppTemplate(context.Background()).CreateWhatsAppTemplateRequest(createWhatsAppTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.CreateWhatsAppTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWhatsAppTemplate`: CreateWhatsAppTemplate200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.CreateWhatsAppTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWhatsAppTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWhatsAppTemplateRequest** | [**CreateWhatsAppTemplateRequest**](CreateWhatsAppTemplateRequest.md) |  | 

### Return type

[**CreateWhatsAppTemplate200Response**](CreateWhatsAppTemplate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWhatsAppGroupChat

> UnpublishPost200Response DeleteWhatsAppGroupChat(ctx, groupId).AccountId(accountId).Execute()

Delete group



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
	groupId := "groupId_example" // string | Group ID
	accountId := "accountId_example" // string | WhatsApp social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.DeleteWhatsAppGroupChat(context.Background(), groupId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.DeleteWhatsAppGroupChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWhatsAppGroupChat`: UnpublishPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.DeleteWhatsAppGroupChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWhatsAppGroupChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**UnpublishPost200Response**](UnpublishPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWhatsAppTemplate

> UnpublishPost200Response DeleteWhatsAppTemplate(ctx, templateName).AccountId(accountId).Execute()

Delete template



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
	templateName := "templateName_example" // string | Template name
	accountId := "accountId_example" // string | WhatsApp social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.DeleteWhatsAppTemplate(context.Background(), templateName).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.DeleteWhatsAppTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWhatsAppTemplate`: UnpublishPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.DeleteWhatsAppTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateName** | **string** | Template name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWhatsAppTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**UnpublishPost200Response**](UnpublishPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppBlockStatus

> GetWhatsAppBlockStatus200Response GetWhatsAppBlockStatus(ctx).AccountId(accountId).User(user).Execute()

Check if a user is blocked



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
	user := "user_example" // string | Consumer wa_id or E.164 phone (leading + optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.GetWhatsAppBlockStatus(context.Background()).AccountId(accountId).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.GetWhatsAppBlockStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppBlockStatus`: GetWhatsAppBlockStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.GetWhatsAppBlockStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppBlockStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **user** | **string** | Consumer wa_id or E.164 phone (leading + optional) | 

### Return type

[**GetWhatsAppBlockStatus200Response**](GetWhatsAppBlockStatus200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppBlockedUsers

> GetWhatsAppBlockedUsers200Response GetWhatsAppBlockedUsers(ctx).AccountId(accountId).Limit(limit).After(after).Execute()

List blocked users



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
	limit := int32(56) // int32 | Page size. (optional)
	after := "after_example" // string | Cursor from a previous response's `nextCursor`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.GetWhatsAppBlockedUsers(context.Background()).AccountId(accountId).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.GetWhatsAppBlockedUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppBlockedUsers`: GetWhatsAppBlockedUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.GetWhatsAppBlockedUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppBlockedUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | WhatsApp social account ID | 
 **limit** | **int32** | Page size. | 
 **after** | **string** | Cursor from a previous response&#39;s &#x60;nextCursor&#x60;. | 

### Return type

[**GetWhatsAppBlockedUsers200Response**](GetWhatsAppBlockedUsers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppBusinessProfile

> GetWhatsAppBusinessProfile200Response GetWhatsAppBusinessProfile(ctx).AccountId(accountId).Execute()

Get business profile



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
	resp, r, err := apiClient.WhatsAppAPI.GetWhatsAppBusinessProfile(context.Background()).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.GetWhatsAppBusinessProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppBusinessProfile`: GetWhatsAppBusinessProfile200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.GetWhatsAppBusinessProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppBusinessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**GetWhatsAppBusinessProfile200Response**](GetWhatsAppBusinessProfile200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppDataset

> GetWhatsAppDataset200Response GetWhatsAppDataset(ctx).AccountId(accountId).Execute()

Get CTWA conversions dataset



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
	resp, r, err := apiClient.WhatsAppAPI.GetWhatsAppDataset(context.Background()).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.GetWhatsAppDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppDataset`: GetWhatsAppDataset200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.GetWhatsAppDataset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**GetWhatsAppDataset200Response**](GetWhatsAppDataset200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppDisplayName

> GetWhatsAppDisplayName200Response GetWhatsAppDisplayName(ctx).AccountId(accountId).Execute()

Get display name status



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
	resp, r, err := apiClient.WhatsAppAPI.GetWhatsAppDisplayName(context.Background()).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.GetWhatsAppDisplayName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppDisplayName`: GetWhatsAppDisplayName200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.GetWhatsAppDisplayName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppDisplayNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**GetWhatsAppDisplayName200Response**](GetWhatsAppDisplayName200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppGroupChat

> GetWhatsAppGroupChat200Response GetWhatsAppGroupChat(ctx, groupId).AccountId(accountId).Execute()

Get group info



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
	groupId := "groupId_example" // string | Group ID
	accountId := "accountId_example" // string | WhatsApp social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.GetWhatsAppGroupChat(context.Background(), groupId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.GetWhatsAppGroupChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppGroupChat`: GetWhatsAppGroupChat200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.GetWhatsAppGroupChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppGroupChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**GetWhatsAppGroupChat200Response**](GetWhatsAppGroupChat200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppTemplate

> GetWhatsAppTemplate200Response GetWhatsAppTemplate(ctx, templateName).AccountId(accountId).Execute()

Get template



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
	templateName := "templateName_example" // string | Template name
	accountId := "accountId_example" // string | WhatsApp social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.GetWhatsAppTemplate(context.Background(), templateName).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.GetWhatsAppTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppTemplate`: GetWhatsAppTemplate200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.GetWhatsAppTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateName** | **string** | Template name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**GetWhatsAppTemplate200Response**](GetWhatsAppTemplate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhatsAppTemplates

> GetWhatsAppTemplates200Response GetWhatsAppTemplates(ctx).AccountId(accountId).Execute()

List templates



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
	resp, r, err := apiClient.WhatsAppAPI.GetWhatsAppTemplates(context.Background()).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.GetWhatsAppTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhatsAppTemplates`: GetWhatsAppTemplates200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.GetWhatsAppTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhatsAppTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**GetWhatsAppTemplates200Response**](GetWhatsAppTemplates200Response.md)

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
	resp, r, err := apiClient.WhatsAppAPI.ListWhatsAppConversions(context.Background()).AccountId(accountId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.ListWhatsAppConversions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWhatsAppConversions`: ListWhatsAppConversions200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.ListWhatsAppConversions`: %v\n", resp)
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


## ListWhatsAppGroupChats

> ListWhatsAppGroupChats200Response ListWhatsAppGroupChats(ctx).AccountId(accountId).Limit(limit).After(after).Execute()

List active groups



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
	limit := int32(56) // int32 | Max groups to return (optional) (default to 25)
	after := "after_example" // string | Pagination cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.ListWhatsAppGroupChats(context.Background()).AccountId(accountId).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.ListWhatsAppGroupChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWhatsAppGroupChats`: ListWhatsAppGroupChats200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.ListWhatsAppGroupChats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWhatsAppGroupChatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | WhatsApp social account ID | 
 **limit** | **int32** | Max groups to return | [default to 25]
 **after** | **string** | Pagination cursor | 

### Return type

[**ListWhatsAppGroupChats200Response**](ListWhatsAppGroupChats200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWhatsAppGroupJoinRequests

> ListWhatsAppGroupJoinRequests200Response ListWhatsAppGroupJoinRequests(ctx, groupId).AccountId(accountId).Execute()

List join requests



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
	groupId := "groupId_example" // string | Group ID
	accountId := "accountId_example" // string | WhatsApp social account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.ListWhatsAppGroupJoinRequests(context.Background(), groupId).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.ListWhatsAppGroupJoinRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWhatsAppGroupJoinRequests`: ListWhatsAppGroupJoinRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.ListWhatsAppGroupJoinRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWhatsAppGroupJoinRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 

### Return type

[**ListWhatsAppGroupJoinRequests200Response**](ListWhatsAppGroupJoinRequests200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectWhatsAppGroupJoinRequests

> UnpublishPost200Response RejectWhatsAppGroupJoinRequests(ctx, groupId).AccountId(accountId).RemoveWhatsAppGroupParticipantsRequest(removeWhatsAppGroupParticipantsRequest).Execute()

Reject join requests



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
	groupId := "groupId_example" // string | Group ID
	accountId := "accountId_example" // string | WhatsApp social account ID
	removeWhatsAppGroupParticipantsRequest := *openapiclient.NewRemoveWhatsAppGroupParticipantsRequest([]string{"PhoneNumbers_example"}) // RemoveWhatsAppGroupParticipantsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.RejectWhatsAppGroupJoinRequests(context.Background(), groupId).AccountId(accountId).RemoveWhatsAppGroupParticipantsRequest(removeWhatsAppGroupParticipantsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.RejectWhatsAppGroupJoinRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectWhatsAppGroupJoinRequests`: UnpublishPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.RejectWhatsAppGroupJoinRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectWhatsAppGroupJoinRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 
 **removeWhatsAppGroupParticipantsRequest** | [**RemoveWhatsAppGroupParticipantsRequest**](RemoveWhatsAppGroupParticipantsRequest.md) |  | 

### Return type

[**UnpublishPost200Response**](UnpublishPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveWhatsAppGroupParticipants

> UnpublishPost200Response RemoveWhatsAppGroupParticipants(ctx, groupId).AccountId(accountId).RemoveWhatsAppGroupParticipantsRequest(removeWhatsAppGroupParticipantsRequest).Execute()

Remove participants



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
	groupId := "groupId_example" // string | Group ID
	accountId := "accountId_example" // string | WhatsApp social account ID
	removeWhatsAppGroupParticipantsRequest := *openapiclient.NewRemoveWhatsAppGroupParticipantsRequest([]string{"PhoneNumbers_example"}) // RemoveWhatsAppGroupParticipantsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.RemoveWhatsAppGroupParticipants(context.Background(), groupId).AccountId(accountId).RemoveWhatsAppGroupParticipantsRequest(removeWhatsAppGroupParticipantsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.RemoveWhatsAppGroupParticipants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveWhatsAppGroupParticipants`: UnpublishPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.RemoveWhatsAppGroupParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWhatsAppGroupParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 
 **removeWhatsAppGroupParticipantsRequest** | [**RemoveWhatsAppGroupParticipantsRequest**](RemoveWhatsAppGroupParticipantsRequest.md) |  | 

### Return type

[**UnpublishPost200Response**](UnpublishPost200Response.md)

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
	resp, r, err := apiClient.WhatsAppAPI.SendWhatsAppConversion(context.Background()).SendWhatsAppConversionRequest(sendWhatsAppConversionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.SendWhatsAppConversion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendWhatsAppConversion`: SendWhatsAppConversion200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.SendWhatsAppConversion`: %v\n", resp)
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


## UnblockWhatsAppUsers

> UnblockWhatsAppUsers200Response UnblockWhatsAppUsers(ctx).BlockWhatsAppUsersRequest(blockWhatsAppUsersRequest).Execute()

Unblock users



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
	blockWhatsAppUsersRequest := *openapiclient.NewBlockWhatsAppUsersRequest("AccountId_example", []string{"Users_example"}) // BlockWhatsAppUsersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.UnblockWhatsAppUsers(context.Background()).BlockWhatsAppUsersRequest(blockWhatsAppUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.UnblockWhatsAppUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnblockWhatsAppUsers`: UnblockWhatsAppUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.UnblockWhatsAppUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnblockWhatsAppUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockWhatsAppUsersRequest** | [**BlockWhatsAppUsersRequest**](BlockWhatsAppUsersRequest.md) |  | 

### Return type

[**UnblockWhatsAppUsers200Response**](UnblockWhatsAppUsers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWhatsAppBusinessProfile

> UnpublishPost200Response UpdateWhatsAppBusinessProfile(ctx).UpdateWhatsAppBusinessProfileRequest(updateWhatsAppBusinessProfileRequest).Execute()

Update business profile



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
	updateWhatsAppBusinessProfileRequest := *openapiclient.NewUpdateWhatsAppBusinessProfileRequest("AccountId_example") // UpdateWhatsAppBusinessProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.UpdateWhatsAppBusinessProfile(context.Background()).UpdateWhatsAppBusinessProfileRequest(updateWhatsAppBusinessProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.UpdateWhatsAppBusinessProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWhatsAppBusinessProfile`: UnpublishPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.UpdateWhatsAppBusinessProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWhatsAppBusinessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateWhatsAppBusinessProfileRequest** | [**UpdateWhatsAppBusinessProfileRequest**](UpdateWhatsAppBusinessProfileRequest.md) |  | 

### Return type

[**UnpublishPost200Response**](UnpublishPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWhatsAppDisplayName

> UpdateWhatsAppDisplayName200Response UpdateWhatsAppDisplayName(ctx).UpdateWhatsAppDisplayNameRequest(updateWhatsAppDisplayNameRequest).Execute()

Request display name change



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
	updateWhatsAppDisplayNameRequest := *openapiclient.NewUpdateWhatsAppDisplayNameRequest("AccountId_example", "DisplayName_example") // UpdateWhatsAppDisplayNameRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.UpdateWhatsAppDisplayName(context.Background()).UpdateWhatsAppDisplayNameRequest(updateWhatsAppDisplayNameRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.UpdateWhatsAppDisplayName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWhatsAppDisplayName`: UpdateWhatsAppDisplayName200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.UpdateWhatsAppDisplayName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWhatsAppDisplayNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateWhatsAppDisplayNameRequest** | [**UpdateWhatsAppDisplayNameRequest**](UpdateWhatsAppDisplayNameRequest.md) |  | 

### Return type

[**UpdateWhatsAppDisplayName200Response**](UpdateWhatsAppDisplayName200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWhatsAppGroupChat

> UnpublishPost200Response UpdateWhatsAppGroupChat(ctx, groupId).AccountId(accountId).UpdateWhatsAppGroupChatRequest(updateWhatsAppGroupChatRequest).Execute()

Update group settings



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
	groupId := "groupId_example" // string | Group ID
	accountId := "accountId_example" // string | WhatsApp social account ID
	updateWhatsAppGroupChatRequest := *openapiclient.NewUpdateWhatsAppGroupChatRequest() // UpdateWhatsAppGroupChatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.UpdateWhatsAppGroupChat(context.Background(), groupId).AccountId(accountId).UpdateWhatsAppGroupChatRequest(updateWhatsAppGroupChatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.UpdateWhatsAppGroupChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWhatsAppGroupChat`: UnpublishPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.UpdateWhatsAppGroupChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWhatsAppGroupChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | WhatsApp social account ID | 
 **updateWhatsAppGroupChatRequest** | [**UpdateWhatsAppGroupChatRequest**](UpdateWhatsAppGroupChatRequest.md) |  | 

### Return type

[**UnpublishPost200Response**](UnpublishPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWhatsAppTemplate

> UpdateWhatsAppTemplate200Response UpdateWhatsAppTemplate(ctx, templateName).UpdateWhatsAppTemplateRequest(updateWhatsAppTemplateRequest).Execute()

Update template



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
	templateName := "templateName_example" // string | Template name
	updateWhatsAppTemplateRequest := *openapiclient.NewUpdateWhatsAppTemplateRequest("AccountId_example", []openapiclient.WhatsAppTemplateComponent{openapiclient.WhatsAppTemplateComponent{WhatsAppBodyComponent: openapiclient.NewWhatsAppBodyComponent("Type_example", "Text_example")}}) // UpdateWhatsAppTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.UpdateWhatsAppTemplate(context.Background(), templateName).UpdateWhatsAppTemplateRequest(updateWhatsAppTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.UpdateWhatsAppTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWhatsAppTemplate`: UpdateWhatsAppTemplate200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.UpdateWhatsAppTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateName** | **string** | Template name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWhatsAppTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWhatsAppTemplateRequest** | [**UpdateWhatsAppTemplateRequest**](UpdateWhatsAppTemplateRequest.md) |  | 

### Return type

[**UpdateWhatsAppTemplate200Response**](UpdateWhatsAppTemplate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadWhatsAppProfilePhoto

> UnpublishPost200Response UploadWhatsAppProfilePhoto(ctx).AccountId(accountId).File(file).Execute()

Upload profile picture



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
	file := os.NewFile(1234, "some_file") // *os.File | Image file (JPEG or PNG, max 5MB, recommended 640x640)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhatsAppAPI.UploadWhatsAppProfilePhoto(context.Background()).AccountId(accountId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhatsAppAPI.UploadWhatsAppProfilePhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadWhatsAppProfilePhoto`: UnpublishPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WhatsAppAPI.UploadWhatsAppProfilePhoto`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadWhatsAppProfilePhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | WhatsApp social account ID | 
 **file** | ***os.File** | Image file (JPEG or PNG, max 5MB, recommended 640x640) | 

### Return type

[**UnpublishPost200Response**](UnpublishPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

