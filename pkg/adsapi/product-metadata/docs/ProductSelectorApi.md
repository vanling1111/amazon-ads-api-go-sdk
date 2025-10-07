# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductMetadata**](ProductSelectorApi.md#ProductMetadata) | **Post** /product/metadata | Returns product metadata for the advertiser

# **ProductMetadata**
> ProductMetadataResponse ProductMetadata(ctx, body, amazonAdvertisingAPIClientId, optional)
Returns product metadata for the advertiser

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"campaign_edit\",\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductMetadataRequest**](ProductMetadataRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***ProductSelectorApiProductMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductSelectorApiProductMetadataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| Account identifier you use to access the DSP. This is your Amazon DSP Advertiser ID. | 

### Return type

[**ProductMetadataResponse**](ProductMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.productmetadatarequest.v1+json
 - **Accept**: application/vnd.productmetadataresponse.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

