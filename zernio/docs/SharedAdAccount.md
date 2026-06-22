# SharedAdAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Ad account id, in &#x60;act_&lt;digits&gt;&#x60; form. | 
**Name** | Pointer to **string** |  | [optional] 
**BusinessId** | Pointer to **string** | Business Manager id that owns the ad account, when reported. | [optional] 

## Methods

### NewSharedAdAccount

`func NewSharedAdAccount(id string, ) *SharedAdAccount`

NewSharedAdAccount instantiates a new SharedAdAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedAdAccountWithDefaults

`func NewSharedAdAccountWithDefaults() *SharedAdAccount`

NewSharedAdAccountWithDefaults instantiates a new SharedAdAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SharedAdAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedAdAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedAdAccount) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SharedAdAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedAdAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedAdAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharedAdAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBusinessId

`func (o *SharedAdAccount) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *SharedAdAccount) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *SharedAdAccount) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *SharedAdAccount) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


