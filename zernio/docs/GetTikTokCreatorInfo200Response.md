# GetTikTokCreatorInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creator** | Pointer to [**GetTikTokCreatorInfo200ResponseCreator**](GetTikTokCreatorInfo200ResponseCreator.md) |  | [optional] 
**PrivacyLevels** | Pointer to [**[]GetTikTokCreatorInfo200ResponsePrivacyLevelsInner**](GetTikTokCreatorInfo200ResponsePrivacyLevelsInner.md) | Available privacy level options for this creator | [optional] 
**PostingLimits** | Pointer to [**GetTikTokCreatorInfo200ResponsePostingLimits**](GetTikTokCreatorInfo200ResponsePostingLimits.md) |  | [optional] 
**CommercialContentTypes** | Pointer to [**[]GetTikTokCreatorInfo200ResponseCommercialContentTypesInner**](GetTikTokCreatorInfo200ResponseCommercialContentTypesInner.md) | Available commercial content disclosure options | [optional] 

## Methods

### NewGetTikTokCreatorInfo200Response

`func NewGetTikTokCreatorInfo200Response() *GetTikTokCreatorInfo200Response`

NewGetTikTokCreatorInfo200Response instantiates a new GetTikTokCreatorInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTikTokCreatorInfo200ResponseWithDefaults

`func NewGetTikTokCreatorInfo200ResponseWithDefaults() *GetTikTokCreatorInfo200Response`

NewGetTikTokCreatorInfo200ResponseWithDefaults instantiates a new GetTikTokCreatorInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreator

`func (o *GetTikTokCreatorInfo200Response) GetCreator() GetTikTokCreatorInfo200ResponseCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GetTikTokCreatorInfo200Response) GetCreatorOk() (*GetTikTokCreatorInfo200ResponseCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GetTikTokCreatorInfo200Response) SetCreator(v GetTikTokCreatorInfo200ResponseCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GetTikTokCreatorInfo200Response) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetPrivacyLevels

`func (o *GetTikTokCreatorInfo200Response) GetPrivacyLevels() []GetTikTokCreatorInfo200ResponsePrivacyLevelsInner`

GetPrivacyLevels returns the PrivacyLevels field if non-nil, zero value otherwise.

### GetPrivacyLevelsOk

`func (o *GetTikTokCreatorInfo200Response) GetPrivacyLevelsOk() (*[]GetTikTokCreatorInfo200ResponsePrivacyLevelsInner, bool)`

GetPrivacyLevelsOk returns a tuple with the PrivacyLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevels

`func (o *GetTikTokCreatorInfo200Response) SetPrivacyLevels(v []GetTikTokCreatorInfo200ResponsePrivacyLevelsInner)`

SetPrivacyLevels sets PrivacyLevels field to given value.

### HasPrivacyLevels

`func (o *GetTikTokCreatorInfo200Response) HasPrivacyLevels() bool`

HasPrivacyLevels returns a boolean if a field has been set.

### GetPostingLimits

`func (o *GetTikTokCreatorInfo200Response) GetPostingLimits() GetTikTokCreatorInfo200ResponsePostingLimits`

GetPostingLimits returns the PostingLimits field if non-nil, zero value otherwise.

### GetPostingLimitsOk

`func (o *GetTikTokCreatorInfo200Response) GetPostingLimitsOk() (*GetTikTokCreatorInfo200ResponsePostingLimits, bool)`

GetPostingLimitsOk returns a tuple with the PostingLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingLimits

`func (o *GetTikTokCreatorInfo200Response) SetPostingLimits(v GetTikTokCreatorInfo200ResponsePostingLimits)`

SetPostingLimits sets PostingLimits field to given value.

### HasPostingLimits

`func (o *GetTikTokCreatorInfo200Response) HasPostingLimits() bool`

HasPostingLimits returns a boolean if a field has been set.

### GetCommercialContentTypes

`func (o *GetTikTokCreatorInfo200Response) GetCommercialContentTypes() []GetTikTokCreatorInfo200ResponseCommercialContentTypesInner`

GetCommercialContentTypes returns the CommercialContentTypes field if non-nil, zero value otherwise.

### GetCommercialContentTypesOk

`func (o *GetTikTokCreatorInfo200Response) GetCommercialContentTypesOk() (*[]GetTikTokCreatorInfo200ResponseCommercialContentTypesInner, bool)`

GetCommercialContentTypesOk returns a tuple with the CommercialContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialContentTypes

`func (o *GetTikTokCreatorInfo200Response) SetCommercialContentTypes(v []GetTikTokCreatorInfo200ResponseCommercialContentTypesInner)`

SetCommercialContentTypes sets CommercialContentTypes field to given value.

### HasCommercialContentTypes

`func (o *GetTikTokCreatorInfo200Response) HasCommercialContentTypes() bool`

HasCommercialContentTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


