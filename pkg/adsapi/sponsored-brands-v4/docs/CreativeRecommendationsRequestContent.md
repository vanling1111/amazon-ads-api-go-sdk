# CreativeRecommendationsRequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreativeType** | **string** | Supported are PRODUCT_COLLECTION, STORE_SPOTLIGHT, VIDEO, BRAND_VIDEO. More could be added in future. | [default to null]
**NextToken** | **string** | Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;NextToken&#x60; field is empty, there are no further results. | [optional] [default to null]
**MaxResults** | **float64** | Set a limit on the number of results returned by an operation. | [optional] [default to null]
**Source** | [***Source**](Source.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

