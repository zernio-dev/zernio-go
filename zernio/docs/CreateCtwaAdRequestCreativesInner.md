# CreateCtwaAdRequestCreativesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headline** | **string** |  | 
**Body** | **string** | Primary text shown above the image / video. | 
**ImageUrl** | Pointer to **string** | Image asset. Mutually exclusive with this entry&#39;s &#x60;video&#x60;. Required if &#x60;video&#x60; is not supplied.  | [optional] 
**Video** | Pointer to [**CreateCtwaAdRequestVideo**](CreateCtwaAdRequestVideo.md) |  | [optional] 

## Methods

### NewCreateCtwaAdRequestCreativesInner

`func NewCreateCtwaAdRequestCreativesInner(headline string, body string, ) *CreateCtwaAdRequestCreativesInner`

NewCreateCtwaAdRequestCreativesInner instantiates a new CreateCtwaAdRequestCreativesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCtwaAdRequestCreativesInnerWithDefaults

`func NewCreateCtwaAdRequestCreativesInnerWithDefaults() *CreateCtwaAdRequestCreativesInner`

NewCreateCtwaAdRequestCreativesInnerWithDefaults instantiates a new CreateCtwaAdRequestCreativesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeadline

`func (o *CreateCtwaAdRequestCreativesInner) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *CreateCtwaAdRequestCreativesInner) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *CreateCtwaAdRequestCreativesInner) SetHeadline(v string)`

SetHeadline sets Headline field to given value.


### GetBody

`func (o *CreateCtwaAdRequestCreativesInner) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateCtwaAdRequestCreativesInner) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateCtwaAdRequestCreativesInner) SetBody(v string)`

SetBody sets Body field to given value.


### GetImageUrl

`func (o *CreateCtwaAdRequestCreativesInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateCtwaAdRequestCreativesInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateCtwaAdRequestCreativesInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateCtwaAdRequestCreativesInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetVideo

`func (o *CreateCtwaAdRequestCreativesInner) GetVideo() CreateCtwaAdRequestVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *CreateCtwaAdRequestCreativesInner) GetVideoOk() (*CreateCtwaAdRequestVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *CreateCtwaAdRequestCreativesInner) SetVideo(v CreateCtwaAdRequestVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *CreateCtwaAdRequestCreativesInner) HasVideo() bool`

HasVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


