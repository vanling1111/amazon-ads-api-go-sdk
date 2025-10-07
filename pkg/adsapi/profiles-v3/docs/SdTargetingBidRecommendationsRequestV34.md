# SdTargetingBidRecommendationsRequestV34

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | [**[]SdGoalProduct**](SDGoalProduct.md) | A list of products to tailor bid recommendations for category and audience based targeting clauses. This array must contain consistent fields of either asins or landing pages (when linking to other pages), these cannot be mixed for any given request. If landingPageUrl is used, only one item is allowed for the list. | [optional] [default to null]
**BidOptimization** | [***SdBidOptimizationV32**](SDBidOptimizationV32.md) |  | [default to null]
**CostType** | [***SdCostTypeV31**](SDCostTypeV31.md) |  | [default to null]
**CreativeType** | [***SdCreativeType**](SDCreativeType.md) |  | [optional] [default to null]
**TargetingClauses** | [**[]SdTargetingBidRecommendationsRequestV34TargetingClauses**](SDTargetingBidRecommendationsRequestV34_targetingClauses.md) | A list of targeting clauses to receive bid recommendations for. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

