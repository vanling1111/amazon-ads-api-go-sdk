# ModerationResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EtaForModeration** | **string** | Expected date and time by which moderation will be complete. The format is ISO 8601 in UTC time zone. Note that this field is present in the response only when moderationStatus is set to IN_PROGRESS. | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**IdType** | [***IdType**](IdType.md) |  | [optional] [default to null]
**ModerationStatus** | [***ModerationStatus**](ModerationStatus.md) |  | [optional] [default to null]
**PolicyViolations** | [**[]PolicyViolation**](PolicyViolation.md) | A list of policy violations for a campaign that has failed moderation. Note that this field is present in the response only when moderationStatus is set to REJECTED. | [optional] [default to null]
**VersionId** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

