# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBudgetRecommendations**](RecommendationsApi.md#GetBudgetRecommendations) | **Post** /sb/campaigns/budgetRecommendations | Get budget recommendations
[**GetHeadlineRecommendations**](RecommendationsApi.md#GetHeadlineRecommendations) | **Post** /sb/recommendations/creative/headline | Get recommendations for creative headline
[**SBOptimizationRecommendation**](RecommendationsApi.md#SBOptimizationRecommendation) | **Post** /sb/recommendations/optimization | Get recommendation for optimization rule
[**SBTargetingGetNegativeBrands**](RecommendationsApi.md#SBTargetingGetNegativeBrands) | **Get** /sb/negativeTargets/brands/recommendations | Get brand recommendations for negative targeting

# **GetBudgetRecommendations**
> GetBudgetRecommendationsResponseContent GetBudgetRecommendations(ctx, body, amazonAdvertisingAPIClientId, optional)
Get budget recommendations

Provides daily budget recommendations for a list of requested Sponsored Brands campaigns, with context on estimated historical missed opportunities.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetBudgetRecommendationsRequestContent**](GetBudgetRecommendationsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
 **optional** | ***RecommendationsApiGetBudgetRecommendationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecommendationsApiGetBudgetRecommendationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **amazonAdsAccountId** | **optional.**| The identifier of an account. Account must be a global advertising account. | 

### Return type

[**GetBudgetRecommendationsResponseContent**](GetBudgetRecommendationsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbbudgetrecommendation.v4+json
 - **Accept**: application/vnd.sbbudgetrecommendation.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHeadlineRecommendations**
> HeadlineSuggestionResponse GetHeadlineRecommendations(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Get recommendations for creative headline

API to receive creative headline suggestions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HeadlineSuggestionRequest**](HeadlineSuggestionRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**HeadlineSuggestionResponse**](HeadlineSuggestionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SBOptimizationRecommendation**
> SbOptimizationRecommendationResponseContent SBOptimizationRecommendation(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Get recommendation for optimization rule

Returns recommended bid value for optimization rule enable campaigns. Recommendations are generated based landing page, page type and ASINs provided in request. Only available for Sellers and Vendors.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]\"

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SbOptimizationRecommendationRequestContent**](SbOptimizationRecommendationRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**SbOptimizationRecommendationResponseContent**](SBOptimizationRecommendationResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sboptimizationrecommendationresource.v4+json
 - **Accept**: application/json, application/vnd.sboptimizationrecommendationresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SBTargetingGetNegativeBrands**
> SbTargetingGetNegativeBrandsResponseContent SBTargetingGetNegativeBrands(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Get brand recommendations for negative targeting

Returns brands recommended for negative targeting. Only available for Sellers and Vendors. These recommendations include your own brands because targeting your own brands usually results in lower performance than targeting competitors' brands.   Only available in the following marketplaces: US, CA, MX, UK, DE, FR, ES, IT, IN, JP  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***RecommendationsApiSBTargetingGetNegativeBrandsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecommendationsApiSBTargetingGetNegativeBrandsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nextToken** | **optional.String**| Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;NextToken&#x60; field is empty, there are no further results. | 

### Return type

[**SbTargetingGetNegativeBrandsResponseContent**](SBTargetingGetNegativeBrandsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbtargeting.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

