# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkGetAsinSuggestedKeywords**](SuggestedKeywordsApi.md#BulkGetAsinSuggestedKeywords) | **Post** /v2/sp/asins/suggested/keywords | Gets suggested keyword for a specified list of ASINs.
[**GetAdGroupSuggestedKeywords**](SuggestedKeywordsApi.md#GetAdGroupSuggestedKeywords) | **Get** /v2/sp/adGroups/{adGroupId}/suggested/keywords | Gets suggested keywords for the specified ad group.
[**GetAdGroupSuggestedKeywordsEx**](SuggestedKeywordsApi.md#GetAdGroupSuggestedKeywordsEx) | **Get** /v2/sp/adGroups/{adGroupId}/suggested/keywords/extended | Gets suggested keywords with extended data for the specified ad group.
[**GetAsinSuggestedKeywords**](SuggestedKeywordsApi.md#GetAsinSuggestedKeywords) | **Get** /v2/sp/asins/{asinValue}/suggested/keywords | Gets suggested keywords for the specified ASIN.

# **BulkGetAsinSuggestedKeywords**
> []BulkGetAsinSuggestedKeywordsResponseInner BulkGetAsinSuggestedKeywords(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets suggested keyword for a specified list of ASINs.

Suggested keywords are returned in an array ordered by descending effectiveness.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; developer account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***SuggestedKeywordsApiBulkGetAsinSuggestedKeywordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SuggestedKeywordsApiBulkGetAsinSuggestedKeywordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SuggestedKeywordsBody**](SuggestedKeywordsBody.md)|  | 

### Return type

[**[]BulkGetAsinSuggestedKeywordsResponseInner**](array.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdGroupSuggestedKeywords**
> AdGroupSuggestedKeywordsResponse GetAdGroupSuggestedKeywords(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, adGroupId, optional)
Gets suggested keywords for the specified ad group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; developer account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **adGroupId** | **float64**| The identifier of a valid ad group. | 
 **optional** | ***SuggestedKeywordsApiGetAdGroupSuggestedKeywordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SuggestedKeywordsApiGetAdGroupSuggestedKeywordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxNumSuggestions** | **optional.Int32**| The maximum number of suggested keywords for the response. | [default to 100]
 **adStateFilter** | **optional.String**| Filters results to ad groups with state matching the comma-delimited list. | 

### Return type

[**AdGroupSuggestedKeywordsResponse**](AdGroupSuggestedKeywordsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdGroupSuggestedKeywordsEx**
> []AdGroupSuggestedKeywordsResponseEx GetAdGroupSuggestedKeywordsEx(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, adGroupId, optional)
Gets suggested keywords with extended data for the specified ad group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; developer account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **adGroupId** | **float64**| The identifier of a valid ad group. | 
 **optional** | ***SuggestedKeywordsApiGetAdGroupSuggestedKeywordsExOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SuggestedKeywordsApiGetAdGroupSuggestedKeywordsExOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxNumSuggestions** | **optional.Int32**| The maximum number of suggested keywords for the response. | [default to 100]
 **suggestBids** | **optional.String**| Set to &#x60;yes&#x60; to include a suggest bid for the suggested keyword in the response. Otherwise, set to &#x60;no&#x60;. | [default to no]
 **adStateFilter** | **optional.String**| Filters results to ad groups with state matching the comma-delimited list. | 

### Return type

[**[]AdGroupSuggestedKeywordsResponseEx**](AdGroupSuggestedKeywordsResponseEx.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsinSuggestedKeywords**
> GetAsinSuggestedKeywordsResponse GetAsinSuggestedKeywords(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, asinValue, optional)
Gets suggested keywords for the specified ASIN.

Suggested keywords are returned in an array ordered by descending effectiveness.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; developer account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **asinValue** | **string**| An ASIN. | 
 **optional** | ***SuggestedKeywordsApiGetAsinSuggestedKeywordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SuggestedKeywordsApiGetAsinSuggestedKeywordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxNumSuggestions** | **optional.Int32**| The maximum number of suggested keywords for the response. | [default to 100]

### Return type

[**GetAsinSuggestedKeywordsResponse**](GetAsinSuggestedKeywordsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

