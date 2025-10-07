# MediaCompleteBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadLocation** | **string** |  | [optional] [default to null]
**Version** | **string** | The version id of the uploaded media. The upload location retrieved from /media/upload API supports versioning and returns version id in the upload response through &#x60;x-amz-version-id&#x60; header. API user can explicitly specify the version id corresponding to an upload through &#x60;version&#x60; property. &#x60;version&#x60; is optional and if it is not specified, media corresponding to the most recent version at the time of API call will be used. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

