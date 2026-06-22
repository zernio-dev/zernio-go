# GetFacebookPages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pages** | Pointer to [**[]GetFacebookPages200ResponsePagesInner**](GetFacebookPages200ResponsePagesInner.md) |  | [optional] 
**SelectedPageId** | Pointer to **string** |  | [optional] 
**Cached** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetFacebookPages200Response

`func NewGetFacebookPages200Response() *GetFacebookPages200Response`

NewGetFacebookPages200Response instantiates a new GetFacebookPages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFacebookPages200ResponseWithDefaults

`func NewGetFacebookPages200ResponseWithDefaults() *GetFacebookPages200Response`

NewGetFacebookPages200ResponseWithDefaults instantiates a new GetFacebookPages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPages

`func (o *GetFacebookPages200Response) GetPages() []GetFacebookPages200ResponsePagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GetFacebookPages200Response) GetPagesOk() (*[]GetFacebookPages200ResponsePagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GetFacebookPages200Response) SetPages(v []GetFacebookPages200ResponsePagesInner)`

SetPages sets Pages field to given value.

### HasPages

`func (o *GetFacebookPages200Response) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetSelectedPageId

`func (o *GetFacebookPages200Response) GetSelectedPageId() string`

GetSelectedPageId returns the SelectedPageId field if non-nil, zero value otherwise.

### GetSelectedPageIdOk

`func (o *GetFacebookPages200Response) GetSelectedPageIdOk() (*string, bool)`

GetSelectedPageIdOk returns a tuple with the SelectedPageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPageId

`func (o *GetFacebookPages200Response) SetSelectedPageId(v string)`

SetSelectedPageId sets SelectedPageId field to given value.

### HasSelectedPageId

`func (o *GetFacebookPages200Response) HasSelectedPageId() bool`

HasSelectedPageId returns a boolean if a field has been set.

### GetCached

`func (o *GetFacebookPages200Response) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *GetFacebookPages200Response) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *GetFacebookPages200Response) SetCached(v bool)`

SetCached sets Cached field to given value.

### HasCached

`func (o *GetFacebookPages200Response) HasCached() bool`

HasCached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


