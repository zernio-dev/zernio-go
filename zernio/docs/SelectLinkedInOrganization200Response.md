# SelectLinkedInOrganization200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**RedirectUrl** | Pointer to **string** | The redirect URL with connection params appended (only if redirect_url was provided in request) | [optional] 
**Account** | Pointer to [**SelectLinkedInOrganization200ResponseAccount**](SelectLinkedInOrganization200ResponseAccount.md) |  | [optional] 
**BulkRefresh** | Pointer to [**SelectLinkedInOrganization200ResponseBulkRefresh**](SelectLinkedInOrganization200ResponseBulkRefresh.md) |  | [optional] 

## Methods

### NewSelectLinkedInOrganization200Response

`func NewSelectLinkedInOrganization200Response() *SelectLinkedInOrganization200Response`

NewSelectLinkedInOrganization200Response instantiates a new SelectLinkedInOrganization200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectLinkedInOrganization200ResponseWithDefaults

`func NewSelectLinkedInOrganization200ResponseWithDefaults() *SelectLinkedInOrganization200Response`

NewSelectLinkedInOrganization200ResponseWithDefaults instantiates a new SelectLinkedInOrganization200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SelectLinkedInOrganization200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SelectLinkedInOrganization200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SelectLinkedInOrganization200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SelectLinkedInOrganization200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *SelectLinkedInOrganization200Response) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SelectLinkedInOrganization200Response) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SelectLinkedInOrganization200Response) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *SelectLinkedInOrganization200Response) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetAccount

`func (o *SelectLinkedInOrganization200Response) GetAccount() SelectLinkedInOrganization200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SelectLinkedInOrganization200Response) GetAccountOk() (*SelectLinkedInOrganization200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SelectLinkedInOrganization200Response) SetAccount(v SelectLinkedInOrganization200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SelectLinkedInOrganization200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBulkRefresh

`func (o *SelectLinkedInOrganization200Response) GetBulkRefresh() SelectLinkedInOrganization200ResponseBulkRefresh`

GetBulkRefresh returns the BulkRefresh field if non-nil, zero value otherwise.

### GetBulkRefreshOk

`func (o *SelectLinkedInOrganization200Response) GetBulkRefreshOk() (*SelectLinkedInOrganization200ResponseBulkRefresh, bool)`

GetBulkRefreshOk returns a tuple with the BulkRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkRefresh

`func (o *SelectLinkedInOrganization200Response) SetBulkRefresh(v SelectLinkedInOrganization200ResponseBulkRefresh)`

SetBulkRefresh sets BulkRefresh field to given value.

### HasBulkRefresh

`func (o *SelectLinkedInOrganization200Response) HasBulkRefresh() bool`

HasBulkRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


