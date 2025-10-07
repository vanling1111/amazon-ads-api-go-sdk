# CreateCampaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetSettings** | [***BudgetSettings**](BudgetSettings.md) |  | [default to null]
**CostType** | **string** | Cost type of the Campaign. Determines how the Campaign will bid and charge. Note that new values can be added to this list in the future. Support for the responses with newly addded values should be ensured for using them at the time of creation or updates. Following are the list of acceptable values  | Type | Description | | --- | --- | | \&quot;CPM\&quot; | [Default] The performance of the Campaign will be measured per 1000 impresions | | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | endDate is optional. If endDate is specified, startDate must be specified as well. | [optional] [default to null]
**FullFunnelCampaignId** | **string** | full funnel campaign id for child campaign | [optional] [default to null]
**Name** | **string** | The name of the Campaign.  Note: Names including single quotes must be escaped. For example, to create &#x60;Campaign &#x27;A&#x27;&#x60;, the name field should be as follows: &#x60;\&quot;name\&quot;: \&quot;Campaign &#x27;\\&#x27;&#x27;A&#x27;\\&#x27;&#x27;\&quot;&#x60; | [default to null]
**StartDate** | [**time.Time**](time.Time.md) | startDate is optional. If startDate is not specified, current date will be used. | [optional] [default to null]
**State** | [***CreateOrUpdateEntityState**](CreateOrUpdateEntityState.md) |  | [default to null]
**Tags** | [***map[string]string**](map.md) |  | [optional] [default to null]
**TargetingType** | [***TargetingType**](TargetingType.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

