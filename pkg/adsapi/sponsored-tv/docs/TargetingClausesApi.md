# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsoredTvTargetingClauses**](TargetingClausesApi.md#CreateSponsoredTvTargetingClauses) | **Post** /st/targets | 
[**DeleteSponsoredTvTargetingClauses**](TargetingClausesApi.md#DeleteSponsoredTvTargetingClauses) | **Post** /st/targets/delete | 
[**ListSponsoredTvTargetingClauses**](TargetingClausesApi.md#ListSponsoredTvTargetingClauses) | **Post** /st/targets/list | 
[**UpdateSponsoredTvTargetingClauses**](TargetingClausesApi.md#UpdateSponsoredTvTargetingClauses) | **Put** /st/targets | 

# **CreateSponsoredTvTargetingClauses**
> CreateSponsoredTvTargetingClausesResponseContent CreateSponsoredTvTargetingClauses(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSponsoredTvTargetingClausesRequestContent**](CreateSponsoredTvTargetingClausesRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***TargetingClausesApiCreateSponsoredTvTargetingClausesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingClausesApiCreateSponsoredTvTargetingClausesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**CreateSponsoredTvTargetingClausesResponseContent**](CreateSponsoredTvTargetingClausesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stTargetingClause.v1+json
 - **Accept**: application/vnd.stTargetingClause.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSponsoredTvTargetingClauses**
> DeleteSponsoredTvTargetingClausesResponseContent DeleteSponsoredTvTargetingClauses(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteSponsoredTvTargetingClausesRequestContent**](DeleteSponsoredTvTargetingClausesRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***TargetingClausesApiDeleteSponsoredTvTargetingClausesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingClausesApiDeleteSponsoredTvTargetingClausesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**DeleteSponsoredTvTargetingClausesResponseContent**](DeleteSponsoredTvTargetingClausesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stTargetingClause.v1+json
 - **Accept**: application/vnd.stTargetingClause.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSponsoredTvTargetingClauses**
> ListSponsoredTvTargetingClausesResponseContent ListSponsoredTvTargetingClauses(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***TargetingClausesApiListSponsoredTvTargetingClausesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingClausesApiListSponsoredTvTargetingClausesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ListSponsoredTvTargetingClausesRequestContent**](ListSponsoredTvTargetingClausesRequestContent.md)|  | 
 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**ListSponsoredTvTargetingClausesResponseContent**](ListSponsoredTvTargetingClausesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stTargetingClause.v1+json
 - **Accept**: application/vnd.stTargetingClause.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSponsoredTvTargetingClauses**
> UpdateSponsoredTvTargetingClausesResponseContent UpdateSponsoredTvTargetingClauses(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSponsoredTvTargetingClausesRequestContent**](UpdateSponsoredTvTargetingClausesRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***TargetingClausesApiUpdateSponsoredTvTargetingClausesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingClausesApiUpdateSponsoredTvTargetingClausesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**UpdateSponsoredTvTargetingClausesResponseContent**](UpdateSponsoredTvTargetingClausesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stTargetingClause.v1+json
 - **Accept**: application/vnd.stTargetingClause.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

