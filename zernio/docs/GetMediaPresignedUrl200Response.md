# GetMediaPresignedUrl200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadUrl** | Pointer to **string** | Presigned URL to PUT your file to (expires in 1 hour) | [optional] 
**PublicUrl** | Pointer to **string** | Public URL where the file will be accessible after upload | [optional] 
**Key** | Pointer to **string** | Storage key/path of the file | [optional] 
**ExpiresIn** | Pointer to **int32** | Seconds until the presigned uploadUrl expires (always 3600) | [optional] 

## Methods

### NewGetMediaPresignedUrl200Response

`func NewGetMediaPresignedUrl200Response() *GetMediaPresignedUrl200Response`

NewGetMediaPresignedUrl200Response instantiates a new GetMediaPresignedUrl200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMediaPresignedUrl200ResponseWithDefaults

`func NewGetMediaPresignedUrl200ResponseWithDefaults() *GetMediaPresignedUrl200Response`

NewGetMediaPresignedUrl200ResponseWithDefaults instantiates a new GetMediaPresignedUrl200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadUrl

`func (o *GetMediaPresignedUrl200Response) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *GetMediaPresignedUrl200Response) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *GetMediaPresignedUrl200Response) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *GetMediaPresignedUrl200Response) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetPublicUrl

`func (o *GetMediaPresignedUrl200Response) GetPublicUrl() string`

GetPublicUrl returns the PublicUrl field if non-nil, zero value otherwise.

### GetPublicUrlOk

`func (o *GetMediaPresignedUrl200Response) GetPublicUrlOk() (*string, bool)`

GetPublicUrlOk returns a tuple with the PublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUrl

`func (o *GetMediaPresignedUrl200Response) SetPublicUrl(v string)`

SetPublicUrl sets PublicUrl field to given value.

### HasPublicUrl

`func (o *GetMediaPresignedUrl200Response) HasPublicUrl() bool`

HasPublicUrl returns a boolean if a field has been set.

### GetKey

`func (o *GetMediaPresignedUrl200Response) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetMediaPresignedUrl200Response) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetMediaPresignedUrl200Response) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetMediaPresignedUrl200Response) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetExpiresIn

`func (o *GetMediaPresignedUrl200Response) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *GetMediaPresignedUrl200Response) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *GetMediaPresignedUrl200Response) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *GetMediaPresignedUrl200Response) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


