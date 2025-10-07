# MmmReport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | [***MmmReportConfiguration**](MmmReportConfiguration.md) |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time when the report was created. | [optional] [default to null]
**Description** | **string** | A description of the report. | [optional] [default to null]
**DueDate** | **string** | The due date of the report. | [optional] [default to null]
**EndDate** | **string** | Inclusive end of the reporting period. | [optional] [default to null]
**FailureCode** | **string** | An error code indicating why the report failed. Present when the status is &#x60;FAILED&#x60;. | [optional] [default to null]
**FailureMessage** | **string** | A human-readable message providing more information about the failure. Present when the status is &#x60;FAILED&#x60;. | [optional] [default to null]
**ReportId** | **string** | The unique identifier of the report. | [optional] [default to null]
**ReportName** | **string** | The display name of the report. | [optional] [default to null]
**StartDate** | **string** | Inclusive start of the reporting period. | [optional] [default to null]
**Status** | **string** | The report generation status. |Value|Description| |---|---| |PENDING|Report is created and awaiting processing.| |PROCESSING|Report is processing.| |SUCCEEDED|Report is completed. Check &#x60;urls&#x60; for the output files.| |FAILED|Report processing failed. Check the &#x60;failureCode&#x60; and &#x60;failureMessage&#x60; for details.| |CANCELED|Report is canceled. Contact &lt;mmm-support@amazon.com&gt; if this is unexpected.| | [optional] [default to null]
**Urls** | **[]string** | The URLs for downloading output files. Present when the status is &#x60;SUCCEEDED&#x60;. | [optional] [default to null]
**UrlsExpireAt** | [**time.Time**](time.Time.md) | The expiration date of the download URLs. Present when the status is &#x60;SUCCEEDED&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

