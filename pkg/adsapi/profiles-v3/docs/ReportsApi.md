# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadReport**](ReportsApi.md#DownloadReport) | **Get** /v2/reports/{reportId}/download | Downloads a previously requested report identified by reportId.
[**GetReportStatus**](ReportsApi.md#GetReportStatus) | **Get** /v2/reports/{reportId} | Gets the status of a report previously requested.
[**RequestReport**](ReportsApi.md#RequestReport) | **Post** /sd/{recordType}/report | Creates a report request.

# **DownloadReport**
> DownloadReport(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, reportId)
Downloads a previously requested report identified by reportId.

Gets a `307 Temporary Redirect` response that includes a `location` header with the value set to an AWS S3 path where the report is located. The path expires after 30 seconds. If the path expires before the report is downloaded, a new report request must be created.  **To understand the call flow for asynchronous reports, see [Getting started with sponsored ads reports](/API/docs/en-us/guides/reporting/v2/sponsored-ads-reports).**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **reportId** | **string**| The identifier of the requested report. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportStatus**
> ReportResponse GetReportStatus(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, reportId)
Gets the status of a report previously requested.

Uses the `reportId` value from the response of a report previously requested via `POST` method of the `/sd/{recordType}/report` operation.  **To understand the call flow for asynchronous reports, see [Getting started with sponsored ads reports](/API/docs/en-us/guides/reporting/v2/sponsored-ads-reports).**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **reportId** | **string**| The identifier of the requested report. | 

### Return type

[**ReportResponse**](ReportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestReport**
> ReportResponse RequestReport(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, recordType, optional)
Creates a report request.

**To understand the call flow for asynchronous reports, see [Getting started with sponsored ads reports](/API/docs/en-us/guides/reporting/v2/sponsored-ads-reports).**  The Sponsored Display API supports creation of reports for campaigns, ad groups, product ads, targets, and asins. Create a ReportRequest object specifying the fields corresponding to performance data metrics to include in the report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **recordType** | **string**| The type of report to generate, either &#x60;campaigns&#x60;, &#x60;adGroups&#x60;, &#x60;productAds&#x60;, &#x60;targets&#x60;, or &#x60;asins&#x60;. The &#x27;asins&#x27; report, also known as the Purchased products report, is only available for seller brand owners. | 
 **optional** | ***ReportsApiRequestReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiRequestReportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of ReportRequest**](ReportRequest.md)|  | 

### Return type

[**ReportResponse**](ReportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

