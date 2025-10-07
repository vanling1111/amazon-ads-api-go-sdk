# CreateReachForecastsV1ResponseElement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableImpressions** | **int64** | The number of impressions available for you to purchase after considering contention (G - booked demand) among the matchingImpressions | [optional] [default to null]
**AvgCpm** | **float64** | The CPM rate (cost per thousand impressions). | [optional] [default to null]
**CountryCode** | [***CountryCodeV1**](CountryCodeV1.md) |  | [default to null]
**Cpc** | **float64** | The CPC rate (cost per click). | [optional] [default to null]
**CreationDateTime** | [**time.Time**](time.Time.md) | The date time that the reach forecast was created. | [default to null]
**CurrencyCode** | [***CurrencyCodeV1**](CurrencyCodeV1.md) |  | [default to null]
**DataPoints** | [**[]ReachCurveDataPointV1**](ReachCurveDataPointV1.md) | The list of data points for the reach curve. | [default to null]
**MatchingImpressions** | **int64** | The number of impressions that match your targeting | [optional] [default to null]
**MaxCpm** | **float64** | The maximum CPM rate (cost per thousand impressions). | [optional] [default to null]
**ReachForecastId** | **string** | This is the unique identifier of the Reach Forecast resource. | [default to null]
**Status** | [***ReachForecastStatusV1**](ReachForecastStatusV1.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

