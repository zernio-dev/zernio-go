# ListWorkflowVersions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to [**[]ListWorkflowVersions200ResponseVersionsInner**](ListWorkflowVersions200ResponseVersionsInner.md) | Versions in reverse chronological order (newest first). | [optional] 

## Methods

### NewListWorkflowVersions200Response

`func NewListWorkflowVersions200Response() *ListWorkflowVersions200Response`

NewListWorkflowVersions200Response instantiates a new ListWorkflowVersions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflowVersions200ResponseWithDefaults

`func NewListWorkflowVersions200ResponseWithDefaults() *ListWorkflowVersions200Response`

NewListWorkflowVersions200ResponseWithDefaults instantiates a new ListWorkflowVersions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListWorkflowVersions200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListWorkflowVersions200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListWorkflowVersions200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListWorkflowVersions200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetVersions

`func (o *ListWorkflowVersions200Response) GetVersions() []ListWorkflowVersions200ResponseVersionsInner`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ListWorkflowVersions200Response) GetVersionsOk() (*[]ListWorkflowVersions200ResponseVersionsInner, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ListWorkflowVersions200Response) SetVersions(v []ListWorkflowVersions200ResponseVersionsInner)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ListWorkflowVersions200Response) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


