# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RASv1CreateTargets**](TargetsApi.md#RASv1CreateTargets) | **Post** /ras/v1/targets | 
[**RASv1DeleteTargets**](TargetsApi.md#RASv1DeleteTargets) | **Post** /ras/v1/targets/delete | 
[**RASv1ListTargets**](TargetsApi.md#RASv1ListTargets) | **Post** /ras/v1/targets/list | 
[**RASv1UpdateTargets**](TargetsApi.md#RASv1UpdateTargets) | **Put** /ras/v1/targets | 

# **RASv1CreateTargets**
> Rasv1CreateTargetsResponseContent RASv1CreateTargets(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rasv1CreateTargetsRequestContent**](Rasv1CreateTargetsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***TargetsApiRASv1CreateTargetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetsApiRASv1CreateTargetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. | 

### Return type

[**Rasv1CreateTargetsResponseContent**](RASv1CreateTargetsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RASv1DeleteTargets**
> Rasv1DeleteTargetsResponseContent RASv1DeleteTargets(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rasv1DeleteTargetsRequestContent**](Rasv1DeleteTargetsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**Rasv1DeleteTargetsResponseContent**](RASv1DeleteTargetsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RASv1ListTargets**
> Rasv1ListTargetsResponseContent RASv1ListTargets(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***TargetsApiRASv1ListTargetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetsApiRASv1ListTargetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Rasv1ListTargetsRequestContent**](Rasv1ListTargetsRequestContent.md)|  | 

### Return type

[**Rasv1ListTargetsResponseContent**](RASv1ListTargetsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RASv1UpdateTargets**
> Rasv1UpdateTargetsResponseContent RASv1UpdateTargets(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rasv1UpdateTargetsRequestContent**](Rasv1UpdateTargetsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***TargetsApiRASv1UpdateTargetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetsApiRASv1UpdateTargetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. | 

### Return type

[**Rasv1UpdateTargetsResponseContent**](RASv1UpdateTargetsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

