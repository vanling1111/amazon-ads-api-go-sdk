# CreateBillingStatementRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCodes** | [**[]CountryCode**](CountryCode.md) | List of countries for which billing statements are to be fetched. If no country codes are passed in the request, the statement will be generated for all the advertising countries by default. | [optional] [default to null]
**EndDate** | **string** | End date of the invoice summary period for a report in the format YYYY-MM-DD. | [default to null]
**Format** | **string** | Format of the file, such as, for billing statements, etc. | [optional] [default to FORMAT.CSV]
**Locale** | **string** | Preferred locale can be chosen among the list of valid language codes. Check the table below for supported language code. &lt;br/&gt;&lt;br/&gt;&lt;table border&#x3D;1&gt;&lt;caption&gt; **Supported Locales Table** &lt;/caption&gt;&lt;tr&gt;    &lt;th&gt;Locale code&lt;/th&gt;    &lt;th&gt;Language&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;en_US&lt;/td&gt;    &lt;td&gt;English&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;ja_JP&lt;/td&gt;    &lt;td&gt;Japanese&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;ar_SA&lt;/td&gt;    &lt;td&gt;Arabic&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;de_DE&lt;/td&gt;    &lt;td&gt;German&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;fr_FR&lt;/td&gt;    &lt;td&gt;French&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;it_IT&lt;/td&gt;    &lt;td&gt;Italian&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;es_MX&lt;/td&gt;    &lt;td&gt;Spanish&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;nl_NL&lt;/td&gt;    &lt;td&gt;Dutch&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;sv_SE&lt;/td&gt;    &lt;td&gt;Swedish&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;pl_PL&lt;/td&gt;    &lt;td&gt;Polish&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;tr_TR&lt;/td&gt;    &lt;td&gt;Turkish&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;    &lt;td&gt;pt_BR&lt;/td&gt;    &lt;td&gt;Portuguese&lt;/td&gt;&lt;/tr&gt; &lt;/table&gt; | [default to null]
**StartDate** | **string** | Start date of the invoice summary period for a report in the format YYYY-MM-DD. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

