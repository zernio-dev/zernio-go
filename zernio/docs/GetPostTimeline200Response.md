# GetPostTimeline200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostId** | Pointer to **string** | The postId that was requested | [optional] 
**Timeline** | Pointer to [**[]GetPostTimeline200ResponseTimelineInner**](GetPostTimeline200ResponseTimelineInner.md) |  | [optional] 

## Methods

### NewGetPostTimeline200Response

`func NewGetPostTimeline200Response() *GetPostTimeline200Response`

NewGetPostTimeline200Response instantiates a new GetPostTimeline200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPostTimeline200ResponseWithDefaults

`func NewGetPostTimeline200ResponseWithDefaults() *GetPostTimeline200Response`

NewGetPostTimeline200ResponseWithDefaults instantiates a new GetPostTimeline200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostId

`func (o *GetPostTimeline200Response) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *GetPostTimeline200Response) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *GetPostTimeline200Response) SetPostId(v string)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *GetPostTimeline200Response) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetTimeline

`func (o *GetPostTimeline200Response) GetTimeline() []GetPostTimeline200ResponseTimelineInner`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *GetPostTimeline200Response) GetTimelineOk() (*[]GetPostTimeline200ResponseTimelineInner, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *GetPostTimeline200Response) SetTimeline(v []GetPostTimeline200ResponseTimelineInner)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *GetPostTimeline200Response) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


