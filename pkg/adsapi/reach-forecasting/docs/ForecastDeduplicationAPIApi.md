# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeduplicatedReachForecastsV1**](ForecastDeduplicationAPIApi.md#CreateDeduplicatedReachForecastsV1) | **Post** /mediaPlan/v1/deduplicatedReachForecasts | 

# **CreateDeduplicatedReachForecastsV1**
> CreateDeduplicatedReachForecastsV1ResponseContent CreateDeduplicatedReachForecastsV1(ctx, body, amazonAdvertisingAPIClientId, optional)


Creates a list of De-duplicated Reach Forecasts.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDeduplicatedReachForecastsV1RequestContent**](CreateDeduplicatedReachForecastsV1RequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***ForecastDeduplicationAPIApiCreateDeduplicatedReachForecastsV1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastDeduplicationAPIApiCreateDeduplicatedReachForecastsV1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| Account identifier you use to access the DSP. This is your Amazon DSP Advertiser ID. | 
 **amazonAdsManagerAccountId** | **optional.**| The manager accounts associated to the ads accounts. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**CreateDeduplicatedReachForecastsV1ResponseContent**](CreateDeduplicatedReachForecastsV1ResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

