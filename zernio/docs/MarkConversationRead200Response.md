# MarkConversationRead200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**MarkedCount** | Pointer to **int32** | Number of messages marked read by this call | [optional] 

## Methods

### NewMarkConversationRead200Response

`func NewMarkConversationRead200Response() *MarkConversationRead200Response`

NewMarkConversationRead200Response instantiates a new MarkConversationRead200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarkConversationRead200ResponseWithDefaults

`func NewMarkConversationRead200ResponseWithDefaults() *MarkConversationRead200Response`

NewMarkConversationRead200ResponseWithDefaults instantiates a new MarkConversationRead200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *MarkConversationRead200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MarkConversationRead200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MarkConversationRead200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *MarkConversationRead200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMarkedCount

`func (o *MarkConversationRead200Response) GetMarkedCount() int32`

GetMarkedCount returns the MarkedCount field if non-nil, zero value otherwise.

### GetMarkedCountOk

`func (o *MarkConversationRead200Response) GetMarkedCountOk() (*int32, bool)`

GetMarkedCountOk returns a tuple with the MarkedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedCount

`func (o *MarkConversationRead200Response) SetMarkedCount(v int32)`

SetMarkedCount sets MarkedCount field to given value.

### HasMarkedCount

`func (o *MarkConversationRead200Response) HasMarkedCount() bool`

HasMarkedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


