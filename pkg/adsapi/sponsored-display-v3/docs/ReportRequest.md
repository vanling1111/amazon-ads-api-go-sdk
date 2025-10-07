# ReportRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportDate** | **string** | Date in YYYYMMDD format. The report contains only metrics generated on the specified date. Note that the time zone used for date calculation is the one associated with the profile used to make the request. | [optional] [default to null]
**Tactic** | [***TacticReport**](TacticReport.md) |  | [optional] [default to null]
**Segment** | [***Segment**](Segment.md) |  | [optional] [default to null]
**Metrics** | **string** | A comma-separated list of the metrics to be included in the report.  Each report type supports different metrics. **To understand supported metrics for each report type, see [Report types](/API/docs/en-us/guides/reporting/v2/report-types).**  **Note**: Campaigns with vCPM costType should use view+click based metrics (viewAttributedConversions14d, viewAttributedDetailPageView14d, viewAttributedSales14d, viewAttributedUnitsOrdered14d, viewImpressions).  **Note**: Detail page view metrics (attributedDetailPageView14d, viewAttributedDetailPageView14d) have an SLA of 3 days.  **Tip**: Use new-to-brand (NTB) metrics to calculate how efficient your campaigns are at driving new shoppers:    1. Percentage of NTB orders &#x3D; attributedOrdersNewToBrand14d / attributedConversions14d   2. Percentage NTB sales &#x3D; attributedSalesNewToBrand14d / attributedSales14d   3. Percentage NTB units &#x3D; attributedUnitsOrderedNewToBrand14d / attributedUnitsOrdered14d   4. NTB order rate &#x3D; attributedOrdersNewToBrand14 / impressions | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

