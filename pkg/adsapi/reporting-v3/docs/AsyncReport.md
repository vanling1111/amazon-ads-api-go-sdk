# AsyncReport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | [***AsyncReportConfiguration**](AsyncReportConfiguration.md) |  | [default to null]
**CreatedAt** | **string** | The date at which the report was created in ISO 8601 date time format. | [default to null]
**EndDate** | **string** | The end date for the reporting period in YYYY-mm-dd format. | [default to null]
**FailureReason** | **string** | Present for failed reports only. The reason why a report failed to generate. | [optional] [default to null]
**FileSize** | **float64** | The size of the report file, in bytes. | [optional] [default to null]
**GeneratedAt** | **string** | The date at which the report was generated in ISO 8601 date time format. | [optional] [default to null]
**Name** | **string** | Optional. The name of the generated report. | [optional] [default to null]
**ReportId** | **string** | The identifier of the requested report. | [default to null]
**StartDate** | **string** | The start date for the reporting period in YYYY-mm-dd format. | [default to null]
**Status** | **string** | The build status of the report.   - &#x60;PENDING&#x60; - Report is created and awaiting processing.   - &#x60;PROCESSING&#x60; - Report is processing. Please wait.   - &#x60;COMPLETED&#x60; - Report has completed.  Check the &#x60;url&#x60; for the output file.   - &#x60;FAILED&#x60; - Report generation failed.  Check the &#x60;failureReason&#x60; for details.  | [default to null]
**UpdatedAt** | **string** | The date at which the report was last updated in ISO 8601 date time format. | [default to null]
**Url** | **string** | URL of the generated report. | [optional] [default to null]
**UrlExpiresAt** | **string** | The date at which the download URL for the generated report expires. urlExpires at this time defaults to 3600 seconds but may vary in the future. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

