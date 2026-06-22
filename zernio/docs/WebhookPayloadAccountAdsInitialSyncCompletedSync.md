# WebhookPayloadAccountAdsInitialSyncCompletedSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Overall outcome of the initial sync. | 
**TotalAds** | **int32** | Total number of ads discovered for backfill. | 
**Synced** | **int32** | Number of ads successfully synced. | 
**Failed** | **int32** | Number of ads that failed to sync. | 
**Error** | Pointer to **string** | Free-form error message from the platform (typically Meta&#39;s Marketing API). Truncated to ~2KB. Present when &#x60;status&#x60; is &#x60;failure&#x60; (and sometimes on &#x60;success&#x60; when discovery saw zero ad accounts). For UX branching prefer &#x60;errorCategory&#x60;; this field is for human display and debugging.  | [optional] 
**ErrorCode** | Pointer to **string** | Platform-native error code if parsed (e.g. Meta &#x60;190&#x60;, &#x60;10&#x60;, &#x60;200&#x60;). | [optional] 
**ErrorSubcode** | Pointer to **string** | Platform-native error subcode if parsed. | [optional] 
**ErrorCategory** | Pointer to **string** | Stable category for UX branching. New values may be added; existing ones are stable. Mapping:   - &#x60;token_invalid&#x60;: access token is expired or revoked. Reconnect.   - &#x60;permission_denied&#x60;: token lacks required scope, or the user has no role     on the Business Manager that owns the ad account. Reconnect with full     permissions, or have an admin grant access.   - &#x60;no_ad_accounts&#x60;: token is valid but sees zero ad accounts. The user     needs to connect a Business Manager that owns ad accounts.   - &#x60;rate_limited&#x60;: platform throttled us. Sync will retry automatically.   - &#x60;discovery_failed&#x60;: any other platform-side failure. Inspect &#x60;error&#x60;.   - &#x60;unknown&#x60;: classifier could not categorize the failure.  | [optional] 

## Methods

### NewWebhookPayloadAccountAdsInitialSyncCompletedSync

`func NewWebhookPayloadAccountAdsInitialSyncCompletedSync(status string, totalAds int32, synced int32, failed int32, ) *WebhookPayloadAccountAdsInitialSyncCompletedSync`

NewWebhookPayloadAccountAdsInitialSyncCompletedSync instantiates a new WebhookPayloadAccountAdsInitialSyncCompletedSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadAccountAdsInitialSyncCompletedSyncWithDefaults

`func NewWebhookPayloadAccountAdsInitialSyncCompletedSyncWithDefaults() *WebhookPayloadAccountAdsInitialSyncCompletedSync`

NewWebhookPayloadAccountAdsInitialSyncCompletedSyncWithDefaults instantiates a new WebhookPayloadAccountAdsInitialSyncCompletedSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalAds

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetTotalAds() int32`

GetTotalAds returns the TotalAds field if non-nil, zero value otherwise.

### GetTotalAdsOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetTotalAdsOk() (*int32, bool)`

GetTotalAdsOk returns a tuple with the TotalAds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAds

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) SetTotalAds(v int32)`

SetTotalAds sets TotalAds field to given value.


### GetSynced

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetSynced() int32`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetSyncedOk() (*int32, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) SetSynced(v int32)`

SetSynced sets Synced field to given value.


### GetFailed

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) SetFailed(v int32)`

SetFailed sets Failed field to given value.


### GetError

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorCode

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorSubcode

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetErrorSubcode() string`

GetErrorSubcode returns the ErrorSubcode field if non-nil, zero value otherwise.

### GetErrorSubcodeOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetErrorSubcodeOk() (*string, bool)`

GetErrorSubcodeOk returns a tuple with the ErrorSubcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSubcode

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) SetErrorSubcode(v string)`

SetErrorSubcode sets ErrorSubcode field to given value.

### HasErrorSubcode

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) HasErrorSubcode() bool`

HasErrorSubcode returns a boolean if a field has been set.

### GetErrorCategory

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetErrorCategory() string`

GetErrorCategory returns the ErrorCategory field if non-nil, zero value otherwise.

### GetErrorCategoryOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) GetErrorCategoryOk() (*string, bool)`

GetErrorCategoryOk returns a tuple with the ErrorCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCategory

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) SetErrorCategory(v string)`

SetErrorCategory sets ErrorCategory field to given value.

### HasErrorCategory

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedSync) HasErrorCategory() bool`

HasErrorCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


