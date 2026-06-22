# GetYoutubePlaylists200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Playlists** | Pointer to [**[]GetYoutubePlaylists200ResponsePlaylistsInner**](GetYoutubePlaylists200ResponsePlaylistsInner.md) |  | [optional] 
**DefaultPlaylistId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetYoutubePlaylists200Response

`func NewGetYoutubePlaylists200Response() *GetYoutubePlaylists200Response`

NewGetYoutubePlaylists200Response instantiates a new GetYoutubePlaylists200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetYoutubePlaylists200ResponseWithDefaults

`func NewGetYoutubePlaylists200ResponseWithDefaults() *GetYoutubePlaylists200Response`

NewGetYoutubePlaylists200ResponseWithDefaults instantiates a new GetYoutubePlaylists200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaylists

`func (o *GetYoutubePlaylists200Response) GetPlaylists() []GetYoutubePlaylists200ResponsePlaylistsInner`

GetPlaylists returns the Playlists field if non-nil, zero value otherwise.

### GetPlaylistsOk

`func (o *GetYoutubePlaylists200Response) GetPlaylistsOk() (*[]GetYoutubePlaylists200ResponsePlaylistsInner, bool)`

GetPlaylistsOk returns a tuple with the Playlists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylists

`func (o *GetYoutubePlaylists200Response) SetPlaylists(v []GetYoutubePlaylists200ResponsePlaylistsInner)`

SetPlaylists sets Playlists field to given value.

### HasPlaylists

`func (o *GetYoutubePlaylists200Response) HasPlaylists() bool`

HasPlaylists returns a boolean if a field has been set.

### GetDefaultPlaylistId

`func (o *GetYoutubePlaylists200Response) GetDefaultPlaylistId() string`

GetDefaultPlaylistId returns the DefaultPlaylistId field if non-nil, zero value otherwise.

### GetDefaultPlaylistIdOk

`func (o *GetYoutubePlaylists200Response) GetDefaultPlaylistIdOk() (*string, bool)`

GetDefaultPlaylistIdOk returns a tuple with the DefaultPlaylistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlaylistId

`func (o *GetYoutubePlaylists200Response) SetDefaultPlaylistId(v string)`

SetDefaultPlaylistId sets DefaultPlaylistId field to given value.

### HasDefaultPlaylistId

`func (o *GetYoutubePlaylists200Response) HasDefaultPlaylistId() bool`

HasDefaultPlaylistId returns a boolean if a field has been set.

### SetDefaultPlaylistIdNil

`func (o *GetYoutubePlaylists200Response) SetDefaultPlaylistIdNil(b bool)`

 SetDefaultPlaylistIdNil sets the value for DefaultPlaylistId to be an explicit nil

### UnsetDefaultPlaylistId
`func (o *GetYoutubePlaylists200Response) UnsetDefaultPlaylistId()`

UnsetDefaultPlaylistId ensures that no value is present for DefaultPlaylistId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


