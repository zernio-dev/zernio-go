# TwitterPlatformDataPoll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | **[]string** | Poll options (2-4 choices, max 25 characters each) | 
**DurationMinutes** | **int32** | Poll duration in minutes (5 min to 7 days) | 

## Methods

### NewTwitterPlatformDataPoll

`func NewTwitterPlatformDataPoll(options []string, durationMinutes int32, ) *TwitterPlatformDataPoll`

NewTwitterPlatformDataPoll instantiates a new TwitterPlatformDataPoll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwitterPlatformDataPollWithDefaults

`func NewTwitterPlatformDataPollWithDefaults() *TwitterPlatformDataPoll`

NewTwitterPlatformDataPollWithDefaults instantiates a new TwitterPlatformDataPoll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *TwitterPlatformDataPoll) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *TwitterPlatformDataPoll) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *TwitterPlatformDataPoll) SetOptions(v []string)`

SetOptions sets Options field to given value.


### GetDurationMinutes

`func (o *TwitterPlatformDataPoll) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *TwitterPlatformDataPoll) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *TwitterPlatformDataPoll) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


