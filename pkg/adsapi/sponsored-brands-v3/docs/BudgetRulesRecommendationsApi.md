# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SBGetBudgetRulesRecommendation**](BudgetRulesRecommendationsApi.md#SBGetBudgetRulesRecommendation) | **Post** /sb/campaigns/budgetRules/recommendations | Gets a list of special events with suggested date range and suggested budget increase for a campaign specified by identifier.

# **SBGetBudgetRulesRecommendation**
> SbBudgetRulesRecommendationEventResponse SBGetBudgetRulesRecommendation(ctx, amazonAdvertisingAPIClientId, optional)
Gets a list of special events with suggested date range and suggested budget increase for a campaign specified by identifier.

A rule enables an automatic budget increase for a specified date range or for a special event. The response also includes a suggested budget increase for each special event.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
 **optional** | ***BudgetRulesRecommendationsApiSBGetBudgetRulesRecommendationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetRulesRecommendationsApiSBGetBudgetRulesRecommendationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BudgetRulesRecommendationsBody**](BudgetRulesRecommendationsBody.md)|  | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **amazonAdsAccountId** | **optional.**| The identifier of an account. The account must be a global advertising account.. | 

### Return type

[**SbBudgetRulesRecommendationEventResponse**](SBBudgetRulesRecommendationEventResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/vnd.sbbudgetrulesrecommendation.v3+json
 - **Accept**: application/vnd.sbbudgetrulesrecommendation.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

