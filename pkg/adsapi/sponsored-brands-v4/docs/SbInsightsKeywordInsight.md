# SbInsightsKeywordInsight

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | [**[]SbInsightsKeywordAlertType**](SBInsightsKeywordAlertType.md) |  | [optional] [default to null]
**SearchTermImpressionShare** | **float64** | The account-level ad-attributed impression share for the search-term / keyword. Provides percentage share of all ad impressions the advertiser has for the keyword in the last 7 days. This metric helps advertisers identify potential opportunities based on their share of relevant keywords. This information is only available for keywords the advertiser targeted with ad impressions. Only available in the following marketplaces: US, CA, MX, UK, DE, FR, ES, IT, IN, JP. | [optional] [default to null]
**MatchType** | [***SbInsightsMatchType**](SBInsightsMatchType.md) |  | [optional] [default to null]
**AdGroupIndex** | **int32** | Correlates the ad group to the ad group array index specified in the request. Zero-based. | [optional] [default to null]
**SearchTermImpressionRank** | **int32** | The account-level ad-attributed impression rank for the search-term / keyword. Provides the [1:N] place the advertiser ranks among all advertisers for the keyword by ad impressions in a marketplace in the last 7 days. It tells an advertiser how many advertisers had higher share of ad impressions. This information is only available for keywords the advertiser targeted with ad impressions. Only available in the following marketplaces: US, CA, MX, UK, DE, FR, ES, IT, IN, JP. | [optional] [default to null]
**Bid** | **float64** | The associated bid. Note that this value must be less than the budget associated with the Advertiser account. For more information, see [supported features](https://advertising.amazon.com/API/docs/v2/guides/supported_features). | [optional] [default to null]
**KeywordIndex** | **int32** | Correlates the keyword to the keyword array index specified in the request. Zero-based. | [optional] [default to null]
**KeywordText** | **string** | The keyword text. Maximum of 10 words. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

