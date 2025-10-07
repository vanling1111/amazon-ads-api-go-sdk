# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DspAmazonIngestConversionData**](ConversionEventDataApi.md#DspAmazonIngestConversionData) | **Post** /accounts/{accountId}/dsp/conversionDefinitions/eventData | Import conversion event data. This API expects one source per request across all conversion event data and supports partial import. Only conversion event data with matching conversion source will be imported, the rest will be rejected.

# **DspAmazonIngestConversionData**
> BatchImportConversionEventDataResponseV1 DspAmazonIngestConversionData(ctx, body, amazonAdvertisingAPIClientId, accountId, optional)
Import conversion event data. This API expects one source per request across all conversion event data and supports partial import. Only conversion event data with matching conversion source will be imported, the rest will be rejected.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchImportConversionEventDataRequestV1**](BatchImportConversionEventDataRequestV1.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
 **optional** | ***ConversionEventDataApiDspAmazonIngestConversionDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversionEventDataApiDspAmazonIngestConversionDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**BatchImportConversionEventDataResponseV1**](BatchImportConversionEventDataResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspconversioneventimport.v1+json
 - **Accept**: application/vnd.dspconversioneventimport.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

