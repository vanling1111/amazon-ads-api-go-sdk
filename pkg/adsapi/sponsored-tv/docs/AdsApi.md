# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsoredTvAds**](AdsApi.md#CreateSponsoredTvAds) | **Post** /st/ads | 
[**DeleteSponsoredTvAds**](AdsApi.md#DeleteSponsoredTvAds) | **Post** /st/ads/delete | 
[**ListSponsoredTvAds**](AdsApi.md#ListSponsoredTvAds) | **Post** /st/ads/list | 
[**UpdateSponsoredTvAds**](AdsApi.md#UpdateSponsoredTvAds) | **Put** /st/ads | 

# **CreateSponsoredTvAds**
> CreateSponsoredTvAdsResponseContent CreateSponsoredTvAds(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Creates Sponsored Tv Ads.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSponsoredTvAdsRequestContent**](CreateSponsoredTvAdsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***AdsApiCreateSponsoredTvAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdsApiCreateSponsoredTvAdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**CreateSponsoredTvAdsResponseContent**](CreateSponsoredTvAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stAd.v1+json
 - **Accept**: application/vnd.stAd.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSponsoredTvAds**
> DeleteSponsoredTvAdsResponseContent DeleteSponsoredTvAds(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Deletes Sponsored Tv Ads.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteSponsoredTvAdsRequestContent**](DeleteSponsoredTvAdsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***AdsApiDeleteSponsoredTvAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdsApiDeleteSponsoredTvAdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**DeleteSponsoredTvAdsResponseContent**](DeleteSponsoredTvAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stAd.v1+json
 - **Accept**: application/vnd.stAd.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSponsoredTvAds**
> ListSponsoredTvAdsResponseContent ListSponsoredTvAds(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Lists Sponsored Tv Ads.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***AdsApiListSponsoredTvAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdsApiListSponsoredTvAdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ListSponsoredTvAdsRequestContent**](ListSponsoredTvAdsRequestContent.md)|  | 
 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**ListSponsoredTvAdsResponseContent**](ListSponsoredTvAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stAd.v1+json
 - **Accept**: application/vnd.stAd.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSponsoredTvAds**
> UpdateSponsoredTvAdsResponseContent UpdateSponsoredTvAds(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Updates Sponsored Tv Ads.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSponsoredTvAdsRequestContent**](UpdateSponsoredTvAdsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***AdsApiUpdateSponsoredTvAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdsApiUpdateSponsoredTvAdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**UpdateSponsoredTvAdsResponseContent**](UpdateSponsoredTvAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stAd.v1+json
 - **Accept**: application/vnd.stAd.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

