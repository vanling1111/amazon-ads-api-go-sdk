# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PayInvoices**](PayInvoicesApi.md#PayInvoices) | **Post** /billing/invoices/pay | 

# **PayInvoices**
> AdPaymentsPayInvoicesOutput PayInvoices(ctx, body, amazonAdvertisingAPIClientId, optional)


Executes payment on a set of or all of an advertisers open invoices.  **Requires one of these permissions**: [\"adv_billing_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdPaymentsPayInvoicesInput**](AdPaymentsPayInvoicesInput.md)| This API executes payment on one or multiple open invoices. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***PayInvoicesApiPayInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PayInvoicesApiPayInvoicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**AdPaymentsPayInvoicesOutput**](AdPaymentsPayInvoicesOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.invoices.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

