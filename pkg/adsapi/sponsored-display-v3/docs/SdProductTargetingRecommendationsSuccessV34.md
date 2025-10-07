# SdProductTargetingRecommendationsSuccessV34

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | HTTP status code 200 indicating a successful response for product recommendations. | [optional] [default to null]
**Name** | **string** | The theme name specified in the request. | [optional] [default to null]
**Expression** | [**[]SdProductTargetingThemeExpression**](SDProductTargetingThemeExpression.md) | A list of expressions defining the product targeting theme. The list will define an AND operator on different expressions. For example, asinPriceGreaterThan and asinReviewRatingLessThan can be used to request product recommendations which are both with greater price and less review rating compared to the goal products. Note: currently the service only support one item in the array. | [optional] [default to null]
**Recommendations** | [**[]SdProductRecommendationV32**](SDProductRecommendationV32.md) | A list of recommended products. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

