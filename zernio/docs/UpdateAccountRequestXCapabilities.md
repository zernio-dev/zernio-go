# UpdateAccountRequestXCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analytics** | Pointer to **bool** | Enable periodic analytics reads (impressions, likes, etc.) for this X account. Each X API call is metered as &#x60;posts_read&#x60; and billed pass-through (~$0.005/call at the time of writing — actual rate depends on X&#39;s pricing tier).  | [optional] 
**Inbox** | Pointer to **bool** | Enable DM polling and inbox sync for this X account. DM reads are metered as &#x60;dm_event_read&#x60; (~$0.010/call) and DM sends as &#x60;dm_interaction_create&#x60; (~$0.015/call), both billed pass-through. DM sends fire only on user-initiated actions; reads/polling fire only when this flag is true.  | [optional] 

## Methods

### NewUpdateAccountRequestXCapabilities

`func NewUpdateAccountRequestXCapabilities() *UpdateAccountRequestXCapabilities`

NewUpdateAccountRequestXCapabilities instantiates a new UpdateAccountRequestXCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountRequestXCapabilitiesWithDefaults

`func NewUpdateAccountRequestXCapabilitiesWithDefaults() *UpdateAccountRequestXCapabilities`

NewUpdateAccountRequestXCapabilitiesWithDefaults instantiates a new UpdateAccountRequestXCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalytics

`func (o *UpdateAccountRequestXCapabilities) GetAnalytics() bool`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *UpdateAccountRequestXCapabilities) GetAnalyticsOk() (*bool, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *UpdateAccountRequestXCapabilities) SetAnalytics(v bool)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *UpdateAccountRequestXCapabilities) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### GetInbox

`func (o *UpdateAccountRequestXCapabilities) GetInbox() bool`

GetInbox returns the Inbox field if non-nil, zero value otherwise.

### GetInboxOk

`func (o *UpdateAccountRequestXCapabilities) GetInboxOk() (*bool, bool)`

GetInboxOk returns a tuple with the Inbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbox

`func (o *UpdateAccountRequestXCapabilities) SetInbox(v bool)`

SetInbox sets Inbox field to given value.

### HasInbox

`func (o *UpdateAccountRequestXCapabilities) HasInbox() bool`

HasInbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


