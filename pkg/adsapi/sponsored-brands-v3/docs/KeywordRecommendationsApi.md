# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKeywordRecommendations**](KeywordRecommendationsApi.md#GetKeywordRecommendations) | **Post** /sb/recommendations/keyword | Gets keyword recommendations

# **GetKeywordRecommendations**
> []SbKeywordSuggestion GetKeywordRecommendations(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets keyword recommendations

Gets an array of keyword recommendation objects for a set of ASINs included either on a landing page or a Stores page. Vendors may also specify a custom landing page.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***KeywordRecommendationsApiGetKeywordRecommendationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KeywordRecommendationsApiGetKeywordRecommendationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of RecommendationsKeywordBody**](RecommendationsKeywordBody.md)| **Must contain exactly only one of** 

 1.) An array of ASINs for which keyword recommendations are generated. 

 2.) The URL of a Stores page. Vendors may also specify the URL of a custom landing page. The products on the landing page are used to generate keyword recommendations. 

 Optional parameters include the max number of suggestions and locale for keyword translations. Supported locales include: Simplified Chinese (locale: “zh_CN”) for US, UK and CA. English (locale:  “en_GB”) for DE, FR, IT and ES. If locale is invalid or unsupported, no translations will be returned. | 

### Return type

[**[]SbKeywordSuggestion**](SBKeywordSuggestion.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/vnd.sbkeywordrecommendation.v3+json
 - **Accept**: application/vnd.sbkeywordrecommendation.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

