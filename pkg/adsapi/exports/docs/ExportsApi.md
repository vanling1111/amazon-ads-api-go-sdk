# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdExport**](ExportsApi.md#AdExport) | **Post** /ads/export | Creates a file-based export of Ads.
[**AdGroupExport**](ExportsApi.md#AdGroupExport) | **Post** /adGroups/export | Creates a file-based export of Ad Groups.
[**CampaignExport**](ExportsApi.md#CampaignExport) | **Post** /campaigns/export | Creates a file-based export of Campaigns.
[**GetExport**](ExportsApi.md#GetExport) | **Get** /exports/{exportId} | Gets the status of a requested export and a link to download the export.
[**TargetExport**](ExportsApi.md#TargetExport) | **Post** /targets/export | Creates a file-based export of Targets.

# **AdExport**
> UniversalApiExportResponse AdExport(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Creates a file-based export of Ads.

Creates a file-based export of Ads in the account satisfying the filtering criteria.  To understand the call flow for asynchronous exports, see [Getting started with sponsored ads exports](/API/docs/en-us/guides/exports/get-started).  **Requires one of these permissions**: [\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseUniversalApiExportRequest**](BaseUniversalApiExportRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**UniversalApiExportResponse**](UniversalApiExportResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/vnd.adsexport.v1+json
 - **Accept**: application/vnd.adsexport.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdGroupExport**
> UniversalApiExportResponse AdGroupExport(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Creates a file-based export of Ad Groups.

Creates a file-based export of Ad Groups in the account satisfying the filtering criteria.  To understand the call flow for asynchronous exports, see [Getting started with sponsored ads exports](/API/docs/en-us/guides/exports/get-started).  **Requires one of these permissions**: [\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseUniversalApiExportRequest**](BaseUniversalApiExportRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**UniversalApiExportResponse**](UniversalApiExportResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/vnd.adgroupsexport.v1+json
 - **Accept**: application/vnd.adgroupsexport.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignExport**
> UniversalApiExportResponse CampaignExport(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Creates a file-based export of Campaigns.

Creates a file-based export of Campaigns in the account satisfying the filtering criteria.  To understand the call flow for asynchronous exports, see [Getting started with sponsored ads exports](/API/docs/en-us/guides/exports/get-started).  **Requires one of these permissions**: [\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseUniversalApiExportRequest**](BaseUniversalApiExportRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**UniversalApiExportResponse**](UniversalApiExportResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/vnd.campaignsexport.v1+json
 - **Accept**: application/vnd.campaignsexport.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExport**
> UniversalApiExportResponse GetExport(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, exportId)
Gets the status of a requested export and a link to download the export.

This API will return a status of the specified export.  To understand the call flow for asynchronous exports, see [Getting started with sponsored ads exports](/API/docs/en-us/guides/exports/get-started).  **Requires one of these permissions**: [\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **exportId** | **string**| The export identifier. | 

### Return type

[**UniversalApiExportResponse**](UniversalApiExportResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.adgroupsexport.v1+json, application/vnd.adsexport.v1+json, application/vnd.campaignsexport.v1+json, application/vnd.targetsexport.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetExport**
> UniversalApiExportResponse TargetExport(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Creates a file-based export of Targets.

Creates a file-based export of Targets in the account satisfying the filtering criteria.  To understand the call flow for asynchronous exports, see [Getting started with sponsored ads exports](/API/docs/en-us/guides/exports/get-started).  **Requires one of these permissions**: [\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TargetsUniversalApiExportRequest**](TargetsUniversalApiExportRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**UniversalApiExportResponse**](UniversalApiExportResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/vnd.targetsexport.v1+json
 - **Accept**: application/vnd.targetsexport.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

