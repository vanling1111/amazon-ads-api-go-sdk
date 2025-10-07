# CreativeRecommendationProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asins** | **[]string** | ----------------------------------------------- List types ----------------------------------------------- A list of ASINs | [optional] [default to null]
**BrandName** | **string** | The displayed brand name in the ad headline. Maximum length is 30 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements. | [optional] [default to null]
**Subpages** | [**[]Subpage**](Subpage.md) | An array of subpages | [optional] [default to null]
**LandingPage** | [***CreativeLandingPageV2**](CreativeLandingPageV2.md) |  | [optional] [default to null]
**CustomImages** | [**[]CustomImage**](CustomImage.md) | An array of customImages associated with the creative. | [optional] [default to null]
**VideoAssetIds** | **[]string** | An array of videoAssetIds associated with the creative. Advertisers can get video assetIds from Asset Library /assets/search API. | [optional] [default to null]
**RecommendedCreativeId** | **string** | a Unique Id identifying the creative Recommendation | [optional] [default to null]
**BrandLogo** | [***BrandLogo**](BrandLogo.md) |  | [optional] [default to null]
**Headline** | **string** | The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

