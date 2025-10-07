# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RASv1CreateAdGroups**](AdGroupsApi.md#RASv1CreateAdGroups) | **Post** /ras/v1/adGroups | 
[**RASv1DeleteAdGroups**](AdGroupsApi.md#RASv1DeleteAdGroups) | **Post** /ras/v1/adGroups/delete | 
[**RASv1ListAdGroups**](AdGroupsApi.md#RASv1ListAdGroups) | **Post** /ras/v1/adGroups/list | 
[**RASv1UpdateAdGroups**](AdGroupsApi.md#RASv1UpdateAdGroups) | **Put** /ras/v1/adGroups | 

# **RASv1CreateAdGroups**
> Rasv1CreateAdGroupsResponseContent RASv1CreateAdGroups(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rasv1CreateAdGroupsRequestContent**](Rasv1CreateAdGroupsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***AdGroupsApiRASv1CreateAdGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiRASv1CreateAdGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. | 

### Return type

[**Rasv1CreateAdGroupsResponseContent**](RASv1CreateAdGroupsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RASv1DeleteAdGroups**
> Rasv1DeleteAdGroupsResponseContent RASv1DeleteAdGroups(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rasv1DeleteAdGroupsRequestContent**](Rasv1DeleteAdGroupsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**Rasv1DeleteAdGroupsResponseContent**](RASv1DeleteAdGroupsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RASv1ListAdGroups**
> Rasv1ListAdGroupsResponseContent RASv1ListAdGroups(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***AdGroupsApiRASv1ListAdGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiRASv1ListAdGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Rasv1ListAdGroupsRequestContent**](Rasv1ListAdGroupsRequestContent.md)|  | 

### Return type

[**Rasv1ListAdGroupsResponseContent**](RASv1ListAdGroupsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RASv1UpdateAdGroups**
> Rasv1UpdateAdGroupsResponseContent RASv1UpdateAdGroups(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rasv1UpdateAdGroupsRequestContent**](Rasv1UpdateAdGroupsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***AdGroupsApiRASv1UpdateAdGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiRASv1UpdateAdGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. | 

### Return type

[**Rasv1UpdateAdGroupsResponseContent**](RASv1UpdateAdGroupsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

