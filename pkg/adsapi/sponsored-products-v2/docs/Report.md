# Report

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateFilter** | **string** | Filters the response to include reports with &#x60;stateFilter&#x60; set to one of the values in the comma-delimited list. &#x60;stateFilter&#x60; and &#x60;segment&#x60; cannot be used in the same report request.  &#x60;asins&#x60; report types do not support the use of &#x60;stateFilter&#x60;. | [optional] [default to null]
**CampaignType** | **string** | The type of campaign. Only required for &#x60;asins&#x60; reports - don&#x27;t use with other report types. | [optional] [default to null]
**Segment** | **string** | A secondary dimension used to further segment certain types of reports. &#x60;stateFilter&#x60; and &#x60;segment&#x60; cannot be used in the same report request. Keyword search term reports only return search terms that have generated at least one click or sale.  **Note**: Search term reports for auto-targeted campaigns created before 11/14/2018 can be accessed from the &#x60;/v2/sp/keywords/report&#x60; resource. Search term reports for auto-targeted campaigns generated on-and-after 11/14/2018 can be accessed from the &#x60;/v2/sp/targets/report&#x60; resource.   | Dimension | Valid report types | Description | |-----------|-------------|-------------| | query | keywords, targets | Segments a report based on customer search term. | | placement | campaigns | Segments a &#x60;campaigns&#x60; report based on the page location where the ad appeared. | | [optional] [default to null]
**ReportDate** | **string** | The date for which to retrieve the performance report in YYYYMMDD format. The time zone is specified by the profile used to request the report. If this date is today, then the performance report may contain partial information. Reports are not available for data older than 60 days. For details on data latency, see the service guarantees in the [developer notes](https://advertising.amazon.com/API/docs/en-us/reference/concepts/developer-notes) section. | [optional] [default to null]
**Metrics** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

