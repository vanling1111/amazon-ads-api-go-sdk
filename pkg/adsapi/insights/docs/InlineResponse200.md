# InlineResponse200

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Marketplace** | **string** | The locale used to generate the overlapping audiences. | [default to null]
**NextToken** | **string** | If present, there are more overlapping audiences than initially returned. Use this token to call the operation again and have the additional overlapping audiences returned. The token is valid for 8 hours from the initial request. Note: subsequent calls must be made using the same parameters as used in previous requests. | [optional] [default to null]
**OverlappingAudiences** | [**[]InsightsAudiencesOverlapEntryV2**](insightsAudiencesOverlapEntryV2.md) | The list of overlapping audiences. | [default to null]
**RequestedAudienceMetadata** | [***InsightsAudiencesOverlapAudienceMetadataV2**](insightsAudiencesOverlapAudienceMetadataV2.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

