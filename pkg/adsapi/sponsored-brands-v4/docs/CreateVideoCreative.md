# CreateVideoCreative

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asins** | **[]string** |  | [optional] [default to null]
**ConsentToTranslate** | **bool** | If set to true and the headline and/or video are not in the marketplace&#x27;s default language, Amazon will attempt to translate them to the marketplace&#x27;s default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information. | [optional] [default to null]
**VideoAssetIds** | **[]string** | In SB API V4, &#x60;videoMediaIds&#x60; is replaced by &#x60;videoAssetIds&#x60;. &#x60;videoAssetIds&#x60; will only allow Asset Library identifiers for ad creation, but responses can include mediaIds for v1 campaigns and API V3 operations. At a future state, existing mediaIds will be added to Asset library for use in SB campaigns. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

