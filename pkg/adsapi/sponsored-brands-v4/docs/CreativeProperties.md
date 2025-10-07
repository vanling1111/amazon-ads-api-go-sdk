# CreativeProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandLogoCrop** | [***AssetCrop**](AssetCrop.md) |  | [optional] [default to null]
**BrandName** | **string** | The displayed brand name in the ad headline. Maximum length is 30 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements. | [optional] [default to null]
**CustomImageAssetId** | **string** | The identifier of image/video asset from the store&#x27;s asset library | [optional] [default to null]
**LandingPage** | [***CreativeLandingPage**](CreativeLandingPage.md) |  | [optional] [default to null]
**CustomImages** | [**[]CustomImage**](CustomImage.md) | An array of customImages associated with the creative. | [optional] [default to null]
**ConsentToTranslate** | **bool** | If set to true and the headline and/or video are not in the marketplace&#x27;s default language, Amazon will attempt to translate them to the marketplace&#x27;s default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information. | [optional] [default to null]
**CustomImageCrop** | [***AssetCrop**](AssetCrop.md) |  | [optional] [default to null]
**CustomImageUrl** | **string** |  | [optional] [default to null]
**OriginalVideoAssetIds** | **[]string** | The assetIds of the original videos submitted by the advertiser. If &#x27;consentToTranslate&#x27; is set to true and translation is SUCCESSFUL then &#x60;originalVideoAssetIds&#x60; will return the original video assetId whereas &#x60;videoAssetIds&#x60; will return translated video assetId. In all other cases, &#x27;originalVideoAssetIds&#x27; and &#x60;videoAssetIds&#x60; both will return original video assetId. | [optional] [default to null]
**Asins** | **[]string** | ----------------------------------------------- List types ----------------------------------------------- A list of ASINs | [optional] [default to null]
**BrandLogoUrl** | **string** |  | [optional] [default to null]
**Subpages** | [**[]Subpage**](Subpage.md) | An array of subpages | [optional] [default to null]
**CreativePropertiesToOptimize** | [**[]CreativePropertyToOptimize**](CreativePropertyToOptimize.md) | If this property is enabled, Sponsored Brands will dynamically optimize by enhancing or generating creative properties based on shopper search intent. | [optional] [default to null]
**OriginalHeadline** | **string** | The original headline submitted by the advertiser. | [optional] [default to null]
**VideoAssetIds** | **[]string** | The assetIds of the original videos submitted by the advertiser. If &#x27;consentToTranslate&#x27; is set to true and translation is SUCCESSFUL then &#x27;videoAssetIds&#x27; will return translated video assetId whereas &#x60;originalVideoAssetIds&#x60; will return the original video assetId. In all other cases, &#x60;videoAssetIds&#x60; will return original video assetId. | [optional] [default to null]
**BrandLogoAssetId** | **string** | The identifier of image/video asset from the store&#x27;s asset library | [optional] [default to null]
**Headline** | **string** | If &#x27;consentToTranslate&#x27; is set to true and translation is SUCCESSFUL then &#x60;headline&#x60; will return the translated headline whereas &#x60;originalHeadline&#x60; will return the original headline. In all other cases, &#x27;originalHeadline&#x27; and &#x60;headline&#x60; both will return the original headline. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

