# ProductCollectionCreative

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asins** | **[]string** | An array of ASINs associated with the creative. | [default to null]
**BrandLogoCrop** | [***AssetCrop**](AssetCrop.md) |  | [optional] [default to null]
**BrandName** | **string** | The displayed brand name in the ad headline. Maximum length is 30 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements. | [default to null]
**CustomImageAssetId** | **string** | The identifier of the Custom image from the Store assets library. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#customimage) for more information on what constitutes a valid Custom image. | [optional] [default to null]
**CustomImageCrop** | [***AssetCrop**](AssetCrop.md) |  | [optional] [default to null]
**BrandLogoAssetId** | **string** | The identifier of the [brand logo](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#brandlogo) image from the brand store&#x27;s asset library. Note that for campaigns created in the Amazon Advertising console prior to release of the brand store&#x27;s assets library, responses will not include a value for this field. | [default to null]
**Headline** | **string** | The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. See [the policy](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies#headlines) for headline requirements. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

