# InlineResponse2009

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateTargetSuccessResults** | [**[]InlineResponse2009UpdateTargetSuccessResults**](inline_response_200_9_updateTargetSuccessResults.md) | Lists the successfully updated targets. Note that targets in the response are correlated to targets in the request using the &#x60;targetRequestIndex&#x60; field. For example, if &#x60;targetRequestIndex&#x60; is set to &#x60;2&#x60;, the values correlate to the third target object in the request. | [optional] [default to null]
**UpdateTargetErrorResults** | [**[]AllOfinlineResponse2009UpdateTargetErrorResultsItems**](.md) | Lists errors that occured during target update. Note that errors are correlated to target update requests by the &#x60;targetRequestIndex&#x60; field. This field corresponds to the order of the target in the request. For example, if &#x60;targetRequestIndex&#x60; is set to &#x60;2&#x60;, the values correlate to the third target object in the request array. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

