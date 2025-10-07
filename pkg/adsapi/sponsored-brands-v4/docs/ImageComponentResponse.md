# ImageComponentResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreModerationStatus** | **string** | The pre moderation status of the component. | [optional] [default to null]
**ComponentType** | **string** | Type of the image component. | [optional] [default to null]
**LandingPage** | [***LandingPage**](LandingPage.md) |  | [optional] [default to null]
**PolicyViolations** | [**[]ImagePolicyViolation**](ImagePolicyViolation.md) | A list of policy violations for the component that were detected during pre moderation. Note that this field is present in the response only when preModerationStatus is set to REJECTED. | [optional] [default to null]
**Id** | **string** | Id of the component. This is the same id sent as part of the request. This can be used to uniquely identify the component. | [optional] [default to null]
**Url** | **string** | Publicly accessible url of the image that got pre moderated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

