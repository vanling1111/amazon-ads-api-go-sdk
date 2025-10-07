# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsoredBrandStoreSpotlightAds**](AdsApi.md#CreateSponsoredBrandStoreSpotlightAds) | **Post** /sb/v4/ads/storeSpotlight | Create store spotlight ads
[**CreateSponsoredBrandsBrandVideoAds**](AdsApi.md#CreateSponsoredBrandsBrandVideoAds) | **Post** /sb/v4/ads/brandVideo | Create brand video ads
[**CreateSponsoredBrandsExtendedProductCollectionAds**](AdsApi.md#CreateSponsoredBrandsExtendedProductCollectionAds) | **Post** /sb/v4/ads/productCollectionExtended | Creates Sponsored Brands new product collection ads with collection of custom images
[**CreateSponsoredBrandsProductCollectionAds**](AdsApi.md#CreateSponsoredBrandsProductCollectionAds) | **Post** /sb/v4/ads/productCollection | Create product collection ads
[**CreateSponsoredBrandsVideoAds**](AdsApi.md#CreateSponsoredBrandsVideoAds) | **Post** /sb/v4/ads/video | Create video ads
[**DeleteSponsoredBrandsAds**](AdsApi.md#DeleteSponsoredBrandsAds) | **Post** /sb/v4/ads/delete | Delete ads
[**ListSponsoredBrandsAds**](AdsApi.md#ListSponsoredBrandsAds) | **Post** /sb/v4/ads/list | List ads
[**UpdateSponsoredBrandsAds**](AdsApi.md#UpdateSponsoredBrandsAds) | **Put** /sb/v4/ads | Update ads

# **CreateSponsoredBrandStoreSpotlightAds**
> CreateSponsoredBrandStoreSpotlightAdsResponseContent CreateSponsoredBrandStoreSpotlightAds(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Create store spotlight ads

Creates Sponsored Brands store spotlight ads.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSponsoredBrandStoreSpotlightAdsRequestContent**](CreateSponsoredBrandStoreSpotlightAdsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**CreateSponsoredBrandStoreSpotlightAdsResponseContent**](CreateSponsoredBrandStoreSpotlightAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbadresource.v4+json
 - **Accept**: application/vnd.sbadresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSponsoredBrandsBrandVideoAds**
> CreateSponsoredBrandsBrandVideoAdsResponseContent CreateSponsoredBrandsBrandVideoAds(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Create brand video ads

Creates Sponsored Brands brand video ads.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSponsoredBrandsBrandVideoAdsRequestContent**](CreateSponsoredBrandsBrandVideoAdsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**CreateSponsoredBrandsBrandVideoAdsResponseContent**](CreateSponsoredBrandsBrandVideoAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbadresource.v4+json
 - **Accept**: application/vnd.sbadresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSponsoredBrandsExtendedProductCollectionAds**
> CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent CreateSponsoredBrandsExtendedProductCollectionAds(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Creates Sponsored Brands new product collection ads with collection of custom images

Creates Sponsored Brands product collection ads with collection of custom images[1-5].  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent**](CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent**](CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbadresource.v4+json
 - **Accept**: application/vnd.sbadresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSponsoredBrandsProductCollectionAds**
> CreateSponsoredBrandsProductCollectionAdsResponseContent CreateSponsoredBrandsProductCollectionAds(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Create product collection ads

Creates Sponsored Brands product collection ads.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSponsoredBrandsProductCollectionAdsRequestContent**](CreateSponsoredBrandsProductCollectionAdsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**CreateSponsoredBrandsProductCollectionAdsResponseContent**](CreateSponsoredBrandsProductCollectionAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbadresource.v4+json
 - **Accept**: application/vnd.sbadresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSponsoredBrandsVideoAds**
> CreateSponsoredBrandsVideoAdsResponseContent CreateSponsoredBrandsVideoAds(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Create video ads

Creates Sponsored Brands video ads.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSponsoredBrandsVideoAdsRequestContent**](CreateSponsoredBrandsVideoAdsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**CreateSponsoredBrandsVideoAdsResponseContent**](CreateSponsoredBrandsVideoAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbadresource.v4+json
 - **Accept**: application/vnd.sbadresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSponsoredBrandsAds**
> DeleteSponsoredBrandsAdsResponseContent DeleteSponsoredBrandsAds(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Delete ads

Deletes Sponsored Brands ads.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***AdsApiDeleteSponsoredBrandsAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdsApiDeleteSponsoredBrandsAdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeleteSponsoredBrandsAdsRequestContent**](DeleteSponsoredBrandsAdsRequestContent.md)|  | 

### Return type

[**DeleteSponsoredBrandsAdsResponseContent**](DeleteSponsoredBrandsAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbadresource.v4+json
 - **Accept**: application/vnd.sbadresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSponsoredBrandsAds**
> ListSponsoredBrandsAdsResponseContent ListSponsoredBrandsAds(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
List ads

Lists Sponsored Brands ads.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***AdsApiListSponsoredBrandsAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdsApiListSponsoredBrandsAdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ListSponsoredBrandsAdsRequestContent**](ListSponsoredBrandsAdsRequestContent.md)|  | 

### Return type

[**ListSponsoredBrandsAdsResponseContent**](ListSponsoredBrandsAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbadresource.v4+json
 - **Accept**: application/vnd.sbadresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSponsoredBrandsAds**
> UpdateSponsoredBrandsAdsResponseContent UpdateSponsoredBrandsAds(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Update ads

Updates Sponsored Brands ads.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSponsoredBrandsAdsRequestContent**](UpdateSponsoredBrandsAdsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**UpdateSponsoredBrandsAdsResponseContent**](UpdateSponsoredBrandsAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbadresource.v4+json
 - **Accept**: application/vnd.sbadresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

