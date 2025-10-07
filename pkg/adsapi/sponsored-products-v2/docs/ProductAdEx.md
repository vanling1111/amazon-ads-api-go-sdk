# ProductAdEx

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdId** | **float64** | The product ad identifier. | [optional] [default to null]
**CampaignId** | **float64** | The campaign identifier. | [optional] [default to null]
**AdGroupId** | **float64** | The ad group identifier. | [optional] [default to null]
**Sku** | **string** | The SKU associated with the product. Defined for seller accounts only. | [optional] [default to null]
**Asin** | **string** | The ASIN associated with the product. Defined for vendors only. | [optional] [default to null]
**State** | [***State**](State.md) |  | [optional] [default to null]
**CreationDate** | **float64** | The epoch date the product ad was created. | [optional] [default to null]
**LastUpdatedDate** | **float64** | The epoch date the product ad was last updated. | [optional] [default to null]
**ServingStatus** | **string** | The computed status of the product ad. See the [developer notes](https://advertising.amazon.com/API/docs/en-us/reference/concepts/developer-notes) for more information. | [optional] [default to null]
**ServingStatusDetails** | [**[]ProductAdExServingStatusDetails**](ProductAdEx_servingStatusDetails.md) | Details of serving status. Only statuses related to moderation according to the ad policy are currently included. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

