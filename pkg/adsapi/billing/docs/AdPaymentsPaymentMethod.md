# AdPaymentsPaymentMethod

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | **string** | The processor used to execute payments on the card, such as Visa, MasterCard, etc. Only valued for credit card payment methods. | [optional] [default to null]
**CountryCode** | [***AdPaymentsCountryCode**](AdPaymentsCountryCode.md) |  | [optional] [default to null]
**ExpiryDetails** | [***AdPaymentsExpiryDetails**](AdPaymentsExpiryDetails.md) |  | [optional] [default to null]
**ForeignExchange** | [***AdPaymentsForeignExchange**](AdPaymentsForeignExchange.md) |  | [optional] [default to null]
**InstrumentId** | **string** | Identifies a credit card or direct debit payment method. | [optional] [default to null]
**Priority** | **int32** | A numerical priority assigned to each payment method within a given profile, dictating the sequential order in which these payment methods are utilized during payment execution. Lower numerical values indicate higher priority, with the option possessing the lowest value being selected first. | [default to null]
**SellerAccountId** | **string** | The seller account ID associated to this payment method, only valued for seller payable payment methods. | [optional] [default to null]
**Tail** | **string** | The last four digits of a credit card or bank account number. | [optional] [default to null]
**Type_** | [***AdPaymentsPaymentMethodType**](AdPaymentsPaymentMethodType.md) |  | [default to null]
**VendorCode** | **string** | The vendor code associated to this payment method, only valued for deduct from payment payment methods. | [optional] [default to null]
**VendorCodeBalance** | [***AdPaymentsCurrencyAmount**](AdPaymentsCurrencyAmount.md) |  | [optional] [default to null]
**VendorCodeName** | **string** | The name of the vendor code this payment method represents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

