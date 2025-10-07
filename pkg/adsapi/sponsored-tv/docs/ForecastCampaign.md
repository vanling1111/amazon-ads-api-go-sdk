# ForecastCampaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetSettings** | [***BudgetSettings**](BudgetSettings.md) |  | [default to null]
**CampaignId** | **string** | The identifier of the forecast campaign. | [optional] [default to null]
**CostType** | **string** | Support for the responses with newly added values should be ensured for using them at the time of creation or updates.  | Type | Description | | --- | --- | | \&quot;CPM\&quot; | [Default] The performance of the campaign will be measured per 1000 impressions | | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The yyyy-MM-dd&#x27;T&#x27;hh:mm:ss&#x27;Z&#x27; end date for the campaign in forecast. Must be greater than the value for startDate. If not specified, the campaign has no end date and runs continuously. | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The yyyy-MM-dd&#x27;T&#x27;hh:mm:ss&#x27;Z&#x27; start date for the campaign in forecast. If this field is not set to a value, the current date is used. | [default to null]
**TargetingType** | [***TargetingType**](TargetingType.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

