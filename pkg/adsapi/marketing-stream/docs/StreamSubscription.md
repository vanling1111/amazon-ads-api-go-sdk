# StreamSubscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | [**time.Time**](time.Time.md) | ISO8601 Timestamp | [default to null]
**DataSetId** | **string** | Identifier of data set, callers can be subscribed to. Please refer to https://advertising.amazon.com/API/docs/en-us/amazon-marketing-stream/data-guide for the list of all data sets. | [default to null]
**Destination** | [***Destination**](Destination.md) |  | [optional] [default to null]
**DestinationArn** | **string** | AWS ARN of the destination endpoint associated with the subscription. Supported destination types: - SQS | [optional] [default to null]
**Notes** | **string** | Additional details associated with the subscription | [optional] [default to null]
**Status** | [***SubscriptionEntityStatus**](SubscriptionEntityStatus.md) |  | [default to null]
**SubscriptionId** | **string** | Unique subscription identifier | [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) | ISO8601 Timestamp | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

