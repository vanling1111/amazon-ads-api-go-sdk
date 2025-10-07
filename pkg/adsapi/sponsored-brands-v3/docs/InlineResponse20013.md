# InlineResponse20013

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | [**[]InlineResponse20013Success**](inline_response_200_13_success.md) | Lists the successfully updated theme targets. Note that theme targets in the response are correlated to theme targets in the request using the &#x60;index&#x60; field. For example, if &#x60;index&#x60; is set to &#x60;2&#x60;, the values correlate to the third theme target object in the request. | [optional] [default to null]
**Error_** | [**[]AllOfinlineResponse20013ErrorItems**](.md) | Lists errors that occurred during theme target update. Note that errors are correlated to theme target update requests by the &#x60;index&#x60; field. This field corresponds to the order of the theme target in the request. For example, if &#x60;index&#x60; is set to &#x60;2&#x60;, the values correlate to the third theme target object in the request array. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

