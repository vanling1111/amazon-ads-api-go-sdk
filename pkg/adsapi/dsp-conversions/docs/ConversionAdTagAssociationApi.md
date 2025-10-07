# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DspAmazonGetAdTagAssociatedEvent**](ConversionAdTagAssociationApi.md#DspAmazonGetAdTagAssociatedEvent) | **Get** /accounts/{accountId}/dsp/conversionDefinitions/{conversionDefinitionId}/adTagEventAssociations | Retrieve associated Amazon adTag event for a ConversionDefinition.
[**DspAmazonUpdateAdTagAssociatedEvent**](ConversionAdTagAssociationApi.md#DspAmazonUpdateAdTagAssociatedEvent) | **Post** /accounts/{accountId}/dsp/conversionDefinitions/{conversionDefinitionId}/adTagEventAssociations | Associate/Dissociate an Amazon adTag event to a ConversionDefinition.

# **DspAmazonGetAdTagAssociatedEvent**
> ConversionDefinitionAssociatedAdTagEventV1 DspAmazonGetAdTagAssociatedEvent(ctx, amazonAdvertisingAPIClientId, accountId, conversionDefinitionId, optional)
Retrieve associated Amazon adTag event for a ConversionDefinition.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
  **conversionDefinitionId** | **string**| The identifier of the ConversionDefinition. | 
 **optional** | ***ConversionAdTagAssociationApiDspAmazonGetAdTagAssociatedEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversionAdTagAssociationApiDspAmazonGetAdTagAssociatedEventOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.String**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**ConversionDefinitionAssociatedAdTagEventV1**](ConversionDefinitionAssociatedAdTagEventV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.dspconversionadtageventassociation.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DspAmazonUpdateAdTagAssociatedEvent**
> ConversionDefinitionAssociatedAdTagEventV1 DspAmazonUpdateAdTagAssociatedEvent(ctx, amazonAdvertisingAPIClientId, accountId, conversionDefinitionId, optional)
Associate/Dissociate an Amazon adTag event to a ConversionDefinition.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
  **conversionDefinitionId** | **string**| The identifier of the ConversionDefinition. | 
 **optional** | ***ConversionAdTagAssociationApiDspAmazonUpdateAdTagAssociatedEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversionAdTagAssociationApiDspAmazonUpdateAdTagAssociatedEventOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of ConversionDefinitionAdTagEventAssociationRequestV1**](ConversionDefinitionAdTagEventAssociationRequestV1.md)|  | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**ConversionDefinitionAssociatedAdTagEventV1**](ConversionDefinitionAssociatedAdTagEventV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspconversionadtageventassociation.v1+json
 - **Accept**: application/vnd.dspconversionadtageventassociation.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

