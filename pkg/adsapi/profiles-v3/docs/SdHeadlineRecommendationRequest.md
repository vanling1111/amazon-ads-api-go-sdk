# SdHeadlineRecommendationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asins** | **[]string** | An array of ASINs associated with the creative. | [optional] [default to null]
**MaxNumRecommendations** | **float64** | Maximum number of recommendations that API should return. Response will [0, maxNumRecommendations] recommendations (recommendations are not guaranteed as there can be instances where the ML model can not generate policy compliant headlines for the given set of asins). | [optional] [default to null]
**AdFormat** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

