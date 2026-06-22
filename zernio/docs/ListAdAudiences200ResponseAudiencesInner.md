# ListAdAudiences200ResponseAudiencesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**PlatformAudienceId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Spec** | Pointer to [**NullableTargetingSpec**](TargetingSpec.md) | Present (and the only meaningful payload) when &#x60;type&#x60; is &#x60;saved_targeting&#x60;. Null for uploaded/derived audience types. | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewListAdAudiences200ResponseAudiencesInner

`func NewListAdAudiences200ResponseAudiencesInner() *ListAdAudiences200ResponseAudiencesInner`

NewListAdAudiences200ResponseAudiencesInner instantiates a new ListAdAudiences200ResponseAudiencesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAdAudiences200ResponseAudiencesInnerWithDefaults

`func NewListAdAudiences200ResponseAudiencesInnerWithDefaults() *ListAdAudiences200ResponseAudiencesInner`

NewListAdAudiences200ResponseAudiencesInnerWithDefaults instantiates a new ListAdAudiences200ResponseAudiencesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListAdAudiences200ResponseAudiencesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAdAudiences200ResponseAudiencesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAdAudiences200ResponseAudiencesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListAdAudiences200ResponseAudiencesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ListAdAudiences200ResponseAudiencesInner) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ListAdAudiences200ResponseAudiencesInner) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPlatformAudienceId

`func (o *ListAdAudiences200ResponseAudiencesInner) GetPlatformAudienceId() string`

GetPlatformAudienceId returns the PlatformAudienceId field if non-nil, zero value otherwise.

### GetPlatformAudienceIdOk

`func (o *ListAdAudiences200ResponseAudiencesInner) GetPlatformAudienceIdOk() (*string, bool)`

GetPlatformAudienceIdOk returns a tuple with the PlatformAudienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAudienceId

`func (o *ListAdAudiences200ResponseAudiencesInner) SetPlatformAudienceId(v string)`

SetPlatformAudienceId sets PlatformAudienceId field to given value.

### HasPlatformAudienceId

`func (o *ListAdAudiences200ResponseAudiencesInner) HasPlatformAudienceId() bool`

HasPlatformAudienceId returns a boolean if a field has been set.

### GetName

`func (o *ListAdAudiences200ResponseAudiencesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListAdAudiences200ResponseAudiencesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListAdAudiences200ResponseAudiencesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListAdAudiences200ResponseAudiencesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListAdAudiences200ResponseAudiencesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListAdAudiences200ResponseAudiencesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListAdAudiences200ResponseAudiencesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListAdAudiences200ResponseAudiencesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ListAdAudiences200ResponseAudiencesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListAdAudiences200ResponseAudiencesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListAdAudiences200ResponseAudiencesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListAdAudiences200ResponseAudiencesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSpec

`func (o *ListAdAudiences200ResponseAudiencesInner) GetSpec() TargetingSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ListAdAudiences200ResponseAudiencesInner) GetSpecOk() (*TargetingSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ListAdAudiences200ResponseAudiencesInner) SetSpec(v TargetingSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ListAdAudiences200ResponseAudiencesInner) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### SetSpecNil

`func (o *ListAdAudiences200ResponseAudiencesInner) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *ListAdAudiences200ResponseAudiencesInner) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetPlatform

`func (o *ListAdAudiences200ResponseAudiencesInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListAdAudiences200ResponseAudiencesInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListAdAudiences200ResponseAudiencesInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListAdAudiences200ResponseAudiencesInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSize

`func (o *ListAdAudiences200ResponseAudiencesInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListAdAudiences200ResponseAudiencesInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListAdAudiences200ResponseAudiencesInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ListAdAudiences200ResponseAudiencesInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *ListAdAudiences200ResponseAudiencesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListAdAudiences200ResponseAudiencesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListAdAudiences200ResponseAudiencesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListAdAudiences200ResponseAudiencesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


