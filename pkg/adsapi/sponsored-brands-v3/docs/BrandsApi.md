# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBrands**](BrandsApi.md#GetBrands) | **Get** /brands | getBrands

# **GetBrands**
> []InlineResponse200 GetBrands(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
getBrands

Gets an array of Brand data objects for the Brand associated with the profile ID passed in the header. For more information about Brands, see [Brand Services](https://brandservices.amazon.com/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***BrandsApiGetBrandsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandsApiGetBrandsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **brandTypeFilter** | [**optional.Interface of BrandType**](.md)| The returned array is filtered to include only brands with brand type set to one of the values in the specified comma-delimited list. Returns all brands if not specified. | 

### Return type

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.brand.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

