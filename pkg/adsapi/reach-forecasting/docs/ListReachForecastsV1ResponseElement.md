# ListReachForecastsV1ResponseElement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertisedProductCategoryIds** | **[]string** | The identifiers of the advertised product categories for the forecast. Use the DSP [ListAdvertisedProductCategories API](https://advertising.amazon.com/API/docs/en-us/dsp-discovery-advertised-product-categories#tag/Discovery-Advertised-Product-Categories/operation/DspListAdvertisedProductCategoriesV1) to look up advertised product category IDs. | [optional] [default to null]
**AvailableImpressions** | **int64** | The number of impressions available for you to purchase after considering contention (G - booked demand) among the matchingImpressions | [optional] [default to null]
**AvgCpm** | **float64** | The CPM rate (cost per thousand impressions). | [optional] [default to null]
**CountryCode** | [***CountryCodeV1**](CountryCodeV1.md) |  | [default to null]
**Cpc** | **float64** | The CPC rate (cost per click). | [optional] [default to null]
**CreationDateTime** | [**time.Time**](time.Time.md) | The date time that the reach forecast was created. | [default to null]
**CurrencyCode** | [***CurrencyCodeV1**](CurrencyCodeV1.md) |  | [default to null]
**DataPoints** | [**[]ReachCurveDataPointV1**](ReachCurveDataPointV1.md) | The list of data points for the reach curve. | [default to null]
**DeliveryType** | [***DeliveryTypeV1**](DeliveryTypeV1.md) |  | [default to null]
**EndDate** | **string** | The forecast end date in YYYY-MM-DD format. | [default to null]
**FrequencyCap** | [***FrequencyCapV1**](FrequencyCapV1.md) |  | [optional] [default to null]
**MatchingImpressions** | **int64** | The number of impressions that match your targeting | [optional] [default to null]
**MaxCpm** | **float64** | The maximum CPM rate (cost per thousand impressions). | [optional] [default to null]
**ReachForecastId** | **string** | This is the unique identifier of the Reach Forecast resource. | [default to null]
**ReachType** | [***ReachTypeV1**](ReachTypeV1.md) |  | [default to null]
**StartDate** | **string** | The forecast start date in YYYY-MM-DD format. | [default to null]
**Status** | [***ReachForecastStatusV1**](ReachForecastStatusV1.md) |  | [default to null]
**SupplyPackage** | [**[]SupplyV1**](SupplyV1.md) | The combination of Ads supply. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

