# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DspAmazonCreateConversionDefinitions**](AmazonConversionDefinitionsApi.md#DspAmazonCreateConversionDefinitions) | **Post** /accounts/{accountId}/dsp/conversionDefinitions | Batch create conversion definitions.
[**DspAmazonListConversionDefinitions**](AmazonConversionDefinitionsApi.md#DspAmazonListConversionDefinitions) | **Post** /accounts/{accountId}/dsp/conversionDefinitions/list | Retrieve a list of conversion definitions based on filters.
[**DspAmazonUpdateConversionDefinitions**](AmazonConversionDefinitionsApi.md#DspAmazonUpdateConversionDefinitions) | **Put** /accounts/{accountId}/dsp/conversionDefinitions | Batch update conversion definitions.

# **DspAmazonCreateConversionDefinitions**
> ConversionDefinitionsBatchResponseV1 DspAmazonCreateConversionDefinitions(ctx, amazonAdvertisingAPIClientId, accountId, optional)
Batch create conversion definitions.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
 **optional** | ***AmazonConversionDefinitionsApiDspAmazonCreateConversionDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AmazonConversionDefinitionsApiDspAmazonCreateConversionDefinitionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []ConversionDefinitionInputV1**](ConversionDefinitionInputV1.md)|  | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**ConversionDefinitionsBatchResponseV1**](ConversionDefinitionsBatchResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspconversiondefinition.v1+json, application/vnd.dspconversiondefinition.v2+json
 - **Accept**: application/vnd.dspconversiondefinition.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DspAmazonListConversionDefinitions**
> ListConversionDefinitionsResponseV1 DspAmazonListConversionDefinitions(ctx, body, amazonAdvertisingAPIClientId, accountId, optional)
Retrieve a list of conversion definitions based on filters.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ListConversionDefinitionsRequestV1**](ListConversionDefinitionsRequestV1.md)| Filters to apply when retrieving conversion definitions. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
 **optional** | ***AmazonConversionDefinitionsApiDspAmazonListConversionDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AmazonConversionDefinitionsApiDspAmazonListConversionDefinitionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**ListConversionDefinitionsResponseV1**](ListConversionDefinitionsResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspconversiondefinition.v1+json
 - **Accept**: application/vnd.dspconversiondefinition.v1+json, application/vnd.dspconversiondefinition.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DspAmazonUpdateConversionDefinitions**
> ConversionDefinitionsBatchResponseV1 DspAmazonUpdateConversionDefinitions(ctx, amazonAdvertisingAPIClientId, accountId, optional)
Batch update conversion definitions.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
 **optional** | ***AmazonConversionDefinitionsApiDspAmazonUpdateConversionDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AmazonConversionDefinitionsApiDspAmazonUpdateConversionDefinitionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []UpdateConversionDefinitionV1**](UpdateConversionDefinitionV1.md)|  | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**ConversionDefinitionsBatchResponseV1**](ConversionDefinitionsBatchResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspconversiondefinition.v1+json, application/vnd.dspconversiondefinition.v2+json
 - **Accept**: application/vnd.dspconversiondefinition.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

