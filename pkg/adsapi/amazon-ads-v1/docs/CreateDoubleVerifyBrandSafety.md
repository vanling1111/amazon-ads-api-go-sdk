# CreateDoubleVerifyBrandSafety

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppAgeRating** | [**[]DvBrandSafetyAppAgeRatingType**](DVBrandSafetyAppAgeRatingType.md) | A list of app age ratings to be used for excluding apps. For example, TEENS_12_PLUS will only exclude apps with content rated for everyone ages 12 and over. UNKNOWN will exclude apps with content unrated or unknown to Double Verify. | [optional] [default to null]
**AppStarRating** | [***DvBrandSafetyAppStarRatingType**](DVBrandSafetyAppStarRatingType.md) |  | [optional] [default to null]
**ContentCategories** | [**[]DvBrandSafetyContentCategoryType**](DVBrandSafetyContentCategoryType.md) | A list of content categories to exclude from targeting. | [optional] [default to null]
**ContentCategoriesWithRisk** | [**[]CreateDvBrandSafetyContentCategoriesWithRiskMap**](CreateDVBrandSafetyContentCategoriesWithRiskMap.md) |  | [optional] [default to null]
**ExcludeAppsWithInsufficientRating** | **bool** | Set to true to exclude unofficial apps or apps with insufficient user ratings (&lt;100 lifetime). | [optional] [default to null]
**UnknownContent** | **bool** | Set to true to exclude unknown content. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

