# InvoiceSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountDue** | [***CurrencyAmount**](currencyAmount.md) |  | [default to null]
**BillingAggregation** | [***BillingAggregation**](billingAggregation.md) |  | [optional] [default to null]
**DownloadableDocuments** | [**[]DocumentType**](documentType.md) | List of downloadable documents associated with this invoice and accessible from the advertising console.  | [optional] [default to null]
**DueDate** | **string** |  | [optional] [default to null]
**Fees** | [**[]Fee**](fee.md) | Regulatory Advertising Fees.  | [optional] [default to null]
**FromDate** | **string** |  | [default to null]
**Id** | **string** |  | [default to null]
**InvoiceDate** | **string** |  | [default to null]
**PaymentMethod** | [***PaymentMethod**](paymentMethod.md) |  | [optional] [default to null]
**PaymentTermsDays** | **int32** |  | [optional] [default to null]
**PaymentTermsType** | **string** |  | [optional] [default to null]
**PurchaseOrderNumber** | **string** |  | [optional] [default to null]
**RemainingAmountDue** | [***CurrencyAmount**](currencyAmount.md) |  | [default to null]
**RemainingFees** | [**[]Fee**](fee.md) | Remaining Regulatory Advertising Fees.  | [optional] [default to null]
**RemainingTaxAmountDue** | [***CurrencyAmount**](currencyAmount.md) |  | [optional] [default to null]
**Status** | [***InvoiceStatus**](invoiceStatus.md) |  | [default to null]
**TaxAmountDue** | [***CurrencyAmount**](currencyAmount.md) |  | [optional] [default to null]
**TaxRate** | **float64** |  | [optional] [default to null]
**ToDate** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

