# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaign**](CampaignsApi.md#CreateCampaign) | **Post** /adsApi/v1/create/campaigns | 
[**DeleteCampaign**](CampaignsApi.md#DeleteCampaign) | **Post** /adsApi/v1/delete/campaigns | 
[**QueryCampaign**](CampaignsApi.md#QueryCampaign) | **Post** /adsApi/v1/query/campaigns | 
[**UpdateCampaign**](CampaignsApi.md#UpdateCampaign) | **Post** /adsApi/v1/update/campaigns | 

# **CreateCampaign**
> CampaignMultiStatusResponseWithPartialErrors CreateCampaign(ctx, amazonAdsClientId, optional)


Creates Campaigns.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"], [\"campaign_edit\"], [\"advertiser_campaign_edit\",\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***CampaignsApiCreateCampaignOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiCreateCampaignOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateCampaignRequest**](CreateCampaignRequest.md)|  | 
 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**CampaignMultiStatusResponseWithPartialErrors**](CampaignMultiStatusResponseWithPartialErrors.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCampaign**
> CampaignMultiStatusResponseWithPartialErrors DeleteCampaign(ctx, amazonAdsClientId, optional)


Archives or deletes Campaigns.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***CampaignsApiDeleteCampaignOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiDeleteCampaignOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DeleteCampaignRequest**](DeleteCampaignRequest.md)|  | 
 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**CampaignMultiStatusResponseWithPartialErrors**](CampaignMultiStatusResponseWithPartialErrors.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryCampaign**
> CampaignSuccessResponse QueryCampaign(ctx, body, amazonAdsClientId, optional)


A search read, allowing use of more complex filters.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"], [\"campaign_view\"], [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**QueryCampaignRequest**](QueryCampaignRequest.md)|  | 
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***CampaignsApiQueryCampaignOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiQueryCampaignOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**CampaignSuccessResponse**](CampaignSuccessResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCampaign**
> CampaignMultiStatusResponseWithPartialErrors UpdateCampaign(ctx, amazonAdsClientId, optional)


Updates Campaigns. Behaves as PATCH unless otherwise stated.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"], [\"campaign_edit\"], [\"advertiser_campaign_edit\",\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***CampaignsApiUpdateCampaignOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiUpdateCampaignOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateCampaignRequest**](UpdateCampaignRequest.md)|  | 
 **amazonAdsAccountId** | **optional.**| The identifier of an Amazon Ads Advertiser Account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 

### Return type

[**CampaignMultiStatusResponseWithPartialErrors**](CampaignMultiStatusResponseWithPartialErrors.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

