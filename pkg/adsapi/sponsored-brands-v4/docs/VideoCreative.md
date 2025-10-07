# VideoCreative

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsentToTranslate** | **bool** | If set to true and the heaadline and/or video are not in the marketplace&#x27;s default language, Amazon will attempt to translate them to the marketplace&#x27;s default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information. | [optional] [default to null]
**VideoAssetIds** | **[]string** | The assetIds of the original videos submitted by the advertiser. If &#x27;consentToTranslate&#x27; is set to true and translation is SUCCESSFUL then &#x27;videoAssetIds&#x27; will return translated video assetId whereas &#x60;originalVideoAssetIds&#x60; will return the original video assetId. In all other cases, &#x60;videoAssetIds&#x60; will return original video assetId. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

