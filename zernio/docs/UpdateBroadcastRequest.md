# UpdateBroadcastRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Message** | Pointer to [**SendInboxMessageRequestInteractiveFooter**](SendInboxMessageRequestInteractiveFooter.md) |  | [optional] 
**Template** | Pointer to [**UpdateBroadcastRequestTemplate**](UpdateBroadcastRequestTemplate.md) |  | [optional] 
**SegmentFilters** | Pointer to **map[string]interface{}** | Recipient segment filters (tags, channels, subscription state). | [optional] 

## Methods

### NewUpdateBroadcastRequest

`func NewUpdateBroadcastRequest() *UpdateBroadcastRequest`

NewUpdateBroadcastRequest instantiates a new UpdateBroadcastRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBroadcastRequestWithDefaults

`func NewUpdateBroadcastRequestWithDefaults() *UpdateBroadcastRequest`

NewUpdateBroadcastRequestWithDefaults instantiates a new UpdateBroadcastRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateBroadcastRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBroadcastRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBroadcastRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateBroadcastRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateBroadcastRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateBroadcastRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateBroadcastRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateBroadcastRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMessage

`func (o *UpdateBroadcastRequest) GetMessage() SendInboxMessageRequestInteractiveFooter`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateBroadcastRequest) GetMessageOk() (*SendInboxMessageRequestInteractiveFooter, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateBroadcastRequest) SetMessage(v SendInboxMessageRequestInteractiveFooter)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateBroadcastRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTemplate

`func (o *UpdateBroadcastRequest) GetTemplate() UpdateBroadcastRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *UpdateBroadcastRequest) GetTemplateOk() (*UpdateBroadcastRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *UpdateBroadcastRequest) SetTemplate(v UpdateBroadcastRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *UpdateBroadcastRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetSegmentFilters

`func (o *UpdateBroadcastRequest) GetSegmentFilters() map[string]interface{}`

GetSegmentFilters returns the SegmentFilters field if non-nil, zero value otherwise.

### GetSegmentFiltersOk

`func (o *UpdateBroadcastRequest) GetSegmentFiltersOk() (*map[string]interface{}, bool)`

GetSegmentFiltersOk returns a tuple with the SegmentFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentFilters

`func (o *UpdateBroadcastRequest) SetSegmentFilters(v map[string]interface{})`

SetSegmentFilters sets SegmentFilters field to given value.

### HasSegmentFilters

`func (o *UpdateBroadcastRequest) HasSegmentFilters() bool`

HasSegmentFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


