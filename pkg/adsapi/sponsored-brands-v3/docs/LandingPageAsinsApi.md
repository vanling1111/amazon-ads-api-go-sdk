# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAsins**](LandingPageAsinsApi.md#ListAsins) | **Get** /pageAsins | Gets ASIN information for a specified address.

# **ListAsins**
> InlineResponse2002 ListAsins(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, pageUrl)
Gets ASIN information for a specified address.

Note that for sellers, the addresss must be a Store page. Vendors may also specify a custom landing page address. Only the `ASINs` that are directly associated on creation with a Store landing page are returned by this API. Any other `ASINs` that are not directly associated to the Store landing page or are listed on other pages are not returned by this API. This includes `ASINs` that are on Dynamic widgets, product selections, and featured deals widgets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **pageUrl** | **string**| For sellers, the address of a Store page. Vendors may also specify the address of a custom landing page. For more information, see the [Stores section](https://advertising.amazon.com/help#GPRM3ZHEXEY5RBFZ) of the Amazon Ads support center. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.pageasins.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

