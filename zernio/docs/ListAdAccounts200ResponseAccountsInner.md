# ListAdAccounts200ResponseAccountsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Platform ad account ID (e.g. act_123) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimezoneName** | Pointer to **string** | IANA timezone of the ad account (Meta only). Drives daily-budget reset and Insights day boundaries. | [optional] 
**TimezoneOffsetHoursUtc** | Pointer to **float32** | Signed UTC offset in hours, reflecting current DST (Meta only). | [optional] 

## Methods

### NewListAdAccounts200ResponseAccountsInner

`func NewListAdAccounts200ResponseAccountsInner() *ListAdAccounts200ResponseAccountsInner`

NewListAdAccounts200ResponseAccountsInner instantiates a new ListAdAccounts200ResponseAccountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAdAccounts200ResponseAccountsInnerWithDefaults

`func NewListAdAccounts200ResponseAccountsInnerWithDefaults() *ListAdAccounts200ResponseAccountsInner`

NewListAdAccounts200ResponseAccountsInnerWithDefaults instantiates a new ListAdAccounts200ResponseAccountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListAdAccounts200ResponseAccountsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAdAccounts200ResponseAccountsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAdAccounts200ResponseAccountsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListAdAccounts200ResponseAccountsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListAdAccounts200ResponseAccountsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListAdAccounts200ResponseAccountsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListAdAccounts200ResponseAccountsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListAdAccounts200ResponseAccountsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrency

`func (o *ListAdAccounts200ResponseAccountsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListAdAccounts200ResponseAccountsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListAdAccounts200ResponseAccountsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListAdAccounts200ResponseAccountsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStatus

`func (o *ListAdAccounts200ResponseAccountsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListAdAccounts200ResponseAccountsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListAdAccounts200ResponseAccountsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListAdAccounts200ResponseAccountsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimezoneName

`func (o *ListAdAccounts200ResponseAccountsInner) GetTimezoneName() string`

GetTimezoneName returns the TimezoneName field if non-nil, zero value otherwise.

### GetTimezoneNameOk

`func (o *ListAdAccounts200ResponseAccountsInner) GetTimezoneNameOk() (*string, bool)`

GetTimezoneNameOk returns a tuple with the TimezoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneName

`func (o *ListAdAccounts200ResponseAccountsInner) SetTimezoneName(v string)`

SetTimezoneName sets TimezoneName field to given value.

### HasTimezoneName

`func (o *ListAdAccounts200ResponseAccountsInner) HasTimezoneName() bool`

HasTimezoneName returns a boolean if a field has been set.

### GetTimezoneOffsetHoursUtc

`func (o *ListAdAccounts200ResponseAccountsInner) GetTimezoneOffsetHoursUtc() float32`

GetTimezoneOffsetHoursUtc returns the TimezoneOffsetHoursUtc field if non-nil, zero value otherwise.

### GetTimezoneOffsetHoursUtcOk

`func (o *ListAdAccounts200ResponseAccountsInner) GetTimezoneOffsetHoursUtcOk() (*float32, bool)`

GetTimezoneOffsetHoursUtcOk returns a tuple with the TimezoneOffsetHoursUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneOffsetHoursUtc

`func (o *ListAdAccounts200ResponseAccountsInner) SetTimezoneOffsetHoursUtc(v float32)`

SetTimezoneOffsetHoursUtc sets TimezoneOffsetHoursUtc field to given value.

### HasTimezoneOffsetHoursUtc

`func (o *ListAdAccounts200ResponseAccountsInner) HasTimezoneOffsetHoursUtc() bool`

HasTimezoneOffsetHoursUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


