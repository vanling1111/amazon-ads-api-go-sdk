# ProductMetadataModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asin** | **string** | ASIN of the item | [optional] [default to null]
**Availability** | **string** | Stock availability:   * IN_STOCK - The item is in stock.   * IN_STOCK_SCARCE - The item is in stock, but stock levels are limited.   * OUT_OF_STOCK - The item is currently out of stock.   * PREORDER - The item is not yet available, but can be pre-ordered.   * LEADTIME - The item is only available after some amount of lead time.   * AVAILABLE_DATE - The item is not available, but will be available on a future date. | [optional] [default to null]
**BasisPrice** | [***BasisPrice**](basisPrice.md) |  | [optional] [default to null]
**BestSellerRank** | **string** | Best seller rank position in the category | [optional] [default to null]
**Brand** | **string** | Brand name of the item | [optional] [default to null]
**Category** | **string** | Category (browse node) name of the ASIN | [optional] [default to null]
**CreatedDate** | **string** | Date the item was first available on Amazon | [optional] [default to null]
**EligibilityStatus** | **string** | Eligibility status for advertising:   * ELIGIBLE - Eligible for advertising   * INELIGIBLE - Ineligible for advertising | [optional] [default to null]
**GlobalStoreSetting** | [***ProductMetadataModelGlobalStoreSetting**](ProductMetadataModel_globalStoreSetting.md) |  | [optional] [default to null]
**ImageUrl** | **string** | Url to the product image | [optional] [default to null]
**IneligibilityCodes** | **[]string** | List of ineligible status identifier | [optional] [default to null]
**IneligibilityReasons** | **[]string** | List of reasons that made this item ineligible to be advertised | [optional] [default to null]
**PriceToPay** | [***PriceToPay**](priceToPay.md) |  | [optional] [default to null]
**Sku** | **string** | sku of the item | [optional] [default to null]
**Title** | **string** | Product title of the item | [optional] [default to null]
**VariationList** | **[]string** | List of ASIN variations of the current item | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

