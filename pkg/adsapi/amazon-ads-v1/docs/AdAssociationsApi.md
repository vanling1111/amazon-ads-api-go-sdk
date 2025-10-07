# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdAssociation**](AdAssociationsApi.md#CreateAdAssociation) | **Post** /adsApi/v1/create/adAssociations | 
[**DeleteAdAssociation**](AdAssociationsApi.md#DeleteAdAssociation) | **Post** /adsApi/v1/delete/adAssociations | 
[**QueryAdAssociation**](AdAssociationsApi.md#QueryAdAssociation) | **Post** /adsApi/v1/query/adAssociations | 
[**UpdateAdAssociation**](AdAssociationsApi.md#UpdateAdAssociation) | **Post** /adsApi/v1/update/adAssociations | 

# **CreateAdAssociation**
> AdAssociationMultiStatusResponse CreateAdAssociation(ctx, amazonAdsAccountId, amazonAdsClientId, optional)


Creates AdAssociations.  **Requires one of these permissions**: [\"campaign_edit\"][\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsAccountId** | **string**| The identifier of an Amazon Ads Advertiser Account. | 
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***AdAssociationsApiCreateAdAssociationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdAssociationsApiCreateAdAssociationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of CreateAdAssociationRequest**](CreateAdAssociationRequest.md)|  | 

### Return type

[**AdAssociationMultiStatusResponse**](AdAssociationMultiStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAdAssociation**
> AdAssociationMultiStatusResponse DeleteAdAssociation(ctx, amazonAdsAccountId, amazonAdsClientId, optional)


Archives or deletes AdAssociations.  **Requires one of these permissions**: [\"campaign_edit\"][\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsAccountId** | **string**| The identifier of an Amazon Ads Advertiser Account. | 
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***AdAssociationsApiDeleteAdAssociationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdAssociationsApiDeleteAdAssociationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeleteAdAssociationRequest**](DeleteAdAssociationRequest.md)|  | 

### Return type

[**AdAssociationMultiStatusResponse**](AdAssociationMultiStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryAdAssociation**
> AdAssociationSuccessResponse QueryAdAssociation(ctx, amazonAdsAccountId, amazonAdsClientId, optional)


A search read, allowing use of more complex filters.  **Requires one of these permissions**: [\"campaign_view\",\"creatives_view\"][\"campaign_view\",\"creatives_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsAccountId** | **string**| The identifier of an Amazon Ads Advertiser Account. | 
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***AdAssociationsApiQueryAdAssociationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdAssociationsApiQueryAdAssociationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of QueryAdAssociationRequest**](QueryAdAssociationRequest.md)|  | 

### Return type

[**AdAssociationSuccessResponse**](AdAssociationSuccessResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAdAssociation**
> AdAssociationMultiStatusResponse UpdateAdAssociation(ctx, amazonAdsAccountId, amazonAdsClientId, optional)


Updates AdAssociations. Behaves as PATCH unless otherwise stated.  **Requires one of these permissions**: [\"campaign_edit\"][\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdsAccountId** | **string**| The identifier of an Amazon Ads Advertiser Account. | 
  **amazonAdsClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
 **optional** | ***AdAssociationsApiUpdateAdAssociationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdAssociationsApiUpdateAdAssociationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateAdAssociationRequest**](UpdateAdAssociationRequest.md)|  | 

### Return type

[**AdAssociationMultiStatusResponse**](AdAssociationMultiStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

