# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsoredTvAdGroups**](AdGroupsApi.md#CreateSponsoredTvAdGroups) | **Post** /st/adGroups | 
[**DeleteSponsoredTvAdGroups**](AdGroupsApi.md#DeleteSponsoredTvAdGroups) | **Post** /st/adGroups/delete | 
[**ListSponsoredTvAdGroups**](AdGroupsApi.md#ListSponsoredTvAdGroups) | **Post** /st/adGroups/list | 
[**UpdateSponsoredTvAdGroups**](AdGroupsApi.md#UpdateSponsoredTvAdGroups) | **Put** /st/adGroups | 

# **CreateSponsoredTvAdGroups**
> CreateSponsoredTvAdGroupsResponseContent CreateSponsoredTvAdGroups(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Creates Sponsored Tv Ad Groups.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSponsoredTvAdGroupsRequestContent**](CreateSponsoredTvAdGroupsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***AdGroupsApiCreateSponsoredTvAdGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiCreateSponsoredTvAdGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**CreateSponsoredTvAdGroupsResponseContent**](CreateSponsoredTvAdGroupsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stAdGroup.v1+json
 - **Accept**: application/vnd.stAdGroup.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSponsoredTvAdGroups**
> DeleteSponsoredTvAdGroupsResponseContent DeleteSponsoredTvAdGroups(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Deletes Sponsored Tv Ad Groups.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteSponsoredTvAdGroupsRequestContent**](DeleteSponsoredTvAdGroupsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***AdGroupsApiDeleteSponsoredTvAdGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiDeleteSponsoredTvAdGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**DeleteSponsoredTvAdGroupsResponseContent**](DeleteSponsoredTvAdGroupsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stAdGroup.v1+json
 - **Accept**: application/vnd.stAdGroup.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSponsoredTvAdGroups**
> ListSponsoredTvAdGroupsResponseContent ListSponsoredTvAdGroups(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Lists Sponsored Tv Ad Groups.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***AdGroupsApiListSponsoredTvAdGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiListSponsoredTvAdGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ListSponsoredTvAdGroupsRequestContent**](ListSponsoredTvAdGroupsRequestContent.md)|  | 
 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**ListSponsoredTvAdGroupsResponseContent**](ListSponsoredTvAdGroupsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stAdGroup.v1+json
 - **Accept**: application/vnd.stAdGroup.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSponsoredTvAdGroups**
> UpdateSponsoredTvAdGroupsResponseContent UpdateSponsoredTvAdGroups(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Updates Sponsored Tv Ad Groups.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSponsoredTvAdGroupsRequestContent**](UpdateSponsoredTvAdGroupsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***AdGroupsApiUpdateSponsoredTvAdGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiUpdateSponsoredTvAdGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**UpdateSponsoredTvAdGroupsResponseContent**](UpdateSponsoredTvAdGroupsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stAdGroup.v1+json
 - **Accept**: application/vnd.stAdGroup.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

