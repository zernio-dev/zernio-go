# AdjustConversionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | SocialAccount ID. Must be a &#x60;googleads&#x60; account. | 
**DestinationId** | **string** | Conversion action resource name, e.g. &#x60;customers/1234567890/conversionActions/987654321&#x60;. | 
**Adjustments** | [**[]AdjustConversionsRequestAdjustmentsInner**](AdjustConversionsRequestAdjustmentsInner.md) |  | 

## Methods

### NewAdjustConversionsRequest

`func NewAdjustConversionsRequest(accountId string, destinationId string, adjustments []AdjustConversionsRequestAdjustmentsInner, ) *AdjustConversionsRequest`

NewAdjustConversionsRequest instantiates a new AdjustConversionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustConversionsRequestWithDefaults

`func NewAdjustConversionsRequestWithDefaults() *AdjustConversionsRequest`

NewAdjustConversionsRequestWithDefaults instantiates a new AdjustConversionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AdjustConversionsRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AdjustConversionsRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AdjustConversionsRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetDestinationId

`func (o *AdjustConversionsRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *AdjustConversionsRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *AdjustConversionsRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetAdjustments

`func (o *AdjustConversionsRequest) GetAdjustments() []AdjustConversionsRequestAdjustmentsInner`

GetAdjustments returns the Adjustments field if non-nil, zero value otherwise.

### GetAdjustmentsOk

`func (o *AdjustConversionsRequest) GetAdjustmentsOk() (*[]AdjustConversionsRequestAdjustmentsInner, bool)`

GetAdjustmentsOk returns a tuple with the Adjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustments

`func (o *AdjustConversionsRequest) SetAdjustments(v []AdjustConversionsRequestAdjustmentsInner)`

SetAdjustments sets Adjustments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


