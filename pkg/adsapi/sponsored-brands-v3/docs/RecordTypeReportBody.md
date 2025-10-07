# RecordTypeReportBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportDate** | **string** | Date in YYYYMMDD format. The report only contains performance data for the specified date. The time zone is specified by the profile used to request the report. If this date is today, then the performance report may contain partial information. Reports are not available for data older than 60 days. For details on data latency, see [Service Guarantees](../../reference/concepts/developer-notes). | [default to null]
**Segment** | **string** | Optional. Allows you to run a report based on secondary dimensions. Use &#x60;placement&#x60; to segment a &#x60;campaigns&#x60; report by the location on a page where your ad appears. Use &#x60;query&#x60; to segment a &#x60;keywords&#x60; report and create a search terms report. | [optional] [default to null]
**CreativeType** | **string** | Optional. Supported values are &#x60;video&#x60; (for video campaigns) and &#x60;all&#x60; (for both non-video and video campaigns). If not specified, the report will contain data for non-video campaigns. For &#x60;ads&#x60; report types only, &#x60;creativeType&#x60; is required and must be set to &#x60;all&#x60; (&#x60;video&#x60; is not allowed). To return performance data for [multi-ad group campaigns](https://advertising.amazon.com/API/docs/en-us/sponsored-brands/3-0/openapi/prod#/Campaigns), you must include &#x60;creativeType&#x60; set to &#x60;all&#x60;&#x60; in your request. | [optional] [default to null]
**Metrics** | **string** | Each report type supports different metrics. **To understand supported metrics for each report type, see [Report types](/API/docs/en-us/guides/reporting/v2/report-types).**  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

