# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMmmReport**](ReportsApi.md#CreateMmmReport) | **Post** /mmm/v1/reports | Create a report
[**DeleteMmmReport**](ReportsApi.md#DeleteMmmReport) | **Delete** /mmm/v1/reports/{reportId} | Delete a report
[**GetMmmReport**](ReportsApi.md#GetMmmReport) | **Get** /mmm/v1/reports/{reportId} | Get a report

# **CreateMmmReport**
> MmmReport CreateMmmReport(ctx, optional)
Create a report

Creates a report. Each report requires the `brandGroupId` of a predefined brand group listed in `POST /mmm/v1/brandGroups/list`. The response will include a `reportId` that can be used to poll the status and results.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"MasterAccount_Manager\",\"MasterAccount_Viewer\",\"ManagerAccount_Dev\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsApiCreateMmmReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiCreateMmmReportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V1ReportsBody**](V1ReportsBody.md)|  | 

### Return type

[**MmmReport**](MmmReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMmmReport**
> DeleteMmmReport(ctx, )
Delete a report

Deletes a report by ID. Use this operation to cancel or clean up a report.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"MasterAccount_Manager\",\"MasterAccount_Viewer\",\"ManagerAccount_Dev\"]

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMmmReport**
> MmmReport GetMmmReport(ctx, )
Get a report

Gets the generation status of a report by ID. The `reportId` is found in the response to creating a report using `POST /mmm/v1/reports`. When the `status` is `SUCCESSFUL` the output files will be available for download at `urls`. A report may take up to 24 hours to be processed. Repeated calls to check report status may generate a 429 response, indicating that your requests have been throttled. To retrieve reports programmatically, your application logic should institute a delay between requests.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"MasterAccount_Manager\",\"MasterAccount_Viewer\",\"ManagerAccount_Dev\"]

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MmmReport**](MmmReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

