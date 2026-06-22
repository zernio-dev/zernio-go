# ListWhatsAppGroupChats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]ListWhatsAppGroupChats200ResponseGroupsInner**](ListWhatsAppGroupChats200ResponseGroupsInner.md) |  | [optional] 
**Paging** | Pointer to [**ListWhatsAppGroupChats200ResponsePaging**](ListWhatsAppGroupChats200ResponsePaging.md) |  | [optional] 

## Methods

### NewListWhatsAppGroupChats200Response

`func NewListWhatsAppGroupChats200Response() *ListWhatsAppGroupChats200Response`

NewListWhatsAppGroupChats200Response instantiates a new ListWhatsAppGroupChats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWhatsAppGroupChats200ResponseWithDefaults

`func NewListWhatsAppGroupChats200ResponseWithDefaults() *ListWhatsAppGroupChats200Response`

NewListWhatsAppGroupChats200ResponseWithDefaults instantiates a new ListWhatsAppGroupChats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *ListWhatsAppGroupChats200Response) GetGroups() []ListWhatsAppGroupChats200ResponseGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ListWhatsAppGroupChats200Response) GetGroupsOk() (*[]ListWhatsAppGroupChats200ResponseGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ListWhatsAppGroupChats200Response) SetGroups(v []ListWhatsAppGroupChats200ResponseGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ListWhatsAppGroupChats200Response) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetPaging

`func (o *ListWhatsAppGroupChats200Response) GetPaging() ListWhatsAppGroupChats200ResponsePaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ListWhatsAppGroupChats200Response) GetPagingOk() (*ListWhatsAppGroupChats200ResponsePaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ListWhatsAppGroupChats200Response) SetPaging(v ListWhatsAppGroupChats200ResponsePaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ListWhatsAppGroupChats200Response) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


