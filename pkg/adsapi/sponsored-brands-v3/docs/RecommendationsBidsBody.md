# RecommendationsBidsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int64** | The identifier of the campaign for which bid recommendations are created. | [optional] [default to null]
**Targets** | [**[][]SbExpression**](array.md) | Sum of the sizes of (targets + keywords + themeTypes) arrays should be &lt;&#x3D; 100 | [optional] [default to null]
**Keywords** | [**[]SbBidRecommendationKeyword**](SBBidRecommendationKeyword.md) | Sum of the sizes of (targets + keywords + themeTypes) arrays should be &lt;&#x3D; 100 | [optional] [default to null]
**AdFormat** | [***AdFormat**](AdFormat.md) |  | [optional] [default to null]
**CostType** | **string** | Optional. Support &#x60;CPC&#x60; (cost per click) and &#x60;VCPM&#x60; (cost per thousand viewable impressions). | costType | goal | Expected result | |----------|------|-----------------| | Empty | Empty | Recommendation will be generated for costType&#x3D; &#x60;CPC&#x60; and goal&#x3D;&#x60;PAGE_VISIT&#x60;. | Specified | Empty | It will return error and no recommendation will be generated. | Empty | Specified | It will use default costType based on goal selection. e.g. goal&#x3D;&#x60;BRAND_IMPRESSION_SHARE&#x60; then it will use costType &#x3D; &#x60;VCPM&#x60; to generate recommendation. | Specified | Specified | The recommendation will be based on selected goal and costType. If the mismatch then it will generate an error. | [optional] [default to null]
**ThemeTypes** | **[]string** | Sum of the sizes of (targets + keywords + themeTypes) arrays should be &lt;&#x3D; 100 | [optional] [default to null]
**Goal** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

