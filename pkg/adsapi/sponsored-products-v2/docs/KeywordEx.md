# KeywordEx

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeywordId** | **float64** | The identifier of the keyword. | [optional] [default to null]
**CampaignId** | **float64** | The identifer of the campaign to which the keyword is associated. | [optional] [default to null]
**AdGroupId** | **float64** | The identifier of the ad group to which this keyword is associated. | [optional] [default to null]
**State** | [***State**](State.md) |  | [optional] [default to null]
**KeywordText** | **string** | The text of the expression to match against a search query. | [optional] [default to null]
**NativeLanguageKeyword** | **string** | The unlocalized keyword text in the preferred locale of the advertiser. | [optional] [default to null]
**MatchType** | [***MatchType**](MatchType.md) |  | [optional] [default to null]
**Bid** | **float32** | Bid associated with this keyword. This table details the maximum allowable bid (in local currency) for keywords by marketplace: | Marketplace | Currency | Min / Max bid for SP | | --- | --- | --- | | US | USD | 0.02 / 1000 | | CA | CAD | 0.02 / 1000 | | UK | GBP | 0.02 / 1000 | | DE | EUR | 0.02 / 1000 | | FR | EUR | 0.02 / 1000 | | ES | EUR | 0.02 / 1000 | | IT | EUR | 0.02 / 1000 | | JP | JPY | 2.0 / 100000 | | AU | AUD | 0.10 / 1410 | | AE | AED | 0.24 / 184.0 | | [optional] [default to null]
**CreationDate** | **float64** | Creation date in epoch time. | [optional] [default to null]
**LastUpdatedDate** | **float64** | Date of last update in epoch time. | [optional] [default to null]
**ServingStatus** | **string** | The serving status of the keyword. See the **computed status** section of the [developer notes](https://advertising.amazon.com/API/docs/en-us/reference/concepts/developer-notes) for definitions. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

