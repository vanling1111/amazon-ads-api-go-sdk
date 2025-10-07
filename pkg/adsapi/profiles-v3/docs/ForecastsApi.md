# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSDForecast**](ForecastsApi.md#CreateSDForecast) | **Post** /sd/forecasts | Return forecasts for an ad group that may or may not exist.

# **CreateSDForecast**
> SdForecastResponse CreateSDForecast(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Return forecasts for an ad group that may or may not exist.

Returns forecasts for a given ad group specified in SD forecast request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***ForecastsApiCreateSDForecastOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForecastsApiCreateSDForecastOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SdForecastRequest**](SdForecastRequest.md)|  | 

### Return type

[**SdForecastResponse**](SDForecastResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/vnd.sdforecasts.v3.1+json
 - **Accept**: application/vnd.sdforecasts.v3.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

