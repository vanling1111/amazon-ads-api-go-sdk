# CreateCampaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortfolioId** | **float64** | The identifier of an existing portfolio to which the campaign is associated. | [optional] [default to null]
**Name** | **string** | A name for the campaign. Note that idempotency for this field works different for sellers and vendors. Sellers aren&#x27;t allowed to have duplicate campaign names, but vendors can have duplicate campaign names. | [optional] [default to null]
**Tags** | [***map[string]string**](map.md) |  | [optional] [default to null]
**CampaignType** | **string** | The advertising product managed by this campaign. | [optional] [default to null]
**TargetingType** | **string** | The type of targeting for the campaign. | [optional] [default to null]
**State** | [***State**](State.md) |  | [optional] [default to null]
**DailyBudget** | **float32** | A daily budget for the campaign. | [optional] [default to null]
**StartDate** | **string** | A starting date for the campaign to go live. The format of the date is YYYYMMDD. | [optional] [default to null]
**EndDate** | **string** | An ending date for the campaign to stop running. The format of the date is YYYYMMDD. | [optional] [default to null]
**PremiumBidAdjustment** | **bool** | If set to true, Amazon increases the default bid for ads that are eligible to appear in this placement. See developer notes for more information. | [optional] [default to null]
**Bidding** | [***Bidding**](Bidding.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

