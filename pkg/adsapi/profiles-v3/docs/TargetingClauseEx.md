# TargetingClauseEx

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | **float64** |  | [optional] [default to null]
**AdGroupId** | **float64** |  | [optional] [default to null]
**CampaignId** | **float64** |  | [optional] [default to null]
**State** | **string** |  | [optional] [default to null]
**ExpressionType** | **string** |  | [optional] [default to null]
**Bid** | **float64** | If a value for &#x60;bid&#x60; is specified, it overrides the current adGroup bid. When using vcpm costType. $1 is the minimum bid for vCPM. Note that this field is ignored for negative targeting clauses. | [optional] [default to null]
**Expression** | [***[]OneOfTargetingExpressionItems**](array.md) |  | [optional] [default to null]
**ResolvedExpression** | [***[]OneOfTargetingExpressionItems**](array.md) |  | [optional] [default to null]
**ServingStatus** | **string** | The status of the target. | [optional] [default to null]
**CreationDate** | **int64** | Epoch date the target was created. | [optional] [default to null]
**LastUpdatedDate** | **int64** | Epoch date of the last update to any property associated with the target. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

