# HeadlineSuggestionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asins** | **[]string** | An array of ASINs associated with the creative. Note do not pass an empty array, this results in an error.  | [optional] [default to null]
**StorePages** | [**[]StorePage**](StorePage.md) | An array of Store Pages associated with SB Spotlight Creative. | [optional] [default to null]
**MaxNumSuggestions** | **float64** | Maximum number of suggestions that API should return. Response will [0, maxNumSuggestions] suggestions (suggestions are not guaranteed). | [optional] [default to null]
**AdFormat** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

