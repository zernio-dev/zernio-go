# CreateLeadForm200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Form** | Pointer to [**UpdateFacebookPage200ResponseSelectedPage**](UpdateFacebookPage200ResponseSelectedPage.md) |  | [optional] 

## Methods

### NewCreateLeadForm200Response

`func NewCreateLeadForm200Response() *CreateLeadForm200Response`

NewCreateLeadForm200Response instantiates a new CreateLeadForm200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLeadForm200ResponseWithDefaults

`func NewCreateLeadForm200ResponseWithDefaults() *CreateLeadForm200Response`

NewCreateLeadForm200ResponseWithDefaults instantiates a new CreateLeadForm200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateLeadForm200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateLeadForm200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateLeadForm200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateLeadForm200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetForm

`func (o *CreateLeadForm200Response) GetForm() UpdateFacebookPage200ResponseSelectedPage`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *CreateLeadForm200Response) GetFormOk() (*UpdateFacebookPage200ResponseSelectedPage, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *CreateLeadForm200Response) SetForm(v UpdateFacebookPage200ResponseSelectedPage)`

SetForm sets Form field to given value.

### HasForm

`func (o *CreateLeadForm200Response) HasForm() bool`

HasForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


