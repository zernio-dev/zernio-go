# GetLinkedInPostReactions200ResponseReactionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReactionType** | Pointer to **string** | LinkedIn reaction enum (LIKE, PRAISE, EMPATHY, INTEREST, APPRECIATION, ENTERTAINMENT) | [optional] 
**ReactionLabel** | Pointer to **string** | User-friendly label (Like, Celebrate, Love, Insightful, Support, Funny) | [optional] 
**ReactedAt** | Pointer to **time.Time** |  | [optional] 
**From** | Pointer to [**GetLinkedInPostReactions200ResponseReactionsInnerFrom**](GetLinkedInPostReactions200ResponseReactionsInnerFrom.md) |  | [optional] 

## Methods

### NewGetLinkedInPostReactions200ResponseReactionsInner

`func NewGetLinkedInPostReactions200ResponseReactionsInner() *GetLinkedInPostReactions200ResponseReactionsInner`

NewGetLinkedInPostReactions200ResponseReactionsInner instantiates a new GetLinkedInPostReactions200ResponseReactionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinkedInPostReactions200ResponseReactionsInnerWithDefaults

`func NewGetLinkedInPostReactions200ResponseReactionsInnerWithDefaults() *GetLinkedInPostReactions200ResponseReactionsInner`

NewGetLinkedInPostReactions200ResponseReactionsInnerWithDefaults instantiates a new GetLinkedInPostReactions200ResponseReactionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReactionType

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) GetReactionType() string`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) GetReactionTypeOk() (*string, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) SetReactionType(v string)`

SetReactionType sets ReactionType field to given value.

### HasReactionType

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) HasReactionType() bool`

HasReactionType returns a boolean if a field has been set.

### GetReactionLabel

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) GetReactionLabel() string`

GetReactionLabel returns the ReactionLabel field if non-nil, zero value otherwise.

### GetReactionLabelOk

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) GetReactionLabelOk() (*string, bool)`

GetReactionLabelOk returns a tuple with the ReactionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionLabel

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) SetReactionLabel(v string)`

SetReactionLabel sets ReactionLabel field to given value.

### HasReactionLabel

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) HasReactionLabel() bool`

HasReactionLabel returns a boolean if a field has been set.

### GetReactedAt

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) GetReactedAt() time.Time`

GetReactedAt returns the ReactedAt field if non-nil, zero value otherwise.

### GetReactedAtOk

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) GetReactedAtOk() (*time.Time, bool)`

GetReactedAtOk returns a tuple with the ReactedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactedAt

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) SetReactedAt(v time.Time)`

SetReactedAt sets ReactedAt field to given value.

### HasReactedAt

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) HasReactedAt() bool`

HasReactedAt returns a boolean if a field has been set.

### GetFrom

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) GetFrom() GetLinkedInPostReactions200ResponseReactionsInnerFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) GetFromOk() (*GetLinkedInPostReactions200ResponseReactionsInnerFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) SetFrom(v GetLinkedInPostReactions200ResponseReactionsInnerFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetLinkedInPostReactions200ResponseReactionsInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


