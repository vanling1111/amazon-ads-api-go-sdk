# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadReport**](ReportsApi.md#DownloadReport) | **Get** /v2/reports/{reportId}/download | Downloads a previously requested report identified by &#x60;reportId&#x60;.
[**V2HsaRecordTypeReportPost**](ReportsApi.md#V2HsaRecordTypeReportPost) | **Post** /v2/hsa/{recordType}/report | Requests the creation of a report containing performance data related to Sponsored Brands campaigns.
[**V2ReportsReportIdGet**](ReportsApi.md#V2ReportsReportIdGet) | **Get** /v2/reports/{reportId} | Returns the status of a previously requested report.

# **DownloadReport**
> DownloadReport(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, reportId)
Downloads a previously requested report identified by `reportId`.

Gets a `307 Temporary Redirect` response that includes a `location` header with the value set to an AWS S3 path where the report is located. The path expires after 30 seconds. If the path expires before the report is downloaded, a new report request must be created.   **To understand the call flow for asynchronous reports, see [Getting started with sponsored ads reports](/API/docs/en-us/guides/reporting/v2/sponsored-ads-reports).**  The report file contains one row per entity for which performance data is present. These records are represented as JSON containing the ID attribute corresponding to the `recordType`, the segment (if specified), and each of the metrics in the request.  **Note**: The report files in S3 are gzipped.  *Example report download*  `$ curl -o /tmp/report.json.gz \"https://sandboxreports.s3.amazonaws.com/amzn1.clicksAPI.v1.m1.580149D6.c7aa92c1-ca5b-435d-bb8b-51cb26ad5731?AWSAccessKeyId=AKIAIKLHNT32USZOWVRA&amp;Expires=1476479900&amp;Signature=I%2F2Gd%2B8TbcPbXbBUM6ix%2BSVP3qA%3D\"`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **reportId** | **string**| The identifier of the requested report. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2HsaRecordTypeReportPost**
> InlineResponse20018 V2HsaRecordTypeReportPost(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, recordType)
Requests the creation of a report containing performance data related to Sponsored Brands campaigns.

Use this interface to request and retrieve performance reports. **To understand the call flow for asynchronous reports, see [Getting started with sponsored ads reports](/API/docs/en-us/guides/reporting/v2/sponsored-ads-reports).**  **KDP support**  Note that KDP profiles cannot currently request reports for Sponsored Brands campaigns using the API. Authors should use the advertising console to download reporting data.   **Filtering**  For more information, see the [Reporting FAQ](/API/docs/en-us/guides/reporting/v2/faq#can-i-filter-a-sponsored-brands-report).  **Constraints of Sponsored Brands reporting**  Sponsored Brands reporting data cannot be combined with Sponsored Products data into one report. Only 14-day data is available for Sponsored Brands. Attribution windows of 1, 7, and 30-day intervals are not available.  **New-to-brand performance metrics**  With new-to-brand metrics, advertisers can measure and optimize campaigns, as well as plan future marketing strategies to grow their customer base on Amazon. New-to-brand metrics determine whether an ad-attributed purchase was made by an existing customer or one buying a brand’s product on Amazon for the first time over the prior year. With new-to-brand metrics, advertisers receive campaign performance metrics such as total new-to-brand purchases and sales, new-to-brand purchase rate, and cost per new-to-brand customer. Advertisers now have the tools they need to estimate the cost of acquiring new customers on Amazon and identify the most efficient channels and tactics to achieve their campaign goals.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RecordTypeReportBody**](RecordTypeReportBody.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **recordType** | **string**| The type of report. Valid types are &#x60;campaigns&#x60;, &#x60;adGroups&#x60;, &#x60;ads&#x60;, &#x60;targets&#x60;, and &#x60;keywords&#x60;. | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2ReportsReportIdGet**
> InlineResponse20019 V2ReportsReportIdGet(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, reportId)
Returns the status of a previously requested report.

To understand the call flow for asynchronous reports, see [Getting started with sponsored ads reports](/API/docs/en-us/guides/reporting/v2/sponsored-ads-reports).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **reportId** | **string**| Report ID returned by POST report method. | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

