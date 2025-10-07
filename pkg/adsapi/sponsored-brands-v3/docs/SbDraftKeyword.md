# SbDraftKeyword

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeywordId** | **int64** | The keyword identifier. | [optional] [default to null]
**AdGroupId** | **int64** | The identifier of the ad group associated with the keyword. | [optional] [default to null]
**CampaignId** | **int64** | The identifier of the campaign associated with the keyword. | [optional] [default to null]
**KeywordText** | **string** | The keyword text. The maximum number of words for this string is 10. | [optional] [default to null]
**MatchType** | [***MatchType**](MatchType.md) |  | [optional] [default to null]
**State** | **string** | | state | description | |-------|-------------| | draft | Newly created keyword. | | enabled | Keyword passed moderation. |  | [optional] [default to null]
**Bid** | **float64** | The bid associated with the keyword. Note that this value must be less than the budget associated with the Advertiser account. For more information, see the **Keyword bid constraints by marketplace** section of the [supported features](https://advertising.amazon.com/API/docs/v2/guides/supported_features) article. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

