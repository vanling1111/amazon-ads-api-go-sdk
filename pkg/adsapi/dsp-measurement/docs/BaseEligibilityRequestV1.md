# BaseEligibilityRequestV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FundingTypeFilters** | [**[]FundingTypeV1**](FundingTypeV1.md) | FundingType filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]
**VendorProductIdFilters** | **[]string** | VendorProduct identifier filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]
**VendorTypeFilters** | [**[]VendorTypeV1**](VendorTypeV1.md) | VendorType filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

