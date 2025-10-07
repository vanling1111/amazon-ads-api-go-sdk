# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAd**](AdsApi.md#CreateAd) | **Post** /adsApi/v1/create/ads | 
[**DeleteAd**](AdsApi.md#DeleteAd) | **Post** /adsApi/v1/delete/ads | 
[**QueryAd**](AdsApi.md#QueryAd) | **Post** /adsApi/v1/query/ads | 
[**UpdateAd**](AdsApi.md#UpdateAd) | **Post** /adsApi/v1/update/ads | 

# **CreateAd**
> AdMultiStatusResponseWithPartialErrors CreateAd(ctx, amazonAdsClientId, optional)


Creates Ads.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"], [\"creatives_edit\"], [\"advertiser_campaign_edit\",\"creatives_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***AdsApiCreateAdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdsApiCreateAdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateAdRequest**](CreateAdRequest.md)|  | 
 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**AdMultiStatusResponseWithPartialErrors**](AdMultiStatusResponseWithPartialErrors.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAd**
> AdMultiStatusResponseWithPartialErrors DeleteAd(ctx, amazonAdsClientId, optional)


Archives or deletes Ads.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***AdsApiDeleteAdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdsApiDeleteAdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DeleteAdRequest**](DeleteAdRequest.md)|  | 
 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**AdMultiStatusResponseWithPartialErrors**](AdMultiStatusResponseWithPartialErrors.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryAd**
> AdSuccessResponse QueryAd(ctx, body, amazonAdsClientId, optional)


A search read, allowing use of more complex filters.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"], [\"creatives_view\"], [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"creatives_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**QueryAdRequest**](QueryAdRequest.md)|  | 
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***AdsApiQueryAdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdsApiQueryAdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**AdSuccessResponse**](AdSuccessResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAd**
> AdMultiStatusResponseWithPartialErrors UpdateAd(ctx, amazonAdsClientId, optional)


Updates Ads. Behaves as PATCH unless otherwise stated.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"], [\"creatives_edit\"], [\"advertiser_campaign_edit\",\"creatives_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***AdsApiUpdateAdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdsApiUpdateAdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateAdRequest**](UpdateAdRequest.md)|  | 
 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**AdMultiStatusResponseWithPartialErrors**](AdMultiStatusResponseWithPartialErrors.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

