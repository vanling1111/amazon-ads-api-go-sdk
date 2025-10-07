# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SponsoredTvForecasts**](ForecastsApi.md#SponsoredTvForecasts) | **Post** /st/forecasts | 

# **SponsoredTvForecasts**
> SponsoredTvForecastsResponseContent SponsoredTvForecasts(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)


Returns forecasts for a given ad group specified in Sponsored TV forecast request.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SponsoredTvForecastsRequestContent**](SponsoredTvForecastsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**SponsoredTvForecastsResponseContent**](SponsoredTvForecastsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stForecast.v1+json
 - **Accept**: application/vnd.stForecast.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

