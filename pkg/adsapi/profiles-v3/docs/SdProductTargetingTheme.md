# SdProductTargetingTheme

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | This is the meaningful theme name which will be used as a unique identifier across various themes in the same request. This identifier will also be used to map the recommendations back to the theme in the response body. Note: the value for this field cannot be \&quot;default\&quot; as that&#x27;s a reserved keyword in the system. | [default to null]
**Expression** | [**[]SdProductTargetingThemeExpression**](SDProductTargetingThemeExpression.md) | A list of expressions defining the contextual targeting theme. The list will define an AND operator on different expressions. For example, asinPriceGreaterThan and asinReviewRatingLessThan can be used to request product recommendations which are both with greater price and less review rating compared to the goal products. Note: Currently the service only supports one item in the array. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

