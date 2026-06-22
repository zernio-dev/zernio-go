# GetMediaPresignedUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | Name of the file to upload | 
**ContentType** | **string** | MIME type of the file | 
**Size** | Pointer to **int32** | Optional file size in bytes for pre-validation (max 5GB) | [optional] 

## Methods

### NewGetMediaPresignedUrlRequest

`func NewGetMediaPresignedUrlRequest(filename string, contentType string, ) *GetMediaPresignedUrlRequest`

NewGetMediaPresignedUrlRequest instantiates a new GetMediaPresignedUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMediaPresignedUrlRequestWithDefaults

`func NewGetMediaPresignedUrlRequestWithDefaults() *GetMediaPresignedUrlRequest`

NewGetMediaPresignedUrlRequestWithDefaults instantiates a new GetMediaPresignedUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *GetMediaPresignedUrlRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *GetMediaPresignedUrlRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *GetMediaPresignedUrlRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *GetMediaPresignedUrlRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GetMediaPresignedUrlRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GetMediaPresignedUrlRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetSize

`func (o *GetMediaPresignedUrlRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetMediaPresignedUrlRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetMediaPresignedUrlRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetMediaPresignedUrlRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


