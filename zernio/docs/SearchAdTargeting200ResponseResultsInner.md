# SearchAdTargeting200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The platform&#39;s opaque id. Use as a geo &#x60;key&#x60; (regions/cities/zips/metros) or an entity &#x60;id&#x60; (interests/behaviors) in TargetingSpec. | 
**Name** | **string** | Human-readable label. | 
**Type** | **string** | What the result is (e.g. city, region, country, zip, metro, interest, behavior, income). | 
**Path** | Pointer to **[]string** | Optional breadcrumb of parent labels (e.g. [&#39;United States&#39;, &#39;California&#39;, &#39;Los Angeles&#39;]). Disambiguates same-named results. | [optional] 
**AudienceSize** | Pointer to **NullableInt32** | Optional estimated reachable users for this option, when the platform returns it. | [optional] 

## Methods

### NewSearchAdTargeting200ResponseResultsInner

`func NewSearchAdTargeting200ResponseResultsInner(id string, name string, type_ string, ) *SearchAdTargeting200ResponseResultsInner`

NewSearchAdTargeting200ResponseResultsInner instantiates a new SearchAdTargeting200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAdTargeting200ResponseResultsInnerWithDefaults

`func NewSearchAdTargeting200ResponseResultsInnerWithDefaults() *SearchAdTargeting200ResponseResultsInner`

NewSearchAdTargeting200ResponseResultsInnerWithDefaults instantiates a new SearchAdTargeting200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchAdTargeting200ResponseResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchAdTargeting200ResponseResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchAdTargeting200ResponseResultsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SearchAdTargeting200ResponseResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchAdTargeting200ResponseResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchAdTargeting200ResponseResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SearchAdTargeting200ResponseResultsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchAdTargeting200ResponseResultsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchAdTargeting200ResponseResultsInner) SetType(v string)`

SetType sets Type field to given value.


### GetPath

`func (o *SearchAdTargeting200ResponseResultsInner) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SearchAdTargeting200ResponseResultsInner) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SearchAdTargeting200ResponseResultsInner) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SearchAdTargeting200ResponseResultsInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetAudienceSize

`func (o *SearchAdTargeting200ResponseResultsInner) GetAudienceSize() int32`

GetAudienceSize returns the AudienceSize field if non-nil, zero value otherwise.

### GetAudienceSizeOk

`func (o *SearchAdTargeting200ResponseResultsInner) GetAudienceSizeOk() (*int32, bool)`

GetAudienceSizeOk returns a tuple with the AudienceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceSize

`func (o *SearchAdTargeting200ResponseResultsInner) SetAudienceSize(v int32)`

SetAudienceSize sets AudienceSize field to given value.

### HasAudienceSize

`func (o *SearchAdTargeting200ResponseResultsInner) HasAudienceSize() bool`

HasAudienceSize returns a boolean if a field has been set.

### SetAudienceSizeNil

`func (o *SearchAdTargeting200ResponseResultsInner) SetAudienceSizeNil(b bool)`

 SetAudienceSizeNil sets the value for AudienceSize to be an explicit nil

### UnsetAudienceSize
`func (o *SearchAdTargeting200ResponseResultsInner) UnsetAudienceSize()`

UnsetAudienceSize ensures that no value is present for AudienceSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


