# GoogleBusinessPlatformDataEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Event name (displayed as the event heading on Google Search and Maps) | 
**Schedule** | [**GoogleBusinessPlatformDataEventSchedule**](GoogleBusinessPlatformDataEventSchedule.md) |  | 

## Methods

### NewGoogleBusinessPlatformDataEvent

`func NewGoogleBusinessPlatformDataEvent(title string, schedule GoogleBusinessPlatformDataEventSchedule, ) *GoogleBusinessPlatformDataEvent`

NewGoogleBusinessPlatformDataEvent instantiates a new GoogleBusinessPlatformDataEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleBusinessPlatformDataEventWithDefaults

`func NewGoogleBusinessPlatformDataEventWithDefaults() *GoogleBusinessPlatformDataEvent`

NewGoogleBusinessPlatformDataEventWithDefaults instantiates a new GoogleBusinessPlatformDataEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *GoogleBusinessPlatformDataEvent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GoogleBusinessPlatformDataEvent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GoogleBusinessPlatformDataEvent) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSchedule

`func (o *GoogleBusinessPlatformDataEvent) GetSchedule() GoogleBusinessPlatformDataEventSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *GoogleBusinessPlatformDataEvent) GetScheduleOk() (*GoogleBusinessPlatformDataEventSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *GoogleBusinessPlatformDataEvent) SetSchedule(v GoogleBusinessPlatformDataEventSchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


