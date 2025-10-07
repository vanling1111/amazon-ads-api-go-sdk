# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DspAmazonBatchCreateMobileMeasurementPartnerAppRegistration**](AppsApi.md#DspAmazonBatchCreateMobileMeasurementPartnerAppRegistration) | **Post** /accounts/{accountId}/dsp/mobileMeasurementPartners | Create a new Mobile Measurement Partner app registration.
[**DspAmazonBatchUpdateMobileMeasurementPartnerAppRegistration**](AppsApi.md#DspAmazonBatchUpdateMobileMeasurementPartnerAppRegistration) | **Put** /accounts/{accountId}/dsp/mobileMeasurementPartners | Update a Mobile Measurement Partner app registration. Updates may sever the data connection between the Mobile Measurement Partner and Amazon.
[**DspAmazonDeleteMeasurementPartnerAppRegistrations**](AppsApi.md#DspAmazonDeleteMeasurementPartnerAppRegistrations) | **Post** /accounts/{accountId}/dsp/mobileMeasurementPartners/delete | Marks a Mobile Measurement Partner app registration as deleted. Deleted Mobile Measurement Partner app registrations will sever the data connection between the Mobile Measurement Partner and Amazon.
[**DspAmazonGetAssociatedMobileAppForConversionDefinition**](AppsApi.md#DspAmazonGetAssociatedMobileAppForConversionDefinition) | **Get** /accounts/{accountId}/dsp/conversionDefinitions/{conversionDefinitionId}/mobileMeasurementPartnerAppRegistration | Retrieve associated Mobile Measurement Partner App for a ConversionDefinition.
[**DspAmazonListMobileMeasurementPartnerAppRegistrations**](AppsApi.md#DspAmazonListMobileMeasurementPartnerAppRegistrations) | **Post** /accounts/{accountId}/dsp/mobileMeasurementPartners/list | List Mobile Measurement Partner App Registrations

# **DspAmazonBatchCreateMobileMeasurementPartnerAppRegistration**
> MobileMeasurementPartnerAppBatchResponseV1 DspAmazonBatchCreateMobileMeasurementPartnerAppRegistration(ctx, body, amazonAdvertisingAPIClientId, accountId, optional)
Create a new Mobile Measurement Partner app registration.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MobileMeasurementPartnerAppRegistrationV1**](MobileMeasurementPartnerAppRegistrationV1.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
 **optional** | ***AppsApiDspAmazonBatchCreateMobileMeasurementPartnerAppRegistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppsApiDspAmazonBatchCreateMobileMeasurementPartnerAppRegistrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**MobileMeasurementPartnerAppBatchResponseV1**](MobileMeasurementPartnerAppBatchResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspmobilemeasurementpartnerappregistration.v1+json
 - **Accept**: application/vnd.dspmobilemeasurementpartnerappregistration.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DspAmazonBatchUpdateMobileMeasurementPartnerAppRegistration**
> MobileMeasurementPartnerAppBatchResponseV1 DspAmazonBatchUpdateMobileMeasurementPartnerAppRegistration(ctx, body, amazonAdvertisingAPIClientId, accountId, optional)
Update a Mobile Measurement Partner app registration. Updates may sever the data connection between the Mobile Measurement Partner and Amazon.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]UpdateMobileMeasurementPartnerAppRegistrationV1**](UpdateMobileMeasurementPartnerAppRegistrationV1.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
 **optional** | ***AppsApiDspAmazonBatchUpdateMobileMeasurementPartnerAppRegistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppsApiDspAmazonBatchUpdateMobileMeasurementPartnerAppRegistrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**MobileMeasurementPartnerAppBatchResponseV1**](MobileMeasurementPartnerAppBatchResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspmobilemeasurementpartnerappregistration.v1+json
 - **Accept**: application/vnd.dspmobilemeasurementpartnerappregistration.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DspAmazonDeleteMeasurementPartnerAppRegistrations**
> MobileMeasurementPartnerAppBatchResponseV1 DspAmazonDeleteMeasurementPartnerAppRegistrations(ctx, body, amazonAdvertisingAPIClientId, accountId, optional)
Marks a Mobile Measurement Partner app registration as deleted. Deleted Mobile Measurement Partner app registrations will sever the data connection between the Mobile Measurement Partner and Amazon.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]DeleteMobileMeasurementPartnerAppRegistrationV1**](DeleteMobileMeasurementPartnerAppRegistrationV1.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
 **optional** | ***AppsApiDspAmazonDeleteMeasurementPartnerAppRegistrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppsApiDspAmazonDeleteMeasurementPartnerAppRegistrationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**MobileMeasurementPartnerAppBatchResponseV1**](MobileMeasurementPartnerAppBatchResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspmobilemeasurementpartnerappregistration.v1+json
 - **Accept**: application/vnd.dspmobilemeasurementpartnerappregistration.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DspAmazonGetAssociatedMobileAppForConversionDefinition**
> AssociatedMobileMeasurementPartnerAppRegistrationV1 DspAmazonGetAssociatedMobileAppForConversionDefinition(ctx, amazonAdvertisingAPIClientId, accountId, conversionDefinitionId, optional)
Retrieve associated Mobile Measurement Partner App for a ConversionDefinition.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
  **conversionDefinitionId** | **string**| The identifier of the ConversionDefinition. | 
 **optional** | ***AppsApiDspAmazonGetAssociatedMobileAppForConversionDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppsApiDspAmazonGetAssociatedMobileAppForConversionDefinitionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.String**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**AssociatedMobileMeasurementPartnerAppRegistrationV1**](AssociatedMobileMeasurementPartnerAppRegistrationV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.dspassociatedmobilemeasurementpartnerappregistration.v1+json, application/vnd.dspconversionadtageventassociation.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DspAmazonListMobileMeasurementPartnerAppRegistrations**
> ListMobileMeasurementPartnerAppRegistrationsResponseV1 DspAmazonListMobileMeasurementPartnerAppRegistrations(ctx, body, amazonAdvertisingAPIClientId, accountId, optional)
List Mobile Measurement Partner App Registrations

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"event_manager_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ListMobileMeasurementPartnerAppRegistrationsRequestV1**](ListMobileMeasurementPartnerAppRegistrationsRequestV1.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
 **optional** | ***AppsApiDspAmazonListMobileMeasurementPartnerAppRegistrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppsApiDspAmazonListMobileMeasurementPartnerAppRegistrationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**ListMobileMeasurementPartnerAppRegistrationsResponseV1**](ListMobileMeasurementPartnerAppRegistrationsResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspmobilemeasurementpartnerappregistration.v1+json
 - **Accept**: application/vnd.dspmobilemeasurementpartnerappregistration.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

