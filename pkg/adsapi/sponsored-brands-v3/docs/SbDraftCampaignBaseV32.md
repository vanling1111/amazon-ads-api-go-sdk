# SbDraftCampaignBaseV32

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DraftCampaignId** | **int32** | The identifier of the draft campaign. | [optional] [default to null]
**Name** | **string** | The name of the draft campaign. Maximum 128 characters. Duplicate campaign names are not allowed. | [optional] [default to null]
**Tags** | [***map[string]string**](map.md) |  | [optional] [default to null]
**Budget** | **float64** | The budget associated with the draft campaign. | [optional] [default to null]
**BudgetType** | [***BudgetType**](BudgetType.md) |  | [optional] [default to null]
**StartDate** | **string** | The YYYYMMDD start date for the campaign. If this field is not set to a value, the current date is used. | [optional] [default to null]
**EndDate** | **string** | The YYYYMMDD end date for the campaign. Must be greater than the value for &#x60;startDate&#x60;. If not specified, the campaign has no end date and runs continuously. | [optional] [default to null]
**BrandEntityId** | **string** | The brand entity identifier. Note that this field is required for sellers. For more information, see the [Stores reference](https://advertising.amazon.com/API/docs/v2/reference/stores) or [Brands reference](https://advertising.amazon.com/API/docs/v3/reference/SponsoredBrands/Brands). | [optional] [default to null]
**BidOptimization** | **bool** | Set to &#x60;true&#x60; to allow Amazon to automatically optimize bids for placements below top of search. | [optional] [default to true]
**BidMultiplier** | **float64** | A bid multiplier. Note that this field can only be set when &#x27;bidOptimization&#x27; is set to false. Value is a percentage to two decimal places. Example: If set to -40.00 for a $5.00 bid, the resulting bid is $3.00. | [optional] [default to null]
**BidAdjustments** | [**[]BidAdjustment**](BidAdjustment.md) | List of bid adjustment for each placement group. BidMultiplier cannot be specified when bidAdjustments presents. &#x60;Not supported for video campaigns&#x60; | [optional] [default to null]
**PortfolioId** | **int64** | The identifier of the Portfolio to which the draft campaign is associated. | [optional] [default to null]
**AdFormat** | [***AdFormatV32**](AdFormatV3_2.md) |  | [optional] [default to null]
**Creative** | [***OneOfSbDraftCampaignBaseV32Creative**](OneOfSbDraftCampaignBaseV32Creative.md) |  | [optional] [default to null]
**LandingPage** | [***OneOfSbDraftCampaignBaseV32LandingPage**](OneOfSbDraftCampaignBaseV32LandingPage.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

