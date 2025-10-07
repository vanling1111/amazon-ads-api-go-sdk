# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAsyncReport**](AsynchronousReportsApi.md#CreateAsyncReport) | **Post** /reporting/reports | Creates a report request
[**DeleteAsyncReport**](AsynchronousReportsApi.md#DeleteAsyncReport) | **Delete** /reporting/reports/{reportId} | Deletes a report by id
[**GetAsyncReport**](AsynchronousReportsApi.md#GetAsyncReport) | **Get** /reporting/reports/{reportId} | Gets a generation status of report by id

# **CreateAsyncReport**
> AsyncReport CreateAsyncReport(ctx, amazonAdvertisingAPIClientId, optional)
Creates a report request

Creates a report request. Use this operation to request the creation of a new report for Amazon Advertising Products. Use `adProduct` to specify the Advertising Product of the report.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Scope  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"nemo_report_edit\",\"nemo_report_view\",\"reports_edit\"]  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"nemo_report_edit\",\"nemo_report_view\",\"reports_edit\",\"view_performance_dashboard\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The client identifier of the customer making the request. | 
 **optional** | ***AsynchronousReportsApiCreateAsyncReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AsynchronousReportsApiCreateAsyncReportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateAsyncReportRequest**](CreateAsyncReportRequest.md)|  | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **amazonAdsAccountId** | **optional.**| The identifier of a DSP advertiser account. Optional. For details, see [this guide](guides/reporting/dsp/creating-reports). | 

### Return type

[**AsyncReport**](AsyncReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.createasyncreportrequest.v3+json
 - **Accept**: application/vnd.createasyncreportresponse.v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAsyncReport**
> DeleteAsyncReportResponse DeleteAsyncReport(ctx, reportId, amazonAdvertisingAPIClientId, optional)
Deletes a report by id

Deletes a report by id. Use this operation to cancel a report in a `PENDING` status.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Scope  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"nemo_report_edit\",\"nemo_report_view\",\"reports_edit\"]  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"nemo_report_edit\",\"nemo_report_view\",\"reports_edit\",\"view_performance_dashboard\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **string**| The identifier of the requested report. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***AsynchronousReportsApiDeleteAsyncReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AsynchronousReportsApiDeleteAsyncReportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **amazonAdsAccountId** | **optional.String**| The identifier of a DSP advertiser account. Optional. For details, see [this guide](guides/reporting/dsp/creating-reports). | 

### Return type

[**DeleteAsyncReportResponse**](DeleteAsyncReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.deleteasyncreportresponse.v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsyncReport**
> AsyncReport GetAsyncReport(ctx, reportId, amazonAdvertisingAPIClientId, optional)
Gets a generation status of report by id

Gets a generation status of a report by id. Uses the `reportId` value from the response of previously requested report via `POST /reporting/reports` operation. When `status` is set to `COMPLETED`, the report will be available to be downloaded at `url`.  Report generation can take as long as 3 hours. Repeated calls to check report status may generate a 429 response, indicating that your requests have been throttled. To retrieve reports programmatically, your application logic should institute a delay between requests. For more information, see [Retry logic with exponential backoff](concepts/rate-limiting#use-retry-logic-with-exponential-backoff).   **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Scope  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"nemo_report_edit\",\"nemo_report_view\",\"reports_edit\"]  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"nemo_report_edit\",\"nemo_report_view\",\"reports_edit\",\"view_performance_dashboard\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **string**| The identifier of the requested report. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***AsynchronousReportsApiGetAsyncReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AsynchronousReportsApiGetAsyncReportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **amazonAdsAccountId** | **optional.String**| The identifier of a DSP advertiser account. Optional. For details, see [this guide](guides/reporting/dsp/creating-reports). | 

### Return type

[**AsyncReport**](AsyncReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.getasyncreportresponse.v3+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

