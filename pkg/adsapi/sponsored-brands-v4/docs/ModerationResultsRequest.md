# ModerationResultsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionIdFilter** | **[]string** | Filter by specific version id of the ad. The API will return the ad&#x27;s all versions moderation status if this field is empty. | [optional] [default to null]
**IdType** | [***IdType**](IdType.md) |  | [default to null]
**AdProgramType** | [***ModerationResultsAdProgramType**](ModerationResultsAdProgramType.md) |  | [default to null]
**NextToken** | **string** |  | [optional] [default to null]
**MaxResults** | **int32** | Sets a limit on the number of results returned by an operation. | [default to null]
**ModerationStatusFilter** | [**[]ModerationStatus**](ModerationStatus.md) | Filter by specific moderation status. | [optional] [default to null]
**Id** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

