# RuleCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | **string** | The name of the metric. Supported rule metrics and corresponding supported comparisonOperators: |      MetricName      |ComparisonOperator  |Description| |------------------|--------------------|-------------------| |COST_PER_THOUSAND_VIEWABLE_IMPRESSIONS     |              LESS_THAN_OR_EQUAL_TO             |Maximize viewable impressions while cost per 1000 views less than or equal to &#x60;threshold&#x60;| |COST_PER_CLICK    |              LESS_THAN_OR_EQUAL_TO            |Maximize page visits while cost per click less than or equal to &#x60;threshold&#x60;| |COST_PER_ORDER    |              LESS_THAN_OR_EQUAL_TO            |Maximize viewable impressions/page visits/conversion while cost per order less than or equal to &#x60;threshold&#x60;| | [default to null]
**ComparisonOperator** | **string** | The comparison operator. | [default to null]
**Threshold** | **float64** | The value of the threshold associated with the metric. The threshold values has defined minimums depending on the metric names in the following table: |                  MetricName            | Minimum of &#x60;threshold&#x60; Value  | |----------------------------------------|-----------------------------------| |COST_PER_THOUSAND_VIEWABLE_IMPRESSIONS  | 1                                 | |COST_PER_CLICK                          | 0.5                               | |COST_PER_ORDER                          | 5                                 | | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

