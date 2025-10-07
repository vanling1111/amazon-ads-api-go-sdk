# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsoredTvCreatives**](CreativesApi.md#CreateSponsoredTvCreatives) | **Post** /st/creatives | 
[**ListSponsoredTvCreatives**](CreativesApi.md#ListSponsoredTvCreatives) | **Post** /st/creatives/list | 
[**UpdateSponsoredTvCreatives**](CreativesApi.md#UpdateSponsoredTvCreatives) | **Put** /st/creatives | 

# **CreateSponsoredTvCreatives**
> CreateSponsoredTvCreativesResponseContent CreateSponsoredTvCreatives(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSponsoredTvCreativesRequestContent**](CreateSponsoredTvCreativesRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**CreateSponsoredTvCreativesResponseContent**](CreateSponsoredTvCreativesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stCreative.v1+json
 - **Accept**: application/vnd.stCreative.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSponsoredTvCreatives**
> ListSponsoredTvCreativesResponseContent ListSponsoredTvCreatives(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***CreativesApiListSponsoredTvCreativesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreativesApiListSponsoredTvCreativesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ListSponsoredTvCreativesRequestContent**](ListSponsoredTvCreativesRequestContent.md)|  | 

### Return type

[**ListSponsoredTvCreativesResponseContent**](ListSponsoredTvCreativesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stCreative.v1+json
 - **Accept**: application/vnd.stCreative.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSponsoredTvCreatives**
> UpdateSponsoredTvCreativesResponseContent UpdateSponsoredTvCreatives(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSponsoredTvCreativesRequestContent**](UpdateSponsoredTvCreativesRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**UpdateSponsoredTvCreativesResponseContent**](UpdateSponsoredTvCreativesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stCreative.v1+json
 - **Accept**: application/vnd.stCreative.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

