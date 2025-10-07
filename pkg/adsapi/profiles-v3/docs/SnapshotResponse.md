# SnapshotResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | **string** | The identifier of the snapshot that was requested. | [optional] [default to null]
**RecordType** | **string** | The record type of the snapshot file. | [optional] [default to null]
**Status** | **string** | The status of the generation of the snapshot. | [optional] [default to null]
**StatusDetails** | **string** | Optional description of the status. | [optional] [default to null]
**Location** | **string** | The URI for the snapshot. It&#x27;s only available if status is SUCCESS. | [optional] [default to null]
**FileSize** | **float64** | The size of the snapshot file in bytes. It&#x27;s only available if status is SUCCESS. | [optional] [default to null]
**Expiration** | **float64** | The epoch time for expiration of the snapshot file and each snapshot file will be expired in 30 mins after generated. It&#x27;s only available if status is SUCCESS. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

