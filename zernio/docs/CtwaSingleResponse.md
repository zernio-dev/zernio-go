# CtwaSingleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdType** | **string** |  | 
**Ad** | **map[string]interface{}** | The persisted Ad document. | 
**Message** | **string** |  | 

## Methods

### NewCtwaSingleResponse

`func NewCtwaSingleResponse(adType string, ad map[string]interface{}, message string, ) *CtwaSingleResponse`

NewCtwaSingleResponse instantiates a new CtwaSingleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCtwaSingleResponseWithDefaults

`func NewCtwaSingleResponseWithDefaults() *CtwaSingleResponse`

NewCtwaSingleResponseWithDefaults instantiates a new CtwaSingleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdType

`func (o *CtwaSingleResponse) GetAdType() string`

GetAdType returns the AdType field if non-nil, zero value otherwise.

### GetAdTypeOk

`func (o *CtwaSingleResponse) GetAdTypeOk() (*string, bool)`

GetAdTypeOk returns a tuple with the AdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdType

`func (o *CtwaSingleResponse) SetAdType(v string)`

SetAdType sets AdType field to given value.


### GetAd

`func (o *CtwaSingleResponse) GetAd() map[string]interface{}`

GetAd returns the Ad field if non-nil, zero value otherwise.

### GetAdOk

`func (o *CtwaSingleResponse) GetAdOk() (*map[string]interface{}, bool)`

GetAdOk returns a tuple with the Ad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAd

`func (o *CtwaSingleResponse) SetAd(v map[string]interface{})`

SetAd sets Ad field to given value.


### GetMessage

`func (o *CtwaSingleResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CtwaSingleResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CtwaSingleResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


