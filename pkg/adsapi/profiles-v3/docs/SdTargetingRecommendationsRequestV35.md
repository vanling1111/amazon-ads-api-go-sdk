# SdTargetingRecommendationsRequestV35

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tactic** | [***SdTacticV31**](SDTacticV31.md) |  | [default to null]
**Products** | [***[]SdAdvertisedProduct**](array.md) |  | [default to null]
**TypeFilter** | [***[]SdRecommendationTypeV33**](array.md) |  | [default to null]
**Themes** | [***SdTargetingRecommendationsThemes**](SDTargetingRecommendationsThemes.md) |  | [optional] [default to null]
**CategoryType** | **string** | This field is optional unless the field locationExpression is present in the request. It is used for category audience targeting to specify if the audience is for views (re-marketing) or purchases (re-purchasing). The specified categories will be returned accordingly. | [optional] [default to null]
**LocationExpression** | [**[]LocationExpression**](LocationExpression.md) | This optional field is used to specify the locations used in SD location targeting for non-Amazon sellers only at the moment. Therefore it&#x27;s only supported if the product is a landing page url. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

