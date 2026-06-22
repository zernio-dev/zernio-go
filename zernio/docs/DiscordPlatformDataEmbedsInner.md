# DiscordPlatformDataEmbedsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Embed title (max 256 chars) | [optional] 
**Description** | Pointer to **string** | Embed body text (max 4,096 chars) | [optional] 
**Url** | Pointer to **string** | URL the title links to | [optional] 
**Color** | Pointer to **int32** | Embed accent color as decimal integer (e.g. 5814783 for blue). Convert hex to decimal. | [optional] 
**Image** | Pointer to [**DiscordPlatformDataEmbedsInnerImage**](DiscordPlatformDataEmbedsInnerImage.md) |  | [optional] 
**Thumbnail** | Pointer to [**DiscordPlatformDataEmbedsInnerImage**](DiscordPlatformDataEmbedsInnerImage.md) |  | [optional] 
**Footer** | Pointer to [**DiscordPlatformDataEmbedsInnerFooter**](DiscordPlatformDataEmbedsInnerFooter.md) |  | [optional] 
**Author** | Pointer to [**DiscordPlatformDataEmbedsInnerAuthor**](DiscordPlatformDataEmbedsInnerAuthor.md) |  | [optional] 
**Fields** | Pointer to [**[]DiscordPlatformDataEmbedsInnerFieldsInner**](DiscordPlatformDataEmbedsInnerFieldsInner.md) | Up to 25 fields per embed | [optional] 

## Methods

### NewDiscordPlatformDataEmbedsInner

`func NewDiscordPlatformDataEmbedsInner() *DiscordPlatformDataEmbedsInner`

NewDiscordPlatformDataEmbedsInner instantiates a new DiscordPlatformDataEmbedsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscordPlatformDataEmbedsInnerWithDefaults

`func NewDiscordPlatformDataEmbedsInnerWithDefaults() *DiscordPlatformDataEmbedsInner`

NewDiscordPlatformDataEmbedsInnerWithDefaults instantiates a new DiscordPlatformDataEmbedsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DiscordPlatformDataEmbedsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DiscordPlatformDataEmbedsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DiscordPlatformDataEmbedsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DiscordPlatformDataEmbedsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *DiscordPlatformDataEmbedsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscordPlatformDataEmbedsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscordPlatformDataEmbedsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscordPlatformDataEmbedsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *DiscordPlatformDataEmbedsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DiscordPlatformDataEmbedsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DiscordPlatformDataEmbedsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DiscordPlatformDataEmbedsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetColor

`func (o *DiscordPlatformDataEmbedsInner) GetColor() int32`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DiscordPlatformDataEmbedsInner) GetColorOk() (*int32, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DiscordPlatformDataEmbedsInner) SetColor(v int32)`

SetColor sets Color field to given value.

### HasColor

`func (o *DiscordPlatformDataEmbedsInner) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetImage

`func (o *DiscordPlatformDataEmbedsInner) GetImage() DiscordPlatformDataEmbedsInnerImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DiscordPlatformDataEmbedsInner) GetImageOk() (*DiscordPlatformDataEmbedsInnerImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DiscordPlatformDataEmbedsInner) SetImage(v DiscordPlatformDataEmbedsInnerImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *DiscordPlatformDataEmbedsInner) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetThumbnail

`func (o *DiscordPlatformDataEmbedsInner) GetThumbnail() DiscordPlatformDataEmbedsInnerImage`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *DiscordPlatformDataEmbedsInner) GetThumbnailOk() (*DiscordPlatformDataEmbedsInnerImage, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *DiscordPlatformDataEmbedsInner) SetThumbnail(v DiscordPlatformDataEmbedsInnerImage)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *DiscordPlatformDataEmbedsInner) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetFooter

`func (o *DiscordPlatformDataEmbedsInner) GetFooter() DiscordPlatformDataEmbedsInnerFooter`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *DiscordPlatformDataEmbedsInner) GetFooterOk() (*DiscordPlatformDataEmbedsInnerFooter, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *DiscordPlatformDataEmbedsInner) SetFooter(v DiscordPlatformDataEmbedsInnerFooter)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *DiscordPlatformDataEmbedsInner) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetAuthor

`func (o *DiscordPlatformDataEmbedsInner) GetAuthor() DiscordPlatformDataEmbedsInnerAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *DiscordPlatformDataEmbedsInner) GetAuthorOk() (*DiscordPlatformDataEmbedsInnerAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *DiscordPlatformDataEmbedsInner) SetAuthor(v DiscordPlatformDataEmbedsInnerAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *DiscordPlatformDataEmbedsInner) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetFields

`func (o *DiscordPlatformDataEmbedsInner) GetFields() []DiscordPlatformDataEmbedsInnerFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DiscordPlatformDataEmbedsInner) GetFieldsOk() (*[]DiscordPlatformDataEmbedsInnerFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DiscordPlatformDataEmbedsInner) SetFields(v []DiscordPlatformDataEmbedsInnerFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DiscordPlatformDataEmbedsInner) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


