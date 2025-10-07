# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHeadlineRecommendationsForSD**](HeadlineRecommendationsApi.md#GetHeadlineRecommendationsForSD) | **Post** /sd/recommendations/creative/headline | 

# **GetHeadlineRecommendationsForSD**
> SdHeadlineRecommendationResponse GetHeadlineRecommendationsForSD(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)


You can use this Sponsored Display API to retrieve creative headline recommendations from an array of ASINs.  **Requires one of these permissions**: [\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdHeadlineRecommendationRequest**](SdHeadlineRecommendationRequest.md)| Request body for SD headline recommendations API. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**SdHeadlineRecommendationResponse**](SDHeadlineRecommendationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/vnd.sdheadlinerecommendationrequest.v4.0+json
 - **Accept**: application/vnd.sdheadlinerecommendationresponse.v4.0+json, application/vnd.sdheadlinerecommendationschemavalidationexception.v4.0+json, application/vnd.sdheadlinerecommendationaccessdeniedexception.v4.0+json, application/vnd.sdheadlinerecommendationidentifiernotfoundexception.v4.0+json, application/vnd.sdheadlinerecommendationthrottlingexception.v4.0+json, application/vnd.sdheadlinerecommendationinternalserverexception.v4.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

