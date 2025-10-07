# BrandMetricsGetReportByIdResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandsInfo** | [**[]BrandMetricsGetReportByIdResponseBrandsInfo**](brandMetricsGetReportByIdResponse_brandsInfo.md) | List of first 200 brands for which the Brand Metrics report is generated. The report may contain more than 200 brands. This list is only populated with brands if the Brand Metrics are available for the brands that an advertiser has access to. | [optional] [default to null]
**Expiration** | **int64** | The expiration time of the URI in the location property in milliseconds. The expiration time is the interval between the time the response was generated and the time the URI expires. | [default to null]
**Format** | **string** | Format of the report. | [default to FORMAT.JSON]
**Location** | **string** | The URI address of the report. Only available if the report is generated successfully. The location is empty if the Brand Metrics are not available or if the report is not generated successfully. | [optional] [default to null]
**ReportId** | **string** | The identifier of the report. | [default to null]
**Status** | **string** | The build status of the report. | [default to null]
**StatusDetails** | **string** | A human-readable description of the current status. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

