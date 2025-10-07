# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTermsToken**](TermsTokenApi.md#CreateTermsToken) | **Post** /termsTokens | Create a new UUID terms token for the customer to accept advertising terms
[**GetTermsToken**](TermsTokenApi.md#GetTermsToken) | **Get** /termsTokens/{termsToken} | Get the terms token status for the customer

# **CreateTermsToken**
> CreateTermsTokenResponseContent CreateTermsToken(ctx, optional)
Create a new UUID terms token for the customer to accept advertising terms

Create a new UUID terms token for the customer to accept advertising terms  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TermsTokenApiCreateTermsTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermsTokenApiCreateTermsTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateTermsTokenRequestContent**](CreateTermsTokenRequestContent.md)|  | 
 **amazonAdvertisingAPIClientId** | **optional.**| An identifier of API client making the request on behalf of the customer. | 

### Return type

[**CreateTermsTokenResponseContent**](CreateTermsTokenResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.GlobalRegistrationService.TermsTokenResource.v1.0+json
 - **Accept**: application/vnd.termstokenresource.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTermsToken**
> GetTermsTokenResponseContent GetTermsToken(ctx, termsToken, amazonAdvertisingAPIClientId)
Get the terms token status for the customer

Get the terms token status for the customer  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termsToken** | **string**| A Terms Token refers to an UUID token used for terms and conditions acceptance  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a Login with Amazon account. | 

### Return type

[**GetTermsTokenResponseContent**](GetTermsTokenResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.termstokenresource.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

