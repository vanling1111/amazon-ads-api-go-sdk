# NegativeTargetingClauseEx

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | **float64** |  | [optional] [default to null]
**AdGroupId** | **float64** |  | [optional] [default to null]
**State** | **string** |  | [optional] [default to null]
**ExpressionType** | **string** |  | [optional] [default to null]
**Expression** | [**[]NegativeTargetingClauseExExpression**](NegativeTargetingClauseEx_expression.md) | The expression to negatively match against. * Only one brand may be specified per targeting expression. * Only one asin may be specified per targeting expression. * To exclude a brand from a targeting expression, you must create a negative targeting expression in the same ad group as the positive targeting expression. | [optional] [default to null]
**ServingStatus** | **string** | The status of the target. | [optional] [default to null]
**CreationDate** | **int64** | Epoch date the target was created. | [optional] [default to null]
**LastUpdatedDate** | **int64** | Epoch date of the last update to any property associated with the target. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

