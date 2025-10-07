# ProductEligibilityRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdType** | **string** | Set to &#x27;sp&#x27; to check product eligibility for Sponsored Products advertisements. Set to &#x27;sb&#x27; to check product eligibility for Sponsored Brands advertisements. Set to &#x27;sd&#x27; to check product eligibility for Sponsored Displays advertisements. Set to &#x27;dsp&#x27; to check product eligibility for Demand Side Platform advertisements. | [optional] [default to AD_TYPE.SP]
**Locale** | **string** | Set locale string as \&quot;en_US\&quot; to specify the language in which the response is returned | [optional] [default to null]
**ProductDetailsList** | [**[]ProductDetails**](ProductDetails.md) | A list of product identifier objects. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

