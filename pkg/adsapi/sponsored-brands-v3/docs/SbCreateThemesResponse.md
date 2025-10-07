# SbCreateThemesResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | [**[]SbCreateThemesResponseSuccess**](SBCreateThemesResponse_success.md) | Lists the successfully created theme targets. Note that theme targets in the response are correlated to theme targets in the request using the &#x60;index&#x60; field. For example, if &#x60;index&#x60; is set to &#x60;2&#x60;, the values correlate to the third theme target object in the request. | [optional] [default to null]
**Error_** | [**[]AllOfSbCreateThemesResponseErrorItems**](.md) | Lists errors that occurred during theme target creation. Note that errors are correlated to theme target create requests by the &#x60;index&#x60; field. This field corresponds to the order of the target object in the request. For example, if &#x60;index&#x60; is set to &#x60;3&#x60;, an error occurred during creation of the fourth theme target in the request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

