# UpdateCampaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortfolioId** | **string** | The identifier of an existing portfolio to which the campaign is associated. | [optional] [default to null]
**Bidding** | [***Bidding**](Bidding.md) |  | [optional] [default to null]
**EndDate** | **string** | endDate is optional. If endDate is specified, startDate must be specified as well. Note: This property is nullable. If null is explicitly provided in a campaign update request, any existing endDate for the campaign will be removed. | [optional] [default to null]
**CampaignId** | **string** | The identifier of the campaign. | [default to null]
**Name** | **string** | The name of the campaign. | [optional] [default to null]
**State** | [***CreateOrUpdateEntityState**](CreateOrUpdateEntityState.md) |  | [optional] [default to null]
**StartDate** | **string** | startDate can only be changed if the current startDate is in the future. | [optional] [default to null]
**Budget** | **float64** | The budget of the campaign. See https://advertising.amazon.com/help?entityId&#x3D;ENTITYJDATFOIA05Q7#GE5QEBS6QRJJAT3A | [optional] [default to null]
**Tags** | [***map[string]string**](map.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

