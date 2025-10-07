# DspOmnichannelMetricsEligibilityRequestV1M2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandIds** | **[]string** | A list of canonical brand identifiers. | [optional] [default to null]
**CurrentStudyId** | **string** | Optional current study identifier. If provided orders are expected to be added into this study and the orders already associated with this study will be excluded from certain eligibility checks. | [optional] [default to null]
**ExcludedLineItemIds** | **[]string** | A list of canonical lineItem identifiers that are excluded from the eligibility check. | [optional] [default to null]
**OrderIds** | **[]string** | A list of canonical order identifiers. By default all lineItems in those orders will be included. | [optional] [default to null]
**FundingTypeFilters** | [**[]FundingTypeV1M2**](FundingTypeV1M2.md) | FundingType filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]
**VendorProductIdFilters** | **[]string** | VendorProduct identifier filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]
**VendorTypeFilters** | [**[]VendorTypeV1M2**](VendorTypeV1M2.md) | VendorType filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

