# UpdateBillingProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingProfileId** | **string** | Billing profile identifier which will be used to identify a billing profile. This identifier will NOT affect billing process. | [default to null]
**Address** | [***BaseBillingProfileAddress**](BaseBillingProfile_address.md) |  | [default to null]
**Agreements** | [**[]BaseBillingProfileAgreements**](BaseBillingProfile_agreements.md) | List of tax-type and corresponding details. | [optional] [default to null]
**BillingProfileName** | **string** | Name of a billing profile. This name will only be used to identify the billing profile and will not be used for billing. This is for the client to organize the billing profiles and doesn&#x27;t affect the billing process. | [default to null]
**DisplayInfo** | [***BaseBillingProfileDisplayInfo**](BaseBillingProfile_displayInfo.md) |  | [optional] [default to null]
**HoldingCompany** | **string** | Type of holding company for agency billing profile. Agency Holding Companies are conglomerate entities that own multiple smaller advertising agencies. There are six holding companies, WPP, Omnicom, Publicis, Interpublic, Dentsu, and Havas. A holding company is provided if the billing party is an agency. If your agency is not associated with any of the above, choose NO_HOLDING_COMPANY Check following table for supported values: &lt;br/&gt;&lt;br/&gt;&lt;table border&#x3D;1&gt;&lt;caption&gt; **Supported Holding Company** &lt;/caption&gt;&lt;tr&gt;    &lt;th&gt;Holding Company&lt;/th&gt;    &lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;NO_HOLDING_COMPANY&lt;/td&gt;    &lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;WPP&lt;/td&gt;    &lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;OMNICOM&lt;/td&gt;    &lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;PUBLICIS_GROUPE&lt;/td&gt;    &lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;INTERPUBLIC&lt;/td&gt;    &lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;DENTSU&lt;/td&gt;    &lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;HAVAS&lt;/td&gt;    &lt;/tr&gt;&lt;/table&gt; | [optional] [default to null]
**IsDefault** | **bool** | Attribute to indicate if a billing profile is default or not under that global account. Once marked as default, for new countries this billing profile will get automatically applied, if all information provided are sufficient | [optional] [default to false]
**PurchaseOrderNumber** | **string** | Number to track spend against the budgeted amounts. | [default to null]
**Taxes** | [**[]BaseBillingProfileTaxes**](BaseBillingProfile_taxes.md) | List of tax-type and values. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

