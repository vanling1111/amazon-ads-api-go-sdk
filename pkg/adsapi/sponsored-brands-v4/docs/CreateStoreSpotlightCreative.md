# CreateStoreSpotlightCreative

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandLogoCrop** | [***BrandLogoCrop**](BrandLogoCrop.md) |  | [optional] [default to null]
**BrandName** | **string** |  | [optional] [default to null]
**Subpages** | [**[]Subpage**](Subpage.md) |  | [optional] [default to null]
**ConsentToTranslate** | **bool** | If set to true and the headline and/or video are not in the marketplace&#x27;s default language, Amazon will attempt to translate them to the marketplace&#x27;s default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information. | [optional] [default to null]
**CreativePropertiesToOptimize** | [**[]CreativePropertyToOptimize**](CreativePropertyToOptimize.md) | If this property is enabled, Sponsored Brands will dynamically optimize by enhancing or generating creative properties based on shopper search intent. | [optional] [default to null]
**BrandLogoAssetID** | **string** |  | [optional] [default to null]
**Headline** | **string** | The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

