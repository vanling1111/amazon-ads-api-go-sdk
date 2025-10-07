# AsyncReportConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdProduct** | [***AsyncReportAdProduct**](AsyncReportAdProduct.md) |  | [default to null]
**Columns** | **[]string** | The list of columns to be used for report. The availability of columns depends on the selection of reportTypeId. This list cannot be null or empty.  | [default to null]
**Filters** | [**[]AsyncReportFilter**](AsyncReportFilter.md) | The list of filters supported by a report type. The availability of filters fields depends on the selection of reportTypeId. | [optional] [default to null]
**Format** | **string** | The report file format. | [default to null]
**GroupBy** | **[]string** | This field determines the aggregation level of the report data and also makes additional fields available for selection. This field cannot be null or empty.  | [default to null]
**ReportTypeId** | **string** | The identifier of the Report Type to be generated. | [default to null]
**TimeUnit** | **string** | The aggregation level of report data. If the timeUnit is set to &#x60;SUMMARY&#x60;, the report data is aggregated at the time period specified. The availability of time unit breakdowns depends on the selection of reportTypeId.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

