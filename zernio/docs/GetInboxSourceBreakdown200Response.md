# GetInboxSourceBreakdown200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **NullableString** |  | [optional] 
**Sources** | Pointer to [**[]GetInboxSourceBreakdown200ResponseSourcesInner**](GetInboxSourceBreakdown200ResponseSourcesInner.md) |  | [optional] 

## Methods

### NewGetInboxSourceBreakdown200Response

`func NewGetInboxSourceBreakdown200Response() *GetInboxSourceBreakdown200Response`

NewGetInboxSourceBreakdown200Response instantiates a new GetInboxSourceBreakdown200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxSourceBreakdown200ResponseWithDefaults

`func NewGetInboxSourceBreakdown200ResponseWithDefaults() *GetInboxSourceBreakdown200Response`

NewGetInboxSourceBreakdown200ResponseWithDefaults instantiates a new GetInboxSourceBreakdown200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetInboxSourceBreakdown200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetInboxSourceBreakdown200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetInboxSourceBreakdown200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetInboxSourceBreakdown200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFrom

`func (o *GetInboxSourceBreakdown200Response) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetInboxSourceBreakdown200Response) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetInboxSourceBreakdown200Response) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetInboxSourceBreakdown200Response) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *GetInboxSourceBreakdown200Response) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetInboxSourceBreakdown200Response) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetInboxSourceBreakdown200Response) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GetInboxSourceBreakdown200Response) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *GetInboxSourceBreakdown200Response) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *GetInboxSourceBreakdown200Response) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetSources

`func (o *GetInboxSourceBreakdown200Response) GetSources() []GetInboxSourceBreakdown200ResponseSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GetInboxSourceBreakdown200Response) GetSourcesOk() (*[]GetInboxSourceBreakdown200ResponseSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GetInboxSourceBreakdown200Response) SetSources(v []GetInboxSourceBreakdown200ResponseSourcesInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *GetInboxSourceBreakdown200Response) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


