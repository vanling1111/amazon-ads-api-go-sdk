# SbLandingPage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asins** | **[]string** | An array of ASINs used to generate a simple landing page. The response includes the URL of the generated simple landing page. Do not include this property in the request if the &#x60;url&#x60; property is also included, these properties are mutually exclusive. | [optional] [default to null]
**Url** | **string** | URL of an existing simple landing page or Store page. Vendors may also specify the URL of a custom landing page. If a custom URL is specified, the landing page must include the ASINs of at least three products that are advertised as part of the campaign. Do not include this property in the request if the &#x60;asins&#x60; property is also included, these properties are mutually exclusive. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

