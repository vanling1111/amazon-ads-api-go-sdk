# ProductMetadataRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdType** | **string** | Program type. Required if checks advertising eligibility:   * SP - Sponsored Product   * SB - Sponsored Brand   * SD - Sponsored Display | [optional] [default to null]
**Asins** | **[]string** | Specific asins to search for in the advertiser&#x27;s inventory. Cannot use together with skus or searchStr input types. | [optional] [default to null]
**CheckEligibility** | **bool** | Whether advertising eligibility info is required | [optional] [default to false]
**CheckItemDetails** | **bool** | Whether item details such as name, image, and price is required. | [optional] [default to false]
**CursorToken** | **string** | Pagination token used for the suggested sort type or for author merchant | [optional] [default to null]
**IsGlobalStoreSelection** | **bool** |  This will return only GlobalStore listings related to [GlobalStore Program](https://sellercentral.amazon.com/help/hub/reference/external/202139180) and not local listings | [optional] [default to false]
**Locale** | **string** | Optional locale for detail and eligibility response strings. Default to the marketplace locale. | [optional] [default to null]
**PageIndex** | **int32** | Index of the page to be returned; For author, this value will be ignored, should use cursorToken instead. For seller and vendor, results are capped at 10k ((pageIndex + 1) * pageSize). | [default to null]
**PageSize** | **int32** | Number of items to be returned on this page index. | [default to null]
**SearchStr** | **string** | Specific string in the item title to search for in the advertiser&#x27;s inventory. Case insensitive. Cannot use together with asins or skus input types. | [optional] [default to null]
**Skus** | **[]string** | Specific SKUs to search for in the advertiser&#x27;s inventory. Currently only support SP program type for sellers. Cannot use together with asins or searchStr input types. | [optional] [default to null]
**SortBy** | **string** | Sort option for the result. Currently only support SP program type for sellers:   * SUGGESTED - Suggested products are those most likely to engage customers, and have a higher chance of generating clicks if advertised.   * CREATED_DATE - Date the item listing was created  | [optional] [default to null]
**SortOrder** | **string** | Sort order (has to be DESC for the suggested sort type):   * ASC - Ascending, from A to Z   * DESC - Descending, from Z to A | [optional] [default to SORT_ORDER.DESC]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

