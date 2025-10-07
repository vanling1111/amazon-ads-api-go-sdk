# Post

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Caption** | **string** | Caption for a post. | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | Date and time for when the post was created. The value is in ISO8601 date-time format (UTC). For example, 2020-08-16T18:20:00.000Z. | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**IsFlaggedForQuality** | **bool** | Whether the post has quality defects or not. | [optional] [default to null]
**LastModified** | [**time.Time**](time.Time.md) | Date and time for when the post was last modified. The value is in ISO8601 date-time format (UTC). For example, 2020-08-16T18:20:00.000Z. | [optional] [default to null]
**LiveDate** | [**time.Time**](time.Time.md) | Date and time for when the post was published. The value is in ISO8601 date-time format (UTC). For example, 2020-08-16T18:20:00.000Z. | [optional] [default to null]
**MediaUrl** | **string** |  | [optional] [default to null]
**Medias** | [**[]Media**](Media.md) | A list of medias for a post. | [optional] [default to null]
**Metrics** | [***map[string]float64**](map.md) |  | [optional] [default to null]
**Products** | **[]string** | A list of product identifiers. | [optional] [default to null]
**ProfileId** | **string** | Identifier for a profile. | [optional] [default to null]
**PromotionMetadata** | [***PromotionMetadata**](PromotionMetadata.md) |  | [optional] [default to null]
**ScheduledLiveDate** | [**time.Time**](time.Time.md) | Date and time for when the post was unpublished. The value is in ISO8601 date-time format (UTC). For example, 2020-08-16T18:20:00.000Z. | [optional] [default to null]
**ScheduledWithdrawalDate** | [**time.Time**](time.Time.md) | A date and time for when to unpublish a post. The value is in ISO8601 date-time format (UTC). For example, 2020-08-16T18:20:00.000Z. | [optional] [default to null]
**Status** | [***PostStatus**](PostStatus.md) |  | [optional] [default to null]
**StatusMetadata** | [***StatusMetadata**](StatusMetadata.md) |  | [optional] [default to null]
**Version** | **float64** | Version of a post. Used to ensure that post writes are consistent. Calls can only update the latest version of a post. | [optional] [default to null]
**WithdrawnDate** | [**time.Time**](time.Time.md) | Date and time for when the post was unpublished. The value is in ISO8601 date-time format (UTC). For example, 2020-08-16T18:20:00.000Z. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

