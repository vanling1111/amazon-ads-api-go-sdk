# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SBInsightsCampaignInsights**](InsightsApi.md#SBInsightsCampaignInsights) | **Post** /sb/campaigns/insights | Get insights for campaigns

# **SBInsightsCampaignInsights**
> SbInsightsCampaignInsightsResponseContent SBInsightsCampaignInsights(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Get insights for campaigns

Creates campaign level insights. Insights will be provided for passed in campaign parameters.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SbInsightsCampaignInsightsRequestContent**](SbInsightsCampaignInsightsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***InsightsApiSBInsightsCampaignInsightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiSBInsightsCampaignInsightsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nextToken** | **optional.**| Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;NextToken&#x60; field is empty, there are no further results. | 

### Return type

[**SbInsightsCampaignInsightsResponseContent**](SBInsightsCampaignInsightsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbinsights.v4+json
 - **Accept**: application/vnd.sbinsights.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

