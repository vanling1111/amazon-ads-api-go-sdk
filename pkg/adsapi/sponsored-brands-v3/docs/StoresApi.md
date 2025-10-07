# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAsset**](StoresApi.md#CreateAsset) | **Post** /stores/assets | Creates a new image asset.
[**ListAssets**](StoresApi.md#ListAssets) | **Get** /stores/assets | Gets a list of assets associated with a specified brand entity identifier.

# **CreateAsset**
> InlineResponse2006 CreateAsset(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, contentDisposition, contentType, optional)
Creates a new image asset.

Image assets are stored in the Store Assets Library. Note that there may be a delay before the image is displayed in the console.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **contentDisposition** | **string**| The name of the image file. | 
  **contentType** | **string**| The image format type. The following table lists the valid image types: |Image Type|Description| |----------|-----------| |PNG|[Portable network graphics](https://en.wikipedia.org/wiki/Portable_Network_Graphics)| |JPEG|[JPEG](https://en.wikipedia.org/wiki/JPEG)| |GIF|[Graphics interchange format](https://en.wikipedia.org/wiki/GIF)| | 
 **optional** | ***StoresApiCreateAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StoresApiCreateAssetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **assetInfo** | **optional.**|  | 
 **asset** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAssets**
> []InlineResponse2005 ListAssets(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of assets associated with a specified brand entity identifier.

For sellers or vendors, gets an array of assets associated with the specified brand entity identifier. Vendors are not required to specify a brand entity identifier, and in this case all assets associated with the vendor are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***StoresApiListAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StoresApiListAssetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **brandEntityId** | **optional.String**| For sellers, this field is required. It is the Brand entity identifier of the Brand for which assets are returned. This identifier is retrieved using the [getBrands operation](https://advertising.amazon.com/API/docs/v3/reference/SponsoredBrands/brands). For vendors, this field is optional. If a vendor does not specify this field, all assets associated with the vendor are returned. For more information about the [difference between a seller and a vendor](https://advertising.amazon.com/resources/faq#advertising-basics), see the Amazon Ads FAQ. | 
 **mediaType** | [**optional.Interface of MediaType**](.md)| Specifies the media types used to filter the returned array. Currently, only the &#x60;brandLogo&#x60; type is supported. If not specified, all media types are returned. | 

### Return type

[**[]InlineResponse2005**](inline_response_200_5.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.mediaasset.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

