# UpdatePostRequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Caption** | **string** | Caption for a post. | [optional] [default to null]
**Medias** | [**[]Media**](Media.md) | A list of medias for a post. | [default to null]
**Products** | **[]string** | A list of product identifiers. | [default to null]
**ProfileId** | **string** | Identifier for a profile. | [default to null]
**ScheduledLiveDate** | [**time.Time**](time.Time.md) | A date and time for when to publish a post. The value is in ISO8601 date-time format (UTC). For example, 2020-08-16T18:20:00.000Z. | [optional] [default to null]
**ScheduledWithdrawalDate** | [**time.Time**](time.Time.md) | A date and time for when to unpublish a post. The value is in ISO8601 date-time format (UTC). For example, 2020-08-16T18:20:00.000Z. | [optional] [default to null]
**Version** | **float64** | Version of a post. Used to ensure that post writes are consistent. Calls can only update the latest version of a post. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

