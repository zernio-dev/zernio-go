# ListGoogleBusinessMedia200ResponseMediaItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Resource name | [optional] 
**MediaFormat** | Pointer to **string** |  | [optional] 
**SourceUrl** | Pointer to **string** |  | [optional] 
**GoogleUrl** | Pointer to **string** | Google-hosted URL | [optional] 
**ThumbnailUrl** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**LocationAssociation** | Pointer to [**ListGoogleBusinessMedia200ResponseMediaItemsInnerLocationAssociation**](ListGoogleBusinessMedia200ResponseMediaItemsInnerLocationAssociation.md) |  | [optional] 

## Methods

### NewListGoogleBusinessMedia200ResponseMediaItemsInner

`func NewListGoogleBusinessMedia200ResponseMediaItemsInner() *ListGoogleBusinessMedia200ResponseMediaItemsInner`

NewListGoogleBusinessMedia200ResponseMediaItemsInner instantiates a new ListGoogleBusinessMedia200ResponseMediaItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGoogleBusinessMedia200ResponseMediaItemsInnerWithDefaults

`func NewListGoogleBusinessMedia200ResponseMediaItemsInnerWithDefaults() *ListGoogleBusinessMedia200ResponseMediaItemsInner`

NewListGoogleBusinessMedia200ResponseMediaItemsInnerWithDefaults instantiates a new ListGoogleBusinessMedia200ResponseMediaItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMediaFormat

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetMediaFormat() string`

GetMediaFormat returns the MediaFormat field if non-nil, zero value otherwise.

### GetMediaFormatOk

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetMediaFormatOk() (*string, bool)`

GetMediaFormatOk returns a tuple with the MediaFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaFormat

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) SetMediaFormat(v string)`

SetMediaFormat sets MediaFormat field to given value.

### HasMediaFormat

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) HasMediaFormat() bool`

HasMediaFormat returns a boolean if a field has been set.

### GetSourceUrl

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetGoogleUrl

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetGoogleUrl() string`

GetGoogleUrl returns the GoogleUrl field if non-nil, zero value otherwise.

### GetGoogleUrlOk

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetGoogleUrlOk() (*string, bool)`

GetGoogleUrlOk returns a tuple with the GoogleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleUrl

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) SetGoogleUrl(v string)`

SetGoogleUrl sets GoogleUrl field to given value.

### HasGoogleUrl

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) HasGoogleUrl() bool`

HasGoogleUrl returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetDescription

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreateTime

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetLocationAssociation

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetLocationAssociation() ListGoogleBusinessMedia200ResponseMediaItemsInnerLocationAssociation`

GetLocationAssociation returns the LocationAssociation field if non-nil, zero value otherwise.

### GetLocationAssociationOk

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) GetLocationAssociationOk() (*ListGoogleBusinessMedia200ResponseMediaItemsInnerLocationAssociation, bool)`

GetLocationAssociationOk returns a tuple with the LocationAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationAssociation

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) SetLocationAssociation(v ListGoogleBusinessMedia200ResponseMediaItemsInnerLocationAssociation)`

SetLocationAssociation sets LocationAssociation field to given value.

### HasLocationAssociation

`func (o *ListGoogleBusinessMedia200ResponseMediaItemsInner) HasLocationAssociation() bool`

HasLocationAssociation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


