# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SdCampaignsBudgetUsage**](BudgetUsageApi.md#SdCampaignsBudgetUsage) | **Post** /sd/campaigns/budget/usage | Budget usage API for SD campaigns

# **SdCampaignsBudgetUsage**
> BudgetUsageCampaignResponse SdCampaignsBudgetUsage(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Budget usage API for SD campaigns

**Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | [**Object**](.md)| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | [**Object**](.md)| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 
 **optional** | ***BudgetUsageApiSdCampaignsBudgetUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetUsageApiSdCampaignsBudgetUsageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of BudgetUsageCampaignRequest**](BudgetUsageCampaignRequest.md)|  | 

### Return type

[**BudgetUsageCampaignResponse**](BudgetUsageCampaignResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/vnd.sdcampaignbudgetusage.v1+json
 - **Accept**: application/vnd.sdcampaignbudgetusage.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

