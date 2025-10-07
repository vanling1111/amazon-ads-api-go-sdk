# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountInformation**](GetTestAccountInformation_Api.md#GetAccountInformation) | **Get** /testAccounts | 

# **GetAccountInformation**
> []GetAccountInformationResponseInner GetAccountInformation(ctx, amazonAdvertisingAPIClientId, optional)


API to get Account information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***GetTestAccountInformation_ApiGetAccountInformationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTestAccountInformation_ApiGetAccountInformationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestId** | **optional.String**| request id. | 

### Return type

[**[]GetAccountInformationResponseInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

