# InstagramPlatformDataUserTagsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Instagram username (@ symbol is optional and will be removed automatically) | 
**X** | **float32** | X coordinate position from left edge (0.0 &#x3D; left, 0.5 &#x3D; center, 1.0 &#x3D; right) | 
**Y** | **float32** | Y coordinate position from top edge (0.0 &#x3D; top, 0.5 &#x3D; center, 1.0 &#x3D; bottom) | 
**MediaIndex** | Pointer to **int32** | Zero-based index of the carousel item to tag. Defaults to 0. Tags on video items or out-of-range indices are ignored. | [optional] 

## Methods

### NewInstagramPlatformDataUserTagsInner

`func NewInstagramPlatformDataUserTagsInner(username string, x float32, y float32, ) *InstagramPlatformDataUserTagsInner`

NewInstagramPlatformDataUserTagsInner instantiates a new InstagramPlatformDataUserTagsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstagramPlatformDataUserTagsInnerWithDefaults

`func NewInstagramPlatformDataUserTagsInnerWithDefaults() *InstagramPlatformDataUserTagsInner`

NewInstagramPlatformDataUserTagsInnerWithDefaults instantiates a new InstagramPlatformDataUserTagsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *InstagramPlatformDataUserTagsInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InstagramPlatformDataUserTagsInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InstagramPlatformDataUserTagsInner) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetX

`func (o *InstagramPlatformDataUserTagsInner) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *InstagramPlatformDataUserTagsInner) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *InstagramPlatformDataUserTagsInner) SetX(v float32)`

SetX sets X field to given value.


### GetY

`func (o *InstagramPlatformDataUserTagsInner) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *InstagramPlatformDataUserTagsInner) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *InstagramPlatformDataUserTagsInner) SetY(v float32)`

SetY sets Y field to given value.


### GetMediaIndex

`func (o *InstagramPlatformDataUserTagsInner) GetMediaIndex() int32`

GetMediaIndex returns the MediaIndex field if non-nil, zero value otherwise.

### GetMediaIndexOk

`func (o *InstagramPlatformDataUserTagsInner) GetMediaIndexOk() (*int32, bool)`

GetMediaIndexOk returns a tuple with the MediaIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaIndex

`func (o *InstagramPlatformDataUserTagsInner) SetMediaIndex(v int32)`

SetMediaIndex sets MediaIndex field to given value.

### HasMediaIndex

`func (o *InstagramPlatformDataUserTagsInner) HasMediaIndex() bool`

HasMediaIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


