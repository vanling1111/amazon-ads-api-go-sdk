# SdForecastResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BidOptimization** | **string** |  | [optional] [default to null]
**LifetimeForecasts** | [**[]Forecast**](Forecast.md) | Forecasts for campaign start date and end date. Default end date is start date plus 7 days. | [optional] [default to null]
**WeeklyForecasts** | [**[]Forecast**](Forecast.md) | Weekly average forecasts. | [optional] [default to null]
**DailyForecasts** | [**[]Forecast**](Forecast.md) | Daily average forecasts. | [optional] [default to null]
**Curves** | [**[]Curve**](Curve.md) | Forecasting curves. | [optional] [default to null]
**ForecastStatus** | [***ForecastStatus**](ForecastStatus.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

