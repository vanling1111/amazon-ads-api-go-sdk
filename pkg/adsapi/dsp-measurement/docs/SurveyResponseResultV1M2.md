# SurveyResponseResultV1M2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdExposedGroupResponseRate** | **float64** | The percent of people in ad exposed group choosing this response. | [optional] [default to null]
**ControlGroupResponseRate** | **float64** | The percent of people in control group choosing this response. | [optional] [default to null]
**IsQualifyingResponse** | **bool** | Is the response a qualifying response. Used in calculating Brand Lift. | [optional] [default to null]
**MarginOfError** | **float64** | The percentage of margin of error for this response. | [optional] [default to null]
**QuestionAsset** | [***AssetV1M2**](AssetV1M2.md) |  | [optional] [default to null]
**QuestionObjective** | [***SurveyQuestionObjectiveTypeV1**](SurveyQuestionObjectiveTypeV1.md) |  | [optional] [default to null]
**QuestionResponse** | [***SurveyQuestionResponseV1M2**](SurveyQuestionResponseV1M2.md) |  | [optional] [default to null]
**QuestionSequence** | **float64** | Sequence number of the question in the Survey. | [optional] [default to null]
**QuestionText** | **string** | Text of the Survey question. | [optional] [default to null]
**ResponseRate** | **float64** | The percentage of people choosing this response. | [optional] [default to null]
**SegmentType** | **string** | The segment type to which this response data belongs to. | [optional] [default to null]
**SegmentValue** | **string** | The segment value to which this response data belongs to. Would be corresponding to the above segmentType field. | [optional] [default to null]
**StatisticalSignificance** | **float64** | The significance percentage for the response data in this segment. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

