# ConversionDefinitionInputV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversionDefinitionId** | **string** | The id of the ConversionDefinition. | [optional] [default to null]
**ConversionType** | [***ConversionDefinitionTypeV1**](ConversionDefinitionTypeV1.md) |  | [default to null]
**CountingMethod** | [***ConversionDefinitionCountingMethodV1**](ConversionDefinitionCountingMethodV1.md) |  | [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | The timestamp of when the ConversionDefinition was created in ISO format (YYYY-MM-DDThh:mm:ssTZD). | [optional] [default to null]
**LastActivityTime** | [**time.Time**](time.Time.md) | The latest timestamp of when a conversion event for the ConversionDefinition was imported in ISO format (YYYY-MM-DDThh:mm:ssTZD). | [optional] [default to null]
**LastUpdatedTime** | [**time.Time**](time.Time.md) | Date and time last edit was made to conversion settings in ISO format (YYYY-MM-DDThh:mm:ssTZD). | [optional] [default to null]
**Name** | **string** | The name of the ConversionDefinition. | [default to null]
**Source** | [***ConversionDefinitionSourceV1**](ConversionDefinitionSourceV1.md) |  | [default to null]
**SourceType** | [***ConversionDefinitionSourceTypeV1**](ConversionDefinitionSourceTypeV1.md) |  | [default to null]
**Value** | **float64** | The value of the event.&lt;br&gt; When the conversionType of the associated Conversion Definition is OFF_AMAZON_PURCHASES, this represents a monetary value. Must be a minimum of 0 and must not exceed 2 decimal points. If not specified, a default of 0 will be applied.&lt;br&gt; When the conversionType of the associated Conversion Definition is not OFF_AMAZON_PURCHASES, this represents a non-monetary value based on a scale of your choosing. Can be negative and must not exceed 2 decimal points. If not specified on the conversion definition, a default of 1 will be applied. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

