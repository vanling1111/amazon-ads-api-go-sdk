# CreateDspCreativeTestingStudyV1M2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertiserId** | **string** | The associated advertiser identifier. Immutable field. | [optional] [default to null]
**Assets** | [**[]AssetV1M2**](AssetV1M2.md) | A list of assets to be used for the creative testing study as part of either the survey question or the response. In case of API responses, number of assets returned would be limited to 10 even if a creative testing study has more than 10 assets associated with it. | [optional] [default to null]
**AudienceTargetingGroup** | [***AudienceTargetingGroupV1M2**](AudienceTargetingGroupV1M2.md) |  | [optional] [default to null]
**BrandName** | **string** | The study brand name. | [optional] [default to null]
**ProductCategory** | **string** | Optional study product category. | [optional] [default to null]
**Comment** | **string** | The approver&#x27;s comment on why the study is approved/rejected. | [optional] [default to null]
**CreateDate** | [**time.Time**](time.Time.md) | The study creation date in ISO format (YYYY-MM-DDThh:mm:ssTZD). Timezone is UTC. | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The study end date in ISO format (YYYY-MM-DDThh:mm:ssTZD). Timezone is UTC. By default this will be the latest endDate of the associated orders. | [optional] [default to null]
**ExternalReferenceId** | **string** | Optional field. For some vendors, advertisers are required to provide this vendor assigned reference identifier for EXTERNAL_BILLING studies. | [optional] [default to null]
**LastUpdatedDate** | [**time.Time**](time.Time.md) | The study last updated date in ISO format (YYYY-MM-DDThh:mm:ssTZD). Timezone is UTC. | [optional] [default to null]
**Name** | **string** | The study name. | [optional] [default to null]
**ReviewDate** | [**time.Time**](time.Time.md) | The study review date in ISO format (YYYY-MM-DDThh:mm:ssTZD). Timezone is UTC. | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The study start date in ISO format (YYYY-MM-DDThh:mm:ssTZD). Timezone is UTC. By default this will be the earliest startDate of the associated orders. | [optional] [default to null]
**Status** | [***StudyStatusV1**](StudyStatusV1.md) |  | [optional] [default to null]
**StatusReasons** | **[]string** | List of reasons for study status. For example, when study is marked Rejected or Ineligible, this field would be available. | [optional] [default to null]
**StudyResultStatus** | **string** | The status of result of the study. | [optional] [default to null]
**SurveyId** | **string** | The study survey canonical identifier. | [optional] [default to null]
**VendorProductId** | **string** | Associated vendor product canonical identifier. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

