# SbProductCollectionCreative

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandLogoCrop** | [***SbBrandLogoCrop**](SBBrandLogoCrop.md) |  | [optional] [default to null]
**Asins** | **[]string** |  | [optional] [default to null]
**BrandLogoUrl** | **string** | The address of the hosted image. | [optional] [default to null]
**BrandName** | **string** | A brand name. Maximum length is 30 characters. | [optional] [default to null]
**CustomImageAssetId** | **string** | The identifier of the Custom image from the Store assets library. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#customimage) for more information on what constitutes a valid Custom image. | [optional] [default to null]
**ShouldOptimizeAsins** | **bool** | Starting on March 25th, 2021, this property will no longer be supported. This feature is currently available in the US and UK. Existing Sponsored Brands campaigns with product optimization enabled will no longer have the products in the creative automatically optimized. Campaigns with product optimization enabled will be converted to standard Sponsored Brands product collection campaigns with the default selected products showing in the creative. For POST and PUT operations, setting this property to true will not have any effect. The value returned in the response will always be false. For the GET operation, the value of this field will always be false. And starting on September 25th, 2021, this property will be removed completely. | [optional] [default to null]
**CustomImageCrop** | [***SbCustomImageCrop**](SBCustomImageCrop.md) |  | [optional] [default to null]
**CustomImageUrl** | **string** | The address of the hosted Custom image. | [optional] [default to null]
**BrandLogoAssetID** | **string** | The identifier of the brand logo image from the Store assets library. See [listAssets](https://advertising.amazon.com/API/docs/v3/reference/SponsoredBrands/Stores) for more information. Note that for campaigns created in the Amazon Advertising console prior to release of the Store assets library, responses will not include a value for the brandLogoAssetID field. | [optional] [default to null]
**Headline** | **string** | The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

