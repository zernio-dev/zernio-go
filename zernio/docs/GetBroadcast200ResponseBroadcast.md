# GetBroadcast200ResponseBroadcast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Message** | Pointer to [**SendInboxMessageRequestInteractiveFooter**](SendInboxMessageRequestInteractiveFooter.md) |  | [optional] 
**Template** | Pointer to [**GetWhatsAppPhoneNumbers200ResponseSandboxTemplate**](GetWhatsAppPhoneNumbers200ResponseSandboxTemplate.md) |  | [optional] 
**SegmentFilters** | Pointer to [**ListContacts200ResponseFilters**](ListContacts200ResponseFilters.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ScheduledAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**RecipientCount** | Pointer to **int32** |  | [optional] 
**SentCount** | Pointer to **int32** |  | [optional] 
**DeliveredCount** | Pointer to **int32** |  | [optional] 
**ReadCount** | Pointer to **int32** |  | [optional] 
**FailedCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetBroadcast200ResponseBroadcast

`func NewGetBroadcast200ResponseBroadcast() *GetBroadcast200ResponseBroadcast`

NewGetBroadcast200ResponseBroadcast instantiates a new GetBroadcast200ResponseBroadcast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBroadcast200ResponseBroadcastWithDefaults

`func NewGetBroadcast200ResponseBroadcastWithDefaults() *GetBroadcast200ResponseBroadcast`

NewGetBroadcast200ResponseBroadcastWithDefaults instantiates a new GetBroadcast200ResponseBroadcast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetBroadcast200ResponseBroadcast) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBroadcast200ResponseBroadcast) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBroadcast200ResponseBroadcast) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetBroadcast200ResponseBroadcast) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetBroadcast200ResponseBroadcast) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetBroadcast200ResponseBroadcast) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetBroadcast200ResponseBroadcast) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetBroadcast200ResponseBroadcast) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetBroadcast200ResponseBroadcast) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetBroadcast200ResponseBroadcast) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetBroadcast200ResponseBroadcast) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetBroadcast200ResponseBroadcast) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatform

`func (o *GetBroadcast200ResponseBroadcast) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetBroadcast200ResponseBroadcast) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetBroadcast200ResponseBroadcast) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetBroadcast200ResponseBroadcast) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountId

`func (o *GetBroadcast200ResponseBroadcast) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetBroadcast200ResponseBroadcast) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetBroadcast200ResponseBroadcast) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetBroadcast200ResponseBroadcast) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetMessage

`func (o *GetBroadcast200ResponseBroadcast) GetMessage() SendInboxMessageRequestInteractiveFooter`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetBroadcast200ResponseBroadcast) GetMessageOk() (*SendInboxMessageRequestInteractiveFooter, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetBroadcast200ResponseBroadcast) SetMessage(v SendInboxMessageRequestInteractiveFooter)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetBroadcast200ResponseBroadcast) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTemplate

`func (o *GetBroadcast200ResponseBroadcast) GetTemplate() GetWhatsAppPhoneNumbers200ResponseSandboxTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetBroadcast200ResponseBroadcast) GetTemplateOk() (*GetWhatsAppPhoneNumbers200ResponseSandboxTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetBroadcast200ResponseBroadcast) SetTemplate(v GetWhatsAppPhoneNumbers200ResponseSandboxTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetBroadcast200ResponseBroadcast) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetSegmentFilters

`func (o *GetBroadcast200ResponseBroadcast) GetSegmentFilters() ListContacts200ResponseFilters`

GetSegmentFilters returns the SegmentFilters field if non-nil, zero value otherwise.

### GetSegmentFiltersOk

`func (o *GetBroadcast200ResponseBroadcast) GetSegmentFiltersOk() (*ListContacts200ResponseFilters, bool)`

GetSegmentFiltersOk returns a tuple with the SegmentFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentFilters

`func (o *GetBroadcast200ResponseBroadcast) SetSegmentFilters(v ListContacts200ResponseFilters)`

SetSegmentFilters sets SegmentFilters field to given value.

### HasSegmentFilters

`func (o *GetBroadcast200ResponseBroadcast) HasSegmentFilters() bool`

HasSegmentFilters returns a boolean if a field has been set.

### GetStatus

`func (o *GetBroadcast200ResponseBroadcast) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBroadcast200ResponseBroadcast) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBroadcast200ResponseBroadcast) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetBroadcast200ResponseBroadcast) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScheduledAt

`func (o *GetBroadcast200ResponseBroadcast) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *GetBroadcast200ResponseBroadcast) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *GetBroadcast200ResponseBroadcast) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *GetBroadcast200ResponseBroadcast) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *GetBroadcast200ResponseBroadcast) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetBroadcast200ResponseBroadcast) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetBroadcast200ResponseBroadcast) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GetBroadcast200ResponseBroadcast) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GetBroadcast200ResponseBroadcast) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetBroadcast200ResponseBroadcast) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetBroadcast200ResponseBroadcast) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetBroadcast200ResponseBroadcast) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetRecipientCount

`func (o *GetBroadcast200ResponseBroadcast) GetRecipientCount() int32`

GetRecipientCount returns the RecipientCount field if non-nil, zero value otherwise.

### GetRecipientCountOk

`func (o *GetBroadcast200ResponseBroadcast) GetRecipientCountOk() (*int32, bool)`

GetRecipientCountOk returns a tuple with the RecipientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientCount

`func (o *GetBroadcast200ResponseBroadcast) SetRecipientCount(v int32)`

SetRecipientCount sets RecipientCount field to given value.

### HasRecipientCount

`func (o *GetBroadcast200ResponseBroadcast) HasRecipientCount() bool`

HasRecipientCount returns a boolean if a field has been set.

### GetSentCount

`func (o *GetBroadcast200ResponseBroadcast) GetSentCount() int32`

GetSentCount returns the SentCount field if non-nil, zero value otherwise.

### GetSentCountOk

`func (o *GetBroadcast200ResponseBroadcast) GetSentCountOk() (*int32, bool)`

GetSentCountOk returns a tuple with the SentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentCount

`func (o *GetBroadcast200ResponseBroadcast) SetSentCount(v int32)`

SetSentCount sets SentCount field to given value.

### HasSentCount

`func (o *GetBroadcast200ResponseBroadcast) HasSentCount() bool`

HasSentCount returns a boolean if a field has been set.

### GetDeliveredCount

`func (o *GetBroadcast200ResponseBroadcast) GetDeliveredCount() int32`

GetDeliveredCount returns the DeliveredCount field if non-nil, zero value otherwise.

### GetDeliveredCountOk

`func (o *GetBroadcast200ResponseBroadcast) GetDeliveredCountOk() (*int32, bool)`

GetDeliveredCountOk returns a tuple with the DeliveredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredCount

`func (o *GetBroadcast200ResponseBroadcast) SetDeliveredCount(v int32)`

SetDeliveredCount sets DeliveredCount field to given value.

### HasDeliveredCount

`func (o *GetBroadcast200ResponseBroadcast) HasDeliveredCount() bool`

HasDeliveredCount returns a boolean if a field has been set.

### GetReadCount

`func (o *GetBroadcast200ResponseBroadcast) GetReadCount() int32`

GetReadCount returns the ReadCount field if non-nil, zero value otherwise.

### GetReadCountOk

`func (o *GetBroadcast200ResponseBroadcast) GetReadCountOk() (*int32, bool)`

GetReadCountOk returns a tuple with the ReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCount

`func (o *GetBroadcast200ResponseBroadcast) SetReadCount(v int32)`

SetReadCount sets ReadCount field to given value.

### HasReadCount

`func (o *GetBroadcast200ResponseBroadcast) HasReadCount() bool`

HasReadCount returns a boolean if a field has been set.

### GetFailedCount

`func (o *GetBroadcast200ResponseBroadcast) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *GetBroadcast200ResponseBroadcast) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *GetBroadcast200ResponseBroadcast) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *GetBroadcast200ResponseBroadcast) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetBroadcast200ResponseBroadcast) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetBroadcast200ResponseBroadcast) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetBroadcast200ResponseBroadcast) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetBroadcast200ResponseBroadcast) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetBroadcast200ResponseBroadcast) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetBroadcast200ResponseBroadcast) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetBroadcast200ResponseBroadcast) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetBroadcast200ResponseBroadcast) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


