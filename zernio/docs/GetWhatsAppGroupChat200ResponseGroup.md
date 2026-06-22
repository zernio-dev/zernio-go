# GetWhatsAppGroupChat200ResponseGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**JoinApprovalMode** | Pointer to **string** |  | [optional] 
**Participants** | Pointer to [**[]GetWhatsAppGroupChat200ResponseGroupParticipantsInner**](GetWhatsAppGroupChat200ResponseGroupParticipantsInner.md) |  | [optional] 
**ParticipantCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **int32** | UNIX timestamp | [optional] 
**IsSuspended** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetWhatsAppGroupChat200ResponseGroup

`func NewGetWhatsAppGroupChat200ResponseGroup() *GetWhatsAppGroupChat200ResponseGroup`

NewGetWhatsAppGroupChat200ResponseGroup instantiates a new GetWhatsAppGroupChat200ResponseGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWhatsAppGroupChat200ResponseGroupWithDefaults

`func NewGetWhatsAppGroupChat200ResponseGroupWithDefaults() *GetWhatsAppGroupChat200ResponseGroup`

NewGetWhatsAppGroupChat200ResponseGroupWithDefaults instantiates a new GetWhatsAppGroupChat200ResponseGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetWhatsAppGroupChat200ResponseGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetWhatsAppGroupChat200ResponseGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubject

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetWhatsAppGroupChat200ResponseGroup) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetWhatsAppGroupChat200ResponseGroup) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetDescription

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetWhatsAppGroupChat200ResponseGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetWhatsAppGroupChat200ResponseGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJoinApprovalMode

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetJoinApprovalMode() string`

GetJoinApprovalMode returns the JoinApprovalMode field if non-nil, zero value otherwise.

### GetJoinApprovalModeOk

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetJoinApprovalModeOk() (*string, bool)`

GetJoinApprovalModeOk returns a tuple with the JoinApprovalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinApprovalMode

`func (o *GetWhatsAppGroupChat200ResponseGroup) SetJoinApprovalMode(v string)`

SetJoinApprovalMode sets JoinApprovalMode field to given value.

### HasJoinApprovalMode

`func (o *GetWhatsAppGroupChat200ResponseGroup) HasJoinApprovalMode() bool`

HasJoinApprovalMode returns a boolean if a field has been set.

### GetParticipants

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetParticipants() []GetWhatsAppGroupChat200ResponseGroupParticipantsInner`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetParticipantsOk() (*[]GetWhatsAppGroupChat200ResponseGroupParticipantsInner, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *GetWhatsAppGroupChat200ResponseGroup) SetParticipants(v []GetWhatsAppGroupChat200ResponseGroupParticipantsInner)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *GetWhatsAppGroupChat200ResponseGroup) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetParticipantCount

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetParticipantCount() int32`

GetParticipantCount returns the ParticipantCount field if non-nil, zero value otherwise.

### GetParticipantCountOk

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetParticipantCountOk() (*int32, bool)`

GetParticipantCountOk returns a tuple with the ParticipantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCount

`func (o *GetWhatsAppGroupChat200ResponseGroup) SetParticipantCount(v int32)`

SetParticipantCount sets ParticipantCount field to given value.

### HasParticipantCount

`func (o *GetWhatsAppGroupChat200ResponseGroup) HasParticipantCount() bool`

HasParticipantCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetWhatsAppGroupChat200ResponseGroup) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetWhatsAppGroupChat200ResponseGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIsSuspended

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetIsSuspended() bool`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *GetWhatsAppGroupChat200ResponseGroup) GetIsSuspendedOk() (*bool, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *GetWhatsAppGroupChat200ResponseGroup) SetIsSuspended(v bool)`

SetIsSuspended sets IsSuspended field to given value.

### HasIsSuspended

`func (o *GetWhatsAppGroupChat200ResponseGroup) HasIsSuspended() bool`

HasIsSuspended returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


