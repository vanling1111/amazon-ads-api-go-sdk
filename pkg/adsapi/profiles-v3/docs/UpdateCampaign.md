# UpdateCampaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int64** |  | [default to null]
**Name** | **string** | The name of the campaign. | [optional] [default to null]
**BudgetType** | **string** | The time period over which the amount specified in the &#x60;budget&#x60; property is allocated. | [optional] [default to null]
**Budget** | **float64** | The amount of the budget. | [optional] [default to null]
**StartDate** | **string** | The YYYYMMDD start date of the campaign. The date must be today or in the future. | [optional] [default to null]
**EndDate** | **string** | The YYYYMMDD end date of the campaign. | [optional] [default to null]
**CostType** | **string** | Determines how the campaign will bid and charge. |Name|Description| |----|----------| |cpc |[Default] The performance of this campaign is measured by the clicks triggered by the ad.| |vcpm |The performance of this campaign is measured by the viewed impressions triggered by the ad. |  To view minimum and maximum bids based on the costType, see [Limits](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace). | [optional] [default to null]
**State** | **string** | The state of the campaign. | [optional] [default to null]
**PortfolioId** | **int64** | Identifier of the portfolio that will be associated with the campaign. If null then the campaign will be disassociated from existing portfolio. Campaigns with CPC and vCPM costType are supported. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

