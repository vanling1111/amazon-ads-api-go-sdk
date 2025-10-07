# CreateAd

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdGroupId** | **string** | The identifier of the Ad Group associated with the Ad. | [default to null]
**AdName** | **string** | The name of the Ad. | [optional] [default to null]
**FullFunnelCampaignId** | **string** | full funnel campaign id for child ad | [optional] [default to null]
**LandingPageType** | **string** | The type of landing page for the Ad. You can specify one of these values: | Type | Description | | --- | ------- | | ASIN_DP | The detail page of the asin. Please specify the ASIN in the &#x60;landingPageValue&#x60; field | | SKU_DP | The detail page of the sku. Please specify the SKU in the &#x60;landingPageValue&#x60; field | | OFF_AMAZON_LINK | Landing page URL. Please specify an HTTPS URL value in the &#x60;landingPageValue&#x60; field | | STORE | The Amazon brand store URL. Please specify an HTTPS URL value in the &#x60;landingPageValue&#x60; field | | [optional] [default to null]
**LandingPageValue** | **string** | The landing page for the Ad. | [optional] [default to null]
**State** | [***CreateOrUpdateEntityState**](CreateOrUpdateEntityState.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

