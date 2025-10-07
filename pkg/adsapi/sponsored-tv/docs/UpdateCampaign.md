# UpdateCampaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetSettings** | [***BudgetSettings**](BudgetSettings.md) |  | [optional] [default to null]
**CampaignId** | **string** | The identifier of the Campaign. | [default to null]
**CostType** | **string** | Cost type of the Campaign. Determines how the campaign will bid and charge. Note that new values can be added to this list in the future. Support for the responses with newly addded values should be ensured for using them at the time of creation or updates.  | Type | Description | | --- | --- | | \&quot;CPM\&quot; | [Default] The performance of the campaign will be measured per 1000 impresions | | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | endDate is optional. If endDate is specified, startDate must be specified as well. | [optional] [default to null]
**Name** | **string** | The name of the Campaign.  Note: Names including single quotes must be escaped. For example, to update the name to &#x60;Campaign &#x27;A&#x27;&#x60;, the name field should be as follows: &#x60;\&quot;name\&quot;: \&quot;Campaign &#x27;\\&#x27;&#x27;A&#x27;\\&#x27;&#x27;\&quot;&#x60; | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | startDate can only be changed if the current startDate is in the future. | [optional] [default to null]
**State** | [***CreateOrUpdateEntityState**](CreateOrUpdateEntityState.md) |  | [optional] [default to null]
**Tags** | [***map[string]string**](map.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

