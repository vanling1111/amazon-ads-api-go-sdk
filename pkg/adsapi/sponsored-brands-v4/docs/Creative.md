# Creative

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandLogoCrop** | [***BrandLogoCrop**](BrandLogoCrop.md) |  | [optional] [default to null]
**BrandName** | **string** |  | [optional] [default to null]
**CustomImageAssetId** | **string** |  | [optional] [default to null]
**ConsentToTranslate** | **bool** | If set to true and the headline and/or video asset are not in the marketplace&#x27;s default language, Amazon will attempt to translate them to the marketplace&#x27;s default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information. | [optional] [default to null]
**CustomImages** | [**[]CustomImage**](CustomImage.md) | Requires minimum one custom image. You can add an optional collection of custom images that can be displayed on the ad as slideshow. Learn more about slideshow here https://advertising.amazon.com/resources/whats-new/slideshow-ads-creative-for-sponsored-brands/ | [optional] [default to null]
**CustomImageCrop** | [***CustomImageCrop**](CustomImageCrop.md) |  | [optional] [default to null]
**CustomImageUrl** | **string** |  | [optional] [default to null]
**Type_** | [***CreativeType**](CreativeType.md) |  | [optional] [default to null]
**OriginalVideoAssetIds** | **[]string** | The assetIds of the original videos submitted by the advertiser. If &#x27;consentToTranslate&#x27; is set to true and translation is SUCCESSFUL then &#x60;originalVideoAssetIds&#x60; will return the original video assetId whereas &#x60;videoAssetIds&#x60; will return translated video assetId. In all other cases, &#x27;originalVideoAssetIds&#x27; and &#x60;videoAssetIds&#x60; both will return original video assetId. | [optional] [default to null]
**Asins** | **[]string** |  | [optional] [default to null]
**BrandLogoUrl** | **string** |  | [optional] [default to null]
**Subpages** | [**[]Subpage**](Subpage.md) |  | [optional] [default to null]
**CreativePropertiesToOptimize** | [**[]CreativePropertyToOptimize**](CreativePropertyToOptimize.md) | If this property is enabled, Sponsored Brands will dynamically optimize by enhancing or generating creative properties based on shopper search intent. | [optional] [default to null]
**OriginalHeadline** | **string** | The original headline submitted by the advertiser. | [optional] [default to null]
**VideoAssetIds** | **[]string** | In SB API V4, &#x60;videoMediaIds&#x60; is replaced by &#x60;videoAssetIds&#x60;. &#x60;videoAssetIds&#x60; will only allow Asset Library identifiers for ad creation, but responses can include mediaIds for v1 campaigns and API V3 operations. At a future state, existing mediaIds will be added to Asset library for use in SB campaigns. | [optional] [default to null]
**BrandLogoAssetID** | **string** |  | [optional] [default to null]
**Headline** | **string** | The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. | [optional] [default to null]
**CreativeStatus** | [***CreativeStatus**](CreativeStatus.md) |  | [optional] [default to null]
**CreativeVersion** | **string** | The version identifier that helps you keep track of multiple versions of a submitted (non-draft) Sponsored Brands creative. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

