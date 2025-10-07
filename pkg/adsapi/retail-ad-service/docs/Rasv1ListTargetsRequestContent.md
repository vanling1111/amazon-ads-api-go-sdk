# Rasv1ListTargetsRequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdGroupIdFilter** | [***Rasv1ReducedEntityIdFilter**](RASv1ReducedEntityIdFilter.md) |  | [optional] [default to null]
**CampaignIdFilter** | [***Rasv1ReducedEntityIdFilter**](RASv1ReducedEntityIdFilter.md) |  | [optional] [default to null]
**KeywordFilter** | [***Rasv1KeywordFilter**](RASv1KeywordFilter.md) |  | [optional] [default to null]
**MatchTypeFilter** | [**[]Rasv1KeywordMatchType**](RASv1KeywordMatchType.md) | Only keywords with match type that is in this list will be listed | [optional] [default to null]
**MaxResults** | **int32** | Number of records to include in the paginated response. Defaults to max page size for given API | [optional] [default to 1000]
**NegativeFilter** | [***Rasv1NegativeTargetFilter**](RASv1NegativeTargetFilter.md) |  | [optional] [default to null]
**NextToken** | **string** | token value allowing to navigate to the next response page | [optional] [default to null]
**ProductIdFilter** | [***Rasv1ProductIdFilter**](RASv1ProductIdFilter.md) |  | [optional] [default to null]
**StateFilter** | [***Rasv1StateFilter**](RASv1StateFilter.md) |  | [optional] [default to null]
**TargetIdFilter** | [***Rasv1EntityIdFilter**](RASv1EntityIdFilter.md) |  | [optional] [default to null]
**TargetLevelFilter** | [***Rasv1TargetLevelFilter**](RASv1TargetLevelFilter.md) |  | [optional] [default to null]
**TargetTypeFilter** | [**[]Rasv1TargetType**](RASv1TargetType.md) | Only targets of specified types will be listed | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

