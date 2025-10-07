# PreModerationResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdProgram** | **string** | Type of Ad program to which the pre moderation components belong to. | [optional] [default to null]
**AsinComponents** | [**[]AsinComponentResponse**](AsinComponentResponse.md) | Pre moderation result of the asin components. It will have information regarding the policy violations present if any. | [optional] [default to null]
**DateComponents** | [**[]DateComponentResponse**](DateComponentResponse.md) | Pre moderation result of the date components. It will have information regarding the policy violations present if any. | [optional] [default to null]
**ImageComponents** | [**[]ImageComponentResponse**](ImageComponentResponse.md) | Pre moderation result of the image components. It will have information regarding the policy violations present if any. | [optional] [default to null]
**Locale** | **string** | Locale value that was passed in request. | [optional] [default to null]
**PreModerationId** | **string** | Unique Id for the moderation Request. | [optional] [default to null]
**RecordId** | **string** | Id of the brand/advertiser. | [optional] [default to null]
**TextComponents** | [**[]TextComponentResponse**](TextComponentResponse.md) | Pre moderation result of the text components. It will have information regarding the policy violations present if any. | [optional] [default to null]
**VideoComponents** | [**[]VideoComponentResponse**](VideoComponentResponse.md) | Pre moderation result of the video components. It will have information regarding the policy violations present if any. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

