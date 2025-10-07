# BaseUniversalApiExportRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdProductFilter** | **[]string** | Filters the entities returned in export only to selected ad products. In case the filter is not provided, it returns entities from all ad products. | [optional] [default to ["SPONSORED_BRANDS","SPONSORED_DISPLAY","SPONSORED_PRODUCTS"]]
**StateFilter** | **[]string** | Filters the entities returned in export only to selected states. In case the filter is not provided, it returns only &#x60;ENABLED&#x60; or &#x60;PAUSED&#x60; entities. | [optional] [default to ["ENABLED","PAUSED"]]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

