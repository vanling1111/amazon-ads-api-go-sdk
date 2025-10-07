# DateComponentResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentType** | **string** | Type of the date component. | [optional] [default to null]
**EndDate** | **string** | End date of the component. | [optional] [default to null]
**Id** | **string** | Id of the component. This is the same id sent as part of the request. This can be used to uniquely identify the component. | [optional] [default to null]
**PolicyViolations** | [**[]DatePolicyViolation**](DatePolicyViolation.md) | A list of policy violations for the component that were detected during pre moderation. Note that this field is present in the response only when preModerationStatus is set to REJECTED. | [optional] [default to null]
**PreModerationStatus** | **string** | The pre-moderation status of the component. | [optional] [default to null]
**StartDate** | **string** | Start date of the component. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

