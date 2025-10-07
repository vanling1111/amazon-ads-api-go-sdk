# InlineResponse2004

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaId** | **string** |  | [optional] [default to null]
**Status** | [***MediaStatus**](MediaStatus.md) |  | [optional] [default to null]
**StatusMetadata** | [**[]InlineResponse2004StatusMetadata**](inline_response_200_4_statusMetadata.md) |  | [optional] [default to null]
**PublishedMediaUrl** | **string** | The preview URL of the media. It is only available when status is &#x60;Available&#x60;. | [optional] [default to null]
**OriginalMediaUrl** | **string** | This is a signed URL which returns the original media in .mp4 file extension. The URL is only active for 7 days and requires to be regenerated if the video is not downloaded within 7 days. If you try to upload the downloaded video using the Asset Library API and get an error, then please retry upload after changing the file extension to .mov. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

