# GetWorkflowVersion200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to [**GetWorkflowVersion200ResponseVersion**](GetWorkflowVersion200ResponseVersion.md) |  | [optional] 

## Methods

### NewGetWorkflowVersion200Response

`func NewGetWorkflowVersion200Response() *GetWorkflowVersion200Response`

NewGetWorkflowVersion200Response instantiates a new GetWorkflowVersion200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWorkflowVersion200ResponseWithDefaults

`func NewGetWorkflowVersion200ResponseWithDefaults() *GetWorkflowVersion200Response`

NewGetWorkflowVersion200ResponseWithDefaults instantiates a new GetWorkflowVersion200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetWorkflowVersion200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetWorkflowVersion200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetWorkflowVersion200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetWorkflowVersion200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetVersion

`func (o *GetWorkflowVersion200Response) GetVersion() GetWorkflowVersion200ResponseVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetWorkflowVersion200Response) GetVersionOk() (*GetWorkflowVersion200ResponseVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetWorkflowVersion200Response) SetVersion(v GetWorkflowVersion200ResponseVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetWorkflowVersion200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


