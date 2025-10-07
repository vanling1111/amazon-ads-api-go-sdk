# LandingPage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asins** | **[]string** |  | [optional] [default to null]
**PageType** | [***LandingPageType**](LandingPageType.md) |  | [optional] [default to null]
**Url** | **string** | URL of an existing simple landing page or Store page. Vendors may also specify the URL of a custom landing page. If a custom URL is specified, the landing page must include the ASINs of at least three products that are advertised as part of the campaign. Do not include this property in the request if the asins property is also included, these properties are mutually exclusive. Note that brandVideo ads only support Store page as landing page and does not allow asins property. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

