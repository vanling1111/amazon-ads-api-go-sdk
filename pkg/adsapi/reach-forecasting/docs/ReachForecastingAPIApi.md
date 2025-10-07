# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReachForecastsV1**](ReachForecastingAPIApi.md#CreateReachForecastsV1) | **Post** /mediaPlan/v1/reachForecasts | 
[**ListReachForecastTargetsV1**](ReachForecastingAPIApi.md#ListReachForecastTargetsV1) | **Post** /mediaPlan/v1/reachForecasts/targets/list | 
[**ListReachForecastsV1**](ReachForecastingAPIApi.md#ListReachForecastsV1) | **Post** /mediaPlan/v1/reachForecasts/list | 

# **CreateReachForecastsV1**
> CreateReachForecastsV1ResponseContent CreateReachForecastsV1(ctx, body, amazonAdvertisingAPIClientId, optional)


Creates a list of new Reach Forecasts in bulk action.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateReachForecastsV1RequestContent**](CreateReachForecastsV1RequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***ReachForecastingAPIApiCreateReachForecastsV1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReachForecastingAPIApiCreateReachForecastsV1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| Account identifier you use to access the DSP. This is your Amazon DSP Advertiser ID. | 
 **amazonAdsManagerAccountId** | **optional.**| The manager accounts associated to the ads accounts. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**CreateReachForecastsV1ResponseContent**](CreateReachForecastsV1ResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReachForecastTargetsV1**
> ListReachForecastTargetsV1ResponseContent ListReachForecastTargetsV1(ctx, body, amazonAdvertisingAPIClientId, optional)


Gets a list of targets of a Reach Forecast  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ListReachForecastTargetsV1RequestContent**](ListReachForecastTargetsV1RequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***ReachForecastingAPIApiListReachForecastTargetsV1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReachForecastingAPIApiListReachForecastTargetsV1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| Account identifier you use to access the DSP. This is your Amazon DSP Advertiser ID. | 
 **amazonAdsManagerAccountId** | **optional.**| The manager accounts associated to the ads accounts. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**ListReachForecastTargetsV1ResponseContent**](ListReachForecastTargetsV1ResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReachForecastsV1**
> ListReachForecastsV1ResponseContent ListReachForecastsV1(ctx, body, amazonAdvertisingAPIClientId, optional)


Gets a list of Reach Forecasts by IDs  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ListReachForecastsV1RequestContent**](ListReachForecastsV1RequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***ReachForecastingAPIApiListReachForecastsV1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReachForecastingAPIApiListReachForecastsV1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| Account identifier you use to access the DSP. This is your Amazon DSP Advertiser ID. | 
 **amazonAdsManagerAccountId** | **optional.**| The manager accounts associated to the ads accounts. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**ListReachForecastsV1ResponseContent**](ListReachForecastsV1ResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

