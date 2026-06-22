# GetConnectUrl200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthUrl** | Pointer to **string** | URL to redirect your user to for OAuth authorization | [optional] 
**State** | Pointer to **string** | State parameter for security (handled automatically) | [optional] 

## Methods

### NewGetConnectUrl200Response

`func NewGetConnectUrl200Response() *GetConnectUrl200Response`

NewGetConnectUrl200Response instantiates a new GetConnectUrl200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectUrl200ResponseWithDefaults

`func NewGetConnectUrl200ResponseWithDefaults() *GetConnectUrl200Response`

NewGetConnectUrl200ResponseWithDefaults instantiates a new GetConnectUrl200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthUrl

`func (o *GetConnectUrl200Response) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *GetConnectUrl200Response) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *GetConnectUrl200Response) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.

### HasAuthUrl

`func (o *GetConnectUrl200Response) HasAuthUrl() bool`

HasAuthUrl returns a boolean if a field has been set.

### GetState

`func (o *GetConnectUrl200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetConnectUrl200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetConnectUrl200Response) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetConnectUrl200Response) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


