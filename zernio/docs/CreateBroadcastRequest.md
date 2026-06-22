# CreateBroadcastRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** |  | 
**AccountId** | **string** |  | 
**Platform** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Message** | Pointer to [**CreateBroadcastRequestMessage**](CreateBroadcastRequestMessage.md) |  | [optional] 
**Template** | Pointer to [**CreateBroadcastRequestTemplate**](CreateBroadcastRequestTemplate.md) |  | [optional] 
**SegmentFilters** | Pointer to [**CreateBroadcastRequestSegmentFilters**](CreateBroadcastRequestSegmentFilters.md) |  | [optional] 

## Methods

### NewCreateBroadcastRequest

`func NewCreateBroadcastRequest(profileId string, accountId string, platform string, name string, ) *CreateBroadcastRequest`

NewCreateBroadcastRequest instantiates a new CreateBroadcastRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBroadcastRequestWithDefaults

`func NewCreateBroadcastRequestWithDefaults() *CreateBroadcastRequest`

NewCreateBroadcastRequestWithDefaults instantiates a new CreateBroadcastRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *CreateBroadcastRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateBroadcastRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateBroadcastRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetAccountId

`func (o *CreateBroadcastRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateBroadcastRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateBroadcastRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetPlatform

`func (o *CreateBroadcastRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateBroadcastRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateBroadcastRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetName

`func (o *CreateBroadcastRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBroadcastRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBroadcastRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateBroadcastRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBroadcastRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBroadcastRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateBroadcastRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMessage

`func (o *CreateBroadcastRequest) GetMessage() CreateBroadcastRequestMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateBroadcastRequest) GetMessageOk() (*CreateBroadcastRequestMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateBroadcastRequest) SetMessage(v CreateBroadcastRequestMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateBroadcastRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTemplate

`func (o *CreateBroadcastRequest) GetTemplate() CreateBroadcastRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CreateBroadcastRequest) GetTemplateOk() (*CreateBroadcastRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CreateBroadcastRequest) SetTemplate(v CreateBroadcastRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CreateBroadcastRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetSegmentFilters

`func (o *CreateBroadcastRequest) GetSegmentFilters() CreateBroadcastRequestSegmentFilters`

GetSegmentFilters returns the SegmentFilters field if non-nil, zero value otherwise.

### GetSegmentFiltersOk

`func (o *CreateBroadcastRequest) GetSegmentFiltersOk() (*CreateBroadcastRequestSegmentFilters, bool)`

GetSegmentFiltersOk returns a tuple with the SegmentFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentFilters

`func (o *CreateBroadcastRequest) SetSegmentFilters(v CreateBroadcastRequestSegmentFilters)`

SetSegmentFilters sets SegmentFilters field to given value.

### HasSegmentFilters

`func (o *CreateBroadcastRequest) HasSegmentFilters() bool`

HasSegmentFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


