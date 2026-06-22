# GetGoogleBusinessReviews200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **string** |  | [optional] 
**Reviews** | Pointer to [**[]GetGoogleBusinessReviews200ResponseReviewsInner**](GetGoogleBusinessReviews200ResponseReviewsInner.md) |  | [optional] 
**AverageRating** | Pointer to **float32** | Overall average rating | [optional] 
**TotalReviewCount** | Pointer to **int32** | Total number of reviews | [optional] 
**NextPageToken** | Pointer to **NullableString** | Token for next page | [optional] 

## Methods

### NewGetGoogleBusinessReviews200Response

`func NewGetGoogleBusinessReviews200Response() *GetGoogleBusinessReviews200Response`

NewGetGoogleBusinessReviews200Response instantiates a new GetGoogleBusinessReviews200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGoogleBusinessReviews200ResponseWithDefaults

`func NewGetGoogleBusinessReviews200ResponseWithDefaults() *GetGoogleBusinessReviews200Response`

NewGetGoogleBusinessReviews200ResponseWithDefaults instantiates a new GetGoogleBusinessReviews200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetGoogleBusinessReviews200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetGoogleBusinessReviews200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetGoogleBusinessReviews200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetGoogleBusinessReviews200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAccountId

`func (o *GetGoogleBusinessReviews200Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetGoogleBusinessReviews200Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetGoogleBusinessReviews200Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetGoogleBusinessReviews200Response) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLocationId

`func (o *GetGoogleBusinessReviews200Response) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *GetGoogleBusinessReviews200Response) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *GetGoogleBusinessReviews200Response) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *GetGoogleBusinessReviews200Response) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetReviews

`func (o *GetGoogleBusinessReviews200Response) GetReviews() []GetGoogleBusinessReviews200ResponseReviewsInner`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *GetGoogleBusinessReviews200Response) GetReviewsOk() (*[]GetGoogleBusinessReviews200ResponseReviewsInner, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *GetGoogleBusinessReviews200Response) SetReviews(v []GetGoogleBusinessReviews200ResponseReviewsInner)`

SetReviews sets Reviews field to given value.

### HasReviews

`func (o *GetGoogleBusinessReviews200Response) HasReviews() bool`

HasReviews returns a boolean if a field has been set.

### GetAverageRating

`func (o *GetGoogleBusinessReviews200Response) GetAverageRating() float32`

GetAverageRating returns the AverageRating field if non-nil, zero value otherwise.

### GetAverageRatingOk

`func (o *GetGoogleBusinessReviews200Response) GetAverageRatingOk() (*float32, bool)`

GetAverageRatingOk returns a tuple with the AverageRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRating

`func (o *GetGoogleBusinessReviews200Response) SetAverageRating(v float32)`

SetAverageRating sets AverageRating field to given value.

### HasAverageRating

`func (o *GetGoogleBusinessReviews200Response) HasAverageRating() bool`

HasAverageRating returns a boolean if a field has been set.

### GetTotalReviewCount

`func (o *GetGoogleBusinessReviews200Response) GetTotalReviewCount() int32`

GetTotalReviewCount returns the TotalReviewCount field if non-nil, zero value otherwise.

### GetTotalReviewCountOk

`func (o *GetGoogleBusinessReviews200Response) GetTotalReviewCountOk() (*int32, bool)`

GetTotalReviewCountOk returns a tuple with the TotalReviewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReviewCount

`func (o *GetGoogleBusinessReviews200Response) SetTotalReviewCount(v int32)`

SetTotalReviewCount sets TotalReviewCount field to given value.

### HasTotalReviewCount

`func (o *GetGoogleBusinessReviews200Response) HasTotalReviewCount() bool`

HasTotalReviewCount returns a boolean if a field has been set.

### GetNextPageToken

`func (o *GetGoogleBusinessReviews200Response) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *GetGoogleBusinessReviews200Response) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *GetGoogleBusinessReviews200Response) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *GetGoogleBusinessReviews200Response) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageTokenNil

`func (o *GetGoogleBusinessReviews200Response) SetNextPageTokenNil(b bool)`

 SetNextPageTokenNil sets the value for NextPageToken to be an explicit nil

### UnsetNextPageToken
`func (o *GetGoogleBusinessReviews200Response) UnsetNextPageToken()`

UnsetNextPageToken ensures that no value is present for NextPageToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


