# CampaignEx

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortfolioId** | **float64** | The identifier of an existing portfolio to which the campaign is associated. | [optional] [default to null]
**CampaignId** | **float64** | The identifier of the campaign. | [optional] [default to null]
**Tags** | [***map[string]string**](map.md) |  | [optional] [default to null]
**Name** | **string** | The name of the campaign. | [optional] [default to null]
**CampaignType** | **string** | The advertising product managed by this campaign. | [optional] [default to null]
**TargetingType** | **string** | The type of targeting of the campaign. | [optional] [default to null]
**State** | [***State**](State.md) |  | [optional] [default to null]
**DailyBudget** | **float32** | The daily budget of the campaign. | [optional] [default to null]
**StartDate** | **string** | The starting date of the campaign. The format of the date is YYYYMMDD. | [optional] [default to null]
**EndDate** | **string** | The ending date of the campaign to stop running. The format of the date is YYYYMMDD. | [optional] [default to null]
**PremiumBidAdjustment** | **bool** | If set to true, Amazon increases the default bid for ads that are eligible to appear in this placement. See developer notes for more information. | [optional] [default to null]
**Bidding** | [***Bidding**](Bidding.md) |  | [optional] [default to null]
**CreationDate** | **float64** | The creation date of the campaign, in epoch time. | [optional] [default to null]
**LastUpdatedDate** | **float64** | The date that any value associated with the campaign was last changed, in epoch time. | [optional] [default to null]
**ServingStatus** | **string** | The computed status of the campaign. See developer notes for more information. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

