# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStreamSubscription**](StreamSubscriptionApi.md#CreateStreamSubscription) | **Post** /streams/subscriptions | 
[**GetStreamSubscription**](StreamSubscriptionApi.md#GetStreamSubscription) | **Get** /streams/subscriptions/{subscriptionId} | 
[**ListStreamSubscriptions**](StreamSubscriptionApi.md#ListStreamSubscriptions) | **Get** /streams/subscriptions | 
[**UpdateStreamSubscription**](StreamSubscriptionApi.md#UpdateStreamSubscription) | **Put** /streams/subscriptions/{subscriptionId} | 

# **CreateStreamSubscription**
> CreateStreamSubscriptionResponseContent CreateStreamSubscription(ctx, body, amazonAdvertisingAPIClientId, optional)


Create a new subscription Note: trailing slash in request uri is not allowed  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"view_performance_dashboard\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateStreamSubscriptionRequestContent**](CreateStreamSubscriptionRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***StreamSubscriptionApiCreateStreamSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StreamSubscriptionApiCreateStreamSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| The identifier of a DSP advertiser level account | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**CreateStreamSubscriptionResponseContent**](CreateStreamSubscriptionResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.amazonmarketingstreamsubscriptions.v1+json
 - **Accept**: application/vnd.amazonmarketingstreamsubscriptions.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreamSubscription**
> GetStreamSubscriptionResponseContent GetStreamSubscription(ctx, subscriptionId, amazonAdvertisingAPIClientId, optional)


Fetch a specific subscription by Id Note: trailing slash in request uri is not allowed  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"view_performance_dashboard\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| Unique subscription identifier | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***StreamSubscriptionApiGetStreamSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StreamSubscriptionApiGetStreamSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.String**| The identifier of a DSP advertiser level account | 
 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**GetStreamSubscriptionResponseContent**](GetStreamSubscriptionResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.amazonmarketingstreamsubscriptions.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStreamSubscriptions**
> ListStreamSubscriptionsResponseContent ListStreamSubscriptions(ctx, amazonAdvertisingAPIClientId, optional)


List subscriptions Note: trailing slash in request uri is not allowed  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"view_performance_dashboard\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***StreamSubscriptionApiListStreamSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StreamSubscriptionApiListStreamSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **optional.Float64**| desired number of entries in the response, defaults to maximum value | 
 **startingToken** | **optional.String**| Token which can be used to get the next page of results, if more entries exist | 
 **amazonAdsAccountId** | **optional.String**| The identifier of a DSP advertiser level account | 
 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**ListStreamSubscriptionsResponseContent**](ListStreamSubscriptionsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.amazonmarketingstreamsubscriptions.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStreamSubscription**
> UpdateStreamSubscription(ctx, subscriptionId, amazonAdvertisingAPIClientId, optional)


Update an existing subscription Note: trailing slash in request uri is not allowed  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"view_performance_dashboard\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| Unique subscription identifier | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***StreamSubscriptionApiUpdateStreamSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StreamSubscriptionApiUpdateStreamSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateStreamSubscriptionRequestContent**](UpdateStreamSubscriptionRequestContent.md)|  | 
 **amazonAdsAccountId** | **optional.**| The identifier of a DSP advertiser level account | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.amazonmarketingstreamsubscriptions.v1+json
 - **Accept**: application/vnd.amazonmarketingstreamsubscriptions.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

