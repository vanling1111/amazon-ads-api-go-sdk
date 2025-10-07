# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGsbTargetKpiRecommendation**](DefaultApi.md#GetGsbTargetKpiRecommendation) | **Post** /dsp/campaigns/targetKpi/recommendations | 

# **GetGsbTargetKpiRecommendation**
> GsbTargetKpiRecommendationResponse GetGsbTargetKpiRecommendation(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Creates a Target KPI recommendation for advertisers when they are in the process of creating a new campaign (ADSP).  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***DefaultApiGetGsbTargetKpiRecommendationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetGsbTargetKpiRecommendationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of GsbTargetKpiRecommendationRequest**](GsbTargetKpiRecommendationRequest.md)|  | 

### Return type

[**GsbTargetKpiRecommendationResponse**](GsbTargetKpiRecommendationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.gsbtargetkpirecommendation.v1+json
 - **Accept**: application/vnd.spinitialbudgetrecommendation.v3.4+json, application/vnd.gsbtargetkpirecommendation.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

