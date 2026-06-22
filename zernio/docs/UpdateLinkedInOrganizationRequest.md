# UpdateLinkedInOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | **string** |  | 
**SelectedOrganization** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateLinkedInOrganizationRequest

`func NewUpdateLinkedInOrganizationRequest(accountType string, ) *UpdateLinkedInOrganizationRequest`

NewUpdateLinkedInOrganizationRequest instantiates a new UpdateLinkedInOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLinkedInOrganizationRequestWithDefaults

`func NewUpdateLinkedInOrganizationRequestWithDefaults() *UpdateLinkedInOrganizationRequest`

NewUpdateLinkedInOrganizationRequestWithDefaults instantiates a new UpdateLinkedInOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *UpdateLinkedInOrganizationRequest) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *UpdateLinkedInOrganizationRequest) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *UpdateLinkedInOrganizationRequest) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetSelectedOrganization

`func (o *UpdateLinkedInOrganizationRequest) GetSelectedOrganization() map[string]interface{}`

GetSelectedOrganization returns the SelectedOrganization field if non-nil, zero value otherwise.

### GetSelectedOrganizationOk

`func (o *UpdateLinkedInOrganizationRequest) GetSelectedOrganizationOk() (*map[string]interface{}, bool)`

GetSelectedOrganizationOk returns a tuple with the SelectedOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedOrganization

`func (o *UpdateLinkedInOrganizationRequest) SetSelectedOrganization(v map[string]interface{})`

SetSelectedOrganization sets SelectedOrganization field to given value.

### HasSelectedOrganization

`func (o *UpdateLinkedInOrganizationRequest) HasSelectedOrganization() bool`

HasSelectedOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


