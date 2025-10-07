# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAsinEngagementForStore**](StoresAnalyticsApi.md#GetAsinEngagementForStore) | **Post** /stores/{brandEntityId}/asinMetrics | 
[**GetInsightsForStoreAPI**](StoresAnalyticsApi.md#GetInsightsForStoreAPI) | **Post** /stores/{brandEntityId}/insights | 

# **GetAsinEngagementForStore**
> GetAsinEngagementForStoreResponse GetAsinEngagementForStore(ctx, body, brandEntityId)


Store asin metrics provides information about your store asin performance, including rendered impressions, viewed impressions, clicks and sales. You can access Stores insights through this API.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetAsinEngagementForStoreRequest**](GetAsinEngagementForStoreRequest.md)|  | 
  **brandEntityId** | **string**| The identifier of the requested store. It can be fetched through calling existing API \&quot;/v2/stores\&quot; listed here https://advertising.amazon.com/API/docs/en-us/reference/2/stores | 

### Return type

[**GetAsinEngagementForStoreResponse**](GetAsinEngagementForStoreResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.GetAsinEngagementForStoreRequest.v1+json
 - **Accept**: application/vnd.GetAsinEngagementForStoreResponse.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInsightsForStoreAPI**
> GetInsightsForStoreResponse GetInsightsForStoreAPI(ctx, body, brandEntityId)


Stores insights provides information about your store's performance, including traffic and sales. You can access Stores insights through this API.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetInsightsForStoreRequest**](GetInsightsForStoreRequest.md)|  | 
  **brandEntityId** | **string**| The identifier of the requested store. It can be fetched through calling existing API \&quot;/v2/stores\&quot; listed here https://advertising.amazon.com/API/docs/en-us/reference/2/stores | 

### Return type

[**GetInsightsForStoreResponse**](GetInsightsForStoreResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.GetInsightsForStoreRequest.v1+json
 - **Accept**: application/vnd.GetInsightsForStoreResponse.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

