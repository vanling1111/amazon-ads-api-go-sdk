# DspAudienceFieldsV1Fees

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Fee amount in base currency units, multiplied by scaling factor (&#x27;scale&#x27;). | [optional] [default to null]
**Currency** | **string** | Base currency, such as US Dollar. | [optional] [default to null]
**FeeCalculationType** | **string** | How the fee is applied. | [optional] [default to null]
**ImpressionSupplyType** | **string** | To which supply types this fee applies to. The fee may be different for different supply types. | [optional] [default to null]
**Scale** | **int32** | Scale of the amount relative to the base currency unit. For instance, if the scale is 1000, the currency is USD, and the amount is 500, the human-readable fee is $0.50. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

