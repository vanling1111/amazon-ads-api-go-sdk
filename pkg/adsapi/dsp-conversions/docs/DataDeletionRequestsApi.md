# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DspAmazonDeletionRequest**](DataDeletionRequestsApi.md#DspAmazonDeletionRequest) | **Post** /accounts/{accountId}/dsp/conversionDefinitions/delete | Deletes existing event data associated with user(s). Supply all match keys associated with the user. Events processed before the deletion request will be deleted from our advertising systems. However, any events for these users sent after the deletion request will not be deleted. Therefore you must stop sending events for opted-out users after the deletion request is made in order to effectuate the opt-out.

# **DspAmazonDeletionRequest**
> BatchDeleteUserEventsResponseV1 DspAmazonDeletionRequest(ctx, body, amazonAdvertisingAPIClientId, accountId, optional)
Deletes existing event data associated with user(s). Supply all match keys associated with the user. Events processed before the deletion request will be deleted from our advertising systems. However, any events for these users sent after the deletion request will not be deleted. Therefore you must stop sending events for opted-out users after the deletion request is made in order to effectuate the opt-out.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchDeleteUserEventsRequestV1**](BatchDeleteUserEventsRequestV1.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
 **optional** | ***DataDeletionRequestsApiDspAmazonDeletionRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataDeletionRequestsApiDspAmazonDeletionRequestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**BatchDeleteUserEventsResponseV1**](BatchDeleteUserEventsResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspuserdeletionrequest.v1+json
 - **Accept**: application/vnd.dspuserdeletionrequest.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

