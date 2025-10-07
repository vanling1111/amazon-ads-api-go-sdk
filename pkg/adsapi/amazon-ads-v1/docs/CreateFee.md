# CreateFee

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddToBudgetSpentAmount** | **bool** | Applies only to THIRD_PARTY_APPLIED_FEE. When set to true, third-party applied fees are are added on top of the total ad group budget spent amount in reports. | [optional] [default to null]
**FeeType** | [***FeeType**](FeeType.md) |  | [optional] [default to null]
**FeeValue** | **float64** | The fee amount expressed as the feeValueType. AMAZON_AUDIENCE_FEE AND THIRD_PARTY_AUDIENCE_FEE is in the currency of the marketplace. All other CPM based fees are in the currency of the advertiser. For percentages, 100 represents 100%. | [optional] [default to null]
**ThirdPartyProvider** | [***FeesThirdPartyProvider**](FeesThirdPartyProvider.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

