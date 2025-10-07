# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DspAmazonBatchGetConversionDefinitionsForOrders**](ConversionOrderAssociationApi.md#DspAmazonBatchGetConversionDefinitionsForOrders) | **Post** /accounts/{accountId}/dsp/batchOrders/conversionDefinitionAssociations | Retrieve associated conversion definitions for orders.
[**DspAmazonGetAssociatedConversionDefinitionsForOrder**](ConversionOrderAssociationApi.md#DspAmazonGetAssociatedConversionDefinitionsForOrder) | **Get** /accounts/{accountId}/dsp/orders/{orderId}/conversionDefinitionAssociations | Retrieve associated conversion definitions for an order.
[**DspAmazonUpdateAssociatedConversionDefinitionsForOrder**](ConversionOrderAssociationApi.md#DspAmazonUpdateAssociatedConversionDefinitionsForOrder) | **Post** /accounts/{accountId}/dsp/orders/{orderId}/conversionDefinitionAssociations | Associate/Dissociate conversion definitions to an order.

# **DspAmazonBatchGetConversionDefinitionsForOrders**
> BatchGetConversionDefinitionsForOrdersResponseV1 DspAmazonBatchGetConversionDefinitionsForOrders(ctx, body, amazonAdvertisingAPIClientId, accountId, optional)
Retrieve associated conversion definitions for orders.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchGetConversionDefinitionsAssociatedForOrdersRequestV1**](BatchGetConversionDefinitionsAssociatedForOrdersRequestV1.md)| A list of orderIds to retrieve the associated conversion definitions. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
 **optional** | ***ConversionOrderAssociationApiDspAmazonBatchGetConversionDefinitionsForOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversionOrderAssociationApiDspAmazonBatchGetConversionDefinitionsForOrdersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**BatchGetConversionDefinitionsForOrdersResponseV1**](BatchGetConversionDefinitionsForOrdersResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspbatchgetconversiondefinitions.v1+json
 - **Accept**: application/vnd.dspbatchgetconversiondefinitions.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DspAmazonGetAssociatedConversionDefinitionsForOrder**
> OrderAssociatedConversionDefinitionsResponseV1 DspAmazonGetAssociatedConversionDefinitionsForOrder(ctx, amazonAdvertisingAPIClientId, accountId, orderId, optional)
Retrieve associated conversion definitions for an order.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
  **orderId** | **string**| The identifier of the order. | 
 **optional** | ***ConversionOrderAssociationApiDspAmazonGetAssociatedConversionDefinitionsForOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversionOrderAssociationApiDspAmazonGetAssociatedConversionDefinitionsForOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.String**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 
 **nextToken** | **optional.String**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.Int32**| Sets the maximum number of associated conversions in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | 

### Return type

[**OrderAssociatedConversionDefinitionsResponseV1**](OrderAssociatedConversionDefinitionsResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.dsporderconversionassociation.v1+json, application/vnd.dsporderconversionassociation.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DspAmazonUpdateAssociatedConversionDefinitionsForOrder**
> BatchAssociateConversionDefinitionsResponseV1 DspAmazonUpdateAssociatedConversionDefinitionsForOrder(ctx, amazonAdvertisingAPIClientId, accountId, orderId, optional)
Associate/Dissociate conversion definitions to an order.

  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **accountId** | **string**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. DSP Entity ID is not supported. | 
  **orderId** | **string**| The identifier of the order. | 
 **optional** | ***ConversionOrderAssociationApiDspAmazonUpdateAssociatedConversionDefinitionsForOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversionOrderAssociationApiDspAmazonUpdateAssociatedConversionDefinitionsForOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of []BatchAssociateConversionDefinitionsRequestV1Inner**](BatchAssociateConversionDefinitionsRequestV1_inner.md)|  | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id profileId from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| Account Identifier you use to access the DSP. This is your DSP Advertiser ID. This header is required for access to any DSP data when you have been directly invited to the advertiser. | 

### Return type

[**BatchAssociateConversionDefinitionsResponseV1**](BatchAssociateConversionDefinitionsResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dsporderconversionassociation.v1+json, application/vnd.dsporderconversionassociation.v2+json
 - **Accept**: application/vnd.dsporderconversionassociation.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

