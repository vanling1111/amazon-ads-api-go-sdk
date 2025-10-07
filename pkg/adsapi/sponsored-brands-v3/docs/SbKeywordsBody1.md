# SbKeywordsBody1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdGroupId** | **int64** | The identifier of an existing ad group to which the keyword is associated. | [optional] [default to null]
**CampaignId** | **int64** | The identifier of an existing campaign to which the keyword is associated. | [optional] [default to null]
**KeywordText** | **string** | The keyword text. The maximum number of words for this string is 10. | [optional] [default to null]
**NativeLanguageKeyword** | **string** | The unlocalized keyword text in the preferred locale of the advertiser. | [optional] [default to null]
**NativeLanguageLocale** | **string** | The locale preference of the advertiser. For example, if the advertiser’s preferred language is Simplified Chinese, set the locale to &#x60;zh_CN&#x60;. Supported locales include: Simplified Chinese (locale: zh_CN) for US, UK and CA. English (locale: en_GB) for DE, FR, IT and ES. | [optional] [default to null]
**MatchType** | [***MatchType**](MatchType.md) |  | [optional] [default to null]
**Bid** | **float64** | The bid associated with the keyword. Note that this value must be less than the budget associated with the Advertiser account. For more information, see [Keyword bid constraints by marketplace](https://advertising.amazon.com/API/docs/en-us/reference/concepts/limits#bid-constraints-by-marketplace). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

