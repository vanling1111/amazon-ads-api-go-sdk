# CreativeImageRecommendationRequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asins** | **[]string** | ----------------------------------------------- List types ----------------------------------------------- A list of ASINs | [default to null]
**AssetSubType** | [***AssetSubType**](AssetSubType.md) |  | [optional] [default to null]
**MaxNumRecommendations** | **float64** | Maximum number of recommendations that API should return. Response will [0, recommendations] recommendations (recommendations are not guaranteed). | [optional] [default to null]
**AssetPrograms** | [**[]ProgramType**](ProgramType.md) | Filter assets by program types. For example, if only [A_PLUS] assets are requested then only assets that were used as A+ content will be recommended. If no program type is provided, recommend assets from all programs | [optional] [default to null]
**Locale** | **string** | (Optional) locale of creative headline and ASIN titles. If locale is not provided, default locale of marketplace is used. Currently, only en_US and en_CA are supported. | [optional] [default to null]
**Headline** | **string** | The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

