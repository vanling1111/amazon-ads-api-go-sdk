# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsoredBrandsAdGroups**](AdGroupsApi.md#CreateSponsoredBrandsAdGroups) | **Post** /sb/v4/adGroups | Create ad groups
[**DeleteSponsoredBrandsAdGroups**](AdGroupsApi.md#DeleteSponsoredBrandsAdGroups) | **Post** /sb/v4/adGroups/delete | Delete ad groups
[**ListSponsoredBrandsAdGroups**](AdGroupsApi.md#ListSponsoredBrandsAdGroups) | **Post** /sb/v4/adGroups/list | List ad groups
[**UpdateSponsoredBrandsAdGroups**](AdGroupsApi.md#UpdateSponsoredBrandsAdGroups) | **Put** /sb/v4/adGroups | Update ad groups

# **CreateSponsoredBrandsAdGroups**
> CreateSponsoredBrandsAdGroupsResponseContent CreateSponsoredBrandsAdGroups(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Create ad groups

Creates Sponsored Brands ad groups.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSponsoredBrandsAdGroupsRequestContent**](CreateSponsoredBrandsAdGroupsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**CreateSponsoredBrandsAdGroupsResponseContent**](CreateSponsoredBrandsAdGroupsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbadgroupresource.v4+json
 - **Accept**: application/vnd.sbadgroupresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSponsoredBrandsAdGroups**
> DeleteSponsoredBrandsAdGroupsResponseContent DeleteSponsoredBrandsAdGroups(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Delete ad groups

Deletes Sponsored Brands ad groups.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***AdGroupsApiDeleteSponsoredBrandsAdGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiDeleteSponsoredBrandsAdGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeleteSponsoredBrandsAdGroupsRequestContent**](DeleteSponsoredBrandsAdGroupsRequestContent.md)|  | 

### Return type

[**DeleteSponsoredBrandsAdGroupsResponseContent**](DeleteSponsoredBrandsAdGroupsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbadgroupresource.v4+json
 - **Accept**: application/vnd.sbadgroupresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSponsoredBrandsAdGroups**
> ListSponsoredBrandsAdGroupsResponseContent ListSponsoredBrandsAdGroups(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
List ad groups

Lists Sponsored Brands ad groups.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***AdGroupsApiListSponsoredBrandsAdGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiListSponsoredBrandsAdGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ListSponsoredBrandsAdGroupsRequestContent**](ListSponsoredBrandsAdGroupsRequestContent.md)|  | 

### Return type

[**ListSponsoredBrandsAdGroupsResponseContent**](ListSponsoredBrandsAdGroupsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbadgroupresource.v4+json
 - **Accept**: application/vnd.sbadgroupresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSponsoredBrandsAdGroups**
> UpdateSponsoredBrandsAdGroupsResponseContent UpdateSponsoredBrandsAdGroups(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Update ad groups

Updates Sponsored Brands ad groups.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSponsoredBrandsAdGroupsRequestContent**](UpdateSponsoredBrandsAdGroupsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**UpdateSponsoredBrandsAdGroupsResponseContent**](UpdateSponsoredBrandsAdGroupsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbadgroupresource.v4+json
 - **Accept**: application/vnd.sbadgroupresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

