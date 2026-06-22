# BatchGetGoogleBusinessReviewsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationNames** | **[]string** | Array of full location resource names (e.g. [&#39;accounts/123/locations/456&#39;]) | 
**PageSize** | Pointer to **int32** | Number of reviews per page (max 50) | [optional] [default to 50]
**PageToken** | Pointer to **string** | Pagination token from previous response | [optional] 

## Methods

### NewBatchGetGoogleBusinessReviewsRequest

`func NewBatchGetGoogleBusinessReviewsRequest(locationNames []string, ) *BatchGetGoogleBusinessReviewsRequest`

NewBatchGetGoogleBusinessReviewsRequest instantiates a new BatchGetGoogleBusinessReviewsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchGetGoogleBusinessReviewsRequestWithDefaults

`func NewBatchGetGoogleBusinessReviewsRequestWithDefaults() *BatchGetGoogleBusinessReviewsRequest`

NewBatchGetGoogleBusinessReviewsRequestWithDefaults instantiates a new BatchGetGoogleBusinessReviewsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationNames

`func (o *BatchGetGoogleBusinessReviewsRequest) GetLocationNames() []string`

GetLocationNames returns the LocationNames field if non-nil, zero value otherwise.

### GetLocationNamesOk

`func (o *BatchGetGoogleBusinessReviewsRequest) GetLocationNamesOk() (*[]string, bool)`

GetLocationNamesOk returns a tuple with the LocationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationNames

`func (o *BatchGetGoogleBusinessReviewsRequest) SetLocationNames(v []string)`

SetLocationNames sets LocationNames field to given value.


### GetPageSize

`func (o *BatchGetGoogleBusinessReviewsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *BatchGetGoogleBusinessReviewsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *BatchGetGoogleBusinessReviewsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *BatchGetGoogleBusinessReviewsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *BatchGetGoogleBusinessReviewsRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *BatchGetGoogleBusinessReviewsRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *BatchGetGoogleBusinessReviewsRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *BatchGetGoogleBusinessReviewsRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


