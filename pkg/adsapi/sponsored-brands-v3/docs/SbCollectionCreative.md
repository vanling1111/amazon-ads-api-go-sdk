# SbCollectionCreative

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandName** | **string** | The brand name. | [optional] [default to null]
**BrandLogoAssetId** | **string** | The identifier of the brand logo image from the Store assets library. See [listAssets](https://advertising.amazon.com/API/docs/v3/reference/SponsoredBrands/Stores) for more information. Note that for campaigns created in the Amazon Ads console prior to release of the Store assets library, responses will not include a value for the brandLogoAssetID field. | [optional] [default to null]
**BrandLogoUrl** | **string** | The address of the hosted image. | [optional] [default to null]
**Headline** | **string** | The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters. | [optional] [default to null]
**Asins** | **[]string** | An array of ASINs associated with the creative. **Note** do not pass an empty array. This results in an error. | [optional] [default to null]
**ShouldOptimizeAsins** | **bool** | Note that this field is supported only in the US and UK marketplaces. Set to &#x60;true&#x60; to have Amazon show other products from your landing page in the advertisement if they are more relevant to the shopper&#x27;s search. Set to &#x60;false&#x60; to use the ASINs specified in the &#x60;asins&#x60; field. Do not specify in unsupported marketplaces. | [optional] [default to null]
**BrandLogoAssetID** | **string** | The identifier of the brand logo image from the Store assets library. See [listAssets](https://advertising.amazon.com/API/docs/v3/reference/SponsoredBrands/Stores) for more information. Note that for campaigns created in the Amazon Ads console prior to release of the Store assets library, responses will not include a value for the brandLogoAssetID field. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

