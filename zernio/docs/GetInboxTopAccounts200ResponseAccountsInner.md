# GetInboxTopAccounts200ResponseAccountsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | (disconnected) when the SocialAccount no longer exists | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Received** | Pointer to **int32** |  | [optional] 
**Sent** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Conversations** | Pointer to **int32** |  | [optional] 
**MedianResponseSeconds** | Pointer to **int32** |  | [optional] 
**RepliedCount** | Pointer to **int32** | Distinguishes &#39;instant replies&#39; from &#39;no replies at all&#39; so a zero medianResponseSeconds with repliedCount&#x3D;0 renders as &#39;—&#39; instead of &#39;0s&#39; | [optional] 

## Methods

### NewGetInboxTopAccounts200ResponseAccountsInner

`func NewGetInboxTopAccounts200ResponseAccountsInner() *GetInboxTopAccounts200ResponseAccountsInner`

NewGetInboxTopAccounts200ResponseAccountsInner instantiates a new GetInboxTopAccounts200ResponseAccountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxTopAccounts200ResponseAccountsInnerWithDefaults

`func NewGetInboxTopAccounts200ResponseAccountsInnerWithDefaults() *GetInboxTopAccounts200ResponseAccountsInner`

NewGetInboxTopAccounts200ResponseAccountsInnerWithDefaults instantiates a new GetInboxTopAccounts200ResponseAccountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetInboxTopAccounts200ResponseAccountsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetInboxTopAccounts200ResponseAccountsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetInboxTopAccounts200ResponseAccountsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetInboxTopAccounts200ResponseAccountsInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetInboxTopAccounts200ResponseAccountsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetInboxTopAccounts200ResponseAccountsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUsername

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetInboxTopAccounts200ResponseAccountsInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetInboxTopAccounts200ResponseAccountsInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetReceived

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *GetInboxTopAccounts200ResponseAccountsInner) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *GetInboxTopAccounts200ResponseAccountsInner) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetSent

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *GetInboxTopAccounts200ResponseAccountsInner) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *GetInboxTopAccounts200ResponseAccountsInner) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetTotal

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetInboxTopAccounts200ResponseAccountsInner) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetInboxTopAccounts200ResponseAccountsInner) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetConversations

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetConversations() int32`

GetConversations returns the Conversations field if non-nil, zero value otherwise.

### GetConversationsOk

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetConversationsOk() (*int32, bool)`

GetConversationsOk returns a tuple with the Conversations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversations

`func (o *GetInboxTopAccounts200ResponseAccountsInner) SetConversations(v int32)`

SetConversations sets Conversations field to given value.

### HasConversations

`func (o *GetInboxTopAccounts200ResponseAccountsInner) HasConversations() bool`

HasConversations returns a boolean if a field has been set.

### GetMedianResponseSeconds

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetMedianResponseSeconds() int32`

GetMedianResponseSeconds returns the MedianResponseSeconds field if non-nil, zero value otherwise.

### GetMedianResponseSecondsOk

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetMedianResponseSecondsOk() (*int32, bool)`

GetMedianResponseSecondsOk returns a tuple with the MedianResponseSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianResponseSeconds

`func (o *GetInboxTopAccounts200ResponseAccountsInner) SetMedianResponseSeconds(v int32)`

SetMedianResponseSeconds sets MedianResponseSeconds field to given value.

### HasMedianResponseSeconds

`func (o *GetInboxTopAccounts200ResponseAccountsInner) HasMedianResponseSeconds() bool`

HasMedianResponseSeconds returns a boolean if a field has been set.

### GetRepliedCount

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetRepliedCount() int32`

GetRepliedCount returns the RepliedCount field if non-nil, zero value otherwise.

### GetRepliedCountOk

`func (o *GetInboxTopAccounts200ResponseAccountsInner) GetRepliedCountOk() (*int32, bool)`

GetRepliedCountOk returns a tuple with the RepliedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepliedCount

`func (o *GetInboxTopAccounts200ResponseAccountsInner) SetRepliedCount(v int32)`

SetRepliedCount sets RepliedCount field to given value.

### HasRepliedCount

`func (o *GetInboxTopAccounts200ResponseAccountsInner) HasRepliedCount() bool`

HasRepliedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


