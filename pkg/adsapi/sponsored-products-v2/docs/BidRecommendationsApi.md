# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKeywordBidRecommendations**](BidRecommendationsApi.md#CreateKeywordBidRecommendations) | **Post** /v2/sp/keywords/bidRecommendations | Gets bid recommendations for keywords. [PLANNED DEPRECATION 3/27/2024]
[**GetAdGroupBidRecommendations**](BidRecommendationsApi.md#GetAdGroupBidRecommendations) | **Get** /v2/sp/adGroups/{adGroupId}/bidRecommendations | Gets a bid recommendation for an ad group. [PLANNED DEPRECATION 3/27/2024]
[**GetBidRecommendations**](BidRecommendationsApi.md#GetBidRecommendations) | **Post** /v2/sp/targets/bidRecommendations | Gets a list of bid recommendations for keyword, product, or auto targeting expressions.
[**GetKeywordBidRecommendations**](BidRecommendationsApi.md#GetKeywordBidRecommendations) | **Get** /v2/sp/keywords/{keywordId}/bidRecommendations | Gets a bid recommendation for a keyword. [PLANNED DEPRECATION 3/27/2024]

# **CreateKeywordBidRecommendations**
> BidRecommendationsResponse CreateKeywordBidRecommendations(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets bid recommendations for keywords. [PLANNED DEPRECATION 3/27/2024]

**Deprecation notice: This endpoint will be deprecated on March 27, 2024. Use [theme-based bid recommendations](/API/docs/en-us/sponsored-products/3-0/openapi/prod#/ThemeBasedBidRecommendation/GetThemeBasedBidRecommendationForAdGroup_v1) going forward.**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; developer account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***BidRecommendationsApiCreateKeywordBidRecommendationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BidRecommendationsApiCreateKeywordBidRecommendationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of KeywordBidRecommendationsData**](KeywordBidRecommendationsData.md)| An array of keyword bid recommendation objects. | 

### Return type

[**BidRecommendationsResponse**](BidRecommendationsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdGroupBidRecommendations**
> AdGroupBidRecommendationsResponse GetAdGroupBidRecommendations(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, adGroupId)
Gets a bid recommendation for an ad group. [PLANNED DEPRECATION 3/27/2024]

**Deprecation notice: This endpoint will be deprecated on March 27, 2024. Use [theme-based bid recommendations](/API/docs/en-us/sponsored-products/3-0/openapi/prod#/ThemeBasedBidRecommendation/GetThemeBasedBidRecommendationForAdGroup_v1) going forward.**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; developer account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **adGroupId** | **float64**| The identifier of an existing ad group. | 

### Return type

[**AdGroupBidRecommendationsResponse**](AdGroupBidRecommendationsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBidRecommendations**
> InlineResponse200 GetBidRecommendations(ctx, optional)
Gets a list of bid recommendations for keyword, product, or auto targeting expressions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BidRecommendationsApiGetBidRecommendationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BidRecommendationsApiGetBidRecommendationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TargetsBidRecommendationsBody**](TargetsBidRecommendationsBody.md)| An ad group identifier and list of associated targeting expressions for which to generate bid recommendations. Note that targeting expressions are required to be of the same type. That is, all targeting expressions in the list must be one of **keyword**, **product**, or **auto** target types. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeywordBidRecommendations**
> KeywordBidRecommendationsResponse GetKeywordBidRecommendations(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, keywordId)
Gets a bid recommendation for a keyword. [PLANNED DEPRECATION 3/27/2024]

**Deprecation notice: This endpoint will be deprecated on March 27, 2024. Use [theme-based bid recommendations](/API/docs/en-us/sponsored-products/3-0/openapi/prod#/ThemeBasedBidRecommendation/GetThemeBasedBidRecommendationForAdGroup_v1) going forward.**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; developer account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **keywordId** | **float64**| The identifier of an existing keyword. | 

### Return type

[**KeywordBidRecommendationsResponse**](KeywordBidRecommendationsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

