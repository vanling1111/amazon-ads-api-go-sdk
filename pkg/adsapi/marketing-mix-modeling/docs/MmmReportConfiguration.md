# MmmReportConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandGroupId** | **string** | Identifies the brand group being reported on. | [default to null]
**GeoDimension** | **string** | Geographic granularity of the report. |Value|Description| |---|---| |COUNTRY|Aggregate metrics by country.| |POSTAL_CODE|Aggregate metrics by postal code, e.g. ZIP Code. Valid only in select countries.| |DMA|Aggregate metrics by DMA® (Designated Market Area) region. Valid only in the US.| | [default to null]
**MetricsType** | **string** | The type of metrics to include in the report. |Value|Description| |---|---| |MEDIA_ONLY|Core advertising metrics only.| |MEDIA_AND_SALES|Advertising and retail metrics.| | [default to null]
**TimeUnit** | **string** | Time granularity of the report. |Value|Description| |---|---| |DAILY|Aggregate metrics with daily granularity.| |WEEKLY|Aggregate metrics with weekly granularity.| | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

