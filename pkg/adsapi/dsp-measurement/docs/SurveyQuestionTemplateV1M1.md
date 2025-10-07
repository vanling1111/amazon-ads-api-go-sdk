# SurveyQuestionTemplateV1M1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [***SurveyQuestionCategoryV1M1**](SurveyQuestionCategoryV1M1.md) |  | [optional] [default to null]
**GridQuestionResponse** | [***SurveyQuestionGridQuestionResponseV1M1**](SurveyQuestionGridQuestionResponseV1M1.md) |  | [optional] [default to null]
**Id** | **string** | The survey question template canonical Id. | [optional] [default to null]
**Locale** | [***MeasurementLocaleV1**](MeasurementLocaleV1.md) |  | [optional] [default to null]
**MaximumQualifyingResponses** | **int32** | The maximum number of qualifying responses allowed for the question. This will be available if the qualifying responses are not pre-defined/inferred. | [optional] [default to null]
**MaximumQuestionResponses** | **int32** | The maximum number of responses allowed for the question. This will be available if the question responses are not pre-defined/inferred. | [optional] [default to null]
**MinimumQualifyingResponses** | **int32** | The minimum number of qualifying responses required for the question. This will be available if the qualifying responses are not pre-defined/inferred. | [optional] [default to null]
**MinimumQuestionResponses** | **int32** | The minimum number of responses required for the question. This will be available if the question responses are not pre-defined/inferred. | [optional] [default to null]
**ObjectiveType** | [***SurveyQuestionObjectiveTypeV1M1**](SurveyQuestionObjectiveTypeV1M1.md) |  | [optional] [default to null]
**PlaceholderCandidates** | [**[]SurveyQuestionPlaceholderCandidateV1**](SurveyQuestionPlaceholderCandidateV1.md) |  | [optional] [default to null]
**Priority** | **int32** | The priority of the question. If present this will determine the ordering of questions in a survey. The check will be enforced when a survey is created/updated. Lower number indicates higher priority. | [optional] [default to null]
**QualifyingResponses** | **[]string** | The pre-defined qualifying survey question responses with placeholders, this will help to define which responses will be counted as positive ones in the study report. | [optional] [default to null]
**QuestionResponses** | **[]string** | The pre-defined survey question responses with placeholders. | [optional] [default to null]
**QuestionText** | **string** | The survey question text with placeholders. | [optional] [default to null]
**SubCategory** | [***SurveyQuestionSubCategoryV1M1**](SurveyQuestionSubCategoryV1M1.md) |  | [optional] [default to null]
**Type_** | [***SurveyQuestionTypeV1M1**](SurveyQuestionTypeV1M1.md) |  | [optional] [default to null]
**VendorProductId** | **string** | The associated vendor product id. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

