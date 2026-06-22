# PlatformTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** | Supported values: twitter, threads, instagram, youtube, facebook, linkedin, pinterest, reddit, tiktok, bluesky, googlebusiness, telegram | [optional] 
**AccountId** | Pointer to [**PlatformTargetAccountId**](PlatformTargetAccountId.md) |  | [optional] 
**CustomContent** | Pointer to **string** | Platform-specific text override. When set, this content is used instead of the top-level post content for this platform. Useful for tailoring captions per platform (e.g. keeping tweets under 280 characters). | [optional] 
**CustomMedia** | Pointer to [**[]MediaItem**](MediaItem.md) |  | [optional] 
**ScheduledFor** | Pointer to **time.Time** | Optional per-platform scheduled time override (uses post.scheduledFor when omitted) | [optional] 
**PlatformSpecificData** | Pointer to [**PlatformTargetPlatformSpecificData**](PlatformTargetPlatformSpecificData.md) |  | [optional] 
**Status** | Pointer to **string** | Platform-specific status: pending, publishing, published, failed | [optional] 
**PlatformPostId** | Pointer to **string** | The native post ID on the platform (populated after successful publish) | [optional] 
**PlatformPostUrl** | Pointer to **string** | Public URL of the published post. Included in the response for immediate posts; for scheduled posts, fetch via GET /v1/posts/{postId} after publish time. | [optional] 
**PublishedAt** | Pointer to **time.Time** | Timestamp when the post was published to this platform | [optional] 
**ErrorMessage** | Pointer to **string** | Human-readable error message when status is failed. Contains platform-specific error details explaining why the publish failed. | [optional] 
**ErrorCategory** | Pointer to **string** | Error category for programmatic handling: auth_expired (token expired/revoked), user_content (wrong format/too long), user_abuse (rate limits/spam), account_issue (config problems), platform_rejected (policy violation), platform_error (5xx/maintenance), system_error (Zernio infra), unknown | [optional] 
**ErrorSource** | Pointer to **string** | Who caused the error: user (fix content/reconnect), platform (outage/API change), system (Zernio issue, rare) | [optional] 

## Methods

### NewPlatformTarget

`func NewPlatformTarget() *PlatformTarget`

NewPlatformTarget instantiates a new PlatformTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformTargetWithDefaults

`func NewPlatformTargetWithDefaults() *PlatformTarget`

NewPlatformTargetWithDefaults instantiates a new PlatformTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *PlatformTarget) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PlatformTarget) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PlatformTarget) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PlatformTarget) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountId

`func (o *PlatformTarget) GetAccountId() PlatformTargetAccountId`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PlatformTarget) GetAccountIdOk() (*PlatformTargetAccountId, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PlatformTarget) SetAccountId(v PlatformTargetAccountId)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *PlatformTarget) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCustomContent

`func (o *PlatformTarget) GetCustomContent() string`

GetCustomContent returns the CustomContent field if non-nil, zero value otherwise.

### GetCustomContentOk

`func (o *PlatformTarget) GetCustomContentOk() (*string, bool)`

GetCustomContentOk returns a tuple with the CustomContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContent

`func (o *PlatformTarget) SetCustomContent(v string)`

SetCustomContent sets CustomContent field to given value.

### HasCustomContent

`func (o *PlatformTarget) HasCustomContent() bool`

HasCustomContent returns a boolean if a field has been set.

### GetCustomMedia

`func (o *PlatformTarget) GetCustomMedia() []MediaItem`

GetCustomMedia returns the CustomMedia field if non-nil, zero value otherwise.

### GetCustomMediaOk

`func (o *PlatformTarget) GetCustomMediaOk() (*[]MediaItem, bool)`

GetCustomMediaOk returns a tuple with the CustomMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMedia

`func (o *PlatformTarget) SetCustomMedia(v []MediaItem)`

SetCustomMedia sets CustomMedia field to given value.

### HasCustomMedia

`func (o *PlatformTarget) HasCustomMedia() bool`

HasCustomMedia returns a boolean if a field has been set.

### GetScheduledFor

`func (o *PlatformTarget) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *PlatformTarget) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *PlatformTarget) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *PlatformTarget) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetPlatformSpecificData

`func (o *PlatformTarget) GetPlatformSpecificData() PlatformTargetPlatformSpecificData`

GetPlatformSpecificData returns the PlatformSpecificData field if non-nil, zero value otherwise.

### GetPlatformSpecificDataOk

`func (o *PlatformTarget) GetPlatformSpecificDataOk() (*PlatformTargetPlatformSpecificData, bool)`

GetPlatformSpecificDataOk returns a tuple with the PlatformSpecificData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformSpecificData

`func (o *PlatformTarget) SetPlatformSpecificData(v PlatformTargetPlatformSpecificData)`

SetPlatformSpecificData sets PlatformSpecificData field to given value.

### HasPlatformSpecificData

`func (o *PlatformTarget) HasPlatformSpecificData() bool`

HasPlatformSpecificData returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformTarget) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformTarget) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformTarget) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformTarget) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPlatformPostId

`func (o *PlatformTarget) GetPlatformPostId() string`

GetPlatformPostId returns the PlatformPostId field if non-nil, zero value otherwise.

### GetPlatformPostIdOk

`func (o *PlatformTarget) GetPlatformPostIdOk() (*string, bool)`

GetPlatformPostIdOk returns a tuple with the PlatformPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostId

`func (o *PlatformTarget) SetPlatformPostId(v string)`

SetPlatformPostId sets PlatformPostId field to given value.

### HasPlatformPostId

`func (o *PlatformTarget) HasPlatformPostId() bool`

HasPlatformPostId returns a boolean if a field has been set.

### GetPlatformPostUrl

`func (o *PlatformTarget) GetPlatformPostUrl() string`

GetPlatformPostUrl returns the PlatformPostUrl field if non-nil, zero value otherwise.

### GetPlatformPostUrlOk

`func (o *PlatformTarget) GetPlatformPostUrlOk() (*string, bool)`

GetPlatformPostUrlOk returns a tuple with the PlatformPostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostUrl

`func (o *PlatformTarget) SetPlatformPostUrl(v string)`

SetPlatformPostUrl sets PlatformPostUrl field to given value.

### HasPlatformPostUrl

`func (o *PlatformTarget) HasPlatformPostUrl() bool`

HasPlatformPostUrl returns a boolean if a field has been set.

### GetPublishedAt

`func (o *PlatformTarget) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *PlatformTarget) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *PlatformTarget) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *PlatformTarget) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PlatformTarget) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PlatformTarget) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PlatformTarget) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PlatformTarget) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorCategory

`func (o *PlatformTarget) GetErrorCategory() string`

GetErrorCategory returns the ErrorCategory field if non-nil, zero value otherwise.

### GetErrorCategoryOk

`func (o *PlatformTarget) GetErrorCategoryOk() (*string, bool)`

GetErrorCategoryOk returns a tuple with the ErrorCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCategory

`func (o *PlatformTarget) SetErrorCategory(v string)`

SetErrorCategory sets ErrorCategory field to given value.

### HasErrorCategory

`func (o *PlatformTarget) HasErrorCategory() bool`

HasErrorCategory returns a boolean if a field has been set.

### GetErrorSource

`func (o *PlatformTarget) GetErrorSource() string`

GetErrorSource returns the ErrorSource field if non-nil, zero value otherwise.

### GetErrorSourceOk

`func (o *PlatformTarget) GetErrorSourceOk() (*string, bool)`

GetErrorSourceOk returns a tuple with the ErrorSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSource

`func (o *PlatformTarget) SetErrorSource(v string)`

SetErrorSource sets ErrorSource field to given value.

### HasErrorSource

`func (o *PlatformTarget) HasErrorSource() bool`

HasErrorSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


