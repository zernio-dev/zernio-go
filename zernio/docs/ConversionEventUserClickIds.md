# ConversionEventUserClickIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fbc** | Pointer to **string** | Meta click ID (from fbclid URL param). | [optional] 
**Fbp** | Pointer to **string** | Meta browser ID (_fbp cookie). | [optional] 
**Gclid** | Pointer to **string** | Google click ID (from gclid URL param). | [optional] 
**Gbraid** | Pointer to **string** | Google iOS 14.5+ app attribution ID. | [optional] 
**Wbraid** | Pointer to **string** | Google iOS 14.5+ web-to-app attribution ID. | [optional] 
**LiFatId** | Pointer to **string** | LinkedIn first-party ad tracking click ID. Captured by parsing &#x60;li_fat_id&#x60; from landing-page URLs after the advertiser enables enhanced conversion tracking on the LinkedIn Insight Tag. Sent to LinkedIn as the LINKEDIN_FIRST_PARTY_ADS_TRACKING_UUID userId. Opaque token, not hashed.  | [optional] 

## Methods

### NewConversionEventUserClickIds

`func NewConversionEventUserClickIds() *ConversionEventUserClickIds`

NewConversionEventUserClickIds instantiates a new ConversionEventUserClickIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionEventUserClickIdsWithDefaults

`func NewConversionEventUserClickIdsWithDefaults() *ConversionEventUserClickIds`

NewConversionEventUserClickIdsWithDefaults instantiates a new ConversionEventUserClickIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFbc

`func (o *ConversionEventUserClickIds) GetFbc() string`

GetFbc returns the Fbc field if non-nil, zero value otherwise.

### GetFbcOk

`func (o *ConversionEventUserClickIds) GetFbcOk() (*string, bool)`

GetFbcOk returns a tuple with the Fbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFbc

`func (o *ConversionEventUserClickIds) SetFbc(v string)`

SetFbc sets Fbc field to given value.

### HasFbc

`func (o *ConversionEventUserClickIds) HasFbc() bool`

HasFbc returns a boolean if a field has been set.

### GetFbp

`func (o *ConversionEventUserClickIds) GetFbp() string`

GetFbp returns the Fbp field if non-nil, zero value otherwise.

### GetFbpOk

`func (o *ConversionEventUserClickIds) GetFbpOk() (*string, bool)`

GetFbpOk returns a tuple with the Fbp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFbp

`func (o *ConversionEventUserClickIds) SetFbp(v string)`

SetFbp sets Fbp field to given value.

### HasFbp

`func (o *ConversionEventUserClickIds) HasFbp() bool`

HasFbp returns a boolean if a field has been set.

### GetGclid

`func (o *ConversionEventUserClickIds) GetGclid() string`

GetGclid returns the Gclid field if non-nil, zero value otherwise.

### GetGclidOk

`func (o *ConversionEventUserClickIds) GetGclidOk() (*string, bool)`

GetGclidOk returns a tuple with the Gclid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGclid

`func (o *ConversionEventUserClickIds) SetGclid(v string)`

SetGclid sets Gclid field to given value.

### HasGclid

`func (o *ConversionEventUserClickIds) HasGclid() bool`

HasGclid returns a boolean if a field has been set.

### GetGbraid

`func (o *ConversionEventUserClickIds) GetGbraid() string`

GetGbraid returns the Gbraid field if non-nil, zero value otherwise.

### GetGbraidOk

`func (o *ConversionEventUserClickIds) GetGbraidOk() (*string, bool)`

GetGbraidOk returns a tuple with the Gbraid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbraid

`func (o *ConversionEventUserClickIds) SetGbraid(v string)`

SetGbraid sets Gbraid field to given value.

### HasGbraid

`func (o *ConversionEventUserClickIds) HasGbraid() bool`

HasGbraid returns a boolean if a field has been set.

### GetWbraid

`func (o *ConversionEventUserClickIds) GetWbraid() string`

GetWbraid returns the Wbraid field if non-nil, zero value otherwise.

### GetWbraidOk

`func (o *ConversionEventUserClickIds) GetWbraidOk() (*string, bool)`

GetWbraidOk returns a tuple with the Wbraid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbraid

`func (o *ConversionEventUserClickIds) SetWbraid(v string)`

SetWbraid sets Wbraid field to given value.

### HasWbraid

`func (o *ConversionEventUserClickIds) HasWbraid() bool`

HasWbraid returns a boolean if a field has been set.

### GetLiFatId

`func (o *ConversionEventUserClickIds) GetLiFatId() string`

GetLiFatId returns the LiFatId field if non-nil, zero value otherwise.

### GetLiFatIdOk

`func (o *ConversionEventUserClickIds) GetLiFatIdOk() (*string, bool)`

GetLiFatIdOk returns a tuple with the LiFatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiFatId

`func (o *ConversionEventUserClickIds) SetLiFatId(v string)`

SetLiFatId sets LiFatId field to given value.

### HasLiFatId

`func (o *ConversionEventUserClickIds) HasLiFatId() bool`

HasLiFatId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


