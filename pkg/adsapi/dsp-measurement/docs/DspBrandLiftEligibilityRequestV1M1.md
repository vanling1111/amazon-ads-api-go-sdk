# DspBrandLiftEligibilityRequestV1M1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentStudyId** | **string** | Optional current study identifier, if provided orders are expected to be added into this study and the orders already associated with this study will be excluded from certain eligibility check. | [optional] [default to null]
**ExcludedLineItemIds** | **[]string** | A list of canonical lineItem identifiers that are excluded from the eligibility check. | [optional] [default to null]
**OrderIds** | **[]string** | A list of canonical order identifiers. By default all lineItems in those orders will be included. | [optional] [default to null]
**FundingTypeFilters** | [**[]FundingTypeV1M1**](FundingTypeV1M1.md) | FundingType filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]
**VendorProductIdFilters** | **[]string** | VendorProduct identifier filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]
**VendorTypeFilters** | [**[]VendorTypeV1M1**](VendorTypeV1M1.md) | VendorType filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

