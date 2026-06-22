# ThreadsPlatformData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopicTag** | Pointer to **string** | Topic tag for post categorization and discoverability on Threads. Must be 1-50 characters, cannot contain periods (.) or ampersands (&amp;). Overrides auto-extraction from content hashtags when provided. | [optional] 
**ThreadItems** | Pointer to [**[]TwitterPlatformDataThreadItemsInner**](TwitterPlatformDataThreadItemsInner.md) | Complete sequence of posts in a Threads thread. The first item becomes the root post, subsequent items are chained as replies. When threadItems is provided, the top-level content field is used only for display and search purposes, it is NOT published. You must include your first post as threadItems[0].  | [optional] 

## Methods

### NewThreadsPlatformData

`func NewThreadsPlatformData() *ThreadsPlatformData`

NewThreadsPlatformData instantiates a new ThreadsPlatformData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadsPlatformDataWithDefaults

`func NewThreadsPlatformDataWithDefaults() *ThreadsPlatformData`

NewThreadsPlatformDataWithDefaults instantiates a new ThreadsPlatformData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopicTag

`func (o *ThreadsPlatformData) GetTopicTag() string`

GetTopicTag returns the TopicTag field if non-nil, zero value otherwise.

### GetTopicTagOk

`func (o *ThreadsPlatformData) GetTopicTagOk() (*string, bool)`

GetTopicTagOk returns a tuple with the TopicTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicTag

`func (o *ThreadsPlatformData) SetTopicTag(v string)`

SetTopicTag sets TopicTag field to given value.

### HasTopicTag

`func (o *ThreadsPlatformData) HasTopicTag() bool`

HasTopicTag returns a boolean if a field has been set.

### GetThreadItems

`func (o *ThreadsPlatformData) GetThreadItems() []TwitterPlatformDataThreadItemsInner`

GetThreadItems returns the ThreadItems field if non-nil, zero value otherwise.

### GetThreadItemsOk

`func (o *ThreadsPlatformData) GetThreadItemsOk() (*[]TwitterPlatformDataThreadItemsInner, bool)`

GetThreadItemsOk returns a tuple with the ThreadItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadItems

`func (o *ThreadsPlatformData) SetThreadItems(v []TwitterPlatformDataThreadItemsInner)`

SetThreadItems sets ThreadItems field to given value.

### HasThreadItems

`func (o *ThreadsPlatformData) HasThreadItems() bool`

HasThreadItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


