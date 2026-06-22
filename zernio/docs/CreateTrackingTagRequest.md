# CreateTrackingTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdAccountId** | **string** | Meta ad account id, e.g. &#x60;act_123456789&#x60;. | 
**Name** | **string** |  | 

## Methods

### NewCreateTrackingTagRequest

`func NewCreateTrackingTagRequest(adAccountId string, name string, ) *CreateTrackingTagRequest`

NewCreateTrackingTagRequest instantiates a new CreateTrackingTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTrackingTagRequestWithDefaults

`func NewCreateTrackingTagRequestWithDefaults() *CreateTrackingTagRequest`

NewCreateTrackingTagRequestWithDefaults instantiates a new CreateTrackingTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdAccountId

`func (o *CreateTrackingTagRequest) GetAdAccountId() string`

GetAdAccountId returns the AdAccountId field if non-nil, zero value otherwise.

### GetAdAccountIdOk

`func (o *CreateTrackingTagRequest) GetAdAccountIdOk() (*string, bool)`

GetAdAccountIdOk returns a tuple with the AdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountId

`func (o *CreateTrackingTagRequest) SetAdAccountId(v string)`

SetAdAccountId sets AdAccountId field to given value.


### GetName

`func (o *CreateTrackingTagRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTrackingTagRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTrackingTagRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


