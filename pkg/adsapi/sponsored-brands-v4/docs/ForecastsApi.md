# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SBCampaignPerformanceForecasts**](ForecastsApi.md#SBCampaignPerformanceForecasts) | **Post** /sb/forecasts | Get performance forecasts for campaigns

# **SBCampaignPerformanceForecasts**
> SbCampaignPerformanceForecastsResponseContent SBCampaignPerformanceForecasts(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Get performance forecasts for campaigns

Returns forecasts for a list of new campaigns specified in SB forecast request. Currently only one new campaign is supported.   **If the campaign is not set to any goal during campaign creation then use PAGE_VISIT as goal value.**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SbCampaignPerformanceForecastsRequestContent**](SbCampaignPerformanceForecastsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**SbCampaignPerformanceForecastsResponseContent**](SBCampaignPerformanceForecastsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbforecasting.v4+json
 - **Accept**: application/vnd.sbforecasting.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

