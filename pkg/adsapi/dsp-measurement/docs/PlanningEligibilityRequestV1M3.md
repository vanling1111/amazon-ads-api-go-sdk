# PlanningEligibilityRequestV1M3

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertiserId** | **string** | The advertiserId. | [optional] [default to null]
**Locale** | [***MeasurementLocaleV1**](MeasurementLocaleV1.md) |  | [optional] [default to null]
**OrderMetadata** | [**[]PlanningOrderMetadataV1M3**](PlanningOrderMetadataV1M3.md) |  | [optional] [default to null]
**StudyTypeFilters** | [**[]StudyTypeV1M2**](StudyTypeV1M2.md) | StudyType identifier filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]
**FundingTypeFilters** | [**[]FundingTypeV1M2**](FundingTypeV1M2.md) | FundingType filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]
**VendorProductIdFilters** | **[]string** | VendorProduct identifier filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]
**VendorTypeFilters** | [**[]VendorTypeV1M2**](VendorTypeV1M2.md) | VendorType filters to be applied when checking eligibility status. If not supplied we will check against all available vendor products. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

