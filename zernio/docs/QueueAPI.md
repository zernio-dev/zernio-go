# \QueueAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQueueSlot**](QueueAPI.md#CreateQueueSlot) | **Post** /v1/queue/slots | Create schedule
[**DeleteQueueSlot**](QueueAPI.md#DeleteQueueSlot) | **Delete** /v1/queue/slots | Delete schedule
[**GetNextQueueSlot**](QueueAPI.md#GetNextQueueSlot) | **Get** /v1/queue/next-slot | Get next available slot
[**ListQueueSlots**](QueueAPI.md#ListQueueSlots) | **Get** /v1/queue/slots | List schedules
[**PreviewQueue**](QueueAPI.md#PreviewQueue) | **Get** /v1/queue/preview | Preview upcoming slots
[**UpdateQueueSlot**](QueueAPI.md#UpdateQueueSlot) | **Put** /v1/queue/slots | Update schedule



## CreateQueueSlot

> CreateQueueSlot201Response CreateQueueSlot(ctx).CreateQueueSlotRequest(createQueueSlotRequest).Execute()

Create schedule



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
	createQueueSlotRequest := *openapiclient.NewCreateQueueSlotRequest("ProfileId_example", "Name_example", "Timezone_example", []openapiclient.QueueSlot{*openapiclient.NewQueueSlot()}) // CreateQueueSlotRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueueAPI.CreateQueueSlot(context.Background()).CreateQueueSlotRequest(createQueueSlotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueueAPI.CreateQueueSlot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQueueSlot`: CreateQueueSlot201Response
	fmt.Fprintf(os.Stdout, "Response from `QueueAPI.CreateQueueSlot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQueueSlotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createQueueSlotRequest** | [**CreateQueueSlotRequest**](CreateQueueSlotRequest.md) |  | 

### Return type

[**CreateQueueSlot201Response**](CreateQueueSlot201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQueueSlot

> DeleteQueueSlot200Response DeleteQueueSlot(ctx).ProfileId(profileId).QueueId(queueId).Execute()

Delete schedule



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
	profileId := "profileId_example" // string | 
	queueId := "queueId_example" // string | Queue ID to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueueAPI.DeleteQueueSlot(context.Background()).ProfileId(profileId).QueueId(queueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueueAPI.DeleteQueueSlot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteQueueSlot`: DeleteQueueSlot200Response
	fmt.Fprintf(os.Stdout, "Response from `QueueAPI.DeleteQueueSlot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueueSlotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** |  | 
 **queueId** | **string** | Queue ID to delete | 

### Return type

[**DeleteQueueSlot200Response**](DeleteQueueSlot200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNextQueueSlot

> GetNextQueueSlot200Response GetNextQueueSlot(ctx).ProfileId(profileId).QueueId(queueId).Execute()

Get next available slot



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
	profileId := "profileId_example" // string | 
	queueId := "queueId_example" // string | Specific queue ID (optional, defaults to profile's default queue) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueueAPI.GetNextQueueSlot(context.Background()).ProfileId(profileId).QueueId(queueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueueAPI.GetNextQueueSlot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNextQueueSlot`: GetNextQueueSlot200Response
	fmt.Fprintf(os.Stdout, "Response from `QueueAPI.GetNextQueueSlot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNextQueueSlotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** |  | 
 **queueId** | **string** | Specific queue ID (optional, defaults to profile&#39;s default queue) | 

### Return type

[**GetNextQueueSlot200Response**](GetNextQueueSlot200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueueSlots

> ListQueueSlots200Response ListQueueSlots(ctx).ProfileId(profileId).QueueId(queueId).All(all).Execute()

List schedules



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
	profileId := "profileId_example" // string | Profile ID to get queues for
	queueId := "queueId_example" // string | Specific queue ID to retrieve (optional) (optional)
	all := "all_example" // string | Set to 'true' to list all queues for the profile (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueueAPI.ListQueueSlots(context.Background()).ProfileId(profileId).QueueId(queueId).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueueAPI.ListQueueSlots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQueueSlots`: ListQueueSlots200Response
	fmt.Fprintf(os.Stdout, "Response from `QueueAPI.ListQueueSlots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListQueueSlotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Profile ID to get queues for | 
 **queueId** | **string** | Specific queue ID to retrieve (optional) | 
 **all** | **string** | Set to &#39;true&#39; to list all queues for the profile | 

### Return type

[**ListQueueSlots200Response**](ListQueueSlots200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewQueue

> PreviewQueue200Response PreviewQueue(ctx).ProfileId(profileId).QueueId(queueId).Count(count).Execute()

Preview upcoming slots



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
	profileId := "profileId_example" // string | 
	queueId := "queueId_example" // string | Filter by specific queue ID. Omit to use the default queue. (optional)
	count := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueueAPI.PreviewQueue(context.Background()).ProfileId(profileId).QueueId(queueId).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueueAPI.PreviewQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewQueue`: PreviewQueue200Response
	fmt.Fprintf(os.Stdout, "Response from `QueueAPI.PreviewQueue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** |  | 
 **queueId** | **string** | Filter by specific queue ID. Omit to use the default queue. | 
 **count** | **int32** |  | [default to 20]

### Return type

[**PreviewQueue200Response**](PreviewQueue200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQueueSlot

> UpdateQueueSlot200Response UpdateQueueSlot(ctx).UpdateQueueSlotRequest(updateQueueSlotRequest).Execute()

Update schedule



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
	updateQueueSlotRequest := *openapiclient.NewUpdateQueueSlotRequest("ProfileId_example", "Timezone_example", []openapiclient.QueueSlot{*openapiclient.NewQueueSlot()}) // UpdateQueueSlotRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueueAPI.UpdateQueueSlot(context.Background()).UpdateQueueSlotRequest(updateQueueSlotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueueAPI.UpdateQueueSlot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateQueueSlot`: UpdateQueueSlot200Response
	fmt.Fprintf(os.Stdout, "Response from `QueueAPI.UpdateQueueSlot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQueueSlotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateQueueSlotRequest** | [**UpdateQueueSlotRequest**](UpdateQueueSlotRequest.md) |  | 

### Return type

[**UpdateQueueSlot200Response**](UpdateQueueSlot200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

