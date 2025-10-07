# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdvertiserInvoices**](InvoiceApi.md#GetAdvertiserInvoices) | **Get** /invoices | Get invoices for advertiser
[**GetInvoice**](InvoiceApi.md#GetInvoice) | **Get** /invoices/{invoiceId} | Get invoice data by invoice ID

# **GetAdvertiserInvoices**
> InlineResponse200 GetAdvertiserInvoices(ctx, optional)
Get invoices for advertiser

  **Requires one of these permissions**: [\"nemo_transactions_view\",\"nemo_transactions_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InvoiceApiGetAdvertiserInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoiceApiGetAdvertiserInvoicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceStatuses** | [**optional.Interface of []string**](string.md)| * &#x60;ISSUED&#x60;: An invoice is issued when its charges are finalized and tax is computed on the total amount.  * &#x60;PAID_IN_PART&#x60;: When a partial payment is received, the invoice status will change to paid in part. * &#x60;PAID_IN_PART&#x60;: One full payment has been received, the invoice will be paid in full. * &#x60;WRITTEN_OFF&#x60;: If an invoice is written off because of an error, the status will be updated to written off.  | 
 **startDate** | [**optional.Interface of Object**](.md)| The starting date (inclusive) of the date range for filtering invoices. Please provide the date in ISO-8601 format, representing a UTC date with only the date portion (no time). | 
 **endDate** | [**optional.Interface of Object**](.md)| The ending date (inclusive) of the date range for filtering invoices. Please provide the date in ISO-8601 format, representing a UTC date with only the date portion (no time). | 
 **count** | **optional.Int32**| Number of records to include in the paged response. Defaults to 100. Cannot be combined with the cursor parameter. | 
 **cursor** | **optional.String**| A cursor representing how far into a result set this query should begin. In the absence of a cursor the request will default to start index of 0 and page size of 100. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.invoices.v1.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoice**
> Invoice GetInvoice(ctx, invoiceId)
Get invoice data by invoice ID

  **Requires one of these permissions**: [\"nemo_transactions_view\",\"nemo_transactions_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **string**| ID of invoice to fetch | 

### Return type

[**Invoice**](invoice.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.invoice.v1.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

