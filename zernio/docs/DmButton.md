# DmButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | **string** | Button label (20 chars max) | 
**Url** | Pointer to **string** | Target URL (required when type is url) | [optional] 
**Payload** | Pointer to **string** | Postback payload delivered via the messaging_postbacks webhook (required when type is postback) | [optional] 
**Phone** | Pointer to **string** | Phone number, e.g. +14155551234 (required when type is phone; Facebook only) | [optional] 

## Methods

### NewDmButton

`func NewDmButton(type_ string, title string, ) *DmButton`

NewDmButton instantiates a new DmButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmButtonWithDefaults

`func NewDmButtonWithDefaults() *DmButton`

NewDmButtonWithDefaults instantiates a new DmButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DmButton) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DmButton) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DmButton) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *DmButton) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DmButton) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DmButton) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *DmButton) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DmButton) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DmButton) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DmButton) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPayload

`func (o *DmButton) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *DmButton) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *DmButton) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *DmButton) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPhone

`func (o *DmButton) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *DmButton) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *DmButton) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *DmButton) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


