# ListBroadcasts200ResponseBroadcastsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** | Display name of the sending account | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**MessagePreview** | Pointer to **string** | Template name or message text snippet | [optional] 
**ScheduledAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**RecipientCount** | Pointer to **int32** |  | [optional] 
**SentCount** | Pointer to **int32** |  | [optional] 
**DeliveredCount** | Pointer to **int32** |  | [optional] 
**ReadCount** | Pointer to **int32** |  | [optional] 
**FailedCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListBroadcasts200ResponseBroadcastsInner

`func NewListBroadcasts200ResponseBroadcastsInner() *ListBroadcasts200ResponseBroadcastsInner`

NewListBroadcasts200ResponseBroadcastsInner instantiates a new ListBroadcasts200ResponseBroadcastsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBroadcasts200ResponseBroadcastsInnerWithDefaults

`func NewListBroadcasts200ResponseBroadcastsInnerWithDefaults() *ListBroadcasts200ResponseBroadcastsInner`

NewListBroadcasts200ResponseBroadcastsInnerWithDefaults instantiates a new ListBroadcasts200ResponseBroadcastsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatform

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountId

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountName

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetStatus

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessagePreview

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetMessagePreview() string`

GetMessagePreview returns the MessagePreview field if non-nil, zero value otherwise.

### GetMessagePreviewOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetMessagePreviewOk() (*string, bool)`

GetMessagePreviewOk returns a tuple with the MessagePreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagePreview

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetMessagePreview(v string)`

SetMessagePreview sets MessagePreview field to given value.

### HasMessagePreview

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasMessagePreview() bool`

HasMessagePreview returns a boolean if a field has been set.

### GetScheduledAt

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetRecipientCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetRecipientCount() int32`

GetRecipientCount returns the RecipientCount field if non-nil, zero value otherwise.

### GetRecipientCountOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetRecipientCountOk() (*int32, bool)`

GetRecipientCountOk returns a tuple with the RecipientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetRecipientCount(v int32)`

SetRecipientCount sets RecipientCount field to given value.

### HasRecipientCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasRecipientCount() bool`

HasRecipientCount returns a boolean if a field has been set.

### GetSentCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetSentCount() int32`

GetSentCount returns the SentCount field if non-nil, zero value otherwise.

### GetSentCountOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetSentCountOk() (*int32, bool)`

GetSentCountOk returns a tuple with the SentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetSentCount(v int32)`

SetSentCount sets SentCount field to given value.

### HasSentCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasSentCount() bool`

HasSentCount returns a boolean if a field has been set.

### GetDeliveredCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetDeliveredCount() int32`

GetDeliveredCount returns the DeliveredCount field if non-nil, zero value otherwise.

### GetDeliveredCountOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetDeliveredCountOk() (*int32, bool)`

GetDeliveredCountOk returns a tuple with the DeliveredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetDeliveredCount(v int32)`

SetDeliveredCount sets DeliveredCount field to given value.

### HasDeliveredCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasDeliveredCount() bool`

HasDeliveredCount returns a boolean if a field has been set.

### GetReadCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetReadCount() int32`

GetReadCount returns the ReadCount field if non-nil, zero value otherwise.

### GetReadCountOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetReadCountOk() (*int32, bool)`

GetReadCountOk returns a tuple with the ReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetReadCount(v int32)`

SetReadCount sets ReadCount field to given value.

### HasReadCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasReadCount() bool`

HasReadCount returns a boolean if a field has been set.

### GetFailedCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListBroadcasts200ResponseBroadcastsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListBroadcasts200ResponseBroadcastsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListBroadcasts200ResponseBroadcastsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


