# SbKeywordsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeywordId** | **int64** | The identifier of the keyword. | [default to null]
**AdGroupId** | **int64** | The identifier of an existing ad group to which the keyword is associated. | [default to null]
**CampaignId** | **int64** | The identifier of an existing campaign to which the keyword is associated. | [default to null]
**State** | [***SbKeywordState**](SBKeywordState.md) |  | [optional] [default to null]
**Bid** | **float64** | The bid associated with the keyword. Note that this value must be less than the budget associated with the Advertiser account. For more information, see the **Keyword bid constraints by marketplace** section of the [supported features](https://advertising.amazon.com/API/docs/v2/guides/supported_features) article. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

