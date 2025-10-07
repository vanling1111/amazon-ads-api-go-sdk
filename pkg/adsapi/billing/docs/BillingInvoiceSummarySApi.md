# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBillingInvoiceSummaries**](BillingInvoiceSummarySApi.md#GetBillingInvoiceSummaries) | **Post** /invoiceSummaries/list | Lists the billing invoice summary(s) in a global ads account as per the search and aggregation parameters passed in the request

# **GetBillingInvoiceSummaries**
> BillingInvoiceSummariesResponse GetBillingInvoiceSummaries(ctx, body, amazonAdvertisingAPIClientId, optional)
Lists the billing invoice summary(s) in a global ads account as per the search and aggregation parameters passed in the request

Lists the billing invoice summary(s) in a global ads account. You can further narrow down the search by providing filter(s) over country, status, paymentMethod with 'exact' and 'broad' match, search over invoice number & ro number and aggregation query(s) over the billing invoice summary(s)  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"nemo_transactions_edit\",\"nemo_transactions_view\",\"ManagerAccount_Dev\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\"]  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"nemo_transactions_edit\",\"nemo_transactions_view\",\"ManagerAccount_Dev\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\"]  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BillingInvoiceSummariesRequest**](BillingInvoiceSummariesRequest.md)| Payload with filter, sort and aggregate key(s) to fetch list of billing invoice summary(s). | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***BillingInvoiceSummarySApiGetBillingInvoiceSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingInvoiceSummarySApiGetBillingInvoiceSummariesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| The identifier of an account. Account can be a global advertising account. | 
 **amazonAdvertisingAPIManagerAccount** | **optional.**| The identifier of a manager account. | 

### Return type

[**BillingInvoiceSummariesResponse**](BillingInvoiceSummariesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.billingInvoiceSummary.v1+json
 - **Accept**: application/vnd.billingInvoiceSummary.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

