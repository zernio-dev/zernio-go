# TrackingTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Platform-native tag id. Meta: numeric pixel id, as a string. | 
**Name** | **string** |  | 
**Platform** | **string** |  | 
**Kind** | **string** | Platform-native flavor of the tag (Meta: &#x60;pixel&#x60;). | 
**Status** | **string** | &#x60;inactive&#x60; when the platform reports the tag as broken/unavailable. | 
**Code** | Pointer to **string** | The base-code &#x60;&lt;script&gt;&#x60; snippet to install on the site. Meta only; populated by &#x60;getTrackingTag&#x60;, omitted from the list view.  | [optional] 
**LastFiredTime** | Pointer to **NullableInt32** | Unix seconds of the last event the tag received, or &#x60;null&#x60; if it never fired. The practical \&quot;is it installed and working\&quot; signal.  | [optional] 
**IsUnavailable** | Pointer to **bool** | Whether the tag is in a broken/unavailable state (Meta &#x60;is_unavailable&#x60;). | [optional] 
**Installed** | Pointer to **bool** | Convenience flag derived from &#x60;lastFiredTime&#x60; — has the tag ever fired. | [optional] 
**CreationTime** | Pointer to **int32** | Unix seconds the tag was created. | [optional] 
**OwnerBusinessId** | Pointer to **NullableString** | Business Manager id that owns the tag, or &#x60;null&#x60; when the tag lives on a personal (non-BM) ad account — such tags can&#39;t be shared with other ad accounts.  | [optional] 
**OwnerAdAccountId** | Pointer to **string** | Ad account id (&#x60;act_...&#x60;) that owns the tag, when reported. | [optional] 

## Methods

### NewTrackingTag

`func NewTrackingTag(id string, name string, platform string, kind string, status string, ) *TrackingTag`

NewTrackingTag instantiates a new TrackingTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingTagWithDefaults

`func NewTrackingTagWithDefaults() *TrackingTag`

NewTrackingTagWithDefaults instantiates a new TrackingTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrackingTag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackingTag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackingTag) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TrackingTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrackingTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrackingTag) SetName(v string)`

SetName sets Name field to given value.


### GetPlatform

`func (o *TrackingTag) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *TrackingTag) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *TrackingTag) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetKind

`func (o *TrackingTag) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TrackingTag) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TrackingTag) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStatus

`func (o *TrackingTag) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrackingTag) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrackingTag) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCode

`func (o *TrackingTag) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TrackingTag) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TrackingTag) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TrackingTag) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLastFiredTime

`func (o *TrackingTag) GetLastFiredTime() int32`

GetLastFiredTime returns the LastFiredTime field if non-nil, zero value otherwise.

### GetLastFiredTimeOk

`func (o *TrackingTag) GetLastFiredTimeOk() (*int32, bool)`

GetLastFiredTimeOk returns a tuple with the LastFiredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFiredTime

`func (o *TrackingTag) SetLastFiredTime(v int32)`

SetLastFiredTime sets LastFiredTime field to given value.

### HasLastFiredTime

`func (o *TrackingTag) HasLastFiredTime() bool`

HasLastFiredTime returns a boolean if a field has been set.

### SetLastFiredTimeNil

`func (o *TrackingTag) SetLastFiredTimeNil(b bool)`

 SetLastFiredTimeNil sets the value for LastFiredTime to be an explicit nil

### UnsetLastFiredTime
`func (o *TrackingTag) UnsetLastFiredTime()`

UnsetLastFiredTime ensures that no value is present for LastFiredTime, not even an explicit nil
### GetIsUnavailable

`func (o *TrackingTag) GetIsUnavailable() bool`

GetIsUnavailable returns the IsUnavailable field if non-nil, zero value otherwise.

### GetIsUnavailableOk

`func (o *TrackingTag) GetIsUnavailableOk() (*bool, bool)`

GetIsUnavailableOk returns a tuple with the IsUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnavailable

`func (o *TrackingTag) SetIsUnavailable(v bool)`

SetIsUnavailable sets IsUnavailable field to given value.

### HasIsUnavailable

`func (o *TrackingTag) HasIsUnavailable() bool`

HasIsUnavailable returns a boolean if a field has been set.

### GetInstalled

`func (o *TrackingTag) GetInstalled() bool`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *TrackingTag) GetInstalledOk() (*bool, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *TrackingTag) SetInstalled(v bool)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *TrackingTag) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetCreationTime

`func (o *TrackingTag) GetCreationTime() int32`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *TrackingTag) GetCreationTimeOk() (*int32, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *TrackingTag) SetCreationTime(v int32)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *TrackingTag) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetOwnerBusinessId

`func (o *TrackingTag) GetOwnerBusinessId() string`

GetOwnerBusinessId returns the OwnerBusinessId field if non-nil, zero value otherwise.

### GetOwnerBusinessIdOk

`func (o *TrackingTag) GetOwnerBusinessIdOk() (*string, bool)`

GetOwnerBusinessIdOk returns a tuple with the OwnerBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerBusinessId

`func (o *TrackingTag) SetOwnerBusinessId(v string)`

SetOwnerBusinessId sets OwnerBusinessId field to given value.

### HasOwnerBusinessId

`func (o *TrackingTag) HasOwnerBusinessId() bool`

HasOwnerBusinessId returns a boolean if a field has been set.

### SetOwnerBusinessIdNil

`func (o *TrackingTag) SetOwnerBusinessIdNil(b bool)`

 SetOwnerBusinessIdNil sets the value for OwnerBusinessId to be an explicit nil

### UnsetOwnerBusinessId
`func (o *TrackingTag) UnsetOwnerBusinessId()`

UnsetOwnerBusinessId ensures that no value is present for OwnerBusinessId, not even an explicit nil
### GetOwnerAdAccountId

`func (o *TrackingTag) GetOwnerAdAccountId() string`

GetOwnerAdAccountId returns the OwnerAdAccountId field if non-nil, zero value otherwise.

### GetOwnerAdAccountIdOk

`func (o *TrackingTag) GetOwnerAdAccountIdOk() (*string, bool)`

GetOwnerAdAccountIdOk returns a tuple with the OwnerAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAdAccountId

`func (o *TrackingTag) SetOwnerAdAccountId(v string)`

SetOwnerAdAccountId sets OwnerAdAccountId field to given value.

### HasOwnerAdAccountId

`func (o *TrackingTag) HasOwnerAdAccountId() bool`

HasOwnerAdAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


