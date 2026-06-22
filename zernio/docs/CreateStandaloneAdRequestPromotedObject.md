# CreateStandaloneAdRequestPromotedObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PixelId** | Pointer to **string** | Pixel ID. **Meta:** Facebook Pixel ID, required for &#x60;goal: conversions&#x60;. **TikTok:** TikTok Pixel ID, required for &#x60;goal: conversions&#x60;.  | [optional] 
**CustomEventType** | Pointer to **string** | The event the campaign/ad group optimises against.  **Meta:** standard event like &#x60;PURCHASE&#x60;, &#x60;LEAD&#x60;, &#x60;COMPLETE_REGISTRATION&#x60;, &#x60;ADD_TO_CART&#x60;. Uppercased internally so callers can pass any case. Required for &#x60;goal: conversions&#x60;.  **TikTok:** an &#x60;optimization_event&#x60; code (UPPER_SNAKE, not Meta&#39;s vocabulary and not PascalCase), e.g. &#x60;ON_WEB_ORDER&#x60; (Complete Payment), &#x60;INITIATE_ORDER&#x60; (Place an Order), &#x60;ON_WEB_CART&#x60; (Add to Cart), &#x60;ON_WEB_REGISTER&#x60; (Complete Registration), &#x60;FORM&#x60; (Submit Form), &#x60;ON_WEB_DETAIL&#x60; (View Content). Must be one of the events your TikTok Pixel is configured to track; passed through verbatim and validated by TikTok. Optional for &#x60;goal: conversions&#x60;.  | [optional] 
**PageId** | Pointer to **string** | Facebook Page ID. Used by &#x60;goal: lead_generation&#x60;. Auto-filled from the connected Page when omitted.  | [optional] 
**ApplicationId** | Pointer to **string** | App ID. Required for &#x60;goal: app_promotion&#x60;. | [optional] 
**ObjectStoreUrl** | Pointer to **string** | App Store / Play Store listing URL. Required for &#x60;goal: app_promotion&#x60;. | [optional] 
**CustomConversionId** | Pointer to **string** | Custom Conversion ID, when optimising against one instead of a standard event. | [optional] 
**ProductCatalogId** | Pointer to **string** | Catalog ID for catalog/Advantage+ Shopping campaigns. | [optional] 
**ProductSetId** | Pointer to **string** | Product Set ID inside the catalog. | [optional] 

## Methods

### NewCreateStandaloneAdRequestPromotedObject

`func NewCreateStandaloneAdRequestPromotedObject() *CreateStandaloneAdRequestPromotedObject`

NewCreateStandaloneAdRequestPromotedObject instantiates a new CreateStandaloneAdRequestPromotedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandaloneAdRequestPromotedObjectWithDefaults

`func NewCreateStandaloneAdRequestPromotedObjectWithDefaults() *CreateStandaloneAdRequestPromotedObject`

NewCreateStandaloneAdRequestPromotedObjectWithDefaults instantiates a new CreateStandaloneAdRequestPromotedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPixelId

`func (o *CreateStandaloneAdRequestPromotedObject) GetPixelId() string`

GetPixelId returns the PixelId field if non-nil, zero value otherwise.

### GetPixelIdOk

`func (o *CreateStandaloneAdRequestPromotedObject) GetPixelIdOk() (*string, bool)`

GetPixelIdOk returns a tuple with the PixelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixelId

`func (o *CreateStandaloneAdRequestPromotedObject) SetPixelId(v string)`

SetPixelId sets PixelId field to given value.

### HasPixelId

`func (o *CreateStandaloneAdRequestPromotedObject) HasPixelId() bool`

HasPixelId returns a boolean if a field has been set.

### GetCustomEventType

`func (o *CreateStandaloneAdRequestPromotedObject) GetCustomEventType() string`

GetCustomEventType returns the CustomEventType field if non-nil, zero value otherwise.

### GetCustomEventTypeOk

`func (o *CreateStandaloneAdRequestPromotedObject) GetCustomEventTypeOk() (*string, bool)`

GetCustomEventTypeOk returns a tuple with the CustomEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEventType

`func (o *CreateStandaloneAdRequestPromotedObject) SetCustomEventType(v string)`

SetCustomEventType sets CustomEventType field to given value.

### HasCustomEventType

`func (o *CreateStandaloneAdRequestPromotedObject) HasCustomEventType() bool`

HasCustomEventType returns a boolean if a field has been set.

### GetPageId

`func (o *CreateStandaloneAdRequestPromotedObject) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *CreateStandaloneAdRequestPromotedObject) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *CreateStandaloneAdRequestPromotedObject) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *CreateStandaloneAdRequestPromotedObject) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetApplicationId

`func (o *CreateStandaloneAdRequestPromotedObject) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CreateStandaloneAdRequestPromotedObject) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CreateStandaloneAdRequestPromotedObject) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *CreateStandaloneAdRequestPromotedObject) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetObjectStoreUrl

`func (o *CreateStandaloneAdRequestPromotedObject) GetObjectStoreUrl() string`

GetObjectStoreUrl returns the ObjectStoreUrl field if non-nil, zero value otherwise.

### GetObjectStoreUrlOk

`func (o *CreateStandaloneAdRequestPromotedObject) GetObjectStoreUrlOk() (*string, bool)`

GetObjectStoreUrlOk returns a tuple with the ObjectStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreUrl

`func (o *CreateStandaloneAdRequestPromotedObject) SetObjectStoreUrl(v string)`

SetObjectStoreUrl sets ObjectStoreUrl field to given value.

### HasObjectStoreUrl

`func (o *CreateStandaloneAdRequestPromotedObject) HasObjectStoreUrl() bool`

HasObjectStoreUrl returns a boolean if a field has been set.

### GetCustomConversionId

`func (o *CreateStandaloneAdRequestPromotedObject) GetCustomConversionId() string`

GetCustomConversionId returns the CustomConversionId field if non-nil, zero value otherwise.

### GetCustomConversionIdOk

`func (o *CreateStandaloneAdRequestPromotedObject) GetCustomConversionIdOk() (*string, bool)`

GetCustomConversionIdOk returns a tuple with the CustomConversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConversionId

`func (o *CreateStandaloneAdRequestPromotedObject) SetCustomConversionId(v string)`

SetCustomConversionId sets CustomConversionId field to given value.

### HasCustomConversionId

`func (o *CreateStandaloneAdRequestPromotedObject) HasCustomConversionId() bool`

HasCustomConversionId returns a boolean if a field has been set.

### GetProductCatalogId

`func (o *CreateStandaloneAdRequestPromotedObject) GetProductCatalogId() string`

GetProductCatalogId returns the ProductCatalogId field if non-nil, zero value otherwise.

### GetProductCatalogIdOk

`func (o *CreateStandaloneAdRequestPromotedObject) GetProductCatalogIdOk() (*string, bool)`

GetProductCatalogIdOk returns a tuple with the ProductCatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCatalogId

`func (o *CreateStandaloneAdRequestPromotedObject) SetProductCatalogId(v string)`

SetProductCatalogId sets ProductCatalogId field to given value.

### HasProductCatalogId

`func (o *CreateStandaloneAdRequestPromotedObject) HasProductCatalogId() bool`

HasProductCatalogId returns a boolean if a field has been set.

### GetProductSetId

`func (o *CreateStandaloneAdRequestPromotedObject) GetProductSetId() string`

GetProductSetId returns the ProductSetId field if non-nil, zero value otherwise.

### GetProductSetIdOk

`func (o *CreateStandaloneAdRequestPromotedObject) GetProductSetIdOk() (*string, bool)`

GetProductSetIdOk returns a tuple with the ProductSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSetId

`func (o *CreateStandaloneAdRequestPromotedObject) SetProductSetId(v string)`

SetProductSetId sets ProductSetId field to given value.

### HasProductSetId

`func (o *CreateStandaloneAdRequestPromotedObject) HasProductSetId() bool`

HasProductSetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


