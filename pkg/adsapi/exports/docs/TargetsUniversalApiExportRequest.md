# TargetsUniversalApiExportRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NegativeFilter** | **[]bool** | Filters the targets returned in export to negative or positive targets. In case the filter is not provided, it returns both negative and positive targets. | [optional] [default to ["false","true"]]
**TargetLevelFilter** | **[]string** | Filters the targets returned in export only to selected levels. In case the filter is not provided, it returns both &#x60;CAMPAIGN&#x60; and &#x60;AD_GROUP&#x60; level targets. | [optional] [default to ["AD_GROUP","CAMPAIGN"]]
**TargetTypeFilter** | **[]string** | Filters the targets returned in exports only to selected types. In case the filter is not provided, it returns targets with all target types. Target types are only supported by certain ad products - for instance, &#x60;THEME&#x60; targets are not available in &#x60;SPONSORED_BRANDS&#x60;. Please reference https://advertising.amazon.com/API/docs/en-us/reference/common-models/targets for more details. | [optional] [default to ["AUDIENCE","AUTO","KEYWORD","PRODUCT","PRODUCT_AUDIENCE","PRODUCT_CATEGORY","PRODUCT_CATEGORY_AUDIENCE","THEME"]]
**AdProductFilter** | **[]string** | Filters the entities returned in export only to selected ad products. In case the filter is not provided, it returns entities from all ad products. | [optional] [default to ["SPONSORED_BRANDS","SPONSORED_DISPLAY","SPONSORED_PRODUCTS"]]
**StateFilter** | **[]string** | Filters the entities returned in export only to selected states. In case the filter is not provided, it returns only &#x60;ENABLED&#x60; or &#x60;PAUSED&#x60; entities. | [optional] [default to ["ENABLED","PAUSED"]]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

