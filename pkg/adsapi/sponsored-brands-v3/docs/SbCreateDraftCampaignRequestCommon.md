# SbCreateDraftCampaignRequestCommon

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the draft campaign. Maximum 128 characters. Names must be unique to the Amazon Ads account to which they are associated. | [default to null]
**Budget** | **float64** | The budget associated with the draft campaign. | [default to null]
**BudgetType** | [***BudgetType**](BudgetType.md) |  | [default to null]
**StartDate** | **string** | The YYYYMMDD start date of the campaign. Must be equal to or greater than the current date. If not specified, is set to current date by default. | [optional] [default to null]
**EndDate** | **string** | The YYYYMMDD end date of the campaign. Must be greater than the value specified in the &#x60;startDate&#x60; field. If not specified, the campaign has no end date and runs continuously. | [optional] [default to null]
**BrandEntityId** | **string** | The brand entity identifier to which the draft campaign is associated. Note that this field is required for sellers. Retrieve using the getBrands or getStores operations in the /v2/stores resource. | [optional] [default to null]
**BidOptimization** | **bool** | Set to &#x60;true&#x60; to have Amazon automatically optimize bids for placements below top of search. | [optional] [default to true]
**BidMultiplier** | **float64** | A bid multiplier. Note that this field can only be set when &#x27;bidOptimization&#x27; is set to false. Value is a percentage to two decimal places. Example: If set to -40.00 for a $5.00 bid, the resulting bid is $3.00. | [optional] [default to null]
**BidAdjustments** | [**[]BidAdjustment**](BidAdjustment.md) | List of bid adjustment for each placement group. BidMultiplier cannot be specified when bidAdjustments presents. &#x60;Not supported for video campaigns&#x60; | [optional] [default to null]
**PortfolioId** | **int64** | The identifier of the Portfolio to which the draft campaign is associated. | [optional] [default to null]
**Creative** | [***SbCreative**](SBCreative.md) |  | [optional] [default to null]
**LandingPage** | [***SbLandingPage**](SBLandingPage.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

