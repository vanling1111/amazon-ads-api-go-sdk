# SbTargetingGetRefinementsForCategoryResponseContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgeRanges** | [**[]SbTargetingAgeRange**](SBTargetingAgeRange.md) | List of Age Ranges. Use /sb/targets/categories/{categoryRefinementId}/refinements to retrieve Age Ranges. Age Ranges are only available for categories related to children&#x27;s toys and games. | [optional] [default to null]
**Brands** | [**[]SbTargetingBrand**](SBTargetingBrand.md) | List of Brands. | [optional] [default to null]
**Genres** | [**[]SbTargetingGenre**](SBTargetingGenre.md) | List of Genres. Use /sb/targets/categories/{categoryRefinementId}/refinements to retrieve Genre Node IDs. Genres are only available for categories related to books. | [optional] [default to null]
**NextToken** | **string** | Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;NextToken&#x60; field is empty, there are no further results. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

