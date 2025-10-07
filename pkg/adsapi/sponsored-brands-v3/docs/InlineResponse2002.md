# InlineResponse2002

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsinList** | **[]string** | An array of ASINs. Note that this field is present only if there were no errors during the request. If there were errors, the &#x60;code&#x60; field is the enumerated error, and the &#x60;details&#x60; field contains a human-readable description of the error. | [optional] [default to null]
**Code** | **string** | The enumerated response code. | Code | Description | |------|--------| |SUCCESS| The request was successful. The &#x60;asinList&#x60; field includes all available ASINs.| |INVALID_ARGUMENT| The request was not successful because the address was not for a valid landing page.| |BAD_GATEWAY| The request failed because the landing page at the specified address did not have any ASINs.| | [optional] [default to null]
**Details** | **string** | A human-readable description of the &#x60;code&#x60; field. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

