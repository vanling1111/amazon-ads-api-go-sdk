# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DspAmazonAdTagGetAdTagByAdvertiserId**](AmazonAdTagApi.md#DspAmazonAdTagGetAdTagByAdvertiserId) | **Get** /accounts/{accountId}/dsp/amazonAdTag | Gets an Amazon Ad Tag for a given advertiser
[**DspAmazonAdTagGetEventsByAdTagId**](AmazonAdTagApi.md#DspAmazonAdTagGetEventsByAdTagId) | **Get** /accounts/{accountId}/dsp/adTagEvents/{adTagId}/list | Gets a list of available event metadata for the given ad tag.

# **DspAmazonAdTagGetAdTagByAdvertiserId**
> GetAdTagResponseV1 DspAmazonAdTagGetAdTagByAdvertiserId(ctx, amazonAdvertisingAPIClientId, accountId, optional)
Gets an Amazon Ad Tag for a given advertiser

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_view\"]  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Ads-Manager-Account-ID  **Parameter in**: header  **Requires one of these permissions**: [\"ads_data_manager_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
 **optional** | ***AmazonAdTagApiDspAmazonAdTagGetAdTagByAdvertiserIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AmazonAdTagApiDspAmazonAdTagGetAdTagByAdvertiserIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.String**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 
 **amazonAdsManagerAccountID** | **optional.String**| The identifier of an Amazon Ads Manager Account. This ID can be retrieved by making a GET request to the /managerAccounts endpoint. | 

### Return type

[**GetAdTagResponseV1**](GetAdTagResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.dspamazonadtags.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DspAmazonAdTagGetEventsByAdTagId**
> ListAdTagEventsResponseV1 DspAmazonAdTagGetEventsByAdTagId(ctx, amazonAdvertisingAPIClientId, accountId, adTagId, optional)
Gets a list of available event metadata for the given ad tag.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_view\"]  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Ads-Manager-Account-ID  **Parameter in**: header  **Requires one of these permissions**: [\"ads_data_manager_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
  **adTagId** | **string**| The identifier of the ad tag. | 
 **optional** | ***AmazonAdTagApiDspAmazonAdTagGetEventsByAdTagIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AmazonAdTagApiDspAmazonAdTagGetEventsByAdTagIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.String**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 
 **amazonAdsManagerAccountID** | **optional.String**| The identifier of an Amazon Ads Manager Account. This ID can be retrieved by making a GET request to the /managerAccounts endpoint. | 
 **searchTerm** | **optional.String**| A string used to fuzzy search for event metadata by event name. | 
 **startDateTime** | **optional.Time**| A date string to create a time window for events that have seen activity. If not set, a start date 30 days before the request was sent will be chosen. The oldest start date is 1 year before the request was sent to this api. | 
 **nextToken** | **optional.String**| Token from a previous request. Use to control pagination of the returned array. | 
 **maxResults** | **optional.Int32**| Sets the maximum number of associated conversions in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | 
 **endDateTime** | **optional.Time**| A date string to create a time window for events that have seen activity. If not set, the day and time the request was sent will be chosen as the end date. The maximum end date is the day the request to this api was sent. | 

### Return type

[**ListAdTagEventsResponseV1**](ListAdTagEventsResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.dspamazonadtagevents.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

