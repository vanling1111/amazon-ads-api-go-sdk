# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreModeration**](PreModerationApi.md#PreModeration) | **Post** /preModeration | Pre moderate the components

# **PreModeration**
> PreModerationResponse PreModeration(ctx, body, amazonAdvertisingAPIClientId, optional)
Pre moderate the components

This API will be accepting different components of the ad/page and will be automatically validating the components and send back the policy violations if any. We recommend to send all components of the same entity to be sent together. It will make us better detect any policy violation if present. This will increase the Time to go live for the entity. In one request please don't send the components of more than one entity.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"campaign_edit\",\"campaign_view\",\"amazon_stores_edit\",\"amazon_stores_view\",\"brand_posts_admin\",\"brand_posts_edit\",\"brand_posts_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PreModerationRequest**](PreModerationRequest.md)| Request body for preModeration API. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***PreModerationApiPreModerationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreModerationApiPreModerationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| The header used to pass global account associated with the advertiser account. Use &#x60;GET&#x60; method on the Global Ads Account resource to list the global ads account associated with the access token passed in the HTTP Authorization header and choose AdvertisingAccountIdentifier id from the response to pass it as input. Use for v2 global calls. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**PreModerationResponse**](PreModerationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

