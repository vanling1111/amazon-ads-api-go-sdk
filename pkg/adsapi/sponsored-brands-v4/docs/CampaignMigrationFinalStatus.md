# CampaignMigrationFinalStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegacyCampaignId** | **string** | Entity object identifier. | [optional] [default to null]
**NewCampaignId** | **string** |  | [optional] [default to null]
**MigrationStatus** | **string** | Enumerated status code for migration job status | Status                                             |  Description | |----------------------------------------------------------|--------------| | MIGRATION_COMPLETE  | Migration is complete and new V4 campaigns are ready to serve. The V3 campaigns are archived | | STAGING_COMPLETE    | Staging of V4 campaign IDs are complete | | MIGRATION_FAILED         | Migration of V3 campaign failed  | | MIGRATION_IN_PROGRESS   | Migration for V3 campaign is in-progress | | [optional] [default to null]
**MigrationStatusReason** | **string** | Status reason for the given migration status | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

