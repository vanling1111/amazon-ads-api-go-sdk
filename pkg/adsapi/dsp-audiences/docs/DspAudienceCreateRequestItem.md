# DspAudienceCreateRequestItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceType** | **string** | Type of audience to create. | [default to null]
**Country** | **string** | The ISO Alpha-2 code for the country in which the audience will be available during audience discovery and targeting setup.  **Notes:**  - If you are using a Global Advertising account, this field is required. - If you are not using a Global Advertising account, this field is optional. If used, it must be same as the advertiser&#x27;s country.  | [optional] [default to null]
**Description** | **string** | The audience description. | [default to null]
**IdempotencyKey** | **string** | The unique UUID for this requested audience. | [default to null]
**Lookback** | **int32** | The specified time period (in days) to include those who performed the action in the audience. Lookback Constraints Table: Provides available valid values of lookback allowed for given audienceType   | audienceType | lookback range |   |------------------------------|-------|   | PRODUCT_PURCHASES            | 1-365 |   | PRODUCT_VIEWS                |  1-90 |   | PRODUCT_SEARCH               |  1-90 |   | PRODUCT_SIMS                 |  1-90 |   | WHOLE_FOODS_MARKET_PURCHASES | 1-365 |  | [default to null]
**Name** | **string** | The audience name. | [default to null]
**Rules** | [**[]DspAudienceRule**](DspAudienceRule.md) | The set of rules defining an audience; these rules will be ORed. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

