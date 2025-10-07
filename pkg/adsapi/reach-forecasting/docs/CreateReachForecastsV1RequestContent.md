# CreateReachForecastsV1RequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonTargets** | [**[]PlanningTargetV1**](PlanningTargetV1.md) | The list of common targets to be applied to all reach forecasts to be created in this request. The common targets will be added into each forecast specific targets to obtain the final targeting. Targets of the same targetType and of the same negative boolean are combined using OR statements. Targets of different types are associated using AND statements. Note: Audience targets of the same group (as specified by groupId) are linked via an OR linkage. Audience target groups are linked via an AND linkage. Audience targets without groupId are consider as different group, thus are also linked via an AND linkage. | [optional] [default to null]
**ReachForecasts** | [**[]CreateReachForecastsV1RequestElement**](CreateReachForecastsV1RequestElement.md) | A list of reach forecast to be created. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

