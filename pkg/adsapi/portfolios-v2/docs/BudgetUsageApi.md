# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortfolioBudgetUsage**](BudgetUsageApi.md#PortfolioBudgetUsage) | **Post** /portfolios/budget/usage | Budget usage API for portfolios

# **PortfolioBudgetUsage**
> BudgetUsagePortfolioResponse PortfolioBudgetUsage(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Budget usage API for portfolios

  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | [**Object**](.md)| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | [**Object**](.md)| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 
 **optional** | ***BudgetUsageApiPortfolioBudgetUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetUsageApiPortfolioBudgetUsageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of BudgetUsagePortfolioRequest**](BudgetUsagePortfolioRequest.md)|  | 

### Return type

[**BudgetUsagePortfolioResponse**](BudgetUsagePortfolioResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.portfoliobudgetusage.v1+json
 - **Accept**: application/vnd.portfoliobudgetusage.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

