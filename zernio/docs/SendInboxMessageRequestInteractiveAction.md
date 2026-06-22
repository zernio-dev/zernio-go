# SendInboxMessageRequestInteractiveAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Button** | **string** | CTA label that opens the list (max ~20 chars). | 
**Sections** | [**[]SendInboxMessageRequestInteractiveActionOneOfSectionsInner**](SendInboxMessageRequestInteractiveActionOneOfSectionsInner.md) | 1-10 sections. Total rows across all sections cannot exceed 10. | 
**Name** | **string** |  | 
**Parameters** | [**SendInboxMessageRequestInteractiveActionOneOf3Parameters**](SendInboxMessageRequestInteractiveActionOneOf3Parameters.md) |  | 

## Methods

### NewSendInboxMessageRequestInteractiveAction

`func NewSendInboxMessageRequestInteractiveAction(button string, sections []SendInboxMessageRequestInteractiveActionOneOfSectionsInner, name string, parameters SendInboxMessageRequestInteractiveActionOneOf3Parameters, ) *SendInboxMessageRequestInteractiveAction`

NewSendInboxMessageRequestInteractiveAction instantiates a new SendInboxMessageRequestInteractiveAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestInteractiveActionWithDefaults

`func NewSendInboxMessageRequestInteractiveActionWithDefaults() *SendInboxMessageRequestInteractiveAction`

NewSendInboxMessageRequestInteractiveActionWithDefaults instantiates a new SendInboxMessageRequestInteractiveAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButton

`func (o *SendInboxMessageRequestInteractiveAction) GetButton() string`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *SendInboxMessageRequestInteractiveAction) GetButtonOk() (*string, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *SendInboxMessageRequestInteractiveAction) SetButton(v string)`

SetButton sets Button field to given value.


### GetSections

`func (o *SendInboxMessageRequestInteractiveAction) GetSections() []SendInboxMessageRequestInteractiveActionOneOfSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *SendInboxMessageRequestInteractiveAction) GetSectionsOk() (*[]SendInboxMessageRequestInteractiveActionOneOfSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *SendInboxMessageRequestInteractiveAction) SetSections(v []SendInboxMessageRequestInteractiveActionOneOfSectionsInner)`

SetSections sets Sections field to given value.


### GetName

`func (o *SendInboxMessageRequestInteractiveAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SendInboxMessageRequestInteractiveAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SendInboxMessageRequestInteractiveAction) SetName(v string)`

SetName sets Name field to given value.


### GetParameters

`func (o *SendInboxMessageRequestInteractiveAction) GetParameters() SendInboxMessageRequestInteractiveActionOneOf3Parameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *SendInboxMessageRequestInteractiveAction) GetParametersOk() (*SendInboxMessageRequestInteractiveActionOneOf3Parameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *SendInboxMessageRequestInteractiveAction) SetParameters(v SendInboxMessageRequestInteractiveActionOneOf3Parameters)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


