# SbCampaignResponseCommonV33

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BidOptimizationStrategy** | [***BidOptimizationStrategy**](BidOptimizationStrategy.md) |  | [optional] [default to null]
**BidAdjustments** | [**[]BidAdjustmentV33**](BidAdjustmentV3_3.md) | List of bid adjustments for placement group and shopper segments. BidMultiplier cannot be specified when bidAdjustments are present. &#x60;Not supported for video campaigns&#x60; | [optional] [default to null]
**BidOptimization** | **bool** | Set to &#x60;true&#x60; to allow Amazon to automatically optimize bids for placements below top of search. | [optional] [default to null]
**BidMultiplier** | **float64** | A bid multiplier. Note that this field can only be set when &#x27;bidOptimization&#x27; is set to false. Value is a percentage to two decimal places. For example, if set to -40.00 for a $5.00 bid, the resulting bid is $3.00. | [optional] [default to null]
**Creative** | [***Object**](.md) |  | [optional] [default to null]
**LandingPage** | [***Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

