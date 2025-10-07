# SdForecastRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaign** | [***Campaign**](Campaign.md) |  | [default to null]
**AdGroup** | [***AdGroup**](AdGroup.md) |  | [default to null]
**OptimizationRules** | [**[]OptimizationRule**](OptimizationRule.md) | A list of SD optimization rules. Forecast will be affected by the optimization strategy rules.  Currently, supported rule metrics by forecast are &#x60;COST_PER_CLICK&#x60;, &#x60;COST_PER_THOUSAND_VIEWABLE_IMPRESSIONS&#x60; and &#x60;COST_PER_ORDER&#x60;. | [optional] [default to null]
**ProductAds** | [**[]ProductAd**](ProductAd.md) |  | [default to null]
**TargetingClauses** | [**[]SdForecastRequestTargetingClause**](SDForecastRequestTargetingClause.md) | A list of SD targeting clauses. | [default to null]
**NegativeTargetingClauses** | [**[]NegativeTargetingClause**](NegativeTargetingClause.md) | A list of SD negative targeting clauses. | [optional] [default to null]
**LocationExpressions** | [**[]LocationExpression**](LocationExpression.md) | A list of location expressions. Only applicable for advertisers using landingPageType of OFF_AMAZON_LINK. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

