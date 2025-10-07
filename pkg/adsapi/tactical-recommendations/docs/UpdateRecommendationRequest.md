# UpdateRecommendationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetRule** | [***UpdateBudgetRule**](UpdateBudgetRule.md) |  | [optional] [default to null]
**RecommendedValue** | **string** | Recommended value of the recommendation. Type of data expected for each recommendation type: | Recommendation type | Data type | |---|---| | CAMPAIGN_BUDGET, CAMPAIGN_TOP_PLACEMENT, CAMPAIGN_PRODUCT_PLACEMENT, AD_GROUP_DEFAULT_BID, NEW_KEYWORD, KEYWORD_BID, NEW_PRODUCT_TARGETING, PRODUCT_TARGETING_BID, NEW_AUDIENCE_TARGETING, AUDIENCE_TARGETING_BID | number | | CAMPAIGN_END_DATE | string($YYYY-MM-DD) | | AD_GROUP_BID_OPTIMIZATION | One of [ CLICKS, CONVERSIONS, REACH ] | | CAMPAIGN_BIDDING_STRATEGY | One of [ LEGACY_FOR_SALES, AUTO_FOR_SALES, MANUAL ] | | CAMPAIGN_STATE, AD_GROUP_STATE, KEYWORD_STATE, NEGATIVE_KEYWORD_STATE, PRODUCT_AD_STATE, PRODUCT_TARGETING_STATE, NEGATIVE_PRODUCT_TARGETING_STATE, AUDIENCE_TARGETING_STATE, NEGATIVE_AUDIENCE_TARGETING_STATE | One of [ ENABLED, PAUSED, ARCHIVED ] |  | [optional] [default to null]
**RuleBasedBidding** | [***UpdateRuleBasedBidding**](UpdateRuleBasedBidding.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

