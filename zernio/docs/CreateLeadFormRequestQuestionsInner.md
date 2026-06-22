# CreateLeadFormRequestQuestionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | EMAIL, PHONE, FULL_NAME, FIRST_NAME, LAST_NAME, CUSTOM, … | 
**Key** | Pointer to **string** | CUSTOM questions only. | [optional] 
**Label** | Pointer to **string** | CUSTOM questions only. | [optional] 
**Options** | Pointer to [**[]CreateLeadFormRequestQuestionsInnerOptionsInner**](CreateLeadFormRequestQuestionsInnerOptionsInner.md) |  | [optional] 
**InlineContext** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateLeadFormRequestQuestionsInner

`func NewCreateLeadFormRequestQuestionsInner(type_ string, ) *CreateLeadFormRequestQuestionsInner`

NewCreateLeadFormRequestQuestionsInner instantiates a new CreateLeadFormRequestQuestionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLeadFormRequestQuestionsInnerWithDefaults

`func NewCreateLeadFormRequestQuestionsInnerWithDefaults() *CreateLeadFormRequestQuestionsInner`

NewCreateLeadFormRequestQuestionsInnerWithDefaults instantiates a new CreateLeadFormRequestQuestionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateLeadFormRequestQuestionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateLeadFormRequestQuestionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateLeadFormRequestQuestionsInner) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *CreateLeadFormRequestQuestionsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateLeadFormRequestQuestionsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateLeadFormRequestQuestionsInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateLeadFormRequestQuestionsInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *CreateLeadFormRequestQuestionsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateLeadFormRequestQuestionsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateLeadFormRequestQuestionsInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateLeadFormRequestQuestionsInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOptions

`func (o *CreateLeadFormRequestQuestionsInner) GetOptions() []CreateLeadFormRequestQuestionsInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateLeadFormRequestQuestionsInner) GetOptionsOk() (*[]CreateLeadFormRequestQuestionsInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateLeadFormRequestQuestionsInner) SetOptions(v []CreateLeadFormRequestQuestionsInnerOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateLeadFormRequestQuestionsInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetInlineContext

`func (o *CreateLeadFormRequestQuestionsInner) GetInlineContext() string`

GetInlineContext returns the InlineContext field if non-nil, zero value otherwise.

### GetInlineContextOk

`func (o *CreateLeadFormRequestQuestionsInner) GetInlineContextOk() (*string, bool)`

GetInlineContextOk returns a tuple with the InlineContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineContext

`func (o *CreateLeadFormRequestQuestionsInner) SetInlineContext(v string)`

SetInlineContext sets InlineContext field to given value.

### HasInlineContext

`func (o *CreateLeadFormRequestQuestionsInner) HasInlineContext() bool`

HasInlineContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


