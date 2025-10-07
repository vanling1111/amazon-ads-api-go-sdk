# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SbModerationCampaignsCampaignIdGet**](ModerationApi.md#SbModerationCampaignsCampaignIdGet) | **Get** /sb/moderation/campaigns/{campaignId} | Gets the moderation result for a campaign specified by identifier.

# **SbModerationCampaignsCampaignIdGet**
> InlineResponse20017 SbModerationCampaignsCampaignIdGet(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, campaignId)
Gets the moderation result for a campaign specified by identifier.

Note that this resource is only available for campaigns in the US marketplace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **campaignId** | **int64**| The campaign identifier. | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbmoderation.v3+json, application/vnd.error.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

