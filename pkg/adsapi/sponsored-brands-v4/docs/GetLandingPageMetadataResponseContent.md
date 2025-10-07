# GetLandingPageMetadataResponseContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageType** | **string** | The type of landing page, such as store page, product list (simple landing page), custom url. | Page Type    | |--------------| | PRODUCT_LIST | | STORE        | | CUSTOM_URL   | | DETAIL_PAGE  | | [default to null]
**CanonicalUrl** | **string** | A canonical URL is the URL that represents the best version of landing page URL from a group of duplicate landing page URLs. For example, if you have two landing page URLs for the same page (such as amazon.it/HSA/pages/default?pageId&#x3D;B59A592C-8A12-4684-A37E-2416FD594D87 and amazon.it/stores/page/B59A592C-8A12-4684-A37E-2416FD594D87), we chooses one as canonical. In this case, canonical url is amazon.it/stores/page/B59A592C-8A12-4684-A37E-2416FD594D87 | [default to null]
**UnSupportedReason** | **string** | A human-readable description of the unSupportedReasonCode field. | [optional] [default to null]
**IsSupported** | **bool** | This field determines whether the landing page is supported for the ad product. | [optional] [default to null]
**UnSupportedReasonCode** | **string** | Enumerated code for why landing page is unsupported. | Reason Code                 | | SB_DETAIL_PAGE_UNSUPPORTED  | | SB_GATEWAY_PAGE_UNSUPPORTED | | SB_SEARCH_PAGE_UNSUPPORTED  | | SB_BROWSE_PAGE_UNSUPPORTED  | | SB_OTHER_PAGE_UNSUPPORTED   | | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

