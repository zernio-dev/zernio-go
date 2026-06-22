# CreateStandaloneAdRequestPlacements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublisherPlatforms** | Pointer to **[]string** | Top-level platforms to deliver on. A position field below is only honoured when its parent platform is included here. | [optional] 
**FacebookPositions** | Pointer to **[]string** |  | [optional] 
**InstagramPositions** | Pointer to **[]string** |  | [optional] 
**MessengerPositions** | Pointer to **[]string** |  | [optional] 
**AudienceNetworkPositions** | Pointer to **[]string** |  | [optional] 
**ThreadsPositions** | Pointer to **[]string** |  | [optional] 
**WhatsappPositions** | Pointer to **[]string** |  | [optional] 
**DevicePlatforms** | Pointer to **[]string** | Restrict by device. Omit to deliver on both mobile and desktop. | [optional] 

## Methods

### NewCreateStandaloneAdRequestPlacements

`func NewCreateStandaloneAdRequestPlacements() *CreateStandaloneAdRequestPlacements`

NewCreateStandaloneAdRequestPlacements instantiates a new CreateStandaloneAdRequestPlacements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandaloneAdRequestPlacementsWithDefaults

`func NewCreateStandaloneAdRequestPlacementsWithDefaults() *CreateStandaloneAdRequestPlacements`

NewCreateStandaloneAdRequestPlacementsWithDefaults instantiates a new CreateStandaloneAdRequestPlacements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublisherPlatforms

`func (o *CreateStandaloneAdRequestPlacements) GetPublisherPlatforms() []string`

GetPublisherPlatforms returns the PublisherPlatforms field if non-nil, zero value otherwise.

### GetPublisherPlatformsOk

`func (o *CreateStandaloneAdRequestPlacements) GetPublisherPlatformsOk() (*[]string, bool)`

GetPublisherPlatformsOk returns a tuple with the PublisherPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherPlatforms

`func (o *CreateStandaloneAdRequestPlacements) SetPublisherPlatforms(v []string)`

SetPublisherPlatforms sets PublisherPlatforms field to given value.

### HasPublisherPlatforms

`func (o *CreateStandaloneAdRequestPlacements) HasPublisherPlatforms() bool`

HasPublisherPlatforms returns a boolean if a field has been set.

### GetFacebookPositions

`func (o *CreateStandaloneAdRequestPlacements) GetFacebookPositions() []string`

GetFacebookPositions returns the FacebookPositions field if non-nil, zero value otherwise.

### GetFacebookPositionsOk

`func (o *CreateStandaloneAdRequestPlacements) GetFacebookPositionsOk() (*[]string, bool)`

GetFacebookPositionsOk returns a tuple with the FacebookPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookPositions

`func (o *CreateStandaloneAdRequestPlacements) SetFacebookPositions(v []string)`

SetFacebookPositions sets FacebookPositions field to given value.

### HasFacebookPositions

`func (o *CreateStandaloneAdRequestPlacements) HasFacebookPositions() bool`

HasFacebookPositions returns a boolean if a field has been set.

### GetInstagramPositions

`func (o *CreateStandaloneAdRequestPlacements) GetInstagramPositions() []string`

GetInstagramPositions returns the InstagramPositions field if non-nil, zero value otherwise.

### GetInstagramPositionsOk

`func (o *CreateStandaloneAdRequestPlacements) GetInstagramPositionsOk() (*[]string, bool)`

GetInstagramPositionsOk returns a tuple with the InstagramPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramPositions

`func (o *CreateStandaloneAdRequestPlacements) SetInstagramPositions(v []string)`

SetInstagramPositions sets InstagramPositions field to given value.

### HasInstagramPositions

`func (o *CreateStandaloneAdRequestPlacements) HasInstagramPositions() bool`

HasInstagramPositions returns a boolean if a field has been set.

### GetMessengerPositions

`func (o *CreateStandaloneAdRequestPlacements) GetMessengerPositions() []string`

GetMessengerPositions returns the MessengerPositions field if non-nil, zero value otherwise.

### GetMessengerPositionsOk

`func (o *CreateStandaloneAdRequestPlacements) GetMessengerPositionsOk() (*[]string, bool)`

GetMessengerPositionsOk returns a tuple with the MessengerPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessengerPositions

`func (o *CreateStandaloneAdRequestPlacements) SetMessengerPositions(v []string)`

SetMessengerPositions sets MessengerPositions field to given value.

### HasMessengerPositions

`func (o *CreateStandaloneAdRequestPlacements) HasMessengerPositions() bool`

HasMessengerPositions returns a boolean if a field has been set.

### GetAudienceNetworkPositions

`func (o *CreateStandaloneAdRequestPlacements) GetAudienceNetworkPositions() []string`

GetAudienceNetworkPositions returns the AudienceNetworkPositions field if non-nil, zero value otherwise.

### GetAudienceNetworkPositionsOk

`func (o *CreateStandaloneAdRequestPlacements) GetAudienceNetworkPositionsOk() (*[]string, bool)`

GetAudienceNetworkPositionsOk returns a tuple with the AudienceNetworkPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceNetworkPositions

`func (o *CreateStandaloneAdRequestPlacements) SetAudienceNetworkPositions(v []string)`

SetAudienceNetworkPositions sets AudienceNetworkPositions field to given value.

### HasAudienceNetworkPositions

`func (o *CreateStandaloneAdRequestPlacements) HasAudienceNetworkPositions() bool`

HasAudienceNetworkPositions returns a boolean if a field has been set.

### GetThreadsPositions

`func (o *CreateStandaloneAdRequestPlacements) GetThreadsPositions() []string`

GetThreadsPositions returns the ThreadsPositions field if non-nil, zero value otherwise.

### GetThreadsPositionsOk

`func (o *CreateStandaloneAdRequestPlacements) GetThreadsPositionsOk() (*[]string, bool)`

GetThreadsPositionsOk returns a tuple with the ThreadsPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPositions

`func (o *CreateStandaloneAdRequestPlacements) SetThreadsPositions(v []string)`

SetThreadsPositions sets ThreadsPositions field to given value.

### HasThreadsPositions

`func (o *CreateStandaloneAdRequestPlacements) HasThreadsPositions() bool`

HasThreadsPositions returns a boolean if a field has been set.

### GetWhatsappPositions

`func (o *CreateStandaloneAdRequestPlacements) GetWhatsappPositions() []string`

GetWhatsappPositions returns the WhatsappPositions field if non-nil, zero value otherwise.

### GetWhatsappPositionsOk

`func (o *CreateStandaloneAdRequestPlacements) GetWhatsappPositionsOk() (*[]string, bool)`

GetWhatsappPositionsOk returns a tuple with the WhatsappPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsappPositions

`func (o *CreateStandaloneAdRequestPlacements) SetWhatsappPositions(v []string)`

SetWhatsappPositions sets WhatsappPositions field to given value.

### HasWhatsappPositions

`func (o *CreateStandaloneAdRequestPlacements) HasWhatsappPositions() bool`

HasWhatsappPositions returns a boolean if a field has been set.

### GetDevicePlatforms

`func (o *CreateStandaloneAdRequestPlacements) GetDevicePlatforms() []string`

GetDevicePlatforms returns the DevicePlatforms field if non-nil, zero value otherwise.

### GetDevicePlatformsOk

`func (o *CreateStandaloneAdRequestPlacements) GetDevicePlatformsOk() (*[]string, bool)`

GetDevicePlatformsOk returns a tuple with the DevicePlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePlatforms

`func (o *CreateStandaloneAdRequestPlacements) SetDevicePlatforms(v []string)`

SetDevicePlatforms sets DevicePlatforms field to given value.

### HasDevicePlatforms

`func (o *CreateStandaloneAdRequestPlacements) HasDevicePlatforms() bool`

HasDevicePlatforms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


