# NegativeTargetingClause

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | **int64** |  | [optional] [default to null]
**AdGroupId** | **int64** |  | [optional] [default to null]
**ExpressionType** | **string** |  | [optional] [default to null]
**Expression** | [**[]NegativeTargetingExpression**](NegativeTargetingExpression.md) | The expression to negatively match against. * Only one brand may be specified per targeting expression. * Only one asin may be specified per targeting expression. * To exclude a brand from a targeting expression, you must create a negative targeting expression in the same ad group as the positive targeting expression. | [optional] [default to null]
**ResolvedExpression** | [**[]NegativeTargetingExpression**](NegativeTargetingExpression.md) | The resolved negative targeting expression. | [optional] [default to null]
**State** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

