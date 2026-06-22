# BlueskyPlatformData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadItems** | Pointer to [**[]TwitterPlatformDataThreadItemsInner**](TwitterPlatformDataThreadItemsInner.md) | Complete sequence of posts in a Bluesky thread. The first item becomes the root post, subsequent items are chained as replies. When threadItems is provided, the top-level content field is used only for display and search purposes, it is NOT published. You must include your first post as threadItems[0].  | [optional] 

## Methods

### NewBlueskyPlatformData

`func NewBlueskyPlatformData() *BlueskyPlatformData`

NewBlueskyPlatformData instantiates a new BlueskyPlatformData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueskyPlatformDataWithDefaults

`func NewBlueskyPlatformDataWithDefaults() *BlueskyPlatformData`

NewBlueskyPlatformDataWithDefaults instantiates a new BlueskyPlatformData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadItems

`func (o *BlueskyPlatformData) GetThreadItems() []TwitterPlatformDataThreadItemsInner`

GetThreadItems returns the ThreadItems field if non-nil, zero value otherwise.

### GetThreadItemsOk

`func (o *BlueskyPlatformData) GetThreadItemsOk() (*[]TwitterPlatformDataThreadItemsInner, bool)`

GetThreadItemsOk returns a tuple with the ThreadItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadItems

`func (o *BlueskyPlatformData) SetThreadItems(v []TwitterPlatformDataThreadItemsInner)`

SetThreadItems sets ThreadItems field to given value.

### HasThreadItems

`func (o *BlueskyPlatformData) HasThreadItems() bool`

HasThreadItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


