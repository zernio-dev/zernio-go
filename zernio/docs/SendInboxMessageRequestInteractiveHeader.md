# SendInboxMessageRequestInteractiveHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** | Required when header type is text. | [optional] 
**Image** | Pointer to [**SendInboxMessageRequestInteractiveHeaderImage**](SendInboxMessageRequestInteractiveHeaderImage.md) |  | [optional] 
**Video** | Pointer to [**SendInboxMessageRequestInteractiveHeaderImage**](SendInboxMessageRequestInteractiveHeaderImage.md) |  | [optional] 
**Document** | Pointer to [**SendInboxMessageRequestInteractiveHeaderImage**](SendInboxMessageRequestInteractiveHeaderImage.md) |  | [optional] 

## Methods

### NewSendInboxMessageRequestInteractiveHeader

`func NewSendInboxMessageRequestInteractiveHeader() *SendInboxMessageRequestInteractiveHeader`

NewSendInboxMessageRequestInteractiveHeader instantiates a new SendInboxMessageRequestInteractiveHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestInteractiveHeaderWithDefaults

`func NewSendInboxMessageRequestInteractiveHeaderWithDefaults() *SendInboxMessageRequestInteractiveHeader`

NewSendInboxMessageRequestInteractiveHeaderWithDefaults instantiates a new SendInboxMessageRequestInteractiveHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SendInboxMessageRequestInteractiveHeader) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendInboxMessageRequestInteractiveHeader) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendInboxMessageRequestInteractiveHeader) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SendInboxMessageRequestInteractiveHeader) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *SendInboxMessageRequestInteractiveHeader) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SendInboxMessageRequestInteractiveHeader) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SendInboxMessageRequestInteractiveHeader) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SendInboxMessageRequestInteractiveHeader) HasText() bool`

HasText returns a boolean if a field has been set.

### GetImage

`func (o *SendInboxMessageRequestInteractiveHeader) GetImage() SendInboxMessageRequestInteractiveHeaderImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SendInboxMessageRequestInteractiveHeader) GetImageOk() (*SendInboxMessageRequestInteractiveHeaderImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SendInboxMessageRequestInteractiveHeader) SetImage(v SendInboxMessageRequestInteractiveHeaderImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *SendInboxMessageRequestInteractiveHeader) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVideo

`func (o *SendInboxMessageRequestInteractiveHeader) GetVideo() SendInboxMessageRequestInteractiveHeaderImage`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *SendInboxMessageRequestInteractiveHeader) GetVideoOk() (*SendInboxMessageRequestInteractiveHeaderImage, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *SendInboxMessageRequestInteractiveHeader) SetVideo(v SendInboxMessageRequestInteractiveHeaderImage)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *SendInboxMessageRequestInteractiveHeader) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetDocument

`func (o *SendInboxMessageRequestInteractiveHeader) GetDocument() SendInboxMessageRequestInteractiveHeaderImage`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *SendInboxMessageRequestInteractiveHeader) GetDocumentOk() (*SendInboxMessageRequestInteractiveHeaderImage, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *SendInboxMessageRequestInteractiveHeader) SetDocument(v SendInboxMessageRequestInteractiveHeaderImage)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *SendInboxMessageRequestInteractiveHeader) HasDocument() bool`

HasDocument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


