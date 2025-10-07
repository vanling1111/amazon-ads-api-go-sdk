# TargetingClauseEx

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | **float64** | The target identifier. | [optional] [default to null]
**CampaignId** | **float64** | The identifier of the campaign to which this target is associated. | [optional] [default to null]
**AdGroupId** | **float64** | The identifier of the ad group to which this target is associated. | [optional] [default to null]
**State** | [***State**](State.md) |  | [optional] [default to null]
**Expression** | [**[]TargetingExpressionPredicate**](TargetingExpressionPredicate.md) | The targeting expression. | [optional] [default to null]
**ResolvedExpression** | [**[]TargetingExpressionPredicate**](TargetingExpressionPredicate.md) | The resolved targeting expression. | [optional] [default to null]
**ExpressionType** | [***ExpressionType**](ExpressionType.md) |  | [optional] [default to null]
**Bid** | **float32** |  | [optional] [default to null]
**CreationDate** | **float64** | The epoch time that the targeting clause was created. | [optional] [default to null]
**LastUpdatedDate** | **float64** | The epoch time that the targeting clause was updated. | [optional] [default to null]
**ServingStatus** | **string** | The computed status of the targeting clause. See the [developer notes](https://advertising.amazon.com/API/docs/en-us/reference/concepts/developer-notes) for more information. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

