# DspAudienceResearchEligibilityRequestV1M2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceTargetingGroup** | [***AudienceTargetingGroupV1M2**](AudienceTargetingGroupV1M2.md) |  | [optional] [default to null]
**FundingTypeFilters** | [**[]FundingTypeV1M2**](FundingTypeV1M2.md) | FundingType filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]
**VendorProductIdFilters** | **[]string** | VendorProduct identifier filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]
**VendorTypeFilters** | [**[]VendorTypeV1M2**](VendorTypeV1M2.md) | VendorType filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

