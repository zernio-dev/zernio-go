# SendInboxMessageRequestTemplateElementsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Element title (max 80 chars). Required for Instagram/Facebook generic templates. | [optional] 
**Subtitle** | Pointer to **string** | Element subtitle (Instagram/Facebook only) | [optional] 
**ImageUrl** | Pointer to **string** | Element image URL (Instagram/Facebook only) | [optional] 
**Buttons** | Pointer to [**[]SendInboxMessageRequestTemplateElementsInnerButtonsInner**](SendInboxMessageRequestTemplateElementsInnerButtonsInner.md) | Element buttons (Instagram/Facebook only) | [optional] 
**Name** | Pointer to **string** | WhatsApp only. Name of the approved template to send. | [optional] 
**Language** | Pointer to **string** | WhatsApp only. Template language code (e.g. en_US). | [optional] 
**Components** | Pointer to **[]map[string]interface{}** | WhatsApp only. Meta Cloud API send-shape components array, forwarded to Meta verbatim. | [optional] 

## Methods

### NewSendInboxMessageRequestTemplateElementsInner

`func NewSendInboxMessageRequestTemplateElementsInner() *SendInboxMessageRequestTemplateElementsInner`

NewSendInboxMessageRequestTemplateElementsInner instantiates a new SendInboxMessageRequestTemplateElementsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestTemplateElementsInnerWithDefaults

`func NewSendInboxMessageRequestTemplateElementsInnerWithDefaults() *SendInboxMessageRequestTemplateElementsInner`

NewSendInboxMessageRequestTemplateElementsInnerWithDefaults instantiates a new SendInboxMessageRequestTemplateElementsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SendInboxMessageRequestTemplateElementsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendInboxMessageRequestTemplateElementsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendInboxMessageRequestTemplateElementsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SendInboxMessageRequestTemplateElementsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSubtitle

`func (o *SendInboxMessageRequestTemplateElementsInner) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *SendInboxMessageRequestTemplateElementsInner) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *SendInboxMessageRequestTemplateElementsInner) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *SendInboxMessageRequestTemplateElementsInner) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetImageUrl

`func (o *SendInboxMessageRequestTemplateElementsInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *SendInboxMessageRequestTemplateElementsInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *SendInboxMessageRequestTemplateElementsInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *SendInboxMessageRequestTemplateElementsInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetButtons

`func (o *SendInboxMessageRequestTemplateElementsInner) GetButtons() []SendInboxMessageRequestTemplateElementsInnerButtonsInner`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *SendInboxMessageRequestTemplateElementsInner) GetButtonsOk() (*[]SendInboxMessageRequestTemplateElementsInnerButtonsInner, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *SendInboxMessageRequestTemplateElementsInner) SetButtons(v []SendInboxMessageRequestTemplateElementsInnerButtonsInner)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *SendInboxMessageRequestTemplateElementsInner) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetName

`func (o *SendInboxMessageRequestTemplateElementsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SendInboxMessageRequestTemplateElementsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SendInboxMessageRequestTemplateElementsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SendInboxMessageRequestTemplateElementsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLanguage

`func (o *SendInboxMessageRequestTemplateElementsInner) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SendInboxMessageRequestTemplateElementsInner) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SendInboxMessageRequestTemplateElementsInner) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SendInboxMessageRequestTemplateElementsInner) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetComponents

`func (o *SendInboxMessageRequestTemplateElementsInner) GetComponents() []map[string]interface{}`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SendInboxMessageRequestTemplateElementsInner) GetComponentsOk() (*[]map[string]interface{}, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SendInboxMessageRequestTemplateElementsInner) SetComponents(v []map[string]interface{})`

SetComponents sets Components field to given value.

### HasComponents

`func (o *SendInboxMessageRequestTemplateElementsInner) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


