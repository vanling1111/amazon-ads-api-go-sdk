# GetAsinEngagementForStoreRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | [***AsinEngagementDimension**](AsinEngagementDimension.md) |  | [optional] [default to null]
**EndDate** | **string** | The end date (inclusive) in YYYY-MM-DD format for the time period from when to fetch the insights. | [default to null]
**Metrics** | [**[]AsinEngagementMetric**](AsinEngagementMetric.md) | List of the engagement metrics to be fetched. At least one metric should be specified. | [default to null]
**OrderBy** | [***SortOrder**](SortOrder.md) |  | [optional] [default to null]
**SortBy** | [***Object**](.md) | Nullable metric to sort on. If a value is provided, it must also appear in the metrics list. If no value is provided, the result is not guaranteed to be sorted. This field is only valid when the dimension is ASIN. | [optional] [default to null]
**StartDate** | **string** | The start date (inclusive) in YYYY-MM-DD format for the time period from when to fetch the insights. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

