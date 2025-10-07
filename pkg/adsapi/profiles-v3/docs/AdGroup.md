# AdGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdGroupId** | **int64** |  | [optional] [default to null]
**Tactic** | [***Tactic**](Tactic.md) |  | [optional] [default to null]
**CreativeType** | [***CreativeType**](CreativeType.md) |  | [optional] [default to null]
**Name** | **string** | The name of the ad group. | [optional] [default to null]
**CampaignId** | **int64** |  | [optional] [default to null]
**DefaultBid** | **float64** | The amount of the default bid associated with the ad group. Used if no bid is specified. | [optional] [default to null]
**BidOptimization** | **string** | Bid Optimization for the Adgroup. Default behavior is to optimize for clicks. |Name|CostType|Description| |----|--------|-----------| |reach |vcpm|Optimize for viewable impressions. $1 is the minimum bid for vCPM.| |clicks |cpc|[Default] Optimize for page visits.| |conversions |cpc|Optimize for conversion.| | [optional] [default to null]
**State** | **string** | The state of the ad group. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

