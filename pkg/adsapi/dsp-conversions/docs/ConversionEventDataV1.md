# ConversionEventDataV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientDedupeId** | **string** | An identifier chosen by the advertiser to represent a user event. This parameter is used for deduplication across all conversion definitions (events) and all sources (Amazon Ad tag, Conversions API, Mobile Measurement Partner). In case of multiple events containing the same cientDedupeId, only the first will be kept. Common values used for deduplication are transaction ID or order ID. Do not use user identifiers (hashed or unhashed) for this purpose. | [optional] [default to null]
**ConversionDefinitionId** | **string** | The id of the associated ConversionDefinition. | [default to null]
**CountryCode** | **string** | The country where the event originates from. e.g. US&lt;br&gt; This value is based on [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes), two-letter country codes defined in ISO 3166-1, part of the ISO 3166 standard[1] published by the International Organization for Standardization (ISO), to represent countries, dependent territories, and special areas of geographical interest. | [default to null]
**CurrencyCode** | [***ConversionDefinitionCurrencyCodeV1**](ConversionDefinitionCurrencyCodeV1.md) |  | [optional] [default to null]
**DataProcessingOptions** | **string** | A flag for signaling how an event shall be processed. Events marked for limited data use will not be processed. | [optional] [default to null]
**MatchKeys** | [**[]ConversionMatchKeyV1**](ConversionMatchKeyV1.md) | Array representing the user and device identifier types/values to be used for attribution to traffic events. Match key value must be normalized and hashed, except for MAID which should not be hashed. ADID, IDFA, or FIREADID can be passed into the MAID field for mobile identifiers.  All match keys provided must follow the [formatting guidelines](https://advertising.amazon.com/dsp/help/ss/en/audiences#GCCXMZYCK4RXWS6C). Only SHA-256 is supported. | [default to null]
**Name** | **string** | The name of the imported event. | [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The timestamp when the event occurred in ISO format (YYYY-MM-DDThh:mm:ssTZD). The event&#x27;s timestamp must be no more than 7 days old at the time of import. | [default to null]
**UnitsSold** | **int64** | The number of items purchased. Only applicable for OFF_AMAZON_PURCHASES conversion type. If not provided on the conversion event, a default of 1 will be applied. | [optional] [default to null]
**Value** | **float64** | The value of the event.&lt;br&gt; When the conversionType of the associated Conversion Definition is OFF_AMAZON_PURCHASES, this represents a monetary value. Must be a minimum of 0 and must not exceed 2 decimal points. If not provided, the static value provided on the conversion definition will be used.&lt;br&gt; When the conversionType of the associated Conversion Definition is not OFF_AMAZON_PURCHASES, this represents a non-monetary value based on a scale of your choosing. Can be negative and must not exceed 2 decimal points. If not provided, the static value provided on the conversion definition will be used. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

