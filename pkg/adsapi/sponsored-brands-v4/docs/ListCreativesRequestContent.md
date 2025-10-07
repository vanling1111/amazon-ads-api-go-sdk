# ListCreativesRequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreativeTypeFilter** | [**[]CreativeType**](CreativeType.md) | Filters creatives by optional creative type. By default, you can list all creative versions regardless of creative type. | [optional] [default to null]
**AdId** | **string** | The unique ID of a Sponsored Brands ad. | [default to null]
**NextToken** | **string** | Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;NextToken&#x60; field is empty, there are no further results. | [optional] [default to null]
**MaxResults** | **float64** | Set a limit on the number of results returned by an operation. | [optional] [default to null]
**CreativeVersionFilter** | **[]string** | Filters creatives by optional creative version. This means you can either list all creative versions without specific creative version filter, all just retrieve a single creative version by providing a specific version identifier. | [optional] [default to null]
**CreativeStatusFilter** | [**[]CreativeStatus**](CreativeStatus.md) | Filters creatives by optional creative status. By default, you can list all creative versions regardless of creative status. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

