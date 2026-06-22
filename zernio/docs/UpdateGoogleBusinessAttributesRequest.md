# UpdateGoogleBusinessAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**[]UpdateGoogleBusinessAttributesRequestAttributesInner**](UpdateGoogleBusinessAttributesRequestAttributesInner.md) |  | 
**AttributeMask** | **string** | Comma-separated attribute names to update (e.g. &#39;has_delivery,has_takeout&#39;) | 

## Methods

### NewUpdateGoogleBusinessAttributesRequest

`func NewUpdateGoogleBusinessAttributesRequest(attributes []UpdateGoogleBusinessAttributesRequestAttributesInner, attributeMask string, ) *UpdateGoogleBusinessAttributesRequest`

NewUpdateGoogleBusinessAttributesRequest instantiates a new UpdateGoogleBusinessAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGoogleBusinessAttributesRequestWithDefaults

`func NewUpdateGoogleBusinessAttributesRequestWithDefaults() *UpdateGoogleBusinessAttributesRequest`

NewUpdateGoogleBusinessAttributesRequestWithDefaults instantiates a new UpdateGoogleBusinessAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UpdateGoogleBusinessAttributesRequest) GetAttributes() []UpdateGoogleBusinessAttributesRequestAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateGoogleBusinessAttributesRequest) GetAttributesOk() (*[]UpdateGoogleBusinessAttributesRequestAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateGoogleBusinessAttributesRequest) SetAttributes(v []UpdateGoogleBusinessAttributesRequestAttributesInner)`

SetAttributes sets Attributes field to given value.


### GetAttributeMask

`func (o *UpdateGoogleBusinessAttributesRequest) GetAttributeMask() string`

GetAttributeMask returns the AttributeMask field if non-nil, zero value otherwise.

### GetAttributeMaskOk

`func (o *UpdateGoogleBusinessAttributesRequest) GetAttributeMaskOk() (*string, bool)`

GetAttributeMaskOk returns a tuple with the AttributeMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMask

`func (o *UpdateGoogleBusinessAttributesRequest) SetAttributeMask(v string)`

SetAttributeMask sets AttributeMask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


