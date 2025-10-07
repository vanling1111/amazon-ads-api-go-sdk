# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdGroup**](AdGroupsApi.md#CreateAdGroup) | **Post** /adsApi/v1/create/adGroups | 
[**DeleteAdGroup**](AdGroupsApi.md#DeleteAdGroup) | **Post** /adsApi/v1/delete/adGroups | 
[**QueryAdGroup**](AdGroupsApi.md#QueryAdGroup) | **Post** /adsApi/v1/query/adGroups | 
[**UpdateAdGroup**](AdGroupsApi.md#UpdateAdGroup) | **Post** /adsApi/v1/update/adGroups | 

# **CreateAdGroup**
> AdGroupMultiStatusResponseWithPartialErrors CreateAdGroup(ctx, amazonAdsClientId, optional)


Creates AdGroups.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"], [\"campaign_edit\"], [\"advertiser_campaign_edit\",\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***AdGroupsApiCreateAdGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiCreateAdGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateAdGroupRequest**](CreateAdGroupRequest.md)|  | 
 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**AdGroupMultiStatusResponseWithPartialErrors**](AdGroupMultiStatusResponseWithPartialErrors.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAdGroup**
> AdGroupMultiStatusResponseWithPartialErrors DeleteAdGroup(ctx, amazonAdsClientId, optional)


Archives or deletes AdGroups.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***AdGroupsApiDeleteAdGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiDeleteAdGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DeleteAdGroupRequest**](DeleteAdGroupRequest.md)|  | 
 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**AdGroupMultiStatusResponseWithPartialErrors**](AdGroupMultiStatusResponseWithPartialErrors.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryAdGroup**
> AdGroupSuccessResponse QueryAdGroup(ctx, body, amazonAdsClientId, optional)


A search read, allowing use of more complex filters.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"], [\"campaign_view\"], [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**QueryAdGroupRequest**](QueryAdGroupRequest.md)|  | 
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***AdGroupsApiQueryAdGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiQueryAdGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**AdGroupSuccessResponse**](AdGroupSuccessResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAdGroup**
> AdGroupMultiStatusResponseWithPartialErrors UpdateAdGroup(ctx, amazonAdsClientId, optional)


Updates AdGroups. Behaves as PATCH unless otherwise stated.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"], [\"campaign_edit\"], [\"advertiser_campaign_edit\",\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***AdGroupsApiUpdateAdGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiUpdateAdGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateAdGroupRequest**](UpdateAdGroupRequest.md)|  | 
 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**AdGroupMultiStatusResponseWithPartialErrors**](AdGroupMultiStatusResponseWithPartialErrors.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

