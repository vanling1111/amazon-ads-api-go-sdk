# ListAudiencesRequestBodyV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdType** | **string** |  | [optional] [default to null]
**Countries** | **[]string** | The ISO Alpha-2 country codes to search audiences from. This field must be specified if the advertiser does not have an associated country. Currently, it is only supported to specify a single country per request. | [optional] [default to null]
**Filters** | [**[]AudienceFilterV1**](AudienceFilterV1.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

