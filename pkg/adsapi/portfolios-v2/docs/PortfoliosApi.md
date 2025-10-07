# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePortfolios**](PortfoliosApi.md#CreatePortfolios) | **Post** /portfolios | 
[**ListPortfolios**](PortfoliosApi.md#ListPortfolios) | **Post** /portfolios/list | 
[**UpdatePortfolios**](PortfoliosApi.md#UpdatePortfolios) | **Put** /portfolios | 

# **CreatePortfolios**
> CreatePortfoliosResponseContent CreatePortfolios(ctx, body, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreatePortfoliosRequestContent**](CreatePortfoliosRequestContent.md)|  | 
 **optional** | ***PortfoliosApiCreatePortfoliosOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfoliosApiCreatePortfoliosOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amazonAdvertisingAPIClientId** | **optional.**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;canonicalObjectId - (default) - returns objectIds encoded using canonical representation return&#x3D;obfuscatedObjectId - returns obfuscated objectIds return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids | 

### Return type

[**CreatePortfoliosResponseContent**](CreatePortfoliosResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.spPortfolio.v3+json
 - **Accept**: application/vnd.spPortfolio.v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPortfolios**
> ListPortfoliosResponseContent ListPortfolios(ctx, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfoliosApiListPortfoliosOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfoliosApiListPortfoliosOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ListPortfoliosRequestContent**](ListPortfoliosRequestContent.md)|  | 
 **amazonAdvertisingAPIClientId** | **optional.**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;canonicalObjectId - (default) - returns objectIds encoded using canonical representation return&#x3D;obfuscatedObjectId - returns obfuscated objectIds return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids | 

### Return type

[**ListPortfoliosResponseContent**](ListPortfoliosResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.spPortfolio.v3+json
 - **Accept**: application/vnd.spPortfolio.v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePortfolios**
> UpdatePortfoliosResponseContent UpdatePortfolios(ctx, body, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePortfoliosRequestContent**](UpdatePortfoliosRequestContent.md)|  | 
 **optional** | ***PortfoliosApiUpdatePortfoliosOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfoliosApiUpdatePortfoliosOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amazonAdvertisingAPIClientId** | **optional.**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;canonicalObjectId - (default) - returns objectIds encoded using canonical representation return&#x3D;obfuscatedObjectId - returns obfuscated objectIds return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids | 

### Return type

[**UpdatePortfoliosResponseContent**](UpdatePortfoliosResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.spPortfolio.v3+json
 - **Accept**: application/vnd.spPortfolio.v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

