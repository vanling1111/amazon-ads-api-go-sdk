# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrandVideoCreative**](AdCreativesApi.md#CreateBrandVideoCreative) | **Post** /sb/ads/creatives/brandVideo | Create new version of brand video ad creative
[**CreateExtendedProductCollectionCreative**](AdCreativesApi.md#CreateExtendedProductCollectionCreative) | **Post** /sb/ads/creatives/productCollectionExtended | Creates Sponsored Brands new version of product collection with collection of custom image ads
[**CreateProductCollectionCreative**](AdCreativesApi.md#CreateProductCollectionCreative) | **Post** /sb/ads/creatives/productCollection | Create new version of product collection ad creative
[**CreateStoreSpotlightCreative**](AdCreativesApi.md#CreateStoreSpotlightCreative) | **Post** /sb/ads/creatives/storeSpotlight | Create new version of store spotlight ad creative
[**CreateVideoCreative**](AdCreativesApi.md#CreateVideoCreative) | **Post** /sb/ads/creatives/video | Create new version of video ad creative
[**ListCreatives**](AdCreativesApi.md#ListCreatives) | **Post** /sb/ads/creatives/list | List ad creatives

# **CreateBrandVideoCreative**
> CreateBrandVideoCreativeResponseContent CreateBrandVideoCreative(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Create new version of brand video ad creative

This API creates a new version of an existing creative for given Sponsored Brands Ad by supplying brand video creative content  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateBrandVideoCreativeRequestContent**](CreateBrandVideoCreativeRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and use profileId from the response to pass as input. This is a required header for advertisers and integrators using the Advertising API. | 
 **optional** | ***AdCreativesApiCreateBrandVideoCreativeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdCreativesApiCreateBrandVideoCreativeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accept** | [**optional.Interface of AcceptHeader**](.md)| Clients request a specific version of a resource using the Accept request-header field set to the value field of the desired content-type. | 

### Return type

[**CreateBrandVideoCreativeResponseContent**](CreateBrandVideoCreativeResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbAdCreativeResource.v4+json
 - **Accept**: application/vnd.sbAdCreativeResource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateExtendedProductCollectionCreative**
> CreateExtendedProductCollectionCreativeResponseContent CreateExtendedProductCollectionCreative(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Creates Sponsored Brands new version of product collection with collection of custom image ads

This API creates a new version of creative for given Sponsored Brands ad by supplying extended product collection creative content  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateExtendedProductCollectionCreativeRequestContent**](CreateExtendedProductCollectionCreativeRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and use profileId from the response to pass as input. This is a required header for advertisers and integrators using the Advertising API. | 
 **optional** | ***AdCreativesApiCreateExtendedProductCollectionCreativeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdCreativesApiCreateExtendedProductCollectionCreativeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accept** | [**optional.Interface of AcceptHeader**](.md)| Clients request a specific version of a resource using the Accept request-header field set to the value field of the desired content-type. | 

### Return type

[**CreateExtendedProductCollectionCreativeResponseContent**](CreateExtendedProductCollectionCreativeResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbAdCreativeResource.v4+json
 - **Accept**: application/vnd.sbAdCreativeResource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProductCollectionCreative**
> CreateProductCollectionCreativeResponseContent CreateProductCollectionCreative(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Create new version of product collection ad creative

This API creates a new version of creative for given Sponsored Brands ad by supplying product collection creative content  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateProductCollectionCreativeRequestContent**](CreateProductCollectionCreativeRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and use profileId from the response to pass as input. This is a required header for advertisers and integrators using the Advertising API. | 
 **optional** | ***AdCreativesApiCreateProductCollectionCreativeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdCreativesApiCreateProductCollectionCreativeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accept** | [**optional.Interface of AcceptHeader**](.md)| Clients request a specific version of a resource using the Accept request-header field set to the value field of the desired content-type. | 

### Return type

[**CreateProductCollectionCreativeResponseContent**](CreateProductCollectionCreativeResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbAdCreativeResource.v4+json
 - **Accept**: application/vnd.sbAdCreativeResource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStoreSpotlightCreative**
> CreateStoreSpotlightCreativeResponseContent CreateStoreSpotlightCreative(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Create new version of store spotlight ad creative

This API creates a new version of creative for given Sponsored Brands ad by supplying store spotlight creative content  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateStoreSpotlightCreativeRequestContent**](CreateStoreSpotlightCreativeRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and use profileId from the response to pass as input. This is a required header for advertisers and integrators using the Advertising API. | 
 **optional** | ***AdCreativesApiCreateStoreSpotlightCreativeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdCreativesApiCreateStoreSpotlightCreativeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accept** | [**optional.Interface of AcceptHeader**](.md)| Clients request a specific version of a resource using the Accept request-header field set to the value field of the desired content-type. | 

### Return type

[**CreateStoreSpotlightCreativeResponseContent**](CreateStoreSpotlightCreativeResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbAdCreativeResource.v4+json
 - **Accept**: application/vnd.sbAdCreativeResource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVideoCreative**
> CreateVideoCreativeResponseContent CreateVideoCreative(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Create new version of video ad creative

This API creates a new version of an existing creative for given Sponsored Brands ad by supplying video creative content  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVideoCreativeRequestContent**](CreateVideoCreativeRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and use profileId from the response to pass as input. This is a required header for advertisers and integrators using the Advertising API. | 
 **optional** | ***AdCreativesApiCreateVideoCreativeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdCreativesApiCreateVideoCreativeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accept** | [**optional.Interface of AcceptHeader**](.md)| Clients request a specific version of a resource using the Accept request-header field set to the value field of the desired content-type. | 

### Return type

[**CreateVideoCreativeResponseContent**](CreateVideoCreativeResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbAdCreativeResource.v4+json
 - **Accept**: application/vnd.sbAdCreativeResource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCreatives**
> ListCreativesResponseContent ListCreatives(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
List ad creatives

This API gets an array of all Sponsored Brands creatives that qualify the given resource identifiers and filters  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ListCreativesRequestContent**](ListCreativesRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and use profileId from the response to pass as input. This is a required header for advertisers and integrators using the Advertising API. | 
 **optional** | ***AdCreativesApiListCreativesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdCreativesApiListCreativesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accept** | [**optional.Interface of AcceptHeader**](.md)| Clients request a specific version of a resource using the Accept request-header field set to the value field of the desired content-type. | 

### Return type

[**ListCreativesResponseContent**](ListCreativesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbAdCreativeResource.v4+json
 - **Accept**: application/vnd.sbAdCreativeResource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

