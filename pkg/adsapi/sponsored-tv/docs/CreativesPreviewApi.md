# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreviewSponsoredTvCreative**](CreativesPreviewApi.md#PreviewSponsoredTvCreative) | **Post** /st/creatives/preview | 

# **PreviewSponsoredTvCreative**
> PreviewSponsoredTvCreativeResponseContent PreviewSponsoredTvCreative(ctx, body, optional)


  **Requires one of these permissions**: [\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PreviewSponsoredTvCreativeRequestContent**](PreviewSponsoredTvCreativeRequestContent.md)|  | 
 **optional** | ***CreativesPreviewApiPreviewSponsoredTvCreativeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreativesPreviewApiPreviewSponsoredTvCreativeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amazonAdvertisingAPIClientId** | **optional.**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**PreviewSponsoredTvCreativeResponseContent**](PreviewSponsoredTvCreativeResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.stCreativesPreview.v1+json
 - **Accept**: application/vnd.stCreativesPreview.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

