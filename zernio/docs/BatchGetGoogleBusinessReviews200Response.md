# BatchGetGoogleBusinessReviews200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**LocationReviews** | Pointer to [**[]BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner**](BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner.md) |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewBatchGetGoogleBusinessReviews200Response

`func NewBatchGetGoogleBusinessReviews200Response() *BatchGetGoogleBusinessReviews200Response`

NewBatchGetGoogleBusinessReviews200Response instantiates a new BatchGetGoogleBusinessReviews200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchGetGoogleBusinessReviews200ResponseWithDefaults

`func NewBatchGetGoogleBusinessReviews200ResponseWithDefaults() *BatchGetGoogleBusinessReviews200Response`

NewBatchGetGoogleBusinessReviews200ResponseWithDefaults instantiates a new BatchGetGoogleBusinessReviews200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *BatchGetGoogleBusinessReviews200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BatchGetGoogleBusinessReviews200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BatchGetGoogleBusinessReviews200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *BatchGetGoogleBusinessReviews200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAccountId

`func (o *BatchGetGoogleBusinessReviews200Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BatchGetGoogleBusinessReviews200Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BatchGetGoogleBusinessReviews200Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *BatchGetGoogleBusinessReviews200Response) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLocationReviews

`func (o *BatchGetGoogleBusinessReviews200Response) GetLocationReviews() []BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner`

GetLocationReviews returns the LocationReviews field if non-nil, zero value otherwise.

### GetLocationReviewsOk

`func (o *BatchGetGoogleBusinessReviews200Response) GetLocationReviewsOk() (*[]BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner, bool)`

GetLocationReviewsOk returns a tuple with the LocationReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationReviews

`func (o *BatchGetGoogleBusinessReviews200Response) SetLocationReviews(v []BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner)`

SetLocationReviews sets LocationReviews field to given value.

### HasLocationReviews

`func (o *BatchGetGoogleBusinessReviews200Response) HasLocationReviews() bool`

HasLocationReviews returns a boolean if a field has been set.

### GetNextPageToken

`func (o *BatchGetGoogleBusinessReviews200Response) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *BatchGetGoogleBusinessReviews200Response) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *BatchGetGoogleBusinessReviews200Response) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *BatchGetGoogleBusinessReviews200Response) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


