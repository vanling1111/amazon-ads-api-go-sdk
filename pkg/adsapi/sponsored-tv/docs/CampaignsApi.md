# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsoredTvCampaigns**](CampaignsApi.md#CreateSponsoredTvCampaigns) | **Post** /st/campaigns | 
[**DeleteSponsoredTvCampaigns**](CampaignsApi.md#DeleteSponsoredTvCampaigns) | **Post** /st/campaigns/delete | 
[**ListSponsoredTvCampaigns**](CampaignsApi.md#ListSponsoredTvCampaigns) | **Post** /st/campaigns/list | 
[**UpdateSponsoredTvCampaigns**](CampaignsApi.md#UpdateSponsoredTvCampaigns) | **Put** /st/campaigns | 

# **CreateSponsoredTvCampaigns**
> CreateSponsoredTvCampaignsResponseContent CreateSponsoredTvCampaigns(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Creates Sponsored Tv campaigns.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSponsoredTvCampaignsRequestContent**](CreateSponsoredTvCampaignsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***CampaignsApiCreateSponsoredTvCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiCreateSponsoredTvCampaignsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**CreateSponsoredTvCampaignsResponseContent**](CreateSponsoredTvCampaignsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stCampaign.v1+json
 - **Accept**: application/vnd.stCampaign.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSponsoredTvCampaigns**
> DeleteSponsoredTvCampaignsResponseContent DeleteSponsoredTvCampaigns(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Deletes Sponsored Tv campaigns.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteSponsoredTvCampaignsRequestContent**](DeleteSponsoredTvCampaignsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***CampaignsApiDeleteSponsoredTvCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiDeleteSponsoredTvCampaignsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**DeleteSponsoredTvCampaignsResponseContent**](DeleteSponsoredTvCampaignsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stCampaign.v1+json
 - **Accept**: application/vnd.stCampaign.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSponsoredTvCampaigns**
> ListSponsoredTvCampaignsResponseContent ListSponsoredTvCampaigns(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Lists Sponsored Tv Campaigns  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***CampaignsApiListSponsoredTvCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiListSponsoredTvCampaignsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ListSponsoredTvCampaignsRequestContent**](ListSponsoredTvCampaignsRequestContent.md)|  | 
 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**ListSponsoredTvCampaignsResponseContent**](ListSponsoredTvCampaignsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stCampaign.v1+json
 - **Accept**: application/vnd.stCampaign.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSponsoredTvCampaigns**
> UpdateSponsoredTvCampaignsResponseContent UpdateSponsoredTvCampaigns(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Updates Sponsored Tv campaigns.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSponsoredTvCampaignsRequestContent**](UpdateSponsoredTvCampaignsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***CampaignsApiUpdateSponsoredTvCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiUpdateSponsoredTvCampaignsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefer** | **optional.**| The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences for operations are as follows: List Operation - supports &#x60;include&#x3D;extendedData&#x60; preference to return representation with extendedData Create/Update/Delete Operations - supports &#x60;return&#x3D;representation&#x60; to return the full object. &#x60;include&#x3D;extendedData&#x60; can be used with &#x60;return&#x3D;representation&#x60; to return the representation with extendedData. | 

### Return type

[**UpdateSponsoredTvCampaignsResponseContent**](UpdateSponsoredTvCampaignsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stCampaign.v1+json
 - **Accept**: application/vnd.stCampaign.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

