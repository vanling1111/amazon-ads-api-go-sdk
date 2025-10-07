# BillingAggregation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAggregationId** | **string** | An identifier that helps associate this invoice with specific billing entities, such as campaigns or rodeo advertiser accounts. | [optional] [default to null]
**BillingAggregationResourcePath** | **string** | The resource path to suffix to the base URL endpoints to retrieve the corresponding billing aggregation entity, such as a specific campaign or all active campaigns under DSP Rodeo Advertiser. An example URL after suffixing will be &#x27;https://advertising-api.amazon.com/dsp/orders/123&#x27;. The base URL endpoints for the different regions and marketplaces can be found in the Amazon Advertising API Reference Portal. | [optional] [default to null]
**BillingLevel** | [***BillingLevel**](billingLevel.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

