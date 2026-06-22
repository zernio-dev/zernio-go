# DeleteDiscordScheduledEvent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Deleted** | Pointer to **string** | The deleted event&#39;s snowflake. | [optional] 

## Methods

### NewDeleteDiscordScheduledEvent200Response

`func NewDeleteDiscordScheduledEvent200Response() *DeleteDiscordScheduledEvent200Response`

NewDeleteDiscordScheduledEvent200Response instantiates a new DeleteDiscordScheduledEvent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDiscordScheduledEvent200ResponseWithDefaults

`func NewDeleteDiscordScheduledEvent200ResponseWithDefaults() *DeleteDiscordScheduledEvent200Response`

NewDeleteDiscordScheduledEvent200ResponseWithDefaults instantiates a new DeleteDiscordScheduledEvent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DeleteDiscordScheduledEvent200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DeleteDiscordScheduledEvent200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DeleteDiscordScheduledEvent200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DeleteDiscordScheduledEvent200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetDeleted

`func (o *DeleteDiscordScheduledEvent200Response) GetDeleted() string`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteDiscordScheduledEvent200Response) GetDeletedOk() (*string, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteDiscordScheduledEvent200Response) SetDeleted(v string)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeleteDiscordScheduledEvent200Response) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


