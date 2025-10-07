# InvoiceLine

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int64** |  | [optional] [default to null]
**CampaignName** | **string** |  | [optional] [default to null]
**CampaignTags** | **map[string]string** | Campaign tags in the form of string key-value pairs. | [optional] [default to null]
**CommissionAmount** | [***CurrencyAmount**](currencyAmount.md) |  | [optional] [default to null]
**CommissionRate** | **float64** |  | [optional] [default to null]
**Cost** | [***CurrencyAmount**](currencyAmount.md) |  | [default to null]
**CostEventCount** | **int64** | Number of clicks/impressions charged | [default to null]
**CostEventType** | **string** | Type of event charged (clicks or impressions) | [default to null]
**CostPerEventType** | **float64** | Ad spends cost (Cost exclusive of adjustments/promotions/fees/etc) per unit (thousand impressions/clicks).  | [optional] [default to null]
**CostPerUnit** | **float64** |  | [default to null]
**Fees** | [**[]Fee**](fee.md) | Charges can include different fees (see feeType below).  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**PortfolioId** | **int64** | Sponsored Ads only. This identifier maps to one of the portfolios listed in the portfolios section.  | [optional] [default to null]
**PriceType** | **string** | Metric used for performance measurement. | [default to null]
**ProgramName** | [***AdProgram**](adProgram.md) |  | [optional] [default to null]
**PromotionAmount** | [***CurrencyAmount**](currencyAmount.md) |  | [optional] [default to null]
**PurchaseOrderNumber** | **string** |  | [optional] [default to null]
**SupplyCost** | [***CurrencyAmount**](currencyAmount.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

