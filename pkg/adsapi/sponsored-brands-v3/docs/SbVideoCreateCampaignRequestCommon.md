# SbVideoCreateCampaignRequestCommon

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the campaign. This name must be unique to the Amazon Ads account to which the campaign is associated. Maximum length of the string is 128 characters. | [optional] [default to null]
**Budget** | **float64** | The budget amount associated with the campaign. | [optional] [default to null]
**BudgetType** | [***BudgetType**](BudgetType.md) |  | [optional] [default to null]
**StartDate** | **string** |  | [optional] [default to null]
**EndDate** | **string** |  | [optional] [default to null]
**AdFormat** | [***AdFormat**](AdFormat.md) |  | [optional] [default to null]
**State** | [***State**](State.md) |  | [optional] [default to null]
**BrandEntityId** | **string** | The brand entity identifier. Note that this field is required for sellers. For more information, see the [Stores reference](https://advertising.amazon.com/API/docs/v2/reference/stores) or [Brands reference](https://advertising.amazon.com/API/docs/v3/reference/SponsoredBrands/Brands). | [optional] [default to null]
**PortfolioId** | **int64** | The identifier of the portfolio to which the campaign is associated. | [optional] [default to null]
**Creative** | [***SbVideoCreative**](SBVideoCreative.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

