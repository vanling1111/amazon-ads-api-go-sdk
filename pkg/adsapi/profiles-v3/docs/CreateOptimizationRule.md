# CreateOptimizationRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The state of the optimization rule. | [default to null]
**RuleName** | **string** | The name of the optimization rule. | [optional] [default to null]
**RuleConditions** | [**[]RuleCondition**](RuleCondition.md) | A list of rule conditions that define the advertiser&#x27;s intent for the outcome of the rule. The rule uses &#x27;AND&#x27; logic to combine every condition in this list, and will validate the combination when the rule is created or updated. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

