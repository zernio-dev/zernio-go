# GetLinkedInMentions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urn** | Pointer to **string** | The LinkedIn URN (person or organization) | [optional] 
**Type** | Pointer to **string** | The type of entity (person or organization) | [optional] 
**DisplayName** | Pointer to **string** | Display name (provided, from API, or derived from vanity URL) | [optional] 
**MentionFormat** | Pointer to **string** | Ready-to-use mention format for post content | [optional] 
**VanityName** | Pointer to **string** | The vanity name/slug (only for organization mentions) | [optional] 
**Warning** | Pointer to **string** | Warning about clickable mentions (only present for person mentions if displayName was not provided) | [optional] 

## Methods

### NewGetLinkedInMentions200Response

`func NewGetLinkedInMentions200Response() *GetLinkedInMentions200Response`

NewGetLinkedInMentions200Response instantiates a new GetLinkedInMentions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinkedInMentions200ResponseWithDefaults

`func NewGetLinkedInMentions200ResponseWithDefaults() *GetLinkedInMentions200Response`

NewGetLinkedInMentions200ResponseWithDefaults instantiates a new GetLinkedInMentions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrn

`func (o *GetLinkedInMentions200Response) GetUrn() string`

GetUrn returns the Urn field if non-nil, zero value otherwise.

### GetUrnOk

`func (o *GetLinkedInMentions200Response) GetUrnOk() (*string, bool)`

GetUrnOk returns a tuple with the Urn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrn

`func (o *GetLinkedInMentions200Response) SetUrn(v string)`

SetUrn sets Urn field to given value.

### HasUrn

`func (o *GetLinkedInMentions200Response) HasUrn() bool`

HasUrn returns a boolean if a field has been set.

### GetType

`func (o *GetLinkedInMentions200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetLinkedInMentions200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetLinkedInMentions200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetLinkedInMentions200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetLinkedInMentions200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetLinkedInMentions200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetLinkedInMentions200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetLinkedInMentions200Response) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMentionFormat

`func (o *GetLinkedInMentions200Response) GetMentionFormat() string`

GetMentionFormat returns the MentionFormat field if non-nil, zero value otherwise.

### GetMentionFormatOk

`func (o *GetLinkedInMentions200Response) GetMentionFormatOk() (*string, bool)`

GetMentionFormatOk returns a tuple with the MentionFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionFormat

`func (o *GetLinkedInMentions200Response) SetMentionFormat(v string)`

SetMentionFormat sets MentionFormat field to given value.

### HasMentionFormat

`func (o *GetLinkedInMentions200Response) HasMentionFormat() bool`

HasMentionFormat returns a boolean if a field has been set.

### GetVanityName

`func (o *GetLinkedInMentions200Response) GetVanityName() string`

GetVanityName returns the VanityName field if non-nil, zero value otherwise.

### GetVanityNameOk

`func (o *GetLinkedInMentions200Response) GetVanityNameOk() (*string, bool)`

GetVanityNameOk returns a tuple with the VanityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVanityName

`func (o *GetLinkedInMentions200Response) SetVanityName(v string)`

SetVanityName sets VanityName field to given value.

### HasVanityName

`func (o *GetLinkedInMentions200Response) HasVanityName() bool`

HasVanityName returns a boolean if a field has been set.

### GetWarning

`func (o *GetLinkedInMentions200Response) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *GetLinkedInMentions200Response) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *GetLinkedInMentions200Response) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *GetLinkedInMentions200Response) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


