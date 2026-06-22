# SendInboxMessage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**SendInboxMessage200ResponseData**](SendInboxMessage200ResponseData.md) |  | [optional] 

## Methods

### NewSendInboxMessage200Response

`func NewSendInboxMessage200Response() *SendInboxMessage200Response`

NewSendInboxMessage200Response instantiates a new SendInboxMessage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessage200ResponseWithDefaults

`func NewSendInboxMessage200ResponseWithDefaults() *SendInboxMessage200Response`

NewSendInboxMessage200ResponseWithDefaults instantiates a new SendInboxMessage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SendInboxMessage200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SendInboxMessage200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SendInboxMessage200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SendInboxMessage200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *SendInboxMessage200Response) GetData() SendInboxMessage200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SendInboxMessage200Response) GetDataOk() (*SendInboxMessage200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SendInboxMessage200Response) SetData(v SendInboxMessage200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *SendInboxMessage200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


