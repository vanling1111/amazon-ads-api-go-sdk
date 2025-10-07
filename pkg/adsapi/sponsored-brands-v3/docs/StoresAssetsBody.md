# StoresAssetsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetInfo** | **string** | A JSON object specifying the Brand entity identifier an media type. The Brand entity identifier is optional, but media type is not. |Field|Type|Values| |-----|-----|-----| |brandEntityId|string|The Brand entity identifier.| |mediaType|string| Only &#x60;brandLogo&#x60; is currently supported.| Example: &#x60;&#x60;&#x60; {   brandEntityId: \&quot;12345678\&quot;,   mediaType: {     \&quot;brandLogo\&quot;   } } &#x60;&#x60;&#x60; | [optional] [default to null]
**Asset** | [****os.File**](*os.File.md) | The binary data for the image. For more information,  File size must be smaller than 1MB, and the resolution must be a minimum of 400 pixels by 400 pixels. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

