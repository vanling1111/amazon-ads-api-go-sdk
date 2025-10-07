# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTargetRecommendations**](TargetingRecommendationsApi.md#GetTargetRecommendations) | **Post** /sd/targets/recommendations | Returns a set of recommended products and categories to target

# **GetTargetRecommendations**
> SdTargetingRecommendationsResponseV35 GetTargetRecommendations(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Returns a set of recommended products and categories to target

This API provides product, category and standard audience recommendations to target based on the list of input ASINs. Allow 1 week for our systems to process data for any new ASINs listed on Amazon before using this service. Note -  recommendations are only available for productAds with SKU or ASIN.  For API v3.0, the API returns up to 100 recommendations for contextual targeting.  For API v3.1, the API returns up to 100 recommendations for both product and category targeting.  For API v3.2, the API introduces contextual targeting themes in the request and returns product recommendations based on different targeting themes.  For API v3.3, the API introduces standard audience recommendations and translated category recommendations based on locale.  For API v3.4, the API includes the theme expression used in contextual targeting recommendations in the response.  [PREVIEW ONLY] For API v3.5, the API supports recommendations for landing pages for audience targeting with tactic T00030. The API also supports entertainment targeting recommendations. Both features are currently limited to US marketplace.  The currently available tactic identifiers are:  |Tactic Name|Type|Description| |-----------|----|-----------| |T00020&nbsp;|Contextual Targeting|Products: Choose individual products to show your ads in placements related to those products.| |T00030&nbsp;|Audiences or Contextual Targeting| Select individual products, categories, refined categories, or audiences to show your ads.|

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***TargetingRecommendationsApiGetTargetRecommendationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingRecommendationsApiGetTargetRecommendationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SdTargetingRecommendationsRequestV35**](SdTargetingRecommendationsRequestV35.md)|  | 
 **locale** | [**optional.Interface of SdTargetingRecommendationsLocale**](.md)| The requested locale from query parameter to return translated category recommendations. | 

### Return type

[**SdTargetingRecommendationsResponseV35**](SDTargetingRecommendationsResponseV35.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/vnd.sdtargetingrecommendations.v3.5+json, application/vnd.sdtargetingrecommendations.v3.4+json, application/vnd.sdtargetingrecommendations.v3.3+json, application/vnd.sdtargetingrecommendations.v3.2+json, application/vnd.sdtargetingrecommendations.v3.1+json, application/vnd.sdtargetingrecommendations.v3.0+json
 - **Accept**: application/vnd.sdtargetingrecommendations.v3.5+json, application/vnd.sdtargetingrecommendations.v3.4+json, application/vnd.sdtargetingrecommendations.v3.3+json, application/vnd.sdtargetingrecommendations.v3.2+json, application/vnd.sdtargetingrecommendations.v3.1+json, application/vnd.sdtargetingrecommendations.v3.0+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

