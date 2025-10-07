# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RASv1CreateProductAds**](ProductAdsApi.md#RASv1CreateProductAds) | **Post** /ras/v1/productAds | 
[**RASv1DeleteProductAds**](ProductAdsApi.md#RASv1DeleteProductAds) | **Post** /ras/v1/productAds/delete | 
[**RASv1ListProductAds**](ProductAdsApi.md#RASv1ListProductAds) | **Post** /ras/v1/productAds/list | 
[**RASv1UpdateProductAds**](ProductAdsApi.md#RASv1UpdateProductAds) | **Put** /ras/v1/productAds | 

# **RASv1CreateProductAds**
> Rasv1CreateProductAdsResponseContent RASv1CreateProductAds(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rasv1CreateProductAdsRequestContent**](Rasv1CreateProductAdsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***ProductAdsApiRASv1CreateProductAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductAdsApiRASv1CreateProductAdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. | 

### Return type

[**Rasv1CreateProductAdsResponseContent**](RASv1CreateProductAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RASv1DeleteProductAds**
> Rasv1DeleteProductAdsResponseContent RASv1DeleteProductAds(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rasv1DeleteProductAdsRequestContent**](Rasv1DeleteProductAdsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**Rasv1DeleteProductAdsResponseContent**](RASv1DeleteProductAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RASv1ListProductAds**
> Rasv1ListProductAdsResponseContent RASv1ListProductAds(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***ProductAdsApiRASv1ListProductAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductAdsApiRASv1ListProductAdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Rasv1ListProductAdsRequestContent**](Rasv1ListProductAdsRequestContent.md)|  | 

### Return type

[**Rasv1ListProductAdsResponseContent**](RASv1ListProductAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RASv1UpdateProductAds**
> Rasv1UpdateProductAdsResponseContent RASv1UpdateProductAds(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rasv1UpdateProductAdsRequestContent**](Rasv1UpdateProductAdsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***ProductAdsApiRASv1UpdateProductAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductAdsApiRASv1UpdateProductAdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. | 

### Return type

[**Rasv1UpdateProductAdsResponseContent**](RASv1UpdateProductAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

