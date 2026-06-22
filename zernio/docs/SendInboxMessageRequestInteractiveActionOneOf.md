# SendInboxMessageRequestInteractiveActionOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Button** | **string** | CTA label that opens the list (max ~20 chars). | 
**Sections** | [**[]SendInboxMessageRequestInteractiveActionOneOfSectionsInner**](SendInboxMessageRequestInteractiveActionOneOfSectionsInner.md) | 1-10 sections. Total rows across all sections cannot exceed 10. | 

## Methods

### NewSendInboxMessageRequestInteractiveActionOneOf

`func NewSendInboxMessageRequestInteractiveActionOneOf(button string, sections []SendInboxMessageRequestInteractiveActionOneOfSectionsInner, ) *SendInboxMessageRequestInteractiveActionOneOf`

NewSendInboxMessageRequestInteractiveActionOneOf instantiates a new SendInboxMessageRequestInteractiveActionOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestInteractiveActionOneOfWithDefaults

`func NewSendInboxMessageRequestInteractiveActionOneOfWithDefaults() *SendInboxMessageRequestInteractiveActionOneOf`

NewSendInboxMessageRequestInteractiveActionOneOfWithDefaults instantiates a new SendInboxMessageRequestInteractiveActionOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButton

`func (o *SendInboxMessageRequestInteractiveActionOneOf) GetButton() string`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *SendInboxMessageRequestInteractiveActionOneOf) GetButtonOk() (*string, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *SendInboxMessageRequestInteractiveActionOneOf) SetButton(v string)`

SetButton sets Button field to given value.


### GetSections

`func (o *SendInboxMessageRequestInteractiveActionOneOf) GetSections() []SendInboxMessageRequestInteractiveActionOneOfSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *SendInboxMessageRequestInteractiveActionOneOf) GetSectionsOk() (*[]SendInboxMessageRequestInteractiveActionOneOfSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *SendInboxMessageRequestInteractiveActionOneOf) SetSections(v []SendInboxMessageRequestInteractiveActionOneOfSectionsInner)`

SetSections sets Sections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


