# AdGroupUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdGroupId** | **string** | The unique identifier of the ad group. | [default to null]
**AdProduct** | [***AdProduct**](AdProduct.md) |  | [optional] [default to null]
**AdvertisedProductCategoryIds** | **[]string** | The array of identifiers of product categories associated with the ad group. For VIDEO ad group type only one parent product category or multiple sub-categories from one parent product category are allowed. | [optional] [default to null]
**Bid** | [***UpdateAdGroupBid**](UpdateAdGroupBid.md) |  | [optional] [default to null]
**Budgets** | [**[]CreateBudget**](CreateBudget.md) | An object containing budget details for the ad group. | [optional] [default to null]
**CreativeRotationType** | [***CreativeRotationType**](CreativeRotationType.md) |  | [optional] [default to null]
**EndDateTime** | [**time.Time**](time.Time.md) | The end date time for the ad group. | [optional] [default to null]
**Fees** | [**[]CreateFee**](CreateFee.md) | The fees associated with the ad group. | [optional] [default to null]
**Frequencies** | [**[]CreateFrequency**](CreateFrequency.md) | An object containing frequency details for the ad group. | [optional] [default to null]
**MarketplaceConfigurations** | [**[]CreateMarketplaceAdGroupConfigurations**](CreateMarketplaceAdGroupConfigurations.md) | List of marketplace-specific configurations for a global ad group that enables overriding certain attributes at individual marketplace level. For example, if a global ad group state is ENABLED and needs to be PAUSED only in DE marketplace, you can specify: [{marketplace: DE, overrides: {state: PAUSED}}]. When a marketplace-specific override is not provided, ad group&#x27;s global value is applied to that marketplace. | [optional] [default to null]
**MarketplaceScope** | [***MarketplaceScope**](MarketplaceScope.md) |  | [optional] [default to null]
**Marketplaces** | [**[]Marketplace**](Marketplace.md) | A list of country codes representing Amazon marketplaces | Marketplace | Description | | --- | --- | | &#x60;AE&#x60; |  | | &#x60;AU&#x60; |  | | &#x60;BE&#x60; |  | | &#x60;BR&#x60; |  | | &#x60;CA&#x60; |  | | &#x60;DE&#x60; |  | | &#x60;EG&#x60; |  | | &#x60;ES&#x60; |  | | &#x60;FR&#x60; |  | | &#x60;GB&#x60; |  | | &#x60;IN&#x60; |  | | &#x60;IT&#x60; |  | | &#x60;JP&#x60; |  | | &#x60;MX&#x60; |  | | &#x60;NL&#x60; |  | | &#x60;PL&#x60; |  | | &#x60;SA&#x60; |  | | &#x60;SE&#x60; |  | | &#x60;SG&#x60; |  | | &#x60;TR&#x60; |  | | &#x60;US&#x60; |  | | [optional] [default to null]
**Name** | **string** | The name of the ad group. | [optional] [default to null]
**Optimization** | [***UpdateOptimization**](UpdateOptimization.md) |  | [optional] [default to null]
**Pacing** | [***UpdatePacing**](UpdatePacing.md) |  | [optional] [default to null]
**PurchaseOrderNumber** | **string** | The purchase order number associated with the ad group. | [optional] [default to null]
**StartDateTime** | [**time.Time**](time.Time.md) | The start date time for the ad group. | [optional] [default to null]
**State** | [***UpdateState**](UpdateState.md) |  | [optional] [default to null]
**Tags** | [**[]CreateTag**](CreateTag.md) | Open ended labels with a key value pair applied to the ad group | [optional] [default to null]
**TargetingSettings** | [***UpdateTargetingSettings**](UpdateTargetingSettings.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

