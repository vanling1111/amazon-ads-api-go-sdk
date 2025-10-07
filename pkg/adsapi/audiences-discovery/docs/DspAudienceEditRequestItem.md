# DspAudienceEditRequestItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceId** | **string** | The audience identifier of the audience to be actioned. | [default to null]
**AudienceType** | [***AudienceType**](AudienceType.md) |  | [default to null]
**Description** | **string** | The audience description. | [optional] [default to null]
**IdempotencyKey** | **string** | unique request token for this request. | [default to null]
**Lookback** | **int32** | The specified time period (in days) to include those who performed the action in the audience. Lookback Constraints Table: Provides available valid values of lookback allowed for given audienceType | audienceType | lookback range | |------------------------------|-------| | PRODUCT_PURCHASES            | 1-365 | | PRODUCT_VIEWS                |  1-90 | | PRODUCT_SEARCH               |  1-90 | | PRODUCT_SIMS                 |  1-90 | | [optional] [default to null]
**Name** | **string** | The audience name. | [optional] [default to null]
**Rules** | [**[]DspAudienceRule**](DSPAudienceRule.md) | Set of rules to define an audience, these rules will be ORed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

