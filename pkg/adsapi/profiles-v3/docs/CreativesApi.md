# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCreatives**](CreativesApi.md#CreateCreatives) | **Post** /sd/creatives | A POST request of one or more creatives.
[**ListCreativeModerations**](CreativesApi.md#ListCreativeModerations) | **Get** /sd/moderation/creatives | Gets a list of creative moderations
[**ListCreatives**](CreativesApi.md#ListCreatives) | **Get** /sd/creatives | Gets a list of creatives
[**PostCreativePreview**](CreativesApi.md#PostCreativePreview) | **Post** /sd/creatives/preview | Gets creative preview HTML.
[**UpdateCreatives**](CreativesApi.md#UpdateCreatives) | **Put** /sd/creatives | Updates one or more creatives.

# **CreateCreatives**
> []CreativeResponse CreateCreatives(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
A POST request of one or more creatives.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***CreativesApiCreateCreativesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreativesApiCreateCreativesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []CreateCreative**](CreateCreative.md)| An array of Creative objects to create. Maximum length of the array is 100 objects. Note - when using productAds with landingPageURL of OFF_AMAZON_LINK, STORE, or MOMENT, the following properties are required all together;
1) headline, 2) brandLogo, and 3) rectCustomImage, squareCustomImage. | 

### Return type

[**[]CreativeResponse**](CreativeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCreativeModerations**
> []CreativeModeration ListCreativeModerations(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, language, optional)
Gets a list of creative moderations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **language** | [**Locale**](.md)| The language of the returned creative moderation metadata. | 
 **optional** | ***CreativesApiListCreativeModerationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreativesApiListCreativeModerationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **startIndex** | **optional.Int32**| Sets a cursor into the requested set of creative moderations. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0. | [default to 0]
 **count** | **optional.Int32**| Sets the number of creative objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten creative moderations set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten creative moderations, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size. | [default to 100]
 **adGroupIdFilter** | **optional.String**| The returned array includes only creative moderations associated with ad group identifiers matching those specified in the comma-delimited string. Cannot be used in conjunction with the &#x60;creativeIdFilter&#x60; parameter. | 
 **creativeIdFilter** | **optional.String**| The returned array includes only creative moderations with creative identifiers matching those specified in the comma-delimited string. Cannot be used in conjunction with the &#x60;adGroupIdFilter&#x60; parameter. | 

### Return type

[**[]CreativeModeration**](CreativeModeration.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCreatives**
> []Creative ListCreatives(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of creatives

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***CreativesApiListCreativesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreativesApiListCreativesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Sets a cursor into the requested set of creatives. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0. | [default to 0]
 **count** | **optional.Int32**| Sets the number of creative objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten creatives set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten creatives, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size. | [default to 100]
 **adGroupIdFilter** | **optional.String**| The returned array includes only creatives associated with ad group identifiers matching those specified in the comma-delimited string. Cannot be used in conjunction with the &#x60;creativeIdFilter&#x60; parameter. | 
 **creativeIdFilter** | **optional.String**| The returned array includes only creatives with identifiers matching those specified in the comma-delimited string. Cannot be used in conjunction with the &#x60;adGroupIdFilter&#x60; parameter. | 

### Return type

[**[]Creative**](Creative.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCreativePreview**
> CreativePreviewResponse PostCreativePreview(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets creative preview HTML.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***CreativesApiPostCreativePreviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreativesApiPostCreativePreviewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of CreativePreviewRequest**](CreativePreviewRequest.md)|  | 

### Return type

[**CreativePreviewResponse**](CreativePreviewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCreatives**
> []CreativeResponse UpdateCreatives(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Updates one or more creatives.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***CreativesApiUpdateCreativesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreativesApiUpdateCreativesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []CreativeUpdate**](CreativeUpdate.md)| An array of creative objects to update. Maximum length of the array is 100 objects. | 

### Return type

[**[]CreativeResponse**](CreativeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

