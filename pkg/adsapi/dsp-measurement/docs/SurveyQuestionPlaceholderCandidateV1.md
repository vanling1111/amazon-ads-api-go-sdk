# SurveyQuestionPlaceholderCandidateV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCustomValue** | **bool** | Whether custom value is allowed for the placeholder. | [optional] [default to null]
**AllowedValueRanges** | [**[]SurveyQuestionPlaceholderAllowedRangeV1**](SurveyQuestionPlaceholderAllowedRangeV1.md) | Allowed value ranges for placeholder. Only applicable if the valueType is INTEGER. | [optional] [default to null]
**AllowedValues** | [**[]SurveyQuestionPlaceholderAllowedValueV1**](SurveyQuestionPlaceholderAllowedValueV1.md) | Allowed values for placeholder. Will be empty if placeholder is free text field. | [optional] [default to null]
**DefaultValues** | **[]string** | Default values that will be appended to the values list regardless. | [optional] [default to null]
**FieldName** | **string** | The survey question placeholder field name. | [optional] [default to null]
**InferredFields** | **[]string** | Where the placeholder values will be inferred from. | [optional] [default to null]
**MaximumValueLength** | **int32** | The maximum allowed character length for each individual placeholder value. | [optional] [default to null]
**MinimumValueLength** | **int32** | The minimum allowed character length for each individual placeholder value. | [optional] [default to null]
**ValueType** | [***PlaceholderValueTypeV1**](PlaceholderValueTypeV1.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

