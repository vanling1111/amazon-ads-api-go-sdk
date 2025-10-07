# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RASv1ListRetailers**](RetailersApi.md#RASv1ListRetailers) | **Post** /ras/v1/retailers/list | Returns a list of Retailers.

# **RASv1ListRetailers**
> IdentityListRetailersResponseContent RASv1ListRetailers(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Returns a list of Retailers.

Returns an array of all active retailers.  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**|  | 
  **amazonAdvertisingAPIScope** | **string**|  | 
 **optional** | ***RetailersApiRASv1ListRetailersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetailersApiRASv1ListRetailersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ListRetailersRequestContent**](ListRetailersRequestContent.md)|  | 

### Return type

[**IdentityListRetailersResponseContent**](IdentityListRetailersResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

