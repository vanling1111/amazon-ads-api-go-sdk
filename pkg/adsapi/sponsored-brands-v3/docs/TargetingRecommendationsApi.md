# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBrandRecommendations**](TargetingRecommendationsApi.md#GetBrandRecommendations) | **Post** /sb/recommendations/targets/brand | Gets a list of brand suggestions.
[**GetProductRecommendations**](TargetingRecommendationsApi.md#GetProductRecommendations) | **Post** /sb/recommendations/targets/product/list | Gets a list of recommended products for targeting.
[**GetTargetingCategories**](TargetingRecommendationsApi.md#GetTargetingCategories) | **Post** /sb/recommendations/targets/category | Gets a list of recommended categories for targeting.

# **GetBrandRecommendations**
> InlineResponse20016 GetBrandRecommendations(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of brand suggestions.

The Brand suggestions are based on a list of either category identifiers or keywords passed in the request. It is not valid to specify both category identifiers and keywords in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***TargetingRecommendationsApiGetBrandRecommendationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingRecommendationsApiGetBrandRecommendationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of TargetsBrandBody**](TargetsBrandBody.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sbbrandrecommendationsresponse.v3.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductRecommendations**
> InlineResponse20014 GetProductRecommendations(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of recommended products for targeting.

Recommendations are based on the ASINs that are passed in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***TargetingRecommendationsApiGetProductRecommendationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingRecommendationsApiGetProductRecommendationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ProductListBody**](ProductListBody.md)| A list of ASINs. | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sbproductrecommendationsresponse.v3.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargetingCategories**
> InlineResponse20015 GetTargetingCategories(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of recommended categories for targeting.

Recommendations are based on the ASINs that are passed in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***TargetingRecommendationsApiGetTargetingCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingRecommendationsApiGetTargetingCategoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of TargetsCategoryBody**](TargetsCategoryBody.md)| List of ASINs. | 
 **locale** | **optional.**| Return the categories in the specified locale. | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sbcategoryrecommendationsresponse.v3.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

