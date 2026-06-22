# ConversionDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Platform-native identifier. Pass back as &#x60;destinationId&#x60; on event send and as the path segment on CRUD endpoints.  | 
**Name** | **string** |  | 
**Type** | Pointer to **string** | Present when the platform locks the event type/category to the destination (Google conversion actions, LinkedIn conversion rules). Absent for Meta pixels (which accept any event name per request).  | [optional] 
**Status** | Pointer to **string** | For LinkedIn, &#x60;inactive&#x60; means the rule is soft-deleted (&#x60;enabled: false&#x60;).  | [optional] 
**AdAccountId** | Pointer to **string** | Set by adapters whose destinations are scoped to a specific ad account (LinkedIn). Pass back on subsequent CRUD calls to identify the parent ad account.  | [optional] 

## Methods

### NewConversionDestination

`func NewConversionDestination(id string, name string, ) *ConversionDestination`

NewConversionDestination instantiates a new ConversionDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionDestinationWithDefaults

`func NewConversionDestinationWithDefaults() *ConversionDestination`

NewConversionDestinationWithDefaults instantiates a new ConversionDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConversionDestination) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConversionDestination) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConversionDestination) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ConversionDestination) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConversionDestination) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConversionDestination) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ConversionDestination) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConversionDestination) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConversionDestination) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConversionDestination) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *ConversionDestination) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConversionDestination) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConversionDestination) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConversionDestination) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAdAccountId

`func (o *ConversionDestination) GetAdAccountId() string`

GetAdAccountId returns the AdAccountId field if non-nil, zero value otherwise.

### GetAdAccountIdOk

`func (o *ConversionDestination) GetAdAccountIdOk() (*string, bool)`

GetAdAccountIdOk returns a tuple with the AdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountId

`func (o *ConversionDestination) SetAdAccountId(v string)`

SetAdAccountId sets AdAccountId field to given value.

### HasAdAccountId

`func (o *ConversionDestination) HasAdAccountId() bool`

HasAdAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


