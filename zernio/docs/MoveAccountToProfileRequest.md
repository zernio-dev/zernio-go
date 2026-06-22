# MoveAccountToProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | Target profile ID (must be a valid ObjectId and owned by the same user as the account). | 

## Methods

### NewMoveAccountToProfileRequest

`func NewMoveAccountToProfileRequest(profileId string, ) *MoveAccountToProfileRequest`

NewMoveAccountToProfileRequest instantiates a new MoveAccountToProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveAccountToProfileRequestWithDefaults

`func NewMoveAccountToProfileRequestWithDefaults() *MoveAccountToProfileRequest`

NewMoveAccountToProfileRequestWithDefaults instantiates a new MoveAccountToProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *MoveAccountToProfileRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *MoveAccountToProfileRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *MoveAccountToProfileRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


