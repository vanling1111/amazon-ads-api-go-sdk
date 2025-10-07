# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveProductAd**](ProductAdsApi.md#ArchiveProductAd) | **Delete** /sd/productAds/{adId} | Sets the status of a sproduct ad to archived.
[**CreateProductAds**](ProductAdsApi.md#CreateProductAds) | **Post** /sd/productAds | Creates one or more product ads.
[**GetProductAd**](ProductAdsApi.md#GetProductAd) | **Get** /sd/productAds/{adId} | Gets a requested product ad.
[**GetProductAdResponseEx**](ProductAdsApi.md#GetProductAdResponseEx) | **Get** /sd/productAds/extended/{adId} | Gets extended information for a product ad.
[**ListProductAds**](ProductAdsApi.md#ListProductAds) | **Get** /sd/productAds | Gets a list of product ads.
[**ListProductAdsEx**](ProductAdsApi.md#ListProductAdsEx) | **Get** /sd/productAds/extended | Gets a list of product ads with extended fields.
[**UpdateProductAds**](ProductAdsApi.md#UpdateProductAds) | **Put** /sd/productAds | Updates one or more product ads.

# **ArchiveProductAd**
> ProductAdResponse ArchiveProductAd(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, adId)
Sets the status of a sproduct ad to archived.

This operation is equivalent to an update operation that sets the status field to 'archived'. Note that setting the status field to 'archived' is permanent and can't be undone. See [Developer Notes](https://advertising.amazon.com/API/docs/en-us/info/developer-notes#archiving) for more information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **adId** | **int64**| The identifier of the produce ad. | 

### Return type

[**ProductAdResponse**](ProductAdResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProductAds**
> []ProductAdResponse CreateProductAds(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Creates one or more product ads.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***ProductAdsApiCreateProductAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductAdsApiCreateProductAdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []CreateProductAd**](CreateProductAd.md)| An array of ProductAd objects. For each object, specify required fields and their values. Required fields are &#x60;adGroupId&#x60; and &#x60;state&#x60;. Within any campaign, you must pass consistent fields of either &#x60;asin&#x60; (for vendors), &#x60;sku&#x60; (for sellers), or &#x60;landingPageURL&#x60; (when linking to other pages), these cannot be mixed for any given campaign. Maximum length of the array is 100 objects. | 

### Return type

[**[]ProductAdResponse**](ProductAdResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductAd**
> ProductAd GetProductAd(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, adId)
Gets a requested product ad.

Note that the ProductAd object is designed for performance, and includes a small set of commonly used fields to reduce size. If the extended set of fields is required, use a product ad operations that returns the ProductAdResponseEx object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **adId** | **int64**| The identifier of the requested product ad. | 

### Return type

[**ProductAd**](ProductAd.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductAdResponseEx**
> ProductAdResponseEx GetProductAdResponseEx(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, adId)
Gets extended information for a product ad.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **adId** | **int64**| The identifier of the requested product ad. | 

### Return type

[**ProductAdResponseEx**](ProductAdResponseEx.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProductAds**
> []ProductAd ListProductAds(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of product ads.

Gets an array of ProductAd objects for a requested set of Sponsored Display product ads. Note that the ProductAd object is designed for performance, and includes a small set of commonly used fields to reduce size. If the extended set of fields is required, use a product ad operation that returns the ProductAdResponseEx object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***ProductAdsApiListProductAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductAdsApiListProductAdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. Sets a cursor into the requested set of product ads. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0. | 
 **count** | **optional.Int32**| Optional. Sets the number of ProductAd objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten product ad set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten product ads, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size. | 
 **stateFilter** | **optional.String**| Optional. The returned array is filtered to include only products ads associated with campaigns that have state set to one of the values in the comma-delimited list. | [default to enabled, paused, archived]
 **adIdFilter** | **optional.String**| Optional. The returned array includes only product ads with identifiers matching those in the comma-delimited string. | 
 **adGroupIdFilter** | **optional.String**| Optional. The returned array is filtered to include only products ads associated with ad groups identifiers in the comma-delimited list. | 
 **campaignIdFilter** | **optional.String**| Optional. The returned array is filtered to include only product ads associated with the campaign identifiers in the comma-delimited list. | 

### Return type

[**[]ProductAd**](ProductAd.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProductAdsEx**
> []ProductAdResponseEx ListProductAdsEx(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of product ads with extended fields.

Gets an array of ProductAdResponseEx objects for a set of requested ad groups. The ProductAdResponseEx object includes the extended set of available fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***ProductAdsApiListProductAdsExOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductAdsApiListProductAdsExOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. Sets a cursor into the requested set of product ads. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0. | 
 **count** | **optional.Int32**| Optional. Sets the number of ProduceAdEx objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten product ads set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten campaigns, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size. | 
 **stateFilter** | **optional.String**| Optional. The returned array is filtered to include only campaigns with state set to one of the values in the specified comma-delimited list. | [default to enabled, paused, archived]
 **adIdFilter** | **optional.String**| Optional. The returned array includes only product ads with identifiers matching those in the comma-delimited string. | 
 **adGroupIdFilter** | **optional.String**| Optional. The returned array is filtered to include only products ads associated with ad groups identifiers in the comma-delimited list. | 
 **campaignIdFilter** | **optional.String**| Optional. The returned array is filtered to include only product ads associated with the campaign identifiers in the comma-delimited list. | 

### Return type

[**[]ProductAdResponseEx**](ProductAdResponseEx.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProductAds**
> []ProductAdResponse UpdateProductAds(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Updates one or more product ads.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***ProductAdsApiUpdateProductAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductAdsApiUpdateProductAdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []UpdateProductAd**](UpdateProductAd.md)| An array of ProductAd objects. For each object, specify a product ad identifier and the only mutable field, &#x60;state&#x60;. Maximum length of the array is 100 objects. | 

### Return type

[**[]ProductAdResponse**](ProductAdResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

