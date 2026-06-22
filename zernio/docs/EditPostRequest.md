# EditPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** | The platform to edit the post on. Currently only twitter is supported. | 
**Content** | **string** | The new tweet text content | 

## Methods

### NewEditPostRequest

`func NewEditPostRequest(platform string, content string, ) *EditPostRequest`

NewEditPostRequest instantiates a new EditPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditPostRequestWithDefaults

`func NewEditPostRequestWithDefaults() *EditPostRequest`

NewEditPostRequestWithDefaults instantiates a new EditPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *EditPostRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *EditPostRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *EditPostRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetContent

`func (o *EditPostRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EditPostRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EditPostRequest) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


