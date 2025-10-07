# BatchGetConversionDefinitionsAssociatedForOrdersRequestV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxResults** | **int32** | Sets the maximum number of conversions in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [optional] [default to null]
**NextToken** | **string** | Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | [optional] [default to null]
**OrderIds** | **[]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

