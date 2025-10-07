# SbCreateTargetsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTargetSuccessResults** | [**[]SbCreateTargetsResponseCreateTargetSuccessResults**](SBCreateTargetsResponse_createTargetSuccessResults.md) | Lists the successfully created targets. Note that targets in the response are correlated to targets in the request using the &#x60;targetRequestIndex&#x60; field. For example, if &#x60;targetRequestIndex&#x60; is set to &#x60;2&#x60;, the values correlate to the third target object in the request. | [optional] [default to null]
**CreateTargetErrorResults** | [**[]AllOfSbCreateTargetsResponseCreateTargetErrorResultsItems**](.md) | Lists errors that occured during target creation. Note that errors are correlated to target create requests by the &#x60;targetRequestIndex&#x60; field. This field corresponds to the order of the target object in the request. For example, if &#x60;targetRequestIndex&#x60; is set to &#x60;3&#x60;, an error occured during creation of the fourth target in the request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

