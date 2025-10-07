# CreateStreamSubscriptionRequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientRequestToken** | **string** | Unique value supplied by the caller used to track identical API requests. Should request be re-tried, the caller should supply the same value. We recommend using GUID. | [default to null]
**DataSetId** | **string** | Identifier of data set, callers can be subscribed to. Please refer to https://advertising.amazon.com/API/docs/en-us/amazon-marketing-stream/data-guide for the list of all data sets. | [default to null]
**Destination** | [***Destination**](Destination.md) |  | [optional] [default to null]
**DestinationArn** | **string** | AWS ARN of the destination endpoint associated with the subscription. Supported destination types: - SQS | [optional] [default to null]
**Notes** | **string** | Additional details associated with the subscription | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

