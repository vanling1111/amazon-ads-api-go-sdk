# SbKeywordRecommendationRequestUrl

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Goal** | [***SbKeywordRecommendationGoal**](SBKeywordRecommendationGoal.md) |  | [optional] [default to null]
**CreativeType** | [***SbKeywordRecommendationCreativeType**](SBKeywordRecommendationCreativeType.md) |  | [optional] [default to null]
**MaxNumSuggestions** | **int64** | Maximum number of suggestions to return. Max value is 1000. If not provided, default to 100. | [optional] [default to null]
**CreativeAsins** | **[]string** |  | [optional] [default to null]
**Locale** | **string** | Optional locale to request keyword suggestion translations. For example, to request Simplified Chinese translations in US, provide locale “zh_CN”. Response will include both keyword suggestions and their translations. Supported locales include: Simplified Chinese (locale: “zh_CN”) for US, UK and CA. English (locale: “en_GB”) for DE, FR, IT and ES. | [optional] [default to null]
**Url** | **string** | The URL of the Stores page, or, Vendors may also specify the URL of a custom landing page. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

