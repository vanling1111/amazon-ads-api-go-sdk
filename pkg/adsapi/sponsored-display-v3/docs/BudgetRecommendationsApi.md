# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSDBudgetRecommendations**](BudgetRecommendationsApi.md#GetSDBudgetRecommendations) | **Post** /sd/campaigns/budgetRecommendations | Returns recommended daily budget and estimated missed opportunities for campaigns

# **GetSDBudgetRecommendations**
> SdBudgetRecommendationsResponse GetSDBudgetRecommendations(ctx, amazonAdvertisingAPIClientId, optional)
Returns recommended daily budget and estimated missed opportunities for campaigns

Given a list of campaigns as input, this API provides the following metrics: <br> <b>1. Recommended daily budget - </b> Estimated budget needed to keep the campaign in budget for the full 24-hour period. Consider this budget to minimize your campaign's chances of running out of budget.  <br> <b>2. Percent time in budget </b> - The share of time the campaign was in budget during the past 7 days. <br> <b>3. Estimated missed impressions, clicks and sales </b> - These are the estimated additional impressions, clicks and sales the campaign might have generated had it adopted the recommended budget. These are estimates based on campaign's historical performance - and not a guarantee of actual impressions, clicks and sales. Consider using these metrics to further inform your budget allocation decisions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***BudgetRecommendationsApiGetSDBudgetRecommendationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetRecommendationsApiGetSDBudgetRecommendationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SdBudgetRecommendationsRequest**](SdBudgetRecommendationsRequest.md)|  | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **amazonAdsAccountId** | **optional.**| The identifier of an account. Account must be a global advertising account. | 

### Return type

[**SdBudgetRecommendationsResponse**](SDBudgetRecommendationsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/vnd.sdbudgetrecommendations.v3+json
 - **Accept**: application/vnd.sdbudgetrecommendations.v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

