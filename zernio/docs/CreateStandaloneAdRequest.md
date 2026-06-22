# CreateStandaloneAdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**AdAccountId** | **string** |  | 
**Name** | **string** |  | 
**CampaignName** | Pointer to **string** | Meta only. Exact campaign name. Overrides the default &#x60;&lt;name&gt; - Campaign&#x60;. | [optional] 
**AdSetName** | Pointer to **string** | Meta only. Exact ad set name. Overrides the default &#x60;&lt;name&gt; - Ad Set&#x60;. (For per-ad names on the multi-creative shape, set &#x60;name&#x60; on each &#x60;creatives[]&#x60; entry.) | [optional] 
**AdName** | Pointer to **string** | Meta only. Exact ad name (the single-creative ad object&#39;s name). Overrides the default, which is &#x60;name&#x60;. (For per-ad names on the multi-creative shape, set &#x60;name&#x60; on each &#x60;creatives[]&#x60; entry instead.) | [optional] 
**Tracking** | Pointer to [**CreateStandaloneAdRequestTracking**](CreateStandaloneAdRequestTracking.md) |  | [optional] 
**Goal** | Pointer to **string** | Required on legacy + multi-creative shapes. Inherited from the ad set on the attach shape. Available goals vary by platform. Meta-specific: &#x60;conversions&#x60; (OUTCOME_SALES) requires &#x60;promotedObject.pixelId&#x60; + &#x60;promotedObject.customEventType&#x60; (use a commerce event, e.g. PURCHASE, START_TRIAL); &#x60;lead_conversion&#x60; (OUTCOME_LEADS, website pixel leads) requires the same pixel + event but with a leads-class event (e.g. LEAD, SUBMIT_APPLICATION, SCHEDULE, CONTACT) — these are rejected under &#x60;conversions&#x60; because Meta gates conversion events by objective; &#x60;lead_generation&#x60; is OUTCOME_LEADS with instant forms (&#x60;leadGenFormId&#x60;), distinct from &#x60;lead_conversion&#x60;&#39;s website pixel optimization; &#x60;app_promotion&#x60; requires &#x60;promotedObject.applicationId&#x60; + &#x60;promotedObject.objectStoreUrl&#x60;; &#x60;catalog_sales&#x60; (Advantage+ catalog ads, e.g. vehicle inventory) requires &#x60;promotedObject.productSetId&#x60; + &#x60;promotedObject.pixelId&#x60; + &#x60;promotedObject.customEventType&#x60; and builds a catalog TEMPLATE creative from the copy fields (headline/body/description/linkUrl/callToAction, which may carry catalog template tags like {{product.name}} or {{vehicle.make}}) — no imageUrl/video is sent, Meta renders the visuals per catalog item; discover catalogs via GET /v1/ads/catalogs and product sets via GET /v1/ads/catalogs/{catalogId}/product-sets; single shape only (no creatives[]/adSetId/dynamicCreative/placementAssets); &#x60;lead_generation&#x60; accepts an optional &#x60;promotedObject.pageId&#x60; (auto-filled from the connected Page when omitted). TikTok-specific: &#x60;conversions&#x60; (website-conversion ad group) requires &#x60;promotedObject.pixelId&#x60; (your TikTok Pixel ID) and accepts an optional &#x60;promotedObject.customEventType&#x60; (a TikTok &#x60;optimization_event&#x60; code like &#x60;ON_WEB_ORDER&#x60;, &#x60;INITIATE_ORDER&#x60;, &#x60;ON_WEB_REGISTER&#x60;, &#x60;FORM&#x60;); to inherit a pixel + event from an existing ad group, pass &#x60;adSetId&#x60; instead. LinkedIn-specific: &#x60;engagement&#x60;, &#x60;traffic&#x60;, &#x60;awareness&#x60;, and &#x60;video_views&#x60; are supported for standalone ads (creates a Direct Sponsored Content single image or single video ad). &#x60;traffic&#x60; requires &#x60;linkUrl&#x60;; &#x60;video_views&#x60; requires the &#x60;video&#x60; field. For &#x60;lead_generation&#x60; / &#x60;conversions&#x60; on LinkedIn — or to promote an existing post — use &#x60;POST /v1/ads/boost&#x60;. | [optional] 
**OptimizationGoal** | Pointer to **string** | Meta only. Explicit ad-set &#x60;optimization_goal&#x60; (e.g. &#x60;LANDING_PAGE_VIEWS&#x60;, &#x60;LINK_CLICKS&#x60;, &#x60;REACH&#x60;, &#x60;IMPRESSIONS&#x60;, &#x60;OFFSITE_CONVERSIONS&#x60;, &#x60;THRUPLAY&#x60;, &#x60;LEAD_GENERATION&#x60;). Overrides the default derived from &#x60;goal&#x60; (e.g. &#x60;traffic&#x60; defaults to &#x60;LINK_CLICKS&#x60;). Forwarded verbatim to Meta, which validates compatibility with the campaign objective and rejects incompatible combinations. | [optional] 
**BudgetAmount** | Pointer to **float32** | Required on legacy + multi-creative shapes. Inherited on attach. | [optional] 
**BudgetType** | Pointer to **string** | Required on legacy + multi-creative shapes. Inherited on attach. | [optional] 
**Status** | Pointer to **string** | Meta only. Publish state of the created ad set + ad. Omitted or ACTIVE publishes live (default, back-compat); PAUSED creates them paused and skips activation, so you can review before they spend. | [optional] 
**BudgetLevel** | Pointer to **string** | Meta only. Where the budget lives, which selects the Meta budget model:   - &#x60;adset&#x60; (default): ABO (Ad-set Budget Optimization). The budget is set on the     ad set. This is the back-compatible behaviour — omit this field to keep it.   - &#x60;campaign&#x60;: CBO (Campaign Budget Optimization / Advantage Campaign Budget). The     budget AND &#x60;bidStrategy&#x60; are set on the CAMPAIGN, and Meta distributes spend     across ad sets automatically. Meta requires the budget at exactly one level, never both. Non-Meta platforms ignore this field. Ignored on the attach shape (&#x60;adSetId&#x60;), which inherits the existing budget.  | [optional] [default to "adset"]
**Currency** | Pointer to **string** |  | [optional] 
**Headline** | Pointer to **string** | Required for Meta, Google, Pinterest, and LinkedIn on legacy + attach shapes (skip for multi-creative — use &#x60;creatives[].headline&#x60;). Ignored for TikTok and X/Twitter. Max: Meta&#x3D;255, Google&#x3D;30, Pinterest&#x3D;100, LinkedIn&#x3D;400. On LinkedIn this is the ad&#39;s headline (the bold text on the creative); for traffic ads it&#39;s the link card title. | [optional] 
**LongHeadline** | Pointer to **string** | Google Display only — defaults to &#x60;headline&#x60; if omitted. On LinkedIn, reused as the optional secondary description text on traffic (link) ads; omitted if not provided. | [optional] 
**Body** | Pointer to **string** | Required on legacy + attach shapes. For X/Twitter this is the tweet text (max 280 chars including a ~24-char URL when &#x60;linkUrl&#x60; is set). On LinkedIn this is the post commentary (the intro text shown above the ad). Max: Google&#x3D;90, Pinterest&#x3D;500. | [optional] 
**Description** | Pointer to **string** | Meta only (facebook/instagram). Link description — the secondary text shown below the headline (Meta&#39;s link_data.description; on video creatives mapped to video_data.link_description). When omitted, Meta auto-pulls the destination URL&#39;s OpenGraph description. Applies on legacy, attach, and placementAssets shapes; for multi-creative use creatives[].description (this field is the shared fallback). For multi-text variations use dynamicCreative.descriptions instead. | [optional] 
**CallToAction** | Pointer to **string** | Required on legacy + attach shapes for Meta. Honoured on TikTok (passes through to the Spark Ad creative&#39;s &#x60;call_to_action&#x60;) and on LinkedIn (the CTA button on the ad; defaults to LEARN_MORE when &#x60;linkUrl&#x60; is set). LinkedIn accepts: LEARN_MORE, SIGN_UP, DOWNLOAD, SUBSCRIBE, REGISTER, JOIN, ATTEND, REQUEST_DEMO, VIEW_QUOTE, APPLY, SEE_MORE, SHOP_NOW, BUY_NOW. Ignored by Google, Pinterest, and X/Twitter. | [optional] 
**LinkUrl** | Pointer to **string** | Required on legacy + attach shapes (skip for multi-creative). On LinkedIn it&#39;s the ad&#39;s destination URL; required for &#x60;traffic&#x60; ads, optional for &#x60;engagement&#x60; / &#x60;awareness&#x60;. NOT required when &#x60;goal&#x60; is &#x60;lead_generation&#x60; (the ad opens a Lead Gen form instead of a destination). | [optional] 
**LeadGenFormId** | Pointer to **string** | Meta Lead Gen forms only (facebook/instagram). The leadgen_forms ID to attach to the ad&#39;s creative — create one via POST /v1/ads/lead-forms. REQUIRED when &#x60;goal&#x60; is &#x60;lead_generation&#x60;, and on every ATTACH (&#x60;adSetId&#x60;) call that targets a lead ad set (the form attaches per-ad; Meta rejects a formless ad in a lead ad set). Ignored otherwise. The ad set&#39;s promoted_object.page_id + LEAD_GENERATION optimization + destination_type ON_AD are derived automatically from the goal. Both &#x60;placementAssets&#x60; (per-placement creative) and &#x60;dynamicCreative&#x60; (multi-text / multi-asset pool, e.g. multiple headlines and primary texts) ARE supported on instant-form lead ads — the form is attached for you, and for &#x60;dynamicCreative&#x60; the ad set is created as a Dynamic Creative ad set automatically (Meta requires that for any multi-text feed; there is no non-DCO multi-text path). Send a single &#x60;imageUrls&#x60; entry plus your text variations to get Meta&#39;s \&quot;Multiple Text Options\&quot; behavior on a lead ad. | [optional] 
**ImageUrl** | Pointer to **string** | Image creative for Meta/Google/Pinterest/LinkedIn on legacy + attach shapes (mutually exclusive with &#x60;video&#x60;). Required for LinkedIn ads unless &#x60;video&#x60; is set. Not required for Google Search campaigns. For TikTok, this field carries the VIDEO URL (the TikTok ads endpoint is video-only; the field retains the &#x60;imageUrl&#x60; name for cross-platform consistency). Ignored for X/Twitter. For Google Display, treated as the landscape image (alias of &#x60;images.landscape&#x60;); supply &#x60;images.square&#x60; alongside or the request is rejected. For LinkedIn the image is uploaded to LinkedIn under the authoring Company Page (see &#x60;organizationId&#x60;); recommended ratio 1.91:1 (e.g. 1200×627). | [optional] 
**Images** | Pointer to [**CreateStandaloneAdRequestImages**](CreateStandaloneAdRequestImages.md) |  | [optional] 
**Video** | Pointer to [**CreateStandaloneAdRequestVideo**](CreateStandaloneAdRequestVideo.md) |  | [optional] 
**Creatives** | Pointer to [**[]CreateStandaloneAdRequestCreativesInner**](CreateStandaloneAdRequestCreativesInner.md) | Meta-only. When present, switches to the multi-creative shape: creates 1 campaign + 1 ad set + N ads (one per entry here). Top-level &#x60;headline&#x60; / &#x60;body&#x60; / &#x60;imageUrl&#x60; / &#x60;linkUrl&#x60; / &#x60;callToAction&#x60; are ignored in this mode. Mutually exclusive with &#x60;adSetId&#x60;.  | [optional] 
**AdSetId** | Pointer to **string** | Meta-only. When present, switches to the attach shape: adds one new ad to this existing ad set without creating a new campaign. Budget, targeting, goal, schedule, AND bid strategy are inherited from the ad set on Meta — passing &#x60;bidStrategy&#x60; in attach mode returns 400. To change an existing ad set&#39;s bid, use &#x60;PUT /v1/ads/ad-sets/{adSetId}&#x60;. Mutually exclusive with &#x60;creatives[]&#x60;.  The attached ad takes the full single-creative surface: &#x60;headline&#x60;/&#x60;body&#x60;/&#x60;description&#x60;/&#x60;callToAction&#x60; plus either &#x60;imageUrl&#x60;/&#x60;video&#x60; OR &#x60;placementAssets&#x60; (its own per-placement Feed/Story assets), and &#x60;leadGenFormId&#x60; when the target is a lead ad set (the parent must be ON_AD — true for ad sets created via goal &#x60;lead_generation&#x60;; Meta rejects a formless ad there, so pass the form on EVERY attached ad). This is the way to build N full ads sharing one ad set: create the first ad via the normal shape, then attach the rest one call each.  Supported on Meta (facebook, instagram) and TikTok. On TikTok the &#x60;adSetId&#x60; is the ad group ID; the new ad inherits the ad group&#39;s bid + budget + targeting.  | [optional] 
**ExistingCampaignId** | Pointer to **string** | Meta only. Add the new ad set under this EXISTING campaign instead of creating a new one (multi-ad-set audience testing). The new ad set&#39;s budget is matched to the campaign&#39;s mode automatically: for a CBO campaign (campaign-level budget) omit &#x60;budgetAmount&#x60;/&#x60;budgetType&#x60; — the campaign owns the budget; for an ABO campaign pass them (they go on the new ad set). On failure only the new ad set is cleaned up; the existing campaign is left untouched and is never (re)activated. Mutually exclusive with &#x60;adSetId&#x60; and &#x60;creatives[]&#x60;.  | [optional] 
**ExistingCreativeId** | Pointer to **string** | Meta only. Reuse an EXISTING ad creative by id instead of building a new one from the copy/media fields (which are then ignored). Combine with &#x60;existingCampaignId&#x60; to build a multi-ad-set campaign that shares one creative. Mutually exclusive with &#x60;creatives[]&#x60;, &#x60;dynamicCreative&#x60;, and &#x60;placementAssets&#x60;. The creative id used is returned as &#x60;creativeId&#x60; on the create response.  | [optional] 
**BusinessName** | Pointer to **string** | Google Display only | [optional] 
**BoardId** | Pointer to **string** | Pinterest only. Board ID (auto-creates if not provided). | [optional] 
**OrganizationId** | Pointer to **string** | LinkedIn only. The Company Page that authors the Direct Sponsored Content (\&quot;dark\&quot;) post backing the ad — accepts a numeric organization ID or a full &#x60;urn:li:organization:N&#x60; URN. Required unless the resolved &#x60;accountId&#x60; is a connected LinkedIn Company-Page account (defaults to that page) or the LinkedIn ad account is org-owned (defaults to the account&#39;s owning organization). The authenticated member must be an ADMINISTRATOR or DIRECT_SPONSORED_CONTENT_POSTER of this page (and the page must be associated with the ad account), or LinkedIn returns 403. Ignored by every other platform. | [optional] 
**Countries** | Pointer to **[]string** | ISO 3166-1 alpha-2 country codes (e.g. [&#39;NL&#39;]). Defaults to [&#39;US&#39;] when no &#x60;cities&#x60; or &#x60;regions&#x60; are provided. (LinkedIn currently honours country-level targeting only.) | [optional] 
**Cities** | Pointer to [**[]CreateStandaloneAdRequestCitiesInner**](CreateStandaloneAdRequestCitiesInner.md) | Meta-only. City-level geo targeting. Each city is targeted by Meta&#39;s opaque &#x60;key&#x60; (the city ID) which can be looked up via &#x60;GET /v1/ads/targeting/search?type&#x3D;city&amp;q&#x3D;&lt;name&gt;&amp;country_code&#x3D;&lt;ISO&gt;&#x60;. Optional &#x60;radius&#x60; + &#x60;distance_unit&#x60; extend the targeting beyond the city limits (e.g. radius 25 km around the city center). Both must be set together, or both omitted (Meta defaults to ~16 km when omitted).  Cannot overlap with the same country in &#x60;countries&#x60; (Meta returns a \&quot;locations overlap\&quot; error). Either drop the country or scope it to a different country.  | [optional] 
**Regions** | Pointer to [**[]CreateStandaloneAdRequestRegionsInner**](CreateStandaloneAdRequestRegionsInner.md) | Meta-only. Region-level (state/province) geo targeting. Each region is targeted by Meta&#39;s opaque &#x60;key&#x60; (the region ID) which can be looked up via &#x60;GET /v1/ads/targeting/search?type&#x3D;region&amp;q&#x3D;&lt;name&gt;&amp;country_code&#x3D;&lt;ISO&gt;&#x60;.  | [optional] 
**AgeMin** | Pointer to **int32** |  | [optional] 
**AgeMax** | Pointer to **int32** |  | [optional] 
**Interests** | Pointer to [**[]UpdateAdRequestTargetingInterestsInner**](UpdateAdRequestTargetingInterestsInner.md) | Interest objects from /v1/ads/interests. Each must include id and name. | [optional] 
**Zips** | Pointer to [**[]CreateStandaloneAdRequestZipsInner**](CreateStandaloneAdRequestZipsInner.md) | Postal/ZIP geo targeting. &#x60;key&#x60; is the platform&#39;s postal location ID from /v1/ads/targeting/search?dimension&#x3D;geo&amp;geoType&#x3D;zip. Supported on Meta, Google, TikTok, Pinterest, X. | [optional] 
**Metros** | Pointer to [**[]CreateStandaloneAdRequestZipsInner**](CreateStandaloneAdRequestZipsInner.md) | DMA / metro-area geo targeting. &#x60;key&#x60; is the platform&#39;s metro ID from /v1/ads/targeting/search?dimension&#x3D;geo&amp;geoType&#x3D;metro. | [optional] 
**CustomLocations** | Pointer to [**[]CreateStandaloneAdRequestCustomLocationsInner**](CreateStandaloneAdRequestCustomLocationsInner.md) | Point-radius (lat/lng) geo targeting. Meta only (custom_locations). Rejected on platforms without radius support. | [optional] 
**Behaviors** | Pointer to [**[]CreateStandaloneAdRequestBehaviorsInner**](CreateStandaloneAdRequestBehaviorsInner.md) | Behaviour entities from /v1/ads/targeting/search?dimension&#x3D;behavior. Supported on Meta and TikTok. Each must include id. | [optional] 
**IncomeTier** | Pointer to **string** | Normalized household-income tier. Meta and TikTok express all four; Google maps only &#x60;top_10&#x60;; rejected on LinkedIn, X, and Pinterest. On Meta, income targeting is incompatible with housing/employment/credit &#x60;specialAdCategories&#x60;.  | [optional] 
**Languages** | Pointer to **[]string** | Language codes (e.g. [&#39;en&#39;]). Restricts the audience by language. | [optional] 
**Placements** | Pointer to [**CreateStandaloneAdRequestPlacements**](CreateStandaloneAdRequestPlacements.md) |  | [optional] 
**SavedTargetingId** | Pointer to **string** | ID of a &#x60;saved_targeting&#x60; audience (created via POST /v1/ads/audiences). When set, its stored TargetingSpec is expanded as the base targeting; inline fields on this body merge on top. Lets you reuse a named targeting preset without re-sending every field.  | [optional] 
**RawTargeting** | Pointer to **map[string]interface{}** | Meta only. A raw Meta-native targeting spec passed to the ad set VERBATIM (snake_case: &#x60;geo_locations&#x60;, &#x60;age_min&#x60;, &#x60;excluded_custom_audiences&#x60;, &#x60;flexible_spec&#x60;, &#x60;targeting_automation&#x60;, business places, etc.) — exactly the shape &#x60;GET /v1/ads/{adId}&#x60; returns for external ads. Use it to clone a campaign&#39;s targeting EXACTLY, preserving advanced fields the camelCase targeting fields can&#39;t model. Mutually exclusive with the camelCase targeting fields (countries/regions/cities/interests/ ageMin/...), &#x60;audienceId&#x60;, and &#x60;savedTargetingId&#x60; (sending both → 422). Sent as-is; Meta validates and surfaces any errors. If cloning an EU campaign, also pass &#x60;dsaBeneficiary&#x60; / &#x60;dsaPayor&#x60; (those are separate fields, not part of targeting).  | [optional] 
**SpecialAdCategories** | Pointer to **[]string** | Meta only. Declares the ad&#39;s special category, required for housing, employment, credit, or political/social-issue ads (Meta enforces restricted targeting for these). Note: setting a special category disables income/zip targeting on Meta.  | [optional] 
**EndDate** | Pointer to **time.Time** | Required for lifetime budgets | [optional] 
**StartDate** | Pointer to **time.Time** | Meta only. Ad-set start time (ISO 8601, e.g. \&quot;2026-06-10T09:00:00Z\&quot;), mapped to the ad set&#39;s &#x60;start_time&#x60;. When omitted the ad starts delivering immediately. For lifetime budgets Meta also requires &#x60;endDate&#x60;. (Same &#x60;schedule.startDate&#x60; semantics already available on &#x60;POST /v1/ads/boost&#x60;.)  | [optional] 
**InstagramAccountId** | Pointer to **string** | Meta only. Override the Instagram account the ad is delivered as — pass an Instagram Business Account ID (e.g. 17841...), mapped to the creative&#39;s &#x60;instagram_user_id&#x60;. When omitted we auto-resolve the IG account linked to the connected Facebook Page (the existing default). Useful when a Page has more than one eligible IG account.  | [optional] 
**DynamicCreative** | Pointer to [**CreateStandaloneAdRequestDynamicCreative**](CreateStandaloneAdRequestDynamicCreative.md) |  | [optional] 
**PlacementAssets** | Pointer to [**CreateStandaloneAdRequestPlacementAssets**](CreateStandaloneAdRequestPlacementAssets.md) |  | [optional] 
**AudienceId** | Pointer to **string** | Custom audience ID for targeting | [optional] 
**CampaignType** | Pointer to **string** | Google only | [optional] [default to "display"]
**Keywords** | Pointer to **[]string** | Google Search only | [optional] 
**AdditionalHeadlines** | Pointer to **[]string** | Google Search RSA only. Extra headlines. | [optional] 
**AdditionalDescriptions** | Pointer to **[]string** | Google Search RSA only. Extra descriptions. | [optional] 
**AdvantageAudience** | Pointer to **int32** | Meta only. Controls the Advantage audience feature (targeting_automation). 0 &#x3D; disabled (default), 1 &#x3D; enabled. Meta Marketing API requires this field on all ad set creation requests. | [optional] 
**AttributionSpec** | Pointer to [**[]CreateStandaloneAdRequestAttributionSpecInner**](CreateStandaloneAdRequestAttributionSpecInner.md) | Meta only. Conversion attribution window for the ad set — maps 1:1 to Meta&#39;s ad-set &#x60;attribution_spec&#x60;. Only honored for conversion goals (&#x60;conversions&#x60;, &#x60;lead_generation&#x60;, &#x60;app_promotion&#x60;); ignored for awareness/traffic/engagement. Omit to use Meta&#39;s default (&#x60;7-day click&#x60; + &#x60;1-day view&#x60;). Meta enforces the valid combinations: &#x60;VIEW_THROUGH&#x60; only allows &#x60;windowDays: 1&#x60; (7d/28d view windows were removed Jan 2026); &#x60;ENGAGED_VIDEO_VIEW&#x60; only &#x60;1&#x60; and only alongside &#x60;VIEW_THROUGH: 1&#x60;; &#x60;CLICK_THROUGH: 28&#x60; only on certain objectives. Invalid combos surface as a Meta 400. Example: &#x60;[{ \&quot;eventType\&quot;: \&quot;CLICK_THROUGH\&quot;, \&quot;windowDays\&quot;: 7 }, { \&quot;eventType\&quot;: \&quot;VIEW_THROUGH\&quot;, \&quot;windowDays\&quot;: 1 }]&#x60;  | [optional] 
**Gender** | Pointer to **string** | Meta only. Restrict the audience by gender. &#39;male&#39; targets men only, &#39;female&#39; targets women only, &#39;all&#39; (default) targets everyone. Ignored by non-Meta platforms. | [optional] [default to "all"]
**BidStrategy** | Pointer to [**BidStrategy**](BidStrategy.md) | Meta bid strategy applied to the ad set. | [optional] 
**BidAmount** | Pointer to **float32** | Bid cap in WHOLE currency units (USD: 5 &#x3D; $5.00; JPY: 100 &#x3D; ¥100). Required when &#x60;bidStrategy&#x60; is &#x60;LOWEST_COST_WITH_BID_CAP&#x60; or &#x60;COST_CAP&#x60;.  | [optional] 
**RoasAverageFloor** | Pointer to **float32** | Minimum ROAS as a decimal multiplier (e.g. 2.0 &#x3D; 2.0x ROAS). Required when &#x60;bidStrategy&#x60; is &#x60;LOWEST_COST_WITH_MIN_ROAS&#x60;. Sent to Meta as &#x60;bid_constraints.roas_average_floor&#x60; × 10000.  | [optional] 
**DsaBeneficiary** | Pointer to **string** | Name of the legal entity benefiting from the ad. Required by Meta when targeting EU users (DSA Article 26). Not enforced at schema level; enforced server-side when targeting intersects EU member states.  | [optional] 
**DsaPayor** | Pointer to **string** | Name of the legal entity paying for the ad. Required by Meta when targeting EU users (DSA Article 26). Note Meta API spelling: dsa_payor (not dsa_payer).  | [optional] 
**BrandIdentity** | Pointer to [**CreateStandaloneAdRequestBrandIdentity**](CreateStandaloneAdRequestBrandIdentity.md) |  | [optional] 
**IdentityType** | Pointer to **string** | TikTok only. Forces the identity attribution on the ad:    - &#x60;TT_USER&#x60;: the posting account&#39;s open_id (real @username     branding). Requires a connected TikTok posting account     on the same profile.   - &#x60;CUSTOMIZED_USER&#x60;: synthetic Brand Identity (display     name + avatar). Requires a configured Brand Identity     (cached on the &#x60;tiktokads&#x60; SocialAccount via     &#x60;PATCH /v1/connect/tiktok-ads&#x60;) or an inline     &#x60;brandIdentity&#x60; to create one on the fly.  When omitted, defaults to &#x60;TT_USER&#x60; if a posting account is connected on this profile, else &#x60;CUSTOMIZED_USER&#x60;. Spark Ads (&#x60;POST /v1/ads/boost&#x60;) always use &#x60;TT_USER&#x60; regardless of this field — TikTok requires the original organic post&#39;s author identity for Spark.  | [optional] 
**PromotedObject** | Pointer to [**CreateStandaloneAdRequestPromotedObject**](CreateStandaloneAdRequestPromotedObject.md) |  | [optional] 

## Methods

### NewCreateStandaloneAdRequest

`func NewCreateStandaloneAdRequest(accountId string, adAccountId string, name string, ) *CreateStandaloneAdRequest`

NewCreateStandaloneAdRequest instantiates a new CreateStandaloneAdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandaloneAdRequestWithDefaults

`func NewCreateStandaloneAdRequestWithDefaults() *CreateStandaloneAdRequest`

NewCreateStandaloneAdRequestWithDefaults instantiates a new CreateStandaloneAdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateStandaloneAdRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateStandaloneAdRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateStandaloneAdRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAdAccountId

`func (o *CreateStandaloneAdRequest) GetAdAccountId() string`

GetAdAccountId returns the AdAccountId field if non-nil, zero value otherwise.

### GetAdAccountIdOk

`func (o *CreateStandaloneAdRequest) GetAdAccountIdOk() (*string, bool)`

GetAdAccountIdOk returns a tuple with the AdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountId

`func (o *CreateStandaloneAdRequest) SetAdAccountId(v string)`

SetAdAccountId sets AdAccountId field to given value.


### GetName

`func (o *CreateStandaloneAdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStandaloneAdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStandaloneAdRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCampaignName

`func (o *CreateStandaloneAdRequest) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *CreateStandaloneAdRequest) GetCampaignNameOk() (*string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignName

`func (o *CreateStandaloneAdRequest) SetCampaignName(v string)`

SetCampaignName sets CampaignName field to given value.

### HasCampaignName

`func (o *CreateStandaloneAdRequest) HasCampaignName() bool`

HasCampaignName returns a boolean if a field has been set.

### GetAdSetName

`func (o *CreateStandaloneAdRequest) GetAdSetName() string`

GetAdSetName returns the AdSetName field if non-nil, zero value otherwise.

### GetAdSetNameOk

`func (o *CreateStandaloneAdRequest) GetAdSetNameOk() (*string, bool)`

GetAdSetNameOk returns a tuple with the AdSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetName

`func (o *CreateStandaloneAdRequest) SetAdSetName(v string)`

SetAdSetName sets AdSetName field to given value.

### HasAdSetName

`func (o *CreateStandaloneAdRequest) HasAdSetName() bool`

HasAdSetName returns a boolean if a field has been set.

### GetAdName

`func (o *CreateStandaloneAdRequest) GetAdName() string`

GetAdName returns the AdName field if non-nil, zero value otherwise.

### GetAdNameOk

`func (o *CreateStandaloneAdRequest) GetAdNameOk() (*string, bool)`

GetAdNameOk returns a tuple with the AdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdName

`func (o *CreateStandaloneAdRequest) SetAdName(v string)`

SetAdName sets AdName field to given value.

### HasAdName

`func (o *CreateStandaloneAdRequest) HasAdName() bool`

HasAdName returns a boolean if a field has been set.

### GetTracking

`func (o *CreateStandaloneAdRequest) GetTracking() CreateStandaloneAdRequestTracking`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *CreateStandaloneAdRequest) GetTrackingOk() (*CreateStandaloneAdRequestTracking, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *CreateStandaloneAdRequest) SetTracking(v CreateStandaloneAdRequestTracking)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *CreateStandaloneAdRequest) HasTracking() bool`

HasTracking returns a boolean if a field has been set.

### GetGoal

`func (o *CreateStandaloneAdRequest) GetGoal() string`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *CreateStandaloneAdRequest) GetGoalOk() (*string, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *CreateStandaloneAdRequest) SetGoal(v string)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *CreateStandaloneAdRequest) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### GetOptimizationGoal

`func (o *CreateStandaloneAdRequest) GetOptimizationGoal() string`

GetOptimizationGoal returns the OptimizationGoal field if non-nil, zero value otherwise.

### GetOptimizationGoalOk

`func (o *CreateStandaloneAdRequest) GetOptimizationGoalOk() (*string, bool)`

GetOptimizationGoalOk returns a tuple with the OptimizationGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizationGoal

`func (o *CreateStandaloneAdRequest) SetOptimizationGoal(v string)`

SetOptimizationGoal sets OptimizationGoal field to given value.

### HasOptimizationGoal

`func (o *CreateStandaloneAdRequest) HasOptimizationGoal() bool`

HasOptimizationGoal returns a boolean if a field has been set.

### GetBudgetAmount

`func (o *CreateStandaloneAdRequest) GetBudgetAmount() float32`

GetBudgetAmount returns the BudgetAmount field if non-nil, zero value otherwise.

### GetBudgetAmountOk

`func (o *CreateStandaloneAdRequest) GetBudgetAmountOk() (*float32, bool)`

GetBudgetAmountOk returns a tuple with the BudgetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetAmount

`func (o *CreateStandaloneAdRequest) SetBudgetAmount(v float32)`

SetBudgetAmount sets BudgetAmount field to given value.

### HasBudgetAmount

`func (o *CreateStandaloneAdRequest) HasBudgetAmount() bool`

HasBudgetAmount returns a boolean if a field has been set.

### GetBudgetType

`func (o *CreateStandaloneAdRequest) GetBudgetType() string`

GetBudgetType returns the BudgetType field if non-nil, zero value otherwise.

### GetBudgetTypeOk

`func (o *CreateStandaloneAdRequest) GetBudgetTypeOk() (*string, bool)`

GetBudgetTypeOk returns a tuple with the BudgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetType

`func (o *CreateStandaloneAdRequest) SetBudgetType(v string)`

SetBudgetType sets BudgetType field to given value.

### HasBudgetType

`func (o *CreateStandaloneAdRequest) HasBudgetType() bool`

HasBudgetType returns a boolean if a field has been set.

### GetStatus

`func (o *CreateStandaloneAdRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateStandaloneAdRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateStandaloneAdRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateStandaloneAdRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBudgetLevel

`func (o *CreateStandaloneAdRequest) GetBudgetLevel() string`

GetBudgetLevel returns the BudgetLevel field if non-nil, zero value otherwise.

### GetBudgetLevelOk

`func (o *CreateStandaloneAdRequest) GetBudgetLevelOk() (*string, bool)`

GetBudgetLevelOk returns a tuple with the BudgetLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetLevel

`func (o *CreateStandaloneAdRequest) SetBudgetLevel(v string)`

SetBudgetLevel sets BudgetLevel field to given value.

### HasBudgetLevel

`func (o *CreateStandaloneAdRequest) HasBudgetLevel() bool`

HasBudgetLevel returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateStandaloneAdRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateStandaloneAdRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateStandaloneAdRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateStandaloneAdRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetHeadline

`func (o *CreateStandaloneAdRequest) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *CreateStandaloneAdRequest) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *CreateStandaloneAdRequest) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *CreateStandaloneAdRequest) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetLongHeadline

`func (o *CreateStandaloneAdRequest) GetLongHeadline() string`

GetLongHeadline returns the LongHeadline field if non-nil, zero value otherwise.

### GetLongHeadlineOk

`func (o *CreateStandaloneAdRequest) GetLongHeadlineOk() (*string, bool)`

GetLongHeadlineOk returns a tuple with the LongHeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongHeadline

`func (o *CreateStandaloneAdRequest) SetLongHeadline(v string)`

SetLongHeadline sets LongHeadline field to given value.

### HasLongHeadline

`func (o *CreateStandaloneAdRequest) HasLongHeadline() bool`

HasLongHeadline returns a boolean if a field has been set.

### GetBody

`func (o *CreateStandaloneAdRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateStandaloneAdRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateStandaloneAdRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateStandaloneAdRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDescription

`func (o *CreateStandaloneAdRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStandaloneAdRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStandaloneAdRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateStandaloneAdRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCallToAction

`func (o *CreateStandaloneAdRequest) GetCallToAction() string`

GetCallToAction returns the CallToAction field if non-nil, zero value otherwise.

### GetCallToActionOk

`func (o *CreateStandaloneAdRequest) GetCallToActionOk() (*string, bool)`

GetCallToActionOk returns a tuple with the CallToAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallToAction

`func (o *CreateStandaloneAdRequest) SetCallToAction(v string)`

SetCallToAction sets CallToAction field to given value.

### HasCallToAction

`func (o *CreateStandaloneAdRequest) HasCallToAction() bool`

HasCallToAction returns a boolean if a field has been set.

### GetLinkUrl

`func (o *CreateStandaloneAdRequest) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *CreateStandaloneAdRequest) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *CreateStandaloneAdRequest) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.

### HasLinkUrl

`func (o *CreateStandaloneAdRequest) HasLinkUrl() bool`

HasLinkUrl returns a boolean if a field has been set.

### GetLeadGenFormId

`func (o *CreateStandaloneAdRequest) GetLeadGenFormId() string`

GetLeadGenFormId returns the LeadGenFormId field if non-nil, zero value otherwise.

### GetLeadGenFormIdOk

`func (o *CreateStandaloneAdRequest) GetLeadGenFormIdOk() (*string, bool)`

GetLeadGenFormIdOk returns a tuple with the LeadGenFormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadGenFormId

`func (o *CreateStandaloneAdRequest) SetLeadGenFormId(v string)`

SetLeadGenFormId sets LeadGenFormId field to given value.

### HasLeadGenFormId

`func (o *CreateStandaloneAdRequest) HasLeadGenFormId() bool`

HasLeadGenFormId returns a boolean if a field has been set.

### GetImageUrl

`func (o *CreateStandaloneAdRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateStandaloneAdRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateStandaloneAdRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateStandaloneAdRequest) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetImages

`func (o *CreateStandaloneAdRequest) GetImages() CreateStandaloneAdRequestImages`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CreateStandaloneAdRequest) GetImagesOk() (*CreateStandaloneAdRequestImages, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CreateStandaloneAdRequest) SetImages(v CreateStandaloneAdRequestImages)`

SetImages sets Images field to given value.

### HasImages

`func (o *CreateStandaloneAdRequest) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetVideo

`func (o *CreateStandaloneAdRequest) GetVideo() CreateStandaloneAdRequestVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *CreateStandaloneAdRequest) GetVideoOk() (*CreateStandaloneAdRequestVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *CreateStandaloneAdRequest) SetVideo(v CreateStandaloneAdRequestVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *CreateStandaloneAdRequest) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetCreatives

`func (o *CreateStandaloneAdRequest) GetCreatives() []CreateStandaloneAdRequestCreativesInner`

GetCreatives returns the Creatives field if non-nil, zero value otherwise.

### GetCreativesOk

`func (o *CreateStandaloneAdRequest) GetCreativesOk() (*[]CreateStandaloneAdRequestCreativesInner, bool)`

GetCreativesOk returns a tuple with the Creatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatives

`func (o *CreateStandaloneAdRequest) SetCreatives(v []CreateStandaloneAdRequestCreativesInner)`

SetCreatives sets Creatives field to given value.

### HasCreatives

`func (o *CreateStandaloneAdRequest) HasCreatives() bool`

HasCreatives returns a boolean if a field has been set.

### GetAdSetId

`func (o *CreateStandaloneAdRequest) GetAdSetId() string`

GetAdSetId returns the AdSetId field if non-nil, zero value otherwise.

### GetAdSetIdOk

`func (o *CreateStandaloneAdRequest) GetAdSetIdOk() (*string, bool)`

GetAdSetIdOk returns a tuple with the AdSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetId

`func (o *CreateStandaloneAdRequest) SetAdSetId(v string)`

SetAdSetId sets AdSetId field to given value.

### HasAdSetId

`func (o *CreateStandaloneAdRequest) HasAdSetId() bool`

HasAdSetId returns a boolean if a field has been set.

### GetExistingCampaignId

`func (o *CreateStandaloneAdRequest) GetExistingCampaignId() string`

GetExistingCampaignId returns the ExistingCampaignId field if non-nil, zero value otherwise.

### GetExistingCampaignIdOk

`func (o *CreateStandaloneAdRequest) GetExistingCampaignIdOk() (*string, bool)`

GetExistingCampaignIdOk returns a tuple with the ExistingCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingCampaignId

`func (o *CreateStandaloneAdRequest) SetExistingCampaignId(v string)`

SetExistingCampaignId sets ExistingCampaignId field to given value.

### HasExistingCampaignId

`func (o *CreateStandaloneAdRequest) HasExistingCampaignId() bool`

HasExistingCampaignId returns a boolean if a field has been set.

### GetExistingCreativeId

`func (o *CreateStandaloneAdRequest) GetExistingCreativeId() string`

GetExistingCreativeId returns the ExistingCreativeId field if non-nil, zero value otherwise.

### GetExistingCreativeIdOk

`func (o *CreateStandaloneAdRequest) GetExistingCreativeIdOk() (*string, bool)`

GetExistingCreativeIdOk returns a tuple with the ExistingCreativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingCreativeId

`func (o *CreateStandaloneAdRequest) SetExistingCreativeId(v string)`

SetExistingCreativeId sets ExistingCreativeId field to given value.

### HasExistingCreativeId

`func (o *CreateStandaloneAdRequest) HasExistingCreativeId() bool`

HasExistingCreativeId returns a boolean if a field has been set.

### GetBusinessName

`func (o *CreateStandaloneAdRequest) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *CreateStandaloneAdRequest) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *CreateStandaloneAdRequest) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *CreateStandaloneAdRequest) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetBoardId

`func (o *CreateStandaloneAdRequest) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *CreateStandaloneAdRequest) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *CreateStandaloneAdRequest) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *CreateStandaloneAdRequest) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *CreateStandaloneAdRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateStandaloneAdRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateStandaloneAdRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateStandaloneAdRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetCountries

`func (o *CreateStandaloneAdRequest) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *CreateStandaloneAdRequest) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *CreateStandaloneAdRequest) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *CreateStandaloneAdRequest) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetCities

`func (o *CreateStandaloneAdRequest) GetCities() []CreateStandaloneAdRequestCitiesInner`

GetCities returns the Cities field if non-nil, zero value otherwise.

### GetCitiesOk

`func (o *CreateStandaloneAdRequest) GetCitiesOk() (*[]CreateStandaloneAdRequestCitiesInner, bool)`

GetCitiesOk returns a tuple with the Cities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCities

`func (o *CreateStandaloneAdRequest) SetCities(v []CreateStandaloneAdRequestCitiesInner)`

SetCities sets Cities field to given value.

### HasCities

`func (o *CreateStandaloneAdRequest) HasCities() bool`

HasCities returns a boolean if a field has been set.

### GetRegions

`func (o *CreateStandaloneAdRequest) GetRegions() []CreateStandaloneAdRequestRegionsInner`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CreateStandaloneAdRequest) GetRegionsOk() (*[]CreateStandaloneAdRequestRegionsInner, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CreateStandaloneAdRequest) SetRegions(v []CreateStandaloneAdRequestRegionsInner)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CreateStandaloneAdRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetAgeMin

`func (o *CreateStandaloneAdRequest) GetAgeMin() int32`

GetAgeMin returns the AgeMin field if non-nil, zero value otherwise.

### GetAgeMinOk

`func (o *CreateStandaloneAdRequest) GetAgeMinOk() (*int32, bool)`

GetAgeMinOk returns a tuple with the AgeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeMin

`func (o *CreateStandaloneAdRequest) SetAgeMin(v int32)`

SetAgeMin sets AgeMin field to given value.

### HasAgeMin

`func (o *CreateStandaloneAdRequest) HasAgeMin() bool`

HasAgeMin returns a boolean if a field has been set.

### GetAgeMax

`func (o *CreateStandaloneAdRequest) GetAgeMax() int32`

GetAgeMax returns the AgeMax field if non-nil, zero value otherwise.

### GetAgeMaxOk

`func (o *CreateStandaloneAdRequest) GetAgeMaxOk() (*int32, bool)`

GetAgeMaxOk returns a tuple with the AgeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeMax

`func (o *CreateStandaloneAdRequest) SetAgeMax(v int32)`

SetAgeMax sets AgeMax field to given value.

### HasAgeMax

`func (o *CreateStandaloneAdRequest) HasAgeMax() bool`

HasAgeMax returns a boolean if a field has been set.

### GetInterests

`func (o *CreateStandaloneAdRequest) GetInterests() []UpdateAdRequestTargetingInterestsInner`

GetInterests returns the Interests field if non-nil, zero value otherwise.

### GetInterestsOk

`func (o *CreateStandaloneAdRequest) GetInterestsOk() (*[]UpdateAdRequestTargetingInterestsInner, bool)`

GetInterestsOk returns a tuple with the Interests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterests

`func (o *CreateStandaloneAdRequest) SetInterests(v []UpdateAdRequestTargetingInterestsInner)`

SetInterests sets Interests field to given value.

### HasInterests

`func (o *CreateStandaloneAdRequest) HasInterests() bool`

HasInterests returns a boolean if a field has been set.

### GetZips

`func (o *CreateStandaloneAdRequest) GetZips() []CreateStandaloneAdRequestZipsInner`

GetZips returns the Zips field if non-nil, zero value otherwise.

### GetZipsOk

`func (o *CreateStandaloneAdRequest) GetZipsOk() (*[]CreateStandaloneAdRequestZipsInner, bool)`

GetZipsOk returns a tuple with the Zips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZips

`func (o *CreateStandaloneAdRequest) SetZips(v []CreateStandaloneAdRequestZipsInner)`

SetZips sets Zips field to given value.

### HasZips

`func (o *CreateStandaloneAdRequest) HasZips() bool`

HasZips returns a boolean if a field has been set.

### GetMetros

`func (o *CreateStandaloneAdRequest) GetMetros() []CreateStandaloneAdRequestZipsInner`

GetMetros returns the Metros field if non-nil, zero value otherwise.

### GetMetrosOk

`func (o *CreateStandaloneAdRequest) GetMetrosOk() (*[]CreateStandaloneAdRequestZipsInner, bool)`

GetMetrosOk returns a tuple with the Metros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetros

`func (o *CreateStandaloneAdRequest) SetMetros(v []CreateStandaloneAdRequestZipsInner)`

SetMetros sets Metros field to given value.

### HasMetros

`func (o *CreateStandaloneAdRequest) HasMetros() bool`

HasMetros returns a boolean if a field has been set.

### GetCustomLocations

`func (o *CreateStandaloneAdRequest) GetCustomLocations() []CreateStandaloneAdRequestCustomLocationsInner`

GetCustomLocations returns the CustomLocations field if non-nil, zero value otherwise.

### GetCustomLocationsOk

`func (o *CreateStandaloneAdRequest) GetCustomLocationsOk() (*[]CreateStandaloneAdRequestCustomLocationsInner, bool)`

GetCustomLocationsOk returns a tuple with the CustomLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLocations

`func (o *CreateStandaloneAdRequest) SetCustomLocations(v []CreateStandaloneAdRequestCustomLocationsInner)`

SetCustomLocations sets CustomLocations field to given value.

### HasCustomLocations

`func (o *CreateStandaloneAdRequest) HasCustomLocations() bool`

HasCustomLocations returns a boolean if a field has been set.

### GetBehaviors

`func (o *CreateStandaloneAdRequest) GetBehaviors() []CreateStandaloneAdRequestBehaviorsInner`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *CreateStandaloneAdRequest) GetBehaviorsOk() (*[]CreateStandaloneAdRequestBehaviorsInner, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *CreateStandaloneAdRequest) SetBehaviors(v []CreateStandaloneAdRequestBehaviorsInner)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *CreateStandaloneAdRequest) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetIncomeTier

`func (o *CreateStandaloneAdRequest) GetIncomeTier() string`

GetIncomeTier returns the IncomeTier field if non-nil, zero value otherwise.

### GetIncomeTierOk

`func (o *CreateStandaloneAdRequest) GetIncomeTierOk() (*string, bool)`

GetIncomeTierOk returns a tuple with the IncomeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeTier

`func (o *CreateStandaloneAdRequest) SetIncomeTier(v string)`

SetIncomeTier sets IncomeTier field to given value.

### HasIncomeTier

`func (o *CreateStandaloneAdRequest) HasIncomeTier() bool`

HasIncomeTier returns a boolean if a field has been set.

### GetLanguages

`func (o *CreateStandaloneAdRequest) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *CreateStandaloneAdRequest) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *CreateStandaloneAdRequest) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *CreateStandaloneAdRequest) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetPlacements

`func (o *CreateStandaloneAdRequest) GetPlacements() CreateStandaloneAdRequestPlacements`

GetPlacements returns the Placements field if non-nil, zero value otherwise.

### GetPlacementsOk

`func (o *CreateStandaloneAdRequest) GetPlacementsOk() (*CreateStandaloneAdRequestPlacements, bool)`

GetPlacementsOk returns a tuple with the Placements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacements

`func (o *CreateStandaloneAdRequest) SetPlacements(v CreateStandaloneAdRequestPlacements)`

SetPlacements sets Placements field to given value.

### HasPlacements

`func (o *CreateStandaloneAdRequest) HasPlacements() bool`

HasPlacements returns a boolean if a field has been set.

### GetSavedTargetingId

`func (o *CreateStandaloneAdRequest) GetSavedTargetingId() string`

GetSavedTargetingId returns the SavedTargetingId field if non-nil, zero value otherwise.

### GetSavedTargetingIdOk

`func (o *CreateStandaloneAdRequest) GetSavedTargetingIdOk() (*string, bool)`

GetSavedTargetingIdOk returns a tuple with the SavedTargetingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedTargetingId

`func (o *CreateStandaloneAdRequest) SetSavedTargetingId(v string)`

SetSavedTargetingId sets SavedTargetingId field to given value.

### HasSavedTargetingId

`func (o *CreateStandaloneAdRequest) HasSavedTargetingId() bool`

HasSavedTargetingId returns a boolean if a field has been set.

### GetRawTargeting

`func (o *CreateStandaloneAdRequest) GetRawTargeting() map[string]interface{}`

GetRawTargeting returns the RawTargeting field if non-nil, zero value otherwise.

### GetRawTargetingOk

`func (o *CreateStandaloneAdRequest) GetRawTargetingOk() (*map[string]interface{}, bool)`

GetRawTargetingOk returns a tuple with the RawTargeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTargeting

`func (o *CreateStandaloneAdRequest) SetRawTargeting(v map[string]interface{})`

SetRawTargeting sets RawTargeting field to given value.

### HasRawTargeting

`func (o *CreateStandaloneAdRequest) HasRawTargeting() bool`

HasRawTargeting returns a boolean if a field has been set.

### GetSpecialAdCategories

`func (o *CreateStandaloneAdRequest) GetSpecialAdCategories() []string`

GetSpecialAdCategories returns the SpecialAdCategories field if non-nil, zero value otherwise.

### GetSpecialAdCategoriesOk

`func (o *CreateStandaloneAdRequest) GetSpecialAdCategoriesOk() (*[]string, bool)`

GetSpecialAdCategoriesOk returns a tuple with the SpecialAdCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialAdCategories

`func (o *CreateStandaloneAdRequest) SetSpecialAdCategories(v []string)`

SetSpecialAdCategories sets SpecialAdCategories field to given value.

### HasSpecialAdCategories

`func (o *CreateStandaloneAdRequest) HasSpecialAdCategories() bool`

HasSpecialAdCategories returns a boolean if a field has been set.

### GetEndDate

`func (o *CreateStandaloneAdRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateStandaloneAdRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateStandaloneAdRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateStandaloneAdRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateStandaloneAdRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateStandaloneAdRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateStandaloneAdRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateStandaloneAdRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetInstagramAccountId

`func (o *CreateStandaloneAdRequest) GetInstagramAccountId() string`

GetInstagramAccountId returns the InstagramAccountId field if non-nil, zero value otherwise.

### GetInstagramAccountIdOk

`func (o *CreateStandaloneAdRequest) GetInstagramAccountIdOk() (*string, bool)`

GetInstagramAccountIdOk returns a tuple with the InstagramAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramAccountId

`func (o *CreateStandaloneAdRequest) SetInstagramAccountId(v string)`

SetInstagramAccountId sets InstagramAccountId field to given value.

### HasInstagramAccountId

`func (o *CreateStandaloneAdRequest) HasInstagramAccountId() bool`

HasInstagramAccountId returns a boolean if a field has been set.

### GetDynamicCreative

`func (o *CreateStandaloneAdRequest) GetDynamicCreative() CreateStandaloneAdRequestDynamicCreative`

GetDynamicCreative returns the DynamicCreative field if non-nil, zero value otherwise.

### GetDynamicCreativeOk

`func (o *CreateStandaloneAdRequest) GetDynamicCreativeOk() (*CreateStandaloneAdRequestDynamicCreative, bool)`

GetDynamicCreativeOk returns a tuple with the DynamicCreative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicCreative

`func (o *CreateStandaloneAdRequest) SetDynamicCreative(v CreateStandaloneAdRequestDynamicCreative)`

SetDynamicCreative sets DynamicCreative field to given value.

### HasDynamicCreative

`func (o *CreateStandaloneAdRequest) HasDynamicCreative() bool`

HasDynamicCreative returns a boolean if a field has been set.

### GetPlacementAssets

`func (o *CreateStandaloneAdRequest) GetPlacementAssets() CreateStandaloneAdRequestPlacementAssets`

GetPlacementAssets returns the PlacementAssets field if non-nil, zero value otherwise.

### GetPlacementAssetsOk

`func (o *CreateStandaloneAdRequest) GetPlacementAssetsOk() (*CreateStandaloneAdRequestPlacementAssets, bool)`

GetPlacementAssetsOk returns a tuple with the PlacementAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementAssets

`func (o *CreateStandaloneAdRequest) SetPlacementAssets(v CreateStandaloneAdRequestPlacementAssets)`

SetPlacementAssets sets PlacementAssets field to given value.

### HasPlacementAssets

`func (o *CreateStandaloneAdRequest) HasPlacementAssets() bool`

HasPlacementAssets returns a boolean if a field has been set.

### GetAudienceId

`func (o *CreateStandaloneAdRequest) GetAudienceId() string`

GetAudienceId returns the AudienceId field if non-nil, zero value otherwise.

### GetAudienceIdOk

`func (o *CreateStandaloneAdRequest) GetAudienceIdOk() (*string, bool)`

GetAudienceIdOk returns a tuple with the AudienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceId

`func (o *CreateStandaloneAdRequest) SetAudienceId(v string)`

SetAudienceId sets AudienceId field to given value.

### HasAudienceId

`func (o *CreateStandaloneAdRequest) HasAudienceId() bool`

HasAudienceId returns a boolean if a field has been set.

### GetCampaignType

`func (o *CreateStandaloneAdRequest) GetCampaignType() string`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *CreateStandaloneAdRequest) GetCampaignTypeOk() (*string, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignType

`func (o *CreateStandaloneAdRequest) SetCampaignType(v string)`

SetCampaignType sets CampaignType field to given value.

### HasCampaignType

`func (o *CreateStandaloneAdRequest) HasCampaignType() bool`

HasCampaignType returns a boolean if a field has been set.

### GetKeywords

`func (o *CreateStandaloneAdRequest) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *CreateStandaloneAdRequest) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *CreateStandaloneAdRequest) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *CreateStandaloneAdRequest) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetAdditionalHeadlines

`func (o *CreateStandaloneAdRequest) GetAdditionalHeadlines() []string`

GetAdditionalHeadlines returns the AdditionalHeadlines field if non-nil, zero value otherwise.

### GetAdditionalHeadlinesOk

`func (o *CreateStandaloneAdRequest) GetAdditionalHeadlinesOk() (*[]string, bool)`

GetAdditionalHeadlinesOk returns a tuple with the AdditionalHeadlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHeadlines

`func (o *CreateStandaloneAdRequest) SetAdditionalHeadlines(v []string)`

SetAdditionalHeadlines sets AdditionalHeadlines field to given value.

### HasAdditionalHeadlines

`func (o *CreateStandaloneAdRequest) HasAdditionalHeadlines() bool`

HasAdditionalHeadlines returns a boolean if a field has been set.

### GetAdditionalDescriptions

`func (o *CreateStandaloneAdRequest) GetAdditionalDescriptions() []string`

GetAdditionalDescriptions returns the AdditionalDescriptions field if non-nil, zero value otherwise.

### GetAdditionalDescriptionsOk

`func (o *CreateStandaloneAdRequest) GetAdditionalDescriptionsOk() (*[]string, bool)`

GetAdditionalDescriptionsOk returns a tuple with the AdditionalDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDescriptions

`func (o *CreateStandaloneAdRequest) SetAdditionalDescriptions(v []string)`

SetAdditionalDescriptions sets AdditionalDescriptions field to given value.

### HasAdditionalDescriptions

`func (o *CreateStandaloneAdRequest) HasAdditionalDescriptions() bool`

HasAdditionalDescriptions returns a boolean if a field has been set.

### GetAdvantageAudience

`func (o *CreateStandaloneAdRequest) GetAdvantageAudience() int32`

GetAdvantageAudience returns the AdvantageAudience field if non-nil, zero value otherwise.

### GetAdvantageAudienceOk

`func (o *CreateStandaloneAdRequest) GetAdvantageAudienceOk() (*int32, bool)`

GetAdvantageAudienceOk returns a tuple with the AdvantageAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvantageAudience

`func (o *CreateStandaloneAdRequest) SetAdvantageAudience(v int32)`

SetAdvantageAudience sets AdvantageAudience field to given value.

### HasAdvantageAudience

`func (o *CreateStandaloneAdRequest) HasAdvantageAudience() bool`

HasAdvantageAudience returns a boolean if a field has been set.

### GetAttributionSpec

`func (o *CreateStandaloneAdRequest) GetAttributionSpec() []CreateStandaloneAdRequestAttributionSpecInner`

GetAttributionSpec returns the AttributionSpec field if non-nil, zero value otherwise.

### GetAttributionSpecOk

`func (o *CreateStandaloneAdRequest) GetAttributionSpecOk() (*[]CreateStandaloneAdRequestAttributionSpecInner, bool)`

GetAttributionSpecOk returns a tuple with the AttributionSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionSpec

`func (o *CreateStandaloneAdRequest) SetAttributionSpec(v []CreateStandaloneAdRequestAttributionSpecInner)`

SetAttributionSpec sets AttributionSpec field to given value.

### HasAttributionSpec

`func (o *CreateStandaloneAdRequest) HasAttributionSpec() bool`

HasAttributionSpec returns a boolean if a field has been set.

### GetGender

`func (o *CreateStandaloneAdRequest) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CreateStandaloneAdRequest) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CreateStandaloneAdRequest) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CreateStandaloneAdRequest) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetBidStrategy

`func (o *CreateStandaloneAdRequest) GetBidStrategy() BidStrategy`

GetBidStrategy returns the BidStrategy field if non-nil, zero value otherwise.

### GetBidStrategyOk

`func (o *CreateStandaloneAdRequest) GetBidStrategyOk() (*BidStrategy, bool)`

GetBidStrategyOk returns a tuple with the BidStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidStrategy

`func (o *CreateStandaloneAdRequest) SetBidStrategy(v BidStrategy)`

SetBidStrategy sets BidStrategy field to given value.

### HasBidStrategy

`func (o *CreateStandaloneAdRequest) HasBidStrategy() bool`

HasBidStrategy returns a boolean if a field has been set.

### GetBidAmount

`func (o *CreateStandaloneAdRequest) GetBidAmount() float32`

GetBidAmount returns the BidAmount field if non-nil, zero value otherwise.

### GetBidAmountOk

`func (o *CreateStandaloneAdRequest) GetBidAmountOk() (*float32, bool)`

GetBidAmountOk returns a tuple with the BidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidAmount

`func (o *CreateStandaloneAdRequest) SetBidAmount(v float32)`

SetBidAmount sets BidAmount field to given value.

### HasBidAmount

`func (o *CreateStandaloneAdRequest) HasBidAmount() bool`

HasBidAmount returns a boolean if a field has been set.

### GetRoasAverageFloor

`func (o *CreateStandaloneAdRequest) GetRoasAverageFloor() float32`

GetRoasAverageFloor returns the RoasAverageFloor field if non-nil, zero value otherwise.

### GetRoasAverageFloorOk

`func (o *CreateStandaloneAdRequest) GetRoasAverageFloorOk() (*float32, bool)`

GetRoasAverageFloorOk returns a tuple with the RoasAverageFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoasAverageFloor

`func (o *CreateStandaloneAdRequest) SetRoasAverageFloor(v float32)`

SetRoasAverageFloor sets RoasAverageFloor field to given value.

### HasRoasAverageFloor

`func (o *CreateStandaloneAdRequest) HasRoasAverageFloor() bool`

HasRoasAverageFloor returns a boolean if a field has been set.

### GetDsaBeneficiary

`func (o *CreateStandaloneAdRequest) GetDsaBeneficiary() string`

GetDsaBeneficiary returns the DsaBeneficiary field if non-nil, zero value otherwise.

### GetDsaBeneficiaryOk

`func (o *CreateStandaloneAdRequest) GetDsaBeneficiaryOk() (*string, bool)`

GetDsaBeneficiaryOk returns a tuple with the DsaBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsaBeneficiary

`func (o *CreateStandaloneAdRequest) SetDsaBeneficiary(v string)`

SetDsaBeneficiary sets DsaBeneficiary field to given value.

### HasDsaBeneficiary

`func (o *CreateStandaloneAdRequest) HasDsaBeneficiary() bool`

HasDsaBeneficiary returns a boolean if a field has been set.

### GetDsaPayor

`func (o *CreateStandaloneAdRequest) GetDsaPayor() string`

GetDsaPayor returns the DsaPayor field if non-nil, zero value otherwise.

### GetDsaPayorOk

`func (o *CreateStandaloneAdRequest) GetDsaPayorOk() (*string, bool)`

GetDsaPayorOk returns a tuple with the DsaPayor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsaPayor

`func (o *CreateStandaloneAdRequest) SetDsaPayor(v string)`

SetDsaPayor sets DsaPayor field to given value.

### HasDsaPayor

`func (o *CreateStandaloneAdRequest) HasDsaPayor() bool`

HasDsaPayor returns a boolean if a field has been set.

### GetBrandIdentity

`func (o *CreateStandaloneAdRequest) GetBrandIdentity() CreateStandaloneAdRequestBrandIdentity`

GetBrandIdentity returns the BrandIdentity field if non-nil, zero value otherwise.

### GetBrandIdentityOk

`func (o *CreateStandaloneAdRequest) GetBrandIdentityOk() (*CreateStandaloneAdRequestBrandIdentity, bool)`

GetBrandIdentityOk returns a tuple with the BrandIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandIdentity

`func (o *CreateStandaloneAdRequest) SetBrandIdentity(v CreateStandaloneAdRequestBrandIdentity)`

SetBrandIdentity sets BrandIdentity field to given value.

### HasBrandIdentity

`func (o *CreateStandaloneAdRequest) HasBrandIdentity() bool`

HasBrandIdentity returns a boolean if a field has been set.

### GetIdentityType

`func (o *CreateStandaloneAdRequest) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *CreateStandaloneAdRequest) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *CreateStandaloneAdRequest) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *CreateStandaloneAdRequest) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### GetPromotedObject

`func (o *CreateStandaloneAdRequest) GetPromotedObject() CreateStandaloneAdRequestPromotedObject`

GetPromotedObject returns the PromotedObject field if non-nil, zero value otherwise.

### GetPromotedObjectOk

`func (o *CreateStandaloneAdRequest) GetPromotedObjectOk() (*CreateStandaloneAdRequestPromotedObject, bool)`

GetPromotedObjectOk returns a tuple with the PromotedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedObject

`func (o *CreateStandaloneAdRequest) SetPromotedObject(v CreateStandaloneAdRequestPromotedObject)`

SetPromotedObject sets PromotedObject field to given value.

### HasPromotedObject

`func (o *CreateStandaloneAdRequest) HasPromotedObject() bool`

HasPromotedObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


