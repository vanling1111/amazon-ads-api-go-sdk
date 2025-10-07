# MigrationJobStatusResponseContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | [optional] [default to null]
**MigrationJobStatus** | **string** | Enumerated status code for migration job status | Status                                             |  Description | |----------------------------------------------------------|--------------| | COMPLETE  | Migration job is complete | | FAILED    | Migration failed and no V3 campaigns were migrated | | IN_PROGRESS    | Migration job is running | | [optional] [default to null]
**MigrationJobStatusReason** | **string** | Status reason for the migration job status | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

