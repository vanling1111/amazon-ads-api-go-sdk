# BaseBillingProfileAgreements

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consent** | **bool** | Consent on the provided agreement document. | [default to false]
**DocumentName** | **string** | Agreement document name against which consent needs to be provided. The content of the agreement can be checked by providing agreement name in &#x60;POST /billingProfileAgreementContents/{billingProfileAgreementContentId}&#x60; API. Check following table for supported values: &lt;br/&gt;&lt;br/&gt;&lt;table border&#x3D;1&gt;&lt;caption&gt; **Supported Agreement Document Names** &lt;/caption&gt;&lt;tr&gt;    &lt;th&gt;Agreement document name&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;BILLING_PROFILE_TAX_AGREEMENT&lt;/td&gt;    &lt;td&gt;Global tax agreement. Advertiser needs to consent to this if they provide taxes in the billing profile&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;BILLING_PROFILE_CANADIAN_NON_RESIDENCY_AGREEMENT&lt;/td&gt;    &lt;td&gt;If advetiser is non-resident of canada and wants to advertise in canada without providing canada GST, they need to consent to this agreement&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;BILLING_PROFILE_DSA_AGREEMENT&lt;/td&gt;&lt;td&gt;Relevant for Authors, If you choose to display your BILLING_NAME or any name other than “Author” then you must provide consent for BILLING_PROFILE_DSA_AGREEMENT.&lt;/td&gt;&lt;/tr&gt; &lt;/table&gt; | [default to null]
**Locale** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

