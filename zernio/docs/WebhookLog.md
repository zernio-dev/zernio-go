# WebhookLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | ID of the account owner the webhook belongs to | [optional] 
**WebhookId** | Pointer to **string** | ID of the webhook configuration that produced this delivery | [optional] 
**WebhookName** | Pointer to **string** | Name of the webhook configuration at delivery time | [optional] 
**EventId** | Pointer to **string** | Stable webhook event ID (correlates to the delivered payload) | [optional] 
**Event** | Pointer to **string** | Event type that triggered the delivery (e.g. post.published) | [optional] 
**Url** | Pointer to **string** | Destination URL the webhook was delivered to | [optional] 
**Status** | Pointer to **string** | Delivery outcome | [optional] 
**StatusCode** | Pointer to **int32** | HTTP status code returned by the destination endpoint | [optional] 
**RequestPayload** | Pointer to **map[string]interface{}** | The JSON payload sent to the destination endpoint | [optional] 
**ResponseBody** | Pointer to **string** | Response body returned by the destination endpoint | [optional] 
**ErrorMessage** | Pointer to **string** | Error message when delivery failed | [optional] 
**AttemptNumber** | Pointer to **int32** | Delivery attempt number (increments on retries) | [optional] 
**ResponseTime** | Pointer to **int32** | Time taken by the destination endpoint to respond, in milliseconds | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp the delivery was attempted | [optional] 

## Methods

### NewWebhookLog

`func NewWebhookLog() *WebhookLog`

NewWebhookLog instantiates a new WebhookLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookLogWithDefaults

`func NewWebhookLogWithDefaults() *WebhookLog`

NewWebhookLogWithDefaults instantiates a new WebhookLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *WebhookLog) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebhookLog) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebhookLog) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WebhookLog) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWebhookId

`func (o *WebhookLog) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookLog) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookLog) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *WebhookLog) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetWebhookName

`func (o *WebhookLog) GetWebhookName() string`

GetWebhookName returns the WebhookName field if non-nil, zero value otherwise.

### GetWebhookNameOk

`func (o *WebhookLog) GetWebhookNameOk() (*string, bool)`

GetWebhookNameOk returns a tuple with the WebhookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookName

`func (o *WebhookLog) SetWebhookName(v string)`

SetWebhookName sets WebhookName field to given value.

### HasWebhookName

`func (o *WebhookLog) HasWebhookName() bool`

HasWebhookName returns a boolean if a field has been set.

### GetEventId

`func (o *WebhookLog) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WebhookLog) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WebhookLog) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *WebhookLog) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEvent

`func (o *WebhookLog) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookLog) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookLog) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WebhookLog) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookLog) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookLog) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookLog) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookLog) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookLog) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookLog) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookLog) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *WebhookLog) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *WebhookLog) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *WebhookLog) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *WebhookLog) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetRequestPayload

`func (o *WebhookLog) GetRequestPayload() map[string]interface{}`

GetRequestPayload returns the RequestPayload field if non-nil, zero value otherwise.

### GetRequestPayloadOk

`func (o *WebhookLog) GetRequestPayloadOk() (*map[string]interface{}, bool)`

GetRequestPayloadOk returns a tuple with the RequestPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPayload

`func (o *WebhookLog) SetRequestPayload(v map[string]interface{})`

SetRequestPayload sets RequestPayload field to given value.

### HasRequestPayload

`func (o *WebhookLog) HasRequestPayload() bool`

HasRequestPayload returns a boolean if a field has been set.

### GetResponseBody

`func (o *WebhookLog) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *WebhookLog) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *WebhookLog) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *WebhookLog) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetErrorMessage

`func (o *WebhookLog) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *WebhookLog) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *WebhookLog) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *WebhookLog) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetAttemptNumber

`func (o *WebhookLog) GetAttemptNumber() int32`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *WebhookLog) GetAttemptNumberOk() (*int32, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *WebhookLog) SetAttemptNumber(v int32)`

SetAttemptNumber sets AttemptNumber field to given value.

### HasAttemptNumber

`func (o *WebhookLog) HasAttemptNumber() bool`

HasAttemptNumber returns a boolean if a field has been set.

### GetResponseTime

`func (o *WebhookLog) GetResponseTime() int32`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *WebhookLog) GetResponseTimeOk() (*int32, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *WebhookLog) SetResponseTime(v int32)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *WebhookLog) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebhookLog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


