# SbVideoDraftCampaignBase

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DraftCampaignId** | **int32** | The identifier of the draft campaign. | [optional] [default to null]
**Name** | **string** | The name of the draft campaign. Maximum 128 characters. Duplicate campaign names are not allowed. | [optional] [default to null]
**Budget** | **float64** | The budget associated with the draft campaign. | [optional] [default to null]
**BudgetType** | [***BudgetType**](BudgetType.md) |  | [optional] [default to null]
**StartDate** | **string** | The YYYYMMDD start date for the campaign. If this field is not set to a value, the current date is used. | [optional] [default to null]
**EndDate** | **string** | The YYYYMMDD end date for the campaign. Must be greater than the value for &#x60;startDate&#x60;. If not specified, the campaign has no end date and runs continuously. | [optional] [default to null]
**BrandEntityId** | **string** | The brand entity identifier. Note that this field is required for sellers. For more information, see the [Stores reference](https://advertising.amazon.com/API/docs/v2/reference/stores) or [Brands reference](https://advertising.amazon.com/API/docs/v3/reference/SponsoredBrands/Brands). | [optional] [default to null]
**PortfolioId** | **int64** | The identifier of the Portfolio to which the draft campaign is associated. | [optional] [default to null]
**AdFormat** | [***AdFormat**](AdFormat.md) |  | [optional] [default to null]
**Creative** | [***SbVideoCreative**](SBVideoCreative.md) |  | [optional] [default to null]
**LandingPage** | [***SbDetailPageLandingPage**](SBDetailPageLandingPage.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

