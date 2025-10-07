# SubmitImageTasksResponseContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Submitted** | [**[]Submitted**](Submitted.md) |  | [optional] [default to null]
**BatchId** | **string** | As per API First guidance, batch API should return a separate list for success and errors in the response. The success/submitted and error fields will indicate the status of submission, they don&#x27;t mean the status of image generation task. Status code will be 207 for partial successful requests and all successful requests. A batchId that is used to track status multiple tasks if they are submitted in one batch request If none of the request is submitted successfully, batchId will be null | [optional] [default to null]
**Error_** | [**[]ErrorDetails**](ErrorDetails.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

