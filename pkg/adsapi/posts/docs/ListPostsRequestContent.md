# ListPostsRequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**[]PostListFilter**](PostListFilter.md) | A list of post filters. | [optional] [default to null]
**MaxResults** | **float64** | Number of items to be returned on this call. | [optional] [default to null]
**MetricEndDate** | **string** | The end date to get metrics for. The value is in ISO8601 full-date format (UTC). For example, 2020-08-16. | [optional] [default to null]
**MetricStartDate** | **string** | The start date to get metrics for. The value is in ISO8601 full-date format (UTC). For example, 2020-08-16. | [optional] [default to null]
**NextToken** | **string** | Pagination token to get next page of results from previous call. | [optional] [default to null]
**ProfileId** | **string** | Identifier for a profile. | [default to null]
**SelectedMetrics** | [**[]MetricName**](MetricName.md) | A list of metrics to return for each post. | [optional] [default to null]
**SortCriterion** | [***PostListSortCriterion**](PostListSortCriterion.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

