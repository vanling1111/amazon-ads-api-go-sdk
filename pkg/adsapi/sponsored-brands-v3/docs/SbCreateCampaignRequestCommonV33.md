# SbCreateCampaignRequestCommonV33

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BidOptimizationStrategy** | [***BidOptimizationStrategy**](BidOptimizationStrategy.md) |  | [optional] [default to null]
**BidAdjustments** | [**[]BidAdjustmentV33**](BidAdjustmentV3_3.md) | List of bid adjustments for placement group and shopper segments. BidMultiplier cannot be specified when bidAdjustments are present. &#x60;Not supported for video campaigns&#x60; | [optional] [default to null]
**Name** | **string** | The name of the campaign. This name must be unique to the Amazon Ads account to which the campaign is associated. Maximum length of the string is 128 characters. Note that idempotency for this field works different for sellers and vendors. Sellers aren&#x27;t allowed to have duplicate campaign names, but vendors can have duplicate campaign names. | [optional] [default to null]
**Tags** | [***map[string]string**](map.md) |  | [optional] [default to null]
**Budget** | **float64** | The budget amount associated with the campaign. | [optional] [default to null]
**BudgetType** | [***BudgetType**](BudgetType.md) |  | [optional] [default to null]
**StartDate** | **string** |  | [optional] [default to null]
**EndDate** | **string** |  | [optional] [default to null]
**AdFormat** | [***AdFormatV32**](AdFormatV3_2.md) |  | [optional] [default to null]
**State** | [***State**](State.md) |  | [optional] [default to null]
**BrandEntityId** | **string** | The brand entity identifier. Note that this field is required for sellers. For more information, see the [Stores reference](https://advertising.amazon.com/API/docs/v2/reference/stores) or [Brands reference](https://advertising.amazon.com/API/docs/v3/reference/SponsoredBrands/Brands). | [optional] [default to null]
**BidOptimization** | **bool** | Set to &#x60;true&#x60; to allow Amazon to automatically optimize bids for placements below top of search. | [optional] [default to true]
**BidMultiplier** | **float64** | A bid multiplier. Note that this field can only be set when &#x27;bidOptimization&#x27; is set to false. Value is a percentage to two decimal places. Example: If set to -40.00 for a $5.00 bid, the resulting bid is $3.00. | [optional] [default to null]
**PortfolioId** | **int64** | The identifier of the portfolio to which the campaign is associated. | [optional] [default to null]
**Creative** | [***Object**](.md) |  | [optional] [default to null]
**LandingPage** | [***SbLandingPage**](SBLandingPage.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

