# ListWhatsAppFlowVersions200ResponseVersionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**ParentFlowId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Missing** | Pointer to **bool** | True when Meta no longer has this flow | [optional] 

## Methods

### NewListWhatsAppFlowVersions200ResponseVersionsInner

`func NewListWhatsAppFlowVersions200ResponseVersionsInner() *ListWhatsAppFlowVersions200ResponseVersionsInner`

NewListWhatsAppFlowVersions200ResponseVersionsInner instantiates a new ListWhatsAppFlowVersions200ResponseVersionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWhatsAppFlowVersions200ResponseVersionsInnerWithDefaults

`func NewListWhatsAppFlowVersions200ResponseVersionsInnerWithDefaults() *ListWhatsAppFlowVersions200ResponseVersionsInner`

NewListWhatsAppFlowVersions200ResponseVersionsInnerWithDefaults instantiates a new ListWhatsAppFlowVersions200ResponseVersionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetVersion

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetParentFlowId

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) GetParentFlowId() string`

GetParentFlowId returns the ParentFlowId field if non-nil, zero value otherwise.

### GetParentFlowIdOk

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) GetParentFlowIdOk() (*string, bool)`

GetParentFlowIdOk returns a tuple with the ParentFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFlowId

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) SetParentFlowId(v string)`

SetParentFlowId sets ParentFlowId field to given value.

### HasParentFlowId

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) HasParentFlowId() bool`

HasParentFlowId returns a boolean if a field has been set.

### SetParentFlowIdNil

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) SetParentFlowIdNil(b bool)`

 SetParentFlowIdNil sets the value for ParentFlowId to be an explicit nil

### UnsetParentFlowId
`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) UnsetParentFlowId()`

UnsetParentFlowId ensures that no value is present for ParentFlowId, not even an explicit nil
### GetName

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetMissing

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) GetMissing() bool`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) GetMissingOk() (*bool, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) SetMissing(v bool)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *ListWhatsAppFlowVersions200ResponseVersionsInner) HasMissing() bool`

HasMissing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


