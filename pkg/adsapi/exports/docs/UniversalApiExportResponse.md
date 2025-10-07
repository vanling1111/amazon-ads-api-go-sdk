# UniversalApiExportResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | Date of when the export request was created. | [optional] [default to null]
**Error_** | [***UniversalApiError**](UniversalApiError.md) |  | [optional] [default to null]
**ExportId** | **string** | The export identifier. | [default to null]
**FileSize** | **float64** | Byte size of the generated file. | [optional] [default to null]
**GeneratedAt** | [**time.Time**](time.Time.md) | Date of when the export was finished generating. | [optional] [default to null]
**Status** | **string** | The generation status of the export. - PROCESSING: Export is currently in progress. - COMPLETED: Export has completed successfully. - FAILED: Export has failed. See the error message for more details.  | [default to null]
**Url** | **string** | A URL for the export. It’s only available if status is COMPLETED. | [optional] [default to null]
**UrlExpiresAt** | [**time.Time**](time.Time.md) | Date at which the download URL for the generated export expires. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

