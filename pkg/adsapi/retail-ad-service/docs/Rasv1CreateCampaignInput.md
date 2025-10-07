# Rasv1CreateCampaignInput

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | [***Rasv1BudgetInput**](RASv1BudgetInput.md) |  | [default to null]
**DynamicBidding** | [***Rasv1RequiredDynamicBidding**](RASv1RequiredDynamicBidding.md) |  | [default to null]
**EndDate** | **string** | Campaign end date | [optional] [default to null]
**Name** | **string** | The name of the campaign. | [default to null]
**PortfolioId** | **string** | The identifier of an existing portfolio to which the campaign is associated. | [optional] [default to null]
**StartDate** | **string** | Campaign start date. Default: today&amp;#39;s date. The format of the date is YYYY-MM-DD. | [optional] [default to null]
**State** | [***Rasv1EntityState**](RASv1EntityState.md) |  | [default to null]
**Tags** | [**[]Rasv1Tag**](RASv1Tag.md) | A list of advertiser-specified custom identifiers for the campaign. Each customer identifier is a key-value pair. You can specify a maximum of 50 identifiers. | [optional] [default to null]
**TargetingType** | [***Rasv1TargetingType**](RASv1TargetingType.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

