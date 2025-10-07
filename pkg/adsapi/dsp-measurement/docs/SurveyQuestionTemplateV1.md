# SurveyQuestionTemplateV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The survey question template canonical Id. | [optional] [default to null]
**Locale** | [***MeasurementLocaleV1**](MeasurementLocaleV1.md) |  | [optional] [default to null]
**MaximumQualifyingResponses** | **int32** | The maximum number of qualifying responses allowed for the question. This will be available if the qualifying responses are not pre-defined/inferred. | [optional] [default to null]
**MaximumQuestionResponses** | **int32** | The maximum number of responses allowed for the question. This will be available if the question responses are not pre-defined/inferred. | [optional] [default to null]
**MinimumQualifyingResponses** | **int32** | The minimum number of qualifying responses required for the question. This will be available if the qualifying responses are not pre-defined/inferred. | [optional] [default to null]
**MinimumQuestionResponses** | **int32** | The minimum number of responses required for the question. This will be available if the question responses are not pre-defined/inferred. | [optional] [default to null]
**ObjectiveType** | [***SurveyQuestionObjectiveTypeV1**](SurveyQuestionObjectiveTypeV1.md) |  | [optional] [default to null]
**PlaceholderCandidates** | [**[]SurveyQuestionPlaceholderCandidateV1**](SurveyQuestionPlaceholderCandidateV1.md) |  | [optional] [default to null]
**QualifyingResponses** | **[]string** | The pre-defined qualifying survey question responses with placeholders, this will help to define which responses will be counted as positive ones in the study report. | [optional] [default to null]
**QuestionResponses** | **[]string** | The pre-defined survey question responses with placeholders. | [optional] [default to null]
**QuestionText** | **string** | The survey question text with placeholders. | [optional] [default to null]
**Type_** | [***SurveyQuestionTypeV1**](SurveyQuestionTypeV1.md) |  | [optional] [default to null]
**VendorProductId** | **string** | The associated vendor product id. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

