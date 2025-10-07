# CreateCreative

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdGroupId** | **float64** | Unqiue identifier for the ad group associated with the creative. | [default to null]
**CreativeType** | [***CreativeTypeInCreativeRequest**](CreativeTypeInCreativeRequest.md) |  | [optional] [default to null]
**Properties** | [***CreativeProperties**](CreativeProperties.md) |  | [default to null]
**ConsentToTranslate** | **bool** | If set to true and the headline and/or video are not in the marketplace&#x27;s default language, Amazon will attempt to translate them to the marketplace&#x27;s default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

