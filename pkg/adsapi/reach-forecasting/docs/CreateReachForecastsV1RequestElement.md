# CreateReachForecastsV1RequestElement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertisedProductCategoryIds** | **[]string** | The identifiers of the advertised product categories for the forecast. Use the DSP [ListAdvertisedProductCategories API](https://advertising.amazon.com/API/docs/en-us/dsp-discovery-advertised-product-categories#tag/Discovery-Advertised-Product-Categories/operation/DspListAdvertisedProductCategoriesV1) to look up advertised product category IDs. | [optional] [default to null]
**CountryCode** | [***CountryCodeV1**](CountryCodeV1.md) |  | [default to null]
**DeliveryType** | [***DeliveryTypeV1**](DeliveryTypeV1.md) |  | [default to null]
**EndDate** | **string** | The forecast end date in YYYY-MM-DD format. | [default to null]
**FrequencyCap** | [***FrequencyCapV1**](FrequencyCapV1.md) |  | [optional] [default to null]
**ReachType** | [***ReachTypeV1**](ReachTypeV1.md) |  | [default to null]
**StartDate** | **string** | The forecast start date in YYYY-MM-DD format. | [default to null]
**SupplyPackage** | [**[]SupplyV1**](SupplyV1.md) | The combination of Ads supply. | [optional] [default to null]
**Targets** | [**[]PlanningTargetV1**](PlanningTargetV1.md) | The list of targets for the forecast. Targets of the same targetType and of the same negative boolean are combined using OR statements. Targets of different types are associated using AND statements. Note: Audience targets of the same group (as specified by groupId) are linked via an OR linkage. Audience target groups are linked via an AND linkage. Audience targets without groupId are consider as different group, thus are also linked via an AND linkage. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

