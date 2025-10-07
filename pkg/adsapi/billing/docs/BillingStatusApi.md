# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkGetBillingStatus**](BillingStatusApi.md#BulkGetBillingStatus) | **Post** /billing/statuses | Get the billing status for a list of advertising accounts.

# **BulkGetBillingStatus**
> BulkGetBillingStatusResponse BulkGetBillingStatus(ctx, body)
Get the billing status for a list of advertising accounts.

Gets the current billing status associated for each advertising account.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"adv_billing_view\",\"adv_billing_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkGetBillingStatusesRequestBody**](BulkGetBillingStatusesRequestBody.md)|  | 

### Return type

[**BulkGetBillingStatusResponse**](bulkGetBillingStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.bulkgetbillingstatusrequestbody.v1+json
 - **Accept**: application/vnd.bulkgetbillingstatusresponse.v1+json, application/vnd.bulkgetbillingstatuserrorresponse.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

