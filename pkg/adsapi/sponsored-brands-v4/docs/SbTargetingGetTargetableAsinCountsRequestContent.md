# SbTargetingGetTargetableAsinCountsRequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgeRanges** | **[]string** | List of Age Range Refinement Ids. | [optional] [default to null]
**Brands** | **[]string** | List of Brand Refinement Ids. | [optional] [default to null]
**Genres** | **[]string** | List of Genre Refinement Ids. | [optional] [default to null]
**IsPrimeShipping** | **bool** | Indicates if products have prime shipping. Leave empty to include both prime shipping and non-prime shipping products. | [optional] [default to null]
**RatingRange** | [***SbTargetingRatingRange**](SBTargetingRatingRange.md) |  | [optional] [default to null]
**Category** | **string** | The category refinement id. Please use /sb/targets/categories or /sb/recommendations/targets/category to retrieve category IDs. | [default to null]
**PriceRange** | [***SbTargetingPriceRange**](SBTargetingPriceRange.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

