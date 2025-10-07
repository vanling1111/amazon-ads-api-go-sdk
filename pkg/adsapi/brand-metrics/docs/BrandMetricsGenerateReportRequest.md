# BrandMetricsGenerateReportRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandName** | **string** | Optional. Brand Name. If no Brand Name is passed, then all data available for all brands belonging to the entity are retrieved. | [optional] [default to null]
**CategoryPath** | **[]string** | Optional. The hierarchical path that leads to a node starting with the root node. If no Category Node Name is passed, then all data available for all brands belonging to the entity are retrieved. | [optional] [default to null]
**CategoryTreeName** | **string** | Optional. The node at the top of a browse tree. It is the start node of a tree. | [optional] [default to null]
**Format** | **string** | Format of the report. | [optional] [default to FORMAT.JSON]
**LookBackPeriod** | **string** | Currently supported values: \&quot;1w\&quot; (one week), \&quot;1m\&quot; (one month) and  \&quot;1cm\&quot; (one calendar month). This defines the period of time used to determine the number of shoppers in the metrics computation. | [optional] [default to LOOK_BACK_PERIOD.1W_]
**Metrics** | **[]string** | Optional. Specify an array of string of metrics field names to include in the report. If no metric field names are specified, all metrics are returned. | [optional] [default to null]
**ReportEndDate** | **string** | Optional. Retrieves metrics with metricsComputationDate between reportStartDate and reportEndDate (inclusive). The maximum allowed date range for report generation is 3 months. The date will be in the Coordinated Universal Time (UTC) timezone in YYYY-MM-DD format. If both reportStartDate and reportEndDate are passed and the range is greater than 3 months, the reportStartDate will be adjusted to a date 3 months from the reportEndDate. If no date is passed in reportEndDate, all available metrics till metricsComputationDate of 3 months after the reportStartDate will be provided. If no date is passed for either reportStartDate or reportEndDate, the metrics with the most recent metricsComputationDate will be returned. | [optional] [default to null]
**ReportStartDate** | **string** | Optional. Retrieves metrics with metricsComputationDate between reportStartDate and reportEndDate (inclusive). The maximum allowed date range for report generation is 3 months. The date will be in the Coordinated Universal Time (UTC) timezone in YYYY-MM-DD format. If both reportStartDate and reportEndDate are passed and the range is greater than 3 months, the reportStartDate will be adjusted to a date 3 months from the reportEndDate. If no date is passed in reportStartDate, all available metrics from metricsComputationDate of 3 months before the reportEndDate will be provided. If no date is passed for either reportStartDate or reportEndDate, the metrics with the most recent metricsComputationDate will be returned. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

