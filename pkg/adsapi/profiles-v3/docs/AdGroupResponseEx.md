# AdGroupResponseEx

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdGroupId** | **float64** | The identifier of the ad group. | [optional] [default to null]
**Name** | **string** | The name of the ad group. | [optional] [default to null]
**CampaignId** | **float64** | The identifier of the campaign that this ad group is associated with. | [optional] [default to null]
**DefaultBid** | **float64** | The amount of the default bid associated with the ad group. Used if no bid is specified. | [optional] [default to null]
**State** | **string** | The delivery state of the ad group. | [optional] [default to null]
**Tactic** | [***Tactic**](Tactic.md) |  | [optional] [default to null]
**CreativeType** | [***CreativeTypeInCreativeResponse**](CreativeTypeInCreativeResponse.md) |  | [optional] [default to null]
**ServingStatus** | **string** | The status of the ad group. | [optional] [default to null]
**BidOptimization** | **string** | Bid optimization type for the Adgroup. Default behavior is to optimize for clicks. Note, reach, clicks are only accepted with productAds that include landingPageURL OFF_AMAZON_LINK. |Name|CostType|Description| |----|--------|-----------| |reach|vcpm|Optimize for viewable impressions. $1 is the minimum bid for vCPM.| |clicks [Default]|cpc|Optimize for page visits.| |conversions|cpc|Optimize for conversion.| | [optional] [default to null]
**CreationDate** | **int64** | Epoch time the ad group was created. | [optional] [default to null]
**LastUpdatedDate** | **int64** | Epoch time any property in the ad group was last updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

