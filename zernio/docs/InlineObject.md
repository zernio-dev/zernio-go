# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Human-readable error message suitable for end-user display. | 
**Code** | **string** | Machine-readable error code. Stable across versions. | 
**Reason** | **string** | Discriminator for which gate fired. | 
**DocumentationUrl** | Pointer to **string** | Link to the relevant documentation page. | [optional] 
**DashboardUrl** | Pointer to **string** | Deep-link to send the end-user to. For &#x60;free_tier_exceeded&#x60; and &#x60;twitter_passthrough&#x60; this is the Zernio billing tab. For &#x60;enterprise_required&#x60; this is the Zernio enterprise contact page.  | [optional] 
**Details** | Pointer to [**InlineObjectDetails**](InlineObjectDetails.md) |  | [optional] 

## Methods

### NewInlineObject

`func NewInlineObject(error_ string, code string, reason string, ) *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *InlineObject) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineObject) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineObject) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *InlineObject) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineObject) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineObject) SetCode(v string)`

SetCode sets Code field to given value.


### GetReason

`func (o *InlineObject) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InlineObject) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InlineObject) SetReason(v string)`

SetReason sets Reason field to given value.


### GetDocumentationUrl

`func (o *InlineObject) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *InlineObject) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *InlineObject) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *InlineObject) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetDashboardUrl

`func (o *InlineObject) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *InlineObject) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *InlineObject) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.

### HasDashboardUrl

`func (o *InlineObject) HasDashboardUrl() bool`

HasDashboardUrl returns a boolean if a field has been set.

### GetDetails

`func (o *InlineObject) GetDetails() InlineObjectDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineObject) GetDetailsOk() (*InlineObjectDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineObject) SetDetails(v InlineObjectDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineObject) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


