# ListCommentAutomations200ResponseAutomationsInnerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggered** | Pointer to **int32** |  | [optional] 
**DmsSent** | Pointer to **int32** |  | [optional] 
**DmsFailed** | Pointer to **int32** |  | [optional] 
**UniqueContacts** | Pointer to **int32** |  | [optional] 
**TrackedSends** | Pointer to **int32** | DMs sent with a trackable (wrapped) link. CTR denominator: divide clicks by this, not dmsSent. Lags dmsSent for campaigns that predate click tracking. | [optional] 
**LinkClicks** | Pointer to **int32** | Total clicks on tracked links (bots/prefetch excluded). | [optional] 
**UniqueClicks** | Pointer to **int32** | Distinct people who clicked a tracked link. | [optional] 
**Delivered** | Pointer to **int32** | DMs confirmed delivered (Messenger; IG emits no delivery receipt). | [optional] 
**Read** | Pointer to **int32** | DMs confirmed read (IG messaging_seen / Messenger message_reads). | [optional] 

## Methods

### NewListCommentAutomations200ResponseAutomationsInnerStats

`func NewListCommentAutomations200ResponseAutomationsInnerStats() *ListCommentAutomations200ResponseAutomationsInnerStats`

NewListCommentAutomations200ResponseAutomationsInnerStats instantiates a new ListCommentAutomations200ResponseAutomationsInnerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCommentAutomations200ResponseAutomationsInnerStatsWithDefaults

`func NewListCommentAutomations200ResponseAutomationsInnerStatsWithDefaults() *ListCommentAutomations200ResponseAutomationsInnerStats`

NewListCommentAutomations200ResponseAutomationsInnerStatsWithDefaults instantiates a new ListCommentAutomations200ResponseAutomationsInnerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggered

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetTriggered() int32`

GetTriggered returns the Triggered field if non-nil, zero value otherwise.

### GetTriggeredOk

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetTriggeredOk() (*int32, bool)`

GetTriggeredOk returns a tuple with the Triggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggered

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) SetTriggered(v int32)`

SetTriggered sets Triggered field to given value.

### HasTriggered

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) HasTriggered() bool`

HasTriggered returns a boolean if a field has been set.

### GetDmsSent

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetDmsSent() int32`

GetDmsSent returns the DmsSent field if non-nil, zero value otherwise.

### GetDmsSentOk

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetDmsSentOk() (*int32, bool)`

GetDmsSentOk returns a tuple with the DmsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmsSent

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) SetDmsSent(v int32)`

SetDmsSent sets DmsSent field to given value.

### HasDmsSent

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) HasDmsSent() bool`

HasDmsSent returns a boolean if a field has been set.

### GetDmsFailed

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetDmsFailed() int32`

GetDmsFailed returns the DmsFailed field if non-nil, zero value otherwise.

### GetDmsFailedOk

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetDmsFailedOk() (*int32, bool)`

GetDmsFailedOk returns a tuple with the DmsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmsFailed

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) SetDmsFailed(v int32)`

SetDmsFailed sets DmsFailed field to given value.

### HasDmsFailed

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) HasDmsFailed() bool`

HasDmsFailed returns a boolean if a field has been set.

### GetUniqueContacts

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetUniqueContacts() int32`

GetUniqueContacts returns the UniqueContacts field if non-nil, zero value otherwise.

### GetUniqueContactsOk

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetUniqueContactsOk() (*int32, bool)`

GetUniqueContactsOk returns a tuple with the UniqueContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueContacts

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) SetUniqueContacts(v int32)`

SetUniqueContacts sets UniqueContacts field to given value.

### HasUniqueContacts

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) HasUniqueContacts() bool`

HasUniqueContacts returns a boolean if a field has been set.

### GetTrackedSends

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetTrackedSends() int32`

GetTrackedSends returns the TrackedSends field if non-nil, zero value otherwise.

### GetTrackedSendsOk

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetTrackedSendsOk() (*int32, bool)`

GetTrackedSendsOk returns a tuple with the TrackedSends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedSends

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) SetTrackedSends(v int32)`

SetTrackedSends sets TrackedSends field to given value.

### HasTrackedSends

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) HasTrackedSends() bool`

HasTrackedSends returns a boolean if a field has been set.

### GetLinkClicks

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetLinkClicks() int32`

GetLinkClicks returns the LinkClicks field if non-nil, zero value otherwise.

### GetLinkClicksOk

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetLinkClicksOk() (*int32, bool)`

GetLinkClicksOk returns a tuple with the LinkClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkClicks

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) SetLinkClicks(v int32)`

SetLinkClicks sets LinkClicks field to given value.

### HasLinkClicks

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) HasLinkClicks() bool`

HasLinkClicks returns a boolean if a field has been set.

### GetUniqueClicks

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetUniqueClicks() int32`

GetUniqueClicks returns the UniqueClicks field if non-nil, zero value otherwise.

### GetUniqueClicksOk

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetUniqueClicksOk() (*int32, bool)`

GetUniqueClicksOk returns a tuple with the UniqueClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueClicks

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) SetUniqueClicks(v int32)`

SetUniqueClicks sets UniqueClicks field to given value.

### HasUniqueClicks

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) HasUniqueClicks() bool`

HasUniqueClicks returns a boolean if a field has been set.

### GetDelivered

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetDelivered() int32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetDeliveredOk() (*int32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) SetDelivered(v int32)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetRead

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetRead() int32`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) GetReadOk() (*int32, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) SetRead(v int32)`

SetRead sets Read field to given value.

### HasRead

`func (o *ListCommentAutomations200ResponseAutomationsInnerStats) HasRead() bool`

HasRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


