# SendInboxMessageRequestInteractiveActionOneOfSectionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Optional section header (max 24 chars). | [optional] 
**Rows** | [**[]SendInboxMessageRequestInteractiveActionOneOfSectionsInnerRowsInner**](SendInboxMessageRequestInteractiveActionOneOfSectionsInnerRowsInner.md) |  | 

## Methods

### NewSendInboxMessageRequestInteractiveActionOneOfSectionsInner

`func NewSendInboxMessageRequestInteractiveActionOneOfSectionsInner(rows []SendInboxMessageRequestInteractiveActionOneOfSectionsInnerRowsInner, ) *SendInboxMessageRequestInteractiveActionOneOfSectionsInner`

NewSendInboxMessageRequestInteractiveActionOneOfSectionsInner instantiates a new SendInboxMessageRequestInteractiveActionOneOfSectionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestInteractiveActionOneOfSectionsInnerWithDefaults

`func NewSendInboxMessageRequestInteractiveActionOneOfSectionsInnerWithDefaults() *SendInboxMessageRequestInteractiveActionOneOfSectionsInner`

NewSendInboxMessageRequestInteractiveActionOneOfSectionsInnerWithDefaults instantiates a new SendInboxMessageRequestInteractiveActionOneOfSectionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SendInboxMessageRequestInteractiveActionOneOfSectionsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendInboxMessageRequestInteractiveActionOneOfSectionsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendInboxMessageRequestInteractiveActionOneOfSectionsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SendInboxMessageRequestInteractiveActionOneOfSectionsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRows

`func (o *SendInboxMessageRequestInteractiveActionOneOfSectionsInner) GetRows() []SendInboxMessageRequestInteractiveActionOneOfSectionsInnerRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *SendInboxMessageRequestInteractiveActionOneOfSectionsInner) GetRowsOk() (*[]SendInboxMessageRequestInteractiveActionOneOfSectionsInnerRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *SendInboxMessageRequestInteractiveActionOneOfSectionsInner) SetRows(v []SendInboxMessageRequestInteractiveActionOneOfSectionsInnerRowsInner)`

SetRows sets Rows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


