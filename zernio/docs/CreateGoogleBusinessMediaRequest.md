# CreateGoogleBusinessMediaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceUrl** | **string** | Publicly accessible image URL | 
**MediaFormat** | Pointer to **string** |  | [optional] [default to "PHOTO"]
**Description** | Pointer to **string** | Photo description | [optional] 
**Category** | Pointer to **string** | Where the photo appears on the listing | [optional] 

## Methods

### NewCreateGoogleBusinessMediaRequest

`func NewCreateGoogleBusinessMediaRequest(sourceUrl string, ) *CreateGoogleBusinessMediaRequest`

NewCreateGoogleBusinessMediaRequest instantiates a new CreateGoogleBusinessMediaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGoogleBusinessMediaRequestWithDefaults

`func NewCreateGoogleBusinessMediaRequestWithDefaults() *CreateGoogleBusinessMediaRequest`

NewCreateGoogleBusinessMediaRequestWithDefaults instantiates a new CreateGoogleBusinessMediaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceUrl

`func (o *CreateGoogleBusinessMediaRequest) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *CreateGoogleBusinessMediaRequest) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *CreateGoogleBusinessMediaRequest) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.


### GetMediaFormat

`func (o *CreateGoogleBusinessMediaRequest) GetMediaFormat() string`

GetMediaFormat returns the MediaFormat field if non-nil, zero value otherwise.

### GetMediaFormatOk

`func (o *CreateGoogleBusinessMediaRequest) GetMediaFormatOk() (*string, bool)`

GetMediaFormatOk returns a tuple with the MediaFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaFormat

`func (o *CreateGoogleBusinessMediaRequest) SetMediaFormat(v string)`

SetMediaFormat sets MediaFormat field to given value.

### HasMediaFormat

`func (o *CreateGoogleBusinessMediaRequest) HasMediaFormat() bool`

HasMediaFormat returns a boolean if a field has been set.

### GetDescription

`func (o *CreateGoogleBusinessMediaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGoogleBusinessMediaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGoogleBusinessMediaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGoogleBusinessMediaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *CreateGoogleBusinessMediaRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateGoogleBusinessMediaRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateGoogleBusinessMediaRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateGoogleBusinessMediaRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


