# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateBrandMetricsReport**](ReportApi.md#GenerateBrandMetricsReport) | **Post** /insights/brandMetrics/report | Generate Brand Metrics Report. Each response record will include the following dimensional fields (in addition to the requested metrics) brand Namecategory, TreeNamecategory, HierarchylookbackPeriod, metricsComputationDate 
[**GetBrandMetricsReport**](ReportApi.md#GetBrandMetricsReport) | **Get** /insights/brandMetrics/report/{reportId} | Retrieve the status and the URL of the Brand Metrics Report being generated

# **GenerateBrandMetricsReport**
> BrandMetricsGenerateReportResponse GenerateBrandMetricsReport(ctx, amazonAdvertisingAPIScope, amazonAdvertisingAPIClientId, optional)
Generate Brand Metrics Report. Each response record will include the following dimensional fields (in addition to the requested metrics) brand Namecategory, TreeNamecategory, HierarchylookbackPeriod, metricsComputationDate 

Generates the Brand Metrics report in CSV or JSON format. Customize the report by passing a specific categoryTreeName, categoryPath, brandName, reportStartDate, reportEndDate, lookbackPeriod, format or a list of metrics from the available metrics in the metrics field. If an empty request body is passed, report for the latest available report date in JSON format will get generated with all the available brands and metrics for an advertiser. The report may or may not contain the Brand Metrics data for one or more brands depending on data availability.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"nemo_report_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***ReportApiGenerateBrandMetricsReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportApiGenerateBrandMetricsReportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of BrandMetricsGenerateReportRequest**](BrandMetricsGenerateReportRequest.md)| Create request body to generate the Brand Metrics Report. | 

### Return type

[**BrandMetricsGenerateReportResponse**](brandMetricsGenerateReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.insightsbrandmetrics.v1+json, application/vnd.insightsbrandmetrics.v1.1+json
 - **Accept**: application/vnd.insightsbrandmetrics.v1+json, application/vnd.insightsbrandmetrics.v1.1+json, application/vnd.insightsbrandmetricserror.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBrandMetricsReport**
> BrandMetricsGetReportByIdResponse GetBrandMetricsReport(ctx, reportId, amazonAdvertisingAPIScope, amazonAdvertisingAPIClientId)
Retrieve the status and the URL of the Brand Metrics Report being generated

Fetch the location and status of the report for the brands for which the metrics are available. The URL to the report is only available when the status of the report is SUCCESSFUL.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"nemo_report_edit\",\"nemo_report_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **string**| The report Id to be fetched. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 

### Return type

[**BrandMetricsGetReportByIdResponse**](brandMetricsGetReportByIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.insightsbrandmetrics.v1+json, application/vnd.insightsbrandmetrics.v1.1+json, application/vnd.insightsbrandmetricserror.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

