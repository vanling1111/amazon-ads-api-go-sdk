# HistoryBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Requested number of results. Default 100. Minimum 50. Maximum 200. | [optional] [default to null]
**EventTypes** | [***HistoryEventTypes**](HistoryEventTypes.md) |  | [default to null]
**FromDate** | **int32** | Max 90 days of history. | [default to null]
**NextToken** | **string** | token from previous response to get next set of data. | [optional] [default to null]
**PageOffset** | **int32** | Mutually exclusive with &#x27;nextToken&#x27;. Max results with pageOffset is 10000. Use nextToken instead for more results. | [optional] [default to null]
**Sort** | [***HistorySortParameter**](HistorySortParameter.md) |  | [optional] [default to null]
**ToDate** | **int32** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

