# BidAdjustmentV33

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BidAdjustmentPredicate** | **string** | Determines the predicate (placement groups and shopper segments) that the bid adjustment will be made for. |BidAdjustmentPredicate|Description| |------|-----------| |PLACEMENT_GROUP_DETAIL_PAGE|Predicate for adjusting bids at detail page placement.| |PLACEMENT_GROUP_HOME|Predicate for adjusting bids at home page placement.| |PLACEMENT_GROUP_OTHER|Predicate for adjusting bids at pages other than detail and home pages placement.| |SHOPPER_SEGMENT_NEW_TO_BRAND_PURCHASE|Predicate for adjusting bids for new-to-brand purchase shopper segment.| | [optional] [default to null]
**BidAdjustmentPercent** | **float64** | Bid adjustment for placement groups and shopper segments. Value is a percentage to two decimal places. For bid adjustments in placement groups, min is -99.00 and max is 900.00. For bid adjustments in shopper segments, min is 0.00 and max is 900.00. For example: If -40.00 is set for a $5.00 bid, the resulting bid is $3.00. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

