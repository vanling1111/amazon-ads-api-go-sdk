# SbForecastingRequestCampaignObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | **float64** | The amount of the budget. | [default to null]
**BudgetType** | **string** | Budget can be set to DAILY or LIFETIME.   |BudgetType|Description| |-----------|-----------| |DAILY| The amount that you&#x27;re willing to spend on this campaign each day. If the campaign spends less than your daily budget, the unspent amount can be used to increase your daily budget on the other days of the calendar month.| |LIFETIME| The total amount that you are willing to spend on this campaign.|    | [default to null]
**ForecastType** | **string** | The forecast type. can be set to WEEKLY or MONTHLY.   **If have not set the forecastType during campaign creation then use WEEKLY as goal value.**  | [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The YYYY-MM-DD start date for the campaign. If this field is not set to a value, the current date is used. | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The YYYY-MM-DD end date for the campaign. Must be greater than the value for &#x60;startDate&#x60;. If not specified, the campaign has no end date and runs continuously. | [optional] [default to null]
**Goal** | **string** | Goal will allow you to set goal type to help drive your campaign performance.   **If have not set the goal during campaign creation then use PAGE_VISIT as goal type.**    The goal type of the campaign. Initial launch only supports PAGE_VISIT.   BRAND_IMPRESSION_SHARE - This goal is a PREVIEW ONLY and cannot be set currently. It will allow you grown your brand impression share on top of search placement.   PAGE_VISIT - This goal drives traffic to your landing and detail pages through all placements.   ACQUIRE_NEW_CUSTOMERS - This property is a PREVIEW ONLY and cannot be used as part of a request or response. This goal drives new customer acquisition for your brands.   AD_VIEWS - This property is a PREVIEW ONLY and cannot be used as part of a request or response. This goal maximizes view for your ads.    | [optional] [default to null]
**AdGroups** | [**[]SbForecastingAdGroup**](SBForecastingAdGroup.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

