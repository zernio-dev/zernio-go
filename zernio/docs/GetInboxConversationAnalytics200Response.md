# GetInboxConversationAnalytics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ConversationId** | Pointer to **string** | The platformConversationId | [optional] 
**MongoId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **NullableString** |  | [optional] 
**Summary** | Pointer to [**GetInboxConversationAnalytics200ResponseSummary**](GetInboxConversationAnalytics200ResponseSummary.md) |  | [optional] 
**Timeseries** | Pointer to [**[]GetInboxVolume200ResponseTimeseriesInner**](GetInboxVolume200ResponseTimeseriesInner.md) |  | [optional] 
**BySource** | Pointer to [**[]GetInboxConversationAnalytics200ResponseBySourceInner**](GetInboxConversationAnalytics200ResponseBySourceInner.md) |  | [optional] 

## Methods

### NewGetInboxConversationAnalytics200Response

`func NewGetInboxConversationAnalytics200Response() *GetInboxConversationAnalytics200Response`

NewGetInboxConversationAnalytics200Response instantiates a new GetInboxConversationAnalytics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxConversationAnalytics200ResponseWithDefaults

`func NewGetInboxConversationAnalytics200ResponseWithDefaults() *GetInboxConversationAnalytics200Response`

NewGetInboxConversationAnalytics200ResponseWithDefaults instantiates a new GetInboxConversationAnalytics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetInboxConversationAnalytics200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetInboxConversationAnalytics200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetInboxConversationAnalytics200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetInboxConversationAnalytics200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetConversationId

`func (o *GetInboxConversationAnalytics200Response) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *GetInboxConversationAnalytics200Response) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *GetInboxConversationAnalytics200Response) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *GetInboxConversationAnalytics200Response) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetMongoId

`func (o *GetInboxConversationAnalytics200Response) GetMongoId() string`

GetMongoId returns the MongoId field if non-nil, zero value otherwise.

### GetMongoIdOk

`func (o *GetInboxConversationAnalytics200Response) GetMongoIdOk() (*string, bool)`

GetMongoIdOk returns a tuple with the MongoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoId

`func (o *GetInboxConversationAnalytics200Response) SetMongoId(v string)`

SetMongoId sets MongoId field to given value.

### HasMongoId

`func (o *GetInboxConversationAnalytics200Response) HasMongoId() bool`

HasMongoId returns a boolean if a field has been set.

### GetPlatform

`func (o *GetInboxConversationAnalytics200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetInboxConversationAnalytics200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetInboxConversationAnalytics200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetInboxConversationAnalytics200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *GetInboxConversationAnalytics200Response) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *GetInboxConversationAnalytics200Response) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetFrom

`func (o *GetInboxConversationAnalytics200Response) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetInboxConversationAnalytics200Response) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetInboxConversationAnalytics200Response) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetInboxConversationAnalytics200Response) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *GetInboxConversationAnalytics200Response) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetInboxConversationAnalytics200Response) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetInboxConversationAnalytics200Response) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GetInboxConversationAnalytics200Response) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *GetInboxConversationAnalytics200Response) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *GetInboxConversationAnalytics200Response) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetSummary

`func (o *GetInboxConversationAnalytics200Response) GetSummary() GetInboxConversationAnalytics200ResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetInboxConversationAnalytics200Response) GetSummaryOk() (*GetInboxConversationAnalytics200ResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetInboxConversationAnalytics200Response) SetSummary(v GetInboxConversationAnalytics200ResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetInboxConversationAnalytics200Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTimeseries

`func (o *GetInboxConversationAnalytics200Response) GetTimeseries() []GetInboxVolume200ResponseTimeseriesInner`

GetTimeseries returns the Timeseries field if non-nil, zero value otherwise.

### GetTimeseriesOk

`func (o *GetInboxConversationAnalytics200Response) GetTimeseriesOk() (*[]GetInboxVolume200ResponseTimeseriesInner, bool)`

GetTimeseriesOk returns a tuple with the Timeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseries

`func (o *GetInboxConversationAnalytics200Response) SetTimeseries(v []GetInboxVolume200ResponseTimeseriesInner)`

SetTimeseries sets Timeseries field to given value.

### HasTimeseries

`func (o *GetInboxConversationAnalytics200Response) HasTimeseries() bool`

HasTimeseries returns a boolean if a field has been set.

### GetBySource

`func (o *GetInboxConversationAnalytics200Response) GetBySource() []GetInboxConversationAnalytics200ResponseBySourceInner`

GetBySource returns the BySource field if non-nil, zero value otherwise.

### GetBySourceOk

`func (o *GetInboxConversationAnalytics200Response) GetBySourceOk() (*[]GetInboxConversationAnalytics200ResponseBySourceInner, bool)`

GetBySourceOk returns a tuple with the BySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySource

`func (o *GetInboxConversationAnalytics200Response) SetBySource(v []GetInboxConversationAnalytics200ResponseBySourceInner)`

SetBySource sets BySource field to given value.

### HasBySource

`func (o *GetInboxConversationAnalytics200Response) HasBySource() bool`

HasBySource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


