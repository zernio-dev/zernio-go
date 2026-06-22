# GetInboxResponseTime200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **NullableString** |  | [optional] 
**Summary** | Pointer to [**GetInboxResponseTime200ResponseSummary**](GetInboxResponseTime200ResponseSummary.md) |  | [optional] 
**Histogram** | Pointer to [**[]GetInboxResponseTime200ResponseHistogramInner**](GetInboxResponseTime200ResponseHistogramInner.md) |  | [optional] 

## Methods

### NewGetInboxResponseTime200Response

`func NewGetInboxResponseTime200Response() *GetInboxResponseTime200Response`

NewGetInboxResponseTime200Response instantiates a new GetInboxResponseTime200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxResponseTime200ResponseWithDefaults

`func NewGetInboxResponseTime200ResponseWithDefaults() *GetInboxResponseTime200Response`

NewGetInboxResponseTime200ResponseWithDefaults instantiates a new GetInboxResponseTime200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetInboxResponseTime200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetInboxResponseTime200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetInboxResponseTime200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetInboxResponseTime200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFrom

`func (o *GetInboxResponseTime200Response) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetInboxResponseTime200Response) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetInboxResponseTime200Response) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetInboxResponseTime200Response) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *GetInboxResponseTime200Response) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetInboxResponseTime200Response) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetInboxResponseTime200Response) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GetInboxResponseTime200Response) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *GetInboxResponseTime200Response) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *GetInboxResponseTime200Response) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetSummary

`func (o *GetInboxResponseTime200Response) GetSummary() GetInboxResponseTime200ResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetInboxResponseTime200Response) GetSummaryOk() (*GetInboxResponseTime200ResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetInboxResponseTime200Response) SetSummary(v GetInboxResponseTime200ResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetInboxResponseTime200Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetHistogram

`func (o *GetInboxResponseTime200Response) GetHistogram() []GetInboxResponseTime200ResponseHistogramInner`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *GetInboxResponseTime200Response) GetHistogramOk() (*[]GetInboxResponseTime200ResponseHistogramInner, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *GetInboxResponseTime200Response) SetHistogram(v []GetInboxResponseTime200ResponseHistogramInner)`

SetHistogram sets Histogram field to given value.

### HasHistogram

`func (o *GetInboxResponseTime200Response) HasHistogram() bool`

HasHistogram returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


