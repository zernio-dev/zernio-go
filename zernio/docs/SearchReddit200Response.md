# SearchReddit200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]RedditPost**](RedditPost.md) |  | [optional] 
**After** | Pointer to **NullableString** |  | [optional] 
**Before** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSearchReddit200Response

`func NewSearchReddit200Response() *SearchReddit200Response`

NewSearchReddit200Response instantiates a new SearchReddit200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchReddit200ResponseWithDefaults

`func NewSearchReddit200ResponseWithDefaults() *SearchReddit200Response`

NewSearchReddit200ResponseWithDefaults instantiates a new SearchReddit200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SearchReddit200Response) GetItems() []RedditPost`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SearchReddit200Response) GetItemsOk() (*[]RedditPost, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SearchReddit200Response) SetItems(v []RedditPost)`

SetItems sets Items field to given value.

### HasItems

`func (o *SearchReddit200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAfter

`func (o *SearchReddit200Response) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *SearchReddit200Response) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *SearchReddit200Response) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *SearchReddit200Response) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *SearchReddit200Response) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *SearchReddit200Response) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetBefore

`func (o *SearchReddit200Response) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *SearchReddit200Response) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *SearchReddit200Response) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *SearchReddit200Response) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### SetBeforeNil

`func (o *SearchReddit200Response) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *SearchReddit200Response) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


