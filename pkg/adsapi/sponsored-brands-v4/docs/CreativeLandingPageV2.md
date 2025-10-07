# CreativeLandingPageV2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asins** | **[]string** | The list of asins on the landingPage If type is PRODUCT_LIST. A minimum of 3 asins are required. For the &#x27;PRODUCT_LIST&#x27; type, the asins property is mandatory, and the url should not be included. | [optional] [default to null]
**Type_** | **string** | Supported types are PRODUCT_LIST, STORE, DETAIL_PAGE, CUSTOM_URL. More could be added in future. | [optional] [default to null]
**Url** | **string** | The url of the landingPage. When including the &#x27;asins&#x27; property in the request, do not include this property, as they are mutually exclusive. For the PRODUCT_LIST type, the asins property is mandatory, and the url should not be included. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

