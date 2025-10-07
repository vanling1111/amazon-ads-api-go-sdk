# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBidsRecommendations**](BidRecommendationsApi.md#GetBidsRecommendations) | **Post** /sb/recommendations/bids | 

# **GetBidsRecommendations**
> InlineResponse2007 GetBidsRecommendations(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Get a list of bid recommendation objects for a specified list of keywords or products.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***BidRecommendationsApiGetBidsRecommendationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BidRecommendationsApiGetBidsRecommendationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of RecommendationsBidsBody**](RecommendationsBidsBody.md)| A list of keywords or targeting expressions for which to generate bid recommendations. Note that if a value is specified for the &#x60;campaignId&#x60; field, the past performance data for the campaign may be use to create bid recommendations. | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sbbidsrecommendation.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

