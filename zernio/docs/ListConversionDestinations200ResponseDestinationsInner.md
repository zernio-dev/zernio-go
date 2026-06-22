# ListConversionDestinations200ResponseDestinationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Destination identifier. Meta: pixel ID. Google: conversion action resource name. LinkedIn: numeric conversion rule ID.  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Present when the platform locks event type to the destination (Google conversion actions, LinkedIn conversion rules).  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AdAccountId** | Pointer to **string** | Set by adapters whose destinations are scoped to a specific ad account (LinkedIn). Pass back on subsequent CRUD calls.  | [optional] 

## Methods

### NewListConversionDestinations200ResponseDestinationsInner

`func NewListConversionDestinations200ResponseDestinationsInner() *ListConversionDestinations200ResponseDestinationsInner`

NewListConversionDestinations200ResponseDestinationsInner instantiates a new ListConversionDestinations200ResponseDestinationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConversionDestinations200ResponseDestinationsInnerWithDefaults

`func NewListConversionDestinations200ResponseDestinationsInnerWithDefaults() *ListConversionDestinations200ResponseDestinationsInner`

NewListConversionDestinations200ResponseDestinationsInnerWithDefaults instantiates a new ListConversionDestinations200ResponseDestinationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListConversionDestinations200ResponseDestinationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListConversionDestinations200ResponseDestinationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListConversionDestinations200ResponseDestinationsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListConversionDestinations200ResponseDestinationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListConversionDestinations200ResponseDestinationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListConversionDestinations200ResponseDestinationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListConversionDestinations200ResponseDestinationsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListConversionDestinations200ResponseDestinationsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListConversionDestinations200ResponseDestinationsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListConversionDestinations200ResponseDestinationsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListConversionDestinations200ResponseDestinationsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListConversionDestinations200ResponseDestinationsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *ListConversionDestinations200ResponseDestinationsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListConversionDestinations200ResponseDestinationsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListConversionDestinations200ResponseDestinationsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListConversionDestinations200ResponseDestinationsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAdAccountId

`func (o *ListConversionDestinations200ResponseDestinationsInner) GetAdAccountId() string`

GetAdAccountId returns the AdAccountId field if non-nil, zero value otherwise.

### GetAdAccountIdOk

`func (o *ListConversionDestinations200ResponseDestinationsInner) GetAdAccountIdOk() (*string, bool)`

GetAdAccountIdOk returns a tuple with the AdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountId

`func (o *ListConversionDestinations200ResponseDestinationsInner) SetAdAccountId(v string)`

SetAdAccountId sets AdAccountId field to given value.

### HasAdAccountId

`func (o *ListConversionDestinations200ResponseDestinationsInner) HasAdAccountId() bool`

HasAdAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


