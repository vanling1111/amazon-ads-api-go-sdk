# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDspStreamSubscription**](DSPStreamSubscriptionApi.md#CreateDspStreamSubscription) | **Post** /dsp/streams/subscriptions | 
[**GetDspStreamSubscription**](DSPStreamSubscriptionApi.md#GetDspStreamSubscription) | **Get** /dsp/streams/subscriptions/{subscriptionId} | 
[**ListDspStreamSubscriptions**](DSPStreamSubscriptionApi.md#ListDspStreamSubscriptions) | **Get** /dsp/streams/subscriptions | 
[**UpdateDspStreamSubscription**](DSPStreamSubscriptionApi.md#UpdateDspStreamSubscription) | **Put** /dsp/streams/subscriptions/{subscriptionId} | 

# **CreateDspStreamSubscription**
> CreateDspStreamSubscriptionResponseContent CreateDspStreamSubscription(ctx, body, amazonAdsAccountID, amazonAdvertisingAPIClientId)


Create a new subscription Note: trailing slash in request uri is not allowed  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-Account-ID  **Parameter in**: header  **Requires one of these permissions**: [\"view_performance_dashboard\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDspStreamSubscriptionRequestContent**](CreateDspStreamSubscriptionRequestContent.md)|  | 
  **amazonAdsAccountID** | **string**| The identifier of a DSP advertiser level account | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 

### Return type

[**CreateDspStreamSubscriptionResponseContent**](CreateDspStreamSubscriptionResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.amazonmarketingstreamsubscriptions.v1+json
 - **Accept**: application/vnd.amazonmarketingstreamsubscriptions.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDspStreamSubscription**
> GetDspStreamSubscriptionResponseContent GetDspStreamSubscription(ctx, subscriptionId, amazonAdsAccountID, amazonAdvertisingAPIClientId)


Fetch a specific subscription by Id Note: trailing slash in request uri is not allowed  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-Account-ID  **Parameter in**: header  **Requires one of these permissions**: [\"view_performance_dashboard\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| Unique subscription identifier | 
  **amazonAdsAccountID** | **string**| The identifier of a DSP advertiser level account | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 

### Return type

[**GetDspStreamSubscriptionResponseContent**](GetDspStreamSubscriptionResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.amazonmarketingstreamsubscriptions.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDspStreamSubscriptions**
> ListDspStreamSubscriptionsResponseContent ListDspStreamSubscriptions(ctx, amazonAdsAccountID, amazonAdvertisingAPIClientId, optional)


List subscriptions Note: trailing slash in request uri is not allowed  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-Account-ID  **Parameter in**: header  **Requires one of these permissions**: [\"view_performance_dashboard\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsAccountID** | **string**| The identifier of a DSP advertiser level account | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***DSPStreamSubscriptionApiListDspStreamSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DSPStreamSubscriptionApiListDspStreamSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **maxResults** | **optional.Float64**| desired number of entries in the response, defaults to maximum value | 
 **startingToken** | **optional.String**| Token which can be used to get the next page of results, if more entries exist | 

### Return type

[**ListDspStreamSubscriptionsResponseContent**](ListDspStreamSubscriptionsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.amazonmarketingstreamsubscriptions.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDspStreamSubscription**
> UpdateDspStreamSubscription(ctx, subscriptionId, amazonAdsAccountID, amazonAdvertisingAPIClientId, optional)


Update an existing subscription Note: trailing slash in request uri is not allowed  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-Account-ID  **Parameter in**: header  **Requires one of these permissions**: [\"view_performance_dashboard\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| Unique subscription identifier | 
  **amazonAdsAccountID** | **string**| The identifier of a DSP advertiser level account | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***DSPStreamSubscriptionApiUpdateDspStreamSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DSPStreamSubscriptionApiUpdateDspStreamSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of UpdateDspStreamSubscriptionRequestContent**](UpdateDspStreamSubscriptionRequestContent.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.amazonmarketingstreamsubscriptions.v1+json
 - **Accept**: application/vnd.amazonmarketingstreamsubscriptions.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

