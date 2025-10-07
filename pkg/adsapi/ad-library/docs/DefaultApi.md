# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdsByID**](DefaultApi.md#GetAdsByID) | **Get** /adRepository/ads/{id} | 
[**ListAds**](DefaultApi.md#ListAds) | **Post** /adRepository/ads/list | 

# **GetAdsByID**
> GetAdsByIdResponseContent GetAdsByID(ctx, id, amazonAdvertisingAPIClientId)


This is the API for retrieving a single advertisement metadata specified by its id from the ad repository.  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| This parameter will limit results to a single advertisement summary matching the requested identifier. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. This is a required header for advertisers and integrators using the Advertising API. | 

### Return type

[**GetAdsByIdResponseContent**](GetAdsByIDResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.adsrepository.v1.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAds**
> ListAdsResponseContent ListAds(ctx, amazonAdvertisingAPIClientId, optional)


This is the primary paginated API for retrieving, listing, and searching through the ad repository. By default, a request without any parameters specified will list all items in the repository.  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. This is a required header for advertisers and integrators using the Advertising API. | 
 **optional** | ***DefaultApiListAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListAdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ListAdsRequestContent**](ListAdsRequestContent.md)|  | 

### Return type

[**ListAdsResponseContent**](ListAdsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.adsrepository.v1.1+json
 - **Accept**: application/vnd.adsrepository.v1.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

