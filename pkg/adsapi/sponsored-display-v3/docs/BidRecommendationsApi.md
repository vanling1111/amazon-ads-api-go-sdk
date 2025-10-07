# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTargetBidRecommendations**](BidRecommendationsApi.md#GetTargetBidRecommendations) | **Post** /sd/targets/bid/recommendations | Returns a set of bid recommendations for targeting clauses

# **GetTargetBidRecommendations**
> SdTargetingBidRecommendationsResponseV32 GetTargetBidRecommendations(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Returns a set of bid recommendations for targeting clauses

Provides a list of bid recommendations based on the list of input advertised ASINs and targeting clauses in the same format as the targeting API. For each targeting clause in the request a corresponding bid recommendation will be returned in the response. Currently the API will accept up to 100 targeting clauses. Note - these recommendations are only available when productAds have ASIN or SKU fields. This API provides a corresponding bid recommendation for each targeting clause in a request. Currently the API will accept up to 100 targeting clauses.  For API v3.1, the API provides a list of bid recommendations based on the list of input advertised ASINs and targeting clauses in the same format as the targeting API. Note - these recommendations are only available when productAds have ASIN or SKU fields.  For API v3.2, the API accepts optimizationType and costType and returns bid recommendations accordingly.  For API v3.3, the API accepts creativeType and returns bid recommendations accordingly.  [PREVIEW ONLY] For API v3.4, the API supports entertainment targeting bid recommendations which is currently limited to US marketplace.  The recommended bids are derived from the last 7 days of winning auction bids for the related targeting clause.   Receive bid recommendations using the following: Contextual targeting clause|Description| |-----------|----| |asinSameAs=B0123456789|Receive a bid recommendation for this target product |asinCategorySameAs=12345|Receive a bid recommendation for this target category |similarProduct|Receive a bid recommendation for targets that are similar to the advertised asins.   Audience targeting clause|Description| |-----------|----| |views(asinCategorySameAs=12345 lookback=30)|Receive a bid recommendation for a target audience that has viewed products in the given category |views(similarProduct lookback=30)|Receive a bid recommendation for a target audience that has viewed similar products to the advertised asins |views(exactProduct lookback=30)|Receive a bid recommendation for a target audience that has viewed the advertised asins |purchases(asinCategorySameAs=12345 lookback=30)|Receive a bid recommendation for a target audience that has purchased products in the given category |purchases(exactProduct lookback=30)|Receive a bid recommendation for a target audience that has purchased the advertised asins |purchases(relatedProduct lookback=30)|Receive a bid recommendation for a target audience that has purchased related products to the advertised asins |audience(audienceSameAs=12345)|Receive a bid recommendation for the given target audience  | Content targeting clause | Description | |------------------|-------------| | [PREVIEW ONLY] contentCategorySameAs=amzn1.iab-content.325 | Receive a bid recommendation for the given content target |   #### Notes: - Refinements are currently not supported and if included will not impact the bid recommendation for the target.   #### Advertised ASIN Notes: - For asinSameAs targets the advertised asins will not impact the bid recommendation - For asinCategorySameAs targets the advertised asins are optional, but including them will provide a more refined bid recommendation - For similarProduct, exactProduct, and relatedProduct targets the advertised asins are required

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***BidRecommendationsApiGetTargetBidRecommendationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BidRecommendationsApiGetTargetBidRecommendationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SdTargetingBidRecommendationsRequestV34**](SdTargetingBidRecommendationsRequestV34.md)|  | 

### Return type

[**SdTargetingBidRecommendationsResponseV32**](SDTargetingBidRecommendationsResponseV32.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/vnd.sdtargetingrecommendations.v3.4+json, application/vnd.sdtargetingrecommendations.v3.3+json, application/vnd.sdtargetingrecommendations.v3.2+json, application/vnd.sdtargetingrecommendations.v3.1+json
 - **Accept**: application/vnd.sdtargetingrecommendations.v3.3+json, application/vnd.sdtargetingrecommendations.v3.2+json, application/vnd.sdtargetingrecommendations.v3.1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

