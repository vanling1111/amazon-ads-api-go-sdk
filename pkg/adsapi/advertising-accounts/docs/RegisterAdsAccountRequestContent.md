# RegisterAdsAccountRequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** | Account names are typically the name of the company or brand being advertised. We recommend that you avoid using personal details such as first name, last name, phone number, social security number, credit card or other personally identifiable information. | [optional] [default to null]
**Associations** | [**[]Association**](Association.md) | Associations you would like to link to this advertising account, could be Amazon Vendor, Seller, or just a regular business | [optional] [default to null]
**CountryCodes** | **[]string** | The countries that you want this account to operate in. | [optional] [default to null]
**TermsToken** | **string** | We recommend you do not provide this field since we can determine if the customer has accepted the terms for you. An obfuscated identifier of the termsToken, which is activated when an advertisers accepts the Amazon Ads Agreement in relation to the ads account being register. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

