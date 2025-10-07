# TargetUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bid** | [***UpdateTargetBid**](UpdateTargetBid.md) |  | [optional] [default to null]
**CampaignId** | **string** | A unique identifier for the campaign associated with the target. Only used for campaign-level targets. | [optional] [default to null]
**MarketplaceConfigurations** | [**[]CreateMarketplaceTargetConfigurations**](CreateMarketplaceTargetConfigurations.md) | List of marketplace-specific configurations for a global target that enables overriding certain attributes at individual marketplace level. For example, if a global target is ENABLED but needs to be PAUSED in DE marketplace, you can specify: [{marketplace: DE, overrides: {state: PAUSED}}]. When a marketplace-specific override is not provided, the target&#x27;s global value is applied to that marketplace. | [optional] [default to null]
**MarketplaceScope** | [***MarketplaceScope**](MarketplaceScope.md) |  | [optional] [default to null]
**Marketplaces** | [**[]Marketplace**](Marketplace.md) | A list of country codes representing Amazon marketplaces | Marketplace | Description | | --- | --- | | &#x60;AE&#x60; |  | | &#x60;AU&#x60; |  | | &#x60;BE&#x60; |  | | &#x60;BR&#x60; |  | | &#x60;CA&#x60; |  | | &#x60;DE&#x60; |  | | &#x60;EG&#x60; |  | | &#x60;ES&#x60; |  | | &#x60;FR&#x60; |  | | &#x60;GB&#x60; |  | | &#x60;IN&#x60; |  | | &#x60;IT&#x60; |  | | &#x60;JP&#x60; |  | | &#x60;MX&#x60; |  | | &#x60;NL&#x60; |  | | &#x60;PL&#x60; |  | | &#x60;SA&#x60; |  | | &#x60;SE&#x60; |  | | &#x60;SG&#x60; |  | | &#x60;TR&#x60; |  | | &#x60;US&#x60; |  | | [optional] [default to null]
**State** | [***UpdateState**](UpdateState.md) |  | [optional] [default to null]
**Tags** | [**[]CreateTag**](CreateTag.md) | Open ended labels with a key value pair applied to the target | [optional] [default to null]
**TargetId** | **string** | A unique identifier for the target. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

