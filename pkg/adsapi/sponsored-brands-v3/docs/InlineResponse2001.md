# InlineResponse2001

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeywordId** | **int64** | The identifier of the negative keyword. | [optional] [default to null]
**AdGroupId** | **int64** | The identifier of the ad group to which the negative keyword is associated. | [optional] [default to null]
**CampaignId** | **int64** | The identifier of the campaign to which the negative keyword is associated. | [optional] [default to null]
**KeywordText** | **string** | The keyword text. Maximum length of string is ten words if &#x60;matchType&#x60; is set to &#x60;negativeExact&#x60;. Maximum length is 4 words if &#x60;matchType&#x60; is set to &#x60;negativePhrase&#x60;. | [optional] [default to null]
**MatchType** | [***NegativeMatchType**](NegativeMatchType.md) |  | [optional] [default to null]
**State** | [***SbNegativeKeywordState**](SBNegativeKeywordState.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

