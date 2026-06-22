# UpdateTrackingTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**EnableAutomaticMatching** | Pointer to **bool** | Meta Advanced Matching toggle (&#x60;enable_automatic_matching&#x60;). | [optional] 
**AutomaticMatchingFields** | Pointer to **[]string** | Which user fields Advanced Matching may collect. Meta&#39;s terse codes: em&#x3D;email, ph&#x3D;phone, fn&#x3D;first name, ln&#x3D;last name, ge&#x3D;gender, db&#x3D;date of birth, ct&#x3D;city, st&#x3D;state, zp&#x3D;zip.  | [optional] 
**FirstPartyCookieStatus** | Pointer to **string** |  | [optional] 
**DataUseSetting** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateTrackingTagRequest

`func NewUpdateTrackingTagRequest() *UpdateTrackingTagRequest`

NewUpdateTrackingTagRequest instantiates a new UpdateTrackingTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTrackingTagRequestWithDefaults

`func NewUpdateTrackingTagRequestWithDefaults() *UpdateTrackingTagRequest`

NewUpdateTrackingTagRequestWithDefaults instantiates a new UpdateTrackingTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTrackingTagRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTrackingTagRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTrackingTagRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTrackingTagRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnableAutomaticMatching

`func (o *UpdateTrackingTagRequest) GetEnableAutomaticMatching() bool`

GetEnableAutomaticMatching returns the EnableAutomaticMatching field if non-nil, zero value otherwise.

### GetEnableAutomaticMatchingOk

`func (o *UpdateTrackingTagRequest) GetEnableAutomaticMatchingOk() (*bool, bool)`

GetEnableAutomaticMatchingOk returns a tuple with the EnableAutomaticMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticMatching

`func (o *UpdateTrackingTagRequest) SetEnableAutomaticMatching(v bool)`

SetEnableAutomaticMatching sets EnableAutomaticMatching field to given value.

### HasEnableAutomaticMatching

`func (o *UpdateTrackingTagRequest) HasEnableAutomaticMatching() bool`

HasEnableAutomaticMatching returns a boolean if a field has been set.

### GetAutomaticMatchingFields

`func (o *UpdateTrackingTagRequest) GetAutomaticMatchingFields() []string`

GetAutomaticMatchingFields returns the AutomaticMatchingFields field if non-nil, zero value otherwise.

### GetAutomaticMatchingFieldsOk

`func (o *UpdateTrackingTagRequest) GetAutomaticMatchingFieldsOk() (*[]string, bool)`

GetAutomaticMatchingFieldsOk returns a tuple with the AutomaticMatchingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticMatchingFields

`func (o *UpdateTrackingTagRequest) SetAutomaticMatchingFields(v []string)`

SetAutomaticMatchingFields sets AutomaticMatchingFields field to given value.

### HasAutomaticMatchingFields

`func (o *UpdateTrackingTagRequest) HasAutomaticMatchingFields() bool`

HasAutomaticMatchingFields returns a boolean if a field has been set.

### GetFirstPartyCookieStatus

`func (o *UpdateTrackingTagRequest) GetFirstPartyCookieStatus() string`

GetFirstPartyCookieStatus returns the FirstPartyCookieStatus field if non-nil, zero value otherwise.

### GetFirstPartyCookieStatusOk

`func (o *UpdateTrackingTagRequest) GetFirstPartyCookieStatusOk() (*string, bool)`

GetFirstPartyCookieStatusOk returns a tuple with the FirstPartyCookieStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPartyCookieStatus

`func (o *UpdateTrackingTagRequest) SetFirstPartyCookieStatus(v string)`

SetFirstPartyCookieStatus sets FirstPartyCookieStatus field to given value.

### HasFirstPartyCookieStatus

`func (o *UpdateTrackingTagRequest) HasFirstPartyCookieStatus() bool`

HasFirstPartyCookieStatus returns a boolean if a field has been set.

### GetDataUseSetting

`func (o *UpdateTrackingTagRequest) GetDataUseSetting() string`

GetDataUseSetting returns the DataUseSetting field if non-nil, zero value otherwise.

### GetDataUseSettingOk

`func (o *UpdateTrackingTagRequest) GetDataUseSettingOk() (*string, bool)`

GetDataUseSettingOk returns a tuple with the DataUseSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUseSetting

`func (o *UpdateTrackingTagRequest) SetDataUseSetting(v string)`

SetDataUseSetting sets DataUseSetting field to given value.

### HasDataUseSetting

`func (o *UpdateTrackingTagRequest) HasDataUseSetting() bool`

HasDataUseSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


