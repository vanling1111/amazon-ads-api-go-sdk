# InlineResponse20011

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateTargetSuccessResults** | [**[]InlineResponse20011UpdateTargetSuccessResults**](inline_response_200_11_updateTargetSuccessResults.md) | Lists the successfully updated negative targets. Note that negative targets in the response are correlated to negative targets in the request using the &#x60;targetRequestIndex&#x60; field. For example, if &#x60;targetRequestIndex&#x60; is set to &#x60;2&#x60;, the values correlate to the third negative target object in the request. | [optional] [default to null]
**UpdateTargetErrorResults** | [**[]AllOfinlineResponse20011UpdateTargetErrorResultsItems**](.md) | Lists errors that occured during negative target update. Note that errors are correlated to negative target update requests by the &#x60;negativeTargetRequestIndex&#x60; field. This field corresponds to the order of the negative target in the request. For example, if &#x60;negativeTargetRequestIndex&#x60; is set to &#x60;2&#x60;, the values correlate to the third negative target object in the request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

