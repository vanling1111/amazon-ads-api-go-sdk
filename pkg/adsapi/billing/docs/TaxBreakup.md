# TaxBreakup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerJurisdiction** | **string** | Tax jurisdiction of issuer (Amazon billing entity)  | [default to null]
**IssuerTaxInformation** | [***IssuerTaxRegistrationInfo**](Issuer Tax Registration Info.md) |  | [default to null]
**PayerJurisdiction** | **string** | Tax jurisdiction of payer (billed customer)  | [optional] [default to null]
**PayerTaxInformation** | [***PayerTaxRegistrationInfo**](Payer Tax Registration Info.md) |  | [default to null]
**TaxAmount** | [***CurrencyAmount**](currencyAmount.md) |  | [default to null]
**TaxName** | **string** |  | [default to null]
**TaxRate** | **float64** |  | [default to null]
**TaxedJurisdictionName** | **string** | Tax jurisdiction for which tax applies, this can be at the country, state or local level.  | [default to null]
**ThirdPartyTaxInformation** | [***ThirdPartyTaxRegistrationInfo**](Third Party Tax Registration Info.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

