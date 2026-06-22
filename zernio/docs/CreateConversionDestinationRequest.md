# CreateConversionDestinationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdAccountId** | **string** | Ad account ID. For LinkedIn: numeric (e.g. \&quot;5123456\&quot;) or full &#x60;urn:li:sponsoredAccount:{id}&#x60; URN. For Google: numeric customer ID (e.g. \&quot;1234567890\&quot;) or &#x60;customers/{id}&#x60; form.  | 
**Name** | **string** |  | 
**Type** | **string** | Conversion type. For LinkedIn: a unified standard event name (e.g. \&quot;Purchase\&quot;, \&quot;Lead\&quot;, \&quot;AddToCart\&quot;) or a LinkedIn rule type enum (e.g. \&quot;PURCHASE\&quot;, \&quot;QUALIFIED_LEAD\&quot;). For Google: a unified standard event name (Purchase, Subscribe, CompleteRegistration, Lead, Schedule) or a Google ConversionActionCategory enum value directly (e.g. \&quot;PURCHASE\&quot;, \&quot;SUBSCRIBE_PAID\&quot;, \&quot;SIGNUP\&quot;, \&quot;IMPORTED_LEAD\&quot;, \&quot;BOOK_APPOINTMENT\&quot;). Unknown values pass through to the platform.  | 
**AttributionType** | Pointer to **string** | LinkedIn only. | [optional] 
**PostClickAttributionWindowSize** | Pointer to **int32** | LinkedIn only. Default 30. 365 only allowed for LEAD, PURCHASE, ADD_TO_CART, QUALIFIED_LEAD, SUBMIT_APPLICATION rule types; the API rejects other combinations locally.  | [optional] 
**ViewThroughAttributionWindowSize** | Pointer to **int32** | LinkedIn only. Default 7. Same 365-day-window type restriction applies as &#x60;postClickAttributionWindowSize&#x60;.  | [optional] 
**ValueType** | Pointer to **string** | LinkedIn only. DYNAMIC (default) uses the per-event &#x60;value&#x60; from &#x60;sendConversions&#x60;. FIXED uses the rule&#39;s &#x60;value&#x60; field. NO_VALUE drops monetary value entirely.  | [optional] 
**Value** | Pointer to [**CreateConversionDestinationRequestValue**](CreateConversionDestinationRequestValue.md) |  | [optional] 
**AutoAssociationType** | Pointer to **string** | LinkedIn only. Controls campaign association at rule-creation time: - ALL_CAMPAIGNS: associate the rule with every active,   paused, and draft campaign in the ad account - OBJECTIVE_BASED: associate only campaigns whose   objective matches the rule&#39;s type - NONE: don&#39;t auto-associate. Manage associations via   the &#x60;/associations&#x60; endpoints below. Note: auto-association runs once at create time; new campaigns added after the rule still need explicit association.  | [optional] [default to "ALL_CAMPAIGNS"]
**CountingType** | Pointer to **string** | Google Ads only. Whether to count multiple conversions from the same click (MANY_PER_CLICK) or at most one (ONE_PER_CLICK). Defaults to MANY_PER_CLICK if omitted.  | [optional] 
**PrimaryForGoal** | Pointer to **bool** | Google Ads only. When true, the conversion action is marked as primary and immediately influences Smart Bidding. Defaults to false (secondary, record-only) to avoid unintentionally steering the customer&#39;s campaigns on creation.  | [optional] 

## Methods

### NewCreateConversionDestinationRequest

`func NewCreateConversionDestinationRequest(adAccountId string, name string, type_ string, ) *CreateConversionDestinationRequest`

NewCreateConversionDestinationRequest instantiates a new CreateConversionDestinationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConversionDestinationRequestWithDefaults

`func NewCreateConversionDestinationRequestWithDefaults() *CreateConversionDestinationRequest`

NewCreateConversionDestinationRequestWithDefaults instantiates a new CreateConversionDestinationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdAccountId

`func (o *CreateConversionDestinationRequest) GetAdAccountId() string`

GetAdAccountId returns the AdAccountId field if non-nil, zero value otherwise.

### GetAdAccountIdOk

`func (o *CreateConversionDestinationRequest) GetAdAccountIdOk() (*string, bool)`

GetAdAccountIdOk returns a tuple with the AdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountId

`func (o *CreateConversionDestinationRequest) SetAdAccountId(v string)`

SetAdAccountId sets AdAccountId field to given value.


### GetName

`func (o *CreateConversionDestinationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateConversionDestinationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateConversionDestinationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateConversionDestinationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateConversionDestinationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateConversionDestinationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributionType

`func (o *CreateConversionDestinationRequest) GetAttributionType() string`

GetAttributionType returns the AttributionType field if non-nil, zero value otherwise.

### GetAttributionTypeOk

`func (o *CreateConversionDestinationRequest) GetAttributionTypeOk() (*string, bool)`

GetAttributionTypeOk returns a tuple with the AttributionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionType

`func (o *CreateConversionDestinationRequest) SetAttributionType(v string)`

SetAttributionType sets AttributionType field to given value.

### HasAttributionType

`func (o *CreateConversionDestinationRequest) HasAttributionType() bool`

HasAttributionType returns a boolean if a field has been set.

### GetPostClickAttributionWindowSize

`func (o *CreateConversionDestinationRequest) GetPostClickAttributionWindowSize() int32`

GetPostClickAttributionWindowSize returns the PostClickAttributionWindowSize field if non-nil, zero value otherwise.

### GetPostClickAttributionWindowSizeOk

`func (o *CreateConversionDestinationRequest) GetPostClickAttributionWindowSizeOk() (*int32, bool)`

GetPostClickAttributionWindowSizeOk returns a tuple with the PostClickAttributionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostClickAttributionWindowSize

`func (o *CreateConversionDestinationRequest) SetPostClickAttributionWindowSize(v int32)`

SetPostClickAttributionWindowSize sets PostClickAttributionWindowSize field to given value.

### HasPostClickAttributionWindowSize

`func (o *CreateConversionDestinationRequest) HasPostClickAttributionWindowSize() bool`

HasPostClickAttributionWindowSize returns a boolean if a field has been set.

### GetViewThroughAttributionWindowSize

`func (o *CreateConversionDestinationRequest) GetViewThroughAttributionWindowSize() int32`

GetViewThroughAttributionWindowSize returns the ViewThroughAttributionWindowSize field if non-nil, zero value otherwise.

### GetViewThroughAttributionWindowSizeOk

`func (o *CreateConversionDestinationRequest) GetViewThroughAttributionWindowSizeOk() (*int32, bool)`

GetViewThroughAttributionWindowSizeOk returns a tuple with the ViewThroughAttributionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewThroughAttributionWindowSize

`func (o *CreateConversionDestinationRequest) SetViewThroughAttributionWindowSize(v int32)`

SetViewThroughAttributionWindowSize sets ViewThroughAttributionWindowSize field to given value.

### HasViewThroughAttributionWindowSize

`func (o *CreateConversionDestinationRequest) HasViewThroughAttributionWindowSize() bool`

HasViewThroughAttributionWindowSize returns a boolean if a field has been set.

### GetValueType

`func (o *CreateConversionDestinationRequest) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *CreateConversionDestinationRequest) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *CreateConversionDestinationRequest) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *CreateConversionDestinationRequest) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetValue

`func (o *CreateConversionDestinationRequest) GetValue() CreateConversionDestinationRequestValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateConversionDestinationRequest) GetValueOk() (*CreateConversionDestinationRequestValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateConversionDestinationRequest) SetValue(v CreateConversionDestinationRequestValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreateConversionDestinationRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetAutoAssociationType

`func (o *CreateConversionDestinationRequest) GetAutoAssociationType() string`

GetAutoAssociationType returns the AutoAssociationType field if non-nil, zero value otherwise.

### GetAutoAssociationTypeOk

`func (o *CreateConversionDestinationRequest) GetAutoAssociationTypeOk() (*string, bool)`

GetAutoAssociationTypeOk returns a tuple with the AutoAssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAssociationType

`func (o *CreateConversionDestinationRequest) SetAutoAssociationType(v string)`

SetAutoAssociationType sets AutoAssociationType field to given value.

### HasAutoAssociationType

`func (o *CreateConversionDestinationRequest) HasAutoAssociationType() bool`

HasAutoAssociationType returns a boolean if a field has been set.

### GetCountingType

`func (o *CreateConversionDestinationRequest) GetCountingType() string`

GetCountingType returns the CountingType field if non-nil, zero value otherwise.

### GetCountingTypeOk

`func (o *CreateConversionDestinationRequest) GetCountingTypeOk() (*string, bool)`

GetCountingTypeOk returns a tuple with the CountingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountingType

`func (o *CreateConversionDestinationRequest) SetCountingType(v string)`

SetCountingType sets CountingType field to given value.

### HasCountingType

`func (o *CreateConversionDestinationRequest) HasCountingType() bool`

HasCountingType returns a boolean if a field has been set.

### GetPrimaryForGoal

`func (o *CreateConversionDestinationRequest) GetPrimaryForGoal() bool`

GetPrimaryForGoal returns the PrimaryForGoal field if non-nil, zero value otherwise.

### GetPrimaryForGoalOk

`func (o *CreateConversionDestinationRequest) GetPrimaryForGoalOk() (*bool, bool)`

GetPrimaryForGoalOk returns a tuple with the PrimaryForGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryForGoal

`func (o *CreateConversionDestinationRequest) SetPrimaryForGoal(v bool)`

SetPrimaryForGoal sets PrimaryForGoal field to given value.

### HasPrimaryForGoal

`func (o *CreateConversionDestinationRequest) HasPrimaryForGoal() bool`

HasPrimaryForGoal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


