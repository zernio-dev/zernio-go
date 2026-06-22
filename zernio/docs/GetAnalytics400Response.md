# GetAnalytics400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **map[string]interface{}** | Detailed validation errors | [optional] 

## Methods

### NewGetAnalytics400Response

`func NewGetAnalytics400Response() *GetAnalytics400Response`

NewGetAnalytics400Response instantiates a new GetAnalytics400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalytics400ResponseWithDefaults

`func NewGetAnalytics400ResponseWithDefaults() *GetAnalytics400Response`

NewGetAnalytics400ResponseWithDefaults instantiates a new GetAnalytics400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetAnalytics400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetAnalytics400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetAnalytics400Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetAnalytics400Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetails

`func (o *GetAnalytics400Response) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetAnalytics400Response) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetAnalytics400Response) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetAnalytics400Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


