# ListDiscordGuildMembers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ListDiscordGuildMembers200ResponseDataInner**](ListDiscordGuildMembers200ResponseDataInner.md) |  | [optional] 
**Pagination** | Pointer to [**ListDiscordGuildMembers200ResponsePagination**](ListDiscordGuildMembers200ResponsePagination.md) |  | [optional] 

## Methods

### NewListDiscordGuildMembers200Response

`func NewListDiscordGuildMembers200Response() *ListDiscordGuildMembers200Response`

NewListDiscordGuildMembers200Response instantiates a new ListDiscordGuildMembers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDiscordGuildMembers200ResponseWithDefaults

`func NewListDiscordGuildMembers200ResponseWithDefaults() *ListDiscordGuildMembers200Response`

NewListDiscordGuildMembers200ResponseWithDefaults instantiates a new ListDiscordGuildMembers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListDiscordGuildMembers200Response) GetData() []ListDiscordGuildMembers200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDiscordGuildMembers200Response) GetDataOk() (*[]ListDiscordGuildMembers200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDiscordGuildMembers200Response) SetData(v []ListDiscordGuildMembers200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ListDiscordGuildMembers200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListDiscordGuildMembers200Response) GetPagination() ListDiscordGuildMembers200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListDiscordGuildMembers200Response) GetPaginationOk() (*ListDiscordGuildMembers200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListDiscordGuildMembers200Response) SetPagination(v ListDiscordGuildMembers200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListDiscordGuildMembers200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


