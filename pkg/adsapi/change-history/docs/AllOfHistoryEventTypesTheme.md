# AllOfHistoryEventTypesTheme

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTypeIds** | **[]string** | max of 10 event types. IDs here belong to the EventType. For example, if requesting CAMPAGIN as the eventType, these are campaignIds. | [optional] [default to null]
**Filters** | **[]string** | | Filter | Entity Types | ||-| | BUDGET_AMOUNT | CAMPAIGN | | IN_BUDGET | CAMPAIGN | | STATUS | CAMPAIGN, AD_GROUP, AD, KEYWORD, PRODUCT_TARGET, NEGATIVE_KEYWORD | | END_DATE | CAMPAIGN | | START_DATE | CAMPAIGN | | PLACEMENT_GROUP | CAMPAIGN| | SMART_BIDDING_STRATEGY | CAMPAIGN | | BID_AMOUNT | AD_GROUP | | NAME | CAMPAIGN, AD_GROUP | | [optional] [default to null]
**Parents** | [**[]HistoryEventTypeParents**](HistoryEventType_parents.md) | maximum of 10 parents | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

