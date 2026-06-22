# \WorkflowsAPI

All URIs are relative to *https://zernio.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateWorkflow**](WorkflowsAPI.md#ActivateWorkflow) | **Post** /v1/workflows/{workflowId}/activate | Activate workflow
[**CreateWorkflow**](WorkflowsAPI.md#CreateWorkflow) | **Post** /v1/workflows | Create workflow
[**DeleteWorkflow**](WorkflowsAPI.md#DeleteWorkflow) | **Delete** /v1/workflows/{workflowId} | Delete workflow
[**DuplicateWorkflow**](WorkflowsAPI.md#DuplicateWorkflow) | **Post** /v1/workflows/{workflowId}/duplicate | Duplicate a workflow
[**GetWorkflow**](WorkflowsAPI.md#GetWorkflow) | **Get** /v1/workflows/{workflowId} | Get workflow with graph
[**GetWorkflowVersion**](WorkflowsAPI.md#GetWorkflowVersion) | **Get** /v1/workflows/{workflowId}/versions/{version} | Get a specific workflow version
[**ListWorkflowExecutionEvents**](WorkflowsAPI.md#ListWorkflowExecutionEvents) | **Get** /v1/workflows/{workflowId}/executions/{executionId}/events | Get an execution&#39;s timeline
[**ListWorkflowExecutions**](WorkflowsAPI.md#ListWorkflowExecutions) | **Get** /v1/workflows/{workflowId}/executions | List workflow runs
[**ListWorkflowVersions**](WorkflowsAPI.md#ListWorkflowVersions) | **Get** /v1/workflows/{workflowId}/versions | List a workflow&#39;s version history
[**ListWorkflows**](WorkflowsAPI.md#ListWorkflows) | **Get** /v1/workflows | List workflows
[**PauseWorkflow**](WorkflowsAPI.md#PauseWorkflow) | **Post** /v1/workflows/{workflowId}/pause | Pause workflow
[**RestoreWorkflowVersion**](WorkflowsAPI.md#RestoreWorkflowVersion) | **Post** /v1/workflows/{workflowId}/versions/{version}/restore | Restore a previous workflow version
[**TriggerWorkflow**](WorkflowsAPI.md#TriggerWorkflow) | **Post** /v1/workflows/{workflowId}/executions | Manually start a workflow run
[**UpdateWorkflow**](WorkflowsAPI.md#UpdateWorkflow) | **Patch** /v1/workflows/{workflowId} | Update workflow



## ActivateWorkflow

> ActivateWorkflow200Response ActivateWorkflow(ctx, workflowId).Execute()

Activate workflow



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
	workflowId := "workflowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ActivateWorkflow(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ActivateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateWorkflow`: ActivateWorkflow200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ActivateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivateWorkflow200Response**](ActivateWorkflow200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflow

> CreateWorkflow200Response CreateWorkflow(ctx).CreateWorkflowRequest(createWorkflowRequest).Execute()

Create workflow



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
	createWorkflowRequest := *openapiclient.NewCreateWorkflowRequest("ProfileId_example", "AccountId_example", "Name_example") // CreateWorkflowRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.CreateWorkflow(context.Background()).CreateWorkflowRequest(createWorkflowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.CreateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkflow`: CreateWorkflow200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.CreateWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkflowRequest** | [**CreateWorkflowRequest**](CreateWorkflowRequest.md) |  | 

### Return type

[**CreateWorkflow200Response**](CreateWorkflow200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflow

> DeleteWorkflow(ctx, workflowId).Execute()

Delete workflow



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
	workflowId := "workflowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsAPI.DeleteWorkflow(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.DeleteWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DuplicateWorkflow

> DuplicateWorkflow201Response DuplicateWorkflow(ctx, workflowId).Execute()

Duplicate a workflow



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
	workflowId := "workflowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.DuplicateWorkflow(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.DuplicateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DuplicateWorkflow`: DuplicateWorkflow201Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.DuplicateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuplicateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DuplicateWorkflow201Response**](DuplicateWorkflow201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflow

> GetWorkflow200Response GetWorkflow(ctx, workflowId).Execute()

Get workflow with graph



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
	workflowId := "workflowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflow(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflow`: GetWorkflow200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWorkflow200Response**](GetWorkflow200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowVersion

> GetWorkflowVersion200Response GetWorkflowVersion(ctx, workflowId, version).Execute()

Get a specific workflow version



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
	workflowId := "workflowId_example" // string | 
	version := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflowVersion(context.Background(), workflowId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflowVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowVersion`: GetWorkflowVersion200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflowVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 
**version** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetWorkflowVersion200Response**](GetWorkflowVersion200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowExecutionEvents

> ListWorkflowExecutionEvents200Response ListWorkflowExecutionEvents(ctx, workflowId, executionId).Execute()

Get an execution's timeline



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
	workflowId := "workflowId_example" // string | 
	executionId := "executionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ListWorkflowExecutionEvents(context.Background(), workflowId, executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ListWorkflowExecutionEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkflowExecutionEvents`: ListWorkflowExecutionEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ListWorkflowExecutionEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowExecutionEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListWorkflowExecutionEvents200Response**](ListWorkflowExecutionEvents200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowExecutions

> ListWorkflowExecutions200Response ListWorkflowExecutions(ctx, workflowId).Status(status).Limit(limit).Skip(skip).Execute()

List workflow runs



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
	workflowId := "workflowId_example" // string | 
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 25)
	skip := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ListWorkflowExecutions(context.Background(), workflowId).Status(status).Limit(limit).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ListWorkflowExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkflowExecutions`: ListWorkflowExecutions200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ListWorkflowExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 25]
 **skip** | **int32** |  | [default to 0]

### Return type

[**ListWorkflowExecutions200Response**](ListWorkflowExecutions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowVersions

> ListWorkflowVersions200Response ListWorkflowVersions(ctx, workflowId).Execute()

List a workflow's version history



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
	workflowId := "workflowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ListWorkflowVersions(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ListWorkflowVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkflowVersions`: ListWorkflowVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ListWorkflowVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListWorkflowVersions200Response**](ListWorkflowVersions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflows

> ListWorkflows200Response ListWorkflows(ctx).ProfileId(profileId).Status(status).Limit(limit).Skip(skip).Execute()

List workflows



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
	profileId := "profileId_example" // string | Filter by profile. Omit to list across all profiles (optional)
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	skip := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ListWorkflows(context.Background()).ProfileId(profileId).Status(status).Limit(limit).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ListWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkflows`: ListWorkflows200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ListWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Filter by profile. Omit to list across all profiles | 
 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **skip** | **int32** |  | [default to 0]

### Return type

[**ListWorkflows200Response**](ListWorkflows200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseWorkflow

> PauseWorkflow200Response PauseWorkflow(ctx, workflowId).Execute()

Pause workflow



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
	workflowId := "workflowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.PauseWorkflow(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.PauseWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseWorkflow`: PauseWorkflow200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.PauseWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PauseWorkflow200Response**](PauseWorkflow200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreWorkflowVersion

> RestoreWorkflowVersion200Response RestoreWorkflowVersion(ctx, workflowId, version).Execute()

Restore a previous workflow version



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
	workflowId := "workflowId_example" // string | 
	version := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.RestoreWorkflowVersion(context.Background(), workflowId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.RestoreWorkflowVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreWorkflowVersion`: RestoreWorkflowVersion200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.RestoreWorkflowVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 
**version** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreWorkflowVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestoreWorkflowVersion200Response**](RestoreWorkflowVersion200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerWorkflow

> TriggerWorkflow200Response TriggerWorkflow(ctx, workflowId).TriggerWorkflowRequest(triggerWorkflowRequest).Execute()

Manually start a workflow run



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
	workflowId := "workflowId_example" // string | 
	triggerWorkflowRequest := *openapiclient.NewTriggerWorkflowRequest() // TriggerWorkflowRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.TriggerWorkflow(context.Background(), workflowId).TriggerWorkflowRequest(triggerWorkflowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.TriggerWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerWorkflow`: TriggerWorkflow200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.TriggerWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerWorkflowRequest** | [**TriggerWorkflowRequest**](TriggerWorkflowRequest.md) |  | 

### Return type

[**TriggerWorkflow200Response**](TriggerWorkflow200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflow

> UpdateWorkflow200Response UpdateWorkflow(ctx, workflowId).UpdateWorkflowRequest(updateWorkflowRequest).Execute()

Update workflow



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
	workflowId := "workflowId_example" // string | 
	updateWorkflowRequest := *openapiclient.NewUpdateWorkflowRequest() // UpdateWorkflowRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.UpdateWorkflow(context.Background(), workflowId).UpdateWorkflowRequest(updateWorkflowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.UpdateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkflow`: UpdateWorkflow200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.UpdateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWorkflowRequest** | [**UpdateWorkflowRequest**](UpdateWorkflowRequest.md) |  | 

### Return type

[**UpdateWorkflow200Response**](UpdateWorkflow200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

