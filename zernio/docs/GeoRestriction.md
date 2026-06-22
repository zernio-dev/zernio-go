# GeoRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | **[]string** | ISO 3166-1 alpha-2 country codes (uppercase). Only users in these countries can see the post. Maximum 25 countries per post. Example: [\&quot;US\&quot;, \&quot;CA\&quot;, \&quot;GB\&quot;, \&quot;ES\&quot;].  | 

## Methods

### NewGeoRestriction

`func NewGeoRestriction(countries []string, ) *GeoRestriction`

NewGeoRestriction instantiates a new GeoRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoRestrictionWithDefaults

`func NewGeoRestrictionWithDefaults() *GeoRestriction`

NewGeoRestrictionWithDefaults instantiates a new GeoRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *GeoRestriction) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *GeoRestriction) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *GeoRestriction) SetCountries(v []string)`

SetCountries sets Countries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


