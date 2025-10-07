# BaseBillingProfileDisplayInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Relevant for only authors. Required if author chooses displayNameField as DISPLAY_NAME.The value inside this field will be shown publicly. | [optional] [default to null]
**DisplayNameField** | **string** | Determines the name field that will be displayed publicly.For non-authors, the default and only allowed value for this is BILLING_NAME. For author accounts, If BILLING_NAME is chosen then the value inside address.billingName will bedisplayed publicly. If DISPLAY_NAME is chosen then the value inside displayInfo.displayName will beshown publicly. You may choose to display “Author” as your displayName or you may provide an alternatename to display for your ads. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

