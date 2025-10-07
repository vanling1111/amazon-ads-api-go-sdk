# BrandVideoCreative

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asins** | **[]string** | An array of ASINs associated with the creative. | [default to null]
**BrandLogoCrop** | [***AssetCrop**](AssetCrop.md) |  | [optional] [default to null]
**BrandName** | **string** | The displayed brand name in the ad headline. Maximum length is 30 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements. | [default to null]
**LandingPage** | [***CreativeLandingPageV2**](CreativeLandingPageV2.md) |  | [optional] [default to null]
**ConsentToTranslate** | **bool** | If set to true and the headline and/or video are not in the marketplace&#x27;s default language, Amazon will attempt to translate them to the marketplace&#x27;s default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information. | [optional] [default to null]
**VideoAssetIds** | **[]string** | The assetIds of the original videos submitted by the advertiser. If &#x27;consentToTranslate&#x27; is set to true and translation is SUCCESSFUL then &#x27;videoAssetIds&#x27; will return translated video assetId whereas &#x60;originalVideoAssetIds&#x60; will return the original video assetId. In all other cases, &#x60;videoAssetIds&#x60; will return original video assetId. | [default to null]
**BrandLogoAssetId** | **string** | The identifier of the [brand logo](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#brandlogo) image from the brand store&#x27;s asset library. Note that for campaigns created in the Amazon Advertising console prior to release of the brand store&#x27;s assets library, responses will not include a value for this field. | [default to null]
**Headline** | **string** | The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

