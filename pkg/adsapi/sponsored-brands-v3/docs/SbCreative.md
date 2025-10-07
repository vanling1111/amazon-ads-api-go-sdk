# SbCreative

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandName** | **string** | A brand name. Maximum length is 30 characters. | [optional] [default to null]
**BrandLogoAssetID** | **string** | The identifier of the brand logo image from the Store assets library. See [listAssets](https://advertising.amazon.com/API/docs/v3/reference/SponsoredBrands/Stores) for more information. Note that for campaigns created in the Amazon Ads console prior to release of the Store assets library, responses will not include a value for the brandLogoAssetID field. | [optional] [default to null]
**BrandLogoUrl** | **string** | The address of the hosted image. | [optional] [default to null]
**Headline** | **string** | The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. | [optional] [default to null]
**Asins** | **[]string** | An array of ASINs associated with the creative. **Note** do not pass an empty array, this results in an error. | [optional] [default to null]
**ShouldOptimizeAsins** | **bool** | **NOTE** Starting on March 25th, 2021, this property will no longer be supported. This feature is currently available in the US and UK. Existing Sponsored Brands campaigns with product optimization enabled will no longer have the products in the creative automatically optimized. Campaigns with product optimization enabled will be converted to standard Sponsored Brands product collection campaigns with the default selected products showing in the creative. For POST and PUT operations, setting this property to &#x60;true&#x60; will not have any effect. The value returned in the response will always be &#x60;false&#x60;. For the GET operation, the value of this field will always be &#x60;false&#x60;. And starting on September 25th, 2021, this property will be removed completely.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

