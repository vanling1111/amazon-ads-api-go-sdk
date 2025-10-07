# VendorProductPolicyV1M1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BenchMarkCategoryRequired** | **bool** | Whether or not the benchMark category is required for measurement setup. | [optional] [default to null]
**CustomQuestionAllowed** | **bool** | Whether custom survey questions are allowed. | [optional] [default to null]
**ExternalReferenceIdRequired** | **bool** | Whether or not the vendor assigned external reference identifier is required for measurement setup. | [optional] [default to null]
**LeadTime** | **int32** | Days required for measurement configuration. It is recommended that the startDate of the campaign has sufficient padding to accommodate this lead time, but measurement can begin after the campaign start date in some cases. | [optional] [default to null]
**MaximumOrders** | **int32** | The maximum number of order allowed for the product. | [optional] [default to null]
**MaximumPeerNames** | **int32** | The maximum number of peer names required for the product. | [optional] [default to null]
**MaximumStudyLength** | **int32** | The maximum required length/duration of the study in days. | [optional] [default to null]
**MaximumSurveyQuestions** | **int32** | The maximum number of survey questions required for the product. | [optional] [default to null]
**MinimumOrders** | **int32** | The maximum number of orders required for the product. | [optional] [default to null]
**MinimumPeerNames** | **int32** | The minimum number of peer names required for the product. | [optional] [default to null]
**MinimumStudyLength** | **int32** | The minimum required length/duration of the study in days. | [optional] [default to null]
**MinimumSurveyQuestions** | **int32** | The minimum number of survey questions required for the product. | [optional] [default to null]
**RequiredQuestionCategories** | [**[]SurveyQuestionCategoryRequirementV1M1**](SurveyQuestionCategoryRequirementV1M1.md) | The requirements for survey question categories. | [optional] [default to null]
**RequiredQuestionObjectives** | [**[]SurveyQuestionObjectiveTypeV1M1**](SurveyQuestionObjectiveTypeV1M1.md) | The required question objectives that need to be included as part of the survey. | [optional] [default to null]
**SupportedGoals** | [**[]MeasurementGoalV1**](MeasurementGoalV1.md) |  | [optional] [default to null]
**SupportedMarketplaces** | [**[]MeasurementMarketplaceV1**](MeasurementMarketplaceV1.md) |  | [optional] [default to null]
**SupportedVerbs** | **[]string** | List of supported verbs that can be used in survey questions. | [optional] [default to null]
**VendorApprovalRequired** | **bool** | Whether or not the vendor requires an additional sign off process to fully qualify for study. | [optional] [default to null]
**VendorProductId** | **string** | vendor product canonical identifier. | [optional] [default to null]
**VerbRequired** | **bool** | Whether or not a verb is required for measurement setup. It will be used in applicable survey questions to construct the question text. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

