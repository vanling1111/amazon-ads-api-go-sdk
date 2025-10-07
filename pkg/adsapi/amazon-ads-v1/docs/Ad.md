# Ad

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveCreative** | [***Creative**](Creative.md) |  | [optional] [default to null]
**AdGroupId** | **string** | The ad group associated with the ad. | [optional] [default to null]
**AdId** | **string** | The identifier of the ad. | [default to null]
**AdProduct** | [***AdProduct**](AdProduct.md) |  | [default to null]
**AdType** | [***AdType**](AdType.md) |  | [default to null]
**CampaignId** | **string** | The campaign associated with the ad. It&#x27;s a read-only field. | [optional] [default to null]
**CreationDateTime** | [**time.Time**](time.Time.md) | The date time that the ad was created. | [default to null]
**Creative** | [***Creative**](Creative.md) |  | [default to null]
**LastUpdatedDateTime** | [**time.Time**](time.Time.md) | The date time that the ad was last updated. | [default to null]
**MarketplaceConfigurations** | [**[]MarketplaceAdConfigurations**](MarketplaceAdConfigurations.md) | List of marketplace-specific configurations for a global ad that enables overriding certain attributes at individual marketplace level. For example, if a global ad is ENABLED but needs to be PAUSED in DE marketplace, you can specify: [{marketplace: DE, overrides: {state: PAUSED}}]. When a marketplace-specific override is not provided, the ad&#x27;s global value is applied to that marketplace. | [optional] [default to null]
**MarketplaceScope** | [***MarketplaceScope**](MarketplaceScope.md) |  | [optional] [default to null]
**Marketplaces** | [**[]Marketplace**](Marketplace.md) | A list of country codes representing Amazon marketplaces | Marketplace | Description | | --- | --- | | &#x60;AE&#x60; |  | | &#x60;AU&#x60; |  | | &#x60;BE&#x60; |  | | &#x60;BR&#x60; |  | | &#x60;CA&#x60; |  | | &#x60;DE&#x60; |  | | &#x60;EG&#x60; |  | | &#x60;ES&#x60; |  | | &#x60;FR&#x60; |  | | &#x60;GB&#x60; |  | | &#x60;IN&#x60; |  | | &#x60;IT&#x60; |  | | &#x60;JP&#x60; |  | | &#x60;MX&#x60; |  | | &#x60;NL&#x60; |  | | &#x60;PL&#x60; |  | | &#x60;SA&#x60; |  | | &#x60;SE&#x60; |  | | &#x60;SG&#x60; |  | | &#x60;TR&#x60; |  | | &#x60;US&#x60; |  | | [optional] [default to null]
**Name** | **string** | The name of the ad. | [optional] [default to null]
**State** | [***State**](State.md) |  | [default to null]
**Status** | [***Status**](Status.md) |  | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | Open ended labels with a key value pair applied to the ad | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

