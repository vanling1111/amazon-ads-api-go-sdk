# CampaignCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdProduct** | [***AdProduct**](AdProduct.md) |  | [default to null]
**AutoCreationSettings** | [***CreateAutoCreationSettings**](CreateAutoCreationSettings.md) |  | [optional] [default to null]
**BrandId** | **string** | This is the ID of the brand that the campaign is associated with. | [optional] [default to null]
**Budgets** | [**[]CreateBudget**](CreateBudget.md) | The object containing budget details for the campaign (for campaigns that support multiple budgets). | [optional] [default to null]
**CostType** | [***CostType**](CostType.md) |  | [optional] [default to null]
**Countries** | [**[]CountryCode**](CountryCode.md) | This field is used in Sponsored Ads and ADSP and impacts targeted supply. For Sponsored Ads, the campaign.countries field determines what Amazon retail supply (Amazon.com, Amazon.co.uk, Amazon.mx, etc) the campaign will serve in. Similarly in ADSP, this has an implicit filter on your inventory targets. If you choose an inventory target of AMAZON with campaign.countries set to US, this will target the retail supply of Amazon.com and non-retail Amazon properties. ADSP options include additional countries - for example, choosing Austria means targeting Austria eligible inventory and Amazon retail supply of Amazon.de. | [optional] [default to null]
**EndDateTime** | [**time.Time**](time.Time.md) | The end date time for the campaign. | [optional] [default to null]
**Fees** | [**[]CreateCampaignFee**](CreateCampaignFee.md) | Any fees associated with the campaign. | [optional] [default to null]
**Flights** | [**[]CreateCampaignFlight**](CreateCampaignFlight.md) | Flight details associated with the campaign. | [optional] [default to null]
**Frequencies** | [**[]CreateFrequency**](CreateFrequency.md) | Any frequency caps associated with the campaign. | [optional] [default to null]
**MarketplaceConfigurations** | [**[]CreateMarketplaceCampaignConfigurations**](CreateMarketplaceCampaignConfigurations.md) | List of marketplace-specific configurations for a global campaign that enables overriding certain attributes at individual marketplace level. For example, if a global campaign is ENABLED and startDate &#x27;2024-06-01&#x27; but needs to be PAUSED in DE with startDateTime &#x27;2024-06-02&#x27; marketplace, you can specify: [{marketplace: DE, overrides: {state: PAUSED, startDate: &#x27;2024-06-02&#x27;}}]. When a marketplace-specific override is not provided, the campaign&#x27;s global value is applied to that marketplace. | [optional] [default to null]
**MarketplaceScope** | [***MarketplaceScope**](MarketplaceScope.md) |  | [optional] [default to null]
**Marketplaces** | [**[]Marketplace**](Marketplace.md) | A list of country codes representing Amazon marketplaces | Marketplace | Description | | --- | --- | | &#x60;AE&#x60; |  | | &#x60;AU&#x60; |  | | &#x60;BE&#x60; |  | | &#x60;BR&#x60; |  | | &#x60;CA&#x60; |  | | &#x60;DE&#x60; |  | | &#x60;EG&#x60; |  | | &#x60;ES&#x60; |  | | &#x60;FR&#x60; |  | | &#x60;GB&#x60; |  | | &#x60;IN&#x60; |  | | &#x60;IT&#x60; |  | | &#x60;JP&#x60; |  | | &#x60;MX&#x60; |  | | &#x60;NL&#x60; |  | | &#x60;PL&#x60; |  | | &#x60;SA&#x60; |  | | &#x60;SE&#x60; |  | | &#x60;SG&#x60; |  | | &#x60;TR&#x60; |  | | &#x60;US&#x60; |  | | [optional] [default to null]
**Name** | **string** | The name of the campaign. | [default to null]
**Optimizations** | [***CreateCampaignOptimizations**](CreateCampaignOptimizations.md) |  | [optional] [default to null]
**PortfolioId** | **string** | The ID of the portfolio associated with the campaign. | [optional] [default to null]
**ProductCategoryId** | **string** | This is the ID of the product category that the campaign is associated with. | [optional] [default to null]
**PurchaseOrderNumber** | **string** | The purchase order number associated with the campaign. | [optional] [default to null]
**SalesChannel** | [***SalesChannel**](SalesChannel.md) |  | [optional] [default to null]
**SiteRestrictions** | [**[]SiteRestriction**](SiteRestriction.md) | Restrict the ad to a particular site | [optional] [default to null]
**SkanAppId** | **string** | StoreKit AdNetwork application ID. Represents iTunes application ID with which SKAN-enabled campaigns are associated. | [optional] [default to null]
**StartDateTime** | [**time.Time**](time.Time.md) | The start date time for the campaign. | [optional] [default to null]
**State** | [***CreateState**](CreateState.md) |  | [default to null]
**Tags** | [**[]CreateTag**](CreateTag.md) | Open ended labels with a key value pair applied to the campaign | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

