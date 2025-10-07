# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTarget**](TargetsApi.md#CreateTarget) | **Post** /adsApi/v1/create/targets | 
[**DeleteTarget**](TargetsApi.md#DeleteTarget) | **Post** /adsApi/v1/delete/targets | 
[**QueryTarget**](TargetsApi.md#QueryTarget) | **Post** /adsApi/v1/query/targets | 
[**UpdateTarget**](TargetsApi.md#UpdateTarget) | **Post** /adsApi/v1/update/targets | 

# **CreateTarget**
> TargetMultiStatusResponseWithPartialErrors CreateTarget(ctx, amazonAdsClientId, optional)


Creates Targets.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"], [\"campaign_edit\"], [\"advertiser_campaign_edit\",\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***TargetsApiCreateTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetsApiCreateTargetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateTargetRequest**](CreateTargetRequest.md)|  | 
 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**TargetMultiStatusResponseWithPartialErrors**](TargetMultiStatusResponseWithPartialErrors.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTarget**
> TargetMultiStatusResponseWithPartialErrors DeleteTarget(ctx, amazonAdsClientId, optional)


Archives or deletes Targets.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"], [\"campaign_edit\"], [\"advertiser_campaign_edit\",\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***TargetsApiDeleteTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetsApiDeleteTargetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DeleteTargetRequest**](DeleteTargetRequest.md)|  | 
 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**TargetMultiStatusResponseWithPartialErrors**](TargetMultiStatusResponseWithPartialErrors.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryTarget**
> TargetSuccessResponse QueryTarget(ctx, body, amazonAdsClientId, optional)


A search read, allowing use of more complex filters.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"], [\"campaign_edit\",\"campaign_view\"], [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"campaign_edit\",\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**QueryTargetRequest**](QueryTargetRequest.md)|  | 
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***TargetsApiQueryTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetsApiQueryTargetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**TargetSuccessResponse**](TargetSuccessResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTarget**
> TargetMultiStatusResponseWithPartialErrors UpdateTarget(ctx, amazonAdsClientId, optional)


Updates Targets. Behaves as PATCH unless otherwise stated.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***TargetsApiUpdateTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetsApiUpdateTargetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateTargetRequest**](UpdateTargetRequest.md)|  | 
 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**TargetMultiStatusResponseWithPartialErrors**](TargetMultiStatusResponseWithPartialErrors.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

