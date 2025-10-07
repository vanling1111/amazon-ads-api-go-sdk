# ImagePolicyViolation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageEvidences** | [**[]ImageEvidence**](ImageEvidence.md) | List of evidences for the policy violations detected on the image component. | [optional] [default to null]
**Name** | **string** | A policy violation code. | [optional] [default to null]
**PolicyDescription** | **string** | A human-readable description of the policy. | [optional] [default to null]
**PolicyLinkUrl** | **string** | Address of the policy documentation. Follow the link to learn more about the specified policy. | [optional] [default to null]
**TextEvidences** | [**[]TextEvidence**](TextEvidence.md) | Policy violation on an image can be detected on the ocr detected text on the image as well. This list of text evidences will have the policy violations detected on the text on top of the image. | [optional] [default to null]
**Type_** | **string** | Type of policy violation. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

