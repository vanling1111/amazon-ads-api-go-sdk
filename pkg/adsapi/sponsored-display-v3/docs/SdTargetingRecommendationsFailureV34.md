# SdTargetingRecommendationsFailureV34

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | HTTP status code indicating a failure response for targeting recomendations. | [optional] [default to null]
**Name** | **string** | The theme name specified in the request. If the themes field is not provided in the request, the value of this field will be set to default. | [optional] [default to null]
**Expression** | [**[]SdProductTargetingThemeExpression**](SDProductTargetingThemeExpression.md) | A list of expressions that failed to be applied in the product targeting theme. | [optional] [default to null]
**ErrorMessage** | **string** | A human friendly error message indicating the failure reasons. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

