# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsoredBrandsCampaigns**](CampaignsApi.md#CreateSponsoredBrandsCampaigns) | **Post** /sb/v4/campaigns | Create campaigns
[**DeleteSponsoredBrandsCampaigns**](CampaignsApi.md#DeleteSponsoredBrandsCampaigns) | **Post** /sb/v4/campaigns/delete | Delete campaigns.
[**ListSponsoredBrandsCampaigns**](CampaignsApi.md#ListSponsoredBrandsCampaigns) | **Post** /sb/v4/campaigns/list | List campaigns
[**UpdateSponsoredBrandsCampaigns**](CampaignsApi.md#UpdateSponsoredBrandsCampaigns) | **Put** /sb/v4/campaigns | Update campaigns

# **CreateSponsoredBrandsCampaigns**
> CreateSponsoredBrandsCampaignsResponseContent CreateSponsoredBrandsCampaigns(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Create campaigns

Creates Sponsored Brands campaigns.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSponsoredBrandsCampaignsRequestContent**](CreateSponsoredBrandsCampaignsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**CreateSponsoredBrandsCampaignsResponseContent**](CreateSponsoredBrandsCampaignsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbcampaignresource.v4+json
 - **Accept**: application/vnd.sbcampaignresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSponsoredBrandsCampaigns**
> DeleteSponsoredBrandsCampaignsResponseContent DeleteSponsoredBrandsCampaigns(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Delete campaigns.

Deletes Sponsored Brands campaigns.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***CampaignsApiDeleteSponsoredBrandsCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiDeleteSponsoredBrandsCampaignsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeleteSponsoredBrandsCampaignsRequestContent**](DeleteSponsoredBrandsCampaignsRequestContent.md)|  | 

### Return type

[**DeleteSponsoredBrandsCampaignsResponseContent**](DeleteSponsoredBrandsCampaignsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbcampaignresource.v4+json
 - **Accept**: application/vnd.sbcampaignresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSponsoredBrandsCampaigns**
> ListSponsoredBrandsCampaignsResponseContent ListSponsoredBrandsCampaigns(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
List campaigns

Lists Sponsored Brands campaigns.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***CampaignsApiListSponsoredBrandsCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiListSponsoredBrandsCampaignsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ListSponsoredBrandsCampaignsRequestContent**](ListSponsoredBrandsCampaignsRequestContent.md)|  | 

### Return type

[**ListSponsoredBrandsCampaignsResponseContent**](ListSponsoredBrandsCampaignsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbcampaignresource.v4+json
 - **Accept**: application/vnd.sbcampaignresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSponsoredBrandsCampaigns**
> UpdateSponsoredBrandsCampaignsResponseContent UpdateSponsoredBrandsCampaigns(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Update campaigns

Updates Sponsored Brands campaigns.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSponsoredBrandsCampaignsRequestContent**](UpdateSponsoredBrandsCampaignsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**UpdateSponsoredBrandsCampaignsResponseContent**](UpdateSponsoredBrandsCampaignsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbcampaignresource.v4+json
 - **Accept**: application/vnd.sbcampaignresource.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

