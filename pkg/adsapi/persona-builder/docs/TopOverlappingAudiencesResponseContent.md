# TopOverlappingAudiencesResponseContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | [**[]AudienceInsight**](AudienceInsight.md) | Top audiences associated with customers in the input expression, ordered by the affinity score.                       Affinity is a measure of how likely customers in the input expression are to belong to a specific                     audience. An affinity of 5 for some audience indicates that customers in the input expression set                     are 5 times as likely to belong to this audience than the average of customers on Amazon. | [default to null]
**LastUpdatedAt** | [**time.Time**](time.Time.md) | UTC timestamp in ISO 8601 format indicating when insight was last generated for the input audience set. | [default to null]
**NextToken** | **string** | Optional: If present, there are more insights than initially returned. Use this token to call the operation again and have the additional insights returned. The token is valid for 8 hours from the initial request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

