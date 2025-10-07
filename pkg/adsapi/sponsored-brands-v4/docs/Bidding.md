# Bidding

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BidOptimization** | **bool** | Whether to use automatic placement level bid optimization. If set to true, Amazon will automatically set the right placement adjustment and the bidAdjustmentsByPlacement field is ignored. If set to false, the bidAdjustmentsByPlacement field will be used to adjust bid on different placements. If this field is changed from false to true, the bidAdjustmentsByPlacement field will be reset to null. | [optional] [default to null]
**ShopperCohortBidAdjustments** | [**[]ShopperCohortBidAdjustment**](ShopperCohortBidAdjustment.md) | Shopper cohort based bid adjustments. | [optional] [default to null]
**BidAdjustmentsByPlacement** | [**[]BidAdjustmentByPlacement**](BidAdjustmentByPlacement.md) | Placement level bid adjustment. Note that this field can only be set when &#x27;bidOptimization&#x27; is set to false. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

